package controllers

import (
	"fmt"
	"github.com/phachon/mm-wiki/app/models"
	"github.com/phachon/mm-wiki/app/utils"
	"strings"
)

type RoleController struct {
	BaseController
}

func (this *RoleController) Add() {
	this.viewLayout("role/form", "role")
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
		this.ErrorLog("添加角色失败：" + err.Error())
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
	this.InfoLog("添加角色 " + utils.Convert.IntToString(roleId, 10) + " 成功")
	this.jsonSuccess("添加角色成功", nil, "/system/role/list")
}

func (this *RoleController) List() {

	page, _ := this.GetInt("page", 1)
	keyword := strings.TrimSpace(this.GetString("keyword", ""))
	number, _ := this.GetRangeInt("number", 20, 10, 100)
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
		this.ErrorLog("获取角色列表失败: " + err.Error())
		this.ViewError("获取角色列表失败", "/system/main/index")
	}

	this.Data["roles"] = roles
	this.Data["keyword"] = keyword
	this.SetPaginator(number, count)
	this.viewLayout("role/list", "role")
}

func (this *RoleController) Edit() {

	roleId := this.GetString("role_id", "")
	if roleId == "" {
		this.ViewError("角色不存在", "/system/role/list")
	}
	if roleId == fmt.Sprintf("%d", models.Role_Root_Id) {
		this.ViewError("不能修改超级管理员", "/system/role/list")
	}

	role, err := models.RoleModel.GetRoleByRoleId(roleId)
	if err != nil {
		this.ErrorLog("查找角色错误：" + err.Error())
		this.ViewError("查找角色错误", "/system/role/list")
	}
	if len(role) == 0 {
		this.ViewError("角色不存在", "/system/role/list")
	}

	this.Data["role"] = role
	this.viewLayout("role/form", "role")
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
	if roleId == fmt.Sprintf("%d", models.Role_Root_Id) {
		this.jsonError("超级管理员不能修改！")
	}
	if name == "" {
		this.jsonError("角色名称不能为空！")
	}

	role, err := models.RoleModel.GetRoleByRoleId(roleId)
	if err != nil {
		this.ErrorLog("修改角色 " + roleId + " 失败: " + err.Error())
		this.jsonError("修改角色失败！")
	}
	if len(role) == 0 {
		this.jsonError("角色不存在！")
	}
	if role["role_id"] == fmt.Sprintf("%d", models.Role_Root_Id) {
		this.jsonError("超级管理员角色不能修改！")
	}

	ok, _ := models.RoleModel.HasSameName(roleId, name)
	if ok {
		this.jsonError("角色名已经存在！")
	}
	_, err = models.RoleModel.Update(roleId, map[string]interface{}{
		"name": name,
	})

	if err != nil {
		this.ErrorLog("修改角色 " + roleId + " 失败：" + err.Error())
		this.jsonError("修改角色" + roleId + "失败")
	}
	this.InfoLog("修改角色 " + roleId + " 成功")
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

	number, _ := this.GetRangeInt("number", 15, 10, 100)
	limit := (page - 1) * number
	var err error
	var count int64
	var users []map[string]string
	count, err = models.UserModel.CountUsersByKeywords(keywords)
	if err != nil {
		this.ErrorLog("获取角色用户列表失败: " + err.Error())
		this.ViewError("获取角色用户列表失败！", "/system/role/list")
	}
	users, err = models.UserModel.GetUsersByKeywordsAndLimit(keywords, limit, number)
	if err != nil {
		this.ErrorLog("获取用户列表失败: " + err.Error())
		this.ViewError("获取用户列表失败！", "/system/role/list")
	}

	role, err := models.RoleModel.GetRoleByRoleId(roleId)
	if err != nil {
		this.ErrorLog("获取用户列表失败: " + err.Error())
		this.ViewError("获取角色用户列表失败！", "/system/main/index")
	}
	for _, user := range users {
		user["role_name"] = role["name"]
	}

	this.Data["users"] = users
	this.Data["roleId"] = roleId
	this.SetPaginator(number, count)
	this.viewLayout("role/user", "role")
}

