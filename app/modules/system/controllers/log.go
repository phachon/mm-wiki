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
	level := strings.Trim(this.GetString("level", ""), "")
	message := strings.Trim(this.GetString("message", ""), "")
	username := strings.Trim(this.GetString("username", ""), "")

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
		//this.viewError(err.Error(), "template")
	}

	this.Data["logs"] = logs
	this.Data["username"] = username
	this.Data["level"] = level
	this.Data["message"] = message
	this.SetPaginator(number, count)
	this.viewLayout("log/system", "default")
}
