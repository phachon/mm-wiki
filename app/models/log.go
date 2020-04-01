package models

import (
	"encoding/json"
	"github.com/astaxie/beego/context"
	"github.com/phachon/mm-wiki/app/utils"
	"github.com/snail007/go-activerecord/mysql"
	"strings"
	"time"
)

const (
	Log_Level_Emegergency = iota
	Log_Level_Alaert
	Log_Level_Critical
	Log_Level_Error
	Log_Level_Warning
	Log_Level_Notice
	Log_Level_Info
	Log_Level_Debug
)

const Table_Log_Name = "log"

type Log struct {
}

var LogModel = Log{}

// 根据 log_id 获取日志
func (l *Log) GetLogByLogId(logId string) (log map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Log_Name).Where(map[string]interface{}{
		"log_id": logId,
	}))
	if err != nil {
		return
	}
	log = rs.Row()
	return
}

// 插入
func (l *Log) Insert(log map[string]interface{}) (id int64, err error) {

	log["create_time"] = time.Now().Unix()

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Insert(Table_Log_Name, log))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// 根据关键字分页获取日志
func (l *Log) GetLogsByKeywordAndLimit(level, message, username string, limit int, number int) (logs []map[string]string, err error) {

	db := G.DB()
	where := make(map[string]interface{})
	if level != "" {
		where["level"] = level
	}
	if message != "" {
		where["message LIKE"] = "%" + message + "%"
	}
	if username != "" {
		where["username LIKE"] = "%" + username + "%"
	}
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Log_Name).Where(where).Limit(limit, number).OrderBy("log_id", "DESC"))
	if err != nil {
		return
	}
	logs = rs.Rows()

	return
}

// 分页获取日志
func (l *Log) GetLogsByLimit(limit int, number int) (logs []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			From(Table_Log_Name).
			Limit(limit, number).
			OrderBy("log_id", "DESC"))
	if err != nil {
		return
	}
	logs = rs.Rows()

	return
}

// 获取日志总数
func (l *Log) CountLogs() (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			Select("count(*) as total").
			From(Table_Log_Name))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

func (l *Log) CountLogsByLevel(level int) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			Select("count(*) as total").
			Where(map[string]interface{}{
				"level": level,
			}).
			From(Table_Log_Name))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// 根据关键字获取日志总数
func (l *Log) CountLogsByKeyword(level, message, username string) (count int64, err error) {

	db := G.DB()
	where := make(map[string]interface{})
	if level != "" {
		where["level"] = level
	}
	if message != "" {
		where["message LIKE"] = "%" + message + "%"
	}
	if username != "" {
		where["username LIKE"] = "%" + username + "%"
	}
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().
		Select("count(*) as total").
		From(Table_Log_Name).
		Where(where))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

func (l *Log) RecordLog(message string, level int, userId string, username string, ctx context.Context) (id int64, err error) {
	userAgent := ctx.Request.UserAgent()
	referer := ctx.Request.Referer()
	getParams := ctx.Request.URL.String()
	path := ctx.Request.URL.Path
	ctx.Request.ParseForm()
	postParamsMap := map[string][]string(ctx.Request.PostForm)
	postParams, _ := json.Marshal(postParamsMap)

	logValue := map[string]interface{}{
		"level":       level,
		"path":        path,
		"get":         getParams,
		"post":        string(postParams),
		"message":     message,
		"ip":          strings.Split(ctx.Request.RemoteAddr, ":"),
		"user_agent":  userAgent,
		"referer":     referer,
		"user_id":     userId,
		"username":    username,
		"create_time": time.Now().Unix(),
	}
	return LogModel.Insert(logValue)
}
