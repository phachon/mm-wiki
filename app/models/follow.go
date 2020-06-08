package models

import (
	"github.com/phachon/mm-wiki/app/utils"
	"github.com/snail007/go-activerecord/mysql"
	"time"
)

const (
	Follow_Type_Doc  = 1
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
		"follow_id": followId,
	}))
	if err != nil {
		return
	}
	follow = rs.Row()
	return
}

// get follows by user_id and type
func (f *Follow) GetFollowsByUserIdAndType(userId string, followType int) (follows []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Follow_Name).Where(map[string]interface{}{
		"user_id": userId,
		"type":    followType,
	}))
	if err != nil {
		return
	}
	follows = rs.Rows()
	return
}

// get follows by user_id and type limit
func (f *Follow) GetFollowsByUserIdTypeAndLimit(userId string, followType int, limit int, number int) (follows []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Follow_Name).Where(map[string]interface{}{
		"user_id": userId,
		"type":    followType,
	}).Limit(limit, number).OrderBy("follow_id", "DESC"))
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
		"object_id": objectId,
		"type":      followType,
	}))
	if err != nil {
		return
	}
	follows = rs.Rows()
	return
}

// get followed follow
func (f *Follow) GetFollowByUserIdAndTypeAndObjectId(userId string, followType int, objectId string) (follow map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Follow_Name).Where(map[string]interface{}{
		"user_id":   userId,
		"type":      followType,
		"object_id": objectId,
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

// delete follow by type and object_id
func (f *Follow) DeleteByObjectIdType(objectId string, followType string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Delete(Table_Follow_Name, map[string]interface{}{
		"type":      followType,
		"object_id": objectId,
	}))
	if err != nil {
		return
	}
	return
}

// insert follow user
func (f *Follow) Insert(userId string, fType int, objectId string) (id int64, err error) {
	followValue := map[string]interface{}{
		"user_id":     userId,
		"type":        fType,
		"object_id":   objectId,
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

// get follow count
func (f *Follow) CountFollowsByUserIdAndType(userId string, followType int) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().Select("count(*) as total").From(Table_Follow_Name).Where(map[string]interface{}{
		"user_id": userId,
		"type":    followType,
	}))
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
		"follow_id": followIds,
	}))
	if err != nil {
		return
	}
	follows = rs.Rows()
	return
}

// create auto follow document
func (f *Follow) CreateAutoFollowDocument(userId string, documentId string) (id int64, err error) {

	autoFollowConf := ConfigModel.GetConfigValueByKey(ConfigKeyAutoFollowdoc, "0")
	if autoFollowConf == "1" {
		follow, err := f.GetFollowByUserIdAndTypeAndObjectId(userId, Follow_Type_Doc, documentId)
		if err != nil {
			return 0, err
		}
		if len(follow) == 0 {
			return f.Insert(userId, Follow_Type_Doc, documentId)
		}
	}
	return 0, nil
}

func (f *Follow) FollowDocument(userId string, documentId string) (id int64, err error) {

	follow, err := f.GetFollowByUserIdAndTypeAndObjectId(userId, Follow_Type_Doc, documentId)
	if err != nil {
		return 0, err
	}
	if len(follow) == 0 {
		return f.Insert(userId, Follow_Type_Doc, documentId)
	}
	return 0, nil
}

func (f *Follow) GetFollowGroupUserId(fType int) (collects []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	sql := db.AR().Select("user_id, count('user_id') as total").
		From(Table_Follow_Name).Where(map[string]interface{}{
		"type": fType,
	}).GroupBy("user_id")
	rs, err = db.Query(sql)
	if err != nil {
		return
	}
	collects = rs.Rows()
	return
}

func (f *Follow) GetFansUserGroupUserId() (collects []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	sql := db.AR().Select("object_id, count('object_id') as total").
		From(Table_Follow_Name).Where(map[string]interface{}{
		"type": Follow_Type_User,
	}).GroupBy("object_id")
	rs, err = db.Query(sql)
	if err != nil {
		return
	}
	collects = rs.Rows()
	return
}
