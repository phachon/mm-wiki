package models

import (
	"mm-wiki/app/utils"
	"github.com/snail007/go-activerecord/mysql"
	"time"
)

const (
	Follow_Type_Doc = 1
	Follow_Type_User = 2
)

const Table_Follow_Name = "follow"

type Follow struct {

}

var FollowModel = Follow{}

// get follow by follow_id
func (f *Follow) GetFollowByFollowId(followId string) (follow map[string]string, err error) {
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
func (f *Follow) GetFollowsByUserIdAndType(userId string, followType int) (follows []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Follow_Name).Where(map[string]interface{}{
		"user_id":  userId,
		"type":  followType,
	}))
	if err != nil {
		return
	}
	follows = rs.Rows()
	return
}

// get followed follows
func (f *Follow) GetFollowsByObjectIdAndType(objectId string, followType int) (follows []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Follow_Name).Where(map[string]interface{}{
		"object_id":  objectId,
		"type":  followType,
	}))
	if err != nil {
		return
	}
	follows = rs.Rows()
	return
}

// get followed follow
func (f *Follow) GetFollowsByUserIdAndTypeAndObjectId(userId string, followType int, objectId string) (follow map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Follow_Name).Where(map[string]interface{}{
		"user_id": userId,
		"type":  followType,
		"object_id":  objectId,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	follow = rs.Row()
	return
}

// delete follow by follow_id
func (f *Follow) Delete(followId string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Delete(Table_Follow_Name, map[string]interface{}{
		"follow_id": followId,
	}))
	if err != nil {
		return
	}
	return
}

// insert follow user
func (f *Follow) Insert(userId string, fType int, objectId string) (id int64, err error) {
	followValue := map[string]interface{}{
		"user_id": userId,
		"type": fType,
		"object_id": objectId,
		"create_time": time.Now().Unix(),
	}
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
func (f *Follow) GetFollows() (follows []map[string]string, err error) {

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
func (f *Follow) CountFollows() (count int64, err error) {

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
func (f *Follow) GetFollowsByFollowIds(followIds []string) (follows []map[string]string, err error) {
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