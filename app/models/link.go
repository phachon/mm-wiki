package models

import (
	"github.com/phachon/mm-wiki/app/utils"
	"github.com/snail007/go-activerecord/mysql"
	"time"
)

const Table_Link_Name = "link"

type Link struct {
}

var LinkModel = Link{}

// get link by link_id
func (l *Link) GetLinkByLinkId(linkId string) (link map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Link_Name).Where(map[string]interface{}{
		"link_id": linkId,
	}))
	if err != nil {
		return
	}
	link = rs.Row()
	return
}

// link_id and name is exists
func (l *Link) HasSameName(linkId, name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Link_Name).Where(map[string]interface{}{
		"link_id <>": linkId,
		"name":       name,
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
func (l *Link) HasLinkName(name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Link_Name).Where(map[string]interface{}{
		"name": name,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	if rs.Len() > 0 {
		has = true
	}
	return
}

// get link by name
func (l *Link) GetLinkByName(name string) (link map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Link_Name).Where(map[string]interface{}{
		"name": name,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	link = rs.Row()
	return
}

// delete link by link_id
func (l *Link) Delete(linkId string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Delete(Table_Link_Name, map[string]interface{}{
		"link_id": linkId,
	}))
	if err != nil {
		return
	}
	return
}

// insert link
func (l *Link) Insert(linkValue map[string]interface{}) (id int64, err error) {

	linkValue["create_time"] = time.Now().Unix()
	linkValue["update_time"] = time.Now().Unix()
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Insert(Table_Link_Name, linkValue))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// update link by link_id
func (l *Link) Update(linkId string, linkValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	linkValue["update_time"] = time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_Link_Name, linkValue, map[string]interface{}{
		"link_id": linkId,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// get limit links by search keyword
func (l *Link) GetLinksByKeywordAndLimit(keyword string, limit int, number int) (links []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Link_Name).Where(map[string]interface{}{
		"name LIKE": "%" + keyword + "%",
	}).Limit(limit, number).OrderBy("link_id", "DESC"))
	if err != nil {
		return
	}
	links = rs.Rows()

	return
}

// get limit links
func (l *Link) GetLinksByLimit(limit int, number int) (links []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			From(Table_Link_Name).
			Limit(limit, number).
			OrderBy("link_id", "DESC"))
	if err != nil {
		return
	}
	links = rs.Rows()

	return
}

// get all links
func (l *Link) GetLinks() (links []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Link_Name))
	if err != nil {
		return
	}
	links = rs.Rows()
	return
}

// get all links by sequence
func (l *Link) GetLinksOrderBySequence() (links []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Link_Name).OrderBy("sequence", "ASC"))
	if err != nil {
		return
	}
	links = rs.Rows()
	return
}

// get link count
func (l *Link) CountLinks() (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			Select("count(*) as total").
			From(Table_Link_Name))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get link count by keyword
func (l *Link) CountLinksByKeyword(keyword string) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().
		Select("count(*) as total").
		From(Table_Link_Name).
		Where(map[string]interface{}{
			"name LIKE": "%" + keyword + "%",
		}))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get links by like name
func (l *Link) GetLinksByLikeName(name string) (links []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Link_Name).Where(map[string]interface{}{
		"name Like": "%" + name + "%",
	}).Limit(0, 1))
	if err != nil {
		return
	}
	links = rs.Rows()
	return
}

// get link by many link_id
func (l *Link) GetLinkByLinkIds(linkIds []string) (links []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Link_Name).Where(map[string]interface{}{
		"link_id": linkIds,
	}))
	if err != nil {
		return
	}
	links = rs.Rows()
	return
}
