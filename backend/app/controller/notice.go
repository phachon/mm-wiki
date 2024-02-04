package controller

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
	"github.com/phachon/mm-wiki/utils"
)

// NoticeSave 添加公告保存
func NoticeSave(ctx *gin.Context) error {

	title := GetParamString(ctx, "title")
	content := GetParamString(ctx, "content")
	publishStatus := GetParamInt(ctx, "publish_status") // 发布状态
	startTimeStr := GetParamString(ctx, "start_time")   // 开始时间
	endTimeStr := GetParamString(ctx, "end_time")       // 结束时间
	// 判断参数合法性
	if title == "" {
		logger.WithContext(ctx).Warnf("[NoticeSave] title empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "公告标题不能为空")
	}
	if content == "" {
		logger.WithContext(ctx).Warnf("[NoticeSave] name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "公告名不能为空")
	}
	if publishStatus != entity.NoticePublishStatusNotPublish &&
		publishStatus != entity.NoticePublishStatusPublishing {
		logger.WithContext(ctx).Warnf("[NoticeSave] is_publish empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "发布状态错误")
	}
	startTime, cErr := time.ParseInLocation("2006-01-02 15:04", startTimeStr, time.Local)
	if cErr != nil {
		logger.WithContext(ctx).Warnf("[NoticeSave] start_time err:%+v", cErr)
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "开始时间错误")
	}
	endTime, cErr := time.ParseInLocation("2006-01-02 15:04", endTimeStr, time.Local)
	if cErr != nil {
		logger.WithContext(ctx).Warnf("[NoticeSave] end_time err:%+v", cErr)
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "结束时间错误")
	}
	// 时间合法性判断
	if startTime.Unix() >= endTime.Unix() {
		logger.WithContext(ctx).Warnf("[NoticeSave] start_time err:%+v", cErr)
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "开始时间必须小于结束时间")
	}
	if endTime.Unix() <= time.Now().Unix() {
		logger.WithContext(ctx).Warnf("[NoticeSave] end_time err:%+v", cErr)
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "结束时间必须大于当前时间")
	}

	// notice 公告实体
	noticeEntity := &entity.NoticeEntity{
		Title:         title,
		Content:       content,
		AccountId:     global.ContextValueLoginAccountID(ctx),
		AccountName:   global.ContextValueLoginAccountName(ctx),
		Status:        entity.NoticeStatusDefault,
		PublishStatus: publishStatus,
		StartTime:     utils.NewJsonTime(startTime),
		EndTime:       utils.NewJsonTime(endTime),
	}
	// 创建公告
	err := service.NewNotice(ctx).Create(noticeEntity)
	if err != nil {
		sysLogErrorf(ctx, "[NoticeAdd] 添加公告失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[NoticeAdd] 添加公告 %d 成功", noticeEntity.NoticeId)

	return RespJsonSuccess(ctx, nil)
}

// NoticeEdit 公告修改页面
func NoticeEdit(ctx *gin.Context) error {

	noticeId := GetParamInt64(ctx, "notice_id")
	// 判断参数合法性
	if noticeId == 0 {
		logger.WithContext(ctx).Warnf("[NoticeEdit] 公告id不存在")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "公告id不存在")
	}

	notice, err := service.NewNotice(ctx).GetNoticeByNoticeId(noticeId)
	if err != nil {
		sysLogErrorf(ctx, "[NoticeEdit] 获取公告 %d 信息失败: err=%+v", noticeId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if notice == nil {
		logger.WithContext(ctx).Warnf("[NoticeEdit] 公告 %d 不存在", noticeId)
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "公告id错误")
	}

	return RespJsonSuccess(ctx, map[string]interface{}{
		"notice_info": notice,
	})
}

