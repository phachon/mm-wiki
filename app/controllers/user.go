package controllers

type UserController struct {
	BaseController
}

func (this *UserController) Index() {
	this.viewLayout("user/index", "user")
}

func (this *UserController) List() {
	this.viewLayout("user/list", "default")
}