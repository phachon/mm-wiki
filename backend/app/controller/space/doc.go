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

// DocCreate 文档创建
func DocCreate(ctx *gin.Context) error {

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
			err = serviceContent.UpdateContent(docId, content, contentEntity.Content, doc)
		}
	}
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocContentSave] UpdateContent err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "保存文档内容失败")
	}
	// 更新文档
	err = serviceDoc.UpdateNameAndEditAccount(docId, name)
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

// DocHistoryList 文档历史列表
func DocHistoryList(ctx *gin.Context) error {
	docId := controller.GetParamInt64(ctx, "doc_id")
	pageSize := controller.GetParamIntDef(ctx, "page_size", 10)
	pageNum := controller.GetParamIntDef(ctx, "page_num", 1)

	if docId <= 0 {
		logger.WithContext(ctx).Warnf("[DocHistoryList] 文档ID不能为空")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档ID不能为空")
	}

	serviceDoc := service.NewDoc(ctx)
	docInfo, err := serviceDoc.GetDocByDocId(docId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocHistoryList] GetContentsByDocId err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "获取文档信息失败")
	}
	if docInfo == nil {
		logger.WithContext(ctx).Warnf("[DocHistoryList] 文档不存在")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档不存在")
	}

	// 获取文档历史
	docVersion := service.NewContentVersion(ctx)
	contentVersions, err := docVersion.GetContentVersionsByDocIdLimit(
		docId, pageSize, pageNum,
	)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocHistoryList] GetContentVersionsByDocId err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "获取文档历史失败")
	}

	// 获取分页信息
	pageInfo, err := docVersion.GetPageInfoLimit(docId, pageSize, pageNum)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocHistoryList] GetPageInfoLimit err=%s", err.Error())
		return controller.RespJsonError(ctx, err.GetErrCode(), "获取文档历史失败")
	}

	var data = map[string]interface{}{
		"version_list": contentVersions,
		"page_info":    pageInfo,
	}
	return controller.RespJsonSuccess(ctx, data)
}

// DocContentVersion 文档内容版本
func DocContentVersion(ctx *gin.Context) error {
	docId := controller.GetParamInt64(ctx, "doc_id")
	contentVersionId := controller.GetParamInt64(ctx, "content_version_id")

	if docId <= 0 {
		logger.WithContext(ctx).Warnf("[DocContentVersion] 文档ID不能为空")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档ID不能为空")
	}
	if contentVersionId <= 0 {
		logger.WithContext(ctx).Warnf("[DocContentVersion] 版本ID不能为空")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "版本id不能为空")
	}

	// 获取文档内容版本
	docVersion := service.NewContentVersion(ctx)
	contentVersion, err := docVersion.GetContentVersionByVersionId(contentVersionId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocContentVersion] GetContentVersionByVersionId err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "获取文档版本内容失败")
	}
	if contentVersion == nil {
		logger.WithContext(ctx).Warnf("[DocContentVersion] 文档版本不存在")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档版本不存在")
	}
	if contentVersion.DocId != docId {
		logger.WithContext(ctx).Warnf("[DocContentVersion] 文档版本不匹配")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档版本不存在")
	}

	var data = map[string]interface{}{
		"content_version": contentVersion,
	}
	return controller.RespJsonSuccess(ctx, data)
}

// DocRecover 文档恢复操作
func DocRecover(ctx *gin.Context) error {
	docId := controller.GetParamInt64(ctx, "doc_id")
	contentVersionId := controller.GetParamInt64(ctx, "content_version_id")

	if docId <= 0 {
		logger.WithContext(ctx).Warnf("[DocRecover] 文档ID不能为空")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档ID不能为空")
	}
	if contentVersionId <= 0 {
		logger.WithContext(ctx).Warnf("[DocRecover] 版本ID不能为空")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "版本ID不能为空")
	}

	serviceDoc := service.NewDoc(ctx)
	doc, err := serviceDoc.GetDocByDocId(docId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocRecover] GetDocById err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "获取文档信息失败")
	}
	if doc == nil {
		logger.WithContext(ctx).Warnf("[DocRecover] 文档不存在")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档不存在")
	}

	serviceContentVersion := service.NewContentVersion(ctx)
	contentVersion, err := serviceContentVersion.GetContentVersionByVersionId(contentVersionId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocRecover] GetContentVersionByVersionId err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "获取文档版本信息失败")
	}
	if contentVersion == nil {
		logger.WithContext(ctx).Warnf("[DocRecover] 文档版本不存在")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档版本不存在")
	}
	if contentVersion.DocId != docId {
		logger.WithContext(ctx).Warnf("[DocRecover] 文档版本不匹配")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档版本不存在")
	}

	serviceContent := service.NewContent(ctx)
	content, err := serviceContent.GetContentByDocId(docId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocRecover] GetContentByDocId err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "获取文档内容失败")
	}
	if content == nil {
		logger.WithContext(ctx).Warnf("[DocRecover] 文档内容不存在")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档内容不存在")
	}

	// 更新版本正文
	err = serviceContent.UpdateContent(docId, contentVersion.Content, content.Content, doc)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocRecover] UpdateContent err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "恢复文档版本失败")
	}

	// 更新文档信息
	err = serviceDoc.UpdateNameAndEditAccount(docId, doc.Name)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocContentSave] UpdateDoc err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "更新文档失败")
	}

	return controller.RespJsonSuccess(ctx, map[string]interface{}{})
}

// DocContentVersionDel 文档内容版本删除
func DocContentVersionDel(ctx *gin.Context) error {
	docId := controller.GetParamInt64(ctx, "doc_id")
	contentVersionId := controller.GetParamInt64(ctx, "content_version_id")

	if docId <= 0 {
		logger.WithContext(ctx).Warnf("[DocContentVersionDel] 文档ID不能为空")
	}
	if contentVersionId <= 0 {
		logger.WithContext(ctx).Warnf("[DocContentVersionDel] 版本ID不能为空")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "版本ID不能为空")
	}

	serviceDoc := service.NewDoc(ctx)
	doc, err := serviceDoc.GetDocByDocId(docId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocContentVersionDel] GetDocById err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "获取文档信息失败")
	}
	if doc == nil {
		logger.WithContext(ctx).Warnf("[DocContentVersionDel] 文档不存在")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档不存在")
	}

	serviceContentVersion := service.NewContentVersion(ctx)
	contentVersion, err := serviceContentVersion.GetContentVersionByVersionId(contentVersionId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocContentVersionDel] GetContentVersionByVersionId err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "获取文档版本信息失败")
	}
	if contentVersion == nil {
		logger.WithContext(ctx).Warnf("[DocContentVersionDel] 文档版本不存在")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档版本不存在")
	}
	if contentVersion.DocId != docId {
		logger.WithContext(ctx).Warnf("[DocContentVersionDel] 文档版本不匹配")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档版本不存在")
	}

	err = serviceContentVersion.DeleteContentVersion(contentVersionId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocContentVersionDel] DeleteContentVersion err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "删除文档版本失败")
	}

	return controller.RespJsonSuccess(ctx, map[string]interface{}{})
}
