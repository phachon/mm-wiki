package controllers

import (
	"strings"
	"mm-wiki/app/models"
	"mm-wiki/app/utils"
)

type UserController struct {
	BaseController
}

func (this *UserController) Add() {
	this.viewLayout("user/form", "default")
}

func (this *UserController) Save() {

	username := strings.TrimSpace(this.GetString("username", ""))
	givenName := strings.TrimSpace(this.GetString("given_name", ""))
	email := strings.TrimSpace(this.GetString("email", ""))
	mobile := strings.TrimSpace(this.GetString("mobile", ""))
	phone := strings.TrimSpace(this.GetString("phone", ""))
	department := strings.TrimSpace(this.GetString("department", ""))
	position := strings.TrimSpace(this.GetString("position", ""))
	location := strings.TrimSpace(this.GetString("location", ""))
	im := strings.TrimSpace(this.GetString("im", ""))

	if username == "" {
		this.jsonError("用户名不能为空！")
	}
	if givenName == "" {
		this.jsonError("姓名不能为空！")
	}
	if email == "" {
		this.jsonError("邮箱不能为空！")
	}
	if mobile == "" {
		this.jsonError("手机号不能为空！")
	}

	userId, err := models.UserModel.Insert(map[string]interface{}{
		"username": username,
		"given_name": givenName,
		"email":      email,
		"mobile":     mobile,
		"phone":      phone,
		"department": department,
		"position":   position,
		"location":   location,
		"im":         im,
	})

	if err != nil {
		this.ErrorLog("添加用户失败：" + err.Error())
		this.jsonError("添加用户失败")
	}
	this.InfoLog("添加用户 "+utils.Convert.IntToString(userId, 10)+" 成功")
	this.jsonSuccess("添加用户 "+utils.Convert.IntToString(userId, 10)+" 成功", nil, "/system/user/list")
}

func (this *UserController) List() {

	page, _ := this.GetInt("page", 1)
	keyword := strings.TrimSpace(this.GetString("keyword", ""))

	number := 20
	limit := (page - 1) * number
	var err error
	var count int64
	var users []map[string]string
	if keyword != "" {
		count, err = models.UserModel.CountUsersByKeyword(keyword)
		users, err = models.UserModel.GetUsersByKeywordAndLimit(keyword, limit, number)
	} else {
		count, err = models.UserModel.CountUsers()
		users, err = models.UserModel.GetUsersByLimit(limit, number)
	}
	if err != nil {
		this.ErrorLog("查找用户列表失败: "+err.Error())
		this.ViewError("查找用户列表失败", "/system/main/index")
	}

	this.Data["users"] = users
	this.Data["keyword"] = keyword
	this.SetPaginator(number, count)
	this.viewLayout("user/list", "default")
}

func (this *UserController) Edit() {

	userId := this.GetString("user_id", "")
	if userId == "" {
		this.ViewError("用户不存在", "/system/user/list")
	}

	user, err := models.UserModel.GetUserByUserId(userId)
	if err != nil {
		this.ViewError("用户不存在", "/system/user/list")
	}

	this.Data["user"] = user
	this.ViewLayout("user/form", "default")
}