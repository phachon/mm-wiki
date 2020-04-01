package models

import (
	"fmt"
	"github.com/phachon/mm-wiki/app/utils"
	"github.com/snail007/go-activerecord/mysql"
	"time"
)

const Table_Attachment_Name = "attachment"

const (
	Attachment_Source_Default = 0
	Attachment_Source_Image   = 1
)

type Attachment struct {
}

var AttachmentModel = Attachment{}

// get attachment by attachment_id
func (a *Attachment) GetAttachmentByAttachmentId(attachmentId string) (attachment map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Attachment_Name).Where(map[string]interface{}{
		"attachment_id": attachmentId,
	}))
	if err != nil {
		return
	}
	attachment = rs.Row()
	return
}

// attachment_id and name is exists
func (a *Attachment) HasSameName(attachmentId, name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Attachment_Name).Where(map[string]interface{}{
		"attachment_id <>": attachmentId,
		"name":             name,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	if rs.Len() > 0 {
		has = true
	}
	return
}

// name is exists
func (a *Attachment) HasAttachmentName(name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Attachment_Name).Where(map[string]interface{}{
		"name": name,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	if rs.Len() > 0 {
		has = true
	}
	return
}

// get attachment by name
func (a *Attachment) GetAttachmentByName(name string) (attachment map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Attachment_Name).Where(map[string]interface{}{
		"name": name,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	attachment = rs.Row()
	return
}

// get attachments by document and source
func (a *Attachment) GetAttachmentsByDocumentIdAndSource(documentId string, source int) (attachments []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Attachment_Name).Where(map[string]interface{}{
		"document_id": documentId,
		"source":      source,
	}).OrderBy("attachment_id", "desc"))
	if err != nil {
		return
	}
	attachments = rs.Rows()
	return
}

// get attachments by document_id
func (a *Attachment) GetAttachmentsByDocumentId(documentId string) (attachments []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Attachment_Name).Where(map[string]interface{}{
		"document_id": documentId,
	}))
	if err != nil {
		return
	}
	attachments = rs.Rows()
	return
}

// get attachments by document_id
func (a *Attachment) GetAttachmentsByDocumentIds(documentIds []string) (attachments []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Attachment_Name).Where(map[string]interface{}{
		"document_id": documentIds,
	}))
	if err != nil {
		return
	}
	attachments = rs.Rows()
	return
}

// get attachments by space_id
func (a *Attachment) GetAttachmentsBySpaceId(spaceId string) (attachments []map[string]string, err error) {
	documents, err := DocumentModel.GetDocumentsBySpaceId(spaceId)
	if err != nil {
		return
	}
	documentIds := []string{}
	for _, document := range documents {
		documentIds = append(documentIds, document["document_id"])
	}

	return a.GetAttachmentsByDocumentIds(documentIds)
}

// delete attachment by attachment_id
func (a *Attachment) Delete(attachmentId string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Delete(Table_Attachment_Name, map[string]interface{}{
		"attachment_id": attachmentId,
	}))
	if err != nil {
		return
	}

	return
}

// insert attachment
func (a *Attachment) Insert(attachmentValue map[string]interface{}) (id int64, err error) {

	attachmentValue["create_time"] = time.Now().Unix()
	attachmentValue["update_time"] = time.Now().Unix()
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Insert(Table_Attachment_Name, attachmentValue))
	if err != nil {
		return
	}
	id = rs.LastInsertId

	// create document log
	go func() {
		comment := fmt.Sprintf("上传了附件 %s", attachmentValue["name"].(string))
		if attachmentValue["source"].(int) == Attachment_Source_Image {
			comment = fmt.Sprintf("上传了图片 %s", attachmentValue["name"].(string))
		}
		_, _ = LogDocumentModel.UpdateAction(attachmentValue["user_id"].(string),
			attachmentValue["document_id"].(string), comment)
	}()

	return
}

// get all attachments
func (a *Attachment) GetAttachments() (attachments []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Attachment_Name))
	if err != nil {
		return
	}
	attachments = rs.Rows()
	return
}

// get attachments by like name
func (a *Attachment) GetAttachmentsByLikeName(name string) (attachments []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Attachment_Name).Where(map[string]interface{}{
		"name Like": "%" + name + "%",
	}).Limit(0, 1))
	if err != nil {
		return
	}
	attachments = rs.Rows()
	return
}

// get attachment by many attachment_id
func (a *Attachment) GetAttachmentByAttachmentIds(attachmentIds []string) (attachments []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Attachment_Name).Where(map[string]interface{}{
		"attachment_id": attachmentIds,
	}))
	if err != nil {
		return
	}
	attachments = rs.Rows()
	return
}

func (a *Attachment) DeleteAttachmentsDBFileByDocumentId(documentId string) (err error) {
	db := G.DB()
	attachments, err := a.GetAttachmentsByDocumentId(documentId)
	if err != nil {
		return
	}

	// delete attachment file
	err = utils.Document.DeleteAttachment(attachments)
	if err != nil {
		return err
	}

	// delete db attachment
	_, err = db.Exec(db.AR().Delete(Table_Attachment_Name, map[string]interface{}{
		"document_id": documentId,
	}))
	if err != nil {
		return
	}
	return nil
}

func (a *Attachment) DeleteAttachmentDBFile(attachmentId string) (err error) {
	db := G.DB()
	attachment, err := a.GetAttachmentByAttachmentId(attachmentId)
	if err != nil {
		return
	}

	// delete attachment file
	err = utils.Document.DeleteAttachment([]map[string]string{attachment})
	if err != nil {
		return err
	}

	// delete db attachment
	_, err = db.Exec(db.AR().Delete(Table_Attachment_Name, map[string]interface{}{
		"attachment_id": attachmentId,
	}))
	if err != nil {
		return
	}
	return nil
}
