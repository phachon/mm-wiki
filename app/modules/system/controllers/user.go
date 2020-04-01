package controllers

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/phachon/mm-wiki/app/models"
	"github.com/phachon/mm-wiki/app/utils"
	"strings"
)

type UserController struct {
	BaseController
}

func (this *UserController) Add() {

	roles := []map[string]string{}
	var err error

	if this.IsRoot() {
		roles, err = models.RoleModel.GetRoles()
	} else {
		roles, err = models.RoleModel.GetRolesNotContainRoot()
	}
	if err != nil {
		this.ErrorLog("获取用户角色失败：" + err.Error())
		this.ViewError("获取用户角色失败！")
	}
	this.Data["roles"] = roles
	this.viewLayout("user/form", "user")
}

func (this *UserController) Save() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/user/add")
	}
	username := strings.TrimSpace(this.GetString("username", ""))
	givenName := strings.TrimSpace(this.GetString("given_name", ""))
	password := strings.TrimSpace(this.GetString("password", ""))
	email := strings.TrimSpace(this.GetString("email", ""))
	mobile := strings.TrimSpace(this.GetString("mobile", ""))
	roleId := strings.TrimSpace(this.GetString("role_id", ""))
	phone := strings.TrimSpace(this.GetString("phone", ""))
	department := strings.TrimSpace(this.GetString("department", ""))
	position := strings.TrimSpace(this.GetString("position", ""))
	location := strings.TrimSpace(this.GetString("location", ""))
	im := strings.TrimSpace(this.GetString("im", ""))
	this.Ctx.Request.PostForm.Del("password")

	v := validation.Validation{}
	if username == "" {
		this.jsonError("用户名不能为空！")
	}
	if !v.AlphaNumeric(username, "username").Ok {
		this.jsonError("用户名格式不正确！")
	}
	if givenName == "" {
		this.jsonError("姓名不能为空！")
	}
	if password == "" {
		this.jsonError("密码不能为空！")
	}
	if email == "" {
		this.jsonError("邮箱不能为空！")
	}
	if !v.Email(email, "email").Ok {
		this.jsonError("邮箱格式不正确！")
	}
	if mobile == "" {
		this.jsonError("手机号不能为空！")
	}
	//if !v.Mobile(mobile, "mobile").Ok {
	//	this.jsonError("手机号格式不正确！")
	//}
	if roleId == "" {
		this.jsonError("没有选择角色！")
	}
	//if phone != "" && !v.Phone(phone, "phone").Ok {
	//	this.jsonError("电话格式不正确！")
	//}

	ok, err := models.UserModel.HasUsername(username)
	if err != nil {
		this.ErrorLog("添加用户失败：" + err.Error())
		this.jsonError("添加用户失败！")
	}
	if ok {
		this.jsonError("用户名已经存在！")
	}

	if !this.IsRoot() && roleId == fmt.Sprintf("%d", models.Role_Root_Id) {
		this.jsonError("没有权限添加超级管理员！")
	}

	userId, err := models.UserModel.Insert(map[string]interface{}{
		"username":   username,
		"given_name": givenName,
		"password":   models.UserModel.EncodePassword(password),
		"email":      email,
		"mobile":     mobile,
		"phone":      phone,
		"department": department,
		"position":   position,
		"location":   location,
		"im":         im,
		"role_id":    roleId,
	})

	if err != nil {
		this.ErrorLog("添加用户失败：" + err.Error())
		this.jsonError("添加用户失败")
	}
	this.InfoLog("添加用户 " + utils.Convert.IntToString(userId, 10) + " 成功")
	this.jsonSuccess("添加用户成功", nil, "/system/user/list")
}

