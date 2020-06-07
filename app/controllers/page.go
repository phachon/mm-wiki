package controllers

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/phachon/mm-wiki/app"
	"github.com/phachon/mm-wiki/app/models"
	"github.com/phachon/mm-wiki/app/services"
	"github.com/phachon/mm-wiki/app/utils"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/astaxie/beego/logs"
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
		this.ErrorLog("查找文档 " + documentId + " 失败：" + err.Error())
		this.ViewError("查找文档失败！")
	}
	if len(document) == 0 {
		this.ViewError("文档不存在！")
	}

	spaceId := document["space_id"]
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("查找文档 " + documentId + " 所在空间失败：" + err.Error())
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
		this.ErrorLog("查找父文档失败：" + err.Error())
		this.ViewError("查找父文档失败！")
	}
	if len(parentDocuments) == 0 {
		this.ViewError("父文档不存在！")
	}

	// get document content
	documentContent, err := utils.Document.GetContentByPageFile(pageFile)
	if err != nil {
		this.ErrorLog("查找文档 " + documentId + " 失败：" + err.Error())
		this.ViewError("文档不存在！")
	}

	// get edit user and create user
	users, err := models.UserModel.GetUsersByUserIds([]string{document["create_user_id"], document["edit_user_id"]})
	if err != nil {
		this.ErrorLog("查找文档 " + documentId + " 失败：" + err.Error())
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
		this.ErrorLog("查找文档 " + documentId + " 失败：" + err.Error())
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
		this.ErrorLog("修改文档 " + documentId + " 失败：" + err.Error())
		this.ViewError("修改文档失败！")
	}
	if len(document) == 0 {
		this.ViewError("文档不存在！")
	}

	spaceId := document["space_id"]
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("修改文档 " + documentId + " 失败：" + err.Error())
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
		this.ErrorLog("查找父文档失败：" + err.Error())
		this.ViewError("查找父文档失败！")
	}

	// get document content
	documentContent, err := utils.Document.GetContentByPageFile(pageFile)
	if err != nil {
		this.ErrorLog("查找文档 " + documentId + " 失败：" + err.Error())
		this.ViewError("文档不存在！")
	}

	autoFollowDoc := models.ConfigModel.GetConfigValueByKey(models.ConfigKeyAutoFollowdoc, "0")
	sendEmail := models.ConfigModel.GetConfigValueByKey(models.ConfigKeySendEmail, "0")

	this.Data["sendEmail"] = sendEmail
	this.Data["autoFollowDoc"] = autoFollowDoc
	this.Data["page_content"] = documentContent
	this.Data["document"] = document
	this.viewLayout("page/edit", "document_page")
}

