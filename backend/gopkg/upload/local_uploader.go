package upload

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// UploaderLocal 本地上传器
const UploaderLocal UplaoderName = "local"

func init() {
	RegisterUploaderHandler(UploaderLocal, NewLocalUploader)
}

// LocalUploader 实现 Uploader 接口，用于本地文件上传
type LocalUploader struct {
	opt *UploaderConfig
}

// LocalUploaderConfig 本地上传器配置
type LocalUploaderConfig struct {
	LocalDir string `yaml:"local_dir"`
}

// NewLocalUploader 创建一个新的 LocalUploader
func NewLocalUploader(opt *UploaderConfig) Uploader {
	return &LocalUploader{
		opt: opt,
	}
}

// GetName 获取上传器名称
func (u *LocalUploader) GetName() UplaoderName {
	return UploaderLocal
}

// UploadFile 实现文件上传到本地文件系统
func (u *LocalUploader) UploadFile(ctx context.Context, file *multipart.FileHeader) (*UploadResult, error) {
	// 打开上传的文件
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	if u.opt == nil || u.opt.LocalDir == "" {
		return nil, fmt.Errorf("local uploader config is error")
	}
	uploadDir := u.opt.LocalDir
	// 创建目标文件
	dstPath := filepath.Join(uploadDir, file.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	// 将上传的文件内容复制到目标文件
	if _, err := io.Copy(dst, src); err != nil {
		return nil, err
	}
	resp := &UploadResult{
		Url: dstPath,
	}

	return resp, nil
}
