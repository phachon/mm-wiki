// Package service 业务逻辑层
package service

import (
	"strings"

	"github.com/phachon/mm-wiki/utils"
	"golang.org/x/crypto/bcrypt"
)

// PasswordEncode 密码加密（使用 bcrypt）
func PasswordEncode(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// 如果 bcrypt 失败，回退到 MD5（兼容旧数据）
		return utils.Encrypt.Md5Encode(password)
	}
	return string(hash)
}

// PasswordVerify 校验密码是否匹配
// 支持 bcrypt 和旧版 MD5 两种格式
func PasswordVerify(password string, encodedPassword string) bool {
	// 尝试 bcrypt 校验（bcrypt hash 以 $2a$ 或 $2b$ 开头）
	if strings.HasPrefix(encodedPassword, "$2a$") || strings.HasPrefix(encodedPassword, "$2b$") {
		err := bcrypt.CompareHashAndPassword([]byte(encodedPassword), []byte(password))
		return err == nil
	}
	// 兼容旧版 MD5 密码
	return utils.Encrypt.Md5Encode(password) == encodedPassword
}
