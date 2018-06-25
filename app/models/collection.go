package models

import (
	"mm-wiki/app/utils"
	"github.com/snail007/go-activerecord/mysql"
)

const (
	Collection_Type_Page = 1
	Collection_Type_Space = 2
)

const Table_Collection_Name = "collection"

type Collection struct {

}

var CollectionModel = Collection{}

// get collection by collection_id
func (c *Collection) GetCollectionByCollectionId(collectionId string) (collection map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Collection_Name).Where(map[string]interface{}{
		"collection_id":   collectionId,
	}))
	if err != nil {
		return
	}
	collection = rs.Row()
	return
}

// get collections by user_id
func (c *Collection) GetCollectionsByUserId(userId string) (collections []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Collection_Name).Where(map[string]interface{}{
		"user_id":  userId,
	}))
	if err != nil {
		return
	}
	collections = rs.Rows()
	return
}

// get collections by user_id
func (c *Collection) GetCollectionsByUserIdAndType(userId string, typeS int) (collections []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Collection_Name).Where(map[string]interface{}{
		"user_id":  userId,
		"type":  typeS,
	}))
	if err != nil {
		return
	}
	collections = rs.Rows()
	return
}

// delete collection by collection_id
func (c *Collection) Delete(collectionId string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Delete(Table_Collection_Name, map[string]interface{}{
		"collection_id": collectionId,
	}))
	if err != nil {
		return
	}
	return
}

// insert collection
func (c *Collection) Insert(collectionValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Insert(Table_Collection_Name, collectionValue))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// get all collections
func (c *Collection) GetCollections() (collections []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Collection_Name))
	if err != nil {
		return
	}
	collections = rs.Rows()
	return
}

// get collection count
func (c *Collection) CountCollections() (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			Select("count(*) as total").
			From(Table_Collection_Name))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get collections by many collection_id
func (c *Collection) GetCollectionsByCollectionIds(collectionIds []string) (collections []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Collection_Name).Where(map[string]interface{}{
		"collection_id":   collectionIds,
	}))
	if err != nil {
		return
	}
	collections = rs.Rows()
	return
}