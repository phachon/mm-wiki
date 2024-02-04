package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

var Encrypt = NewEncrypt()

const (
	Base64Table = "1234567890poiuytreqwasdfghjklmnbvcxzQWERTYUIOPLKJHGFDSAZXCVBNM-_"
)

type encrypt struct{}

func NewEncrypt() *encrypt {
	return &encrypt{}
}

// Base64Encode base64 加密
func (encrypt *encrypt) Base64Encode(str string) string {
	var coder = base64.NewEncoding(Base64Table)
	var src []byte = []byte(str)
	return string([]byte(coder.EncodeToString(src)))
}

// Base64EncodeBytes base64 加密
func (encrypt *encrypt) Base64EncodeBytes(bytes []byte) []byte {
	var coder = base64.NewEncoding(Base64Table)
	return []byte(coder.EncodeToString(bytes))
}

// Base64Decode base64 解密
func (encrypt *encrypt) Base64Decode(str string) (string, error) {
	var src []byte = []byte(str)
	var coder = base64.NewEncoding(Base64Table)
	by, err := coder.DecodeString(string(src))
	return string(by), err
}

// Base64DecodeBytes base64 解密
func (encrypt *encrypt) Base64DecodeBytes(str string) ([]byte, error) {
	var coder = base64.NewEncoding(Base64Table)
	return coder.DecodeString(str)
}

// Md5Encode md5加密
func (encrypt *encrypt) Md5Encode(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
