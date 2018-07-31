package mysql

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DBGroup struct {
	defaultConfigKey string
	config           map[string]DBConfig
	dbGroup          map[string]*DB
	cache            Cache
}

func NewDBGroupCache(defaultConfigName string, cache Cache) (group *DBGroup) {
	group = &DBGroup{}
	group.defaultConfigKey = defaultConfigName
	group.config = map[string]DBConfig{}
	group.dbGroup = map[string]*DB{}
	group.cache = cache
	return
}
func NewDBGroup(defaultConfigName string) (group *DBGroup) {
	group = &DBGroup{}
	group.defaultConfigKey = defaultConfigName
	group.config = map[string]DBConfig{}
	group.dbGroup = map[string]*DB{}
	return
}
func (g *DBGroup) RegistGroup(cfg map[string]DBConfig) (err error) {
	g.config = cfg
	for name, config := range g.config {
		if config.Cache == nil {
			config.Cache = g.cache
		}
		g.Regist(name, config)
		if err != nil {
			return
		}
	}
	return
}
func (g *DBGroup) Regist(name string, cfg DBConfig) (err error) {
	var db DB
	if cfg.Cache == nil {
		cfg.Cache = g.cache
	}
	db, err = NewDB(cfg)
	if err != nil {
		return
	}
	g.config[name] = cfg
	g.dbGroup[name] = &db
	return
}
func (g *DBGroup) DB(name ...string) (db *DB) {
	key := ""
	if len(name) == 0 {
		key = g.defaultConfigKey
	} else {
		key = name[0]
	}
	db, _ = g.dbGroup[key]
	return
}

type DB struct {
	Config   DBConfig
	ConnPool *sql.DB
	DSN      string
}

func NewDB(config DBConfig) (db DB, err error) {
	db = DB{}
	err = db.init(config)
	return
}
func (db *DB) init(config DBConfig) (err error) {
	db.Config = config
	db.DSN = db.getDSN()
	db.ConnPool, err = db.getDB()
	return
}

func (db *DB) getDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%dms&readTimeout=%dms&writeTimeout=%dms&charset=%s&collation=%s",
		url.QueryEscape(db.Config.Username),
		db.Config.Password,
		url.QueryEscape(db.Config.Host),
		db.Config.Port,
		url.QueryEscape(db.Config.Database),
		db.Config.Timeout,
		db.Config.ReadTimeout,
		db.Config.WriteTimeout,
		url.QueryEscape(db.Config.Charset),
		url.QueryEscape(db.Config.Collate))
}
func (db *DB) getDB() (connPool *sql.DB, err error) {
	connPool, err = sql.Open("mysql", db.getDSN())
	if err != nil {
		return
	}
	connPool.SetMaxOpenConns(db.Config.SetMaxOpenConns)
	connPool.SetMaxIdleConns(db.Config.SetMaxIdleConns)
	err = connPool.Ping()
	return
}
func (db *DB) AR() (ar *ActiveRecord) {
	ar = new(ActiveRecord)
	ar.Reset()
	ar.tablePrefix = db.Config.TablePrefix
	ar.tablePrefixSqlIdentifier = db.Config.TablePrefixSqlIdentifier
	return
}
func (db *DB) Begin(config DBConfig) (tx *sql.Tx, err error) {
	return db.ConnPool.Begin()
}
func (db *DB) ExecTx(ar *ActiveRecord, tx *sql.Tx) (rs *ResultSet, err error) {
	sqlStr := ar.SQL()
	var stmt *sql.Stmt
	var result sql.Result
	rs = new(ResultSet)
	stmt, err = tx.Prepare(sqlStr)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err = stmt.Exec(ar.values...)
	if err != nil {
		return
	}
	rs.RowsAffected, err = result.RowsAffected()
	rs.LastInsertId, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}
func (db *DB) Exec(ar *ActiveRecord) (rs *ResultSet, err error) {
	sqlStr := ar.SQL()
	var stmt *sql.Stmt
	var result sql.Result
	rs = new(ResultSet)
	stmt, err = db.ConnPool.Prepare(sqlStr)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err = stmt.Exec(ar.values...)
	if err != nil {
		return
	}
	rs.RowsAffected, err = result.RowsAffected()
	rs.LastInsertId, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}
