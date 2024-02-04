// Package global 全局常量
package global

import "github.com/gin-gonic/gin"

// GinEngine Gin 框架实例
var GinEngine = gin.New()

const (
	HeaderKTRequestID  = "X-KT-Request-Id"  //  header 公共参数用户请求唯一id
	HeaderKTTimestamp  = "X-KT-Timestamp"   //  header 公共参数用户请求 uninx 时间戳
	HeaderKTLoginToken = "X-KT-Login-Token" //  header 公共参数用户登录后Token
	HeaderKTDebug      = "X-KT-Debug"       //  header 公共参数debug开关
)

const (
	DefAccountPass = "123456"
)
