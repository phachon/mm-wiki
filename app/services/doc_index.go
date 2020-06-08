package services

import (
	"sync"

	"github.com/phachon/mm-wiki/app/models"
	"github.com/phachon/mm-wiki/global"

	"github.com/astaxie/beego/logs"
	"github.com/go-ego/riot/types"
)

var DocIndexService = NewDocIndexService()

type DocIndex struct {
}

func NewDocIndexService() *DocIndex {
	return &DocIndex{}
}

// ForceDelDocIdIndex 强制删除索引
func (di *DocIndex) ForceDelDocIdIndex(docId string) {
	if docId == "" {
		return
	}
	// add search index
	data := types.DocData{Content: ""}
	global.DocSearcher.IndexDoc(docId, data, true)
}

// UpdateDocIndex 更新单个文件的索引
func (di *DocIndex) ForceUpdateDocIndexByDocId(docId string) error {
	if docId == "" {
		return nil
	}
	doc, err := models.DocumentModel.GetDocumentByDocumentId(docId)
	if err != nil {
		return err
	}
	content, _, err := models.DocumentModel.GetDocumentContentByDocument(doc)
	if err != nil {
		return err
	}
	// add search index
	data := types.DocData{Content: content}
	global.DocSearcher.IndexDoc(docId, data, true)
	return nil
}

// UpdateDocIndex 更新单个文件的索引
func (di *DocIndex) UpdateDocIndex(doc map[string]string) {
	docId, ok := doc["document_id"]
	if !ok || docId == "" {
		return
	}
	content, _, err := models.DocumentModel.GetDocumentContentByDocument(doc)
	if err != nil {
		logs.Error("[UpdateDocIndex] get documentId=%s content err: %s", docId, err.Error())
		return
	}
	// add search index
	data := types.DocData{Content: content}
	global.DocSearcher.IndexDoc(docId, data)
}

// UpdateDocsIndex 批量更新多个文件的索引
func (di *DocIndex) UpdateDocsIndex(docs []map[string]string) {
	if len(docs) == 0 {
		return
	}
	wait := sync.WaitGroup{}
	for _, doc := range docs {
		if len(doc) == 0 {
			continue
		}
		wait.Add(1)
		go func(doc map[string]string) {
			docId, _ := doc["document_id"]
			defer func() {
				if e := recover(); e != nil {
					logs.Error("[UpdateAllDocIndex] get documentId=%s content panic: %v", docId, e)
				}
				wait.Done()
			}()
			di.UpdateDocIndex(doc)
		}(doc)
	}
	wait.Wait()
}

// UpdateAllDocIndex 更新所有的文档
func (di *DocIndex) UpdateAllDocIndex(batchNum int) {
	allDocs, err := models.DocumentModel.GetAllDocuments()
	if err != nil {
		logs.Error("[UpdateAllDocIndex] getAllDocuments err: %s", err.Error())
		return
	}
	batchDocs := di.getBatchDocs(allDocs, batchNum)
	for _, docs := range batchDocs {
		di.UpdateDocsIndex(docs)
	}
}

// 获取分批文档
func (di *DocIndex) getBatchDocs(allDocs []map[string]string, n int) [][]map[string]string {

	groupNum := len(allDocs) / n
	remainder := len(allDocs) % n
	res := [][]map[string]string{}
	if groupNum == 0 {
		res = append(res, allDocs)
		return res
	}
	for i := 0; i < groupNum; i++ {
		offset := i * n
		resItem := allDocs[offset : n+offset]
		res = append(res, resItem)
	}
	if remainder > 0 {
		res = append(res, allDocs[len(allDocs)-remainder:])
	}
	return res
}

// FlushIndex 所有索引
func (di *DocIndex) Flush() {
	global.DocSearcher.Flush()
}
