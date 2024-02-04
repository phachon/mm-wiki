package global

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/utils"
)

type BizContextKey string // 业务自定义 context key

const (
	BizContextKeyRequestID        BizContextKey = "biz_request_id"         // 业务 context key：请求ID
	BizContextKeyLoginAccountID   BizContextKey = "biz_login_account_id"   // 业务 context key：登录账号ID
	BizContextKeyLoginAccountName BizContextKey = "biz_login_account_name" // 业务 context key：登录账号名
	BizContextKeyLoginToken       BizContextKey = "biz_login_login"        // 业务 context key：登录token
	BizContextKeyDebugParam       BizContextKey = "biz_debug_param"        // 业务 context key：debug 参数
	BizContextKeyLogger           BizContextKey = "biz_logger"             // 业务 context key：logger
	BizContextKeyLoginIdentifys   BizContextKey = "biz_login_identifys"    // 业务 context key：登录账号权限标识
)

// ContextWithRequestID context 添加请求 ID
func ContextWithRequestID(ctx context.Context, loginToken string) context.Context {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		ginCtx.Set(string(BizContextKeyRequestID), loginToken)
		return ctx
	}
	return context.WithValue(ctx, BizContextKeyRequestID, loginToken)
}

// ContextValueRequestID context 添加请求 ID
func ContextValueRequestID(ctx context.Context) string {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		return ginCtx.GetString(string(BizContextKeyRequestID))
	}
	loginToken := ctx.Value(BizContextKeyRequestID)
	return fmt.Sprintf("%v", loginToken)
}

// ContextWithLoginAccountID context 添加登录账号ID
func ContextWithLoginAccountID(ctx context.Context, accountID int64) context.Context {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		ginCtx.Set(string(BizContextKeyLoginAccountID), accountID)
		return ctx
	}
	return context.WithValue(ctx, BizContextKeyLoginAccountID, accountID)
}

// ContextValueLoginAccountID context 获取登录账号ID
func ContextValueLoginAccountID(ctx context.Context) int64 {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		return ginCtx.GetInt64(string(BizContextKeyLoginAccountID))
	}
	accountIDIf := ctx.Value(BizContextKeyLoginAccountID)
	accountID, _ := utils.Convert.ToInt64(accountIDIf)
	return accountID
}

// ContextWithLoginAccountName context 添加登录账号名
func ContextWithLoginAccountName(ctx context.Context, accountName string) context.Context {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		ginCtx.Set(string(BizContextKeyLoginAccountName), accountName)
		return ctx
	}
	return context.WithValue(ctx, BizContextKeyLoginAccountName, accountName)
}

// ContextValueLoginAccountID context 获取登录账号ID
func ContextValueLoginAccountName(ctx context.Context) string {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		return ginCtx.GetString(string(BizContextKeyLoginAccountName))
	}
	accountName := ctx.Value(BizContextKeyLoginAccountName)
	return fmt.Sprintf("%v", accountName)
}

// ContextWithLoginToken context 添加登录 token
func ContextWithLoginToken(ctx context.Context, loginToken string) context.Context {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		ginCtx.Set(string(BizContextKeyLoginToken), loginToken)
		return ctx
	}
	return context.WithValue(ctx, BizContextKeyLoginToken, loginToken)
}

// ContextValueLoginToken context 获取登录 token
func ContextValueLoginToken(ctx context.Context) string {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		return ginCtx.GetString(string(BizContextKeyLoginToken))
	}
	loginToken := ctx.Value(BizContextKeyLoginToken)
	return fmt.Sprintf("%v", loginToken)
}

// ContextWithDebugParam context 添加 debug 参数
func ContextWithDebugParam(ctx context.Context, debugParam *entity.DebugParam) context.Context {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		ginCtx.Set(string(BizContextKeyDebugParam), debugParam)
		return ctx
	}
	return context.WithValue(ctx, BizContextKeyDebugParam, debugParam)
}

// ContextValueDebugParam context 获取登录 debug 参数
func ContextValueDebugParam(ctx context.Context) *entity.DebugParam {
	var debugParamVal interface{}
	if ginCtx, ok := ctx.(*gin.Context); ok {
		debugParamVal, _ = ginCtx.Get(string(BizContextKeyDebugParam))
	} else {
		debugParamVal = ctx.Value(BizContextKeyDebugParam)
	}
	if debugParam, ok := debugParamVal.(*entity.DebugParam); ok {
		return debugParam
	}
	return new(entity.DebugParam)
}

// ContextWithLoginIdentifys context 添加登录权限标识
func ContextWithLoginIdentifys(ctx context.Context, identifys map[string]int) context.Context {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		ginCtx.Set(string(BizContextKeyLoginIdentifys), identifys)
		return ctx
	}
	return context.WithValue(ctx, BizContextKeyLoginIdentifys, identifys)
}

// ContextValueLoginIdentifys context 获取登录权限标识
func ContextValueLoginIdentifys(ctx context.Context) map[string]int {
	var ctxVal interface{}
	if ginCtx, ok := ctx.(*gin.Context); ok {
		ctxVal, _ = ginCtx.Get(string(BizContextKeyLoginIdentifys))
	} else {
		ctxVal = ctx.Value(BizContextKeyLoginIdentifys)
	}
	if identifys, ok := ctxVal.(map[string]int); ok {
		return identifys
	}
	return make(map[string]int)
}
