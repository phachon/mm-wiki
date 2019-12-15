package controllers

import (
	"github.com/astaxie/beego"
	"github.com/phachon/mm-wiki/app/models"
	"github.com/phachon/mm-wiki/app/utils"
	"strings"
)

type PrivilegeController struct {
	BaseController
}

func (this *PrivilegeController) Add() {

	menus, _, err := models.PrivilegeModel.GetTypePrivileges()
	if err != nil {
		this.ViewError("查找权限菜单失败！")
	}

	this.Data["menus"] = menus
	this.Data["mode"] = beego.BConfig.RunMode
	this.viewLayout("privilege/form", "privilege")
}

func (this *PrivilegeController) Save() {

	if !this.IsPost() {
		this.ViewError("请求方式错误", "/system/privilege/add")
	}
	if beego.BConfig.RunMode != "dev" {
		this.jsonError("只允许在开发模式下添加权限!")
	}

	name := strings.TrimSpace(this.GetString("name", ""))
	privilegeType := strings.TrimSpace(this.GetString("type", ""))
	parentId := strings.TrimSpace(this.GetString("parent_id", "0"))
	controller := strings.TrimSpace(this.GetString("controller", ""))
	action := strings.TrimSpace(this.GetString("action", ""))
	target := strings.TrimSpace(this.GetString("target", ""))
	icon := strings.TrimSpace(this.GetString("icon", ""))
	isDisplay := strings.TrimSpace(this.GetString("is_display", "0"))
	sequence := strings.TrimSpace(this.GetString("sequence", "0"))

	if name == "" {
		this.jsonError("权限名称不能为空！")
	}
	if privilegeType == "" {
		this.jsonError("没有选择权限类型！")
	}
	if privilegeType == "controller" {
		if parentId == "" {
			this.jsonError("控制器必须选择上级菜单！")
		}
		if controller == "" {
			this.jsonError("控制器名称不能为空！")
		}
		if action == "" {
			this.jsonError("方法名称不能为空！")
		}
	}

	privilegeId, err := models.PrivilegeModel.Insert(map[string]interface{}{
		"name":       name,
		"type":       privilegeType,
		"parent_id":  parentId,
		"controller": controller,
		"action":     action,
		"target":     target,
		"icon":       icon,
		"is_display": isDisplay,
		"sequence":   sequence,
	})

	if err != nil {
		this.ErrorLog("添加权限失败：" + err.Error())
		this.jsonError("添加权限失败")
	}
	this.InfoLog("添加权限 " + utils.Convert.IntToString(privilegeId, 10) + " 成功")
	this.jsonSuccess("添加权限成功", nil, "/system/privilege/list")
}

func (this *PrivilegeController) List() {

	menus, controllers, err := models.PrivilegeModel.GetTypePrivileges()
	if err != nil {
		this.ViewError("查找权限失败！")
	}

	this.Data["menus"] = menus
	this.Data["controllers"] = controllers
	this.Data["mode"] = beego.BConfig.RunMode
	this.viewLayout("privilege/list", "privilege")
}

func (this *PrivilegeController) Edit() {

	privilegeId := this.GetString("privilege_id", "")
	if privilegeId == "" {
		this.ViewError("没有选择权限！", "/system/privilege/list")
	}

	privilege, err := models.PrivilegeModel.GetPrivilegeByPrivilegeId(privilegeId)
	if err != nil {
		this.ErrorLog("查找权限失败：" + err.Error())
		this.ViewError("查找权限失败！", "/system/privilege/list")
	}
	if len(privilege) == 0 {
		this.ViewError("权限不存在！", "/system/privilege/list")
	}

	menus, _, err := models.PrivilegeModel.GetTypePrivileges()
	if err != nil {
		this.ViewError("查找权限失败！", "/system/privilege/list")
	}

	this.Data["menus"] = menus
	this.Data["privilege"] = privilege
	this.Data["mode"] = beego.BConfig.RunMode
	this.viewLayout("privilege/form", "privilege")
}

func (this *PrivilegeController) Modify() {

	if !this.IsPost() {
		this.ViewError("请求方式错误", "/system/privilege/list")
	}
	if beego.BConfig.RunMode != "dev" {
		this.jsonError("只允许在开发模式下修改权限!")
	}
	privilegeId := strings.TrimSpace(this.GetString("privilege_id", ""))
	name := strings.TrimSpace(this.GetString("name", ""))
	privilegeType := strings.TrimSpace(this.GetString("type", ""))
	parentId := strings.TrimSpace(this.GetString("parent_id", "0"))
	controller := strings.TrimSpace(this.GetString("controller", ""))
	action := strings.TrimSpace(this.GetString("action", ""))
	target := strings.TrimSpace(this.GetString("target", ""))
	icon := strings.TrimSpace(this.GetString("icon", "glyphicon-list"))
	isDisplay := strings.TrimSpace(this.GetString("is_display", "0"))
	sequence := strings.TrimSpace(this.GetString("sequence", "0"))

	if name == "" {
		this.jsonError("权限名称不能为空！")
	}
	if privilegeType == "" {
		this.jsonError("没有选择权限类型！")
	}
	if privilegeType == "controller" {
		if parentId == "" {
			this.jsonError("控制器必须选择上级菜单！")
		}
		if controller == "" {
			this.jsonError("控制器名称不能为空！")
		}
		if action == "" {
			this.jsonError("方法名称不能为空！")
		}
	}

	_, err := models.PrivilegeModel.Update(privilegeId, map[string]interface{}{
		"name":       name,
		"type":       privilegeType,
		"parent_id":  parentId,
		"controller": controller,
		"action":     action,
		"target":     target,
		"icon":       icon,
		"is_display": isDisplay,
		"sequence":   sequence,
	})

	if err != nil {
		this.ErrorLog("修改权限 " + privilegeId + " 失败：" + err.Error())
		this.jsonError("修改权限失败！")
	}
	this.InfoLog("修改权限 " + privilegeId + " 成功")
	this.jsonSuccess("修改权限成功", nil, "/system/privilege/list")
}

func (this *PrivilegeController) Delete() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/privilege/list")
	}
	if beego.BConfig.RunMode != "dev" {
		this.jsonError("只允许在开发模式下删除权限!")
	}
	privilegeId := this.GetString("privilege_id", "")
	if privilegeId == "" {
		this.jsonError("没有选择权限！")
	}

	privilege, err := models.PrivilegeModel.GetPrivilegeByPrivilegeId(privilegeId)
	if err != nil {
		this.ErrorLog("删除权限 " + privilegeId + " 失败: " + err.Error())
		this.jsonError("删除权限失败")
	}
	if len(privilege) == 0 {
		this.jsonError("权限不存在")
	}

	// delete role_privilege by privilegeId
	err = models.RolePrivilegeModel.DeleteByPrivilegeId(privilegeId)
	if err != nil {
		this.ErrorLog("删除角色权限 " + privilegeId + " 失败: " + err.Error())
		this.jsonError("删除权限失败")
	}

	// delete privilege
	err = models.PrivilegeModel.Delete(privilegeId)
	if err != nil {
		this.ErrorLog("删除权限 " + privilegeId + " 失败: " + err.Error())
		this.jsonError("删除权限失败")
	}

	this.InfoLog("删除权限 " + privilegeId + " 成功")
	this.jsonSuccess("删除权限成功", nil, "/system/privilege/list")
}
