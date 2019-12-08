package controllers

import (
	"fmt"
	"github.com/phachon/mm-wiki/app"
	"github.com/phachon/mm-wiki/app/models"
	"github.com/phachon/mm-wiki/app/utils"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type SpaceController struct {
	BaseController
}

func (this *SpaceController) Add() {
	this.viewLayout("space/form", "space")
}

func (this *SpaceController) Save() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/space/list")
	}
	name := strings.TrimSpace(this.GetString("name", ""))
	description := strings.TrimSpace(this.GetString("description", ""))
	tags := strings.TrimSpace(this.GetString("tags", ""))
	visitLevel := strings.TrimSpace(this.GetString("visit_level", "public"))
	isShare := strings.TrimSpace(this.GetString("is_share", "1"))
	isExport := strings.TrimSpace(this.GetString("is_export", "0"))

	if name == "" {
		this.jsonError("空间名称不能为空！")
	}
	match, err := regexp.MatchString(`[\\\\/:*?\"<>、|]`, name)
	if err != nil {
		this.jsonError("空间名称格式不正确！")
	}
	if match {
		this.jsonError("空间名称格式不正确！")
	}
	ok, err := models.SpaceModel.HasSpaceName(name)
	if err != nil {
		this.ErrorLog("添加空间失败：" + err.Error())
		this.jsonError("添加空间失败！")
	}
	if ok {
		this.jsonError("空间名已经存在！")
	}

	// create space database
	spaceId, err := models.SpaceModel.Insert(map[string]interface{}{
		"name":        name,
		"description": description,
		"tags":        tags,
		"visit_level": strings.ToLower(visitLevel),
		"is_share":    isShare,
		"is_export":   isExport,
	})
	if err != nil {
		this.ErrorLog("添加空间失败：" + err.Error())
		this.jsonError("添加空间失败")
	}

	// create space document
	spaceDocument := map[string]interface{}{
		"space_id":       fmt.Sprintf("%d", spaceId),
		"parent_id":      "0",
		"name":           name,
		"type":           models.Document_Type_Dir,
		"path":           "0",
		"create_user_id": this.UserId,
		"edit_user_id":   this.UserId,
	}
	_, err = models.DocumentModel.Insert(spaceDocument)
	if err != nil {
		// delete space
		models.SpaceModel.Delete(fmt.Sprintf("%d", spaceId))
		this.ErrorLog("添加空间文档失败：" + err.Error())
		this.jsonError("添加空间失败！")
	}

	// add space member
	insertValue := map[string]interface{}{
		"user_id":   this.UserId,
		"space_id":  spaceId,
		"privilege": models.SpaceUser_Privilege_Manager,
	}
	_, err = models.SpaceUserModel.Insert(insertValue)
	if err != nil {
		// delete space
		models.SpaceModel.Delete(fmt.Sprintf("%d", spaceId))
		this.ErrorLog("添加空间添加空间成员失败: " + err.Error())
		this.jsonError("添加空间失败！")
	}

	this.InfoLog("添加空间 " + utils.Convert.IntToString(spaceId, 10) + " 成功")
	this.jsonSuccess("添加空间成功", nil, "/system/space/list")
}

func (this *SpaceController) List() {

	page, _ := this.GetInt("page", 1)
	keyword := strings.TrimSpace(this.GetString("keyword", ""))
	number, _ := this.GetRangeInt("number", 20, 10, 100)
	limit := (page - 1) * number

	var err error
	var count int64
	var spaces []map[string]string
	if keyword != "" {
		count, err = models.SpaceModel.CountSpacesByKeyword(keyword)
		spaces, err = models.SpaceModel.GetSpacesByKeywordAndLimit(keyword, limit, number)
	} else {
		count, err = models.SpaceModel.CountSpaces()
		spaces, err = models.SpaceModel.GetSpacesByLimit(limit, number)
	}
	if err != nil {
		this.ErrorLog("获取空间列表失败: " + err.Error())
		this.ViewError("获取空间列表失败", "/system/main/index")
	}

	this.Data["spaces"] = spaces
	this.Data["keyword"] = keyword
	this.SetPaginator(number, count)
	this.viewLayout("space/list", "space")
}

