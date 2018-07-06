package controllers

type MainController struct {
	BaseController
}

func (this *MainController) Index() {
	this.viewLayout("main/index", "main")
}
func (this *MainController) Default() {
	this.viewLayout("main/default", "default")
}

func (this *MainController) About() {
	this.viewLayout("main/about", "default")
}