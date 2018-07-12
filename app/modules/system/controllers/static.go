package controllers

import (
	"mm-wiki/app/models"
)

type StaticController struct {
	BaseController
}

func (this *StaticController) User() {

	this.viewLayout("static/user", "static")
}

func (this *StaticController) SpaceDocsRank() {

	number := 15

	spaceDocsRank := []map[string]string{}
	spaceDocCountIds , err := models.DocumentModel.GetSpaceIdsOrderByCountDocumentLimit(number)
	if err != nil {
		this.ErrorLog("查找空间文档排行出错："+err.Error())
		this.jsonError("查找数据出错")
	}
	if len(spaceDocCountIds) > 0 {
		spaceIds := []string{}
		for _, spaceDocCountId := range spaceDocCountIds {
			spaceIds = append(spaceIds, spaceDocCountId["space_id"])
		}
		spaces, err := models.SpaceModel.GetSpaceBySpaceIds(spaceIds)
		if err != nil {
			this.ErrorLog("查找空间文档排行出错："+err.Error())
			this.jsonError("查找数据出错")
		}
		for _, spaceDocCountId := range spaceDocCountIds {
			spaceDocCount := map[string]string{
				"space_name": "",
				"total": spaceDocCountId["total"],
			}
			for _, space := range spaces {
				if spaceDocCountId["space_id"] == space["space_id"] {
					spaceDocCount["space_name"] = space["name"]
					break
				}
			}
			spaceDocsRank = append(spaceDocsRank, spaceDocCount)
		}
	}

	this.jsonSuccess("ok", spaceDocsRank)
}

func (this *StaticController) DocCountByTime() {

	documentCountDate , err := models.DocumentModel.GetCountGroupByCreateTime()
	if err != nil {
		this.ErrorLog("查找空间文档排行出错："+err.Error())
		this.jsonError("查找数据出错")
	}

	this.jsonSuccess("ok", documentCountDate)
}