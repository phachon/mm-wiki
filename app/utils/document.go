package utils

import (
	"path/filepath"
	"os"
	"io/ioutil"
	"sync"
	"fmt"
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

// get document page file by parentPath
func (d *document) GetPageFileByParentPath(name string, docType int, parentPath string) (pageFile string){
	if docType == Document_Type_Page {
		pageFile = fmt.Sprintf("%s/%s%s", parentPath, name, Document_Page_Suffix)
	}else {
		pageFile = fmt.Sprintf("%s/%s/%s%s", parentPath, name, Document_Default_FileName, Document_Page_Suffix)
	}
	return
}

//get document path by spaceName
func (d *document) GetDefaultPageFileBySpaceName(name string) string {
	return fmt.Sprintf("%s/%s%s", name, Document_Default_FileName, Document_Page_Suffix)
}

// get document abs pageFile
func (d *document) GetAbsPageFileByPageFile(pageFile string) string {
	return d.RootAbsDir + "/" +pageFile
}

// get document content by pageFile
func (d *document) GetContentByPageFile(pageFile string) (content string , err error){
	return File.GetFileContents(d.GetAbsPageFileByPageFile(pageFile))
}

// create document
func (d *document) Create(pageFile string) error {
	if pageFile == "" {
		return nil
	}
	d.lock.Lock()
	absFilePath := d.GetAbsPageFileByPageFile(pageFile)
	absDir := filepath.Dir(absFilePath)
	err := os.MkdirAll(absDir, 0777)
	if err != nil {
		d.lock.Unlock()
		return err
	}
	d.lock.Unlock()
	return File.CreateFile(absFilePath)
}

// create and write document
func (d *document) CreateAndWrite(pageFile string, content string) error {
	if pageFile == "" {
		return nil
	}
	d.lock.Lock()
	absFilePath := d.GetAbsPageFileByPageFile(pageFile)
	absDir := filepath.Dir(absFilePath)
	err := os.MkdirAll(absDir, 0777)
	if err != nil {
		d.lock.Unlock()
		return err
	}
	d.lock.Unlock()
	return File.WriteFile(absFilePath, content)
}

// replace document content
func (d *document) Replace(pageFile string, content string) error {
	if pageFile == "" {
		return nil
	}
	d.lock.Lock()
	absFilePath := d.GetAbsPageFileByPageFile(pageFile)
	absDir := filepath.Dir(absFilePath)
	err := os.MkdirAll(absDir, 0777)
	if err != nil {
		d.lock.Unlock()
		return err
	}
	d.lock.Unlock()
	return ioutil.WriteFile(absFilePath, []byte(content), os.ModePerm)
}

// update document
func (d *document) Update(oldPageFile string, name string, content string, docType int) (err error) {

	filePath := filepath.Dir(oldPageFile)

	if docType == Document_Type_Page {
		newPageFile := filePath + "/" + name + Document_Page_Suffix
		if oldPageFile == newPageFile {
			err = d.Replace(oldPageFile, content)
			if err != nil {
				return err
			}
			return nil
		} else {
			err = d.CreateAndWrite(newPageFile, content)
			if err != nil {
				return err
			}
			return os.Remove(oldPageFile)
		}
	}else {
		newPageFile := filePath + "/" + name +"/"+ Document_Page_Suffix
		if oldPageFile == oldPageFile {
			err = d.Replace(oldPageFile, content)
			if err != nil {
				return err
			}
			return nil
		} else {
			err = d.CreateAndWrite(newPageFile, content)
			if err != nil {
				return err
			}
			return os.Remove(oldPageFile)
		}
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