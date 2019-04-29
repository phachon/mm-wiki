package controllers

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
