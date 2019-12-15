package controllers

import (
	"github.com/phachon/mm-wiki/app/models"
	"github.com/phachon/mm-wiki/app/utils"
)

type FollowController struct {
	BaseController
}

func (this *FollowController) Add() {

	redirect := this.Ctx.Request.Referer()
	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/main/index")
	}
	objectId := this.GetString("object_id", "")
	followType, _ := this.GetInt("type", 1)
	if objectId == "" {
		this.jsonError("没有选择关注对象！")
	}
	if followType != models.Follow_Type_Doc && followType != models.Follow_Type_User {
		this.jsonError("关注类型错误！")
	}
	if followType == models.Follow_Type_User && objectId == this.UserId {
		this.jsonError("不能关注自己！")
	}

	follow, err := models.FollowModel.GetFollowByUserIdAndTypeAndObjectId(this.UserId, followType, objectId)
	if err != nil {
		this.ErrorLog("添加关注失败：" + err.Error())
		this.jsonError("添加关注失败！")
	}
	if len(follow) > 0 {
		this.jsonError("您已关注过，不能重复关注！")
	}
	fId, err := models.FollowModel.Insert(this.UserId, followType, objectId)
	if err != nil {
		this.ErrorLog("添加关注失败：" + err.Error())
		this.jsonError("添加关注失败！")
	}

	this.InfoLog("添加关注" + utils.Convert.IntToString(fId, 10) + " 成功")
	this.jsonSuccess("关注成功！", nil, redirect)
}

func (this *FollowController) Cancel() {

	redirect := this.Ctx.Request.Referer()
	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/main/index")
	}
	followId := this.GetString("follow_id", "")
	if followId == "" {
		this.jsonError("没有选择关注对象！")
	}

	follow, err := models.FollowModel.GetFollowByFollowId(followId)
	if err != nil {
		this.ErrorLog("取消关注失败：" + err.Error())
		this.jsonError("取消关注失败！")
	}
	if len(follow) == 0 {
		this.jsonError("关注对象不存在！")
	}
	if follow["user_id"] != this.UserId {
		this.jsonError("您只能取消自己的关注！")
	}

	err = models.FollowModel.Delete(followId)
	if err != nil {
		this.ErrorLog("取消关注 " + followId + " 失败：" + err.Error())
		this.jsonError("取消关注用户失败！")
	}

	this.InfoLog("取消关注 " + followId + " 成功")
	this.jsonSuccess("已取消关注！", nil, redirect)
}