// NoticeModify 修改公告保存
func NoticeModify(ctx *gin.Context) error {

	noticeId := GetParamInt64(ctx, "notice_id")
	title := GetParamString(ctx, "title")
	content := GetParamString(ctx, "content")
	publishStatus := GetParamInt(ctx, "publish_status") // 发布状态
	startTimeStr := GetParamString(ctx, "start_time")   // 开始时间
	endTimeStr := GetParamString(ctx, "end_time")       // 结束时间
	// 判断参数合法性
	if noticeId <= 0 {
		logger.WithContext(ctx).Warnf("[NoticeModify] notice_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "公告ID不合法")
	}
	if content == "" {
		logger.WithContext(ctx).Warnf("[NoticeModify] name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "公告名不能为空")
	}
	if title == "" {
		logger.WithContext(ctx).Warnf("[NoticeModify] title empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "公告标题不能为空")
	}
	if publishStatus != entity.NoticePublishStatusNotPublish &&
		publishStatus != entity.NoticePublishStatusPublishing {
		logger.WithContext(ctx).Warnf("[NoticeModify] is_publish empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "发布状态错误")
	}
	startTime, cErr := time.ParseInLocation("2006-01-02 15:04", startTimeStr, time.Local)
	if cErr != nil {
		logger.WithContext(ctx).Warnf("[NoticeModify] start_time err:%+v", cErr)
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "开始时间错误")
	}
	endTime, cErr := time.ParseInLocation("2006-01-02 15:04", endTimeStr, time.Local)
	if cErr != nil {
		logger.WithContext(ctx).Warnf("[NoticeModify] end_time err:%+v", cErr)
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "结束时间错误")
	}
	// 时间合法性判断
	if startTime.Unix() >= endTime.Unix() {
		logger.WithContext(ctx).Warnf("[NoticeModify] start_time err:%+v", cErr)
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "开始时间必须小于结束时间")
	}
	if endTime.Unix() <= time.Now().Unix() {
		logger.WithContext(ctx).Warnf("[NoticeModify] end_time err:%+v", cErr)
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "结束时间必须大于当前时间")
	}

	// notice 公告实体
	noticeEntity := entity.NoticeEntity{
		NoticeId:      noticeId,
		Content:       content,
		Title:         title,
		AccountId:     global.ContextValueLoginAccountID(ctx),
		PublishStatus: publishStatus,
		StartTime:     utils.NewJsonTime(startTime),
		EndTime:       utils.NewJsonTime(endTime),
	}

	// 更新公告
	err := service.NewNotice(ctx).Update(noticeEntity)
	if err != nil {
		sysLogErrorf(ctx, "[NoticeModify] 更新公告 %d 失败: err=%+v", noticeId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[NoticeModify] 更新公告 %d 成功", noticeId)

	return RespJsonSuccess(ctx, nil)
}

// NoticeList 公告列表
func NoticeList(ctx *gin.Context) error {

	pageSize := GetParamIntDef(ctx, "page_size", 20)
	pageNum := GetParamIntDef(ctx, "page_num", 1)
	keywordsStr := GetParamStringDef(ctx, "keywords", "")

	var keywords *entity.NoticeKeywords
	if len(keywordsStr) > 0 {
		jErr := json.Unmarshal([]byte(keywordsStr), &keywords)
		if jErr != nil {
			logger.WithContext(ctx).Errorf("[NoticeList] GetNoticesByLimit err=%s", jErr.Error())
		}
	}

	noticeService := service.NewNotice(ctx)

	// 获取公告列表
	notices, err := noticeService.GetNoticesByLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[NoticeList] 获取公告列表失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取分页信息
	pageInfo, err := noticeService.GetPageInfoLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[NoticeList] 获取公告分页信息失败: err=%s", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 格式化列表结构
	noticeList, err := noticeService.FormatNoticeList(notices)
	if err != nil {
		sysLogErrorf(ctx, "[RoleList] 格式化角色列表失败: err=%s", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	data := map[string]interface{}{
		"list":      noticeList,
		"page_info": pageInfo,
	}
	return RespJsonSuccess(ctx, data)
}

// NoticeDelete 公告删除
func NoticeDelete(ctx *gin.Context) error {

	noticeId := GetParamInt64(ctx, "notice_id")

	// 判断参数合法性
	if noticeId <= 0 {
		logger.WithContext(ctx).Warnf("[NoticeDelete] notice_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "公告id不存在")
	}

	// 删除公告
	err := service.NewNotice(ctx).DeleteNotice(noticeId)
	if err != nil {
		sysLogErrorf(ctx, "[NoticeDelete] 删除公告 %d 失败: err=%+v", noticeId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[NoticeDelete] 删除公告 %d 成功", noticeId)

	return RespJsonSuccess(ctx, nil)
}

// NoticePublishList 公告发布列表
func NoticePublishList(ctx *gin.Context) error {

	pageSize := GetParamIntDef(ctx, "page_size", 4)
	pageNum := GetParamIntDef(ctx, "page_num", 1)
	pageSize = utils.VerifyUint(pageSize, 4)
	pageNum = utils.VerifyUint(pageNum, 1)

	serviceNotice := service.NewNotice(ctx)
	// 获取公告列表
	notices, err := serviceNotice.GetPublishNoticesByLimit(pageSize, pageNum)
	if err != nil {
		sysLogErrorf(ctx, "[NoticePublishList] 获取公告列表失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取分页信息
	totalNum, err := serviceNotice.CountPublishNotices()
	if err != nil {
		sysLogErrorf(ctx, "[NoticePublishList] 获取公告分页信息失败: err=%s", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	data := map[string]interface{}{
		"list":      notices,
		"page_info": entity.GetPageInfo(totalNum, pageSize, pageNum),
	}
	return RespJsonSuccess(ctx, data)
}
