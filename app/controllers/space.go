package controllers

import (
	"strings"
	"time"

	"github.com/phachon/mm-wiki/app/models"
)

type SpaceController struct {
	BaseController
}

func (this *SpaceController) Index() {

	// get space tags
	spaceTags := map[string]string{}
	spaces, err := models.SpaceModel.GetSpaces()
	if err == nil {
		for _, space := range spaces {
			tags := space["tags"]
			if tags == "" {
				continue
			}
			tagList := strings.Split(tags, ",")
			for _, tagName := range tagList {
				spaceTags[tagName] = tagName
			}
		}
	}

	this.Data["spaceTags"] = spaceTags
	this.viewLayout("space/index", "space")
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
		this.ErrorLog("获取全部空间列表失败: " + err.Error())
		this.ViewError("获取空间列表失败", "/main/index")
	}

	collectionSpaces, err := models.CollectionModel.GetCollectionsByUserIdAndType(this.UserId, models.Collection_Type_Space)
	if err != nil {
		this.ErrorLog("获取全部空间列表失败: " + err.Error())
		this.ViewError("获取全部空间列表失败", "/main/index")
	}
	for _, space := range spaces {
		space["collection"] = "0"
		space["collection_id"] = "0"
		for _, collectionSpace := range collectionSpaces {
			if collectionSpace["resource_id"] == space["space_id"] {
				space["collection"] = "1"
				space["collection_id"] = collectionSpace["collection_id"]
				break
			}
		}
	}

	this.Data["spaces"] = spaces
	this.Data["keyword"] = keyword
	this.Data["count"] = count
	this.SetPaginator(number, count)
	this.viewLayout("space/list", "default")
}

func (this *SpaceController) Member() {

	page, _ := this.GetInt("page", 1)
	spaceId := strings.TrimSpace(this.GetString("space_id", ""))
	number, _ := this.GetRangeInt("number", 20, 10, 100)

	if spaceId == "" {
		this.ViewError("没有选择空间！")
	}
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("查找空间 " + spaceId + " 失败：" + err.Error())
		this.ViewError("空间不存在！")
	}
	if len(space) == 0 {
		this.ViewError("空间不存在！")
	}

	limit := (page - 1) * number

	count, err := models.SpaceUserModel.CountSpaceUsersBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("获取空间 " + spaceId + " 成员列表失败: " + err.Error())
		this.ViewError("获取空间成员列表失败！", "/space/index")
	}
	spaceUsers, err := models.SpaceUserModel.GetSpaceUsersBySpaceIdAndLimit(spaceId, limit, number)
	if err != nil {
		this.ErrorLog("获取空间 " + spaceId + " 成员列表失败: " + err.Error())
		this.ViewError("获取空间成员列表失败！", "/space/index")
	}

	var userIds = []string{}
	for _, spaceUser := range spaceUsers {
		userIds = append(userIds, spaceUser["user_id"])
	}
	users, err := models.UserModel.GetUsersByUserIds(userIds)
	if err != nil {
		this.ErrorLog("获取空间 " + spaceId + " 成员列表失败: " + err.Error())
		this.ViewError("获取空间成员列表失败！", "/space/index")
	}
	for _, user := range users {
		for _, spaceUser := range spaceUsers {
			if spaceUser["user_id"] == user["user_id"] {
				user["space_privilege"] = spaceUser["privilege"]
				user["space_user_id"] = spaceUser["space_user_id"]
			}
		}
	}
	this.Data["users"] = users
	this.Data["space_id"] = spaceId
	this.SetPaginator(number, count)

	// check user space privilege
	_, _, isManager := this.GetDocumentPrivilege(space)
	if isManager {
		var otherUsers = []map[string]string{}
		if len(userIds) > 0 {
			otherUsers, err = models.UserModel.GetUserByNotUserIds(userIds)
		} else {
			otherUsers, err = models.UserModel.GetUsers()
		}
		if err != nil {
			this.ErrorLog("获取空间 " + spaceId + " 成员列表失败: " + err.Error())
			this.ViewError("获取空间成员列表失败！", "/space/index")
		}
		this.Data["otherUsers"] = otherUsers
		this.viewLayout("space/manager_member", "default")
	} else {
		this.viewLayout("space/member", "default")
	}
}

