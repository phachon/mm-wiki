package controllers

import (
	"strings"
	"mm-wiki/app/models"
	"mm-wiki/app/utils"
	"fmt"
)

type RoleController struct {
	BaseController
}

func (this *RoleController) Add() {
	this.viewLayout("role/form", "default")
}

func (this *RoleController) Save() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/role/list")
	}
	name := strings.TrimSpace(this.GetString("name", ""))
	if name == "" {
		this.jsonError("角色名称不能为空！")
	}

	ok, err := models.RoleModel.HasRoleName(name)
	if err != nil {
		this.ErrorLog("添加角色失败："+err.Error())
		this.jsonError("添加角色失败！")
	}
	if ok {
		this.jsonError("角色名已经存在！")
	}

	roleId, err := models.RoleModel.Insert(map[string]interface{}{
		"name": name,
	})

	if err != nil {
		this.ErrorLog("添加角色失败：" + err.Error())
		this.jsonError("添加角色失败")
	}
	this.InfoLog("添加角色 "+utils.Convert.IntToString(roleId, 10)+" 成功")
	this.jsonSuccess("添加角色成功", nil, "/system/role/list")
}

func (this *RoleController) List() {

	page, _ := this.GetInt("page", 1)
	keyword := strings.TrimSpace(this.GetString("keyword", ""))

	number := 20
	limit := (page - 1) * number
	var err error
	var count int64
	var roles []map[string]string
	if keyword != "" {
		count, err = models.RoleModel.CountRolesByKeyword(keyword)
		roles, err = models.RoleModel.GetRolesByKeywordAndLimit(keyword, limit, number)
	} else {
		count, err = models.RoleModel.CountRoles()
		roles, err = models.RoleModel.GetRolesByLimit(limit, number)
	}
	if err != nil {
		this.ErrorLog("获取角色列表失败: "+err.Error())
		this.ViewError("获取角色列表失败", "/system/main/index")
	}

	this.Data["roles"] = roles
	this.Data["keyword"] = keyword
	this.SetPaginator(number, count)
	this.viewLayout("role/list", "default")
}

func (this *RoleController) Edit() {

	roleId := this.GetString("role_id", "")
	if roleId == "" {
		this.ViewError("角色不存在", "/system/role/list")
	}

	role, err := models.RoleModel.GetRoleByRoleId(roleId)
	if err != nil {
		this.ViewError("角色不存在", "/system/role/list")
	}

	this.Data["role"] = role
	this.viewLayout("role/form", "default")
}

func (this *RoleController) Modify() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/role/list")
	}
	roleId := this.GetString("role_id", "")
	name := strings.TrimSpace(this.GetString("name", ""))

	if roleId == "" {
		this.jsonError("角色不存在！")
	}
	if name == "" {
		this.jsonError("角色名称不能为空！")
	}

	role, err := models.RoleModel.GetRoleByRoleId(roleId)
	if err != nil {
		this.ErrorLog("修改角色 "+roleId+" 失败: "+err.Error())
		this.jsonError("修改角色失败！")
	}
	if len(role) == 0 {
		this.jsonError("角色不存在！")
	}
	if role["role_id"] == fmt.Sprintf("%d", models.Role_Root_Id) {
		this.jsonError("超级管理员角色不能修改！")
	}

	ok , _ := models.RoleModel.HasSameName(roleId, name)
	if ok {
		this.jsonError("角色名已经存在！")
	}
	_, err = models.RoleModel.Update(roleId, map[string]interface{}{
		"name": name,
	})

	if err != nil {
		this.ErrorLog("修改角色 "+roleId+" 失败：" + err.Error())
		this.jsonError("修改角色"+roleId+"失败")
	}
	this.InfoLog("修改角色 "+roleId+" 成功")
	this.jsonSuccess("修改角色成功", nil, "/system/role/list")
}

func (this *RoleController) User() {

	keywords := map[string]string{}
	page, _ := this.GetInt("page", 1)
	roleId := strings.TrimSpace(this.GetString("role_id", ""))

	if roleId == "" {
		this.ViewError("没有选择角色！")
	}
	keywords["role_id"] = roleId

	number := 20
	limit := (page - 1) * number
	var err error
	var count int64
	var users []map[string]string
	count, err = models.UserModel.CountUsersByKeywords(keywords)
	if err != nil {
		this.ErrorLog("获取角色用户列表失败: "+err.Error())
		this.ViewError("获取角色用户列表失败！", "/system/role/list")
	}
	users, err = models.UserModel.GetUsersByKeywordsAndLimit(keywords, limit, number)
	if err != nil {
		this.ErrorLog("获取用户列表失败: "+err.Error())
		this.ViewError("获取用户列表失败！", "/system/role/list")
	}

	role, err := models.RoleModel.GetRoleByRoleId(roleId)
	if err != nil {
		this.ErrorLog("获取用户列表失败: "+err.Error())
		this.ViewError("获取角色用户列表失败！", "/system/main/index")
	}
	for _, user := range users {
		user["role_name"] = role["name"]
	}

	this.Data["users"] = users
	this.Data["roleId"] = roleId
	this.SetPaginator(number, count)
	this.viewLayout("role/user", "default")
}

func (this *RoleController) Privilege() {

	roleId := this.GetString("role_id", "")
	if roleId == "" {
		this.ViewError("角色不存在", "/system/role/list")
	}

	role, err := models.RoleModel.GetRoleByRoleId(roleId)
	if err != nil {
		this.ErrorLog("查找角色权限失败："+err.Error())
		this.ViewError("查看角色权限失败！", "/system/role/list")
	}
	if len(role) == 0 {
		this.ViewError("角色不存在", "/system/role/list")
	}

	menus, controllers, err := models.PrivilegeModel.GetTypePrivileges()
	if err != nil {
		this.ViewError("查找角色权限失败！")
	}

	rolePrivileges, err := models.RolePrivilegeModel.GetRolePrivilegesByRoleId(roleId)
	if err != nil {
		this.ViewError("查找用户权限出错")
	}

	this.Data["role"] = role
	this.Data["menus"] = menus
	this.Data["controllers"] = controllers
	this.Data["rolePrivileges"] = rolePrivileges

	this.viewLayout("role/privilege", "default")
}

func (this *RoleController) GrantPrivilege() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/role/list")
	}
	privilegeIds := this.GetStrings("privilege_id", []string{})
	roleId := this.GetString("role_id", "")

	if roleId == "" {
		this.jsonError("没有选择角色!")
	}
	if len(privilegeIds) == 0 {
		this.jsonError("没有选择权限!")
	}

	role, err := models.RoleModel.GetRoleByRoleId(roleId)
	if err != nil {
		this.ErrorLog("角色 "+roleId+" 授权失败："+err.Error())
		this.jsonError("角色不存在")
	}
	if len(role) == 0 {
		this.jsonError("角色不存在")
	}

	res, err := models.RolePrivilegeModel.GrantRolePrivileges(roleId, privilegeIds)
	if err != nil {
		this.ErrorLog("角色 "+roleId+" 授权失败："+err.Error())
		this.jsonError("角色授权失败！")
	}
	if !res {
		this.jsonError("角色授权失败")
	}

	this.InfoLog("角色 "+roleId+" 授权成功")
	this.jsonSuccess("角色授权成功", nil, "/system/role/list")
}