package upload

import (
	"context"
	"mime/multipart"
)

// UplaoderName 上传器名称
type UplaoderName string

// UploaderHandler 上传器函数
type UploaderHandler func(opt *UploaderConfig) Uploader

var (
	// uploaderHandlers 上传器函数列表
	uploaderHandlers = make(map[UplaoderName]UploaderHandler)
)

// Uploader 定义上传接口
type Uploader interface {
	// GetName 获取上传器名称
	GetName() UplaoderName
	// UploadFile 上传文件
	UploadFile(ctx context.Context, file *multipart.FileHeader) (*UploadResult, error)
}

// UploadResult 上传结果
type UploadResult struct {
	Url string `json:"url"` // 文件URL
}

// UploaderConfig 上传器配置
type UploaderConfig struct {
	Domain              string `yaml:"domain"` // 文件域名
	LocalUploaderConfig `yaml:",inline"`
}

// RegisterUploaderHandler 注册上传器
func RegisterUploaderHandler(uploadName UplaoderName, uploader UploaderHandler) {
	uploaderHandlers[uploadName] = uploader
}

// GetUploaderHandler 获取上传器
func GetUploaderHandler(name UplaoderName) UploaderHandler {
	return uploaderHandlers[name]
}
