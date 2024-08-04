package dao

import (
	"context"
	"time"

	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/utils"
	"gorm.io/gorm"
)

const (
	// TableNameAccountRole 账号角色关系表
	TableNameAccountRole = "mk_account_role"
	// AccountRolePrimaryKey 账号角色关系表主键ID
	AccountRolePrimaryKey = "account_role_id"
)

// AccountRole 账号角色关系表数据
type AccountRole struct {
	ctx context.Context
	db  *gorm.DB
}

// NewAccountRole 创建系统账号表数据对象
func NewAccountRole(ctx context.Context) *AccountRole {
	return &AccountRole{
		ctx: ctx,
		db:  GetDB(dbNameMK).WithContext(ctx),
	}
}

// Insert 创建账号插入一条账号记录
func (ar *AccountRole) Insert(accountRoleEntity *entity.AccountRoleEntity) errors.BizError {
	accountRoleEntity.CreateTime = utils.NewJsonTime(time.Now())
	db := ar.db.Table(TableNameAccountRole).
		Save(accountRoleEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// DeleteByRoleIdAccountId 通过角色ID和账号ID删除
func (ar *AccountRole) DeleteByRoleIdAccountId(roleId int64, accountId int64) errors.BizError {
	db := ar.db.
		Table(TableNameAccountRole).
		Where("role_id", roleId).
		Where("account_id", accountId).
		Delete(entity.AccountRoleEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}

// DeleteByAccountId 通过账号ID删除
func (ar *AccountRole) DeleteByAccountId(accountId int64) errors.BizError {
	db := ar.db.
		Table(TableNameAccountRole).
		Where("account_id", accountId).
		Delete(entity.AccountRoleEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}

// DeleteByRoleId 通过角色ID删除
func (ar *AccountRole) DeleteByRoleId(roleId int64) errors.BizError {
	db := ar.db.
		Table(TableNameAccountRole).
		Where("role_id", roleId).
		Delete(entity.AccountRoleEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}

// GetAccountRolesByAccountId 根据账号ID获取角色账号关系
func (ar *AccountRole) GetAccountRolesByAccountId(accountId int64) (accountRoles []*entity.AccountRoleEntity,
	err errors.BizError) {
	db := ar.db.Table(TableNameAccountRole).
		Where("account_id = ?", accountId).
		Find(&accountRoles)
	if db.Error != nil {
		return accountRoles, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return accountRoles, nil
}

// GetAccountRolesByAccountIds 根据多个账号ID获取角色账号关系
func (ar *AccountRole) GetAccountRolesByAccountIds(accountIds []int64) (accountRoles []*entity.AccountRoleEntity,
	err errors.BizError) {
	db := ar.db.Table(TableNameAccountRole).
		Where("account_id IN (?)", accountIds).
		Find(&accountRoles)
	if db.Error != nil {
		return accountRoles, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return accountRoles, nil
}

// GetRoleAccountsByRoleId 根据角色ID查找角色账号关系
func (ar *AccountRole) GetRoleAccountsByRoleId(roleId int64) (accountRoles []*entity.AccountRoleEntity,
	err errors.BizError) {
	db := ar.db.Table(TableNameAccountRole).
		Where("role_id = ?", roleId).
		Find(&accountRoles)
	if db.Error != nil {
		return accountRoles, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return accountRoles, nil
}

// GetRoleAccountsByRoleIdLimit 根据分页获取账号
func (ar *AccountRole) GetRoleAccountsByRoleIdLimit(limit int, offset int, roleId int64) (accountRoles []*entity.AccountRoleEntity,
	err errors.BizError) {

	db := ar.db.Table(TableNameAccountRole).
		Where("role_id = ?", roleId).
		Limit(limit).
		Offset(offset).
		Find(&accountRoles)
	if db.Error != nil {
		return accountRoles, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return accountRoles, nil
}

// CountRoleAccountsByRoleId 根据角色ID获取账号总数
func (ar *AccountRole) CountRoleAccountsByRoleId(roleId int64) (count int64, err errors.BizError) {
	db := ar.db.Table(TableNameAccountRole).
		Where("role_id = ?", roleId).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}
