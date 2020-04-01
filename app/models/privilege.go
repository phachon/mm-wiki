package models

import (
	"fmt"
	"github.com/snail007/go-activerecord/mysql"
	"time"
)

type Privilege struct {
}

const Table_Privilege_Name = "privilege"

const (
	Privilege_Type_Menu       = "menu"
	Privilege_Type_Controller = "controller"

	Privilege_DisPlay_True  = "1"
	Privilege_DisPlay_False = "0"
)

var Privilege_Default_Ids = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

var PrivilegeModel = Privilege{}

func (p *Privilege) GetTypePrivileges() (menus, controllers []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Privilege_Name).OrderBy("sequence", "ASC"))
	if err != nil {
		return
	}
	menus = []map[string]string{}
	controllers = []map[string]string{}
	for _, row := range rs.Rows() {
		switch row["type"] {
		case Privilege_Type_Menu:
			menus = append(menus, row)
		case Privilege_Type_Controller:
			controllers = append(controllers, row)
		}
	}
	return
}

func (p *Privilege) GetPrivileges() (privileges []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Privilege_Name))
	if err != nil {
		return
	}
	privileges = rs.Rows()

	return
}

func (p *Privilege) GetPrivilegeByTypeControllerAndAction(ty, controller, action string) (privilege map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Privilege_Name).Where(map[string]interface{}{
		"type":       ty,
		"controller": controller,
		"action":     action,
	}))
	if err != nil {
		return
	}
	privilege = rs.Row()
	return
}

func (p *Privilege) GetTypePrivilegesByUserId(userId string) (menus, controllers []map[string]string, err error) {

	user, err := UserModel.GetUserByUserId(userId)
	if err != nil {
		return
	}
	if len(user) == 0 {
		return
	}
	roleId := user["role_id"]

	if roleId == fmt.Sprintf("%d", Role_Root_Id) {
		return PrivilegeModel.GetTypePrivileges()
	}
	rolePrivileges := []map[string]string{}
	rolePrivileges, err = RolePrivilegeModel.GetRolePrivilegesByRoleId(roleId)
	if err != nil {
		return
	}
	var privilegeIds = []string{}
	for _, rolePrivilege := range rolePrivileges {
		privilegeIds = append(privilegeIds, rolePrivilege["privilege_id"])
	}
	return PrivilegeModel.GetTypePrivilegesByPrivilegeIds(privilegeIds)
}

func (p *Privilege) GetTypePrivilegesByDisplay(display string) (menus, controllers []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Privilege_Name).Where(map[string]interface{}{
		"is_display": display,
	}).OrderBy("sequence", "ASC"))
	if err != nil {
		return
	}
	menus = []map[string]string{}
	controllers = []map[string]string{}
	for _, row := range rs.Rows() {
		switch row["type"] {
		case Privilege_Type_Menu:
			menus = append(menus, row)
		case Privilege_Type_Controller:
			controllers = append(controllers, row)
		}
	}
	return
}

func (p *Privilege) GetTypePrivilegesByDisplayPrivilegeIds(display string, privilegeIds []string) (menus, controllers []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Privilege_Name).Where(map[string]interface{}{
		"is_display":   display,
		"privilege_id": privilegeIds,
	}).OrderBy("sequence", "ASC"))
	if err != nil {
		return
	}
	menus = []map[string]string{}
	controllers = []map[string]string{}
	for _, row := range rs.Rows() {
		switch row["type"] {
		case Privilege_Type_Menu:
			menus = append(menus, row)
		case Privilege_Type_Controller:
			controllers = append(controllers, row)
		}
	}
	return
}

func (p *Privilege) GetTypePrivilegesByPrivilegeIds(privilegeIds []string) (menus, controllers []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Privilege_Name).Where(map[string]interface{}{
		"privilege_id": privilegeIds,
	}).OrderBy("sequence", "ASC"))
	if err != nil {
		return
	}
	menus = []map[string]string{}
	controllers = []map[string]string{}
	for _, row := range rs.Rows() {
		switch row["type"] {
		case Privilege_Type_Menu:
			menus = append(menus, row)
		case Privilege_Type_Controller:
			controllers = append(controllers, row)
		}
	}
	return
}

func (p *Privilege) GetPrivilegeByPrivilegeId(privilegeId string) (privilege map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Privilege_Name).Where(map[string]interface{}{
		"privilege_id": privilegeId,
	}))
	if err != nil {
		return
	}
	privilege = rs.Row()
	return
}
func (p *Privilege) HasSub(privilegeId string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Privilege_Name).Where(map[string]interface{}{
		"parent_id": privilegeId,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	if rs.Len() > 0 {
		has = true
	}
	return
}
func (p *Privilege) Delete(privilegeId string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Delete(Table_Privilege_Name, map[string]interface{}{
		"privilege_id": privilegeId,
	}))
	if err != nil {
		return
	}
	return
}
func (p *Privilege) Insert(privilege map[string]interface{}) (id int64, err error) {

	privilege["create_time"] = time.Now().Unix()
	privilege["update_time"] = time.Now().Unix()
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Insert(Table_Privilege_Name, privilege))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

func (p *Privilege) InsertNotExists(privilege map[string]interface{}) (id int64, err error) {

	privilege["create_time"] = time.Now().Unix()
	privilege["update_time"] = time.Now().Unix()
	db := G.DB()
	rs, err := db.Query(db.AR().From(Table_Privilege_Name).Where(map[string]interface{}{
		"type":       privilege["type"],
		"controller": privilege["controller"],
		"action":     privilege["action"],
	}))
	if rs.Len() == 0 {
		rs, err = db.Exec(db.AR().Insert(Table_Privilege_Name, privilege))
		if err != nil {
			return
		}
		id = rs.LastInsertId
	}
	return
}

func (p *Privilege) Update(privilegeId string, privilege map[string]interface{}) (id int64, err error) {

	privilege["update_time"] = time.Now().Unix()
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Update(Table_Privilege_Name, privilege, map[string]interface{}{
		"privilege_id": privilegeId,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}
