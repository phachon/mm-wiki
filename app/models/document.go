package models

import (
	"mm-wiki/app/utils"
	"github.com/snail007/go-activerecord/mysql"
	"time"
	"strings"
	"fmt"
)

const (
	Document_Delete_True = 1
	Document_Delete_False = 0

	Document_Type_Page = 1
	Document_Type_Dir = 2
)

const Table_Document_Name = "document"

type Document struct {

}

var DocumentModel = Document{}

// get document by document_id
func (d *Document) GetDocumentByDocumentId(documentId string) (document map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"document_id":   documentId,
		"is_delete": Document_Delete_False,
	}))
	if err != nil {
		return
	}
	document = rs.Row()
	return
}

// get documents by parent_id
func (d *Document) GetDocumentsByParentId(parentId string) (documents []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"parent_id": parentId,
		"is_delete": Document_Delete_False,
	}))
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

// get document by name
func (d *Document) GetDocumentsByName(name string) (documents []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"name": name,
		"is_delete": Document_Delete_False,
	}))
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

// get document by name and spaceId
func (d *Document) GetDocumentByNameAndSpaceId(name string, spaceId string) (document map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"name": name,
		"space_id": spaceId,
		"is_delete": Document_Delete_False,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	document = rs.Row()
	return
}

// get document by name and spaceId
func (d *Document) GetDocumentByNameParentIdAndSpaceId(name string, parentId string, spaceId string, docType int) (document map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"name": name,
		"space_id": spaceId,
		"parent_id": parentId,
		"type": docType,
		"is_delete": Document_Delete_False,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	document = rs.Row()
	return
}

// delete document by document_id
func (d *Document) Delete(documentId string, userId string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Update(Table_Document_Name, map[string]interface{}{
		"is_delete": Document_Delete_True,
		"update_time": time.Now().Unix(),
		"edit_user_id": userId,
	}, map[string]interface{}{
		"document_id": documentId,
	}))
	if err != nil {
		return
	}

	// create document log
	_, err = LogDocumentModel.DeleteAction(userId, documentId)
	if err != nil {
		return
	}

	// delete follow doc
	err = FollowModel.DeleteByObjectIdType(fmt.Sprintf("%d", Follow_Type_Doc), documentId)
	if err != nil {
		return
	}

	// delete collect doc
	err = CollectionModel.DeleteByResourceIdType(fmt.Sprintf("%d", Collection_Type_Doc), documentId)
	if err != nil {
		return
	}
	return
}

// insert document
func (d *Document) Insert(documentValue map[string]interface{}) (id int64, err error) {

	document := map[string]string{
		"space_id": documentValue["space_id"].(string),
		"parent_id": documentValue["parent_id"].(string),
		"name": documentValue["name"].(string),
		"type": fmt.Sprintf("%d", documentValue["type"].(int)),
		"path": documentValue["path"].(string),
	}
	_, pageFile, err := d.GetParentDocumentsByDocument(document)
	err = utils.Document.Create(pageFile)
	if err != nil {
		return
	}
	documentValue["create_time"] = time.Now().Unix()
	documentValue["update_time"] = time.Now().Unix()
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Insert(Table_Document_Name, documentValue))
	if err != nil {
		return
	}
	id = rs.LastInsertId

	// create document log
	_, err = LogDocumentModel.CreateAction(documentValue["create_user_id"].(string), fmt.Sprintf("%d", id))
	if err != nil {
		return
	}

	// follow document
	_, err = FollowModel.createFollowDocument(documentValue["create_user_id"].(string), fmt.Sprintf("%d", id))
	if err != nil {
		return
	}
	return
}

// update document by document_id
func (d *Document) Update(documentId string, documentValue map[string]interface{}, comment string) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	documentValue["update_time"] =  time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_Document_Name, documentValue, map[string]interface{}{
		"document_id":   documentId,
		"is_delete": Document_Delete_False,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId

	// create document log
	_, err = LogDocumentModel.UpdateAction(documentValue["edit_user_id"].(string), documentId, comment)
	if err != nil {
		return
	}

	// follow document
	_, err = FollowModel.createFollowDocument(documentValue["edit_user_id"].(string), documentId)
	if err != nil {
		return
	}
	return
}

// get all documents
func (d *Document) GetDocumentsBySpaceId(spaceId string) (documents []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Document_Name).Where(map[string]interface{}{
			"space_id": spaceId,
			"is_delete": Document_Delete_False,
		}))
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

// get document by spaceId and parentId
func (d *Document) GetDocumentsBySpaceIdAndParentId(spaceId string, parentId string) (documents []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Document_Name).Where(map[string]interface{}{
			"space_id": spaceId,
			"parent_id": parentId,
			"is_delete": Document_Delete_False,
		}))
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

// get document by spaceId
func (d *Document) GetSpaceDefaultDocument(spaceId string) (document map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Document_Name).Where(map[string]interface{}{
			"space_id": spaceId,
			"parent_id": "0",
			"is_delete": Document_Delete_False,
		}).Limit(0, 1))
	if err != nil {
		return
	}
	document = rs.Row()
	return
}

