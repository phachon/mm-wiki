package controllers

import (
	"fmt"
	"mm-wiki/app/utils"
)

type PageController struct {
	BaseController
}

// page info
func (this *PageController) Index() {
	this.viewLayout("page/index", "page")
}

// page info
func (this *PageController) View() {

	pageId := this.GetString("page_id", "")


	fileInfo, err := utils.File.GetFileContents("test.md")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	this.Data["page_content"] = fileInfo
	this.Data["page_id"] = pageId
	this.viewLayout("page/view", "default_page")
}

// page edit
func (this *PageController) Edit() {

	pageId := this.GetString("page_id", "")

	fileInfo, err := utils.File.GetFileContents("test.md")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	this.Data["page_content"] = fileInfo
	this.Data["page_id"] = pageId
	this.viewLayout("page/edit", "default_page")
}

// page edit
func (this *PageController) Save() {

	pageId := this.GetString("page_id", "")

	this.Redirect("/page/view?page_id="+pageId, 302)
}