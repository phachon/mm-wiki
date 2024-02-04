package entity

import (
	"encoding/json"
)

// DebugParam Debug 参数
type DebugParam struct {
	SkipAuthLogin      bool `json:"skip_auth_login"`      // 跳过登录校验
	SkipAuthPremission bool `json:"skip_auth_premission"` // 跳过权限校验
}

// ParseDebugParam 解析 debug 参数
func ParseDebugParam(debug string) *DebugParam {
	debugParam := new(DebugParam)
	if len(debug) == 0 {
		return debugParam
	}
	jErr := json.Unmarshal([]byte(debug), &debugParam)
	if jErr != nil {
		return debugParam
	}
	return debugParam
}

// IsSkipAuthLogin 是否跳过登录校验
func (d *DebugParam) IsSkipAuthLogin() bool {
	return d.SkipAuthLogin
}

// IsSkipPremission 是否跳过权限校验
func (d *DebugParam) IsSkipPremission() bool {
	return d.SkipAuthPremission
}
