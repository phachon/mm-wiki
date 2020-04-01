package controllers

import (
	"github.com/phachon/mm-wiki/app/models"
	"strings"
	"time"
)

type Space_UserController struct {
	BaseController
}

func (this *Space_UserController) Save() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/space/list")
	}
	spaceId := strings.TrimSpace(this.GetString("space_id", ""))
	userId := this.GetString("user_id", "")
	privilege := strings.TrimSpace(this.GetString("privilege", "0"))

	if spaceId == "" {
		this.jsonError("空间不存在！")
	}
	if userId == "" {
		this.jsonError("没有选择用户！")
	}
	if privilege == "" {
		this.jsonError("没有选择用户空间权限！")
	}
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("添加空间 " + spaceId + " 成员 " + userId + " 失败: " + err.Error())
		this.jsonError("添加空间成员失败！")
	}
	if len(space) == 0 {
		this.jsonError("空间不存在！")
	}

	spaceUser, err := models.SpaceUserModel.GetSpaceUserBySpaceIdAndUserId(spaceId, userId)
	if err != nil {
		this.ErrorLog("添加空间 " + spaceId + " 成员 " + userId + " 失败: " + err.Error())
		this.jsonError("添加空间成员失败！")
	}
	if len(spaceUser) > 0 {
		this.jsonError("该用户已经是空间成员！")
	}

	insertValue := map[string]interface{}{
		"user_id":   userId,
		"space_id":  spaceId,
		"privilege": privilege,
	}
	_, err = models.SpaceUserModel.Insert(insertValue)
	if err != nil {
		this.ErrorLog("空间 " + spaceId + " 添加成员 " + userId + " 失败: " + err.Error())
		this.jsonError("添加成员失败！")
	}

	this.InfoLog("空间 " + spaceId + " 添加成员 " + userId + " 成功")
	this.jsonSuccess("添加成员成功！", nil, "/system/space/member?space_id="+spaceId)
}

func (this *Space_UserController) Remove() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/space/list")
	}
	spaceId := this.GetString("space_id", "")
	userId := this.GetString("user_id", "")
	spaceUserId := this.GetString("space_user_id", "")

	if spaceUserId == "" {
		this.jsonError("空间成员不存在！")
	}
	if spaceId == "" {
		this.jsonError("空间不存在！")
	}
	if userId == "" {
		this.jsonError("用户不存在！")
	}

	err := models.SpaceUserModel.Delete(spaceUserId)
	if err != nil {
		this.ErrorLog("移除空间 " + spaceId + " 下成员 " + userId + " 失败：" + err.Error())
		this.jsonError("移除成员失败！")
	}

	this.InfoLog("移除空间 " + spaceId + " 下成员 " + userId + " 成功")
	this.jsonSuccess("移除成员成功！", nil, "/system/space/member?space_id="+spaceId)
}

func (this *Space_UserController) Modify() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/space/list")
	}
	spaceUserId := this.GetString("space_user_id", "")
	privilege := this.GetString("privilege", "0")
	userId := this.GetString("user_id", "")
	spaceId := this.GetString("space_id", "")

	if spaceUserId == "" {
		this.jsonError("更新权限错误！")
	}
	if privilege == "" {
		this.jsonError("没有选择权限！")
	}

	_, err := models.SpaceUserModel.Update(spaceUserId, map[string]interface{}{
		"privilege":   privilege,
		"update_time": time.Now().Unix(),
	})
	if err != nil {
		this.ErrorLog("更新空间 " + spaceId + " 下成员 " + userId + " 权限失败：" + err.Error())
		this.jsonError("更新权限失败！")
	}

	this.InfoLog("更新空间 " + spaceId + " 下成员 " + userId + " 权限成功")
	this.jsonSuccess("更新权限成功！", nil)
}