func (db *DB) Query(ar *ActiveRecord) (rs *ResultSet, err error) {
	var results []map[string][]byte
	if ar.cacheKey != "" {
		var data []byte
		data, err = db.Config.Cache.Get(ar.cacheKey)
		if err == nil {
			d := gob.NewDecoder(bytes.NewReader(data))
			err = d.Decode(&results)
			if err != nil {
				return
			}
		}
	}
	if results == nil || len(results) == 0 {
		sqlStr := ar.SQL()
		var stmt *sql.Stmt
		stmt, err = db.ConnPool.Prepare(sqlStr)
		if err != nil {
			return
		}
		defer stmt.Close()
		var rows *sql.Rows
		rows, err = stmt.Query(ar.values...)
		if err != nil {
			return
		}
		defer rows.Close()
		cols := []string{}
		cols, err = rows.Columns()
		if err != nil {
			return
		}
		vals := make([][]byte, len(cols))
		scans := make([]interface{}, len(cols))
		for i := range vals {
			scans[i] = &vals[i]
		}
		results = []map[string][]byte{}
		for rows.Next() {
			err = rows.Scan(scans...)
			if err != nil {
				return
			}
			row := make(map[string][]byte)
			for k, v := range vals {
				key := cols[k]
				row[key] = v
			}
			results = append(results, row)
		}
		if ar.cacheKey != "" {
			b := new(bytes.Buffer)
			e := gob.NewEncoder(b)
			err = e.Encode(results)
			if err != nil {
				return
			}
			err = db.Config.Cache.Set(ar.cacheKey, b.Bytes(), ar.cacheSeconds)
			if err != nil {
				return
			}
		}
	}
	rs = new(ResultSet)
	rs.Init(&results)
	return
}

type DBConfig struct {
	Charset                  string
	Collate                  string
	Database                 string
	Host                     string
	Port                     int
	Username                 string
	Password                 string
	TablePrefix              string
	TablePrefixSqlIdentifier string
	Timeout                  int
	ReadTimeout              int
	WriteTimeout             int
	SetMaxIdleConns          int
	SetMaxOpenConns          int
	Cache                    Cache
}

func NewDBConfigWith(host string, port int, dbName, user, pass string) (cfg DBConfig) {
	cfg = NewDBConfig()
	cfg.Host = host
	cfg.Port = port
	cfg.Username = user
	cfg.Password = pass
	cfg.Database = dbName
	return
}
func NewDBConfig() DBConfig {
	return DBConfig{
		Charset:                  "utf8",
		Collate:                  "utf8_general_ci",
		Database:                 "test",
		Host:                     "127.0.0.1",
		Port:                     3306,
		Username:                 "root",
		Password:                 "",
		TablePrefix:              "",
		TablePrefixSqlIdentifier: "",
		Timeout:                  3000,
		ReadTimeout:              5000,
		WriteTimeout:             5000,
		SetMaxOpenConns:          500,
		SetMaxIdleConns:          50,
	}
}

type ActiveRecord struct {
	arSelect                 [][]interface{}
	arFrom                   []string
	arJoin                   [][]string
	arWhere                  [][]interface{}
	arGroupBy                []string
	arHaving                 [][]interface{}
	arOrderBy                map[string]string
	arLimit                  string
	arSet                    map[string][]interface{}
	arUpdateBatch            []interface{}
	arInsert                 map[string]interface{}
	arInsertBatch            []map[string]interface{}
	asTable                  map[string]bool
	values                   []interface{}
	sqlType                  string
	currentSQL               string
	tablePrefix              string
	tablePrefixSqlIdentifier string
	cacheKey                 string
	cacheSeconds             uint
}

func (ar *ActiveRecord) Cache(key string, seconds uint) *ActiveRecord {
	ar.cacheKey = key
	ar.cacheSeconds = seconds
	return ar
}
func (ar *ActiveRecord) getValues() []interface{} {
	return ar.values
}
func (ar *ActiveRecord) Reset() {
	ar.arSelect = [][]interface{}{}
	ar.arFrom = []string{}
	ar.arJoin = [][]string{}
	ar.arWhere = [][]interface{}{}
	ar.arGroupBy = []string{}
	ar.arHaving = [][]interface{}{}
	ar.arOrderBy = map[string]string{}
	ar.arLimit = ""
	ar.arSet = map[string][]interface{}{}
	ar.arUpdateBatch = []interface{}{}
	ar.arInsert = map[string]interface{}{}
	ar.arInsertBatch = []map[string]interface{}{}
	ar.asTable = map[string]bool{}
	ar.values = []interface{}{}
	ar.sqlType = "select"
	ar.currentSQL = ""
	ar.cacheKey = ""
	ar.cacheSeconds = 0
}

func (ar *ActiveRecord) Select(columns string) *ActiveRecord {
	return ar._select(columns, true)
}
func (ar *ActiveRecord) SelectNoWrap(columns string) *ActiveRecord {
	return ar._select(columns, false)
}

