package controllers

type UserController struct {
	BaseController
}

func (this *UserController) Index() {

	this.viewLayout("user/index", "user")
}