func (this *UserController) List() {

	keywords := map[string]string{}
	page, _ := this.GetInt("page", 1)
	username := strings.TrimSpace(this.GetString("username", ""))
	roleId := strings.TrimSpace(this.GetString("role_id", ""))
	number, _ := this.GetRangeInt("number", 20, 10, 100)

	if username != "" {
		keywords["username"] = username
	}
	if roleId != "" {
		keywords["role_id"] = roleId
	}

	limit := (page - 1) * number
	var err error
	var count int64
	var users []map[string]string
	if len(keywords) != 0 {
		count, err = models.UserModel.CountUsersByKeywords(keywords)
		users, err = models.UserModel.GetUsersByKeywordsAndLimit(keywords, limit, number)
	} else {
		count, err = models.UserModel.CountUsers()
		users, err = models.UserModel.GetUsersByLimit(limit, number)
	}
	if err != nil {
		this.ErrorLog("获取用户列表失败: " + err.Error())
		this.ViewError("获取用户列表失败", "/system/main/index")
	}

	var roleIds = []string{}
	if roleId != "" {
		roleIds = append(roleIds, roleId)
	} else {
		for _, user := range users {
			roleIds = append(roleIds, user["role_id"])
		}
	}
	roles, err := models.RoleModel.GetRoleByRoleIds(roleIds)
	if err != nil {
		this.ErrorLog("获取用户列表失败: " + err.Error())
		this.ViewError("获取用户列表失败!", "/system/main/index")
	}
	var roleUsers = []map[string]string{}
	for _, user := range users {
		roleUser := user
		for _, role := range roles {
			if role["role_id"] == user["role_id"] {
				roleUser["role_name"] = role["name"]
				break
			}
		}
		roleUsers = append(roleUsers, roleUser)
	}

	allRoles, err := models.RoleModel.GetRoles()
	if err != nil {
		this.ErrorLog("获取用户列表失败: " + err.Error())
		this.ViewError("获取用户列表失败！", "/system/main/index")
	}
	this.Data["users"] = roleUsers
	this.Data["username"] = username
	this.Data["roleId"] = roleId
	this.Data["roles"] = allRoles
	this.SetPaginator(number, count)
	this.viewLayout("user/list", "user")
}

func (this *UserController) Edit() {

	userId := this.GetString("user_id", "")
	if userId == "" {
		this.ViewError("用户不存在！", "/system/user/list")
	}

	user, err := models.UserModel.GetUserByUserId(userId)
	if err != nil {
		this.ErrorLog("查找用户出错：" + err.Error())
		this.ViewError("查找用户出错！", "/system/user/list")
	}
	if len(user) == 0 {
		this.ViewError("用户不存在！", "/system/user/list")
	}
	// 登录非 root 用户不能修改 root 用户信息
	if user["role_id"] == fmt.Sprintf("%d", models.Role_Root_Id) && !this.IsRoot() {
		this.ViewError("没有权限修改！", "/system/user/list")
	}

	roles := []map[string]string{}
	if this.IsRoot() {
		roles, err = models.RoleModel.GetRoles()
	} else {
		roles, err = models.RoleModel.GetRolesNotContainRoot()
	}
	if err != nil {
		this.ErrorLog("获取用户角色失败：" + err.Error())
		this.ViewError("获取用户角色失败！")
	}

	this.Data["user"] = user
	this.Data["roles"] = roles
	this.viewLayout("user/edit", "user")
}

func (this *UserController) Modify() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/user/list")
	}
	userId := strings.TrimSpace(this.GetString("user_id", ""))
	givenName := strings.TrimSpace(this.GetString("given_name", ""))
	email := strings.TrimSpace(this.GetString("email", ""))
	mobile := strings.TrimSpace(this.GetString("mobile", ""))
	roleId := strings.TrimSpace(this.GetString("role_id", ""))
	phone := strings.TrimSpace(this.GetString("phone", ""))
	department := strings.TrimSpace(this.GetString("department", ""))
	position := strings.TrimSpace(this.GetString("position", ""))
	location := strings.TrimSpace(this.GetString("location", ""))
	im := strings.TrimSpace(this.GetString("im", ""))
	password := strings.TrimSpace(this.GetString("password", ""))
	this.Ctx.Request.PostForm.Del("password")

	v := validation.Validation{}
	if givenName == "" {
		this.jsonError("姓名不能为空！")
	}
	if email == "" {
		this.jsonError("邮箱不能为空！")
	}
	if !v.Email(email, "email").Ok {
		this.jsonError("邮箱格式不正确！")
	}
	if mobile == "" {
		this.jsonError("手机号不能为空！")
	}
	//if !v.Mobile(mobile, "mobile").Ok {
	//	this.jsonError("手机号格式不正确！")
	//}
	//if roleId == "" {
	//	this.jsonError("没有选择角色！")
	//}
	//if phone != "" && !v.Phone(phone, "phone").Ok {
	//	this.jsonError("电话格式不正确！")
	//}

	user, err := models.UserModel.GetUserByUserId(userId)
	if err != nil {
		this.ErrorLog("修改用户 " + userId + " 失败：" + err.Error())
		this.jsonError("修改用户出错！")
	}
	if len(user) == 0 {
		this.jsonError("用户不存在！")
	}
	if user["role_id"] == fmt.Sprintf("%d", models.Role_Root_Id) {
		roleId = fmt.Sprintf("%d", models.Role_Root_Id)
	}
	// 登录非 root 用户不能修改 root 用户信息
	if user["role_id"] == fmt.Sprintf("%d", models.Role_Root_Id) && !this.IsRoot() {
		this.jsonError("没有权限修改！")
	}

	updateUser := map[string]interface{}{
		"given_name": givenName,
		"email":      email,
		"mobile":     mobile,
		"phone":      phone,
		"department": department,
		"position":   position,
		"location":   location,
		"im":         im,
	}
	// 超级管理员才可以修改其他用户密码
	if password != "" && this.IsRoot() {
		updateUser["password"] = models.UserModel.EncodePassword(password)
	}
	if roleId != "" {
		updateUser["role_id"] = roleId
	}
	_, err = models.UserModel.Update(userId, updateUser)
	if err != nil {
		this.ErrorLog("修改用户 " + userId + " 失败：" + err.Error())
		this.jsonError("修改用户失败")
	}
	this.InfoLog("修改用户 " + userId + " 成功")
	this.jsonSuccess("修改用户成功", nil, "/system/user/list")
}

