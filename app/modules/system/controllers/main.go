package controllers

type MainController struct {
	BaseController
}

func (this *MainController) Index() {
	this.viewLayout("main/index", "main")
}