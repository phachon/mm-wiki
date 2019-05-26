package models

import (
	"github.com/snail007/go-activerecord/mysql"
	"mm-wiki/app/utils"
	"time"
)

const Table_Attachment_Name = "attachment"

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

// get attachments by document
func (a *Attachment) GetAttachmentsByDocumentId(documentId string) (attachments []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Attachment_Name).Where(map[string]interface{}{
		"document_id": documentId,
	}).OrderBy("attachment_id", "desc"))
	if err != nil {
		return
	}
	attachments = rs.Rows()
	return
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
		_, _ = LogDocumentModel.UpdateAction(attachmentValue["user_id"].(string),
			attachmentValue["document_id"].(string), "上传了附件 "+attachmentValue["name"].(string))
	}()

	return
}

// update attachment by attachment_id
func (a *Attachment) Update(attachmentId string, attachmentValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	attachmentValue["update_time"] = time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_Attachment_Name, attachmentValue, map[string]interface{}{
		"attachment_id": attachmentId,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// get limit attachments
func (a *Attachment) GetAttachmentsByLimit(limit int, number int) (attachments []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			From(Table_Attachment_Name).
			Limit(limit, number).
			OrderBy("attachment_id", "DESC"))
	if err != nil {
		return
	}
	attachments = rs.Rows()

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

// get all attachments by sequence
func (a *Attachment) GetAttachmentsOrderBySequence() (attachments []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Attachment_Name).OrderBy("sequence", "ASC"))
	if err != nil {
		return
	}
	attachments = rs.Rows()
	return
}

// get attachment count
func (a *Attachment) CountAttachments() (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			Select("count(*) as total").
			From(Table_Attachment_Name))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get attachment count by keyword
func (a *Attachment) CountAttachmentsByKeyword(keyword string) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().
		Select("count(*) as total").
		From(Table_Attachment_Name).
		Where(map[string]interface{}{
			"name LIKE": "%" + keyword + "%",
		}))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
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
