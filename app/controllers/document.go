package controllers

import (
	"fmt"
	"mm-wiki/app/utils"
	"mm-wiki/app/models"
)

type DocumentController struct {
	BaseController
}

// document info
func (this *DocumentController) Index() {

	documentId := this.GetString("document_id", "")
	if documentId == "" {
		this.ViewError("页面参数错误！", "/space/index")
	}

	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("查找文档 "+documentId+" 失败："+err.Error())
		this.ViewError("查找文档失败！")
	}
	if len(document) == 0 {
		this.ErrorLog("查找文档 "+documentId+" 失败："+err.Error())
		this.ViewError("文档不存在！")
	}
	spaceId := document["space_id"]
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("查找文档 "+documentId+" 所在空间失败："+err.Error())
		this.ViewError("查找文档所在空间失败！")
	}
	if len(space) == 0 {
		this.ViewError("文档所在空间不存在！")
	}
	// get space all document
	documents, err := models.DocumentModel.GetDocumentsBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("查找文档失败："+err.Error())
		this.ViewError("查找文档失败！")
	}

	this.Data["documents"] = documents
	this.Data["default_document_id"] = documentId
	this.Data["space"] = space
	this.viewLayout("document/index", "document")
}

// document info
func (this *DocumentController) View() {

	documentId := this.GetString("document_id", "")
	if documentId == "" {
		this.ViewError("文档未找到！")
	}
	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("查找文档 "+documentId+" 失败："+err.Error())
		this.ViewError("查找文档失败！")
	}
	if len(document) == 0 {
		this.ErrorLog("查找文档 "+documentId+" 失败："+err.Error())
		this.ViewError("文档不存在！")
	}
	if document["type"] != fmt.Sprintf("%d", models.Document_Type_Page) {
		this.ViewError("该文档类型不是页面！")
	}

	path := document["path"]
	documentContent, err := models.DocumentModel.GetContentByPath(path)
	if err != nil {
		this.ErrorLog("查找文档 "+documentId+" 失败："+err.Error())
		this.ViewError("文档不存在！")
	}
	this.Data["document_content"] = documentContent
	this.Data["document_id"] = documentId
	this.viewLayout("document/view", "default_document")
}

// document edit
func (this *DocumentController) Edit() {

	documentId := this.GetString("document_id", "")

	fileInfo, err := utils.File.GetFileContents("test.md")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	this.Data["document_content"] = fileInfo
	this.Data["document_id"] = documentId
	this.viewLayout("document/edit", "default_document")
}

// document edit
func (this *DocumentController) Save() {

	documentId := this.GetString("document_id", "")

	this.Redirect("/document/view?document_id="+documentId, 302)
}