// page modify
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
		this.jsonError("文档名称不能为 " + utils.Document_Default_FileName + " ！")
	}
	//if comment == "" {
	//	this.jsonError("必须输入此次修改的备注！")
	//}

	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("修改文档 " + documentId + " 失败：" + err.Error())
		this.jsonError("保存文档失败！")
	}
	if len(document) == 0 {
		this.jsonError("文档不存在！")
	}

	spaceId := document["space_id"]
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("修改文档 " + documentId + " 失败：" + err.Error())
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
			this.ErrorLog("修改文档失败：" + err.Error())
			this.jsonError("保存文档失败！")
		}
		if len(newDocument) != 0 {
			this.jsonError("该文档名称已经存在！")
		}
	}

	// update document and file content
	updateValue := map[string]interface{}{
		"name":         newName,
		"edit_user_id": this.UserId,
	}
	_, err = models.DocumentModel.UpdateDBAndFile(documentId, document, documentContent, updateValue, comment)
	if err != nil {
		this.ErrorLog("修改文档 " + documentId + " 失败：" + err.Error())
		this.jsonError("修改文档失败！")
	}

	// send email to follow user
	if isNoticeUser == "1" {
		logInfo := this.GetLogInfoByCtx()
		url := fmt.Sprintf("%s:%d/document/index?document_id=%s", this.Ctx.Input.Site(), this.Ctx.Input.Port(), documentId)
		go func(documentId string, username string, comment string, url string) {
			err := sendEmail(documentId, username, comment, url)
			if err != nil {
				logInfo["message"] = "更新文档时发送邮件通知失败：" + err.Error()
				logInfo["level"] = models.Log_Level_Error
				models.LogModel.Insert(logInfo)
				logs.Error("更新文档时发送邮件通知失败：" + err.Error())
			}
		}(documentId, this.User["username"], comment, url)
	}
	// follow doc
	if isFollowDoc == "1" {
		go func() {
			_, _ = models.FollowModel.FollowDocument(this.UserId, documentId)
		}()
	}
	// 更新文档索引
	go func(documentId string) {
		_ = services.DocIndexService.ForceUpdateDocIndexByDocId(documentId)
	}(documentId)

	this.InfoLog("修改文档 " + documentId + " 成功")
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
		this.ErrorLog("查找文档 " + documentId + " 失败：" + err.Error())
		this.ViewError("查找文档失败！")
	}
	if len(document) == 0 {
		this.ViewError("文档不存在！")
	}

	// get document space
	spaceId := document["space_id"]
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("分享文档 " + documentId + " 失败：" + err.Error())
		this.ViewError("保存文档失败！")
	}
	if len(space) == 0 {
		this.ViewError("文档所在空间不存在！")
	}

	// check space is allow display
	if space["is_share"] != fmt.Sprintf("%d", models.Space_Share_True) {
		this.ViewError("该文档不能被分享！")
	}

	// get parent documents by document
	parentDocuments, pageFile, err := models.DocumentModel.GetParentDocumentsByDocument(document)
	if err != nil {
		this.ErrorLog("查找父文档失败：" + err.Error())
		this.ViewError("查找父文档失败！")
	}
	if len(parentDocuments) == 0 {
		this.ViewError("父文档不存在！")
	}

	// get document content
	documentContent, err := utils.Document.GetContentByPageFile(pageFile)
	if err != nil {
		this.ErrorLog("查找文档 " + documentId + " 内容失败：" + err.Error())
		this.ViewError("文档不存在！")
	}

	// get edit user and create user
	users, err := models.UserModel.GetUsersByUserIds([]string{document["create_user_id"], document["edit_user_id"]})
	if err != nil {
		this.ErrorLog("查找文档 " + documentId + " 失败：" + err.Error())
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

// export file
func (this *PageController) Export() {

	documentId := this.GetString("document_id", "")
	if documentId == "" {
		this.ViewError("文档未找到！")
	}

	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("查找文档 " + documentId + " 失败：" + err.Error())
		this.ViewError("查找文档失败！")
	}
	if len(document) == 0 {
		this.ViewError("文档不存在！")
	}

	spaceId := document["space_id"]
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("查找文档 " + documentId + " 所在空间失败：" + err.Error())
		this.ViewError("查找文档失败！")
	}
	if len(space) == 0 {
		this.ViewError("文档所在空间不存在！")
	}

	// check space document privilege
	isVisit, _, _ := this.GetDocumentPrivilege(space)
	if !isVisit {
		this.ViewError("您没有权限导出该空间下文档！")
	}

	// check space is allow export
	if space["is_export"] != fmt.Sprintf("%d", models.Space_Download_True) {
		this.ViewError("该文档不允许被导出！")
	}

	// get parent documents by document
	parentDocuments, pageFile, err := models.DocumentModel.GetParentDocumentsByDocument(document)
	if err != nil {
		this.ErrorLog("查找父文档失败：" + err.Error())
		this.ViewError("查找父文档失败！")
	}
	if len(parentDocuments) == 0 {
		this.ViewError("父文档不存在！")
	}

	packFiles := []*utils.CompressFileInfo{}

	absPageFile := utils.Document.GetAbsPageFileByPageFile(pageFile)
	// pack document file
	packFiles = append(packFiles, &utils.CompressFileInfo{
		File:       absPageFile,
		PrefixPath: "",
	})

	// get document attachments
	attachments, err := models.AttachmentModel.GetAttachmentsByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("查找文档附件失败：" + err.Error())
		this.ViewError("查找文档附件失败！")
	}
	for _, attachment := range attachments {
		if attachment["path"] == "" {
			continue
		}
		path := attachment["path"]
		attachmentFile := filepath.Join(app.DocumentAbsDir, path)
		packFile := &utils.CompressFileInfo{
			File:       attachmentFile,
			PrefixPath: filepath.Dir(path),
		}
		packFiles = append(packFiles, packFile)
	}
	var dest = fmt.Sprintf("%s/mm_wiki/%s.zip", os.TempDir(), document["name"])
	err = utils.Zipx.PackFile(packFiles, dest)
	if err != nil {
		this.ErrorLog("导出文档附件失败：" + err.Error())
		this.ViewError("导出文档失败！")
	}

	this.Ctx.Output.Download(dest, document["name"]+".zip")
}

func sendEmail(documentId string, username string, comment string, url string) error {

	// get document by documentId
	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		return errors.New("发送邮件通知查找文档失败：" + err.Error())
	}

	// get send email open config
	sendEmailConfig := models.ConfigModel.GetConfigValueByKey(models.ConfigKeySendEmail, "0")
	if sendEmailConfig == "0" {
		return nil
	}

	// get email config
	emailConfig, err := models.EmailModel.GetUsedEmail()
	if err != nil {
		return errors.New("发送邮件通知查找邮件服务器配置失败：" + err.Error())
	}
	if len(emailConfig) == 0 {
		return nil
	}

	// get follow doc user
	follows, err := models.FollowModel.GetFollowsByObjectIdAndType(documentId, models.Follow_Type_Doc)
	if err != nil {
		return errors.New("发送邮件查找关注文档用户失败：" + err.Error())
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
		return errors.New("发送邮件查找关注文档用户失败：" + err.Error())
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
		return errors.New("查找文档内容失败: " + err.Error())
	}
	if len(parentDocuments) == 0 {
		return errors.New("查找文档内容失败")
	}
	// get document content
	documentContent, err := utils.Document.GetContentByPageFile(pageFile)
	if err != nil {
		return errors.New("查找文档内容失败: " + err.Error())
	}

	if len([]byte(documentContent)) > 500 {
		documentContent = string([]byte(documentContent)[:500])
	}

	documentValue := document
	documentValue["content"] = documentContent
	documentValue["username"] = username
	documentValue["comment"] = comment
	documentValue["url"] = url

	emailTemplate := beego.BConfig.WebConfig.ViewsPath + "/system/email/template.html"
	body, err := utils.Email.MakeDocumentHtmlBody(documentValue, emailTemplate)
	if err != nil {
		return errors.New("发送邮件生成模板失败：" + err.Error())
	}
	// start send email
	return utils.Email.Send(emailConfig, emails, "文档更新通知", body)
}
