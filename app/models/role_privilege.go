package models

import (
	"fmt"
	"time"
)

var RolePrivilegeModel = RolePrivilege{}

const Table_RolePrivilege_Name = "role_privilege"

type RolePrivilege struct {
}

func (rolePrivilege *RolePrivilege) GetRolePrivilegesByRoleId(roleId string) (rolePrivileges []map[string]string, err error) {

	db := G.DB()
	res, err := db.Query(db.AR().From(Table_RolePrivilege_Name).Where(map[string]interface{}{
		"role_id": roleId,
	}))
	if err != nil {
		return
	}
	rolePrivileges = res.Rows()
	return
}

func (rolePrivilege *RolePrivilege) GetRootRolePrivileges() (rolePrivileges []map[string]string, err error) {

	privileges, err := PrivilegeModel.GetPrivileges()
	if err != nil {
		return
	}
	for _, privilege := range privileges {
		rolePrivilege := map[string]string{
			"role_privilege_id": "",
			"role_id":           fmt.Sprintf("%d", Role_Root_Id),
			"privilege_id":      privilege["privilege_id"],
		}
		rolePrivileges = append(rolePrivileges, rolePrivilege)
	}
	return
}

func (rolePrivilege *RolePrivilege) GrantRolePrivileges(roleId string, privilegeIds []string) (res bool, err error) {

	res = false
	db := G.DB()
	_, err = db.Exec(db.AR().Delete(Table_RolePrivilege_Name, map[string]interface{}{
		"role_id": roleId,
	}))
	if err != nil {
		return
	}

	rolePrivileges := []map[string]interface{}{}
	for _, privilegeId := range privilegeIds {
		rolePrivilege := map[string]interface{}{
			"role_id":      roleId,
			"privilege_id": privilegeId,
			"create_time":  time.Now().Unix(),
		}
		rolePrivileges = append(rolePrivileges, rolePrivilege)
	}
	_, err = db.Exec(db.AR().InsertBatch(Table_RolePrivilege_Name, rolePrivileges))
	if err != nil {
		return
	}
	res = true
	return
}

// delete role privilege by role_id
func (rolePrivilege *RolePrivilege) DeleteByRoleId(roleId string) (err error) {

	db := G.DB()
	_, err = db.Exec(db.AR().Delete(Table_RolePrivilege_Name, map[string]interface{}{
		"role_id": roleId,
	}))
	if err != nil {
		return
	}
	return
}

// delete role privilege by privilege_id
func (rolePrivilege *RolePrivilege) DeleteByPrivilegeId(privilegeId string) (err error) {

	db := G.DB()
	_, err = db.Exec(db.AR().Delete(Table_RolePrivilege_Name, map[string]interface{}{
		"privilege_id": privilegeId,
	}))
	if err != nil {
		return
	}
	return
}
