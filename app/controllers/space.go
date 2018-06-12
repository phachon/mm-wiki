package controllers

type SpaceController struct {
	BaseController
}

func (this *SpaceController) Index() {
	this.viewLayout("main/index", "space")
}

func (this *SpaceController) Add()  {
	this.viewLayout("space/form", "default")
}