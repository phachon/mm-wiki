package controllers

type PrivilegeController struct {
	BaseController
}

func (this *PrivilegeController) Add() {
	this.viewLayout("privilege/form", "default")
}