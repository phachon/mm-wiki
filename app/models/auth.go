package models

import (
	"mm-wiki/app/utils"
	"github.com/snail007/go-activerecord/mysql"
	"time"
)

const (
	Auth_Used_True = 1
	Auth_Used_False = 0
)

const Table_Auth_Name = "login_auth"

type Auth struct {

}

var AuthModel = Auth{}

// get auth by auth_id
func (u *Auth) GetAuthByAuthId(authId string) (auth map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Auth_Name).Where(map[string]interface{}{
		"login_auth_id":   authId,
	}))
	if err != nil {
		return
	}
	auth = rs.Row()
	return
}

// auth_id and name is exists
func (u *Auth) HasSameName(authId, name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Auth_Name).Where(map[string]interface{}{
		"login_auth_id <>": authId,
		"name":   name,
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
func (u *Auth) HasAuthName(name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Auth_Name).Where(map[string]interface{}{
		"name":  name,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	if rs.Len() > 0 {
		has = true
	}
	return
}

// get auth by name
func (u *Auth) GetAuthByName(name string) (auth map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Auth_Name).Where(map[string]interface{}{
		"name":  name,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	auth = rs.Row()
	return
}

// delete auth by auth_id
func (u *Auth) Delete(authId string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Delete(Table_Auth_Name, map[string]interface{}{
		"login_auth_id": authId,
	}))
	if err != nil {
		return
	}
	return
}

// insert auth
func (u *Auth) Insert(authValue map[string]interface{}) (id int64, err error) {

	authValue["create_time"] = time.Now().Unix()
	authValue["update_time"] = time.Now().Unix()
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Insert(Table_Auth_Name, authValue))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// update auth by auth_id
func (u *Auth) Update(authId string, authValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	authValue["update_time"] =  time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_Auth_Name, authValue, map[string]interface{}{
		"login_auth_id":   authId,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// get limit auths by search keyword
func (u *Auth) GetAuthsByKeywordAndLimit(keyword string, limit int, number int) (auths []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Auth_Name).Where(map[string]interface{}{
		"name LIKE": "%" + keyword + "%",
	}).Limit(limit, number).OrderBy("auth_id", "DESC"))
	if err != nil {
		return
	}
	auths = rs.Rows()

	return
}

// get limit auths
func (u *Auth) GetAuthsByLimit(limit int, number int) (auths []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			From(Table_Auth_Name).
			Limit(limit, number).
			OrderBy("login_auth_id", "DESC"))
	if err != nil {
		return
	}
	auths = rs.Rows()

	return
}

// get all auths
func (u *Auth) GetAuths() (auths []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Auth_Name))
	if err != nil {
		return
	}
	auths = rs.Rows()
	return
}

// get auth count
func (u *Auth) CountAuths() (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			Select("count(*) as total").
			From(Table_Auth_Name))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get auth count by keyword
func (u *Auth) CountAuthsByKeyword(keyword string) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().
		Select("count(*) as total").
		From(Table_Auth_Name).
		Where(map[string]interface{}{
			"name LIKE": "%" + keyword + "%",
		}))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get auths by like name
func (u *Auth) GetAuthsByLikeName(name string) (auths []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Auth_Name).Where(map[string]interface{}{
		"name Like": "%" + name + "%",
	}).Limit(0, 1))
	if err != nil {
		return
	}
	auths = rs.Rows()
	return
}

// get auth by many auth_id
func (u *Auth) GetAuthByAuthIds(authIds []string) (auths []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Auth_Name).Where(map[string]interface{}{
		"login_auth_id":   authIds,
	}))
	if err != nil {
		return
	}
	auths = rs.Rows()
	return
}

// set auth used
func (u *Auth) SetAuthUsed(authId string) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet

	rs, err = db.Exec(db.AR().Update(Table_Auth_Name, map[string]interface{}{"is_used": Auth_Used_False}, map[string]interface{}{
		"is_used": Auth_Used_True,
	}))
	if err != nil {
		return
	}
	rs, err = db.Exec(db.AR().Update(Table_Auth_Name, map[string]interface{}{"is_used": Auth_Used_True}, map[string]interface{}{
		"login_auth_id": authId,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}