func (ar *ActiveRecord) _select(columns string, wrap bool) *ActiveRecord {
	for _, column := range strings.Split(columns, ",") {
		ar.arSelect = append(ar.arSelect, []interface{}{column, wrap})
	}
	return ar
}
func (ar *ActiveRecord) From(from string) *ActiveRecord {
	ar.FromAs(from, "")
	return ar
}
func (ar *ActiveRecord) FromAs(from, as string) *ActiveRecord {
	ar.arFrom = []string{from, as}
	if as != "" {
		ar.asTable[as] = true
	}
	return ar
}

func (ar *ActiveRecord) Join(table, as, on, type_ string) *ActiveRecord {
	ar.arJoin = append(ar.arJoin, []string{table, as, on, type_})
	return ar
}
func (ar *ActiveRecord) Where(where map[string]interface{}) *ActiveRecord {
	if len(where) > 0 {
		ar.WhereWrap(where, "AND", "")
	}
	return ar
}
func (ar *ActiveRecord) WhereWrap(where map[string]interface{}, leftWrap, rightWrap string) *ActiveRecord {
	if len(where) > 0 {
		ar.arWhere = append(ar.arWhere, []interface{}{where, leftWrap, rightWrap, len(ar.arWhere)})
	}
	return ar
}
func (ar *ActiveRecord) GroupBy(column string) *ActiveRecord {
	for _, column_ := range strings.Split(column, ",") {
		ar.arGroupBy = append(ar.arGroupBy, strings.TrimSpace(column_))
	}
	return ar
}
func (ar *ActiveRecord) Having(having string) *ActiveRecord {
	ar.HavingWrap(having, "AND", "")
	return ar
}
func (ar *ActiveRecord) HavingWrap(having, leftWrap, rightWrap string) *ActiveRecord {
	ar.arHaving = append(ar.arHaving, []interface{}{having, leftWrap, rightWrap, len(ar.arHaving)})
	return ar
}

func (ar *ActiveRecord) OrderBy(column, type_ string) *ActiveRecord {
	ar.arOrderBy[column] = type_
	return ar
}

//Limit Limit(offset,count) or Limit(count)
func (ar *ActiveRecord) Limit(limit ...int) *ActiveRecord {
	if len(limit) == 1 {
		ar.arLimit = fmt.Sprintf("%d", limit[0])

	} else if len(limit) == 2 {
		ar.arLimit = fmt.Sprintf("%d,%d", limit[0], limit[1])
	} else {
		ar.arLimit = ""
	}
	return ar
}

func (ar *ActiveRecord) Insert(table string, data map[string]interface{}) *ActiveRecord {
	ar.sqlType = "insert"
	ar.arInsert = data
	ar.From(table)
	return ar
}
func (ar *ActiveRecord) Replace(table string, data map[string]interface{}) *ActiveRecord {
	ar.sqlType = "replace"
	ar.arInsert = data
	ar.From(table)
	return ar
}

func (ar *ActiveRecord) InsertBatch(table string, data []map[string]interface{}) *ActiveRecord {
	ar.sqlType = "insertBatch"
	ar.arInsertBatch = data
	ar.From(table)
	return ar
}
func (ar *ActiveRecord) ReplaceBatch(table string, data []map[string]interface{}) *ActiveRecord {
	ar.InsertBatch(table, data)
	ar.sqlType = "replaceBatch"
	return ar
}

func (ar *ActiveRecord) Delete(table string, where map[string]interface{}) *ActiveRecord {
	ar.From(table)
	ar.Where(where)
	ar.sqlType = "delete"
	return ar
}
func (ar *ActiveRecord) Update(table string, data, where map[string]interface{}) *ActiveRecord {
	ar.From(table)
	ar.Where(where)
	for k, v := range data {
		if isBool(v) {
			value := 0
			if v.(bool) {
				value = 1
			}
			ar.Set(k, value)
		} else if v == nil {
			ar.SetNoWrap(k, "NULL")
		} else {
			ar.Set(k, v)
		}
	}
	return ar
}
func (ar *ActiveRecord) UpdateBatch(table string, values []map[string]interface{}, whereColumn []string) *ActiveRecord {
	ar.From(table)
	ar.sqlType = "updateBatch"
	ar.arUpdateBatch = []interface{}{values, whereColumn}
	if len(values) > 0 {
		for _, whereCol := range whereColumn {
			ids := []interface{}{}
			for _, val := range values {
				ids = append(ids, val[whereCol])
			}
			ar.Where(map[string]interface{}{whereCol: ids})
		}
	}
	return ar
}

