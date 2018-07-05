package controllers

import (
	"mm-wiki/app/utils"
	"mm-wiki/app/models"
	"strings"
	"fmt"
)

type DocumentController struct {
	BaseController
}

// document index
func (this *DocumentController) Index() {

	documentId := this.GetString("document_id", "")
	if documentId == "" {
		this.ViewError("页面参数错误！", "/space/index")
	}

	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("查找空间文档 "+documentId+" 失败："+err.Error())
		this.ViewError("查找文档失败！")
	}
	if len(document) == 0 {
		this.ViewError("文档不存在！")
	}
	spaceId := document["space_id"]
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("查找文档 "+documentId+" 所在空间失败："+err.Error())
		this.ViewError("查找文档失败！")
	}
	if len(space) == 0 {
		this.ViewError("文档所在空间不存在！")
	}

	// get default space document
	spaceDocument, err := models.DocumentModel.GetSpaceDefaultDocument(spaceId)
	if err != nil {
		this.ErrorLog("查找文档 "+documentId+" 失败："+err.Error())
		this.ViewError("查找文档失败！")
	}
	if len(spaceDocument) == 0 {
		this.ViewError(" 空间首页文档不存在！")
	}

	// get space all document
	documents, err := models.DocumentModel.GetAllSpaceDocuments(spaceId)
	if err != nil {
		this.ErrorLog("查找文档 "+documentId+" 所在空间失败："+err.Error())
		this.ViewError("查找文档失败！")
	}

	this.Data["documents"] = documents
	this.Data["default_document_id"] = documentId
	this.Data["space"] = space
	this.Data["space_document"] = spaceDocument
	this.viewLayout("document/index", "document")
}

// add document
func (this *DocumentController) Add() {

	spaceId := this.GetString("space_id", "0")
	parentId := this.GetString("parent_id", "0")

	if spaceId == "0" {
		this.ViewError("没有选择空间！")
	}
	if parentId == "0" {
		this.ViewError("没有选择上级！")
	}
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("添加文档失败："+err.Error())
		this.ViewError("添加文档失败！")
	}
	if len(space) == 0 {
		this.ViewError("空间不存在！")
	}
	parentDocument, err := models.DocumentModel.GetDocumentByDocumentId(parentId)
	if err != nil {
		this.ErrorLog("添加文档 "+parentId+" 失败："+err.Error())
		this.jsonError("添加文档失败！")
	}
	if len(parentDocument) == 0 {
		this.jsonError("父文档不存在！")
	}
	path := parentDocument["path"] + ","+parentId
	// get parent documents by path
	parentDocuments, err := models.DocumentModel.GetParentDocumentsByPath(path)
	if err != nil {
		this.ErrorLog("查找父文档失败："+err.Error())
		this.ViewError("查找父文档失败！")
	}
	if len(parentDocuments) == 0 {
		this.ViewError("父文档不存在！")
	}

	this.Data["parent_documents"] = parentDocuments
	this.Data["parent_id"] = parentId
	this.Data["space_id"] = spaceId
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
	if parentId == "0" {
		this.jsonError("没有选择父文档！")
	}
	if name == "" {
		this.jsonError("文档名称不能为空！")
	}
	if name == utils.Document_Default_FileName {
		this.jsonError("文档名称不能为 "+ utils.Document_Default_FileName+" ！")
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

	parentDocument, err := models.DocumentModel.GetDocumentByDocumentId(parentId)
	if err != nil {
		this.ErrorLog("创建保存文档失败："+err.Error())
		this.jsonError("创建文档失败！")
	}
	if len(parentDocument) == 0 {
		this.jsonError("父文档不存在！")
	}
	if parentDocument["type"] != fmt.Sprintf("%d", models.Document_Type_Dir) {
		this.jsonError("父文档不是目录！")
	}

	// check document name
	document, err := models.DocumentModel.GetDocumentByNameParentIdAndSpaceId(name, parentId, spaceId, docType)
	if err != nil {
		this.ErrorLog("创建保存文档失败："+err.Error())
		this.jsonError("创建文档失败！")
	}
	if len(document) != 0 {
		this.jsonError("该文档名称已经存在！")
	}

	insertDocument := map[string]interface{}{
		"parent_id" : parentId,
		"space_id" : spaceId,
		"name" : name,
		"type" : docType,
		"path" : parentDocument["path"]+","+parentId,
		"create_user_id" : this.UserId,
		"edit_user_id" : this.UserId,
	}
	documentId, err := models.DocumentModel.Insert(insertDocument)
	if err != nil {
		this.ErrorLog("创建文档失败：" + err.Error())
		this.jsonError("创建文档失败")
	}
	this.InfoLog("创建文档 "+utils.Convert.IntToString(documentId, 10)+" 成功")
	this.jsonSuccess("创建文档成功", nil, "/document/index?document_id="+utils.Convert.IntToString(documentId, 10))
}

// edit document
func (this *DocumentController) Edit() {

	documentId := this.GetString("document_id", "0")
	spaceId := strings.TrimSpace(this.GetString("space_id", "0"))

	if spaceId == "0" {
		this.ViewError("没有选择空间！")
	}
	if documentId == "0" {
		this.ViewError("没有选择文档目录！")
	}

	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("修改文档目录失败："+err.Error())
		this.ViewError("修改文档目录失败！")
	}
	if len(space) == 0 {
		this.ViewError("空间不存在！")
	}

	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("修改文档目录失败："+err.Error())
		this.jsonError("修改文档目录失败！")
	}
	if len(document) == 0 {
		this.jsonError("文档目录不存在！")
	}

	path := document["path"]
	// get parent documents by path
	parentDocuments, err := models.DocumentModel.GetParentDocumentsByPath(path)
	if err != nil {
		this.ErrorLog("查找父文档失败："+err.Error())
		this.ViewError("查找父文档失败！")
	}
	if len(parentDocuments) == 0 {
		this.ViewError("父文档不存在！")
	}

	this.Data["document"] = document
	this.Data["parent_documents"] = parentDocuments
	this.viewLayout("document/edit", "default")
}

// modify document
func (this *DocumentController) Modify() {

	documentId := this.GetString("document_id", "")

	this.Redirect("/document/view?document_id="+documentId, 302)
}