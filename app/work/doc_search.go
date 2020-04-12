package work

import (
	"github.com/astaxie/beego/logs"
	"github.com/go-ego/riot/types"
	"github.com/phachon/mm-wiki/app/models"
	"github.com/phachon/mm-wiki/global"
	"sync"
	"time"
)

// 初始化文档搜索索引 Work
func InitDocSearchIndexWork(t time.Duration) {
	loadDocSearchIndex()
	go func() {
		defer func() {
			e := recover()
			if e != nil {
				logs.Info("load search index panic: %v", e)
			}
		}()
		for {
			select {
			case <-time.Tick(t):
				loadDocSearchIndex()
			}
		}
	}()
}

// 加载索引文件
func loadDocSearchIndex() {
	logs.Info("[loadSearchIndex] start load doc index")
	allDocs, err := models.DocumentModel.GetAllDocuments()
	if err != nil {
		logs.Error("[loadSearchIndex] getAllDocuments err: %s", err.Error())
		return
	}
	logs.Info("[loadSearchIndex] get all doc finish")
	wait := sync.WaitGroup{}
	for _, doc := range allDocs {
		if len(doc) == 0 {
			continue
		}
		wait.Add(1)
		go func(doc map[string]string) {
			docId, _ := doc["document_id"]
			defer func() {
				e := recover()
				if e != nil {
					logs.Error("[loadSearchIndex] get documentId=%s content panic: %v", docId, e)
				}
				wait.Done()
			}()
			var err error
			content, _, err := models.DocumentModel.GetDocumentContentByDocument(doc)
			if err != nil {
				logs.Error("[loadSearchIndex] get documentId=%s content err: %s", docId, err.Error())
				return
			}
			// add search index
			data := types.DocData{Content: content}
			global.DocSearcher.IndexDoc(docId, data)
		}(doc)
	}
	wait.Wait()
	logs.Info("[loadSearchIndex] add index finish")
	global.DocSearcher.Flush()
	logs.Info("[loadSearchIndex] index flush finish")
}
