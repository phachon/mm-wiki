package models

import (
	"mm-wiki/app/utils"
	"github.com/snail007/go-activerecord/mysql"
	"time"
	"github.com/astaxie/beego"
	"path/filepath"
	"strings"
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

// get document by name
func (d *Document) GetDocumentsByTitle(title string) (documents []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"title": title,
		"is_delete": Document_Delete_False,
	}))
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

// get document by title and spaceId
func (d *Document) GetDocumentByTitleAndSpaceId(title string, spaceId string) (document map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"title": title,
		"space_id": spaceId,
		"is_delete": Document_Delete_False,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	document = rs.Row()
	return
}

// delete document by document_id
func (d *Document) Delete(documentId string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Update(Table_Document_Name, map[string]interface{}{
		"is_delete": Document_Delete_False,
		"update_time": time.Now().Unix(),
	}, map[string]interface{}{
		"document_id": documentId,
	}))
	if err != nil {
		return
	}
	return
}

// insert document
func (d *Document) Insert(documentValue map[string]interface{}) (id int64, err error) {
	documentValue["create_time"] = time.Now().Unix()
	documentValue["update_time"] = time.Now().Unix()
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Insert(Table_Document_Name, documentValue))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// update document by document_id
func (d *Document) Update(documentId string, documentValue map[string]interface{}) (id int64, err error) {
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

// get document by name
func (d *Document) GetDocumentsByLikeTitle(name string) (documents []map[string]string, err error) {
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

// get document by many document_id
func (d *Document) GetDocumentByDocumentIds(documentIds []string) (documents []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"document_id":   documentIds,
		"is_delete": Document_Delete_False,
	}))
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

func (d *Document) GetContentByPath(path string) (content string , err error){
	docRootDir := strings.TrimRight(beego.AppConfig.String("document::root_dir"), "/")
	absDir, _ := filepath.Abs(docRootDir)
	realPath := absDir + "/"+path
	return utils.File.GetFileContents(realPath)
}