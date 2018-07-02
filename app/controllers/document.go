package controllers

import (
	"fmt"
	"mm-wiki/app/utils"
	"mm-wiki/app/models"
	"strings"
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

// add document
func (this *DocumentController) Add() {

	spaceId := this.GetString("space_id", "0")
	parentId := this.GetString("parent_id", "0")

	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("添加文档失败："+err.Error())
		this.ViewError("添加文档失败！")
	}
	if len(space) == 0 {
		this.ViewError("空间不存在！")
	}

	parentDocument := map[string]string{}
	if parentId != "0" {
		// get space by spaceId and parentId
		parentDocument, err = models.DocumentModel.GetDocumentBySpaceIdAndParentId(spaceId, parentId)
		if err != nil {
			this.ErrorLog("查找父文档失败："+err.Error())
			this.ViewError("查找父文档失败！")
		}
		if len(parentDocument) == 0 {
			this.ViewError("父文档不存在！")
		}
	}else {
		parentDocument = map[string]string{
			"document_id": "0",
			"title": space["name"],
			"path": space["name"],
		}
	}

	this.Data["parentDocument"] = parentDocument
	this.Data["spaceId"] = spaceId
	this.viewLayout("document/form", "default")

}

// save document
func (this *DocumentController) Save() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/main/index")
	}
	spaceId := strings.TrimSpace(this.GetString("space_id", "0"))
	parentId := strings.TrimSpace(this.GetString("parent_id", "0"))
	docType, _ := this.GetInt("type", models.Document_Type_Page)
	name := strings.TrimSpace(this.GetString("name", ""))

	if spaceId == "0" {
		this.jsonError("没有选择空间！")
	}
	if name == "" {
		this.jsonError("文档名称不能为空！")
	}
	if docType != models.Document_Type_Page &&
		docType != models.Document_Type_Dir {
		this.jsonError("文档类型错误！")
	}

	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("创建保存文档失败："+err.Error())
		this.jsonError("创建文档失败！")
	}
	if len(space) == 0 {
		this.jsonError("空间不存在！")
	}

	parentPath := ""
	if parentId != "0" {
		parentDocument, err := models.DocumentModel.GetDocumentByDocumentId(parentId)
		if err != nil {
			this.ErrorLog("创建保存文档失败："+err.Error())
			this.jsonError("创建文档失败！")
		}
		if len(parentDocument) != 0 {
			this.jsonError("父文档不存在！")
		}
		parentPath = parentDocument["path"]
	}else {
		parentPath = space["name"]
	}

	document, err := models.DocumentModel.GetDocumentByNameParentIdAndSpaceId(name, parentId, spaceId)
	if err != nil {
		this.ErrorLog("创建保存文档失败："+err.Error())
		this.jsonError("创建文档失败！")
	}
	if len(document) != 0 {
		this.jsonError("目录下文档名称已经存在！")
	}

	path := parentPath+"/"+name
	if docType == models.Document_Type_Page {
		path = path+".md"
	}
	insertDocument := map[string]interface{}{
		"parent_id" : parentId,
		"space_id" : spaceId,
		"name" : name,
		"type" : docType,
		"path" : path,
		"create_user_id" : this.UserId,
		"edit_user_id" : this.UserId,
	}
	documentId, err := models.DocumentModel.Insert(insertDocument)
	if err != nil {
		this.ErrorLog("创建文档失败：" + err.Error())
		this.jsonError("创建文档失败")
	}
	this.InfoLog("创建文档 "+utils.Convert.IntToString(documentId, 10)+" 成功")
	this.jsonSuccess("创建文档成功", nil, "/space/page?space_id="+spaceId)
}

// edit document
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

// modify document
func (this *DocumentController) Modify() {

	documentId := this.GetString("document_id", "")

	this.Redirect("/document/view?document_id="+documentId, 302)
}