package controllers

import (
	"mm-wiki/app/models"
	"strings"
)

type LogController struct {
	BaseController
}

func (this *LogController) System() {

	page, _ := this.GetInt("page", 1)
	level := strings.TrimSpace(this.GetString("level", ""))
	message := strings.TrimSpace(this.GetString("message", ""))
	username := strings.TrimSpace(this.GetString("username", ""))

	number := 15
	limit := (page - 1) * number
	var err error
	var count int64
	var logs []map[string]string
	if level != "" || message != "" || username != "" {
		count, err = models.LogModel.CountLogsByKeyword(level, message, username)
		logs, err = models.LogModel.GetLogsByKeywordAndLimit(level, message, username, limit, number)
	} else {
		count, err = models.LogModel.CountLogs()
		logs, err = models.LogModel.GetLogsByLimit(limit, number)
	}
	if err != nil {
		this.ErrorLog("查找系统日志出错："+err.Error())
		this.ViewError("查找系统日志出错", "/system/main/index")
	}

	this.Data["logs"] = logs
	this.Data["username"] = username
	this.Data["level"] = level
	this.Data["message"] = message
	this.SetPaginator(number, count)
	this.viewLayout("log/system", "default")
}

func (this *LogController) Info() {

	logId := this.GetString("log_id", "")
	if logId == "" {
		this.ViewError("日志不存在", "/system/log/system")
	}

	log, err := models.LogModel.GetLogByLogId(logId)
	if err != nil {
		this.ViewError("日志不存在", "/system/log/system")
	}

	this.Data["log"] = log
	this.viewLayout("log/info", "default")
}