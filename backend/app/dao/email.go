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
	// TableNameEmail 邮箱服务器表
	TableNameEmail = "mk_email"
	// EmailPrimaryKey 邮箱表主键ID
	EmailPrimaryKey = "email_id"
)

// Email 邮箱服务器表数据对象
type Email struct {
	ctx context.Context
}

// NewEmail 创建邮箱服务器表数据对象
func NewEmail(ctx context.Context) *Email {
	return &Email{
		ctx: ctx,
	}
}

// Insert 创建邮箱插入一条邮箱记录
func (r *Email) Insert(emailEntity *entity.EmailEntity) errors.BizError {
	emailEntity.CreateTime = utils.NewJsonTime(time.Now())
	emailEntity.UpdateTime = utils.NewJsonTime(time.Now())
	emailEntity.Status = entity.EmailStatusDefault
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameEmail).Save(emailEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetEmailByEmailId 根据邮箱ID获取邮箱信息
func (r *Email) GetEmailByEmailId(emailId int64) (email *entity.EmailEntity, err errors.BizError) {

	email = &entity.EmailEntity{}
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameEmail).
		Where(map[string]interface{}{
			EmailPrimaryKey: emailId,
			"status":        entity.EmailStatusDefault,
		}).
		First(&email)

	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return email, nil
}

// GetEmailByName 根据邮箱名称查找正常的邮箱
func (r *Email) GetEmailByName(name string) (email *entity.EmailEntity, err errors.BizError) {

	email = &entity.EmailEntity{}
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameEmail).
		Where(map[string]interface{}{
			"name":   name,
			"status": entity.EmailStatusDefault,
		}).
		First(&email)

	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return email, nil
}

// GetEmailsByLimit 分页获取邮箱列表
func (r *Email) GetEmailsByLimit(limit int, offset int) (emails []*entity.EmailEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameEmail).
		Where("status", entity.EmailStatusDefault).
		Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", EmailPrimaryKey)).
		Find(&emails)
	if db.Error != nil {
		return emails, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return emails, nil
}

// CountEmails 获取邮箱总数
func (r *Email) CountEmails() (count int64, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameEmail).
		Where("status", entity.EmailStatusDefault).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// GetEmailsByKeywordsAndLimit 根据关键字分页获取邮箱列表
func (r *Email) GetEmailsByKeywordsAndLimit(limit int, offset int, keywords *entity.EmailKeywords) (emails []*entity.EmailEntity, err errors.BizError) {

	if keywords == nil {
		return r.GetEmailsByLimit(limit, offset)
	}

	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameEmail)
	db.Where("status", entity.EmailStatusDefault)
	if keywords.Name != "" {
		db = db.Where("name LIKE ?", "%"+keywords.Name+"%")
	}
	db = db.Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", EmailPrimaryKey)).
		Find(&emails)
	if db.Error != nil {
		return emails, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return emails, nil
}

// CountEmailsByKeywords 根据关键字获取邮箱总数
func (r *Email) CountEmailsByKeywords(keywords *entity.EmailKeywords) (count int64, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameEmail)
	db = db.Where("status = ?", entity.EmailStatusDefault)
	if keywords.Name != "" {
		db = db.Where("name LIKE ?", "%"+keywords.Name+"%")
	}
	db.Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// Update 更新邮箱只会更新如下字段
func (r *Email) Update(email entity.EmailEntity) errors.BizError {
	email.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameEmail).
		Select("Name", "SenderAddress", "SenderName", "SenderTitlePrefix", "Host", "Port", "Username", "Password", "IsSSL", "IsUsed", "UpdateTime").
		Where(map[string]interface{}{
			EmailPrimaryKey: email.EmailId,
		}).
		Updates(email)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// DeleteEmail 删除邮箱
func (r *Email) DeleteEmail(emailId int64) errors.BizError {
	updateTime := utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameEmail).
		Where(map[string]interface{}{
			EmailPrimaryKey: emailId,
			"status":        entity.EmailStatusDefault,
		}).
		Update("status", entity.EmailStatusDelete).
		Update("update_time", updateTime)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// GetUsedEmail 获取正在使用的邮箱
func (r *Email) GetUsedEmail() (email *entity.EmailEntity, err errors.BizError) {

	email = &entity.EmailEntity{}
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameEmail).
		Where(map[string]interface{}{
			"is_used": entity.EmailUsedTrue,
			"status":  entity.EmailStatusDefault,
		}).
		First(&email)

	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return email, nil
}

// UpdateAllEmailUnused 设置所有邮箱为未使用
func (r *Email) UpdateAllEmailUnused() errors.BizError {
	updateTime := utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameEmail).
		Where("status = ?", entity.EmailStatusDefault).
		Update("is_used", entity.EmailUsedFalse).
		Update("update_time", updateTime)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}
