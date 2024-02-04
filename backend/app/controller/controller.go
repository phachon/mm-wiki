// Package controller 控制器层
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
)

// HandleFunc 控制器处理方法
type HandleFunc func(ctx *gin.Context) error

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
