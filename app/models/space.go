package models

import (
	"fmt"
	"github.com/phachon/mm-wiki/app/utils"
	"github.com/snail007/go-activerecord/mysql"
	"time"
)

const (
	Space_Share_False = 0
	Space_Share_True  = 1

	Space_Download_False = 0
	Space_Download_True  = 1

	Space_Delete_True  = 1
	Space_Delete_False = 0

	Space_Root_Id    = 1
	Space_Admin_Id   = 2
	Space_Default_Id = 3

	Space_VisitLevel_Public  = "public"
	Space_VisitLevel_Private = "private"
)

const Table_Space_Name = "space"

type Space struct {
}

var SpaceModel = Space{}

// get space by space_id
func (s *Space) GetSpaceBySpaceId(spaceId string) (space map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Space_Name).Where(map[string]interface{}{
		"space_id":  spaceId,
		"is_delete": Space_Delete_False,
	}))
	if err != nil {
		return
	}
	space = rs.Row()
	return
}

// space_id and name is exists
func (s *Space) HasSameName(spaceId, name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Space_Name).Where(map[string]interface{}{
		"space_id <>": spaceId,
		"name":        name,
		"is_delete":   Space_Delete_False,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	if rs.Len() > 0 {
		has = true
	}
	return
}

// name is exists
func (s *Space) HasSpaceName(name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Space_Name).Where(map[string]interface{}{
		"name":      name,
		"is_delete": Space_Delete_False,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	if rs.Len() > 0 {
		has = true
	}
	return
}

// get space by name
func (s *Space) GetSpaceByName(name string) (space map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Space_Name).Where(map[string]interface{}{
		"name":      name,
		"is_delete": Space_Delete_False,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	space = rs.Row()
	return
}

// delete space by space_id
func (s *Space) Delete(spaceId string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Update(Table_Space_Name, map[string]interface{}{
		"is_delete":   Space_Delete_True,
		"update_time": time.Now().Unix(),
	}, map[string]interface{}{
		"space_id": spaceId,
	}))
	if err != nil {
		return
	}

	// delete collect space
	go func() {
		CollectionModel.DeleteByResourceIdType(spaceId, fmt.Sprintf("%d", Collection_Type_Space))
	}()

	return
}

// insert space
func (s *Space) Insert(spaceValue map[string]interface{}) (id int64, err error) {

	spaceValue["create_time"] = time.Now().Unix()
	spaceValue["update_time"] = time.Now().Unix()
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Insert(Table_Space_Name, spaceValue))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return id, nil
}

// update space by space_id
func (s *Space) Update(spaceId string, spaceValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	spaceValue["update_time"] = time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_Space_Name, spaceValue, map[string]interface{}{
		"space_id":  spaceId,
		"is_delete": Space_Delete_False,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// update space db and file name by space_id
func (s *Space) UpdateDBAndSpaceFileName(spaceId string, spaceValue map[string]interface{}, oldName string) (id int64, err error) {
	// begin update
	db := G.DB()
	tx, err := db.Begin(db.Config)
	if err != nil {
		return
	}
	var rs *mysql.ResultSet

	// get real old space name (v0.1.2 #53 bug)
	defaultDocument, err := DocumentModel.GetDocumentByParentIdAndSpaceId("0", spaceId, Document_Type_Dir)
	if err != nil {
		return
	}
	if oldName != defaultDocument["name"] {
		oldName = defaultDocument["name"]
	}

	// update space db
	spaceValue["update_time"] = time.Now().Unix()
	rs, err = db.ExecTx(db.AR().Update(Table_Space_Name, spaceValue, map[string]interface{}{
		"space_id":  spaceId,
		"is_delete": Space_Delete_False,
	}), tx)
	if err != nil {
		tx.Rollback()
		return
	}
	id = rs.LastInsertId

	documentValue := map[string]interface{}{
		"name":        spaceValue["name"],
		"update_time": time.Now().Unix(),
	}
	// update space document name
	_, err = db.ExecTx(db.AR().Update(Table_Document_Name, documentValue, map[string]interface{}{
		"space_id":  spaceId,
		"parent_id": 0,
		"type":      Document_Type_Dir,
	}), tx)
	if err != nil {
		tx.Rollback()
		return
	}
	// update space name
	err = utils.Document.UpdateSpaceName(oldName, spaceValue["name"].(string))
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit()

	return
}

// get limit spaces by search keyword
func (s *Space) GetSpacesByKeywordAndLimit(keyword string, limit int, number int) (spaces []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	sql := db.AR().From(Table_Space_Name).Where(map[string]interface{}{
		"is_delete": Space_Delete_False,
	}).WhereWrap(map[string]interface{}{
		"name LIKE": "%" + keyword + "%",
	}, "AND (", "").WhereWrap(map[string]interface{}{
		"description LIKE": "%" + keyword + "%",
	}, "OR", ")").Limit(limit, number).OrderBy("space_id", "DESC")
	rs, err = db.Query(sql)

	if err != nil {
		return
	}
	spaces = rs.Rows()

	return
}

// get limit spaces
func (s *Space) GetSpacesByLimit(limit int, number int) (spaces []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			From(Table_Space_Name).
			Where(map[string]interface{}{
				"is_delete": Space_Delete_False,
			}).
			Limit(limit, number).
			OrderBy("space_id", "DESC"))
	if err != nil {
		return
	}
	spaces = rs.Rows()

	return
}

// get all spaces
func (s *Space) GetSpaces() (spaces []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Space_Name).Where(map[string]interface{}{
			"is_delete": Space_Delete_False,
		}))
	if err != nil {
		return
	}
	spaces = rs.Rows()
	return
}

