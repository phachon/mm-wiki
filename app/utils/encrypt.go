package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

var Encrypt = NewEncrypt()

const (
	BASE_64_TABLE = "1234567890poiuytreqwasdfghjklmnbvcxzQWERTYUIOPLKJHGFDSAZXCVBNM-_"
)

type encrypt struct{}

func NewEncrypt() *encrypt {
	return &encrypt{}
}

//base64 加密
func (encrypt *encrypt) Base64Encode(str string) string {
	var coder = base64.NewEncoding(BASE_64_TABLE)
	var src []byte = []byte(str)
	return string([]byte(coder.EncodeToString(src)))
}

//base64 加密
func (encrypt *encrypt) Base64EncodeBytes(bytes []byte) []byte {
	var coder = base64.NewEncoding(BASE_64_TABLE)
	return []byte(coder.EncodeToString(bytes))
}

//base64 解密
func (encrypt *encrypt) Base64Decode(str string) (string, error) {
	var src []byte = []byte(str)
	var coder = base64.NewEncoding(BASE_64_TABLE)
	by, err := coder.DecodeString(string(src))
	return string(by), err
}

//base64 解密
func (encrypt *encrypt) Base64DecodeBytes(str string) ([]byte, error) {
	var coder = base64.NewEncoding(BASE_64_TABLE)
	return coder.DecodeString(str)
}

//md5加密
func (encrypt *encrypt) Md5Encode(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
