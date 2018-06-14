package controllers

type UserController struct {
	BaseController
}

func (this *UserController) Add() {

	this.viewLayout("user/form", "default")
}