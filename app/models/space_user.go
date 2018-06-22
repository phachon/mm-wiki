package models

import (
	"mm-wiki/app/utils"
	"github.com/snail007/go-activerecord/mysql"
	"time"
)

const Table_SpaceUser_Name = "space_user"

type SpaceUser struct {

}

var SpaceUserModel = SpaceUser{}

// get spaceUser by space_user_id
func (s *SpaceUser) GetSpaceUserBySpaceUserId(spaceUserId string) (spaceUser map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_SpaceUser_Name).Where(map[string]interface{}{
		"space_user_id":   spaceUserId,
	}))
	if err != nil {
		return
	}
	spaceUser = rs.Row()
	return
}

// get spaceUser by name
func (s *SpaceUser) GetSpaceUsersByUserId(userId string) (spaceUsers []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_SpaceUser_Name).Where(map[string]interface{}{
		"user_id":  userId,
	}))
	if err != nil {
		return
	}
	spaceUsers = rs.Rows()
	return
}

// get spaceUser by name
func (s *SpaceUser) GetSpaceUsersBySpaceId(spaceId string) (spaceUsers []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_SpaceUser_Name).Where(map[string]interface{}{
		"space_id":  spaceId,
	}))
	if err != nil {
		return
	}
	spaceUsers = rs.Rows()
	return
}

// delete spaceUser by space_user_id
func (s *SpaceUser) Delete(spaceUserId string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Delete(Table_SpaceUser_Name, map[string]interface{}{
		"space_user_id": spaceUserId,
	}))
	if err != nil {
		return
	}
	return
}

// delete spaceUser by space_user_id
func (s *SpaceUser) DeleteBySpaceId(spaceId string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Delete(Table_SpaceUser_Name, map[string]interface{}{
		"space_id": spaceId,
	}))
	if err != nil {
		return
	}
	return
}

// delete spaceUser by space_user_id
func (s *SpaceUser) DeleteByUserId(userId string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Delete(Table_SpaceUser_Name, map[string]interface{}{
		"user_id": userId,
	}))
	if err != nil {
		return
	}
	return
}

// insert spaceUser
func (s *SpaceUser) Insert(spaceUserValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Insert(Table_SpaceUser_Name, spaceUserValue))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// update spaceUser by spaceUser_id
func (s *SpaceUser) Update(spaceUserId string, spaceUserValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	spaceUserValue["update_time"] =  time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_SpaceUser_Name, spaceUserValue, map[string]interface{}{
		"space_user_id":   spaceUserId,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// get limit spaceUsers by spaceId
func (s *SpaceUser) GetSpaceUsersBySpaceIdAndLimit(spaceId string, limit int, number int) (spaceUsers []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_SpaceUser_Name).Where(map[string]interface{}{
		"space_id": spaceId,
	}).Limit(limit, number))
	if err != nil {
		return
	}
	spaceUsers = rs.Rows()

	return
}

// get all spaceUsers
func (s *SpaceUser) GetSpaceUsers() (spaceUsers []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_SpaceUser_Name))
	if err != nil {
		return
	}
	spaceUsers = rs.Rows()
	return
}

// get spaceUser count by keyword
func (s *SpaceUser) CountSpaceUsersBySpaceId(spaceId string) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().
		Select("count(*) as total").
		From(Table_SpaceUser_Name).
		Where(map[string]interface{}{
			"space_id": spaceId,
		}))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get spaceUser by many space_user_id
func (s *SpaceUser) GetSpaceUsersBySpaceUserIds(spaceUserIds []string) (spaceUsers []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_SpaceUser_Name).Where(map[string]interface{}{
		"space_user_id":   spaceUserIds,
	}))
	if err != nil {
		return
	}
	spaceUsers = rs.Rows()
	return
}