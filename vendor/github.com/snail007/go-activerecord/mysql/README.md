# go-activerecord for mysql
<pre>
import  github.com/snail007/go-activerecord/mysql

0.Configure 
    var dbCfg = mysql.NewDBConfig()
    dbCfg.Password = "admin"
    db, err := mysql.NewDB(dbCfg)
    if err != nil {
            fmt.Printf("ERR:%s", err)
            return
    }
    dbCfg.xxx,"xxxx" default is below:
    Charset:                  "utf8",
    Collate:                  "utf8_general_ci",
    Database:                 "test",
    Host:                     "127.0.0.1",
    Port:                     3306,
    Username:                 "root",
    Password:                 "",
    TablePrefix:              "",
    TablePrefixSqlIdentifier: "",
    Timeout:                  3000, //ms
    SetMaxOpenConns:          500,
    SetMaxIdleConns:          50,
1.Connect to multilple Database
    group := mysql.NewDBGroup("default")
    group.Regist("default", NewDBConfigWith("127.0.0.1", 3306, "test", "root", "admin"))
    group.Regist("blog", NewDBConfigWith("127.0.0.1", 3306, "blog_db", "root", "admin"))
    group.Regist("www", NewDBConfigWith("127.0.0.1", 3306, "www_db", "root", "admin"))
    //group.DB() equal to group.DB("default")
    db := group.DB("www")
    if db != nil {
        rs, err := db.Query(db.AR().From("test"))
        if err != nil {
            t.Errorf("ERR:%s", err)
        } else {
            fmt.Println(rs.Rows())
        }
    } else {
        fmt.Printf("db group config of name %s not found", "www")
    }
2.Select
    type User{
        ID int `column:"id"`
        Name string `column:"name"`
    }
    rs, err := db.Query(db.AR().
                Select("*").
                From("log").
                Where(map[string]interface{}{
                    "id": 11,
                })
    )
    if err != nil {
		fmt.Printf("ERR:%s", err)
	} else {
		fmt.Println(rs.Rows())
	}
    //struct 
    _user :=User{}
    user,err:=rs.Struct(_user)
    if err != nil {
		fmt.Printf("ERR:%s", err)
	} else {
		fmt.Println(user)
	}
    //structs
    _user :=User{}
    users,err:=rs.Structs(_user)
    if err != nil {
		fmt.Printf("ERR:%s", err)
	} else {
		fmt.Println(users)
	}
    //Map structs
    _user :=User{}
    usersMap,err=rs.MapStructs("id",_user)
    if err != nil {
		fmt.Printf("ERR:%s", err)
	} else {
		fmt.Println(usersMap)
	}
    //db.AR() return a new *mysql.ActiveRecord,you can use it to build you sql.
    //all of db.AR() usage to see <a href="https://github.com/snail007/go-activerecord/blob/master/mysql/mysql_test.go">mysql_test.go</a>

    rs is a ResultSet,all of ResultSet method and properties is :
    ResultSet.Len()
        how many rows of select
    ResultSet.MapRows(keyColumn string) (rowsMap map[string]map[string]string)
        get a map which key is each value of row[keyColumn]
    ResultSet.MapStructs(keyColumn string, strucT interface{}) (structsMap map[string]interface{},
     err error)
        get a map which key is row[keyColumn],value is strucT
    ResultSet.Rows() (rows []map[string]string)
        get rows of select
    ResultSet.Structs(strucT interface{}) (structs []interface{}, err error)
        get array of strucT of select
    ResultSet.Row() (row map[string]string)
        get first of rows
    ResultSet.Struct(strucT interface{}) (Struct interface{}, err error)
        get first strucT of select 
    ResultSet.Values(column string) (values []string)
        get an array contains each row[column] 
    ResultSet.MapValues(keyColumn, valueColumn string) (values map[string]string)
        get a map key is each row[column],value is row[valueColumn]
    ResultSet.Value(column string) (value string)
        get first row[column] of rows
    ResultSet.LastInsertId
        if sql type is insert , this is the last insert id
    ResultSet.RowsAffected
        if sql type is write , this is the count of rows affected

	
3.Insert & Insert Batch
    Insert:
        rs, err := db.Exec(db.AR().Insert("test", map[string]interface{}{
			"id":   "id11122",
			"name": "333",
		}))
    Insert Batch:
        rs, err := db.Exec(db.AR().InsertBatch("test", []map[string]interface{}{
            map[string]interface{}{
                "id":   "id11122",
                "name": "333",
            },
            map[string]interface{}{
                "id":   "id11122",
                "name": "4444",
            },
        }))
    lastInsertId:=rs.LastInsertId
    rowsAffected:=rs.RowsAffected
    fmt.printf("last insert id : %d,rows affected : %d",lastInsertId,rowsAffected)
	
