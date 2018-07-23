package controllers

import (
	"mm-wiki/app/models"
	"strings"
	"mm-wiki/app/utils"
	"regexp"
	"github.com/astaxie/beego/context"
	"github.com/pkg/errors"
	"github.com/astaxie/beego"
	"fmt"
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
	// check space visit_level
	isVisit, isEditor, _ := this.GetDocumentPrivilege(space)
	if !isVisit {
		this.ViewError("您没有权限访问该空间！")
	}

	// get parent documents by document
	parentDocuments, pageFile, err := models.DocumentModel.GetParentDocumentsByDocument(document)
	if err != nil {
		this.ErrorLog("查找父文档失败："+err.Error())
		this.ViewError("查找父文档失败！")
	}
	if len(parentDocuments) == 0 {
		this.ViewError("父文档不存在！")
	}

	// get document content
	documentContent, err := utils.Document.GetContentByPageFile(pageFile)
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

	collectionId := "0"
	collection, err := models.CollectionModel.GetCollectionByUserIdTypeAndResourceId(this.UserId, models.Collection_Type_Doc, documentId)
	if err != nil {
		this.ErrorLog("查找文档 "+documentId+" 失败："+err.Error())
		this.ViewError("文档查找失败！")
	}
	if len(collection) > 0 {
		collectionId = collection["collection_id"]
	}

	this.Data["is_editor"] = isEditor
	this.Data["space"] = space
	this.Data["create_user"] = createUser
	this.Data["edit_user"] = editUser
	this.Data["document"] = document
	this.Data["collection_id"] = collectionId
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
		this.ViewError("文档不存在！")
	}

	spaceId := document["space_id"]
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("修改文档 "+documentId+" 失败："+err.Error())
		this.ViewError("修改文档失败！")
	}
	if len(space) == 0 {
		this.ViewError("文档所在空间不存在！")
	}
	// check space visit_level
	_, isEditor, _ := this.GetDocumentPrivilege(space)
	if !isEditor {
		this.ViewError("您没有权限修改该空间下文档！")
	}

	// get parent documents by document
	_, pageFile, err := models.DocumentModel.GetParentDocumentsByDocument(document)
	if err != nil {
		this.ErrorLog("查找父文档失败："+err.Error())
		this.ViewError("查找父文档失败！")
	}

	// get document content
	documentContent, err := utils.Document.GetContentByPageFile(pageFile)
	if err != nil {
		this.ErrorLog("查找文档 "+documentId+" 失败："+err.Error())
		this.ViewError("文档不存在！")
	}

	autoFollowDoc := "0"
	autoFollowConfig, _ := models.ConfigModel.GetConfigByKey(models.Config_Key_AutoFollowDoc)
	if len(autoFollowConfig) > 0 && autoFollowConfig["value"] == "1" {
		autoFollowDoc = "1"
	}

	sendEmail := "0"
	sendEmailConfig, _ := models.ConfigModel.GetConfigByKey(models.Config_Key_SendEmail)
	if len(sendEmailConfig) > 0 && sendEmailConfig["value"] == "1" {
		sendEmail = "1"
	}

	this.Data["sendEmail"] = sendEmail
	this.Data["autoFollowDoc"] = autoFollowDoc
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
	comment := strings.TrimSpace(this.GetString("comment", ""))
	isNoticeUser := strings.TrimSpace(this.GetString("is_notice_user", "0"))
	isFollowDoc := strings.TrimSpace(this.GetString("is_follow_doc", "0"))

	// rm document_page_editor-markdown-doc
	this.Ctx.Request.PostForm.Del("document_page_editor-markdown-doc")

	if documentId == "" {
		this.jsonError("您没有选择文档！")
	}
	if newName == "" {
		this.jsonError("文档名称不能为空！")
	}
	match, err := regexp.MatchString(`[\\\\/:*?\"<>、|]`, newName)
	if err != nil {
		this.jsonError("文档名称格式不正确！")
	}
	if match {
		this.jsonError("文档名称格式不正确！")
	}
	if newName == utils.Document_Default_FileName {
		this.jsonError("文档名称不能为 "+ utils.Document_Default_FileName+" ！")
	}
	if comment == "" {
		this.jsonError("必须输入此次修改的备注！")
	}

	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("修改文档 "+documentId+" 失败："+err.Error())
		this.jsonError("保存文档失败！")
	}
	if len(document) == 0 {
		this.jsonError("文档不存在！")
	}

	spaceId := document["space_id"]
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("修改文档 "+documentId+" 失败："+err.Error())
		this.jsonError("保存文档失败！")
	}
	if len(space) == 0 {
		this.jsonError("文档所在空间不存在！")
	}
	// check space document privilege
	_, isEditor, _ := this.GetDocumentPrivilege(space)
	if !isEditor {
		this.jsonError("您没有权限修改该空间下文档！")
	}

	// not allow update space document home page name
	if document["parent_id"] == "0" {
		newName = document["name"]
	}
	// check document name
	if newName != document["name"] {
		newDocument, err := models.DocumentModel.GetDocumentByNameParentIdAndSpaceId(newName,
			document["parent_id"], document["space_id"], utils.Convert.StringToInt(document["type"]))
		if err != nil {
			this.ErrorLog("修改文档失败："+err.Error())
			this.jsonError("保存文档失败！")
		}
		if len(newDocument) != 0 {
			this.jsonError("该文档名称已经存在！")
		}
	}

	// update document and file content
	updateValue := map[string]interface{}{
		"name": newName,
		"edit_user_id": this.UserId,
	}
	_, err = models.DocumentModel.UpdateDBAndFile(documentId, document, documentContent, updateValue, comment)
	if err != nil {
		this.ErrorLog("修改文档 "+documentId+" 失败："+err.Error())
		this.jsonError("修改文档失败！")
	}

	// send email to follow user
	if isNoticeUser == "1" {
		go func(comment string, userId string, username string, ctx *context.Context) {
			url := fmt.Sprintf("%s:%d/document/index?document_id=%s", ctx.Input.Site(), ctx.Input.Port(), documentId)
			err := sendEmail(documentId, username, comment, url)
			if err != nil {
				models.LogModel.RecordLogByCtx(err.Error(), models.Log_Level_Error, userId, username, ctx)
			}
		}(comment, this.UserId, this.User["username"], this.Ctx)
	}
	// follow doc
	if isFollowDoc == "1" {
		go func() {
			models.FollowModel.FollowDocument(this.UserId, documentId)
		}()
	}

	this.InfoLog("修改文档 "+documentId+" 成功")
	this.jsonSuccess("文档修改成功！", nil, "/document/index?document_id="+documentId)
}

