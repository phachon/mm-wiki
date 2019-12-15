package controllers

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/phachon/mm-wiki/app"
	"github.com/phachon/mm-wiki/app/models"
	"github.com/phachon/mm-wiki/app/utils"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"time"
)

type StaticController struct {
	BaseController
}

func (this *StaticController) Default() {

	normalUserCount, err := models.UserModel.CountNormalUsers()
	if err != nil {
		this.ErrorLog("查找正常用户数出错：" + err.Error())
		this.jsonError("查找数据出错")
	}

	forbiddenUserCount, err := models.UserModel.CountForbiddenUsers()
	if err != nil {
		this.ErrorLog("查找屏蔽用户数出错：" + err.Error())
		this.jsonError("查找数据出错")
	}

	spaceCount, err := models.SpaceModel.CountSpaces()
	if err != nil {
		this.ErrorLog("查找空间总数出错：" + err.Error())
		this.jsonError("查找数据出错")
	}

	documentCount, err := models.DocumentModel.CountDocuments()
	if err != nil {
		this.ErrorLog("查找文档总数出错：" + err.Error())
		this.jsonError("查找数据出错")
	}

	time.Now().Unix()
	today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local).Unix()
	todayLoginUserCount, err := models.UserModel.CountUsersByLastTime(today)
	if err != nil {
		this.ErrorLog("查找今日登录用户数出错：" + err.Error())
		this.jsonError("查找数据出错")
	}

	createUserId := "0"
	createUserIds, err := models.DocumentModel.GetDocumentGroupCreateUserId()
	if err != nil {
		this.ErrorLog("查找创建文档最多用户出错：" + err.Error())
		this.jsonError("查找数据出错")
	}
	if len(createUserIds) > 0 {
		createUserId = createUserIds[0]["create_user_id"]
	}

	editUserId := "0"
	editUserIds, err := models.DocumentModel.GetDocumentGroupEditUserId()
	if err != nil {
		this.ErrorLog("查找修改文档最多用户出错：" + err.Error())
		this.jsonError("查找数据出错")
	}
	if len(editUserIds) > 0 {
		editUserId = editUserIds[0]["edit_user_id"]
	}

	collectUserId := "0"
	collectUserIds, err := models.CollectionModel.GetCollectionGroupUserId(models.Collection_Type_Doc)
	if err != nil {
		this.ErrorLog("查找收藏文档最多用户出错：" + err.Error())
		this.jsonError("查找数据出错")
	}
	if len(collectUserIds) > 0 {
		collectUserId = collectUserIds[0]["user_id"]
	}

	fansUserId := "0"
	fansUserIds, err := models.FollowModel.GetFansUserGroupUserId()
	if err != nil {
		this.ErrorLog("查找粉丝数最多用户出错：" + err.Error())
		this.jsonError("查找数据出错")
	}
	if len(fansUserIds) > 0 {
		fansUserId = fansUserIds[0]["object_id"]
	}

	userIds := []string{createUserId, editUserId, collectUserId, fansUserId}
	users, err := models.UserModel.GetUsersByUserIds(userIds)
	if err != nil {
		this.ErrorLog("查找文档总数出错：" + err.Error())
		this.jsonError("查找数据出错")
	}

	var createMaxUser = map[string]string{"user_id": createUserId, "username": ""}
	var editMaxUser = map[string]string{"user_id": editUserId, "username": ""}
	var collectMaxUser = map[string]string{"user_id": collectUserId, "username": ""}
	var fansMaxUser = map[string]string{"user_id": fansUserId, "username": ""}
	for _, user := range users {
		if user["user_id"] == createUserId {
			createMaxUser["username"] = user["username"]
		}
		if user["user_id"] == editUserId {
			editMaxUser["username"] = user["username"]
		}
		if user["user_id"] == collectUserId {
			collectMaxUser["username"] = user["username"]
		}
		if user["user_id"] == fansUserId {
			fansMaxUser["username"] = user["username"]
		}
	}

	this.Data["normalUserCount"] = normalUserCount
	this.Data["forbiddenUserCount"] = forbiddenUserCount
	this.Data["spaceCount"] = spaceCount
	this.Data["documentCount"] = documentCount

	this.Data["todayLoginUserCount"] = todayLoginUserCount
	this.Data["createMaxUser"] = createMaxUser
	this.Data["editMaxUser"] = editMaxUser
	this.Data["collectMaxUser"] = collectMaxUser
	this.Data["fansMaxUser"] = fansMaxUser

	this.viewLayout("static/default", "static")
}