func (this *SpaceController) Edit() {

	spaceId := this.GetString("space_id", "")
	if spaceId == "" {
		this.ViewError("空间不存在", "/system/space/list")
	}

	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("查找空间失败: " + err.Error())
		this.ViewError("查找空间失败", "/system/space/list")
	}
	if len(space) == 0 {
		this.ViewError("空间不存在", "/system/space/list")
	}

	this.Data["space"] = space
	this.viewLayout("space/form", "space")
}

func (this *SpaceController) Modify() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/space/list")
	}
	spaceId := this.GetString("space_id", "")
	name := strings.TrimSpace(this.GetString("name", ""))
	description := strings.TrimSpace(this.GetString("description", ""))
	tags := strings.TrimSpace(this.GetString("tags", ""))
	visitLevel := strings.TrimSpace(this.GetString("visit_level", "public"))
	isShare := strings.TrimSpace(this.GetString("is_share", "0"))
	isExport := strings.TrimSpace(this.GetString("is_export", "0"))

	if spaceId == "" {
		this.jsonError("空间不存在！")
	}
	if name == "" {
		this.jsonError("空间名称不能为空！")
	}
	match, err := regexp.MatchString(`[\\\\/:*?\"<>、|]`, name)
	if err != nil {
		this.jsonError("空间名称格式不正确！")
	}
	if match {
		this.jsonError("空间名称格式不正确！")
	}

	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("修改空间 " + spaceId + " 失败: " + err.Error())
		this.jsonError("修改空间失败！")
	}
	if len(space) == 0 {
		this.jsonError("空间不存在！")
	}

	ok, _ := models.SpaceModel.HasSameName(spaceId, name)
	if ok {
		this.jsonError("空间名已经存在！")
	}

	spaceValue := map[string]interface{}{
		"name":        name,
		"description": description,
		"tags":        tags,
		"visit_level": visitLevel,
		"is_share":    isShare,
		"is_export":   isExport,
	}
	// update space document dir name if name update
	_, err = models.SpaceModel.UpdateDBAndSpaceFileName(spaceId, spaceValue, space["name"])
	if err != nil {
		this.ErrorLog("修改空间 " + spaceId + " 失败：" + err.Error())
		this.jsonError("修改空间失败")
	}
	this.InfoLog("修改空间 " + spaceId + " 成功")
	this.jsonSuccess("修改空间成功", nil, "/system/space/list")
}

func (this *SpaceController) Member() {

	page, _ := this.GetInt("page", 1)
	spaceId := strings.TrimSpace(this.GetString("space_id", ""))
	number, _ := this.GetRangeInt("number", 15, 10, 100)

	if spaceId == "" {
		this.ViewError("没有选择空间！")
	}

	limit := (page - 1) * number

	count, err := models.SpaceUserModel.CountSpaceUsersBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("获取空间 " + spaceId + " 成员列表失败: " + err.Error())
		this.ViewError("获取空间成员列表失败！", "/system/space/list")
	}
	spaceUsers, err := models.SpaceUserModel.GetSpaceUsersBySpaceIdAndLimit(spaceId, limit, number)
	if err != nil {
		this.ErrorLog("获取空间 " + spaceId + " 成员列表失败: " + err.Error())
		this.ViewError("获取空间成员列表失败！", "/system/space/list")
	}

	var userIds = []string{}
	for _, spaceUser := range spaceUsers {
		userIds = append(userIds, spaceUser["user_id"])
	}
	users, err := models.UserModel.GetUsersByUserIds(userIds)
	if err != nil {
		this.ErrorLog("获取空间 " + spaceId + " 成员列表失败: " + err.Error())
		this.ViewError("获取空间成员列表失败！", "/system/main/index")
	}
	for _, user := range users {
		for _, spaceUser := range spaceUsers {
			if spaceUser["user_id"] == user["user_id"] {
				user["space_privilege"] = spaceUser["privilege"]
				user["space_user_id"] = spaceUser["space_user_id"]
			}
		}
	}

	var otherUsers = []map[string]string{}
	if len(userIds) > 0 {
		otherUsers, err = models.UserModel.GetUserByNotUserIds(userIds)
	} else {
		otherUsers, err = models.UserModel.GetUsers()
	}
	if err != nil {
		this.ErrorLog("获取空间 " + spaceId + " 成员列表失败: " + err.Error())
		this.ViewError("获取空间成员列表失败！", "/system/main/index")
	}

	this.Data["users"] = users
	this.Data["space_id"] = spaceId
	this.Data["otherUsers"] = otherUsers
	this.SetPaginator(number, count)
	this.viewLayout("space/member", "space")
}

