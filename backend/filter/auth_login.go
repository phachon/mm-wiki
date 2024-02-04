package filter

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
)

var (
	// 需要忽略的校验登录的接口
	ignoreLoginAuthPath = map[string]bool{
		"/system/auth/login": true,
	}
)

// AuthLoginJWT 登录 token 认证
func AuthLoginJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		if authLoginJWTHandle(c) {
			c.Next()
		} else {
			c.Abort()
		}
	}
}

func authLoginJWTHandle(c *gin.Context) (keepNext bool) {
	path := c.Request.URL.Path
	_, ok := ignoreLoginAuthPath[path]
	if ok {
		return true
	}
	// 从 ctx 获取 debug 参数
	debugParam := global.ContextValueDebugParam(c)
	if debugParam != nil && debugParam.IsSkipAuthLogin() {
		return true
	}
	res := &entity.Response{
		Code:    int32(errors.SuccessCode),
		Message: "",
		Data:    nil,
	}
	// 从 ctx 获取 loginToken
	loginToken := global.ContextValueLoginToken(c)
	if loginToken == "" {
		res.Code = int32(errors.ClientReqParamEmpty)
		c.JSON(http.StatusOK, res)
		return false
	}
	// 解析 loginToken
	claims, err := service.ParseToken(loginToken)
	if err != nil {
		logger.Errorf("[AuthJWT] login auth jwt err=%s", err.Error())
		res.Code = err.GetErrCode()
		res.Message = err.GetErrMsg()
		c.JSON(http.StatusUnauthorized, res)
		return false
	}
	// 登录成功后设置到 Context 中
	global.ContextWithLoginAccountID(c, claims.AccountID)
	global.ContextWithLoginAccountName(c, claims.AccountName)
	global.ContextWithLoginToken(c, loginToken)

	requestID := global.ContextValueRequestID(c)
	// logger 增加 context 自定义字段
	logger.WithContextFields(c,
		"request_id", requestID,
		"login_account_id", fmt.Sprintf("%d", claims.AccountID),
		"login_account_name", claims.AccountName,
	)
	return true
}
