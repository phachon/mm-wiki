package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var Document = NewDocument("./data", "./data/markdowns")

const (
	Document_Default_FileName = "README"
	Document_Page_Suffix      = ".md"
)

const (
	Document_Type_Page = 1
	Document_Type_Dir  = 2
)

func NewDocument(documentAbsDir string, markdownAbsDir string) *document {
	return &document{
		DocumentAbsDir: documentAbsDir,
		MarkdownAbsDir: markdownAbsDir,
	}
}

type document struct {
	DocumentAbsDir string
	MarkdownAbsDir string
	lock           sync.Mutex
}

// get document page file by parentPath
func (d *document) GetPageFileByParentPath(name string, docType int, parentPath string) (pageFile string) {
	if docType == Document_Type_Page {
		pageFile = fmt.Sprintf("%s/%s%s", parentPath, name, Document_Page_Suffix)
	} else {
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
	return d.MarkdownAbsDir + "/" + pageFile
}

// get document content by pageFile
func (d *document) GetContentByPageFile(pageFile string) (content string, err error) {
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
func (d *document) Update(oldPageFile string, name string, content string, docType int, nameIsChange bool) (err error) {

	d.lock.Lock()
	defer d.lock.Unlock()

	absOldPageFile := d.GetAbsPageFileByPageFile(oldPageFile)

	err = ioutil.WriteFile(absOldPageFile, []byte(content), os.ModePerm)
	if err != nil {
		return
	}
	if nameIsChange {
		filePath := filepath.Dir(absOldPageFile)
		if docType == Document_Type_Page {
			err = os.Rename(absOldPageFile, filePath+"/"+name+Document_Page_Suffix)
		} else {
			err = os.Rename(filePath, filepath.Dir(filePath)+"/"+name)
		}
		if err != nil {
			return
		}
	}
	return nil
}

func (d *document) UpdateSpaceName(oldSpaceName string, newName string) error {

	d.lock.Lock()
	defer d.lock.Unlock()

	spaceOldDir := d.GetAbsPageFileByPageFile(oldSpaceName)
	spaceNewDir := d.GetAbsPageFileByPageFile(newName)
	if spaceNewDir == spaceOldDir {
		return nil
	}
	err := os.Rename(spaceOldDir, spaceNewDir)
	return err
}

// delete document
func (d *document) Delete(path string, docType int) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	absPageFile := d.GetAbsPageFileByPageFile(path)

	ok, _ := File.PathIsExists(absPageFile)
	if !ok {
		return nil
	}
	if docType == Document_Type_Page {
		return os.Remove(absPageFile)
	}

	return os.RemoveAll(filepath.Dir(absPageFile))
}

func (d *document) DeleteSpace(name string) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	absSpaceDir := d.GetAbsPageFileByPageFile(name)

	ok, _ := File.PathIsExists(absSpaceDir)
	if !ok {
		return nil
	}

	return os.RemoveAll(absSpaceDir)
}

func (d *document) Move(movePath string, targetPath string, docType int) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	absOldPageFile := d.GetAbsPageFileByPageFile(movePath)
	absTargetPageFile := d.GetAbsPageFileByPageFile(targetPath)

	if docType == Document_Type_Page {
		return os.Rename(absOldPageFile, absTargetPageFile)
	}
	return os.Rename(filepath.Dir(absOldPageFile), filepath.Dir(absTargetPageFile))
}

// delete document attachment
func (d *document) DeleteAttachment(attachments []map[string]string) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	if len(attachments) == 0 {
		return nil
	}
	// delete attachment file
	for _, attachment := range attachments {
		if len(attachment) == 0 || attachment["path"] == "" {
			continue
		}
		file := filepath.Join(d.DocumentAbsDir, attachment["path"])
		_ = os.Remove(file)
	}
	return nil
}
