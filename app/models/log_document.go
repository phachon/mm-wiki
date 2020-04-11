package models

import (
	"github.com/phachon/mm-wiki/app/utils"
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

func (ld *LogDocument) CreateAction(userId string, documentId string, spaceId string) (id int64, err error) {
	logDocument := map[string]interface{}{
		"user_id":     userId,
		"document_id": documentId,
		"space_id":    spaceId,
		"comment":     "创建了文档",
		"action":      LogDocument_Action_Create,
		"create_time": time.Now().Unix(),
	}
	return ld.Insert(logDocument)
}

func (ld *LogDocument) UpdateAction(userId string, documentId string, comment string) (id int64, err error) {
	logDocument := map[string]interface{}{
		"user_id":     userId,
		"document_id": documentId,
		"comment":     comment,
		"action":      LogDocument_Action_Update,
		"create_time": time.Now().Unix(),
	}
	return ld.Insert(logDocument)
}

func (ld *LogDocument) DeleteAction(userId string, documentId string) (id int64, err error) {
	logDocument := map[string]interface{}{
		"user_id":     userId,
		"document_id": documentId,
		"comment":     "删除了该文档",
		"action":      LogDocument_Action_Delete,
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

func (ld *LogDocument) GetLogDocumentsByUserIdKeywordAndLimit(userId string, keyword string, limit int, number int) (logDocuments []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_LogDocument_Name).Where(map[string]interface{}{
		"comment LIKE": "%" + keyword + "%",
		"user_id":      userId,
	}).Limit(limit, number).OrderBy("log_document_id", "DESC"))
	if err != nil {
		return
	}
	logDocuments = rs.Rows()

	return
}

func (ld *LogDocument) GetLogDocumentsByKeywordAndLimit(keyword string, limit int, number int) (logDocuments []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_LogDocument_Name).Where(map[string]interface{}{
		"comment LIKE": "%" + keyword + "%",
	}).Limit(limit, number).OrderBy("log_document_id", "DESC"))
	if err != nil {
		return
	}
	logDocuments = rs.Rows()

	return
}

func (ld *LogDocument) GetLogDocumentsByLimit(userId string, limit int, number int) (logDocuments []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	where := db.AR().From(Table_LogDocument_Name)

	// 查询用户空间权限
	spaceUserRs, err := db.Query(db.AR().From(Table_SpaceUser_Name).Where(map[string]interface{}{
		"user_id": userId,
	}))

	spaceUsers := spaceUserRs.Rows()
	spaceUsersLen := len(spaceUsers)

	for i := 0; i < spaceUsersLen; i++ {
		spaceUser := spaceUsers[i]
		if i == 0 {
			where.WhereWrap(map[string]interface{}{
				"space_id": spaceUser["space_id"],
			}, "", "")
		} else {
			where.WhereWrap(map[string]interface{}{
				"space_id": spaceUser["space_id"],
			}, "or", "")
		}
	}

	// 查询公共空间
	spaceRs, err := db.Query(db.AR().From(Table_Space_Name).Where(map[string]interface{}{
		"visit_level": "public",
	}))

	spaces := spaceRs.Rows()

	for i := 0; i < len(spaces); i++ {
		space := spaces[i]
		if i == 0 && spaceUsersLen == 0 {
			where.WhereWrap(map[string]interface{}{
				"space_id": space["space_id"],
			}, "", "")
		} else {
			where.WhereWrap(map[string]interface{}{
				"space_id": space["space_id"],
			}, "or", "")
		}
	}

	rs, err = db.Query(where.Limit(limit, number).OrderBy("log_document_id", "DESC"))
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

func (ld *LogDocument) CountLogDocumentsByUserIdAndKeyword(userId string, keyword string) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().Select("count(*) as total").From(Table_LogDocument_Name).Where(map[string]interface{}{
		"comment LIKE": "%" + keyword + "%",
		"user_id":      userId,
	}))
	if err != nil {
		return
	}
	count = utils.Convert.StringToInt64(rs.Value("total"))
	return
}

func (ld *LogDocument) CountLogDocumentsByKeyword(keyword string) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().Select("count(*) as total").From(Table_LogDocument_Name).Where(map[string]interface{}{
		"comment LIKE": "%" + keyword + "%",
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
