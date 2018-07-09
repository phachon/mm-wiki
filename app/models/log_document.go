package models

import (
	"mm-wiki/app/utils"
	"github.com/snail007/go-activerecord/mysql"
	"time"
)

const (
	LogDocument_Action_Create = 1
	LogDocument_Action_Update = 2
	LogDocument_Action_Delete = 3
)

const Table_LogDocument_Name = "log_document"

type LogDocument struct {
	
}

var LogDocumentModel = LogDocument{}

func (ld *LogDocument) GetLogDocumentByLogDocumentId(logDocId string) (logDocuments map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_LogDocument_Name).Where(map[string]interface{}{
		"log_document_id": logDocId,
	}))
	if err != nil {
		return
	}
	logDocuments = rs.Row()
	return
}

func (ld *LogDocument) Insert(logDocument map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Insert(Table_LogDocument_Name, logDocument))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

func (ld *LogDocument) CreateAction(userId string, documentId string) (id int64, err error) {
	logDocument := map[string]interface{}{
		"user_id": userId,
		"document_id": documentId,
		"comment": "创建了文档",
		"action": LogDocument_Action_Create,
		"create_time": time.Now().Unix(),
	}
	return ld.Insert(logDocument)
}

func (ld *LogDocument) UpdateAction(userId string, documentId string, comment string) (id int64, err error) {
	logDocument := map[string]interface{}{
		"user_id": userId,
		"document_id": documentId,
		"comment": comment,
		"action": LogDocument_Action_Update,
		"create_time": time.Now().Unix(),
	}
	return ld.Insert(logDocument)
}

func (ld *LogDocument) DeleteAction(userId string, documentId string) (id int64, err error) {
	logDocument := map[string]interface{}{
		"user_id": userId,
		"document_id": documentId,
		"comment": "删除了该文档",
		"action": LogDocument_Action_Delete,
		"create_time": time.Now().Unix(),
	}
	return ld.Insert(logDocument)
}

func (ld *LogDocument) GetLogDocumentsByDocumentId(documentId string) (logDocuments []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_LogDocument_Name).Where(map[string]interface{}{
		"document_id": documentId,
	}))
	if err != nil {
		return
	}
	logDocuments = rs.Rows()
	return
}

func (ld *LogDocument) GetLogDocumentsByUserId(userId string) (logDocuments []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_LogDocument_Name).Where(map[string]interface{}{
		"user_id": userId,
	}))
	if err != nil {
		return
	}
	logDocuments = rs.Rows()
	return
}

func (ld *LogDocument) GetLogDocumentsByDocumentIdAndLimit(documentId string, limit int, number int) (logDocuments []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_LogDocument_Name).Where(map[string]interface{}{
		"document_id": documentId,
	}).Limit(limit, number).OrderBy("log_document_id", "DESC"))
	if err != nil {
		return
	}
	logDocuments = rs.Rows()

	return
}

func (ld *LogDocument) GetLogDocumentsByUserIdAndLimit(userId string, limit int, number int) (logDocuments []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_LogDocument_Name).Where(map[string]interface{}{
		"user_id": userId,
	}).Limit(limit, number).OrderBy("log_document_id", "DESC"))
	if err != nil {
		return
	}
	logDocuments = rs.Rows()

	return
}

func (ld *LogDocument) CountLogDocumentsByDocumentId(documentId string) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().Select("count(*) as total").From(Table_LogDocument_Name).Where(map[string]interface{}{
		"document_id": documentId,
	}))
	if err != nil {
		return
	}
	count = utils.Convert.StringToInt64(rs.Value("total"))
	return
}

func (ld *LogDocument) CountLogDocumentsByUserId(userId string) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().Select("count(*) as total").From(Table_LogDocument_Name).Where(map[string]interface{}{
		"user_id": userId,
	}))
	if err != nil {
		return
	}
	count = utils.Convert.StringToInt64(rs.Value("total"))
	return
}

func (ld *LogDocument) CountLogDocuments() (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			Select("count(*) as total").
			From(Table_LogDocument_Name))
	if err != nil {
		return
	}
	count = utils.Convert.StringToInt64(rs.Value("total"))
	return
}