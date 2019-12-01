package models

import (
	"github.com/phachon/mm-wiki/app/utils"
	"github.com/snail007/go-activerecord/mysql"
)

const Table_Contact_Name = "contact"

type Contact struct {
}

var ContactModel = Contact{}

// 分页获取联系人
func (c *Contact) GetContactByLimit(limit, number int) (contact []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Contact_Name).Limit(limit, number))
	if err != nil {
		return
	}

	contact = rs.Rows()
	return
}

// 获取联系人总数
func (c *Contact) CountContact() (count int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().Select("count(*) as total").From(Table_Contact_Name))
	if err != nil {
		return
	}

	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// 分页筛选名字查询结果
func (c *Contact) GetContactByLimitAndName(name string, limit, number int) (contacts []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Contact_Name).Where(map[string]interface{}{
		"name Like": "%" + name + "%",
	}).Limit(limit, number))
	if err != nil {
		return
	}

	contacts = rs.Rows()
	return
}

// 获取筛选名字查询结果条数
func (c *Contact) CountContactByName(name string) (count int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Contact_Name).Select("count(*) as total").Where(map[string]interface{}{
		"name Like": "%" + name + "%",
	}))
	if err != nil {
		return
	}

	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// 通过 contact_id 获取联系人数据
func (c *Contact) GetContactByContactId(contactId string) (contact map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Contact_Name).Where(map[string]interface{}{
		"contact_id": contactId,
	}))
	if err != nil {
		return
	}

	contact = rs.Row()
	return
}

// 通过 contact_id 更新联系人信息
func (c *Contact) UpdateByContactId(contact map[string]interface{}, contactId string) (affect int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Update(Table_Contact_Name, contact, map[string]interface{}{
		"contact_id": contactId,
	}))
	if err != nil {
		return
	}

	affect = rs.RowsAffected
	return
}

// 通过 contact_id 更新联系人信息
func (c *Contact) Insert(contact map[string]interface{}) (contactId int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Insert(Table_Contact_Name, contact))
	if err != nil {
		return
	}

	contactId = rs.LastInsertId
	return
}

// 通过 contact_id 删除联系人信息
func (c *Contact) DeleteByContactId(contactId string) (affect int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Delete(Table_Contact_Name, map[string]interface{}{
		"contact_id": contactId,
	}))
	if err != nil {
		return
	}

	affect = rs.RowsAffected
	return
}

// 获取所有联系人信息
func (c *Contact) GetAllContact() (contacts []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Contact_Name))
	if err != nil {
		return
	}

	contacts = rs.Rows()
	return
}
