package controllers

import (
	"mm-wiki/app/models"
	"time"
)

type CollectionController struct {
	BaseController
}

func (this *CollectionController) Space() {

	redirect := this.Ctx.Request.Referer()

	if !this.IsPost() {
		this.ViewError("请求方式有误！", redirect)
	}
	spaceId := this.GetString("space_id", "")
	if spaceId == "" {
		this.jsonError("没有选择空间！")
	}

	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("收藏空间 "+spaceId+"失败: "+err.Error())
		this.jsonError("收藏空间失败！")
	}
	if len(space) == 0 {
		this.jsonError("空间不存在！")
	}

	insertCollection := map[string]interface{}{
		"user_id": this.UserId,
		"resource_id": spaceId,
		"type": models.Collection_Type_Space,
		"create_time": time.Now().Unix(),
	}
	_, err = models.CollectionModel.Insert(insertCollection)
	if err != nil {
		this.ErrorLog("收藏空间 "+spaceId+" 失败：" + err.Error())
		this.jsonError("收藏空间失败！")
	}

	this.InfoLog("收藏空间 "+spaceId+" 成功")
	this.jsonSuccess("收藏空间成功！", nil, "/space/index")
}

func (this *CollectionController) Page() {

	redirect := this.Ctx.Request.Referer()

	if !this.IsPost() {
		this.ViewError("请求方式有误！", redirect)
	}
	pageId := this.GetString("page_id", "")
	if pageId == "" {
		this.jsonError("没有选择页面！")
	}

	this.jsonSuccess("收藏空间成功！", nil, "/space/index")
}

func (this *CollectionController) Cancel() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/space/list")
	}
	collectionId := this.GetString("collection_id", "")

	if collectionId == "" {
		this.jsonError("没有选择收藏资源！")
	}

	err := models.CollectionModel.Delete(collectionId)
	if err != nil {
		this.ErrorLog("取消收藏 "+collectionId+" 失败：" + err.Error())
		this.jsonError("取消收藏失败！")
	}

	this.InfoLog("取消收藏 "+collectionId+" 成功")
	this.jsonSuccess("已取消收藏！", nil, "/space/index")
}