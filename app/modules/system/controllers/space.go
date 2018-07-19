package controllers

import (
	"strings"
	"mm-wiki/app/models"
	"mm-wiki/app/utils"
	"fmt"
	"regexp"
	"os"
)

type SpaceController struct {
	BaseController
}

func (this *SpaceController) Add() {
	this.viewLayout("space/form", "default")
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
		this.ErrorLog("添加空间失败："+err.Error())
		this.jsonError("添加空间失败！")
	}
	if ok {
		this.jsonError("空间名已经存在！")
	}

	// create space database
	spaceId, err := models.SpaceModel.Insert(map[string]interface{}{
		"name": name,
		"description": description,
		"tags": tags,
		"visit_level": strings.ToLower(visitLevel),
		"is_share": isShare,
		"is_export": isExport,
	})
	if err != nil {
		this.ErrorLog("添加空间失败：" + err.Error())
		this.jsonError("添加空间失败")
	}

	// create space document
	spaceDocument := map[string]interface{}{
		"space_id": fmt.Sprintf("%d", spaceId),
		"parent_id": "0",
		"name": name,
		"type": models.Document_Type_Dir,
		"path": "0",
		"create_user_id": this.UserId,
		"edit_user_id": this.UserId,
	}
	_, err = models.DocumentModel.Insert(spaceDocument)
	if err != nil {
		// delete space
		models.SpaceModel.Delete(fmt.Sprintf("%d", spaceId))
		this.ErrorLog("添加空间文档失败："+err.Error())
		this.jsonError("添加空间失败！")
	}

	this.InfoLog("添加空间 "+utils.Convert.IntToString(spaceId, 10)+" 成功")
	this.jsonSuccess("添加空间成功", nil, "/system/space/list")
}

func (this *SpaceController) List() {

	page, _ := this.GetInt("page", 1)
	keyword := strings.TrimSpace(this.GetString("keyword", ""))

	number := 20
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
		this.ErrorLog("获取空间列表失败: "+err.Error())
		this.ViewError("获取空间列表失败", "/system/main/index")
	}

	this.Data["spaces"] = spaces
	this.Data["keyword"] = keyword
	this.SetPaginator(number, count)
	this.viewLayout("space/list", "default")
}

func (this *SpaceController) Edit() {

	spaceId := this.GetString("space_id", "")
	if spaceId == "" {
		this.ViewError("空间不存在", "/system/space/list")
	}

	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("查找空间失败: "+err.Error())
		this.ViewError("查找空间失败", "/system/space/list")
	}
	if len(space) == 0 {
		this.ViewError("空间不存在", "/system/space/list")
	}

	this.Data["space"] = space
	this.viewLayout("space/form", "default")
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
		this.ErrorLog("修改空间 "+spaceId+" 失败: "+err.Error())
		this.jsonError("修改空间失败！")
	}
	if len(space) == 0 {
		this.jsonError("空间不存在！")
	}

	ok , _ := models.SpaceModel.HasSameName(spaceId, name)
	if ok {
		this.jsonError("空间名已经存在！")
	}
	_, err = models.SpaceModel.Update(spaceId, map[string]interface{}{
		"name": name,
		"description": description,
		"tags": tags,
		"visit_level": visitLevel,
		"is_share": isShare,
		"is_export": isExport,
	})

	if err != nil {
		this.ErrorLog("修改空间 "+spaceId+" 失败：" + err.Error())
		this.jsonError("修改空间失败")
	}
	this.InfoLog("修改空间 "+spaceId+" 成功")
	this.jsonSuccess("修改空间成功", nil, "/system/space/list")
}

