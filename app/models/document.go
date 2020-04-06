package models

import (
	"fmt"
	"github.com/phachon/mm-wiki/app/utils"
	"github.com/snail007/go-activerecord/mysql"
	"strconv"
	"strings"
	"time"
)

const (
	Document_Delete_True  = 1
	Document_Delete_False = 0

	Document_Type_Page = 1
	Document_Type_Dir  = 2
)

const Table_Document_Name = "document"

type Document struct {
}

var DocumentModel = Document{}

// get document by document_id
func (d *Document) GetDocumentByDocumentId(documentId string) (document map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"document_id": documentId,
		"is_delete":   Document_Delete_False,
	}))
	if err != nil {
		return
	}
	document = rs.Row()
	return
}

// get documents by parent_id
func (d *Document) GetDocumentsByParentId(parentId string) (documents []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"parent_id": parentId,
		"is_delete": Document_Delete_False,
	}))
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

// get document by name
func (d *Document) GetDocumentsByName(name string) (documents []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"name":      name,
		"is_delete": Document_Delete_False,
	}))
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

// get document by name and spaceId
func (d *Document) GetDocumentByNameAndSpaceId(name string, spaceId string) (document map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"name":      name,
		"space_id":  spaceId,
		"is_delete": Document_Delete_False,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	document = rs.Row()
	return
}

// get document by name and spaceId
func (d *Document) GetDocumentByNameParentIdAndSpaceId(name string, parentId string, spaceId string, docType int) (document map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"name":      name,
		"space_id":  spaceId,
		"parent_id": parentId,
		"type":      docType,
		"is_delete": Document_Delete_False,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	document = rs.Row()
	return
}

// get document by name and spaceId
func (d *Document) GetDocumentByParentIdAndSpaceId(parentId string, spaceId string, docType int) (document map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"space_id":  spaceId,
		"parent_id": parentId,
		"type":      docType,
		"is_delete": Document_Delete_False,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	document = rs.Row()
	return
}

// get document by name and spaceId
func (d *Document) GetDocumentsByParentIdAndSpaceIdOnly(parentId string, spaceId string) (document []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"space_id":  spaceId,
		"parent_id": parentId,
		"is_delete": Document_Delete_False,
	}))
	if err != nil {
		return
	}
	document = rs.Rows()
	return
}

// get max sequence
func (d *Document) GetDocumentMaxSequence(parentId string, spaceId string) (sequence int, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"space_id":  spaceId,
		"parent_id": parentId,
	}).OrderBy("sequence", "desc").Limit(0, 1))

	if err != nil {
		return
	}

	document := rs.Row()
	sequenceStr := document["sequence"]
	sequence, err = strconv.Atoi(sequenceStr)
	return
}

// get all document after this document by sequence, if sequence eq 0, by create_time
func (d *Document) GetDocumentAllAfterSequence(spaceId string, documentId string) (documents []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet

	// get this document info
	nextDocument, err := d.GetDocumentByDocumentId(documentId)

	sequenceStr := nextDocument["sequence"]
	sequence, err := strconv.Atoi(sequenceStr)

	//sql := "select * from " + Table_Document_Name
	now := time.Now().Unix()

	if sequence == 0 {
		createTimeStr := nextDocument["create_time"]
		rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
			"create_time <":  strconv.FormatInt(now, 10),
			"create_time >=": createTimeStr,
			"space_id":       spaceId,
			"is_delete":      Document_Delete_False,
		}).OrderBy("create_time", "asc"))

		//sql += " where create_time > " + strconv.FormatInt(now, 10) + " and create_time > " + createTimeStr + " order by create_time asc"
	} else {
		//sql += " where sequence > " + sequenceStr + " order by sequence asc"
		rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
			"sequence >=": sequenceStr,
			"space_id":    spaceId,
			"is_delete":   Document_Delete_False,
		}).OrderBy("sequence", "asc"))
	}

	if err != nil {
		return
	}

	documents = rs.Rows()
	return
}

