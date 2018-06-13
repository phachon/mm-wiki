package controllers

type SpaceController struct {
	BaseController
}

func (this *SpaceController) Index() {
	this.viewLayout("space/index", "space")
}