// get document by spaceId
func (d *Document) GetAllSpaceDocuments(spaceId string) (documents []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Document_Name).Where(map[string]interface{}{
			"space_id": spaceId,
			"parent_id >": "0",
			"is_delete": Document_Delete_False,
		}))
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

// get document count
func (d *Document) CountDocumentsBySpaceId(spaceId string) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			Select("count(*) as total").
			From(Table_Document_Name).
			Where(map[string]interface{}{
				"space_id": spaceId,
				"is_delete": Document_Delete_False,
			}))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get document count
func (d *Document) CountDocuments() (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			Select("count(*) as total").
			From(Table_Document_Name).
			Where(map[string]interface{}{
			"is_delete": Document_Delete_False,
		}))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get document by name
func (d *Document) GetDocumentsByLikeName(name string) (documents []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"name Like": "%" + name + "%",
		"is_delete":     Document_Delete_False,
	}))
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

// get document link name and limit
func (d *Document) GetDocumentsByLikeNameAndLimit(name string, limit int, number int) (documents []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"name Like": "%" + name + "%",
		"is_delete":     Document_Delete_False,
	}).Limit(limit, number))
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

// count document like name
func (d *Document) CountDocumentsLikeName(name string) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			Select("count(*) as total").
			From(Table_Document_Name).
			Where(map[string]interface{}{
			"name Like": "%" + name + "%",
			"is_delete": Document_Delete_False,
		}))
	if err != nil {
		return
	}
	count = utils.Convert.StringToInt64(rs.Value("total"))
	return
}

// get document by spaceId and document_ids
func (d *Document) GetDocumentsByDocumentIds(documentIds []string) (documents []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"document_id": documentIds,
		"is_delete": Document_Delete_False,
	}))
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

func (d *Document) GetParentDocumentsByDocument(document map[string]string) (parentDocuments []map[string]string, pageFile string, err error) {

	if document["parent_id"] == "0" {
		parentDocuments = append(parentDocuments, document)
		pageFile = utils.Document.GetDefaultPageFileBySpaceName(document["name"])
	}else {
		documentsIds := strings.Split(document["path"], ",")
		parentDocuments, err = d.GetDocumentsByDocumentIds(documentsIds)
		if err != nil {
			return
		}
		var parentPath = ""
		for _, parentDocument := range parentDocuments {
			parentPath += parentDocument["name"]+"/"
		}
		parentPath = strings.TrimRight(parentPath, "/")
		pageFile = utils.Document.GetPageFileByParentPath(document["name"], utils.Convert.StringToInt(document["type"]), parentPath)
	}
	return
}

func (d *Document) GetParentDocumentsByPath(path string) (parentDocuments []map[string]string, err error) {
	documentsIds := strings.Split(path, ",")
	parentDocuments, err = d.GetDocumentsByDocumentIds(documentsIds)
	if err != nil {
		return
	}
	return
}

func (d *Document) GetSpaceIdsOrderByCountDocumentLimit(limit int) (documents []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	sql := db.AR().Select("space_id, count('space_id') as total").
		From(Table_Document_Name).Where(map[string]interface{}{
			"is_delete": Document_Delete_False,
		}).
		GroupBy("space_id").
		OrderBy("total", "DESC").
		Limit(0, limit)
	rs, err = db.Query(sql)
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

func (d *Document) GetCountGroupByCreateTime(startTime int64) (documents []map[string]string, err error) {
	/*select month(FROM_UNIXTIME(time)) from table_name group by month(FROM_UNIXTIME(time))*/

	/*select DATE_FORMAT(FROM_UNIXTIME(time),"%Y-%m") from tcm_fund_list group by DATE_FORMAT(FROM_UNIXTIME(time),"%Y-%m")*/
	/*select FROM_UNIXTIME(time,"%Y-%m") from tcm_fund_list group by FROM_UNIXTIME(time,"%Y-%m")*/

	/*select DATE_FORMAT(FROM_UNIXTIME(time),"%Y-%m-%d") from tcm_fund_list group by DATE_FORMAT(FROM_UNIXTIME(time),"%Y-%m-%d")*/
	db := G.DB()
	var rs *mysql.ResultSet
	sql := db.AR().Select("DATE_FORMAT(FROM_UNIXTIME(create_time),'%Y-%m-%d') as date, count('date') as total").
		From(Table_Document_Name).Where(map[string]interface{}{
		"is_delete": Document_Delete_False,
		"create_time >=": startTime,
	}).GroupBy("DATE_FORMAT(FROM_UNIXTIME(create_time),'%Y-%m-%d')")
	rs, err = db.Query(sql)
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

func (d *Document) GetDocumentGroupCreateUserId() (documents []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	sql := db.AR().Select("create_user_id, count('create_user_id') as total").
		From(Table_Document_Name).Where(map[string]interface{}{
		"is_delete": Document_Delete_False,
	}).GroupBy("create_user_id")
	rs, err = db.Query(sql)
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

func (d *Document) GetDocumentGroupEditUserId() (documents []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	sql := db.AR().Select("edit_user_id, count('edit_user_id') as total").
		From(Table_Document_Name).Where(map[string]interface{}{
		"is_delete": Document_Delete_False,
	}).GroupBy("edit_user_id")
	rs, err = db.Query(sql)
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}