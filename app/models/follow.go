package models

import (
	"mm-wiki/app/utils"
	"github.com/snail007/go-activerecord/mysql"
)

const Table_Follow_Name = "follow"

type Follow struct {

}

var FollowModel = Follow{}

// get follow by follow_id
func (c *Follow) GetFollowByFollowId(followId string) (follow map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Follow_Name).Where(map[string]interface{}{
		"follow_id":   followId,
	}))
	if err != nil {
		return
	}
	follow = rs.Row()
	return
}

// get follows by user_id
func (c *Follow) GetFollowsByUserId(userId string) (follows []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Follow_Name).Where(map[string]interface{}{
		"user_id":  userId,
	}))
	if err != nil {
		return
	}
	follows = rs.Rows()
	return
}

// delete follow by follow_id
func (c *Follow) Delete(followId string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Delete(Table_Follow_Name, map[string]interface{}{
		"follow_id": followId,
	}))
	if err != nil {
		return
	}
	return
}

// insert follow
func (c *Follow) Insert(followValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Insert(Table_Follow_Name, followValue))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// get all follows
func (c *Follow) GetFollows() (follows []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Follow_Name))
	if err != nil {
		return
	}
	follows = rs.Rows()
	return
}

// get follow count
func (c *Follow) CountFollows() (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			Select("count(*) as total").
			From(Table_Follow_Name))
	if err != nil {
		return
	}
	count = utils.Convert.StringToInt64(rs.Value("total"))
	return
}

// get follows by many follow_id
func (c *Follow) GetFollowsByFollowIds(followIds []string) (follows []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Follow_Name).Where(map[string]interface{}{
		"follow_id":   followIds,
	}))
	if err != nil {
		return
	}
	follows = rs.Rows()
	return
}