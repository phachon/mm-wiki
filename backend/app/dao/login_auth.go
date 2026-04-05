package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"gorm.io/gorm"
)

const (
	// TableNameLoginAuth 统一登录认证表
	TableNameLoginAuth = "mk_login_auth"
	// LoginAuthPrimaryKey 认证表主键ID
	LoginAuthPrimaryKey = "login_auth_id"
)

// LoginAuth 统一登录认证表数据对象
type LoginAuth struct {
	ctx context.Context
}

// NewLoginAuth 创建统一登录认证表数据对象
func NewLoginAuth(ctx context.Context) *LoginAuth {
	return &LoginAuth{
		ctx: ctx,
	}
}

// Insert 创建认证记录
func (la *LoginAuth) Insert(loginAuthEntity *entity.LoginAuthEntity) errors.BizError {
	loginAuthEntity.CreateTime = time.Now().Unix()
	loginAuthEntity.UpdateTime = time.Now().Unix()
	loginAuthEntity.Status = entity.LoginAuthStatusDefault
	db := GetDB(dbNameMK).WithContext(la.ctx).Table(TableNameLoginAuth).Save(loginAuthEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetLoginAuthByLoginAuthId 根据认证ID获取认证信息
func (la *LoginAuth) GetLoginAuthByLoginAuthId(loginAuthId int64) (loginAuth *entity.LoginAuthEntity, err errors.BizError) {
	loginAuth = &entity.LoginAuthEntity{}
	db := GetDB(dbNameMK).WithContext(la.ctx).Table(TableNameLoginAuth).
		Where(map[string]interface{}{
			LoginAuthPrimaryKey: loginAuthId,
			"status":            entity.LoginAuthStatusDefault,
		}).
		First(&loginAuth)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return loginAuth, nil
}

// GetLoginAuthByName 根据名称获取认证信息
func (la *LoginAuth) GetLoginAuthByName(name string) (loginAuth *entity.LoginAuthEntity, err errors.BizError) {
	loginAuth = &entity.LoginAuthEntity{}
	db := GetDB(dbNameMK).WithContext(la.ctx).Table(TableNameLoginAuth).
		Where("name = ? AND status = ?", name, entity.LoginAuthStatusDefault).
		First(&loginAuth)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return loginAuth, nil
}

// GetUsedLoginAuth 获取正在使用的认证
func (la *LoginAuth) GetUsedLoginAuth() (loginAuth *entity.LoginAuthEntity, err errors.BizError) {
	loginAuth = &entity.LoginAuthEntity{}
	db := GetDB(dbNameMK).WithContext(la.ctx).Table(TableNameLoginAuth).
		Where(map[string]interface{}{
			"is_used": entity.LoginAuthIsUsedYes,
			"status":  entity.LoginAuthStatusDefault,
		}).
		First(&loginAuth)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return loginAuth, nil
}

// GetLoginAuthsByLimit 分页获取认证列表
func (la *LoginAuth) GetLoginAuthsByLimit(limit int, offset int) (loginAuths []*entity.LoginAuthEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(la.ctx).Table(TableNameLoginAuth).
		Where("status", entity.LoginAuthStatusDefault).
		Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", LoginAuthPrimaryKey)).
		Find(&loginAuths)
	if db.Error != nil {
		return loginAuths, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return loginAuths, nil
}

// CountLoginAuths 获取认证总数
func (la *LoginAuth) CountLoginAuths() (count int64, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(la.ctx).Table(TableNameLoginAuth).
		Where("status", entity.LoginAuthStatusDefault).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// GetLoginAuthsByKeywordsAndLimit 根据关键字分页获取认证列表
func (la *LoginAuth) GetLoginAuthsByKeywordsAndLimit(limit int, offset int, keywords *entity.LoginAuthKeywords) (loginAuths []*entity.LoginAuthEntity, err errors.BizError) {
	if keywords == nil {
		return la.GetLoginAuthsByLimit(limit, offset)
	}
	db := GetDB(dbNameMK).WithContext(la.ctx).Table(TableNameLoginAuth)
	db.Where("status", entity.LoginAuthStatusDefault)
	if keywords.Name != "" {
		db = db.Where("name LIKE ?", "%"+keywords.Name+"%")
	}
	db = db.Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", LoginAuthPrimaryKey)).
		Find(&loginAuths)
	if db.Error != nil {
		return loginAuths, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return loginAuths, nil
}

// CountLoginAuthsByKeywords 根据关键字获取认证总数
func (la *LoginAuth) CountLoginAuthsByKeywords(keywords *entity.LoginAuthKeywords) (count int64, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(la.ctx).Table(TableNameLoginAuth)
	db = db.Where("status = ?", entity.LoginAuthStatusDefault)
	if keywords.Name != "" {
		db = db.Where("name LIKE ?", "%"+keywords.Name+"%")
	}
	db.Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// Update 更新认证信息
func (la *LoginAuth) Update(loginAuth entity.LoginAuthEntity) errors.BizError {
	loginAuth.UpdateTime = time.Now().Unix()
	db := GetDB(dbNameMK).WithContext(la.ctx).Table(TableNameLoginAuth).
		Select("Name", "AccountPrefix", "URL", "ExtData", "IsUsed", "UpdateTime").
		Where(map[string]interface{}{
			LoginAuthPrimaryKey: loginAuth.LoginAuthId,
		}).
		Updates(loginAuth)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// UpdateUsed 更新使用状态
func (la *LoginAuth) UpdateUsed(loginAuthId int64, isUsed int) errors.BizError {
	db := GetDB(dbNameMK).WithContext(la.ctx).Table(TableNameLoginAuth).
		Where(map[string]interface{}{
			LoginAuthPrimaryKey: loginAuthId,
			"status":            entity.LoginAuthStatusDefault,
		}).
		Update("is_used", isUsed).
		Update("update_time", time.Now().Unix())
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// ClearUsed 清除所有使用标记
func (la *LoginAuth) ClearUsed() errors.BizError {
	db := GetDB(dbNameMK).WithContext(la.ctx).Table(TableNameLoginAuth).
		Where("status", entity.LoginAuthStatusDefault).
		Update("is_used", entity.LoginAuthIsUsedNo)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// DeleteLoginAuth 删除认证
func (la *LoginAuth) DeleteLoginAuth(loginAuthId int64) errors.BizError {
	db := GetDB(dbNameMK).WithContext(la.ctx).Table(TableNameLoginAuth).
		Where(map[string]interface{}{
			LoginAuthPrimaryKey: loginAuthId,
			"status":            entity.LoginAuthStatusDefault,
		}).
		Update("status", entity.LoginAuthStatusDelete).
		Update("update_time", time.Now().Unix())
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// GetAllLoginAuths 获取所有认证
func (la *LoginAuth) GetAllLoginAuths() (loginAuths []*entity.LoginAuthEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(la.ctx).Table(TableNameLoginAuth).
		Where(map[string]interface{}{
			"status": entity.LoginAuthStatusDefault,
		}).
		Find(&loginAuths)
	if db.Error != nil {
		return loginAuths, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return loginAuths, nil
}
