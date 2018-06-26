package controllers

import (
	"strings"
	"mm-wiki/app/models"
	"time"
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
	if username != "" {
		keywords["username"] = username
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
		this.ErrorLog("获取全部用户列表失败: "+err.Error())
		this.ViewError("获取全部用户列表失败", "/main/index")
	}

	followUsers, err := models.FollowModel.GetFollowsByUserId(this.UserId)
	if err != nil {
		this.ErrorLog("获取全部用户列表失败: "+err.Error())
		this.ViewError("获取全部用户列表失败", "/main/index")
	}

	for _, user := range users {
		user["follow"] = "0"
		user["follow_id"] = "0"
		for _, followUser := range followUsers {
			if followUser["follow_user_id"] == user["user_id"] {
				user["follow_id"] = followUser["follow_id"]
				user["follow"] = "1"
				break
			}
		}
	}

	this.Data["users"] = users
	this.Data["count"] = count
	this.SetPaginator(number, count)
	this.viewLayout("user/list", "default")
}

func (this *UserController) Follow() {

	followUsers, err := models.FollowModel.GetFollowsByUserId(this.UserId)
	if err != nil {
		this.ErrorLog("获取关注用户列表失败: "+err.Error())
		this.ViewError("获取关注用户列表失败", "/user/list")
	}

	userIds := []string{}
	for _, followUser := range followUsers {
		userIds = append(userIds, followUser["follow_user_id"])
	}
	users, err := models.UserModel.GetUsersByUserIds(userIds)
	if err != nil {
		this.ErrorLog("获取关注用户列表失败: "+err.Error())
		this.ViewError("获取关注用户列表失败", "/user/list")
	}
	for _, user := range users {
		user["follow_id"] = "0"
		for _, followUser := range followUsers {
			if followUser["follow_user_id"] == user["user_id"] {
				user["follow_id"] = followUser["follow_id"]
				break
			}
		}
	}

	this.Data["users"] = users
	this.Data["count"] = len(users)
	this.viewLayout("user/follow", "default")
}

func (this *UserController) AddFollow() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/user/index")
	}
	userId := this.GetString("user_id", "")
	if userId == "" {
		this.jsonError("没有选择用户！")
	}

	user, err := models.UserModel.GetUserByUserId(userId)
	if err != nil {
		this.ErrorLog("关注用户 "+userId+" 失败: "+err.Error())
		this.jsonError("关注用户失败！")
	}
	if len(user) == 0 {
		this.jsonError("用户不存在！")
	}

	insertFollow := map[string]interface{}{
		"user_id": this.UserId,
		"follow_user_id": userId,
		"create_time": time.Now().Unix(),
	}
	_, err = models.FollowModel.Insert(insertFollow)
	if err != nil {
		this.ErrorLog("关注用户 "+userId+" 失败：" + err.Error())
		this.jsonError("关注用户失败！")
	}

	this.InfoLog("关注用户 "+userId+" 成功")
	this.jsonSuccess("关注用户成功！", nil, "/user/index")
}

func (this *UserController) CancelFollow() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/user/index")
	}
	followId := this.GetString("follow_id", "")
	userId := this.GetString("user_id", "")
	if followId == "" {
		this.jsonError("参数错误！")
	}

	err := models.FollowModel.Delete(followId)
	if err != nil {
		this.ErrorLog("取消关注用户 "+userId+" 失败：" + err.Error())
		this.jsonError("取消关注用户失败！")
	}

	this.InfoLog("取消关注用户 "+userId+" 成功")
	this.jsonSuccess("已取消关注！", nil, "/user/index")
}

func (this *UserController) Info() {

	userId := this.GetString("user_id", "")
	if userId == "" {
		this.ViewError("用户不存在！", "/system/user/list")
	}

	user, err := models.UserModel.GetUserByUserId(userId)
	if err != nil {
		this.ErrorLog("查找用户出错："+err.Error())
		this.ViewError("查找用户出错！", "/system/user/list")
	}
	if len(user) == 0 {
		this.ViewError("用户不存在！", "/system/user/list")
	}

	this.Data["user"] = user
	this.viewLayout("user/info", "default")
}