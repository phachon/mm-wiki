package system

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
)

// LogList 日志列表
func LogList(ctx *gin.Context) error {

	pageSize := GetParamIntDef(ctx, "page_size", 20)
	pageNum := GetParamIntDef(ctx, "page_num", 1)
	keywordsStr := GetParamStringDef(ctx, "keywords", "")

	var keywords *entity.LogSearchKeywords
	if len(keywordsStr) > 0 {
		jErr := json.Unmarshal([]byte(keywordsStr), &keywords)
		if jErr != nil {
			logger.WithContext(ctx).Errorf("[LogList] GetLogsByLimit err=%s", jErr.Error())
		}
	}

	serviceLog := service.NewLog(ctx)
	// 获取日志列表
	logs, err := serviceLog.GetLogsByKeywordsAndLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[LogList] 获取日志列表失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取分页信息
	pageInfo, err := serviceLog.GetPageInfoLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[LogList] 获取日志分页信息失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	data := map[string]interface{}{
		"list":      logs,
		"page_info": pageInfo,
	}
	return RespJsonSuccess(ctx, data)
}

// LogInfo 日志详情
func LogInfo(ctx *gin.Context) error {

	logId := GetParamInt64(ctx, "log_id")
	if logId <= 0 {
		logger.WithContext(ctx).Warnf("[LogInfo] 日志ID不能为空")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "日志ID不能为空")
	}

	serviceLog := service.NewLog(ctx)
	logInfo, err := serviceLog.GetLogByLogId(logId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[LogInfo] GetLogByLogId err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), "获取日志详情失败")
	}
	if logInfo == nil {
		logger.WithContext(ctx).Warnf("[LogInfo] 日志不存在")
		return RespJsonError(ctx, int32(errors.BusinessRecordNotExistError), "日志不存在")
	}

	data := map[string]interface{}{
		"log_info": logInfo,
	}
	return RespJsonSuccess(ctx, data)
}
