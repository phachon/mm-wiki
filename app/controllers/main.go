package controllers

import (
	"mm-wiki/app/models"
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

	// link
	links, err := models.LinkModel.GetLinksOrderBySequence()
	if err != nil {
		this.ErrorLog("查找快速链接失败："+err.Error())
		this.ViewError("查找快速链接失败！")
	}
	// contacts
	contacts, err := models.ContactModel.GetAllContact()
	if err != nil {
		this.ErrorLog("查找快速链接失败："+err.Error())
		this.ViewError("查找快速链接失败！")
	}

	// main title config
	mainTitle := ""
	mainDescription := ""
	mainTitleConfig, err := models.ConfigModel.GetConfigByKey(models.Config_Key_MainTitle)
	if err != nil {
		this.ErrorLog("查找 main_title 配置失败："+err.Error())
	}else {
		if len(mainTitleConfig) > 0 {
			mainTitle = mainTitleConfig["value"]
		}
	}
	mainDescriptionConfig, err := models.ConfigModel.GetConfigByKey(models.Config_Key_MainDescription)
	if err != nil {
		this.ErrorLog("查找 main_description 配置失败："+err.Error())
	}else {
		if len(mainDescriptionConfig) > 0 {
			mainDescription = mainDescriptionConfig["value"]
		}
	}

	this.Data["panel_title"] = mainTitle
	this.Data["panel_description"] = mainDescription
	this.Data["logDocuments"] = logDocuments
	this.Data["links"] = links
	this.Data["contacts"] = contacts
	this.SetPaginator(number, count)
	this.viewLayout("main/default", "default")
}

func (this *MainController) About() {
	this.viewLayout("main/about", "default")
}

func (this *MainController) Search() {

	page, _ := this.GetInt("page", 1)
	documentName := this.GetString("document_name", "")

	number := 15
	limit := (page - 1) * number

	var documents = []map[string]string{}
	var err error
	var count int64

	if documentName != "" {
		count, err = models.DocumentModel.CountDocumentsLikeName(documentName)
		if err != nil {
			this.ErrorLog("搜索文档总数出错："+err.Error())
			this.ViewError("搜索文档错误！")
		}
		if count > 0 {
			documents, err = models.DocumentModel.GetDocumentsByLikeNameAndLimit(documentName, limit, number)
			if err != nil {
				this.ErrorLog("搜索文档列表出错："+err.Error())
				this.ViewError("搜索文档错误！")
			}
		}
	}

	this.Data["document_name"] = documentName
	this.Data["documents"] = documents
	this.Data["count"] = count
	this.SetPaginator(number, count)
	this.viewLayout("main/search", "default")
}