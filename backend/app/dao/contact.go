package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/utils"
	"gorm.io/gorm"
)

const (
	// TableNameContact 联系人表
	TableNameContact = "mk_contact"
	// ContactPrimaryKey 联系人表主键ID
	ContactPrimaryKey = "contact_id"
)

// Contact 联系人表数据对象
type Contact struct {
	ctx context.Context
}

// NewContact 创建联系人表数据对象
func NewContact(ctx context.Context) *Contact {
	return &Contact{
		ctx: ctx,
	}
}

// Insert 创建联系人插入一条联系人记录
func (r *Contact) Insert(contactEntity *entity.ContactEntity) errors.BizError {
	contactEntity.CreateTime = utils.NewJsonTime(time.Now())
	contactEntity.UpdateTime = utils.NewJsonTime(time.Now())
	contactEntity.Status = entity.ContactStatusDefault
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameContact).Save(contactEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetContactByContactId 根据联系人ID获取联系人信息
func (r *Contact) GetContactByContactId(contactId int64) (contact *entity.ContactEntity, err errors.BizError) {

	contact = &entity.ContactEntity{}
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameContact).
		Where(map[string]interface{}{
			ContactPrimaryKey: contactId,
			"status":          entity.ContactStatusDefault,
		}).
		First(&contact)

	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return contact, nil
}

// GetContactsByLimit 分页获取联系人列表
func (r *Contact) GetContactsByLimit(limit int, offset int) (contacts []*entity.ContactEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameContact).
		Where("status", entity.ContactStatusDefault).
		Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", ContactPrimaryKey)).
		Find(&contacts)
	if db.Error != nil {
		return contacts, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return contacts, nil
}

// CountContacts 获取联系人总数
func (r *Contact) CountContacts() (count int64, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameContact).
		Where("status", entity.ContactStatusDefault).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// GetContactsByKeywordsAndLimit 根据关键字分页获取联系人列表
func (r *Contact) GetContactsByKeywordsAndLimit(limit int, offset int, keywords *entity.ContactKeywords) (contacts []*entity.ContactEntity, err errors.BizError) {

	if keywords == nil {
		return r.GetContactsByLimit(limit, offset)
	}

	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameContact)
	db.Where("status", entity.ContactStatusDefault)
	if keywords.Name != "" {
		db = db.Where("name LIKE ?", "%"+keywords.Name+"%")
	}
	db = db.Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", ContactPrimaryKey)).
		Find(&contacts)
	if db.Error != nil {
		return contacts, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return contacts, nil
}

// CountContactsByKeywords 根据关键字获取联系人总数
func (r *Contact) CountContactsByKeywords(keywords *entity.ContactKeywords) (count int64, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameContact)
	db = db.Where("status = ?", entity.ContactStatusDefault)
	if keywords.Name != "" {
		db = db.Where("name LIKE ?", "%"+keywords.Name+"%")
	}
	db.Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// Update 更新联系人只会更新如下字段
func (r *Contact) Update(contact entity.ContactEntity) errors.BizError {
	contact.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameContact).
		Select("Name", "Mobile", "Email", "Position", "UpdateTime").
		Where(map[string]interface{}{
			ContactPrimaryKey: contact.ContactId,
		}).
		Updates(contact)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// DeleteContact 删除联系人
func (r *Contact) DeleteContact(contactId int64) errors.BizError {
	updateTime := utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameContact).
		Where(map[string]interface{}{
			ContactPrimaryKey: contactId,
			"status":          entity.ContactStatusDefault,
		}).
		Update("status", entity.ContactStatusDelete).
		Update("update_time", updateTime)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// GetAllContacts 获取所有的联系人
func (r *Contact) GetAllContacts() (contacts []*entity.ContactEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameContact).
		Where(map[string]interface{}{
			"status": entity.ContactStatusDefault,
		}).
		Find(&contacts)
	if db.Error != nil {
		return contacts, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return contacts, nil
}