func (ar *ActiveRecord) Set(column string, value interface{}) *ActiveRecord {
	ar.sqlType = "update"
	ar.arSet[column] = []interface{}{value, true}
	return ar
}
func (ar *ActiveRecord) SetNoWrap(column string, value interface{}) *ActiveRecord {
	ar.sqlType = "update"
	ar.arSet[column] = []interface{}{value, false}
	return ar
}
func (ar *ActiveRecord) Wrap(v string) string {
	columns := strings.Split(v, ".")
	if len(columns) == 2 {
		return ar.protectIdentifier(ar.checkPrefix(columns[0])) + "." + ar.checkPrefix(columns[1])
	}
	return ar.protectIdentifier(ar.checkPrefix(columns[0]))
}
func (ar *ActiveRecord) Raw(sql string, values ...interface{}) *ActiveRecord {
	ar.currentSQL = sql
	if len(values) > 0 {
		ar.values = append(ar.values, values...)
	}
	return ar
}
func (ar *ActiveRecord) Values() []interface{} {
	return ar.values
}
func (ar *ActiveRecord) SQL() string {
	if ar.currentSQL != "" {
		return ar.currentSQL
	}
	switch ar.sqlType {
	case "select":
		ar.currentSQL = ar.getSelectSQL()
	case "update":
		ar.currentSQL = ar.getUpdateSQL()
	case "updateBatch":
		ar.currentSQL = ar.getUpdateBatchSQL()
	case "insert":
		ar.currentSQL = ar.getInsertSQL()
	case "insertBatch":
		ar.currentSQL = ar.getInsertBatchSQL()
	case "replace":
		ar.currentSQL = ar.getReplaceSQL()
	case "replaceBatch":
		ar.currentSQL = ar.getReplaceBatchSQL()
	case "delete":
		ar.currentSQL = ar.getDeleteSQL()
	}
	return ar.currentSQL
}
func (ar *ActiveRecord) getUpdateSQL() string {
	SQL := []string{"UPDATE "}
	SQL = append(SQL, ar.getFrom())
	SQL = append(SQL, "\nSET")
	SQL = append(SQL, ar.compileSet())
	SQL = append(SQL, ar.getWhere())
	orderBy := strings.TrimSpace(ar.compileOrderBy())
	if orderBy != "" {
		SQL = append(SQL, fmt.Sprintf("\nORDER BY %s", orderBy))
	}
	SQL = append(SQL, ar.getLimit())
	return strings.Join(SQL, " ")
}

