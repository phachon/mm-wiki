package controllers

import (
	"strings"
	"mm-wiki/app/models"
)

type ProfileController struct {
	BaseController
}

func (this *ProfileController) Info() {

	user, err := models.UserModel.GetUserByUserId(this.UserId)
	if err != nil {
		this.ErrorLog("获取我的资料失败: "+err.Error())
		this.ViewError("获取资料失败")
	}
	this.Data["user"] = user
	this.viewLayout("profile/info", "default")
}

func (this *ProfileController) Edit() {

	user, err := models.UserModel.GetUserByUserId(this.UserId)
	if err != nil {
		this.ErrorLog("获取我的资料失败: "+err.Error())
		this.ViewError("获取资料失败")
	}
	this.Data["user"] = user
	this.viewLayout("profile/edit", "default")
}

func (this *ProfileController) Modify() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/profile/info")
	}
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
		this.ErrorLog("修改我的资料失败：" + err.Error())
		this.jsonError("修改我的资料失败")
	}
	this.InfoLog("修改我的资料成功")
	this.jsonSuccess("我的资料修改成功", nil, "/system/profile/info")
}

func (this *ProfileController) FollowUser() {

	// follow users
	followUsers, err := models.FollowModel.GetFollowsByUserIdAndType(this.UserId, models.Follow_Type_User)
	if err != nil {
		this.ErrorLog("获取我的关注用户列表失败: "+err.Error())
		this.ViewError("获取关注用户列表失败", "/system/profile/info")
	}
	userIds := []string{}
	for _, followUser := range followUsers {
		userIds = append(userIds, followUser["object_id"])
	}
	users, err := models.UserModel.GetUsersByUserIds(userIds)
	if err != nil {
		this.ErrorLog("获取我的关注用户列表失败: "+err.Error())
		this.ViewError("获取关注用户列表失败", "/system/profile/info")
	}
	for _, user := range users {
		user["follow_id"] = "0"
		for _, followUser := range followUsers {
			if followUser["object_id"] == user["user_id"] {
				user["follow_id"] = followUser["follow_id"]
				break
			}
		}
	}

	// fans users
	followedUsers, err := models.FollowModel.GetFollowsByObjectIdAndType(this.UserId, models.Follow_Type_User)
	if err != nil {
		this.ErrorLog("获取关注用户列表失败: "+err.Error())
		this.ViewError("获取关注用户列表失败", "/system/profile/info")
	}
	followedUserIds := []string{}
	for _, followedUser := range followedUsers {
		followedUserIds = append(followedUserIds, followedUser["user_id"])
	}
	fansUsers, err := models.UserModel.GetUsersByUserIds(followedUserIds)
	if err != nil {
		this.ErrorLog("获取关注用户列表失败: "+err.Error())
		this.ViewError("获取关注用户列表失败", "/system/profile/info")
	}

	this.Data["users"] = users
	this.Data["fansUsers"] = fansUsers
	this.Data["followCount"] = len(users)
	this.Data["fansCount"] = len(fansUsers)
	this.Data["user"] = this.User
	this.viewLayout("profile/follow_user", "default")
}

func (this *ProfileController) Activity() {
	this.viewLayout("profile/activity", "default")
}

func (this *ProfileController) Password()  {

	this.viewLayout("profile/password", "default")
}

func (this *ProfileController) SavePass()  {

	pwd := strings.TrimSpace(this.GetString("pwd", ""))
	pwdNew := strings.TrimSpace(this.GetString("pwd_new", ""))
	pwdConfirm := strings.TrimSpace(this.GetString("pwd_confirm", ""))

	if (pwd == "") || (pwdNew == "") || (pwdConfirm == "") {
		this.jsonError("密码不能为空！")
	}

	p := models.UserModel.EncodePassword(pwd)
	if p != this.User["password"] {
		this.jsonError("当前密码错误")
	}
	if pwdConfirm != pwdNew {
		this.jsonError("确认密码和新密码不一致")
	}

	_, err := models.UserModel.Update(this.UserId, map[string]interface{}{
		"password": models.UserModel.EncodePassword(pwdNew),
	})

	// 阻止日志记录 password
	this.Ctx.Request.PostForm.Del("pwd")
	this.Ctx.Request.PostForm.Del("pwd_new")
	this.Ctx.Request.PostForm.Del("pwd_confirm")

	if err != nil {
		this.ErrorLog("修改密码失败：" + err.Error())
		this.jsonError("修改密码失败")
	}

	this.InfoLog("修改密码成功")
	this.jsonSuccess("修改密码成功, 下次登录时生效", nil, "/system/profile/password")
}