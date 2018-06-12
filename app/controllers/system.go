package controllers

type SystemController struct {
	BaseController
}

func (this *SystemController) Index() {

	this.viewLayout("system/index", "system")
}