func (ar *ActiveRecord) getUpdateBatchSQL() string {
	SQL := []string{"UPDATE "}
	SQL = append(SQL, ar.getFrom())
	SQL = append(SQL, "\nSET")
	SQL = append(SQL, ar.compileUpdateBatch())
	SQL = append(SQL, ar.getWhere())
	return strings.Join(SQL, " ")
}
func (ar *ActiveRecord) getInsertSQL() string {
	SQL := []string{"INSERT INTO "}
	SQL = append(SQL, ar.getFrom())
	SQL = append(SQL, ar.compileInsert())
	return strings.Join(SQL, " ")
}
func (ar *ActiveRecord) getReplaceSQL() string {
	SQL := []string{"REPLACE INTO "}
	SQL = append(SQL, ar.getFrom())
	SQL = append(SQL, ar.compileInsert())
	return strings.Join(SQL, " ")
}
func (ar *ActiveRecord) getInsertBatchSQL() string {
	SQL := []string{"INSERT INTO "}
	SQL = append(SQL, ar.getFrom())
	SQL = append(SQL, ar.compileInsertBatch())
	return strings.Join(SQL, " ")
}
func (ar *ActiveRecord) getReplaceBatchSQL() string {
	SQL := []string{"REPLACE INTO "}
	SQL = append(SQL, ar.getFrom())
	SQL = append(SQL, ar.compileInsertBatch())
	return strings.Join(SQL, " ")
}
func (ar *ActiveRecord) getDeleteSQL() string {
	SQL := []string{"DELETE FROM "}
	SQL = append(SQL, ar.getFrom())
	SQL = append(SQL, ar.getWhere())
	orderBy := strings.TrimSpace(ar.compileOrderBy())
	if orderBy != "" {
		SQL = append(SQL, fmt.Sprintf("\nORDER BY %s", orderBy))
	}
	SQL = append(SQL, ar.getLimit())
	return strings.Join(SQL, " ")
}
func (ar *ActiveRecord) getSelectSQL() string {
	from := ar.getFrom()
	where := ar.getWhere()
	having := ""
	for _, w := range ar.arHaving {
		having += ar.compileWhere(w[0], w[1].(string), w[2].(string), w[3].(int))
	}
	having = strings.TrimSpace(having)
	if having != "" {
		having = fmt.Sprintf("\nHAVING %s", having)
	}
	groupBy := strings.TrimSpace(ar.compileGroupBy())
	if groupBy != "" {
		groupBy = fmt.Sprintf("\nGROUP BY %s", groupBy)
	}
	orderBy := strings.TrimSpace(ar.compileOrderBy())
	if orderBy != "" {
		orderBy = fmt.Sprintf("\nORDER BY %s", orderBy)
	}
	limit := ar.getLimit()
	Select := ar.compileSelect()
	return fmt.Sprintf("SELECT %s \nFROM %s %s %s %s %s %s", Select, from, where, groupBy, having, orderBy, limit)
}
func (ar *ActiveRecord) compileUpdateBatch() string {
	_values, _index := ar.arUpdateBatch[0], ar.arUpdateBatch[1]
	index := _index.([]string)
	values := _values.([]map[string]interface{})
	columns := []string{}
	for k := range values[0] {
		_continue := false
		for _, v1 := range index {
			if k == v1 {
				_continue = true
				break
			}
		}
		if _continue {
			continue
		}
		columns = append(columns, k)
	}
	str := ""
	for _, column := range columns {
		_column := column
		realColumnArr := strings.Split(column, " ")
		if len(realColumnArr) == 2 {
			_column = realColumnArr[0]
		}
		str += fmt.Sprintf("%s = CASE \n", ar.protectIdentifier(_column))
		for _, row := range values {
			_when := []string{}
			for _, col := range index {
				_when = append(_when, fmt.Sprintf("%s = ?", ar.protectIdentifier(col)))
				ar.values = append(ar.values, row[col])
			}
			_whenStr := strings.Join(_when, " AND ")
			if len(realColumnArr) == 2 {
				str += fmt.Sprintf("WHEN %s THEN %s %s ? \n", _whenStr, ar.protectIdentifier(_column), realColumnArr[1])
			} else {
				str += fmt.Sprintf("WHEN %s THEN ? \n", _whenStr)
			}
			ar.values = append(ar.values, row[column])
		}
		str += fmt.Sprintf("ELSE %s END,", ar.protectIdentifier(_column))
	}
	return strings.TrimRight(str, " ,")
}
func isArray(v interface{}) bool {
	if v == nil {
		return false
	}
	return reflect.TypeOf(v).Kind() == reflect.Slice || reflect.TypeOf(v).Kind() == reflect.Array
}
func isBool(v interface{}) bool {
	if v == nil {
		return false
	}
	return reflect.TypeOf(v).Kind() == reflect.Bool
}
func MapKey(v map[string]interface{}) string {
	for k := range v {
		return k
	}
	return ""
}
func MapCurrent(v map[string]interface{}) interface{} {
	for _, val := range v {
		return val
	}
	return ""
}
func (ar *ActiveRecord) compileInsert() string {
	var columns = []string{}
	var values = []string{}
	for k, v := range ar.arInsert {
		columns = append(columns, ar.protectIdentifier(k))
		values = append(values, "?")
		ar.values = append(ar.values, v)
	}
	if len(columns) > 0 {
		return fmt.Sprintf("(%s) \nVALUES (%s)", strings.Join(columns, ","), strings.Join(values, ","))
	}
	return ""
}
func (ar *ActiveRecord) compileInsertBatch() string {
	var columns []string
	var values []string
	for col := range ar.arInsertBatch[0] {
		columns = append(columns, ar.protectIdentifier(col))
	}
	for _, row := range ar.arInsertBatch {

		_values := []string{}
		for _, col := range columns {
			_values = append(_values, "?")
			ar.values = append(ar.values, row[strings.Trim(col, "`")])
		}
		values = append(values, fmt.Sprintf("(%s)", strings.Join(_values, ",")))
	}
	return fmt.Sprintf("(%s) \nVALUES %s", strings.Join(columns, ","), strings.Join(values, ","))
}
func (ar *ActiveRecord) compileSet() string {
	set := []string{}
	for key, _value := range ar.arSet {
		value, wrap := _value[0], _value[1]
		_column := key
		op := ""
		realColumnArr := strings.Split(key, " ")
		if len(realColumnArr) == 2 {
			_column = realColumnArr[0]
			op = realColumnArr[1]
		}
		if wrap.(bool) {
			if op != "" {
				set = append(set, fmt.Sprintf("%s = %s %s ?", ar.protectIdentifier(_column), ar.protectIdentifier(_column), op))
			} else {
				set = append(set, fmt.Sprintf("%s = ?", ar.protectIdentifier(_column)))
			}
			ar.values = append(ar.values, value)
		} else {
			set = append(set, fmt.Sprintf("%s = %s", ar.protectIdentifier(_column), value))
		}
	}
	return strings.Join(set, ",")
}
func (ar *ActiveRecord) compileGroupBy() string {
	groupBy := []string{}
	for _, key := range ar.arGroupBy {
		_key := strings.Split(key, ".")
		if len(_key) == 2 {
			groupBy = append(groupBy, fmt.Sprintf("%s.%s", ar.protectIdentifier(ar.checkPrefix(_key[0])), ar.protectIdentifier(_key[1])))
		} else {
			groupBy = append(groupBy, fmt.Sprintf("%s", ar.protectIdentifier(_key[0])))
		}
	}
	return strings.Join(groupBy, ",")
}

