package models

import (
	"time"

	"github.com/phachon/mm-wiki/app/utils"

	"github.com/snail007/go-activerecord/mysql"
)

const (
	Auth_Used_True  = 1
	Auth_Used_False = 0
)

const Table_Auth_Name = "login_auth"

type Auth struct {
}

var AuthModel = Auth{}

// get auth by auth_id
func (a *Auth) GetAuthByAuthId(authId string) (auth map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Auth_Name).Where(map[string]interface{}{
		"login_auth_id": authId,
	}))
	if err != nil {
		return
	}
	auth = rs.Row()
	return
}

// auth_id and name is exists
func (a *Auth) HasSameName(authId, name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Auth_Name).Where(map[string]interface{}{
		"login_auth_id <>": authId,
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

// auth_id and name is exists
func (a *Auth) HasSameUsernamePrefix(authId, usernamePrefix string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Auth_Name).Where(map[string]interface{}{
		"login_auth_id <>": authId,
		"username_prefix":  usernamePrefix,
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
func (a *Auth) HasAuthName(name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Auth_Name).Where(map[string]interface{}{
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

// name is exists
func (a *Auth) HasAuthUsernamePrefix(usernamePrefix string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Auth_Name).Where(map[string]interface{}{
		"username_prefix": usernamePrefix,
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
func (a *Auth) GetAuthByName(name string) (auth map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Auth_Name).Where(map[string]interface{}{
		"name": name,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	auth = rs.Row()
	return
}

// delete auth by auth_id
func (a *Auth) Delete(authId string) (err error) {
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
func (a *Auth) Insert(authValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet

	// is_used
	rs, err = db.Query(db.AR().From(Table_Auth_Name).Where(map[string]interface{}{
		"is_used": Auth_Used_True,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	if rs.Len() == 0 {
		authValue["is_used"] = Auth_Used_True
	} else {
		authValue["is_used"] = Auth_Used_False
	}

	authValue["create_time"] = time.Now().Unix()
	authValue["update_time"] = time.Now().Unix()
	rs, err = db.Exec(db.AR().Insert(Table_Auth_Name, authValue))
	if err != nil {
		return
	}
	id = rs.LastInsertId

	return
}

// update auth by auth_id
func (a *Auth) Update(authId string, authValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	authValue["update_time"] = time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_Auth_Name, authValue, map[string]interface{}{
		"login_auth_id": authId,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// get limit auths by search keyword
func (a *Auth) GetAuthsByKeywordAndLimit(keyword string, limit int, number int) (auths []map[string]string, err error) {

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
func (a *Auth) GetAuthsByLimit(limit int, number int) (auths []map[string]string, err error) {

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
func (a *Auth) GetAuths() (auths []map[string]string, err error) {

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
func (a *Auth) CountAuths() (count int64, err error) {

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
func (a *Auth) CountAuthsByKeyword(keyword string) (count int64, err error) {

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
func (a *Auth) GetAuthsByLikeName(name string) (auths []map[string]string, err error) {
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
func (a *Auth) GetAuthByAuthIds(authIds []string) (auths []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Auth_Name).Where(map[string]interface{}{
		"login_auth_id": authIds,
	}))
	if err != nil {
		return
	}
	auths = rs.Rows()
	return
}

// set auth used
func (a *Auth) SetAuthUsed(authId string) (id int64, err error) {
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

// get used auth
func (a *Auth) GetUsedAuth() (auth map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().Select("*").From(Table_Auth_Name).Where(map[string]interface{}{
		"is_used": Auth_Used_True,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	auth = rs.Row()
	return
}
