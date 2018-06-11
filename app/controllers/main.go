package controllers

import "mm-wiki/app"

type MainController struct {
	BaseController
}

func (this *MainController) Index() {
	this.Data["version"] = app.Version
	this.viewLayout("main/index", "main")
}