func (this *SpaceController) AddMember() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/space/index")
	}
	spaceId := strings.TrimSpace(this.GetString("space_id", ""))
	userId := this.GetString("user_id", "")
	privilege := strings.TrimSpace(this.GetString("privilege", "0"))

	if spaceId == "" {
		this.jsonError("空间不存在！")
	}
	if userId == "" {
		this.jsonError("没有选择用户！")
	}
	if privilege == "" {
		this.jsonError("没有选择用户空间权限！")
	}
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("添加空间 " + spaceId + " 成员失败: " + err.Error())
		this.jsonError("添加空间成员失败！")
	}
	if len(space) == 0 {
		this.jsonError("空间不存在！")
	}

	// check login user space member privilege
	_, _, isManager := this.GetDocumentPrivilege(space)
	if !isManager {
		this.jsonError("您没有权限添加该空间成员！", "/space/index")
	}

	spaceUser, err := models.SpaceUserModel.GetSpaceUserBySpaceIdAndUserId(spaceId, userId)
	if err != nil {
		this.ErrorLog("添加空间 " + spaceId + " 成员 " + userId + " 失败: " + err.Error())
		this.jsonError("添加空间成员失败！")
	}
	if len(spaceUser) > 0 {
		this.jsonError("该用户已经是空间成员！")
	}

	insertValue := map[string]interface{}{
		"user_id":     userId,
		"space_id":    spaceId,
		"privilege":   privilege,
		"create_time": time.Now().Unix(),
		"update_time": time.Now().Unix(),
	}
	_, err = models.SpaceUserModel.Insert(insertValue)
	if err != nil {
		this.ErrorLog("添加空间 " + spaceId + " 成员 " + userId + " 失败: " + err.Error())
		this.jsonError("添加成员失败！")
	}

	this.InfoLog("空间 " + spaceId + " 添加成员 " + userId + " 成功")
	this.jsonSuccess("添加成员成功！", nil, "/space/member?space_id="+spaceId)
}

func (this *SpaceController) RemoveMember() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/space/index")
	}
	spaceId := this.GetString("space_id", "")
	userId := this.GetString("user_id", "")
	spaceUserId := this.GetString("space_user_id", "")

	if spaceUserId == "" {
		this.jsonError("空间成员不存在！")
	}
	if spaceId == "" {
		this.jsonError("空间不存在！")
	}
	if userId == "" {
		this.jsonError("用户不存在！")
	}

	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("移除空间 " + spaceId + " 成员失败: " + err.Error())
		this.jsonError("移除空间成员失败！")
	}
	if len(space) == 0 {
		this.jsonError("空间不存在！")
	}

	// check login user space member privilege
	_, _, isManager := this.GetDocumentPrivilege(space)
	if !isManager {
		this.ViewError("您没有权限移除该空间成员！", "/space/index")
	}

	err = models.SpaceUserModel.Delete(spaceUserId)
	if err != nil {
		this.ErrorLog("移除空间 " + spaceId + " 下成员 " + userId + " 失败：" + err.Error())
		this.jsonError("移除成员失败！")
	}

	this.InfoLog("移除空间 " + spaceId + " 下成员 " + userId + " 成功")
	this.jsonSuccess("移除成员成功！", nil, "/space/member?space_id="+spaceId)
}