// get spaces by visitLevel
func (s *Space) GetSpacesByVisitLevel(visitLevel string) (spaces []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Space_Name).Where(map[string]interface{}{
			"visit_level": visitLevel,
			"is_delete":   Space_Delete_False,
		}))
	if err != nil {
		return
	}
	spaces = rs.Rows()
	return
}

// get space count
func (s *Space) CountSpaces() (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			Select("count(*) as total").
			From(Table_Space_Name).
			Where(map[string]interface{}{
				"is_delete": Space_Delete_False,
			}))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get space count
func (s *Space) CountSpacesByTags(tag string) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			Select("count(*) as total").
			From(Table_Space_Name).
			Where(map[string]interface{}{
				"tags LIKE": "%" + tag + "%",
				"is_delete": Space_Delete_False,
			}))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get space count by keyword
func (s *Space) CountSpacesByKeyword(keyword string) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	sql := db.AR().Select("count(*) as total").From(Table_Space_Name).
		Where(map[string]interface{}{"is_delete": Space_Delete_False}).
		WhereWrap(map[string]interface{}{"name LIKE": "%" + keyword + "%"}, "AND (", "").
		WhereWrap(map[string]interface{}{"description LIKE": "%" + keyword + "%"}, "OR", ")")
	rs, err = db.Query(sql)
	if err != nil {
		return
	}
	count = utils.Convert.StringToInt64(rs.Value("total"))
	return
}

// get space count by tags
func (s *Space) GetSpacesByTags(tag string) (spaces []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Space_Name).Where(map[string]interface{}{
		"tags LIKE": "%" + tag + "%",
		"is_delete": Space_Delete_False,
	}).OrderBy("space_id", "DESC"))
	if err != nil {
		return
	}
	spaces = rs.Rows()

	return
}

// get space by name
func (s *Space) GetSpaceByLikeName(name string) (spaces []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Space_Name).Where(map[string]interface{}{
		"name Like": "%" + name + "%",
		"is_delete": Space_Delete_False,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	spaces = rs.Rows()
	return
}

// get space by many space_id
func (s *Space) GetSpaceBySpaceIds(spaceIds []string) (spaces []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Space_Name).Where(map[string]interface{}{
		"space_id":  spaceIds,
		"is_delete": Space_Delete_False,
	}))
	if err != nil {
		return
	}
	spaces = rs.Rows()
	return
}

// update space by name
func (s *Space) UpdateSpaceByName(space map[string]interface{}) (affect int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	space["update_time"] = time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_Space_Name, space, map[string]interface{}{
		"name": space["name"],
	}))
	if err != nil {
		return
	}
	affect = rs.RowsAffected
	return
}
