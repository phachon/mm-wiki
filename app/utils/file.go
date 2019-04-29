package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var File = NewFile()

func NewFile() *file {
	return &file{}
}

type file struct {
}

// get file contents
func (f *file) GetFileContents(filePath string) (content string, err error) {
	defer func(err *error) {
		e := recover()
		if e != nil {
			*err = fmt.Errorf("%s", e)
		}
	}(&err)
	bytes, err := ioutil.ReadFile(filePath)
	content = string(bytes)
	return
}

// file or path is exists
func (f *file) PathIsExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// path is empty
func (f *file) PathIsEmpty(path string) bool {
	fs, e := filepath.Glob(filepath.Join(path, "*"))
	if e != nil {
		return false
	}
	if len(fs) > 0 {
		return false
	}
	return true
}

// is write permission
func (f *file) IsWritable(filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			return errors.New("Error: Write permission denied.")
		} else {
			return err
		}
	}
	file.Close()
	return nil
}

// is read permission
func (f *file) IsReadable(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			return errors.New("Error: Read permission denied.")
		} else {
			return err
		}
	}
	file.Close()
	return nil
}

// is read and write permission
func (f *file) IsWriterReadable(file string) error {
	err := f.IsWritable(file)
	if err != nil {
		return err
	}
	err = f.IsReadable(file)
	if err != nil {
		return err
	}

	return nil
}

// read file data
func (f *file) ReadAll(path string) (data string, err error) {
	fi, err := os.Open(path)
	if err != nil {
		return
	}
	defer fi.Close()

	fd, err := ioutil.ReadAll(fi)
	return string(fd), nil
}

// write file
func (f *file) WriteFile(filename string, data string) (err error) {
	var dataByte = []byte(data)
	err = ioutil.WriteFile(filename, dataByte, 0666)
	if err != nil {
		return
	}
	return
}

// create file
func (f *file) CreateFile(filename string) error {
	newFile, err := os.Create(filename)
	defer newFile.Close()
	return err
}

// get dir all files
func (f *file) WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}
