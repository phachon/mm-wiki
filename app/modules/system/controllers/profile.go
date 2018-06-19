package controllers

import (
	"strings"
	"mm-wiki/app/models"
)

type ProfileController struct {
	BaseController
}

func (this *ProfileController) Info() {
	this.Data["user"] = this.User
	this.viewLayout("profile/form", "default")
}

func (this *ProfileController) Save() {

	givenName := strings.TrimSpace(this.GetString("given_name", ""))
	email := strings.TrimSpace(this.GetString("email", ""))
	mobile := strings.TrimSpace(this.GetString("mobile", ""))
	phone := strings.TrimSpace(this.GetString("phone", ""))
	department := strings.TrimSpace(this.GetString("department", ""))
	position := strings.TrimSpace(this.GetString("position", ""))
	location := strings.TrimSpace(this.GetString("location", ""))
	im := strings.TrimSpace(this.GetString("im", ""))

	if givenName == "" {
		this.jsonError("姓名不能为空！")
	}
	if email == "" {
		this.jsonError("邮箱不能为空！")
	}
	if mobile == "" {
		this.jsonError("手机号不能为空！")
	}

	_, err := models.UserModel.Update(this.UserId, map[string]interface{}{
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
		this.ErrorLog("修改个人资料失败：" + err.Error())
		this.jsonError("修改失败")
	}
	this.InfoLog("修改个人资料成功")
	this.jsonSuccess("我的资料修改成功", nil, "/profile/info")
}