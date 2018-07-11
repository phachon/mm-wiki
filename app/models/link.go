package models

import (
	"mm-wiki/app/utils"
	"github.com/snail007/go-activerecord/mysql"
	"time"
)

const Table_Link_Name = "link"

type Link struct {

}

var LinkModel = Link{}

// get link by link_id
func (u *Link) GetLinkByLinkId(linkId string) (link map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Link_Name).Where(map[string]interface{}{
		"link_id":   linkId,
	}))
	if err != nil {
		return
	}
	link = rs.Row()
	return
}

// link_id and name is exists
func (u *Link) HasSameName(linkId, name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Link_Name).Where(map[string]interface{}{
		"link_id <>": linkId,
		"name":   name,
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
func (u *Link) HasLinkName(name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Link_Name).Where(map[string]interface{}{
		"name":  name,
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
func (u *Link) GetLinkByName(name string) (link map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Link_Name).Where(map[string]interface{}{
		"name":  name,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	link = rs.Row()
	return
}

// delete link by link_id
func (u *Link) Delete(linkId string) (err error) {
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
func (u *Link) Insert(linkValue map[string]interface{}) (id int64, err error) {

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
func (u *Link) Update(linkId string, linkValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	linkValue["update_time"] =  time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_Link_Name, linkValue, map[string]interface{}{
		"link_id":   linkId,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// get limit links by search keyword
func (u *Link) GetLinksByKeywordAndLimit(keyword string, limit int, number int) (links []map[string]string, err error) {

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
func (u *Link) GetLinksByLimit(limit int, number int) (links []map[string]string, err error) {

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
func (u *Link) GetLinks() (links []map[string]string, err error) {

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

// get link count
func (u *Link) CountLinks() (count int64, err error) {

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
func (u *Link) CountLinksByKeyword(keyword string) (count int64, err error) {

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
func (u *Link) GetLinksByLikeName(name string) (links []map[string]string, err error) {
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
func (u *Link) GetLinkByLinkIds(linkIds []string) (links []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Link_Name).Where(map[string]interface{}{
		"link_id":   linkIds,
	}))
	if err != nil {
		return
	}
	links = rs.Rows()
	return
}
