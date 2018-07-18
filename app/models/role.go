package models

import (
	"mm-wiki/app/utils"
	"github.com/snail007/go-activerecord/mysql"
	"time"
)

const (
	Role_Delete_True = 1
	Role_Delete_False = 0

	Role_Type_System = 1
	Role_Type_Default = 0

	Role_Root_Id = 1
	Role_Admin_Id = 2
	Role_Default_Id = 3
)

const Table_Role_Name = "role"

type Role struct {

}

var RoleModel = Role{}

// get role by role_id
func (u *Role) GetRoleByRoleId(roleId string) (role map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Role_Name).Where(map[string]interface{}{
		"role_id":   roleId,
		"is_delete": Role_Delete_False,
	}))
	if err != nil {
		return
	}
	role = rs.Row()
	return
}

// role_id and name is exists
func (u *Role) HasSameName(roleId, name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Role_Name).Where(map[string]interface{}{
		"role_id <>": roleId,
		"name":   name,
		"is_delete":  Role_Delete_False,
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
func (u *Role) HasRoleName(name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Role_Name).Where(map[string]interface{}{
		"name":  name,
		"is_delete": Role_Delete_False,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	if rs.Len() > 0 {
		has = true
	}
	return
}

// get role by name
func (u *Role) GetRoleByName(name string) (role map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Role_Name).Where(map[string]interface{}{
		"name":  name,
		"is_delete": Role_Delete_False,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	role = rs.Row()
	return
}

// delete role by role_id
func (u *Role) Delete(roleId string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Update(Table_Role_Name, map[string]interface{}{
		"is_delete": Role_Delete_False,
		"update_time": time.Now().Unix(),
	}, map[string]interface{}{
		"role_id": roleId,
	}))
	if err != nil {
		return
	}
	return
}

// insert role
func (u *Role) Insert(roleValue map[string]interface{}) (id int64, err error) {

	roleValue["create_time"] =  time.Now().Unix()
	roleValue["update_time"] =  time.Now().Unix()
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Insert(Table_Role_Name, roleValue))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// update role by role_id
func (u *Role) Update(roleId string, roleValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	roleValue["update_time"] =  time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_Role_Name, roleValue, map[string]interface{}{
		"role_id":   roleId,
		"is_delete": Role_Delete_False,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// get limit roles by search keyword
func (u *Role) GetRolesByKeywordAndLimit(keyword string, limit int, number int) (roles []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Role_Name).Where(map[string]interface{}{
		"name LIKE": "%" + keyword + "%",
		"is_delete":     Role_Delete_False,
	}).Limit(limit, number).OrderBy("role_id", "DESC"))
	if err != nil {
		return
	}
	roles = rs.Rows()

	return
}

// get limit roles
func (u *Role) GetRolesByLimit(limit int, number int) (roles []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			From(Table_Role_Name).
			Where(map[string]interface{}{
				"is_delete": Role_Delete_False,
			}).
			Limit(limit, number).
			OrderBy("role_id", "DESC"))
	if err != nil {
		return
	}
	roles = rs.Rows()

	return
}

// get all roles
func (u *Role) GetRoles() (roles []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Role_Name).Where(map[string]interface{}{
			"is_delete": Role_Delete_False,
		}))
	if err != nil {
		return
	}
	roles = rs.Rows()
	return
}

// get role count
func (u *Role) CountRoles() (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			Select("count(*) as total").
			From(Table_Role_Name).
			Where(map[string]interface{}{
				"is_delete": Role_Delete_False,
			}))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get role count by keyword
func (u *Role) CountRolesByKeyword(keyword string) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().
		Select("count(*) as total").
		From(Table_Role_Name).
		Where(map[string]interface{}{
			"name LIKE": "%" + keyword + "%",
			"is_delete": Role_Delete_False,
		}))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get role by name
func (u *Role) GetRoleByLikeName(name string) (roles []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Role_Name).Where(map[string]interface{}{
		"name Like": "%" + name + "%",
		"is_delete": Role_Delete_False,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	roles = rs.Rows()
	return
}

// get role by many role_id
func (u *Role) GetRoleByRoleIds(roleIds []string) (roles []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Role_Name).Where(map[string]interface{}{
		"role_id":   roleIds,
		"is_delete": Role_Delete_False,
	}))
	if err != nil {
		return
	}
	roles = rs.Rows()
	return
}

// update role by name
func (u *Role) UpdateRoleByName(role map[string]interface{}) (affect int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	role["update_time"] = time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_Role_Name, role, map[string]interface{}{
		"name": role["name"],
		"is_delete": Role_Delete_False,
	}))
	if err != nil {
		return
	}
	affect = rs.RowsAffected
	return
}
