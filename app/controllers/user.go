package controllers

import (
	"strings"

	"github.com/phachon/mm-wiki/app/models"
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
	number, _ := this.GetRangeInt("number", 20, 10, 100)
	username := strings.TrimSpace(this.GetString("username", ""))
	if username != "" {
		keywords["username"] = username
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
		this.ErrorLog("获取全部用户列表失败: " + err.Error())
		this.ViewError("获取全部用户列表失败", "/main/index")
	}

	followUsers, err := models.FollowModel.GetFollowsByUserIdAndType(this.UserId, models.Follow_Type_User)
	if err != nil {
		this.ErrorLog("获取全部用户列表失败: " + err.Error())
		this.ViewError("获取全部用户列表失败", "/main/index")
	}

	for _, user := range users {
		user["follow"] = "0"
		user["follow_id"] = "0"
		for _, followUser := range followUsers {
			if followUser["object_id"] == user["user_id"] {
				user["follow_id"] = followUser["follow_id"]
				user["follow"] = "1"
				break
			}
		}
	}

	this.Data["users"] = users
	this.Data["count"] = count
	this.Data["login_user_id"] = this.UserId
	this.SetPaginator(number, count)
	this.viewLayout("user/list", "default")
}

func (this *UserController) Follow() {

	followUsers, err := models.FollowModel.GetFollowsByUserIdAndType(this.UserId, models.Follow_Type_User)
	if err != nil {
		this.ErrorLog("获取关注用户列表失败: " + err.Error())
		this.ViewError("获取关注用户列表失败", "/user/list")
	}

	userIds := []string{}
	for _, followUser := range followUsers {
		userIds = append(userIds, followUser["object_id"])
	}
	users, err := models.UserModel.GetUsersByUserIds(userIds)
	if err != nil {
		this.ErrorLog("获取关注用户列表失败: " + err.Error())
		this.ViewError("获取关注用户列表失败", "/user/list")
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

	this.Data["users"] = users
	this.Data["count"] = len(users)
	this.viewLayout("user/follow", "default")
}

func (this *UserController) Info() {

	userId := this.GetString("user_id", "")
	if userId == "" {
		this.ViewError("用户不存在！", "/main/index")
	}

	if this.UserId == userId {
		this.Redirect("/system/main/index", 302)
	}

	user, err := models.UserModel.GetUserByUserId(userId)
	if err != nil {
		this.ErrorLog("查找用户出错：" + err.Error())
		this.ViewError("查找用户出错！", "/main/index")
	}
	if len(user) == 0 {
		this.ViewError("用户不存在！", "/main/index")
	}

	logDocuments, err := models.LogDocumentModel.GetLogDocumentsByUserIdAndLimit(userId, 0, 10)
	if err != nil {
		this.ErrorLog("查找用户活动出错：" + err.Error())
		this.ViewError("查找用户活动出错！", "/main/index")
	}

	docIds := []string{}
	for _, logDocument := range logDocuments {
		docIds = append(docIds, logDocument["document_id"])
	}
	documents, err := models.DocumentModel.GetDocumentsByDocumentIds(docIds)
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

	this.Data["user"] = user
	this.Data["logDocuments"] = logDocuments
	this.Data["count"] = len(logDocuments)
	this.viewLayout("user/info", "default")
}

func (this *UserController) FollowUser() {

	userId := this.GetString("user_id", "")
	if userId == "" {
		this.ViewError("用户不存在！", "/user/list")
	}

	user, err := models.UserModel.GetUserByUserId(userId)
	if err != nil {
		this.ErrorLog("查找用户出错：" + err.Error())
		this.ViewError("查找用户出错！", "/user/list")
	}
	if len(user) == 0 {
		this.ViewError("用户不存在！", "/user/list")
	}

	// follow users
	followUsers, err := models.FollowModel.GetFollowsByUserIdAndType(userId, models.Follow_Type_User)
	if err != nil {
		this.ErrorLog("获取关注用户列表失败: " + err.Error())
		this.ViewError("获取关注用户列表失败", "/user/list")
	}
	userIds := []string{}
	for _, followUser := range followUsers {
		userIds = append(userIds, followUser["object_id"])
	}
	users, err := models.UserModel.GetUsersByUserIds(userIds)
	if err != nil {
		this.ErrorLog("获取关注用户列表失败: " + err.Error())
		this.ViewError("获取关注用户列表失败", "/user/list")
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
	followedUsers, err := models.FollowModel.GetFollowsByObjectIdAndType(userId, models.Follow_Type_User)
	if err != nil {
		this.ErrorLog("获取关注用户列表失败: " + err.Error())
		this.ViewError("获取关注用户列表失败", "/user/list")
	}
	followedUserIds := []string{}
	for _, followedUser := range followedUsers {
		followedUserIds = append(followedUserIds, followedUser["user_id"])
	}
	fansUsers, err := models.UserModel.GetUsersByUserIds(followedUserIds)
	if err != nil {
		this.ErrorLog("获取关注用户列表失败: " + err.Error())
		this.ViewError("获取关注用户列表失败", "/user/list")
	}

	this.Data["users"] = users
	this.Data["fansUsers"] = fansUsers
	this.Data["followCount"] = len(users)
	this.Data["fansCount"] = len(fansUsers)
	this.Data["user"] = user
	this.viewLayout("user/follow_user", "default")
}

func (this *UserController) FollowPage() {

	userId := this.GetString("user_id", "")
	if userId == "" {
		this.ViewError("用户不存在！", "/user/list")
	}

	user, err := models.UserModel.GetUserByUserId(userId)
	if err != nil {
		this.ErrorLog("查找用户出错：" + err.Error())
		this.ViewError("查找用户出错！", "/user/list")
	}
	if len(user) == 0 {
		this.ViewError("用户不存在！", "/user/list")
	}

	followPages, err := models.FollowModel.GetFollowsByUserIdAndType(userId, models.Follow_Type_Doc)
	if err != nil {
		this.ErrorLog("获取用户关注页面列表失败: " + err.Error())
		this.ViewError("获取用户关注页面列表失败", "/user/list")
	}

	documentIds := []string{}
	for _, followPage := range followPages {
		documentIds = append(documentIds, followPage["object_id"])
	}

	pages, err := models.DocumentModel.GetDocumentsByDocumentIds(documentIds)
	if err != nil {
		this.ErrorLog("获取用户关注页面列表失败: " + err.Error())
		this.ViewError("获取用户关注页面列表失败", "/user/list")
	}

	for _, page := range pages {
		page["follow_id"] = "0"
		for _, followPage := range followPages {
			if page["document_id"] == followPage["object_id"] {
				page["follow_id"] = followPage["follow_id"]
				break
			}
		}
	}
	this.Data["pages"] = pages
	this.Data["count"] = len(pages)
	this.Data["user"] = user
	this.viewLayout("user/collect_page", "default")
}