func (this *RoleController) Privilege() {

	roleId := this.GetString("role_id", "")
	if roleId == "" {
		this.ViewError("角色不存在", "/system/role/list")
	}

	role, err := models.RoleModel.GetRoleByRoleId(roleId)
	if err != nil {
		this.ErrorLog("查找角色权限失败：" + err.Error())
		this.ViewError("查看角色权限失败！", "/system/role/list")
	}
	if len(role) == 0 {
		this.ViewError("角色不存在", "/system/role/list")
	}

	menus, controllers, err := models.PrivilegeModel.GetTypePrivileges()
	if err != nil {
		this.ViewError("查找角色权限失败！")
	}

	var rolePrivileges = []map[string]string{}
	if role["role_id"] == fmt.Sprintf("%d", models.Role_Root_Id) {
		rolePrivileges, err = models.RolePrivilegeModel.GetRootRolePrivileges()
	} else {
		rolePrivileges, err = models.RolePrivilegeModel.GetRolePrivilegesByRoleId(roleId)
	}
	if err != nil {
		this.ViewError("查找用户权限出错")
	}

	this.Data["role"] = role
	this.Data["menus"] = menus
	this.Data["controllers"] = controllers
	this.Data["rolePrivileges"] = rolePrivileges
	this.Data["disabledPrivilegeIds"] = models.Privilege_Default_Ids

	this.viewLayout("role/privilege", "role")
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
	//if len(privilegeIds) == 0 {
	//	this.jsonError("没有选择权限!")
	//}
	// add default privileges
	privilegeIds = append(privilegeIds, models.Privilege_Default_Ids...)

	role, err := models.RoleModel.GetRoleByRoleId(roleId)
	if err != nil {
		this.ErrorLog("角色 " + roleId + " 授权失败：" + err.Error())
		this.jsonError("角色不存在")
	}
	if len(role) == 0 {
		this.jsonError("角色不存在")
	}

	if role["role_id"] == fmt.Sprintf("%d", models.Role_Root_Id) {
		this.jsonError("超级管理员不需要授权！")
	}

	res, err := models.RolePrivilegeModel.GrantRolePrivileges(roleId, privilegeIds)
	if err != nil {
		this.ErrorLog("角色 " + roleId + " 授权失败：" + err.Error())
		this.jsonError("角色授权失败！")
	}
	if !res {
		this.jsonError("角色授权失败")
	}

	this.InfoLog("角色 " + roleId + " 授权成功")
	this.jsonSuccess("角色授权成功", nil, "/system/role/list")
}

func (this *RoleController) Delete() {
	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/role/list")
	}

	roleId := this.GetString("role_id", "")
	if roleId == "" {
		this.jsonError("没有选择角色！")
	}
	if roleId == fmt.Sprintf("%d", models.Role_Root_Id) {
		this.jsonError("超级管理员不能删除！")
	}

	role, err := models.RoleModel.GetRoleByRoleId(roleId)
	if err != nil {
		this.ErrorLog("删除角色 " + roleId + " 失败: " + err.Error())
		this.jsonError("删除角色失败")
	}
	if len(role) == 0 {
		this.jsonError("角色不存在")
	}
	if role["type"] == fmt.Sprintf("%d", models.Role_Type_System) {
		this.jsonError("系统角色不能删除！")
	}

	// check role user
	users, err := models.UserModel.GetUsersByRoleId(roleId)
	if err != nil {
		this.ErrorLog("删除角色 " + roleId + " 失败: " + err.Error())
		this.jsonError("删除角色失败")
	}
	if len(users) > 0 {
		this.jsonError("不能删除角色，请先移除该角色下用户!")
	}

	// delete role privilege by role id
	err = models.RolePrivilegeModel.DeleteByRoleId(roleId)
	if err != nil {
		this.ErrorLog("删除角色 " + roleId + " 权限失败: " + err.Error())
		this.jsonError("删除角色失败")
	}

	// delete role by role id
	err = models.RoleModel.Delete(roleId)
	if err != nil {
		this.ErrorLog("删除角色 " + roleId + " 失败: " + err.Error())
		this.jsonError("删除角色失败")
	}

	this.InfoLog("删除角色 " + roleId + " 成功")
	this.jsonSuccess("删除角色成功", nil, "/system/role/list")
}

func (this *RoleController) ResetUser() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/role/list")
	}
	userId := this.GetString("user_id", "")
	if userId == "" {
		this.jsonError("用户不存在")
	}

	if this.UserId == "1" {
		this.jsonError("root 用户不能重置角色！")
	}

	user, err := models.UserModel.GetUserByUserId(userId)
	if err != nil {
		this.ErrorLog("重置用户 " + userId + " 角色失败: " + err.Error())
		this.jsonError("重置用户角色失败")
	}
	if len(user) == 0 {
		this.jsonError("用户不存在")
	}

	_, err = models.UserModel.Update(userId, map[string]interface{}{
		"role_id": models.Role_Default_Id,
	})
	if err != nil {
		this.ErrorLog("重置用户 " + userId + " 角色失败: " + err.Error())
		this.jsonError("重置用户角色失败")
	}

	this.InfoLog("重置用户 " + userId + " 角色成功")
	this.jsonSuccess("重置用户角色成功", nil, "/system/role/list")
}
