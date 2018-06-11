package controllers

type SpaceController struct {
	BaseController
}

func (this *SpaceController) Add() {


	this.viewLayout("space/form", "default")
}