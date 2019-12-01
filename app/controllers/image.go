package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/phachon/mm-wiki/app"
	"github.com/phachon/mm-wiki/app/models"
	"github.com/phachon/mm-wiki/app/utils"
	"os"
	"path"
)

type UploadResponse struct {
	Success int    `json:"success"`
	Message string `json:"message"`
	Url     string `json:"url"`
}

type ImageController struct {
	BaseController
}

func (this *ImageController) Upload() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/space/index")
	}
	documentId := this.GetString("document_id", "")
	if documentId == "" {
		this.jsonError("参数错误！")
	}

	// handle document
	document, err := models.DocumentModel.GetDocumentByDocumentId(documentId)
	if err != nil {
		this.ErrorLog("查找空间文档 " + documentId + " 失败：" + err.Error())
		this.jsonError("查找文档失败！")
	}
	if len(document) == 0 {
		this.jsonError("文档不存在！")
	}

	spaceId := document["space_id"]
	space, err := models.SpaceModel.GetSpaceBySpaceId(spaceId)
	if err != nil {
		this.ErrorLog("查找文档 " + documentId + " 所在空间失败：" + err.Error())
		this.jsonError("查找文档失败！")
	}
	if len(space) == 0 {
		this.jsonError("文档所在空间不存在！")
	}
	// check space visit_level
	_, isEditor, _ := this.GetDocumentPrivilege(space)
	if !isEditor {
		this.jsonError("您没有权限操作该空间下的文档！")
	}

	// handle upload
	f, h, err := this.GetFile("editormd-image-file")
	if err != nil {
		this.ErrorLog("上传图片数据错误: " + err.Error())
		this.jsonError("上传图片数据错误")
		return
	}
	if h == nil || f == nil {
		this.ErrorLog("上传图片错误")
		this.jsonError("上传图片错误")
		return
	}
	_ = f.Close()

	// file save dir
	saveDir := fmt.Sprintf("%s/%s/%s", app.ImageAbsDir, spaceId, documentId)
	ok, _ := utils.File.PathIsExists(saveDir)
	if !ok {
		err := os.MkdirAll(saveDir, 0777)
		if err != nil {
			this.ErrorLog("上传图片错误: " + err.Error())
			this.jsonError("上传图片失败")
			return
		}
	}
	// check file is exists
	imageFile := path.Join(saveDir, h.Filename)
	ok, _ = utils.File.PathIsExists(imageFile)
	if ok {
		this.jsonError("该图片已经上传过！")
	}
	// save file
	err = this.SaveToFile("editormd-image-file", imageFile)
	if err != nil {
		this.ErrorLog("图片保存失败: " + err.Error())
		this.jsonError("图片保存失败")
	}

	// insert db
	attachment := map[string]interface{}{
		"user_id":     this.UserId,
		"document_id": documentId,
		"name":        h.Filename,
		"path":        fmt.Sprintf("images/%s/%s/%s", spaceId, documentId, h.Filename),
		"source":      models.Attachment_Source_Image,
	}
	_, err = models.AttachmentModel.Insert(attachment)
	if err != nil {
		_ = os.Remove(imageFile)
		this.ErrorLog("上传图片保存信息错误: " + err.Error())
		this.jsonError("图片信息保存失败")
	}

	this.InfoLog(fmt.Sprintf("文档 %s 上传图片 %s 成功", documentId, h.Filename))
	this.jsonSuccess("上传成功", fmt.Sprintf("/%s", attachment["path"]))
}

func (this *ImageController) jsonError(message string) {

	uploadRes := UploadResponse{
		Success: 0,
		Message: message,
		Url:     "",
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
		Url:     url,
	}

	j, err := json.Marshal(uploadRes)
	if err != nil {
		this.Abort(err.Error())
	} else {
		this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
		this.Abort(string(j))
	}
}