func (this *UserController) Forbidden() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/user/list")
	}
	userId := this.GetString("user_id", "")
	if userId == "" {
		this.jsonError("用户不存在")
	}

	user, err := models.UserModel.GetUserByUserId(userId)
	if err != nil {
		this.ErrorLog("屏蔽用户 " + userId + " 失败: " + err.Error())
		this.jsonError("屏蔽用户失败")
	}
	if len(user) == 0 {
		this.jsonError("用户不存在")
	}
	if user["role_id"] == fmt.Sprintf("%d", models.Role_Root_Id) {
		this.jsonError("不能操作超级管理员")
	}
	_, err = models.UserModel.Update(userId, map[string]interface{}{
		"is_forbidden": models.User_Forbidden_True,
	})
	if err != nil {
		this.ErrorLog("屏蔽用户 " + userId + " 失败: " + err.Error())
		this.jsonError("屏蔽用户失败")
	}

	this.InfoLog("屏蔽用户 " + userId + " 成功")
	this.jsonSuccess("屏蔽用户成功", nil, "/system/user/list")
}

func (this *UserController) Recover() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/user/list")
	}
	userId := this.GetString("user_id", "")
	if userId == "" {
		this.jsonError("用户不存在")
	}

	user, err := models.UserModel.GetUserByUserId(userId)
	if err != nil {
		this.ErrorLog("恢复用户 " + userId + " 失败: " + err.Error())
		this.jsonError("恢复用户失败")
	}
	if len(user) == 0 {
		this.jsonError("用户不存在")
	}
	if user["role_id"] == fmt.Sprintf("%d", models.Role_Root_Id) {
		this.jsonError("不能操作超级管理员")
	}
	_, err = models.UserModel.Update(userId, map[string]interface{}{
		"is_forbidden": models.User_Is_Forbidden_False,
	})
	if err != nil {
		this.ErrorLog("恢复用户 " + userId + " 失败: " + err.Error())
		this.jsonError("恢复用户失败")
	}

	this.InfoLog("恢复用户 " + userId + " 成功")
	this.jsonSuccess("恢复用户成功", nil, "/system/user/list")
}

func (this *UserController) Info() {

	userId := this.GetString("user_id", "")
	if userId == "" {
		this.ViewError("用户不存在！", "/system/user/list")
	}

	user, err := models.UserModel.GetUserByUserId(userId)
	if err != nil {
		this.ErrorLog("查找用户出错：" + err.Error())
		this.ViewError("查找用户出错！", "/system/user/list")
	}
	if len(user) == 0 {
		this.ViewError("用户不存在！", "/system/user/list")
	}
	role, err := models.RoleModel.GetRoleByRoleId(user["role_id"])
	if err != nil {
		this.ErrorLog("查找用户角色出错：" + err.Error())
		this.ViewError("查找用户出错！", "/system/user/list")
	}
	this.Data["user"] = user
	this.Data["role"] = role
	this.viewLayout("user/info", "user")
}
