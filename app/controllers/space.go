package controllers

type SpaceController struct {
	BaseController
}

func (this *SpaceController) Index() {
	this.viewLayout("space/index", "space")
}

func (this *SpaceController) List() {
	this.viewLayout("space/list", "default")
}