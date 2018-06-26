package controllers

import (
	"mm-wiki/app/models"
	"time"
)

type FollowController struct {
	BaseController
}

func (this *FollowController) User() {

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

func (this *FollowController) Cancel() {

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