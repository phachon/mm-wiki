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

func (this *LogController) Document() {

	page, _ := this.GetInt("page", 1)
	number := 12
	limit := (page - 1) * number
	keyword := strings.TrimSpace(this.GetString("keyword", ""))
	userId := strings.TrimSpace(this.GetString("user_id", ""))

	var logDocuments = []map[string]string{}
	var err error
	var count int64
	if keyword != "" && userId != "" {
		logDocuments, err = models.LogDocumentModel.GetLogDocumentsByUserIdKeywordAndLimit(userId, keyword, limit, number)
		count, err = models.LogDocumentModel.CountLogDocumentsByUserIdAndKeyword(userId, keyword)
	}else if userId != "" {
		logDocuments, err = models.LogDocumentModel.GetLogDocumentsByUserIdAndLimit(userId, limit, number)
		count, err = models.LogDocumentModel.CountLogDocumentsByUserId(userId)
	}else if keyword != "" {
		logDocuments, err = models.LogDocumentModel.GetLogDocumentsByKeywordAndLimit(keyword, limit, number)
		count, err = models.LogDocumentModel.CountLogDocumentsByKeyword(userId)
	}else {
		logDocuments, err = models.LogDocumentModel.GetLogDocumentsByLimit(limit, number)
		count, err = models.LogDocumentModel.CountLogDocuments()
	}
	if err != nil {
		this.ErrorLog("文档日志查找失败："+err.Error())
		this.ViewError("文档日志查找失败！", "/system/main/index")
	}

	userIds := []string{}
	docIds := []string{}
	for _, logDocument := range logDocuments {
		userIds = append(userIds, logDocument["user_id"])
		docIds = append(docIds, logDocument["document_id"])
	}
	users, err := models.UserModel.GetUsersByUserIds(userIds)
	if err != nil {
		this.ErrorLog("文档日志查找失败："+err.Error())
		this.ViewError("文档日志查找失败！", "/system/main/index")
	}
	docs, err := models.DocumentModel.GetDocumentsByDocumentIds(docIds)
	if err != nil {
		this.ErrorLog("文档日志查找失败："+err.Error())
		this.ViewError("文档日志查找失败！", "/system/main/index")
	}
	for _, logDocument := range logDocuments {
		logDocument["username"] = ""
		for _, user := range users {
			if logDocument["user_id"] == user["user_id"] {
				logDocument["username"] = user["username"]
				logDocument["given_name"] = user["given_name"]
				break
			}
		}
		for _, doc := range docs {
			if logDocument["document_id"] == doc["document_id"] {
				logDocument["document_name"] = doc["name"]
				logDocument["document_type"] = doc["type"]
				break
			}
		}
	}

	users, err = models.UserModel.GetUsers()
	if err != nil {
		this.ErrorLog("文档日志查找失败："+err.Error())
		this.ViewError("文档日志查找失败！", "/system/main/index")
	}
	this.Data["logDocuments"] = logDocuments
	this.Data["keyword"] = keyword
	this.Data["userId"] = userId
	this.Data["users"] = users
	this.SetPaginator(number, count)
	this.viewLayout("log/document", "default")
}