4.Update & Update Batch
    Update:
    1.common Update(table,data,where)
        rs, err := db.Exec(db.AR().Update("test", map[string]interface{}{
			"id":   "id11122",
			"name": "333",
		}),map[string]interface{}{
			"pid":   223,
		}))
    //equal sql below :
    UPDATE  `test` 
    SET `id` = ? , `name` =  ?
    WHERE `pid` = ?

    2.column operate
    Update:
        rs, err := db.Exec(db.AR().Update("test", map[string]interface{}{
			"id":   "id11122",
			"score +": "333",
		}),map[string]interface{}{
			"pid":   223,
		}))
    //equal sql below :
    UPDATE  `test` 
    SET `id` = ? , `score` = `score` + ?
    WHERE `pid` = ?

    Update Batch:
    1.common update UpdateBatch(table,data,DataWhereColumns)
    rs, err := db.Exec(db.AR().UpdateBatch("test", []map[string]interface{}{
		map[string]interface{}{
			"id":   "id1",
			"name": "333",
		},
		map[string]interface{}{
			"id":   "id2",
			"name": "4444",
		},
	}, []string{"id"}))
    rowsAffected:=rs.RowsAffected
    fmt.printf("rows affected : %d",rowsAffected)
    //equal sql below :
    UPDATE  `test` 
    SET `name` = CASE 
    WHEN `id` = ? THEN  ? 
    WHEN `id` = ? THEN  ? 
    ELSE `score` END 
    WHERE id IN (?,?)

    2.column operate
    rs, err := db.Exec(db.AR().UpdateBatch("test", []map[string]interface{}{
		map[string]interface{}{
			"id":   "id11",
			"score +": 10,
		},
		map[string]interface{}{
			"id":   "id22",
			"score +": 20,
		},
	}, , []string{"id"}))
    rowsAffected:=rs.RowsAffected
    fmt.printf("rows affected : %d",rowsAffected)
    //equal sql below :
    UPDATE  `test` 
    SET `score` = CASE 
    WHEN `id` = ? THEN `score` + ? 
    WHEN `id` = ? THEN `score` + ? 
    ELSE `score` END 
    WHERE id IN (?,?)

    3.where on more column
    rs, err := db.Exec(db.AR().UpdateBatch("test", []map[string]interface{}{
		map[string]interface{}{
			"id":      "id1",
			"gid":     22,
			"name":    "test1",
			"score +": 1,
		}, map[string]interface{}{
			"id":      "id2",
			"gid":     33,
			"name":    "test2",
			"score +": 1,
		},
	}, []string{"id", "gid"})
    rowsAffected:=rs.RowsAffected
    fmt.printf("rows affected : %d",rowsAffected)
    //equal sql below :
    UPDATE  `test` 
    SET `name` = CASE 
    WHEN `id` = ? AND `gid` = ? THEN ? 
    WHEN `id` = ? AND `gid` = ? THEN ? 
    ELSE `name` END,`score` = CASE 
    WHEN `id` = ? AND `gid` = ? THEN `score` + ? 
    WHEN `id` = ? AND `gid` = ? THEN `score` + ? 
    ELSE `score` END 
    WHERE id IN (?,?)   AND gid IN (?,?)
5.Delete
    rs, err := db.Exec(db.AR().Delete("test", map[string]interface{}{
        "pid":   223,
    }))
    rowsAffected:=rs.RowsAffected
    fmt.printf("rows affected : %d",rowsAffected)
6.Raw SQL Query
    rs, err := db.Exec(db.AR().Raw("insert into test(id,name) values (?,?)", 555,"6666"))
    if err != nil {
        fmt.Printf("ERR:%s", err)
    } else {
        fmt.Println(rs.RowsAffected, rs.LastInsertId)
    }
    //notice:
    if  dbCfg.TablePrefix="user_" 
        dbCfg.TablePrefixSqlIdentifier="{__PREFIX__}" 
    then
        db.AR().Raw("insert into {__PREFIX__}test(id,name) values (?,?)
    when execute sql,{__PREFIX__} will be replaced with "user_"
7.Cache Query
    MyCache is an example to set or get cache data . 

    var (
        cacheData = map[string][]byte{}
    )

    type MyCache struct {
    }

    func (c *MyCache) Set(key string, data []byte, expire uint) (err error) {
        cacheData[key] = data
        log.Println("set cache")
        return
    }
    func (c *MyCache) Get(key string) (data []byte, err error) {
        if v, ok := cacheData[key]; ok {
            log.Println("form cache")
            return v, nil
        }
        return nil, errors.New("key not found or expired")
    }
    func main() {
        g := mysql.NewDBGroupCache("default", &MyCache{})
        g.Regist("default", mysql.NewDBConfigWith("127.0.0.1", 3306, "test", "root", "admin"))
        fmt.Println(g.DB().Query(g.DB().AR().Cache("testkey", 30).From("test")))
        rs, _ := g.DB().Query(g.DB().AR().Cache("testkey", 30).From("test"))
        fmt.Println(rs.Row())
    }
    output like:
    2017/11/21 18:12:01 set cache
    &{0xc42000d340 0 0} <nil>
    2017/11/21 18:12:01 form cache
    map[pid:1 id:a1 name:a1111 gid:11]
</pre>