func (ar *ActiveRecord) compileOrderBy() string {
	orderBy := []string{}
	for key, Type := range ar.arOrderBy {
		Type = strings.ToUpper(Type)
		_key := strings.Split(key, ".")
		if len(_key) == 2 {
			orderBy = append(orderBy, fmt.Sprintf("%s.%s %s", ar.protectIdentifier(ar.checkPrefix(_key[0])), ar.protectIdentifier(_key[1]), Type))

		} else {
			orderBy = append(orderBy, fmt.Sprintf("%s %s", ar.protectIdentifier(_key[0]), Type))
		}
	}
	return strings.Join(orderBy, ",")
}
func (ar *ActiveRecord) compileWhere(where0 interface{}, leftWrap, rightWrap string, index int) string {

	_where := []string{}
	if index == 0 {
		str := strings.ToUpper(strings.TrimSpace(leftWrap))
		if strings.Contains(str, "AND") || strings.Contains(str, "OR") {
			leftWrap = ""
		}
	}
	if reflect.TypeOf(where0).Kind() == reflect.String {
		return fmt.Sprintf(" %s %s %s ", leftWrap, where0, rightWrap)
	}
	where := where0.(map[string]interface{})
	for key, value := range where {
		k := ""
		k = strings.TrimSpace(key)
		_key := strings.SplitN(k, " ", 2)
		op := ""
		if len(_key) == 2 {
			op = _key[1]
		}
		keys := strings.Split(_key[0], ".")
		if len(keys) == 2 {
			k = ar.protectIdentifier(ar.checkPrefix(keys[0])) + "." + ar.protectIdentifier(keys[1])
		} else {
			k = ar.protectIdentifier(keys[0])
		}

		if isArray(value) {
			if op != "" {
				op += " IN"
			} else {
				op = "IN"
			}
			op = strings.ToUpper(op)
			l := reflect.ValueOf(value).Len()

			_v := []string{}
			for i := 0; i < l; i++ {
				_v = append(_v, "?")
			}
			_where = append(_where, fmt.Sprintf("%s %s (%s)", k, op, strings.Join(_v, ",")))
			for _, v := range *ar.interface2Slice(value) {
				ar.values = append(ar.values, v)
			}
		} else if isBool(value) {
			if op == "" {
				op = "="
			}
			op = strings.ToUpper(op)
			_v := 0
			if value.(bool) {
				_v = 1
			}
			_where = append(_where, fmt.Sprintf("%s %s ?", k, op))
			ar.values = append(ar.values, _v)
		} else if value == nil {
			if op == "" {
				op = "IS"
			}
			op = strings.ToUpper(op)
			_where = append(_where, fmt.Sprintf("%s %s NULL", k, op))
		} else {
			if op == "" {
				op = "="
			}
			op = strings.ToUpper(op)
			_where = append(_where, fmt.Sprintf("%s %s ?", k, op))
			ar.values = append(ar.values, value)
		}
	}
	return fmt.Sprintf(" %s %s %s ", leftWrap, strings.Join(_where, " AND "), rightWrap)
}
func (ar *ActiveRecord) interface2Slice(data interface{}) (arr *[]interface{}) {
	arr = &[]interface{}{}
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Array || val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			e := val.Index(i)
			*arr = append(*arr, e.Interface())
		}
	}
	return
}
func (ar *ActiveRecord) compileSelect() string {
	selects := ar.arSelect
	columns := []string{}
	if len(selects) == 0 {
		selects = append(selects, []interface{}{"*", true})
	}
	for _, v := range selects {
		protect := v[1].(bool)
		value := strings.TrimSpace(v[0].(string))
		if value != "*" {
			info := strings.Split(value, ".")
			if len(info) == 2 {
				_v := ar.checkPrefix(info[0])
				if protect {
					info[0] = ar.protectIdentifier(_v)
					info[1] = ar.protectIdentifier(info[1])
				} else {
					info[0] = _v
				}
				value = strings.Join(info, ".")
			} else if protect {
				value = ar.protectIdentifier(value)
			}
		}
		columns = append(columns, value)
	}
	return strings.Join(columns, ",")
}

