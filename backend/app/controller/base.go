package controller

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/gopkg/errors"
	klog "github.com/phachon/mm-wiki/gopkg/log"
	"github.com/phachon/mm-wiki/logger"
	"github.com/phachon/mm-wiki/utils"
)

// 一些基础的公共方法

// GetParamString 获取参数返回 string
func GetParamString(ctx *gin.Context, key string) string {
	queryVal, ok := ctx.GetQuery(key)
	if ok {
		return queryVal
	}
	return ctx.PostForm(key)
}

// GetParamStringDef 获取参数返回 string 带默认值
func GetParamStringDef(ctx *gin.Context, key string, def string) string {
	queryVal, ok := ctx.GetQuery(key)
	if ok {
		return queryVal
	}
	return ctx.DefaultPostForm(key, def)
}

// GetParamInt 获取参数返回 int
func GetParamInt(ctx *gin.Context, key string) int {
	queryVal := GetParamString(ctx, key)
	return utils.Convert.StringToInt(queryVal)
}

// GetParamIntDef 获取参数返回 int
func GetParamIntDef(ctx *gin.Context, key string, def int) int {
	queryVal := GetParamStringDef(ctx, key, fmt.Sprintf("%d", def))
	return utils.Convert.StringToInt(queryVal)
}

// GetParamInt64 获取参数返回 float64
func GetParamInt64(ctx *gin.Context, key string) int64 {
	queryVal := GetParamString(ctx, key)
	return utils.Convert.StringToInt64(queryVal)
}

// GetParamInt64Def 获取参数返回 float64
func GetParamInt64Def(ctx *gin.Context, key string, def int64) int64 {
	queryVal := GetParamStringDef(ctx, key, fmt.Sprintf("%d", def))
	return utils.Convert.StringToInt64(queryVal)
}

// GetParamFloat32 获取参数返回 float32
func GetParamFloat32(ctx *gin.Context, key string) float32 {
	queryVal := GetParamString(ctx, key)
	return utils.Convert.StringToFloat32(queryVal)
}

// GetParamFloat32Def 获取参数返回 float32
func GetParamFloat32Def(ctx *gin.Context, key string, def float32) float32 {
	queryVal := GetParamStringDef(ctx, key, fmt.Sprintf("%f", def))
	return utils.Convert.StringToFloat32(queryVal)
}

// GetParamFloat64 获取参数返回 float64
func GetParamFloat64(ctx *gin.Context, key string) float64 {
	queryVal := GetParamString(ctx, key)
	return utils.Convert.StringToFloat64(queryVal)
}

// GetParamFloat64Def 获取参数返回 float64
func GetParamFloat64Def(ctx *gin.Context, key string, def float64) float64 {
	queryVal := GetParamStringDef(ctx, key, fmt.Sprintf("%f", def))
	return utils.Convert.StringToFloat64(queryVal)
}

// GetParamMap 获取参数返回字符串数组
func GetParamMap(ctx *gin.Context, key string) map[string]string {
	queryVal, ok := ctx.GetQueryMap(key)
	if ok {
		return queryVal
	}
	return ctx.PostFormMap(key)
}

// GetPageInfo 获取分页信息
func GetPageInfo(pageSize int, pageNum int, defPageSize int) (pagination *entity.PageInfo, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, defPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	var totalNum int64
	return entity.GetPageInfo(totalNum, pageSize, pageNum), nil
}

// sysLogInfof 系统 info 日志
func sysLogInfof(ctx *gin.Context, format string, args ...interface{}) {
	sysDBLogf(ctx, int(klog.LevelInfo), format, args...)
	logger.WithContext(ctx).Infof(format, args...)
}

// sysLogWarnf 系统 warn 日志
func sysLogWarnf(ctx *gin.Context, format string, args ...interface{}) {
	sysDBLogf(ctx, int(klog.LevelWarn), format, args...)
	logger.WithContext(ctx).Warnf(format, args...)
}

// sysLogErrorf 系统 error 日志
func sysLogErrorf(ctx *gin.Context, format string, args ...interface{}) {
	sysDBLogf(ctx, int(klog.LevelError), format, args...)
	logger.WithContext(ctx).Errorf(format, args...)
}

// sysLogTracef 系统 trace 日志
// func sysLogTracef(ctx *gin.Context, format string, args ...interface{}) {
// 	sysDBLogf(ctx, int(klog.LevelTrace), format, args...)
// 	logger.WithContext(ctx).Tracef(format, args...)
// }

// sysDBLogf 系统日志写到 db
func sysDBLogf(ctx *gin.Context, level int, format string, args ...interface{}) {
	if ctx == nil {
		return
	}

	var getParams, _ = url.QueryUnescape(ctx.Request.URL.Query().Encode())
	var postParams, _ = url.QueryUnescape(ctx.Request.PostForm.Encode())
	message := fmt.Sprintf(format, args...)
	logEntity := new(entity.LogEntity)
	logEntity.Message = message
	logEntity.Level = level
	logEntity.Uri = ctx.Request.RequestURI
	logEntity.Get = getParams
	logEntity.Post = postParams
	logEntity.Ip = ctx.ClientIP()
	logEntity.AccountId = global.ContextValueLoginAccountID(ctx)
	logEntity.AccountName = global.ContextValueLoginAccountName(ctx)
	err := service.NewLog(ctx).Create(logEntity)
	if err != nil {
		logger.WithContext(ctx).Errorf("[sysDBLog] create system log err=%s", err.Error())
	}
}
