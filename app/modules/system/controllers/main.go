package controllers

import (
	"fmt"
	"github.com/phachon/mm-wiki/app/models"
)

type MainController struct {
	BaseController
}

func (this *MainController) Index() {

	var err error
	var menus = []map[string]string{}
	var controllers = []map[string]string{}

	user, err := models.UserModel.GetUserByUserId(this.UserId)
	if err != nil {
		this.ViewError("查找用户失败！")
	}
	if len(user) == 0 {
		this.ViewError("用户不存在！")
	}
	roleId := user["role_id"]

	if roleId == fmt.Sprintf("%d", models.Role_Root_Id) {
		menus, controllers, err = models.PrivilegeModel.GetTypePrivilegesByDisplay(models.Privilege_DisPlay_True)
		if err != nil {
			this.ViewError("查找用户权限失败！")
		}
	} else {
		rolePrivileges, err := models.RolePrivilegeModel.GetRolePrivilegesByRoleId(roleId)
		if err != nil {
			this.ViewError("查找用户权限出错")
		}
		var privilegeIds = []string{}
		for _, rolePrivilege := range rolePrivileges {
			privilegeIds = append(privilegeIds, rolePrivilege["privilege_id"])
		}
		menus, controllers, err = models.PrivilegeModel.GetTypePrivilegesByDisplayPrivilegeIds(models.Privilege_DisPlay_True, privilegeIds)
		if err != nil {
			this.ViewError("查找用户权限失败！")
		}
	}
	this.Data["menus"] = menus
	this.Data["controllers"] = controllers
	this.viewLayout("main/index", "main")
}

func (this *MainController) Default() {

	this.viewLayout("main/default", "default")
}
