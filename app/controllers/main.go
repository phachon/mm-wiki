package controllers

import (
	"fmt"
	"mm-wiki/app/utils"
)

type MainController struct {
	BaseController
}

func (this *MainController) Index() {
	this.viewLayout("main/index", "main")
}

// main collection page
func (this *MainController) Page() {

	pageId := this.GetString("page_id", "")


	fileInfo, err := utils.File.GetFileContents("test.md")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	this.Data["page_content"] = fileInfo
	this.Data["page_id"] = pageId
	this.viewLayout("page/info", "page")
}