package controllers

import (
	"mm-wiki/app/models"
)

type MainController struct {
	BaseController
}

func (this *MainController) Index() {

	menus, controllers, err := models.PrivilegeModel.GetTypePrivilegesByDisplay(models.Privilege_DisPlay_True)
	if err != nil {
		this.ViewError("查找用户权限失败！")
	}

	this.Data["menus"] = menus
	this.Data["controllers"] = controllers
	this.viewLayout("main/index", "main")
}

func (this *MainController) Default() {
	this.viewLayout("main/default", "default")
}