package controllers

import (
	"mm-wiki/app/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	BaseController
}

func (this *MainController) Index() {

	collectDocs, err := models.CollectionModel.GetCollectionsByUserIdAndType(this.UserId, models.Collection_Type_Doc)
	if err != nil {
		this.ErrorLog("查找收藏文档错误: "+err.Error())
		this.ViewError("查找收藏文档错误！")
	}
	docIds := []string{}
	for _, collectDoc := range collectDocs {
		docIds = append(docIds, collectDoc["resource_id"])
	}

	documents, err := models.DocumentModel.GetDocumentsByDocumentIds(docIds)
	if err != nil {
		this.ErrorLog("查找收藏文档错误: "+err.Error())
		this.ViewError("查找收藏文档错误！")
	}

	this.Data["documents"] = documents
	this.Data["count"] = len(documents)
	this.viewLayout("main/index", "main")
}
func (this *MainController) Default() {

	title := beego.AppConfig.String("panel::title")
	description := beego.AppConfig.String("panel::description")

	this.Data["panel_title"] = title
	this.Data["panel_description"] = description
	this.viewLayout("main/default", "default")
}

func (this *MainController) About() {
	this.viewLayout("main/about", "default")
}