func (this *SpaceController) ModifyMember() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/space/list")
	}
	spaceUserId := this.GetString("space_user_id", "")
	privilege := this.GetString("privilege", "0")
	userId := this.GetString("user_id", "")
	spaceId := this.GetString("space_id", "")

	if spaceUserId == "" {
		this.jsonError("更新权限错误！")
	}
	if privilege == "" {
		this.jsonError("没有选择权限！")
	}

	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("更新空间 " + spaceId + " 成员权限失败: " + err.Error())
		this.jsonError("更新空间成员权限失败！")
	}
	if len(space) == 0 {
		this.jsonError("空间不存在！")
	}

	// check login user space member privilege
	_, _, isManager := this.GetDocumentPrivilege(space)
	if !isManager {
		this.ViewError("您没有权限修改该空间成员！", "/space/index")
	}

	_, err = models.SpaceUserModel.Update(spaceUserId, map[string]interface{}{
		"privilege":   privilege,
		"update_time": time.Now().Unix(),
	})
	if err != nil {
		this.ErrorLog("更新空间 " + spaceId + " 下成员 " + userId + " 权限失败：" + err.Error())
		this.jsonError("更新权限失败！")
	}

	this.InfoLog("更新空间 " + spaceId + " 下成员 " + userId + " 权限成功")
	this.jsonSuccess("更新权限成功！", nil)
}

func (this *SpaceController) Collection() {

	collectionSpaces, err := models.CollectionModel.GetCollectionsByUserIdAndType(this.UserId, models.Collection_Type_Space)
	if err != nil {
		this.ErrorLog("获取收藏空间列表失败: " + err.Error())
		this.ViewError("获取收藏空间列表失败", "/space/list")
	}

	spaceIds := []string{}
	for _, collectionSpace := range collectionSpaces {
		spaceIds = append(spaceIds, collectionSpace["resource_id"])
	}

	spaces, err := models.SpaceModel.GetSpaceBySpaceIds(spaceIds)
	if err != nil {
		this.ErrorLog("获取收藏空间列表失败: " + err.Error())
		this.ViewError("获取收藏空间列表失败", "/space/list")
	}

	for _, space := range spaces {
		space["collection_id"] = "0"
		for _, collectionSpace := range collectionSpaces {
			if collectionSpace["resource_id"] == space["space_id"] {
				space["collection_id"] = collectionSpace["collection_id"]
				break
			}
		}
	}
	this.Data["spaces"] = spaces
	this.Data["count"] = len(spaces)
	this.viewLayout("space/collection", "default")
}

func (this *SpaceController) Document() {

	spaceId := this.GetString("space_id", "")
	if spaceId == "" {
		this.ViewError("空间不存在")
	}
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("查找空间 " + spaceId + " 失败：" + err.Error())
		this.ViewError("空间不存在！")
	}
	if len(space) == 0 {
		this.ViewError("空间不存在！")
	}

	// check space visit_level
	isVisit, _, _ := this.GetDocumentPrivilege(space)
	if !isVisit {
		this.ViewError("您没有权限访问该空间！")
	}

	spaceDocument, err := models.DocumentModel.GetSpaceDefaultDocument(spaceId)
	if err != nil {
		this.ErrorLog("查找空间 " + spaceId + " 失败：" + err.Error())
		this.ViewError("查找空间页面失败！")
	}
	if len(spaceDocument) == 0 {
		this.ViewError("空间首页不存在！")
	}

	documentId := spaceDocument["document_id"]

	this.Redirect("/document/index?document_id="+documentId, 302)
}

func (this *SpaceController) Search() {

	tagName := strings.TrimSpace(this.GetString("tag", ""))

	spaces, err := models.SpaceModel.GetSpacesByTags(tagName)
	if err != nil {
		this.ErrorLog("搜索空间标签列表失败: " + err.Error())
		this.ViewError("获取空间列表失败", "/main/index")
	}

	collectionSpaces, err := models.CollectionModel.GetCollectionsByUserIdAndType(this.UserId, models.Collection_Type_Space)
	if err != nil {
		this.ErrorLog("获取全部空间列表失败: " + err.Error())
		this.ViewError("获取全部空间列表失败", "/main/index")
	}
	for _, space := range spaces {
		space["collection"] = "0"
		space["collection_id"] = "0"
		for _, collectionSpace := range collectionSpaces {
			if collectionSpace["resource_id"] == space["space_id"] {
				space["collection"] = "1"
				space["collection_id"] = collectionSpace["collection_id"]
				break
			}
		}
	}

	this.Data["tag"] = tagName
	this.Data["spaces"] = spaces
	this.Data["count"] = len(spaces)
	this.viewLayout("space/search", "default")
}
