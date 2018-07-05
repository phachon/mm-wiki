package utils

import (
	"path/filepath"
	"fmt"
	"os"
	"io/ioutil"
	"sync"
)

var Document = NewDocument("./data")

const (
	Document_Default_FileName = "README"
	Document_Page_Suffix = ".md"
)

const (
	Document_Type_Page = 1
	Document_Type_Dir = 2
)

func NewDocument(rootAbsDir string) *document {
	return &document{
		RootAbsDir: rootAbsDir,
	}
}

type document struct {
	RootAbsDir string
	lock sync.Mutex
}

// get document path by parentPath
func (d *document) GetPathByParentPath(name string, docType int, parentPath string) (path string){
	parentDir := filepath.Dir(parentPath)
	if docType == Document_Type_Page {
		path = fmt.Sprintf("%s/%s%s", parentDir, name, Document_Page_Suffix)
	}else {
		path = fmt.Sprintf("%s/%s/%s%s", parentDir, name, Document_Default_FileName, Document_Page_Suffix)
	}
	return
}

// get document path by spaceName
func (d *document) GetPathBySpaceName(name string) string {
	return fmt.Sprintf("%s/%s%s", name, Document_Default_FileName, Document_Page_Suffix)
}

// get document abs path
func (d *document) GetAbsPathByPath(path string) string {
	return d.RootAbsDir + "/" +path
}

// get document content by path
func (d *document) GetContentByPath(path string) (content string , err error){
	return File.GetFileContents(d.GetAbsPathByPath(path))
}

// create document
func (d *document) Create(path string) error {
	if path == "" {
		return nil
	}
	d.lock.Lock()
	absPath := d.GetAbsPathByPath(path)
	absDir := filepath.Dir(absPath)
	err := os.MkdirAll(absDir, 0777)
	if err != nil {
		d.lock.Unlock()
		return err
	}
	d.lock.Unlock()
	return File.CreateFile(absPath)
}

// create and write document
func (d *document) CreateAndWrite(path string, content string) error {
	if path == "" {
		return nil
	}
	d.lock.Lock()
	absPath := d.GetAbsPathByPath(path)
	absDir := filepath.Dir(absPath)
	err := os.MkdirAll(absDir, 0777)
	if err != nil {
		d.lock.Unlock()
		return err
	}
	d.lock.Unlock()
	return File.WriteFile(path, content)
}

// replace document content
func (d *document) Replace(path string, content string) error {
	if path == "" {
		return nil
	}
	d.lock.Lock()
	absPath := d.GetAbsPathByPath(path)
	absDir := filepath.Dir(absPath)
	err := os.MkdirAll(absDir, 0777)
	if err != nil {
		d.lock.Unlock()
		return err
	}
	d.lock.Unlock()
	return ioutil.WriteFile(path, []byte(content), os.ModePerm)
}

// update document
func (d *document) Update(oldPath string, newPath string, content string) (err error) {
	if oldPath == newPath {
		err = d.Replace(oldPath, content)
		if err != nil {
			return err
		}
		return nil
	} else {
		err = d.CreateAndWrite(newPath, content)
		if err != nil {
			return err
		}
		return os.Remove(oldPath)
	}
}

// delete document
func (d *document) Delete(path string, docType int) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	if docType == Document_Type_Page {
		return os.Remove(path)
	}
	return os.Remove(filepath.Dir(path))
}