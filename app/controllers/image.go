package controllers

import (
	"encoding/json"
	"path"
	"mm-wiki/app"
	"github.com/nu7hatch/gouuid"
	"strings"
)

type UploadResponse struct{
	Success int `json:"success"`
	Message string `json:"message"`
	Url string `json:"url"`
}

type ImageController struct {
	BaseController
}

func (this *ImageController) Upload() {

	f, h, err := this.GetFile("editormd-image-file")
	if err != nil {
		this.ErrorLog("上传图片错误: "+err.Error())
		this.jsonError("上传出错")
	}
	f.Close()

	ext := h.Filename[strings.LastIndex(h.Filename, "."):]

	uuId, _ := uuid.NewV4()
	uploadFile := path.Join(app.ImageAbsDir, uuId.String()+ext)
	err = this.SaveToFile("editormd-image-file", uploadFile)
	if err != nil {
		this.ErrorLog("上传图片错误: "+err.Error())
		this.jsonError("上传出错")
	}

	this.jsonSuccess("上传成功", "/images/"+uuId.String()+ext)
}

func (this *ImageController) jsonError(message string) {

	uploadRes := UploadResponse{
		Success: 0,
		Message: message,
		Url: "",
	}

	j, err := json.Marshal(uploadRes)
	if err != nil {
		this.Abort(err.Error())
	} else {
		this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
		this.Abort(string(j))
	}
}

func (this *ImageController) jsonSuccess(message string, url string) {

	uploadRes := UploadResponse{
		Success: 1,
		Message: message,
		Url: url,
	}

	j, err := json.Marshal(uploadRes)
	if err != nil {
		this.Abort(err.Error())
	} else {
		this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
		this.Abort(string(j))
	}
}