// document share display
func (this *PageController) Display() {

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
		this.ViewError("文档不存在！")
	}

	// get parent documents by document
	parentDocuments, pageFile, err := models.DocumentModel.GetParentDocumentsByDocument(document)
	if err != nil {
		this.ErrorLog("查找父文档失败："+err.Error())
		this.ViewError("查找父文档失败！")
	}
	if len(parentDocuments) == 0 {
		this.ViewError("父文档不存在！")
	}

	// get document content
	documentContent, err := utils.Document.GetContentByPageFile(pageFile)
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
	this.viewLayout("page/display", "document_share")
}

func (this *PageController) Export() {

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
	// check space visit_level
	if space["visit_level"] == models.Space_VisitLevel_Private {
		ok, _  := models.SpaceUserModel.HasSpaceUser(spaceId, this.UserId)
		if !ok {
			this.ViewError("您没有权限访问该空间！")
		}
	}

	// get parent documents by document
	parentDocuments, pageFile, err := models.DocumentModel.GetParentDocumentsByDocument(document)
	if err != nil {
		this.ErrorLog("查找父文档失败："+err.Error())
		this.ViewError("查找父文档失败！")
	}
	if len(parentDocuments) == 0 {
		this.ViewError("父文档不存在！")
	}

	// get document file
	absPageFile := utils.Document.GetAbsPageFileByPageFile(pageFile)
	this.Ctx.Output.Download(absPageFile, document["name"]+utils.Document_Page_Suffix)
}

func sendEmail(documentId string, username string, comment string, url string) error {

	// get document by documentId
	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		return errors.New("发送邮件通知查找文档失败："+err.Error())
	}

	// get send email open config
	sendEmailConfig, err := models.ConfigModel.GetConfigByKey(models.Config_Key_SendEmail)
	if err != nil {
		return errors.New("发送邮件通知查找发送邮件配置失败："+err.Error())
	}
	if len(sendEmailConfig) == 0 {
		return nil
	}
	if sendEmailConfig["value"] == "0" {
		return nil
	}

	// get email config
	emailConfig, err := models.EmailModel.GetUsedEmail()
	if err != nil {
		return errors.New("发送邮件通知查找邮件服务器配置失败："+err.Error())
	}
	if len(emailConfig) == 0 {
		return nil
	}

	// get follow doc user
	follows, err := models.FollowModel.GetFollowsByObjectIdAndType(documentId, models.Follow_Type_Doc)
	if err != nil {
		return errors.New("发送邮件查找关注文档用户失败："+err.Error())
	}
	if len(follows) == 0 {
		return nil
	}
	userIds := []string{}
	for _, follow := range follows {
		userIds = append(userIds, follow["user_id"])
	}
	users, err := models.UserModel.GetUsersByUserIds(userIds)
	if err != nil {
		return errors.New("发送邮件查找关注文档用户失败："+err.Error())
	}
	if len(users) == 0 {
		return nil
	}
	emails := []string{}
	for _, user := range users {
		if user["email"] != "" {
			emails = append(emails, user["email"])
		}
	}

	// get parent documents by document
	parentDocuments, pageFile, err := models.DocumentModel.GetParentDocumentsByDocument(document)
	if err != nil {
		return errors.New("查找文档内容失败: "+err.Error())
	}
	if len(parentDocuments) == 0 {
		return errors.New("查找文档内容失败")
	}
	// get document content
	documentContent, err := utils.Document.GetContentByPageFile(pageFile)
	if err != nil {
		return errors.New("查找文档内容失败: "+err.Error())
	}

	documentContent = string([]byte(documentContent)[:500])

	documentValue := document
	documentValue["content"] = documentContent
	documentValue["username"] = username
	documentValue["comment"] = comment
	documentValue["url"] = url

	emailTemplate := beego.BConfig.WebConfig.ViewsPath+"system/email/template.html"
	body, err := utils.Email.MakeDocumentHtmlBody(documentValue, emailTemplate)
	if err != nil {
		return errors.New("发送邮件生成模板失败："+err.Error())
	}
	// start send email
	return utils.Email.SendByEmail(emailConfig, emails, "文档更新通知", body ,"html")
}