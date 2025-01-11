// Package errors 错误吗定义
package errors

// ErrorCode 自定义错误码
type ErrorCode int32

// 前三位代表项目模块，后三位代表具体错误
var (
	// SuccessCode 成功返回码
	SuccessCode ErrorCode = 0

	/*  客户端异常错误码 1xx */
	ClientUnknownError      ErrorCode = 1000 // 客户端请求未知异常
	ClientReqParamEmpty     ErrorCode = 1001 // 客户端请求参数为空
	ClientReqParamWrongful  ErrorCode = 1002 // 客户端请求参数不合法
	ClientReqCommonParamErr ErrorCode = 1003 // 客户端请求公共参数不合法

	/*  业务逻辑相关错误 2xx  */
	BusinessUnknownError        ErrorCode = 2000 // 业务未知错误
	BusinessPermissionError     ErrorCode = 2001 // 没有操作权限
	BusinessForbiddenError      ErrorCode = 2002 // 禁止操作
	BusinessRecordNotExistError ErrorCode = 2003 // 记录不存在
	BusinessRecordExistError    ErrorCode = 2004 // 记录已存在
	BusinessPasswordError       ErrorCode = 2005 // 密码错误
	BusinessAuthErr             ErrorCode = 2100 // auth 认证失败
	BusinessAuthTokenErr        ErrorCode = 2101 // auth token 错误
	BusinessAuthTokenInvalid    ErrorCode = 2102 // auth token 失效
	BusinessAuthTokenMakeErr    ErrorCode = 2103 // auth token 生成失败
	BusinessAuthTokenParseErr   ErrorCode = 2104 // auth token 解析失败

	/*  第三方服务错误 3xx */
	ServerUnknownError      ErrorCode = 3000 // 下游服务未知异常
	ServerTimoutError       ErrorCode = 3001 // 下游服务超时
	ServerRefuseError       ErrorCode = 3002 // 下游服务拒绝请求
	ServerNotAvailableError ErrorCode = 3003 // 下游服务不可用

	/*  数据层错误 4xx */
	DalUnknownError    ErrorCode = 4000 // 系统未知错误
	DalMysqlErr        ErrorCode = 4101 // 数据库错误
	DalMysqlInsertErr  ErrorCode = 4102 // 数据库插入错误
	DalMysqlUpdateErr  ErrorCode = 4103 // 数据库更新错误
	DalMysqlSelectErr  ErrorCode = 4104 // 数据库查询错误
	DalMysqlDeleteErr  ErrorCode = 4105 // 数据库删除错误
	DalMysqlRowIdErr   ErrorCode = 4106 // 数据库更新失败
	DalDuplicateKeyErr ErrorCode = 4107 // 数据库主键冲突

	DalRedisError        ErrorCode = 4210 // Redis 错误
	DalRedisTimeoutError ErrorCode = 4211 // Redis 超时错误
	DalRedisAuthError    ErrorCode = 4212 // Redis Auth 错误
	DalRedisGetError     ErrorCode = 4213 // Redis Get 错误
	DalRedisSetError     ErrorCode = 4214 // Redis Set 错误
	DalRedisMGetError    ErrorCode = 4215 // Redis MGet 错误
	DalRedisZSetError    ErrorCode = 4216 // Redis ZSet 错误
	DalRedisHashError    ErrorCode = 4217 // Redis Hash 错误
	DalRedisDelError     ErrorCode = 4218 // Redis Del 错误

	/*  系统底层错误 5xx */
	SystemUnknownError ErrorCode = 5000 // 系统未知错误
	SystemNetIOError   ErrorCode = 5110 // 网络 IO 错误
	SystemFileIOError  ErrorCode = 5210 // 文件 IO 错误
)
