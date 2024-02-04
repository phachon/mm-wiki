package utils

import (
	"errors"
	"os"
	"path"
)

var pathList = []string{"./", "./conf/", "../conf/", "./etc/", "../etc/"}

// SearchPath 查找配置文件路径并返回
func SearchPath(filename string, withEnv string) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for _, p := range pathList {
		file := path.Join(pwd, p, filename)
		if fileExist(file) {
			return file, nil
		}
		if withEnv != "" {
			file := path.Join(pwd, p, withEnv+"/", filename)
			if fileExist(file) {
				return file, nil
			}
		}
	}
	return "", errors.New(filename + " not found")
}

// fileExist 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
