package controllers

import (
	"time"

	"github.com/phachon/mm-wiki/app/models"
	"github.com/phachon/mm-wiki/app/utils"
)

type CollectionController struct {
	BaseController
}

func (this *CollectionController) Add() {

	redirect := this.Ctx.Request.Referer()

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/main/index")
	}
	resourceId := this.GetString("resource_id", "")
	colType, _ := this.GetInt("type", 1)

	if resourceId == "" {
		this.jsonError("没有选项收藏资源！")
	}
	if colType != models.Collection_Type_Doc && colType != models.Collection_Type_Space {
		this.jsonError("收藏类型错误！")
	}

	collect, err := models.CollectionModel.GetCollectionByUserIdTypeAndResourceId(this.UserId, colType, resourceId)
	if err != nil {
		this.ErrorLog("添加收藏失败：" + err.Error())
		this.jsonError("添加收藏失败！")
	}
	if len(collect) > 0 {
		this.jsonError("您已收藏过，不能重复收藏！")
	}
	insertCollection := map[string]interface{}{
		"user_id":     this.UserId,
		"resource_id": resourceId,
		"type":        colType,
		"create_time": time.Now().Unix(),
	}
	collectId, err := models.CollectionModel.Insert(insertCollection)
	if err != nil {
		this.ErrorLog("添加收藏失败：" + err.Error())
		this.jsonError("添加收藏失败！")
	}

	this.InfoLog("添加收藏 " + utils.Convert.IntToString(collectId, 10) + " 成功")
	this.jsonSuccess("收藏成功！", nil, redirect)
}

func (this *CollectionController) Cancel() {

	redirect := this.Ctx.Request.Referer()
	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/space/list")
	}
	collectionId := this.GetString("collection_id", "")

	if collectionId == "" {
		this.jsonError("没有选择收藏资源！")
	}

	collection, err := models.CollectionModel.GetCollectionByCollectionId(collectionId)
	if err != nil {
		this.ErrorLog("取消收藏 " + collectionId + " 失败：" + err.Error())
		this.jsonError("取消收藏失败！")
	}
	if len(collection) == 0 {
		this.jsonError("收藏资源不存在！")
	}
	if collection["user_id"] != this.UserId {
		this.jsonError("您只能取消自己的收藏！")
	}

	err = models.CollectionModel.Delete(collectionId)
	if err != nil {
		this.ErrorLog("取消收藏 " + collectionId + " 失败：" + err.Error())
		this.jsonError("取消收藏失败！")
	}

	this.InfoLog("取消收藏 " + collectionId + " 成功")
	this.jsonSuccess("已取消收藏！", nil, redirect)
}
