package controllers

import (
	"strings"
	"mm-wiki/app/models"
)

type UserController struct {
	BaseController
}

func (this *UserController) Index() {
	this.viewLayout("user/index", "user")
}

func (this *UserController) List() {

	keywords := map[string]string{}
	page, _ := this.GetInt("page", 1)
	username := strings.TrimSpace(this.GetString("username", ""))
	roleId := strings.TrimSpace(this.GetString("role_id", ""))
	if username != "" {
		keywords["username"] = username
	}
	if roleId != "" {
		keywords["role_id"] = roleId
	}

	number := 20
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
		this.ErrorLog("获取用户列表失败: "+err.Error())
		this.ViewError("获取用户列表失败", "/system/main/index")
	}

	var roleIds = []string{}
	if roleId != "" {
		roleIds = append(roleIds, roleId)
	}else {
		for _, user := range users {
			roleIds = append(roleIds, user["role_id"])
		}
	}
	roles, err := models.RoleModel.GetRoleByRoleIds(roleIds)
	if err != nil {
		this.ErrorLog("获取用户列表失败: "+err.Error())
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
		this.ErrorLog("获取用户列表失败: "+err.Error())
		this.ViewError("获取用户列表失败！", "/system/main/index")
	}
	this.Data["users"] = roleUsers
	this.Data["username"] = username
	this.Data["roleId"] = roleId
	this.Data["roles"] = allRoles
	this.SetPaginator(number, count)
	this.viewLayout("user/list", "default")
}