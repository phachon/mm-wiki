package controllers


type AttachmentController struct {
	BaseController
}


func (this *AttachmentController) Page() {
	this.viewLayout("attachment/page", "attachment")
}

func (this *AttachmentController) Upload() {

}