func (this *SpaceController) Member() {

	page, _ := this.GetInt("page", 1)
	spaceId := strings.TrimSpace(this.GetString("space_id", ""))

	if spaceId == "" {
		this.ViewError("没有选择空间！")
	}

	number := 20
	limit := (page - 1) * number

	count, err := models.SpaceUserModel.CountSpaceUsersBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("获取空间 "+spaceId+" 成员列表失败: "+err.Error())
		this.ViewError("获取空间成员列表失败！", "/system/space/list")
	}
	spaceUsers, err := models.SpaceUserModel.GetSpaceUsersBySpaceIdAndLimit(spaceId, limit, number)
	if err != nil {
		this.ErrorLog("获取空间 "+spaceId+" 成员列表失败: "+err.Error())
		this.ViewError("获取空间成员列表失败！", "/system/space/list")
	}

	var userIds = []string{}
	for _, spaceUser := range spaceUsers {
		userIds = append(userIds, spaceUser["user_id"])
	}
	users, err := models.UserModel.GetUsersByUserIds(userIds)
	if err != nil {
		this.ErrorLog("获取空间 "+spaceId+" 成员列表失败: "+err.Error())
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
	}else {
		otherUsers, err = models.UserModel.GetUsers()
	}
	if err != nil {
		this.ErrorLog("获取空间 "+spaceId+" 成员列表失败: "+err.Error())
		this.ViewError("获取空间成员列表失败！", "/system/main/index")
	}

	this.Data["users"] = users
	this.Data["space_id"] = spaceId
	this.Data["otherUsers"] = otherUsers
	this.SetPaginator(number, count)
	this.viewLayout("space/member", "default")
}

func (this *SpaceController) Delete() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/space/list")
	}
	spaceId := this.GetString("space_id", "")
	if spaceId == "" {
		this.jsonError("没有选择空间！")
	}

	space, err := models.RoleModel.GetRoleByRoleId(spaceId)
	if err != nil {
		this.ErrorLog("删除空间 "+spaceId+" 失败: "+err.Error())
		this.jsonError("删除空间失败")
	}
	if len(space) == 0 {
		this.jsonError("空间不存在")
	}

	// check space user
	spaceUsers, err := models.SpaceUserModel.GetSpaceUsersBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("删除空间 "+spaceId+" 失败: "+err.Error())
		this.jsonError("删除空间失败")
	}
	if len(spaceUsers) > 0 {
		this.jsonError("不能删除空间，请先移除该空间下用户成员!")
	}

	// check space documents
	documents, err := models.DocumentModel.GetDocumentsBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("删除空间 "+spaceId+" 失败: "+err.Error())
		this.jsonError("删除空间失败")
	}
	if len(documents) > 0 {
		this.jsonError("不能删除空间，请先删除该空间下文档!")
	}

	// delete space
	err = models.RoleModel.Delete(spaceId)
	if err != nil {
		this.ErrorLog("删除空间 "+spaceId+" 失败: "+err.Error())
		this.jsonError("删除空间失败")
	}

	this.InfoLog("删除空间 "+spaceId+" 成功")
	this.jsonSuccess("删除空间成功", nil, "/system/space/list")
}

func (this *SpaceController) Download() {

	spaceId := this.GetString("space_id", "")
	if spaceId == "" {
		this.ViewError("空间不存在", "/system/space/list")
	}

	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("查找空间失败: "+err.Error())
		this.ViewError("查找空间失败", "/system/space/list")
	}
	if len(space) == 0 {
		this.ViewError("空间不存在", "/system/space/list")
	}

	spaceName := space["name"]
	spacePath := utils.Document.GetAbsPageFileByPageFile(spaceName)

	f3, err := os.Open(spacePath)
	if err != nil {
		this.ErrorLog("下载空间 "+spaceId+" 打开空间目录失败："+err.Error())
		this.ViewError("下载空间文档失败")
	}
	defer f3.Close()
	var files = []*os.File{f3}

	tempDir := os.TempDir()+"/"+"mmwiki"
	os.RemoveAll(tempDir)
	os.Mkdir(tempDir, 0777)
	var dest = tempDir+"/"+spaceName+".zip"

	fmt.Println(dest)
	err = utils.Zipx.Compress(files, dest)
	if err != nil {
		this.ErrorLog("下载空间 "+spaceId+" 压缩文档失败："+err.Error())
		this.ViewError("空间文档压缩失败")
	}

	this.Ctx.Output.Download(dest, spaceName+".zip")
}