func (this *StaticController) SpaceDocsRank() {

	number, _ := this.GetInt("number", 15)
	spaceDocsRank := []map[string]string{}
	spaceDocCountIds, err := models.DocumentModel.GetSpaceIdsOrderByCountDocumentLimit(number)
	if err != nil {
		this.ErrorLog("查找空间文档排行出错：" + err.Error())
		this.jsonError("查找数据出错")
	}
	if len(spaceDocCountIds) > 0 {
		spaceIds := []string{}
		for _, spaceDocCountId := range spaceDocCountIds {
			spaceIds = append(spaceIds, spaceDocCountId["space_id"])
		}
		spaces, err := models.SpaceModel.GetSpaceBySpaceIds(spaceIds)
		if err != nil {
			this.ErrorLog("查找空间文档排行出错：" + err.Error())
			this.jsonError("查找数据出错")
		}
		for _, spaceDocCountId := range spaceDocCountIds {
			spaceDocCount := map[string]string{
				"space_name": "",
				"total":      spaceDocCountId["total"],
			}
			for _, space := range spaces {
				if spaceDocCountId["space_id"] == space["space_id"] {
					spaceDocCount["space_name"] = space["name"]
					break
				}
			}
			spaceDocsRank = append(spaceDocsRank, spaceDocCount)
		}
	}

	this.jsonSuccess("ok", spaceDocsRank)
}

func (this *StaticController) CollectDocRank() {

	number, _ := this.GetInt("number", 10)
	collectDocsRank := []map[string]string{}
	collectionDocIds, err := models.CollectionModel.GetResourceIdsOrderByCountLimit(number, models.Collection_Type_Doc)
	if err != nil {
		this.ErrorLog("查找收藏文档排行出错：" + err.Error())
		this.jsonError("查找数据出错")
	}
	if len(collectionDocIds) > 0 {
		documentIds := []string{}
		for _, collectionDocId := range collectionDocIds {
			documentIds = append(documentIds, collectionDocId["resource_id"])
		}
		documents, err := models.DocumentModel.GetDocumentsByDocumentIds(documentIds)
		if err != nil {
			this.ErrorLog("查找收藏文档排行出错：" + err.Error())
			this.jsonError("查找数据出错")
		}
		for _, collectionDocId := range collectionDocIds {
			collectDocCount := map[string]string{
				"document_name": "",
				"total":         collectionDocId["total"],
			}
			for _, document := range documents {
				if collectionDocId["resource_id"] == document["document_id"] {
					collectDocCount["document_name"] = document["name"]
					break
				}
			}
			collectDocsRank = append(collectDocsRank, collectDocCount)
		}
	}

	this.jsonSuccess("ok", collectDocsRank)
}

func (this *StaticController) DocCountByTime() {

	limitDay, _ := this.GetInt("limit_day", 10)
	startTime := time.Now().Unix() - int64(limitDay*3600*24)

	documentCountDate, err := models.DocumentModel.GetCountGroupByCreateTime(startTime)
	if err != nil {
		this.ErrorLog("查找文档数量增长趋势出错：" + err.Error())
		this.jsonError("查找文档数量增长趋势出错")
	}

	this.jsonSuccess("ok", documentCountDate)
}

func (this *StaticController) Monitor() {

	hostInfo, err := host.Info()
	if err != nil {
		this.ErrorLog("获取服务器数据错误")
	}

	serverInfo := map[string]string{
		"localIp":        utils.Misc.GetLocalIp(),
		"hostname":       hostInfo.Hostname,
		"os":             hostInfo.OS,
		"platform":       hostInfo.Platform,
		"platformFamily": hostInfo.PlatformFamily,
	}

	errorLogCount, err := models.LogModel.CountLogsByLevel(models.Log_Level_Error)
	if err != nil {
		this.ErrorLog("获取数据错误")
	}

	errLogs, err := models.LogModel.GetLogsByKeywordAndLimit(fmt.Sprintf("%d", models.Log_Level_Error), "", "", 0, 5)
	if err != nil {
		this.ErrorLog("获取数据错误")
	}
	this.Data["serverInfo"] = serverInfo
	this.Data["errorLogCount"] = errorLogCount
	this.Data["errLogs"] = errLogs
	this.viewLayout("static/monitor", "static")
}

func (this *StaticController) ServerStatus() {
	memoryUsedPercent := 0
	cpuUsedPercent := 0
	diskUsedPercent := 0
	vm, err := mem.VirtualMemory()
	if err != nil {
		logs.Error("get memory err=%s", err.Error())
	}
	if vm != nil {
		memoryUsedPercent = int(vm.UsedPercent)
	}
	cpuPercent, _ := cpu.Percent(0, false)
	if len(cpuPercent) > 0 {
		cpuUsedPercent = int(cpuPercent[0])
	}
	d, err := disk.Usage("/")
	if err != nil {
		logs.Error("get disk err=%s", err.Error())
	}
	if d != nil {
		diskUsedPercent = int(d.UsedPercent)
	}

	data := map[string]interface{}{
		"memory_used_percent": memoryUsedPercent,
		"cpu_used_percent":    cpuUsedPercent,
		"disk_used_percent":   diskUsedPercent,
	}

	this.jsonSuccess("ok", data)
}

func (this *StaticController) ServerTime() {
	data := map[string]interface{}{
		"server_time": time.Now().Unix(),
		"run_time":    time.Now().Unix() - app.StartTime,
	}

	this.jsonSuccess("ok", data)
}
