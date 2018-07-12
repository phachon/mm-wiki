package controllers

import (
	"mm-wiki/app/models"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	mem "github.com/shirou/gopsutil/mem"
)

type StaticController struct {
	BaseController
}

func (this *StaticController) Default() {

	normalUserCount, err := models.UserModel.CountNormalUsers()
	if err != nil {
		this.ErrorLog("查找用户数出错："+err.Error())
		this.jsonError("查找数据出错")
	}

	forbiddenUserCount, err := models.UserModel.CountForbiddenUsers()
	if err != nil {
		this.ErrorLog("查找用户数出错："+err.Error())
		this.jsonError("查找数据出错")
	}

	spaceCount, err := models.SpaceModel.CountSpaces()
	if err != nil {
		this.ErrorLog("查找空间总数出错："+err.Error())
		this.jsonError("查找数据出错")
	}

	documentCount, err := models.DocumentModel.CountDocuments()
	if err != nil {
		this.ErrorLog("查找文档总数出错："+err.Error())
		this.jsonError("查找数据出错")
	}
	this.Data["normalUserCount"] = normalUserCount
	this.Data["forbiddenUserCount"] = forbiddenUserCount
	this.Data["spaceCount"] = spaceCount
	this.Data["documentCount"] = documentCount
	this.viewLayout("static/default", "static")
}

func (this *StaticController) SpaceDocsRank() {

	number := 15

	spaceDocsRank := []map[string]string{}
	spaceDocCountIds , err := models.DocumentModel.GetSpaceIdsOrderByCountDocumentLimit(number)
	if err != nil {
		this.ErrorLog("查找空间文档排行出错："+err.Error())
		this.jsonError("查找数据出错")
	}
	if len(spaceDocCountIds) > 0 {
		spaceIds := []string{}
		for _, spaceDocCountId := range spaceDocCountIds {
			spaceIds = append(spaceIds, spaceDocCountId["space_id"])
		}
		spaces, err := models.SpaceModel.GetSpaceBySpaceIds(spaceIds)
		if err != nil {
			this.ErrorLog("查找空间文档排行出错："+err.Error())
			this.jsonError("查找数据出错")
		}
		for _, spaceDocCountId := range spaceDocCountIds {
			spaceDocCount := map[string]string{
				"space_name": "",
				"total": spaceDocCountId["total"],
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

func (this *StaticController) DocCountByTime() {

	documentCountDate , err := models.DocumentModel.GetCountGroupByCreateTime()
	if err != nil {
		this.ErrorLog("查找空间文档排行出错："+err.Error())
		this.jsonError("查找数据出错")
	}

	this.jsonSuccess("ok", documentCountDate)
}

func (this *StaticController) Monitor() {

	this.viewLayout("static/monitor", "static")
}

func (this *StaticController) ServerStatus() {
	vm, _ := mem.VirtualMemory()
	cpuPercent, _ := cpu.Percent(0, false)
	d, _ := disk.Usage("/")

	data := map[string]interface{}{
		"memory_used_percent": int(vm.UsedPercent),
		"cpu_used_percent":    int(cpuPercent[0]),
		"disk_used_percent":   int(d.UsedPercent),
	}

	this.jsonSuccess("ok", data)
}