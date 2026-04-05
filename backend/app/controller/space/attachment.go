package space

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/controller"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/config"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/gopkg/upload"
	"github.com/phachon/mm-wiki/logger"
)

// AttachmentList 获取文档附件列表
func AttachmentList(ctx *gin.Context) error {

	docId := controller.GetParamString(ctx, "doc_id")
	source := controller.GetParamIntDef(ctx, "source", -1)

	if docId == "" {
		logger.WithContext(ctx).Warnf("[AttachmentList] doc_id empty")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档ID不能为空")
	}

	attachmentService := service.NewAttachment(ctx)
	var attachments []*entity.AttachmentEntity
	var err errors.BizError

	if source >= 0 {
		attachments, err = attachmentService.GetAttachmentsByDocIdAndSource(docId, source)
	} else {
		attachments, err = attachmentService.GetAttachmentsByDocId(docId)
	}

	if err != nil {
		logger.WithContext(ctx).Errorf("[AttachmentList] 获取附件列表失败: err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "获取附件列表失败")
	}

	data := map[string]interface{}{
		"list": attachments,
	}
	return controller.RespJsonSuccess(ctx, data)
}

// AttachmentUpload 上传附件
func AttachmentUpload(ctx *gin.Context) error {

	docId := controller.GetParamString(ctx, "doc_id")
	source := controller.GetParamIntDef(ctx, "source", entity.AttachmentSourceDefault)

	if docId == "" {
		logger.WithContext(ctx).Warnf("[AttachmentUpload] doc_id empty")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档ID不能为空")
	}

	// 获取上传文件
	file, ferr := ctx.FormFile("file")
	if ferr != nil {
		logger.WithContext(ctx).Errorf("[AttachmentUpload] FormFile err=%+v", ferr)
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "请选择上传文件")
	}

	// 获取上传配置
	uploadConf := config.GetAppConf().GetUploadConf("doc_file")
	uploaderHandler := upload.GetUploaderHandler(upload.UplaoderName(uploadConf.UploadType))
	if uploaderHandler == nil {
		logger.WithContext(ctx).Warnf("[AttachmentUpload] 上传配置错误")
		return controller.RespJsonError(ctx, int32(errors.BusinessUnknownError), "上传配置错误")
	}

	// 上传文件
	uploader := uploaderHandler(&uploadConf.UploaderConfig)
	uploadResult, uploadErr := uploader.UploadFile(ctx, file)
	if uploadErr != nil {
		logger.WithContext(ctx).Errorf("[AttachmentUpload] UploadFile err=%+v", uploadErr)
		return controller.RespJsonError(ctx, int32(errors.BusinessUnknownError), "上传文件失败")
	}

	// 保存附件信息到数据库
	accountId := global.ContextValueLoginAccountID(ctx)
	attachmentEntity := &entity.AttachmentEntity{
		AccountId: accountId,
		DocId:     docId,
		Name:      file.Filename,
		Path:      uploadResult.Url,
		Source:    source,
	}
	createErr := service.NewAttachment(ctx).Create(attachmentEntity)
	if createErr != nil {
		logger.WithContext(ctx).Errorf("[AttachmentUpload] Create attachment err=%+v", createErr)
		return controller.RespJsonError(ctx, createErr.GetErrCode(), "保存附件信息失败")
	}

	data := map[string]interface{}{
		"attachment_id": attachmentEntity.AttachmentId,
		"url":           uploadResult.Url,
		"name":          file.Filename,
	}
	return controller.RespJsonSuccess(ctx, data)
}

// AttachmentDelete 删除附件
func AttachmentDelete(ctx *gin.Context) error {

	attachmentId := controller.GetParamInt64(ctx, "attachment_id")

	if attachmentId <= 0 {
		logger.WithContext(ctx).Warnf("[AttachmentDelete] attachment_id empty")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "附件ID不能为空")
	}

	err := service.NewAttachment(ctx).DeleteAttachment(attachmentId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[AttachmentDelete] 删除附件 %d 失败: err=%+v", attachmentId, err)
		return controller.RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	return controller.RespJsonSuccess(ctx, nil)
}

// DocLogList 获取文档操作日志列表
func DocLogList(ctx *gin.Context) error {

	pageSize := controller.GetParamIntDef(ctx, "page_size", 20)
	pageNum := controller.GetParamIntDef(ctx, "page_num", 1)
	keywordsStr := controller.GetParamStringDef(ctx, "keywords", "")

	var keywords *entity.LogDocKeywords
	if len(keywordsStr) > 0 {
		jErr := json.Unmarshal([]byte(keywordsStr), &keywords)
		if jErr != nil {
			logger.WithContext(ctx).Errorf("[DocLogList] keywords unmarshal err=%s", jErr.Error())
		}
	}

	logDocService := service.NewLogDoc(ctx)

	// 获取文档日志列表
	logDocs, err := logDocService.GetLogDocsByLimit(pageSize, pageNum, keywords)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocLogList] 获取文档日志列表失败: err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "获取文档日志列表失败")
	}
	// 获取分页信息
	pageInfo, err := logDocService.GetPageInfoLimit(pageSize, pageNum, keywords)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocLogList] 获取分页信息失败: err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "获取分页信息失败")
	}

	data := map[string]interface{}{
		"list":      logDocs,
		"page_info": pageInfo,
	}
	return controller.RespJsonSuccess(ctx, data)
}
