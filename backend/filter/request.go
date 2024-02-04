package filter

import (
	"net/http"
	"strconv"
	"time"

	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/logger"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/global"
)

// DebugCosTime 请求耗时打印
func DebugCosTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		cost := time.Since(start)
		logger.WithContext(c).Infof("[RequestCosTime] path=%+v method=%+v status=%+v cost=%+v",
			path, c.Request.Method, c.Writer.Status(), cost)
	}
}

// RequestParse 请求公共参数解析
func RequestParse() gin.HandlerFunc {
	return func(c *gin.Context) {
		if requestParseHandle(c) {
			c.Next()
		} else {
			c.Abort()
		}
	}
}

func requestParseHandle(c *gin.Context) (keepNext bool) {
	// 404 直接返回，不继续往下执行拦截器
	if c.Writer.Status() == http.StatusNotFound {
		return false
	}

	// 解析 header 头 RequestID
	requestID := c.Request.Header.Get(global.HeaderKTRequestID)
	if requestID == "" {
		requestID = strconv.FormatInt(time.Now().UnixNano(), 10)
		c.Request.Header.Set(global.HeaderKTRequestID, requestID)
	}
	global.ContextWithRequestID(c, requestID)

	// 解析 header 头 LoginToken
	loginToken := c.Request.Header.Get(global.HeaderKTLoginToken)
	global.ContextWithLoginToken(c, loginToken)

	// 解析 header 头 Debug 参数
	debugStr := c.Request.Header.Get(global.HeaderKTDebug)
	debugParam := entity.ParseDebugParam(debugStr)
	global.ContextWithDebugParam(c, debugParam)

	// logger 增加 context 自定义字段
	logger.WithContextFields(c,
		"request_id", requestID,
		"login_token", loginToken,
	)
	return true
}
