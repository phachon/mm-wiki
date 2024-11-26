package space

import (
	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/controller"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
)

// DocSave 文档保存
func DocSave(ctx *gin.Context) error {

	parentId := controller.GetParamInt64Def(ctx, "parent_id", 0)
	name := controller.GetParamString(ctx, "name")
	spaceKey := controller.GetParamString(ctx, "space_key")
	docType := controller.GetParamInt64Def(ctx, "doc_type", 0)

	if name == "" {
		logger.WithContext(ctx).Warnf("[DocSave] 文档名不能为空")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档名不能为空")
	}
	if spaceKey == "" {
		logger.WithContext(ctx).Warnf("[DocSave] 空间Key不能为空")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "空间Key不能为空")
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
		return controller.RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	var data = map[string]interface{}{
		"doc_info": doc,
		"content":  content,
	}
	return controller.RespJsonSuccess(ctx, data)
}
