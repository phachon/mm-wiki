package services

import (
	"strconv"
	"sync"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/phachon/mm-wiki/app/models"
	"github.com/phachon/mm-wiki/global"

)

var DocIndexService = NewDocIndexService()

type DocIndex struct {
}

type DocContent struct {
	ID      float64
	Content string
	Time    float64
}

func NewDocIndexService() *DocIndex {
	return &DocIndex{}
}

func (di *DocIndex) IsUpdateDocIndex() bool {
	fulltextSearchOpen := models.ConfigModel.GetConfigValueByKey(models.ConfigKeyFulltextSearch, "0")
	return fulltextSearchOpen == "1"
}

// ForceDelDocIdIndex 强制删除索引
func (di *DocIndex) ForceDelDocIdIndex(docId string) {
	if docId == "" {
		return
	}
	if !di.IsUpdateDocIndex() {
		return
	}
	// add search index
	global.SearchIndex.Delete(docId)

}

// UpdateDocIndex 更新单个文件的索引
func (di *DocIndex) ForceUpdateDocIndexByDocId(docId string) error {
	if docId == "" {
		return nil
	}
	if !di.IsUpdateDocIndex() {
		return nil
	}
	doc, err := models.DocumentModel.GetDocumentByDocumentId(docId)
	if err != nil {
		return err
	}
	_, _, err = models.DocumentModel.GetDocumentContentByDocument(doc)
	if err != nil {
		return err
	}
<<<<<<< HEAD
	// add search index
	DocID, err := strconv.ParseFloat(docId, 64)
	if err != nil {
		logs.Error("string转float64失败：err %+v", err)
	}
	Uptime, err := strconv.ParseFloat(doc["update_time"], 64)
	if err != nil {
		logs.Error("string转float64失败：err %+v", err)
	}
	global.SearchIndex.Index(docId, DocContent{ID: DocID, Content: content, Time: Uptime})
=======
	// todo add search index
>>>>>>> master
	return nil
}

// UpdateDocIndex 更新单个文件的索引
func (di *DocIndex) UpdateDocIndex(doc map[string]string) {
	docId, ok := doc["document_id"]
	if !ok || docId == "" {
		return
	}
	if !di.IsUpdateDocIndex() {
		return
	}
	_, _, err := models.DocumentModel.GetDocumentContentByDocument(doc)
	if err != nil {
		logs.Error("[UpdateDocIndex] get documentId=%s content err: %s", docId, err.Error())
		return
	}
<<<<<<< HEAD
	// add search index
	DocID, err := strconv.ParseFloat(docId, 64)
	if err != nil {
		logs.Error("string转float64失败：err %+v", err)
	}
	Uptime, err := strconv.ParseFloat(doc["update_time"], 64)
	if err != nil {
		logs.Error("string转float64失败：err %+v", err)
	}
	global.SearchIndex.Index(docId, DocContent{ID: DocID, Content: content, Time: Uptime})
=======
	// todo add search index
>>>>>>> master
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

// UpdateDocsIndex 批量更新多个文件的索引
func (di *DocIndex) UpdateDocsIndexByID(docIDs []string) {
	if len(docIDs) == 0 {
		return
	}
	wait := sync.WaitGroup{}
	for _, docID := range docIDs {
		wait.Add(1)
		go func(docID string) {
			defer func() {
				if e := recover(); e != nil {
					logs.Error("[UpdateDocsIndexByID] get documentId=%s content panic: %v", docID, e)
				}
				wait.Done()
			}()
			di.ForceUpdateDocIndexByDocId(docID)
		}(docID)
	}
	wait.Wait()
}

// UpdateAllDocIndex 更新所有的文档
func (di *DocIndex) UpdateAllDocIndex(batchNum int) {
	if !di.IsUpdateDocIndex() {
		return
	}
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

// UpdateDocIndexByDocId 通过DocID更新所有的文档
func (di *DocIndex) UpdateDocIndexByDocId(batchNum int, DocIDs []string) {
	if !di.IsUpdateDocIndex() {
		return
	}
	// 分批
	batchDocs := di.getBatchDocIDs(DocIDs, batchNum)

	for _, docs := range batchDocs {
		for _, ID := range docs {
			di.ForceUpdateDocIndexByDocId(ID)
		}
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

<<<<<<< HEAD
// 获取分批文档ID
func (di *DocIndex) getBatchDocIDs(allDocs []string, n int) [][]string {

	groupNum := len(allDocs) / n
	remainder := len(allDocs) % n
	res := [][]string{}
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

func (di *DocIndex) CheckDocIndexs() {
	for {
		now := time.Now()                                                                    //获取当前时间，放到now里面，要给next用
		next := now.Add(time.Hour * 24)                                                      //通过now偏移24小时
		next = time.Date(next.Year(), next.Month(), next.Day(), 2, 0, 0, 0, next.Location()) //获取下一个凌晨的日期
		t := time.NewTimer(next.Sub(now))                                                    //计算当前时间到凌晨的时间间隔，设置一个定时器
		<-t.C
		logs.Info("开始检查文章索引")
		var start, step int64
		start, step = 0, 5000
		var updateList []string
		for {
			IDs, err := models.DocumentModel.GetAllDocumentsID(start, start+step)
			if err != nil {
				panic(err)
			}
			if len(IDs) == 0 {
				break
			}
			Indexs, err := models.DocumentModel.GetAllDocumentsIndex(float64(start), float64(start+step))
			if err != nil {
				panic(err)
			}
			updateList = append(updateList, difference(IDs, Indexs)...)
			start += step
		}
		di.UpdateDocIndexByDocId(50, updateList)
		logs.Info("文章索引检查完成")
	}
}

//求差集 slice1-并集
func difference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

//求交集
func intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
=======
// FlushIndex 所有索引
func (di *DocIndex) Flush() {
>>>>>>> master
}