func (ar *ActiveRecord) checkPrefix(v string) string {
	if strings.Contains(v, "(") || strings.Contains(v, ")") || strings.TrimSpace(v) == "*" {
		return v
	}
	if ar.tablePrefix != "" && !strings.Contains(v, ar.tablePrefix) {
		if _, exists := ar.asTable[v]; !exists {
			return ar.tablePrefix + v
		}
	}
	return v
}
func (ar *ActiveRecord) protectIdentifier(v string) string {
	if strings.Contains(v, "(") || strings.Contains(v, ")") || strings.TrimSpace(v) == "*" {
		return v
	}
	values := strings.Split(v, " ")
	if len(values) == 3 && strings.ToLower(values[1]) == "as" {
		return fmt.Sprintf("`%s` AS `%s`", values[0], values[2])
	}
	return fmt.Sprintf("`%s`", v)
}
func (ar *ActiveRecord) compileFrom(from, as string) string {
	if as != "" {
		ar.asTable[as] = true
		as = " AS " + ar.protectIdentifier(as) + " "
	}
	return ar.protectIdentifier(ar.checkPrefix(from)) + as
}
func (ar *ActiveRecord) compileJoin(table, as, on, type_ string) string {
	table_ := ""
	if as != "" {
		ar.asTable[table] = true
		table_ = ar.protectIdentifier(ar.checkPrefix(table)) + " AS " + ar.protectIdentifier(as)
	} else {
		table_ = ar.protectIdentifier(ar.checkPrefix(table))
	}
	a := strings.Split(on, "=")
	if len(a) == 2 {
		left := strings.Split(a[0], ".")
		right := strings.Split(a[1], ".")
		left[0] = ar.protectIdentifier(ar.checkPrefix(left[0]))
		left[1] = ar.protectIdentifier(left[1])
		right[0] = ar.protectIdentifier(ar.checkPrefix(right[0]))
		right[1] = ar.protectIdentifier(right[1])
		on = strings.Join(left, ".") + "=" + strings.Join(right, ".")
	}
	return fmt.Sprintf(" %s JOIN %s ON %s ", type_, table_, on)
}

func (ar *ActiveRecord) getFrom() string {
	table := ar.compileFrom(ar.arFrom[0], ar.arFrom[1])
	for _, v := range ar.arJoin {
		table += ar.compileJoin(v[0], v[1], v[2], v[3])
	}
	return table
}
func (ar *ActiveRecord) getLimit() string {
	limit := ar.arLimit
	if limit != "" {
		limit = fmt.Sprintf("\nLIMIT %s", limit)
	}
	return limit
}
func (ar *ActiveRecord) getWhere() string {
	where := []string{}
	hasEmptyIn := false

	for _, v := range ar.arWhere {
		for _, value := range v[0].(map[string]interface{}) {
			if isArray(value) && reflect.ValueOf(value).Len() == 0 {
				hasEmptyIn = true
				break
			}
		}
		if hasEmptyIn {
			break
		}
		where = append(where, ar.compileWhere(v[0].(map[string]interface{}), v[1].(string), v[2].(string), v[3].(int)))
	}
	if hasEmptyIn {
		return "WHERE 0"
	}
	allWhere := strings.TrimSpace(strings.Join(where, ""))
	if allWhere != "" {
		allWhere = fmt.Sprintf("\nWHERE %s", allWhere)
	}
	return allWhere
}

type ResultSet struct {
	rawRows      *[]map[string][]byte
	LastInsertId int64
	RowsAffected int64
}

