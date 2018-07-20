package models

import "fmt"

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
			"role_id": fmt.Sprintf("%d", Role_Root_Id),
			"privilege_id": privilege["privilege_id"],
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
			"role_id": roleId,
			"privilege_id": privilegeId,
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