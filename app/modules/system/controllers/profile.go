package controllers

import (
	"github.com/astaxie/beego/validation"
	"github.com/phachon/mm-wiki/app/models"
	"strings"
)

type ProfileController struct {
	BaseController
}

func (this *ProfileController) Info() {

	user, err := models.UserModel.GetUserByUserId(this.UserId)
	if err != nil {
		this.ErrorLog("获取我的资料失败: " + err.Error())
		this.ViewError("获取资料失败")
	}

	logDocuments, err := models.LogDocumentModel.GetLogDocumentsByUserIdAndLimit(this.UserId, 0, 10)
	if err != nil {
		this.ErrorLog("查找用户活动出错：" + err.Error())
		this.ViewError("查找用户活动出错！", "/main/index")
	}

	docIds := []string{}
	for _, logDocument := range logDocuments {
		docIds = append(docIds, logDocument["document_id"])
	}
	documents, err := models.DocumentModel.GetAllDocumentsByDocumentIds(docIds)
	if err != nil {
		this.ErrorLog("查找用户活动出错: " + err.Error())
		this.ViewError("查找用户活动出错", "/main/index")
	}

	for _, logDocument := range logDocuments {
		for _, document := range documents {
			if document["document_id"] == logDocument["document_id"] {
				logDocument["document_id"] = document["document_id"]
				logDocument["document_name"] = document["name"]
				logDocument["document_type"] = document["type"]
				logDocument["update_time"] = document["update_time"]
			}
		}
	}

	this.Data["logDocuments"] = logDocuments
	this.Data["count"] = len(logDocuments)
	this.Data["user"] = user
	this.viewLayout("profile/info", "profile")
}

func (this *ProfileController) Edit() {

	user, err := models.UserModel.GetUserByUserId(this.UserId)
	if err != nil {
		this.ErrorLog("获取我的资料失败: " + err.Error())
		this.ViewError("获取资料失败")
	}
	this.Data["user"] = user
	this.viewLayout("profile/edit", "profile")
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
	//if phone != "" && !v.Phone(phone, "phone").Ok {
	//	this.jsonError("电话格式不正确！")
	//}

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
		this.ErrorLog("获取我的关注用户列表失败: " + err.Error())
		this.ViewError("获取关注用户列表失败", "/system/profile/info")
	}
	userIds := []string{}
	for _, followUser := range followUsers {
		userIds = append(userIds, followUser["object_id"])
	}
	users, err := models.UserModel.GetUsersByUserIds(userIds)
	if err != nil {
		this.ErrorLog("获取我的关注用户列表失败: " + err.Error())
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
		this.ErrorLog("获取关注用户列表失败: " + err.Error())
		this.ViewError("获取关注用户列表失败", "/system/profile/info")
	}
	followedUserIds := []string{}
	for _, followedUser := range followedUsers {
		followedUserIds = append(followedUserIds, followedUser["user_id"])
	}
	fansUsers, err := models.UserModel.GetUsersByUserIds(followedUserIds)
	if err != nil {
		this.ErrorLog("获取关注用户列表失败: " + err.Error())
		this.ViewError("获取关注用户列表失败", "/system/profile/info")
	}

	this.Data["users"] = users
	this.Data["fansUsers"] = fansUsers
	this.Data["followCount"] = len(users)
	this.Data["fansCount"] = len(fansUsers)
	this.Data["user"] = this.User
	this.viewLayout("profile/follow_user", "profile")
}

