package errors

import "fmt"

// BizError 业务错误定义
type BizError interface {
	GetErrCode() int32 // 获取错误码
	GetErrMsg() string // 获取错误信息
	Error() string     // 返回错误信息
}

// DefaultBizError 默认错误定义
type DefaultBizError struct {
	errCode ErrorCode
	errMsg  string
}

// Errorf 默认错误
func Errorf(errCode ErrorCode, format string, a ...interface{}) BizError {
	return &DefaultBizError{
		errCode: errCode,
		errMsg:  fmt.Sprintf(format, a...),
	}
}

// Error 默认错误
func Error(errCode ErrorCode, format string) BizError {
	return &DefaultBizError{
		errCode: errCode,
		errMsg:  fmt.Sprintf(format),
	}
}

// GetErrCode 获取错误码
func (de *DefaultBizError) GetErrCode() int32 {
	return int32(de.errCode)
}

// GetErrMsg 获取错误信息
func (de *DefaultBizError) GetErrMsg() string {
	return de.errMsg
}

// Error 返回错误信息
func (de *DefaultBizError) Error() string {
	return fmt.Sprintf("code:%v msg:%v", de.GetErrCode(), de.GetErrMsg())
}
