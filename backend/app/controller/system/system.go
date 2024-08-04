// Package system 系统模块控制器层
package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
)

// RespJsonSuccess 返回成功 json
func RespJsonSuccess(ctx *gin.Context, data interface{}) error {
	if data == nil {
		data = make(map[string]interface{})
	}
	ctx.JSON(http.StatusOK, &entity.Response{
		Code: 0,
		Data: data,
	})
	return nil
}

// RespJsonError 返回失败 json
func RespJsonError(ctx *gin.Context, code int32, message string) error {
	ctx.JSON(http.StatusOK, &entity.Response{
		Code:    code,
		Message: message,
		Data:    make(map[string]interface{}),
	})
	return nil
}