// delete document by document_id
func (d *Document) DeleteDBAndFile(documentId string, userId string, pageFile string, docType string) (err error) {
	db := G.DB()
	tx, err := db.Begin(db.Config)
	if err != nil {
		return
	}
	_, err = db.ExecTx(db.AR().Update(Table_Document_Name, map[string]interface{}{
		"is_delete":    Document_Delete_True,
		"update_time":  time.Now().Unix(),
		"edit_user_id": userId,
	}, map[string]interface{}{
		"document_id": documentId,
	}), tx)
	if err != nil {
		tx.Rollback()
		return
	}

	// delete document file
	err = utils.Document.Delete(pageFile, utils.Convert.StringToInt(docType))
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		return
	}

	// create document log
	go func() {
		LogDocumentModel.DeleteAction(userId, documentId)
	}()

	// delete follow doc
	go func() {
		FollowModel.DeleteByObjectIdType(documentId, fmt.Sprintf("%d", Follow_Type_Doc))
	}()

	// delete collect doc
	go func() {
		CollectionModel.DeleteByResourceIdType(documentId, fmt.Sprintf("%d", Collection_Type_Doc))
	}()

	return
}

// insert document
func (d *Document) Insert(documentValue map[string]interface{}) (id int64, err error) {

	db := G.DB()
	// start db begin
	tx, err := db.Begin(db.Config)
	if err != nil {
		return
	}

	// 处理同级排序编号
	parentId := documentValue["parent_id"].(string)
	spaceId := documentValue["space_id"].(string)

	sequence, err := d.GetDocumentMaxSequence(parentId, spaceId)

	if err != nil {
		sequence = 0
	}

	sequence += 1
	documentValue["sequence"] = strconv.Itoa(sequence)

	var rs *mysql.ResultSet
	documentValue["create_time"] = time.Now().Unix()
	documentValue["update_time"] = time.Now().Unix()
	rs, err = db.ExecTx(db.AR().Insert(Table_Document_Name, documentValue), tx)
	if err != nil {
		tx.Rollback()
		return
	}
	id = rs.LastInsertId

	// create document page file
	document := map[string]string{
		"space_id":  documentValue["space_id"].(string),
		"parent_id": documentValue["parent_id"].(string),
		"name":      documentValue["name"].(string),
		"type":      fmt.Sprintf("%d", documentValue["type"].(int)),
		"path":      documentValue["path"].(string),
	}
	_, pageFile, err := d.GetParentDocumentsByDocument(document)
	err = utils.Document.Create(pageFile)
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit()
	if err != nil {
		return
	}

	// create document log
	go func() {
		LogDocumentModel.CreateAction(documentValue["create_user_id"].(string), fmt.Sprintf("%d", id))
	}()

	// follow document
	go func() {
		FollowModel.CreateAutoFollowDocument(documentValue["create_user_id"].(string), fmt.Sprintf("%d", id))
	}()
	return
}

// update document by document_id
func (d *Document) Update(documentId string, documentValue map[string]interface{}, comment string) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	documentValue["update_time"] = time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_Document_Name, documentValue, map[string]interface{}{
		"document_id": documentId,
		"is_delete":   Document_Delete_False,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId

	// create document log
	go func() {
		LogDocumentModel.UpdateAction(documentValue["edit_user_id"].(string), documentId, comment)
	}()

	// follow document
	go func() {
		FollowModel.CreateAutoFollowDocument(documentValue["edit_user_id"].(string), documentId)
	}()
	return
}

