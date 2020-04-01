package models

import (
	"github.com/phachon/mm-wiki/app/utils"
	"github.com/snail007/go-activerecord/mysql"
	"time"
)

const (
	Email_Used_True  = 1
	Email_Used_False = 0
)

const Table_Email_Name = "email"

type Email struct {
}

var EmailModel = Email{}

// get email by email_id
func (u *Email) GetEmailByEmailId(emailId string) (email map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Email_Name).Where(map[string]interface{}{
		"email_id": emailId,
	}))
	if err != nil {
		return
	}
	email = rs.Row()
	return
}

// email_id and name is exists
func (u *Email) HasSameName(emailId, name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Email_Name).Where(map[string]interface{}{
		"email_id <>": emailId,
		"name":        name,
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
func (u *Email) HasEmailName(name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Email_Name).Where(map[string]interface{}{
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

// get email by name
func (u *Email) GetEmailByName(name string) (email map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Email_Name).Where(map[string]interface{}{
		"name": name,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	email = rs.Row()
	return
}

// delete email by email_id
func (u *Email) Delete(emailId string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Delete(Table_Email_Name, map[string]interface{}{
		"email_id": emailId,
	}))
	if err != nil {
		return
	}
	return
}

// insert email
func (u *Email) Insert(emailValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet

	// is_used
	rs, err = db.Query(db.AR().From(Table_Email_Name).Where(map[string]interface{}{
		"is_used": Email_Used_True,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	if rs.Len() == 0 {
		emailValue["is_used"] = Email_Used_True
	} else {
		emailValue["is_used"] = Email_Used_False
	}
	emailValue["create_time"] = time.Now().Unix()
	emailValue["update_time"] = time.Now().Unix()

	rs, err = db.Exec(db.AR().Insert(Table_Email_Name, emailValue))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// update email by email_id
func (u *Email) Update(emailId string, emailValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	emailValue["update_time"] = time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_Email_Name, emailValue, map[string]interface{}{
		"email_id": emailId,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// get limit emails by search keyword
func (u *Email) GetEmailsByKeywordAndLimit(keyword string, limit int, number int) (emails []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Email_Name).Where(map[string]interface{}{
		"name LIKE": "%" + keyword + "%",
	}).Limit(limit, number).OrderBy("email_id", "DESC"))
	if err != nil {
		return
	}
	emails = rs.Rows()

	return
}

// get limit emails
func (u *Email) GetEmailsByLimit(limit int, number int) (emails []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			From(Table_Email_Name).
			Limit(limit, number).
			OrderBy("email_id", "DESC"))
	if err != nil {
		return
	}
	emails = rs.Rows()

	return
}

// get all emails
func (u *Email) GetEmails() (emails []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Email_Name))
	if err != nil {
		return
	}
	emails = rs.Rows()
	return
}

// get used email
func (u *Email) GetUsedEmail() (email map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().Select("*").From(Table_Email_Name).Where(map[string]interface{}{
		"is_used": Email_Used_True,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	email = rs.Row()
	return
}

// get email count
func (u *Email) CountEmails() (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			Select("count(*) as total").
			From(Table_Email_Name))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get email count by keyword
func (u *Email) CountEmailsByKeyword(keyword string) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().
		Select("count(*) as total").
		From(Table_Email_Name).
		Where(map[string]interface{}{
			"name LIKE": "%" + keyword + "%",
		}))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get emails by like name
func (u *Email) GetEmailsByLikeName(name string) (emails []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Email_Name).Where(map[string]interface{}{
		"name Like": "%" + name + "%",
	}).Limit(0, 1))
	if err != nil {
		return
	}
	emails = rs.Rows()
	return
}

// get email by many email_id
func (u *Email) GetEmailByEmailIds(emailIds []string) (emails []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Email_Name).Where(map[string]interface{}{
		"email_id": emailIds,
	}))
	if err != nil {
		return
	}
	emails = rs.Rows()
	return
}

// set email used
func (u *Email) SetEmailUsed(emailId string) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet

	rs, err = db.Exec(db.AR().Update(Table_Email_Name, map[string]interface{}{"is_used": Email_Used_False}, map[string]interface{}{
		"is_used": Email_Used_True,
	}))
	if err != nil {
		return
	}
	rs, err = db.Exec(db.AR().Update(Table_Email_Name, map[string]interface{}{"is_used": Email_Used_True}, map[string]interface{}{
		"email_id": emailId,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}
