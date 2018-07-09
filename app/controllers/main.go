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

	page, _ := this.GetInt("page", 1)
	number := 8
	limit := (page - 1) * number

	logDocuments, err := models.LogDocumentModel.GetLogDocumentsByLimit(limit, number)
	if err != nil {
		this.ErrorLog("查找更新文档列表失败："+err.Error())
		this.ViewError("查找更新文档列表失败！")
	}
	count, err := models.LogDocumentModel.CountLogDocuments()
	if err != nil {
		this.ErrorLog("查找更新文档总数失败："+err.Error())
		this.ViewError("查找更新文档列表失败！")
	}

	userIds := []string{}
	docIds := []string{}
	for _, logDocument := range logDocuments {
		userIds = append(userIds, logDocument["user_id"])
		docIds = append(docIds, logDocument["document_id"])
	}
	users, err := models.UserModel.GetUsersByUserIds(userIds)
	if err != nil {
		this.ErrorLog("查找更新文档用户失败："+err.Error())
		this.ViewError("查找更新文档列表失败！")
	}
	docs, err := models.DocumentModel.GetDocumentsByDocumentIds(docIds)
	if err != nil {
		this.ErrorLog("查找文档信息失败："+err.Error())
		this.ViewError("查找更新文档列表失败！")
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

	title := beego.AppConfig.String("panel::title")
	description := beego.AppConfig.String("panel::description")
	this.Data["panel_title"] = title
	this.Data["panel_description"] = description
	this.Data["logDocuments"] = logDocuments
	this.SetPaginator(number, count)
	this.viewLayout("main/default", "default")
}

func (this *MainController) About() {
	this.viewLayout("main/about", "default")
}