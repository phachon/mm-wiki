package space

import (
	"github.com/phachon/mm-wiki/app/controller"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/config"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/gopkg/upload"
	"github.com/phachon/mm-wiki/logger"

	"github.com/gin-gonic/gin"
)

// DocSave 文档保存
func DocSave(ctx *gin.Context) error {

	parentId := controller.GetParamInt64Def(ctx, "parent_id", 0)
	name := controller.GetParamString(ctx, "name")
	docType := controller.GetParamInt64Def(ctx, "doc_type", 0)

	if name == "" {
		logger.WithContext(ctx).Warnf("[DocSave] 文档名不能为空")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档名不能为空")
	}
	if docType != entity.DocEntityType && docType != entity.DirEntityType {
		logger.WithContext(ctx).Warnf("[DocSave] 文档类型不合法")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档类型不合法")
	}

	serviceDoc := service.NewDoc(ctx)
	doc, err := serviceDoc.GetDocByDocId(parentId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocSave] GetDocById err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if doc == nil {
		logger.WithContext(ctx).Warnf("[DocSave] 文档不存在")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档不存在")
	}
	spaceKey := doc.SpaceKey

	serviceSpace := service.NewSpace(ctx)
	space, err := serviceSpace.GetSpaceByKey(spaceKey)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocSave] GetSpaceByKey err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if space == nil {
		logger.WithContext(ctx).Warnf("[DocSave] 空间不存在")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "空间不存在")
	}

	doc = &entity.DocEntity{
		ParentId: parentId,
		Name:     name,
		SpaceId:  space.SpaceId,
		SpaceKey: spaceKey,
		Type:     int(docType),
	}

	err = serviceDoc.CreateDoc(doc)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocSave] CreateDoc err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	logger.WithContext(ctx).Infof("[DocSave] 创建文档 %d 成功", doc.DocId)

	var data = map[string]interface{}{
		"doc_id": doc.DocId,
	}
	return controller.RespJsonSuccess(ctx, data)
}

// DocInfo 文档详情
func DocInfo(ctx *gin.Context) error {
	docId := controller.GetParamInt64(ctx, "doc_id")
	if docId <= 0 {
		logger.WithContext(ctx).Warnf("[DocInfo] 文档ID不能为空")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档ID不能为空")
	}

	serviceDoc := service.NewDoc(ctx)

	doc, err := serviceDoc.GetDocByDocId(docId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocInfo] GetDocById err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if doc == nil {
		logger.WithContext(ctx).Warnf("[DocInfo] 文档不存在")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档不存在")
	}

	// 获取文档内容
	serviceContent := service.NewContent(ctx)
	content, err := serviceContent.GetContentByDocId(docId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocInfo] GetContentByDocId err=%+v", err)
	}

	var data = map[string]interface{}{
		"doc_info": doc,
		"content":  content,
	}
	return controller.RespJsonSuccess(ctx, data)
}

// DocContentSave 文档内容保存
func DocContentSave(ctx *gin.Context) error {

	docId := controller.GetParamInt64(ctx, "doc_id")
	content := controller.GetParamString(ctx, "content")
	name := controller.GetParamString(ctx, "name")

	if docId <= 0 {
		logger.WithContext(ctx).Warnf("[DocContentSave] 文档ID不能为空")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档ID不能为空")
	}
	if name == "" {
		logger.WithContext(ctx).Warnf("[DocContentSave] 文档名不能为空")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档名不能为空")
	}

	serviceDoc := service.NewDoc(ctx)
	doc, err := serviceDoc.GetDocByDocId(docId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocContentSave] GetDocById err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "文档id不合法")
	}
	if doc == nil {
		logger.WithContext(ctx).Warnf("[DocContentSave] 文档不存在")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档id不合法")
	}
	var isUpdateDoc = false
	if doc.Name != name {
		doc.Name = name
		isUpdateDoc = true
	}

	// 更新文档
	serviceContent := service.NewContent(ctx)
	contentEntity, err := serviceContent.GetContentByDocId(docId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocContentSave] GetContentByDocId err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "获取文档内容失败")
	}

	if contentEntity == nil {
		err = serviceContent.Create(docId, content)
	} else {
		if contentEntity.Content != content {
			isUpdateDoc = true
			err = serviceContent.UpdateContent(docId, content, contentEntity.Content)
		}
	}
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocContentSave] UpdateContent err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "保存文档内容失败")
	}
	if !isUpdateDoc {
		return controller.RespJsonSuccess(ctx, map[string]interface{}{})
	}
	// 更新文档
	err = serviceDoc.UpdateDoc(doc)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocContentSave] UpdateDoc err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "更新文档失败")
	}
	return controller.RespJsonSuccess(ctx, map[string]interface{}{})
}

// DocUploadFile 文档上传文件
func DocUploadFile(ctx *gin.Context) error {
	docId := controller.GetParamInt64(ctx, "doc_id")
	if docId <= 0 {
		logger.WithContext(ctx).Warnf("[DocUploadFile] 文档ID不能为空")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档ID不能为空")
	}

	file, ferr := ctx.FormFile("file")
	if ferr != nil {
		logger.WithContext(ctx).Errorf("[DocUploadFile] FormFile err=%+v", ferr)
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文件不能为空")
	}

	uploadConf := config.GetAppConf().GetUploadConf("doc_file")
	uploaderHander := upload.GetUploaderHandler(upload.UplaoderName(uploadConf.UploadType))
	if uploaderHander == nil {
		logger.WithContext(ctx).Warnf("[DocUploadFile] 上传配置错误")
		return controller.RespJsonError(ctx, int32(errors.BusinessUnknownError), "上传配置错误")
	}
	logger.WithContext(ctx).Infof("[DocUploadFile] 上传文件：%+v", uploadConf)

	serviceDoc := service.NewDoc(ctx)
	doc, err := serviceDoc.GetDocByDocId(docId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocUploadFile] GetDocById err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "上传文件失败：获取文档错误")
	}
	if doc == nil {
		logger.WithContext(ctx).Warnf("[DocUploadFile] 文档不存在")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档id不合法")
	}

	// todo 生成文件唯一ID

	// 上传文件
	uploader := uploaderHander(&uploadConf.UploaderConfig)
	uploadResult, uploadErr := uploader.UploadFile(ctx, file)
	if uploadErr != nil {
		logger.WithContext(ctx).Errorf("[DocUploadFile] UploadFile err=%+v", uploadErr)
		return controller.RespJsonError(ctx, int32(errors.BusinessUnknownError), "上传文件失败")
	}

	return controller.RespJsonSuccess(ctx, map[string]interface{}{
		"path": uploadResult.Url,
		"url":  uploadResult.Url,
	})
}