func (rs *ResultSet) Init(rawRows *[]map[string][]byte) {
	if rawRows != nil {
		rs.rawRows = rawRows
	} else {
		rs.rawRows = &([]map[string][]byte{})
	}
}
func (rs *ResultSet) Len() int {
	return len(*rs.rawRows)
}
func (rs *ResultSet) MapRows(keyColumn string) (rowsMap map[string]map[string]string) {
	rowsMap = map[string]map[string]string{}
	for _, row := range *rs.rawRows {
		newRow := map[string]string{}
		for k, v := range row {
			newRow[k] = string(v)
		}
		rowsMap[newRow[keyColumn]] = newRow
	}
	return
}
func (rs *ResultSet) MapStructs(keyColumn string, strucT interface{}) (structsMap map[string]interface{}, err error) {
	structsMap = map[string]interface{}{}
	for _, row := range *rs.rawRows {
		newRow := map[string]string{}
		for k, v := range row {
			newRow[k] = string(v)
		}
		var _struct interface{}
		_struct, err = rs.mapToStruct(newRow, strucT)
		if err != nil {
			return nil, err
		}
		structsMap[newRow[keyColumn]] = _struct
	}
	return
}
func (rs *ResultSet) Rows() (rows []map[string]string) {
	rows = []map[string]string{}
	for _, row := range *rs.rawRows {
		newRow := map[string]string{}
		for k, v := range row {
			newRow[k] = string(v)
		}
		rows = append(rows, newRow)
	}
	return
}
func (rs *ResultSet) Structs(strucT interface{}) (structs []interface{}, err error) {
	structs = []interface{}{}
	for _, row := range *rs.rawRows {
		newRow := map[string]string{}
		for k, v := range row {
			newRow[k] = string(v)
		}
		var _struct interface{}
		_struct, err = rs.mapToStruct(newRow, strucT)
		if err != nil {
			return nil, err
		}
		structs = append(structs, _struct)
	}
	return structs, nil
}
func (rs *ResultSet) Row() (row map[string]string) {
	row = map[string]string{}
	if rs.Len() > 0 {
		row = map[string]string{}
		for k, v := range (*rs.rawRows)[0] {
			row[k] = string(v)
		}
	}
	return
}
func (rs *ResultSet) Struct(strucT interface{}) (Struct interface{}, err error) {
	if rs.Len() > 0 {
		return rs.mapToStruct(rs.Row(), strucT)
	}
	return nil, errors.New("rs is empty")
}
func (rs *ResultSet) Values(column string) (values []string) {
	values = []string{}
	for _, row := range *rs.rawRows {
		values = append(values, string(row[column]))
	}
	return
}
func (rs *ResultSet) MapValues(keyColumn, valueColumn string) (values map[string]string) {
	values = map[string]string{}
	for _, row := range *rs.rawRows {
		values[string(row[keyColumn])] = string(row[valueColumn])
	}
	return
}
func (rs *ResultSet) Value(column string) (value string) {
	row := rs.Row()
	if row != nil {
		value, _ = row[column]
	}
	return
}
func (rs *ResultSet) mapToStruct(mapData map[string]string, Struct interface{}) (struCt interface{}, err error) {
	rv := reflect.New(reflect.TypeOf(Struct)).Elem()
	if reflect.TypeOf(Struct).Kind() != reflect.Struct {
		return nil, errors.New("v must be struct")
	}
	fieldType := rv.Type()
	for i, fieldCount := 0, rv.NumField(); i < fieldCount; i++ {
		fieldVal := rv.Field(i)
		if !fieldVal.CanSet() {
			continue
		}

		structField := fieldType.Field(i)
		structTag := structField.Tag
		name := structTag.Get("column")

		if _, ok := mapData[name]; !ok {
			continue
		}
		switch structField.Type.Kind() {
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint, reflect.Uintptr:
			if val, err := strconv.ParseUint(mapData[name], 10, 64); err == nil {
				fieldVal.SetUint(val)
			}
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
			if val, err := strconv.ParseInt(mapData[name], 10, 64); err == nil {
				fieldVal.SetInt(val)
			}
		case reflect.String:
			fieldVal.SetString(mapData[name])
		case reflect.Bool:
			val := false
			if mapData[name] == "1" {
				val = true
			}
			fieldVal.SetBool(val)
		case reflect.Float32, reflect.Float64:
			if val, err := strconv.ParseFloat(mapData[name], 64); err == nil {
				fieldVal.SetFloat(val)
			}
		case reflect.Struct:
			if structField.Type.Name() == "Time" {
				local, _ := time.LoadLocation("Local")
				val, err := time.ParseInLocation("2006-01-02 15:04:05", mapData[name], local)
				if err == nil {
					fieldVal.Set(reflect.ValueOf(val))
				}
			}
		}
	}
	return rv.Interface(), err
}

type Cache interface {
	Set(key string, val []byte, expire uint) (err error)
	Get(key string) (data []byte, err error)
}