func (this *SpaceController) Delete() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/space/list")
	}
	spaceId := this.GetString("space_id", "")
	if spaceId == "" {
		this.jsonError("没有选择空间！")
	}

	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("删除空间 " + spaceId + " 失败: " + err.Error())
		this.jsonError("删除空间失败")
	}
	if len(space) == 0 {
		this.jsonError("空间不存在")
	}

	// check space documents
	documents, err := models.DocumentModel.GetDocumentsBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("删除空间 " + spaceId + " 失败: " + err.Error())
		this.jsonError("删除空间失败")
	}
	if len(documents) > 1 {
		this.jsonError("不能删除空间，请先删除该空间下文档!")
	} else if len(documents) == 1 {
		if documents[0]["name"] != space["name"] {
			this.jsonError("不能删除空间，请先删除该空间下文档!")
		} else {
			// delete space dir and documentId
			_, pageFile, err := models.DocumentModel.GetParentDocumentsByDocument(documents[0])
			if err != nil {
				this.ErrorLog("删除空间 " + spaceId + " 获取空间文件失败: " + err.Error())
				this.jsonError("删除空间失败")
			}
			err = models.DocumentModel.DeleteDBAndFile(documents[0]["document_id"], this.UserId, pageFile, fmt.Sprintf("%d", models.Document_Type_Dir))
			// delete space document attachments
			_ = models.AttachmentModel.DeleteAttachmentsDBFileByDocumentId(documents[0]["document_id"])
		}
	} else {
		// delete space dir
		err = utils.Document.DeleteSpace(space["name"])
	}
	if err != nil {
		this.ErrorLog("删除空间 " + spaceId + " 失败: " + err.Error())
		this.jsonError("删除空间失败")
	}

	// delete space user
	err = models.SpaceUserModel.DeleteBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("删除空间 " + spaceId + " 失败: " + err.Error())
		this.jsonError("删除空间失败")
	}
	// delete space and space document
	err = models.SpaceModel.Delete(spaceId)
	if err != nil {
		this.ErrorLog("删除空间 " + spaceId + " 失败: " + err.Error())
		this.jsonError("删除空间失败")
	}

	this.InfoLog("删除空间 " + spaceId + " 成功")
	this.jsonSuccess("删除空间成功", nil, "/system/space/list")
}

func (this *SpaceController) Download() {

	spaceId := this.GetString("space_id", "")
	if spaceId == "" {
		this.ViewError("空间不存在", "/system/space/list")
	}

	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("查找空间失败: " + err.Error())
		this.ViewError("查找空间失败", "/system/space/list")
	}
	if len(space) == 0 {
		this.ViewError("空间不存在", "/system/space/list")
	}

	spaceName := space["name"]
	spacePath := utils.Document.GetAbsPageFileByPageFile(spaceName)

	packFiles := []*utils.CompressFileInfo{}

	// pack space all markdown file
	packFiles = append(packFiles, &utils.CompressFileInfo{
		File:       spacePath,
		PrefixPath: "",
	})

	// get space all document attachments
	attachments, err := models.AttachmentModel.GetAttachmentsBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("查找空间文档附件失败：" + err.Error())
		this.ViewError("查找空间文档附件失败！")
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
	var dest = fmt.Sprintf("%s/mm_wiki/%s.zip", os.TempDir(), spaceName)
	err = utils.Zipx.PackFile(packFiles, dest)
	if err != nil {
		this.ErrorLog("下载空间文档失败：" + err.Error())
		this.ViewError("下载空间文档失败！")
	}

	this.Ctx.Output.Download(dest, spaceName+".zip")
}