// update sequence
func (d *Document) UpdateSequence(spaceId string, movedDocumentId string, moveType string, nextDocumentId string) (affectCount int64, err error) {

	// 获取目标文档信息
	nexDocumentFormDb, err := d.GetDocumentByDocumentId(nextDocumentId)
	nextSequenceStr := nexDocumentFormDb["sequence"]
	parentId := nexDocumentFormDb["parent_id"]

	// 本层级的最大排序号,
	maxSequence, err := d.GetDocumentMaxSequence(parentId, spaceId)

	// 需要修改的批量文档数据
	var updateBatchDocuments []map[string]string

	// 为0时, 从未排序过将所有程序重新排序
	if maxSequence == 0 {
		// 查询此空间下此父类下的所有文档
		allDocuments, err := d.GetDocumentsByParentIdAndSpaceIdOnly(parentId, spaceId)
		if err != nil {
		}

		// 即将移动文档的数组下标
		willBeMovedIndex := 0
		willBeMovedSequence := 0
		currSequence := 0
		for i := 0; i < len(allDocuments); i++ {
			allDocument := allDocuments[i]

			currDocumentId := allDocument["document_id"]

			// 跳过需要移动文档的处理
			if movedDocumentId == currDocumentId {
				willBeMovedIndex = i
				continue
			}

			currSequence += 1
			// 检测排序模型
			if moveType == "prev" && currDocumentId == nextDocumentId {
				currSequence += 1
				willBeMovedSequence = currSequence

				currSequence += 1
				allDocument["sequence"] = strconv.Itoa(currSequence)
			} else if moveType == "next" && currDocumentId == nextDocumentId {
				allDocument["sequence"] = strconv.Itoa(currSequence)

				currSequence += 1
				willBeMovedSequence = currSequence
			} else {
				allDocument["sequence"] = strconv.Itoa(currSequence)
			}
		}

		allDocuments[willBeMovedIndex]["sequence"] = strconv.Itoa(willBeMovedSequence)
		updateBatchDocuments = allDocuments
	} else {
		// moveType, prev排在目标文档前面, next排在目标文档后面
		nextSequence, _ := strconv.Atoi(nextSequenceStr)

		if nextSequence == 0 {
			nextSequence = 1
		}

		if moveType == "prev" {
			//nextSequence -= 1
		}

		if moveType == "next" {
			nextSequence += 1
		}
		nextSequenceStr := strconv.Itoa(nextSequence)
		movedDocumentMap := map[string]interface{}{
			"sequence": nextSequenceStr,
		}

		db := G.DB()
		var rs *mysql.ResultSet
		rs, err = db.Exec(db.AR().Update(Table_Document_Name, movedDocumentMap, map[string]interface{}{
			"document_id": movedDocumentId,
			"is_delete":   Document_Delete_False,
		}))

		if err != nil {
			return
		}

		affected := rs.RowsAffected
		if affected > 0 {
		}

		// 将某个排序号后面的所有文档查询出来, 如果这个排序号是0, 则以时间为排序, 将后面的所有文档都+1变更
		documents, err := d.GetDocumentAllAfterSequence(spaceId, nextDocumentId)

		if err != nil {
		}

		documentCount := len(documents)
		var updateAfterDocuments []map[string]string
		if documentCount > 0 {
			for i := 0; i < documentCount; i++ {
				currDocument := documents[i]
				currDocumentId := currDocument["document_id"]

				if currDocumentId == nextDocumentId || currDocumentId == movedDocumentId {
					continue
				}

				if err != nil {
				}

				nextSequence += 1
				currDocument["sequence"] = strconv.Itoa(nextSequence)
				updateAfterDocuments = append(updateAfterDocuments, currDocument)
			}
		}

		updateBatchDocuments = updateAfterDocuments
	}

	documentMaps := utils.MapString2Interface(updateBatchDocuments)

	db := G.DB()
	var updateBatchRs *mysql.ResultSet
	updateBatchRs, err = db.Exec(db.AR().UpdateBatch(Table_Document_Name, documentMaps, []string{"document_id"}))

	if err != nil {

	}

	rowsAffected := updateBatchRs.RowsAffected
	if rowsAffected > 0 {

	}

	return rowsAffected + 1, err
}

// move document
func (d *Document) MoveDBAndFile(documentId string, updateValue map[string]interface{}, oldPageFile string, newPageFile string, docType string, comment string) (id int64, err error) {

	db := G.DB()
	tx, err := db.Begin(db.Config)
	if err != nil {
		return
	}
	var rs *mysql.ResultSet
	updateValue["update_time"] = time.Now().Unix()
	rs, err = db.ExecTx(db.AR().Update(Table_Document_Name, updateValue, map[string]interface{}{
		"document_id": documentId,
		"is_delete":   Document_Delete_False,
	}), tx)
	if err != nil {
		tx.Rollback()
		return
	}
	id = rs.LastInsertId

	err = utils.Document.Move(oldPageFile, newPageFile, utils.Convert.StringToInt(docType))
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit()
	if err != nil {
		return
	}

	// create document log
	go func() {
		LogDocumentModel.UpdateAction(updateValue["edit_user_id"].(string), documentId, comment)
	}()

	return
}

