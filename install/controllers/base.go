package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/phachon/mm-wiki/install/storage"
	"strings"
)

type BaseController struct {
	beego.Controller
}

type JsonResponse struct {
	Code     int                    `json:"code"`
	Message  interface{}            `json:"message"`
	Data     interface{}            `json:"data"`
	Redirect map[string]interface{} `json:"redirect"`
}

// prepare
func (this *BaseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	methodName := strings.ToLower(actionName)
	if (methodName == "index") || (methodName == "end") || (methodName == "status") {
		return
	}
	if storage.Data.Status == storage.Install_Start {
		if methodName == "ready" && this.isPost() {
			return
		}
		this.Redirect("/install/end", 302)
		this.StopRun()
	}
	if storage.Data.Status == storage.Install_End {
		if storage.Data.IsSuccess == storage.Install_Failed {
			// 重置
			storage.Data.IsSuccess = storage.Install_Default
			storage.Data.Status = storage.Install_Ready
			storage.Data.Result = ""
		} else {
			this.StopRun()
		}
	}
}

// view layout title
func (this *BaseController) viewLayoutTitle(title, viewName, layout string) {
	this.Layout = "install/layout.html"
	this.TplName = viewName + ".html"
	this.Data["title"] = title
	this.Data["copyright"] = storage.CopyRight
	this.Render()
}

// view layout
func (this *BaseController) viewLayout(viewName, layout string) {
	this.Layout = "install/layout.html"
	this.TplName = viewName + ".html"
	this.Data["title"] = ""
	this.Data["copyright"] = storage.CopyRight
	this.Render()
}

// view
func (this *BaseController) view(viewName string) {
	this.Layout = "install/layout.html"
	this.TplName = viewName + ".html"
	this.Data["title"] = ""
	this.Data["copyright"] = storage.CopyRight
	this.Render()
}

// error view
func (this *BaseController) viewError(errorMessage string, data ...interface{}) {
	this.Layout = "layout/install.html"
	redirect := "/"
	sleep := 2000
	if len(data) > 0 {
		redirect = data[0].(string)
	}
	if len(data) > 1 {
		sleep = data[1].(int)
	}
	_, actionName := this.GetControllerAndAction()
	methodName := strings.ToLower(actionName)
	this.TplName = "install/error.html"
	this.Data["title"] = "error"
	this.Data["method"] = methodName
	this.Data["message"] = errorMessage
	this.Data["redirect"] = redirect
	this.Data["sleep"] = sleep
	this.Data["copyright"] = storage.CopyRight
	this.Render()
}

// view title
func (this *BaseController) viewTitle(title, viewName string) {
	this.Layout = "install/layout.html"
	this.TplName = viewName + ".html"
	this.Data["title"] = title
	this.Data["copyright"] = storage.CopyRight
	this.Render()
}

// return json success
func (this *BaseController) jsonSuccess(message interface{}, data ...interface{}) {
	url := ""
	sleep := 300
	var _data interface{}
	if len(data) > 0 {
		_data = data[0]
	}
	if len(data) > 1 {
		url = data[1].(string)
	}
	if len(data) > 2 {
		sleep = data[2].(int)
	}
	this.Data["json"] = JsonResponse{
		Code:    1,
		Message: message,
		Data:    _data,
		Redirect: map[string]interface{}{
			"url":   url,
			"sleep": sleep,
		},
	}

	j, err := json.MarshalIndent(this.Data["json"], "", "\t")
	if err != nil {
		this.Abort(err.Error())
	} else {
		this.Abort(string(j))
	}
}

// return json error
func (this *BaseController) jsonError(message interface{}, data ...interface{}) {
	url := ""
	sleep := 2000
	var _data interface{}
	if len(data) > 0 {
		_data = data[0]
	}
	if len(data) > 1 {
		url = data[1].(string)
	}
	if len(data) > 2 {
		sleep = data[2].(int)
	}
	this.Data["json"] = JsonResponse{
		Code:    0,
		Message: message,
		Data:    _data,
		Redirect: map[string]interface{}{
			"url":   url,
			"sleep": sleep,
		},
	}
	j, err := json.MarshalIndent(this.Data["json"], "", " \t")
	if err != nil {
		this.Abort(err.Error())
	} else {
		this.Abort(string(j))
	}
}

// get client ip
func (this *BaseController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

// is post
func (this *BaseController) isPost() bool {
	return this.Ctx.Input.IsPost()
}
