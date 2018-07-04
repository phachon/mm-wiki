package controllers

import (
	"mm-wiki/app/models"
	"strings"
	"mm-wiki/app/utils"
)

type PageController struct {
	BaseController
}

// document page view
func (this *PageController) View() {

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

	// get parent documents
	var parentDocuments = []map[string]string{}
	if document["parent_id"] != "0" {
		// get parent documents by parentId
		parentDocuments, err = models.DocumentModel.GetParentDocumentsByParentId(document["parent_id"])
		if err != nil {
			this.ErrorLog("查找父文档失败："+err.Error())
			this.ViewError("查找父文档失败！")
		}
	}else {
		parentDocuments = append(parentDocuments, document)
	}
	if len(parentDocuments) == 0 {
		this.ViewError("父文档不存在！")
	}

	// get document content
	path := document["path"]
	documentContent, err := models.DocumentModel.GetContentByPath(path)
	if err != nil {
		this.ErrorLog("查找文档 "+documentId+" 失败："+err.Error())
		this.ViewError("文档不存在！")
	}

	// get edit user and create user
	users, err := models.UserModel.GetUsersByUserIds([]string{document["create_user_id"], document["edit_user_id"]})
	if err != nil {
		this.ErrorLog("查找文档 "+documentId+" 失败："+err.Error())
		this.ViewError("查找文档失败！")
	}
	if len(users) == 0 {
		this.ViewError("文档创建用户不存在！")
	}

	var createUser = map[string]string{}
	var editUser = map[string]string{}
	for _, user := range users {
		if user["user_id"] == document["create_user_id"] {
			createUser = user
		}
		if user["user_id"] == document["edit_user_id"] {
			editUser = user
		}
	}

	this.Data["create_user"] = createUser
	this.Data["edit_user"] = editUser
	this.Data["document"] = document
	this.Data["page_content"] = documentContent
	this.Data["parent_documents"] = parentDocuments
	this.viewLayout("page/view", "document_page")
}

// page edit
func (this *PageController) Edit() {

	documentId := this.GetString("document_id", "")
	if documentId == "" {
		this.ViewError("文档未找到！")
	}

	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("修改文档 "+documentId+" 失败："+err.Error())
		this.ViewError("修改文档失败！")
	}
	if len(document) == 0 {
		this.ErrorLog("修改文档 "+documentId+" 失败："+err.Error())
		this.ViewError("文档不存在！")
	}


	// get document content
	path := document["path"]
	documentContent, err := models.DocumentModel.GetContentByPath(path)
	if err != nil {
		this.ErrorLog("查找文档 "+documentId+" 失败："+err.Error())
		this.ViewError("文档不存在！")
	}


	this.Data["page_content"] = documentContent
	this.Data["document"] = document
	this.viewLayout("page/edit", "document_page")
}

// page edit
func (this *PageController) Modify() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/space/index")
	}
	documentId := this.GetString("document_id", "")
	newName := strings.TrimSpace(this.GetString("name", ""))
	documentContent := this.GetString("document_page_editor-markdown-doc", "")

	if documentId == "" {
		this.jsonError("您没有选择文档！")
	}
	if newName == "" {
		this.jsonError("文档名称不能为空！")
	}
	if newName == models.Document_Default_FileName {
		this.jsonError("文档名称不能为 "+ models.Document_Default_FileName+" ！")
	}

	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("保存文档 "+documentId+" 失败："+err.Error())
		this.ViewError("保存文档失败！")
	}
	if len(document) == 0 {
		this.ErrorLog("保存文档 "+documentId+" 失败："+err.Error())
		this.ViewError("文档不存在！")
	}
	name := document["name"]
	if newName == name {
		//todo name update
	}else {
		// check document name
		document, err := models.DocumentModel.GetDocumentByNameParentIdAndSpaceId(newName,
			document["parent_id"], document["space_id"], utils.Convert.StringToInt(document["type"]))
		if err != nil {
			this.ErrorLog("保存保存文档失败："+err.Error())
			this.jsonError("保存文档失败！")
		}
		if len(document) != 0 {
			this.jsonError("该文档名称已经存在！")
		}
	}

	this.Redirect("/page/view?page_id="+documentId, 302)
}