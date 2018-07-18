package controllers

import (
	"strings"
	"mm-wiki/app/models"
	"mm-wiki/app/utils"
	"fmt"
	"github.com/astaxie/beego/validation"
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

	v := validation.Validation{}
	if name == "" {
		this.jsonError("空间名称不能为空！")
	}
	if !v.AlphaDash(name, "name").Ok {
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

	v := validation.Validation{}
	if spaceId == "" {
		this.jsonError("空间不存在！")
	}
	if name == "" {
		this.jsonError("空间名称不能为空！")
	}
	if !v.AlphaDash(name, "name").Ok {
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