func (this *ProfileController) FollowDoc() {

	page, _ := this.GetInt("page", 1)
	number, _ := this.GetRangeInt("number", 10, 10, 100)
	limit := (page - 1) * number

	// follow docs limit
	followDocuments, err := models.FollowModel.GetFollowsByUserIdTypeAndLimit(this.UserId, models.Follow_Type_Doc, limit, number)
	if err != nil {
		this.ErrorLog("获取关注文档列表失败: " + err.Error())
		this.ViewError("获取关注文档列表失败", "/system/profile/info")
	}
	count, err := models.FollowModel.CountFollowsByUserIdAndType(this.UserId, models.Follow_Type_Doc)
	if err != nil {
		this.ErrorLog("获取关注文档列表失败: " + err.Error())
		this.ViewError("获取关注文档列表失败", "/system/profile/info")
	}

	docIds := []string{}
	for _, followDocument := range followDocuments {
		docIds = append(docIds, followDocument["object_id"])
	}
	documents, err := models.DocumentModel.GetDocumentsByDocumentIds(docIds)
	if err != nil {
		this.ErrorLog("获取我的关注用户列表失败: " + err.Error())
		this.ViewError("获取关注用户列表失败", "/system/profile/info")
	}

	for _, followDocument := range followDocuments {
		for _, document := range documents {
			if document["document_id"] == followDocument["object_id"] {
				followDocument["document_id"] = document["document_id"]
				followDocument["document_name"] = document["name"]
				followDocument["update_time"] = document["update_time"]
			}
		}
	}

	autoFollowDoc := models.ConfigModel.GetConfigValueByKey(models.ConfigKeyAutoFollowdoc, "0")

	this.Data["followDocuments"] = followDocuments
	this.Data["count"] = len(documents)
	this.Data["user"] = this.User
	this.Data["autoFollowDoc"] = autoFollowDoc
	this.SetPaginator(number, count)
	this.viewLayout("profile/follow_doc", "profile")
}

func (this *ProfileController) Activity() {

	page, _ := this.GetInt("page", 1)
	number, _ := this.GetRangeInt("number", 15, 10, 100)
	limit := (page - 1) * number
	keyword := strings.TrimSpace(this.GetString("keyword", ""))

	var logDocuments = []map[string]string{}
	var err error
	var count int64
	if keyword != "" {
		logDocuments, err = models.LogDocumentModel.GetLogDocumentsByUserIdKeywordAndLimit(this.UserId, keyword, limit, number)
		count, err = models.LogDocumentModel.CountLogDocumentsByUserIdAndKeyword(this.UserId, keyword)
	} else {
		logDocuments, err = models.LogDocumentModel.GetLogDocumentsByUserIdAndLimit(this.UserId, limit, number)
		count, err = models.LogDocumentModel.CountLogDocumentsByUserId(this.UserId)
	}
	if err != nil {
		this.ErrorLog("我的活动查找失败：" + err.Error())
		this.ViewError("我的活动查找失败！", "/system/main/index")
	}

	userIds := []string{}
	docIds := []string{}
	for _, logDocument := range logDocuments {
		userIds = append(userIds, logDocument["user_id"])
		docIds = append(docIds, logDocument["document_id"])
	}
	users, err := models.UserModel.GetUsersByUserIds(userIds)
	if err != nil {
		this.ErrorLog("我的活动查找失败：" + err.Error())
		this.ViewError("我的活动查找失败！", "/system/main/index")
	}
	docs, err := models.DocumentModel.GetAllDocumentsByDocumentIds(docIds)
	if err != nil {
		this.ErrorLog("我的活动查找失败：" + err.Error())
		this.ViewError("我的活动查找失败！", "/system/main/index")
	}
	for _, logDocument := range logDocuments {
		logDocument["username"] = ""
		for _, user := range users {
			if logDocument["user_id"] == user["user_id"] {
				logDocument["username"] = user["username"]
				logDocument["given_name"] = user["given_name"]
				break
			}
		}
		for _, doc := range docs {
			if logDocument["document_id"] == doc["document_id"] {
				logDocument["document_name"] = doc["name"]
				logDocument["document_type"] = doc["type"]
				break
			}
		}
	}

	this.Data["logDocuments"] = logDocuments
	this.Data["keyword"] = keyword
	this.SetPaginator(number, count)
	this.viewLayout("profile/activity", "profile")
}

func (this *ProfileController) Password() {

	this.viewLayout("profile/password", "profile")
}

func (this *ProfileController) SavePass() {

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
