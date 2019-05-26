package controllers

import "encoding/json"

type BaseController struct {
	TemplateController
}

// view layout
func (this *BaseController) viewLayout(viewName, layout string) {
	layout = "layouts/" + layout
	this.ViewLayout(viewName, layout)
}

// return json success
func (this *BaseController) jsonSuccess(message interface{}, data ...interface{}) {
	this.JsonSuccess(message, data...)
}

// return json error
func (this *BaseController) jsonError(message interface{}, data ...interface{}) {
	this.JsonError(message, data...)
}

type UploadJsonResponse struct {
	Code     int                    `json:"code"`
	Message  interface{}            `json:"message"`
	Data     interface{}            `json:"data"`
	Redirect map[string]interface{} `json:"redirect"`
	Error    interface{}            `json:"error"`
}

func (this *BaseController) uploadJsonError(message interface{}, data ...interface{}) {
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
	this.Data["json"] = UploadJsonResponse{
		Code:    0,
		Message: message,
		Data:    _data,
		Redirect: map[string]interface{}{
			"url":   url,
			"sleep": sleep,
		},
		Error: message,
	}
	j, err := json.MarshalIndent(this.Data["json"], "", " \t")
	if err != nil {
		this.Abort(err.Error())
	} else {
		this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
		this.Abort(string(j))
	}
}