// update document by document_id
func (d *Document) UpdateDBAndFile(documentId string, document map[string]string, documentContent string, updateValue map[string]interface{}, comment string) (id int64, err error) {

	// get document page file
	_, oldPageFile, err := DocumentModel.GetParentDocumentsByDocument(document)
	if err != nil {
		return
	}
	// begin update
	db := G.DB()
	tx, err := db.Begin(db.Config)
	if err != nil {
		return
	}
	var rs *mysql.ResultSet
	updateValue["update_time"] = time.Now().Unix()
	rs, err = db.ExecTx(db.AR().Update(Table_Document_Name, updateValue, map[string]interface{}{
		"document_id": documentId,
		"is_delete":   Document_Delete_False,
	}), tx)
	if err != nil {
		tx.Rollback()
		return
	}
	id = rs.LastInsertId

	// update document file
	docType := utils.Convert.StringToInt(document["type"])
	nameIsChange := false
	if updateValue["name"].(string) != document["name"] {
		nameIsChange = true
	}
	err = utils.Document.Update(oldPageFile, updateValue["name"].(string), documentContent, docType, nameIsChange)
	if err != nil {
		tx.Rollback()
		return
	}

	// commit
	err = tx.Commit()
	if err != nil {
		return
	}

	// create document log
	go func() {
		LogDocumentModel.UpdateAction(updateValue["edit_user_id"].(string), documentId, comment)
	}()

	// create follow doc
	go func() {
		FollowModel.CreateAutoFollowDocument(updateValue["edit_user_id"].(string), documentId)
	}()

	return
}

// get all documents
func (d *Document) GetDocumentsBySpaceId(spaceId string) (documents []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Document_Name).Where(map[string]interface{}{
			"space_id":  spaceId,
			"is_delete": Document_Delete_False,
		}))
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

// get document by spaceId and parentId
func (d *Document) GetDocumentsBySpaceIdAndParentId(spaceId string, parentId string) (documents []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Document_Name).Where(map[string]interface{}{
			"space_id":  spaceId,
			"parent_id": parentId,
			"is_delete": Document_Delete_False,
		}))
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

// get document by spaceId
func (d *Document) GetSpaceDefaultDocument(spaceId string) (document map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Document_Name).Where(map[string]interface{}{
			"space_id":  spaceId,
			"parent_id": "0",
			"is_delete": Document_Delete_False,
		}).Limit(0, 1))
	if err != nil {
		return
	}
	document = rs.Row()
	return
}

// get document by spaceId
func (d *Document) GetAllSpaceDocuments(spaceId string) (documents []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Document_Name).Where(map[string]interface{}{
			"space_id":    spaceId,
			"parent_id >": "0",
			"is_delete":   Document_Delete_False,
		}).OrderBy("sequence", "ASC"))
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

// get document count
func (d *Document) CountDocumentsBySpaceId(spaceId string) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			Select("count(*) as total").
			From(Table_Document_Name).
			Where(map[string]interface{}{
				"space_id":  spaceId,
				"is_delete": Document_Delete_False,
			}))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get document count
func (d *Document) CountDocuments() (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			Select("count(*) as total").
			From(Table_Document_Name).
			Where(map[string]interface{}{
				"is_delete": Document_Delete_False,
			}))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get document by name
func (d *Document) GetDocumentsByLikeName(name string) (documents []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"name Like": "%" + name + "%",
		"is_delete": Document_Delete_False,
	}))
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

// get document link name and limit
func (d *Document) GetDocumentsByLikeNameAndLimit(name string, limit int, number int) (documents []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"name Like": "%" + name + "%",
		"is_delete": Document_Delete_False,
	}).Limit(limit, number))
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

// count document like name
func (d *Document) CountDocumentsLikeName(name string) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			Select("count(*) as total").
			From(Table_Document_Name).
			Where(map[string]interface{}{
				"name Like": "%" + name + "%",
				"is_delete": Document_Delete_False,
			}))
	if err != nil {
		return
	}
	count = utils.Convert.StringToInt64(rs.Value("total"))
	return
}

// get document by spaceId and document_ids
func (d *Document) GetDocumentsByDocumentIds(documentIds []string) (documents []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"document_id": documentIds,
		"is_delete":   Document_Delete_False,
	}))
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

func (d *Document) GetAllDocumentsByDocumentIds(documentIds []string) (documents []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"document_id": documentIds,
	}))
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

