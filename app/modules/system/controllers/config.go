package controllers

type ConfigController struct {
	BaseController
}

func (this *ConfigController) Link() {
	this.viewLayout("config/link", "default")
}

func (this *ConfigController) Email() {
	this.viewLayout("config/email", "default")
}