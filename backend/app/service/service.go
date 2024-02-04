// Package service 业务逻辑层
package service

import "github.com/phachon/mm-wiki/utils"

// PasswordEncode 密码加密
func PasswordEncode(password string) string {
	return utils.Encrypt.Md5Encode(password)
}