func (d *Document) GetAllDocuments() (documents []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Document_Name).Where(map[string]interface{}{
		"is_delete": Document_Delete_False,
	}))
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

func (d *Document) GetParentDocumentsByDocument(document map[string]string) (parentDocuments []map[string]string, pageFile string, err error) {

	if document["parent_id"] == "0" {
		parentDocuments = append(parentDocuments, document)
		pageFile = utils.Document.GetDefaultPageFileBySpaceName(document["name"])
	} else {
		documentsIds := strings.Split(document["path"], ",")
		parentDocuments, err = d.GetDocumentsByDocumentIds(documentsIds)
		if err != nil {
			return
		}
		var parentPath = ""
		for _, parentDocument := range parentDocuments {
			parentPath += parentDocument["name"] + "/"
		}
		parentPath = strings.TrimRight(parentPath, "/")
		pageFile = utils.Document.GetPageFileByParentPath(document["name"], utils.Convert.StringToInt(document["type"]), parentPath)
	}
	return
}

func (d *Document) GetParentDocumentsByPath(path string) (parentDocuments []map[string]string, err error) {
	documentsIds := strings.Split(path, ",")
	parentDocuments, err = d.GetDocumentsByDocumentIds(documentsIds)
	if err != nil {
		return
	}
	return
}

func (d *Document) GetSpaceIdsOrderByCountDocumentLimit(limit int) (documents []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	sql := db.AR().Select("space_id, count('space_id') as total").
		From(Table_Document_Name).Where(map[string]interface{}{
		"is_delete": Document_Delete_False,
	}).
		GroupBy("space_id").
		OrderBy("total", "DESC").
		Limit(0, limit)
	rs, err = db.Query(sql)
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

func (d *Document) GetCountGroupByCreateTime(startTime int64) (documents []map[string]string, err error) {
	/*select month(FROM_UNIXTIME(time)) from table_name group by month(FROM_UNIXTIME(time))*/

	/*select DATE_FORMAT(FROM_UNIXTIME(time),"%Y-%m") from tcm_fund_list group by DATE_FORMAT(FROM_UNIXTIME(time),"%Y-%m")*/
	/*select FROM_UNIXTIME(time,"%Y-%m") from tcm_fund_list group by FROM_UNIXTIME(time,"%Y-%m")*/

	/*select DATE_FORMAT(FROM_UNIXTIME(time),"%Y-%m-%d") from tcm_fund_list group by DATE_FORMAT(FROM_UNIXTIME(time),"%Y-%m-%d")*/
	db := G.DB()
	var rs *mysql.ResultSet
	sql := db.AR().Select("DATE_FORMAT(FROM_UNIXTIME(create_time),'%Y-%m-%d') as date, count('date') as total").
		From(Table_Document_Name).Where(map[string]interface{}{
		"is_delete":      Document_Delete_False,
		"create_time >=": startTime,
	}).GroupBy("DATE_FORMAT(FROM_UNIXTIME(create_time),'%Y-%m-%d')")
	rs, err = db.Query(sql)
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

func (d *Document) GetDocumentGroupCreateUserId() (documents []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	sql := db.AR().Select("create_user_id, count('create_user_id') as total").
		From(Table_Document_Name).Where(map[string]interface{}{
		"is_delete": Document_Delete_False,
	}).GroupBy("create_user_id")
	rs, err = db.Query(sql)
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

func (d *Document) GetDocumentGroupEditUserId() (documents []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	sql := db.AR().Select("edit_user_id, count('edit_user_id') as total").
		From(Table_Document_Name).Where(map[string]interface{}{
		"is_delete": Document_Delete_False,
	}).GroupBy("edit_user_id")
	rs, err = db.Query(sql)
	if err != nil {
		return
	}
	documents = rs.Rows()
	return
}

// 根据文档信息获取文档内容和文件地址
func (d *Document) GetDocumentContentByDocument(doc map[string]string) (content string, pageFile string, err error) {
	// get document page file
	_, pageFile, err = DocumentModel.GetParentDocumentsByDocument(doc)
	if err != nil {
		return content, pageFile, err
	}
	// get document content
	content, err = utils.Document.GetContentByPageFile(pageFile)
	if err != nil {
		return content, pageFile, err
	}
	return content, pageFile, nil
}
