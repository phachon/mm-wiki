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
	// TableNameAccount 系统账号表
	TableNameAccount = "hms_account"
	// AccountPrimaryKey 账号表主键ID
	AccountPrimaryKey = "account_id"
)

// Account 系统账号表数据
type Account struct {
	ctx context.Context
}

// NewAccount 创建系统账号表数据对象
func NewAccount(ctx context.Context) *Account {
	return &Account{
		ctx: ctx,
	}
}

// Insert 创建账号插入一条账号记录
func (ka *Account) Insert(accountEntity *entity.AccountEntity) errors.BizError {
	accountEntity.CreateTime = utils.NewJsonTime(time.Now())
	accountEntity.UpdateTime = utils.NewJsonTime(time.Now())
	accountEntity.Status = entity.AccountStatusDefault
	db := GetDB(dbNameKms).WithContext(ka.ctx).Table(TableNameAccount).Save(accountEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetAccountByName 根据账号名查找正常的账号
func (ka *Account) GetAccountByName(accountName string) (account *entity.AccountEntity, err errors.BizError) {
	account = &entity.AccountEntity{}
	db := GetDB(dbNameKms).WithContext(ka.ctx).Table(TableNameAccount).
		Where(map[string]interface{}{
			"name": accountName,
		}).
		First(&account)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return account, nil
}

// GetAccountByAccountId 根据账号ID获取账号信息
func (ka *Account) GetAccountByAccountId(accountId int64) (account *entity.AccountEntity, err errors.BizError) {

	account = &entity.AccountEntity{}
	db := GetDB(dbNameKms).WithContext(ka.ctx).Table(TableNameAccount).
		Where(map[string]interface{}{
			AccountPrimaryKey: accountId,
		}).
		First(&account)

	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return account, nil
}

// GetAccountsByAccountIds 根据多个账号ID批量获取账号ID
func (ka *Account) GetAccountsByAccountIds(accountIds []int64) (accounts []*entity.AccountEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(ka.ctx).Table(TableNameAccount).
		Where("account_id IN (?)",
			accountIds).
		Find(&accounts)
	if db.Error != nil {
		return accounts, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return accounts, nil
}

// CountByName 根据账号名获取账号数
func (ka *Account) CountByName(name string) (count int64, err error) {
	if name == "" {
		return 0, nil
	}
	db := GetDB(dbNameKms).WithContext(ka.ctx).Table(TableNameAccount).
		Where(map[string]interface{}{
			"name": name,
		}).Count(&count)
	if db.Error != nil {
		return 0, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// CheckNameExists 检查账号名是否存在
func (ka *Account) CheckNameExists(name string) (bool, error) {
	count, err := ka.CountByName(name)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// HasSameName 账号ID和账号名是否存在
func (ka *Account) HasSameName(accountId int64, name string) (has bool, err errors.BizError) {
	var count int64
	db := GetDB(dbNameKms).WithContext(ka.ctx).Table(TableNameAccount).
		Where("name = ?", name).
		Where("account_id <> ?", accountId).
		Count(&count)
	if db.Error != nil {
		return false, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count > 0, nil
}

// Update 更新账号，只会更新如下字段
func (ka *Account) Update(account entity.AccountEntity) errors.BizError {
	account.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(ka.ctx).Table(TableNameAccount).
		Select("GivenName", "Email", "Phone", "Mobile", "UpdateTime").
		Where(map[string]interface{}{
			AccountPrimaryKey: account.AccountId,
		}).
		Updates(account)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// UpdatePassword 更新账号密码
func (ka *Account) UpdatePassword(accountId int64, password string) errors.BizError {
	updateTime := utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(ka.ctx).Table(TableNameAccount).
		Where(map[string]interface{}{
			AccountPrimaryKey: accountId,
		}).
		Update("password", password).
		Update("update_time", updateTime)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// UpdateStatus 更新账号状态
func (ka *Account) UpdateStatus(accountId int64, status int) errors.BizError {
	updateTime := utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(ka.ctx).Table(TableNameAccount).
		Where(map[string]interface{}{
			AccountPrimaryKey: accountId,
		}).
		Update("status", status).
		Update("update_time", updateTime)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// GetAllAccount 获取所有的账号
func (ka *Account) GetAllAccount() (account []*entity.AccountEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(ka.ctx).Table(TableNameAccount).
		Where(map[string]interface{}{}).
		Find(&account)
	if db.Error != nil {
		return account, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return account, nil
}

// GetAccountsByKeywords 根据账号名模糊匹配账号
func (ka *Account) GetAccountsByKeywords(keywords *entity.AccountKeywords) (accounts []*entity.AccountEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(ka.ctx).Table(TableNameAccount)
	if keywords.Status != "" {
		db = db.Where("status = ?", keywords.Status)
	}
	if keywords.GivenName != "" {
		db = db.Where("given_name = ?", keywords.GivenName)
	}
	if keywords.AccountName != "" {
		db = db.Where("name LIKE ?", "%"+keywords.AccountName+"%")
	}
	db.Find(&accounts)
	if db.Error != nil {
		return accounts, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return accounts, nil
}

// GetAccountsByKeywordAndLimit 根据关键字分页获取账号
func (ka *Account) GetAccountsByKeywordAndLimit(limit int, offset int, keywords *entity.AccountKeywords) (accounts []*entity.AccountEntity, err errors.BizError) {

	if keywords == nil {
		return ka.GetAccountsByLimit(limit, offset)
	}

	db := GetDB(dbNameKms).WithContext(ka.ctx).Table(TableNameAccount)
	if keywords.Status != "" {
		db = db.Where("status = ?", keywords.Status)
	}
	if keywords.GivenName != "" {
		db = db.Where("given_name = ?", keywords.GivenName)
	}
	if keywords.AccountName != "" {
		db = db.Where("name LIKE ?", "%"+keywords.AccountName+"%")
	}
	db = db.Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", AccountPrimaryKey)).
		Find(&accounts)
	if db.Error != nil {
		return accounts, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return accounts, nil
}

// GetAccountsByLimit 分页获取账号列表
func (ka *Account) GetAccountsByLimit(limit int, offset int) (accounts []*entity.AccountEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(ka.ctx).Table(TableNameAccount).
		Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", AccountPrimaryKey)).
		Find(&accounts)
	if db.Error != nil {
		return accounts, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return accounts, nil
}

// CountAccounts 账号总数
func (ka *Account) CountAccounts() (count int64, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(ka.ctx).Table(TableNameAccount).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// CountAccounts 账号总数
func (ka *Account) CountAccountsByKeywords(keywords *entity.AccountKeywords) (count int64, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(ka.ctx).Table(TableNameAccount)
	if keywords.Status != "" {
		db = db.Where("status = ?", keywords.Status)
	}
	if keywords.GivenName != "" {
		db = db.Where("given_name = ?", keywords.GivenName)
	}
	if keywords.AccountName != "" {
		db = db.Where("name LIKE ?", "%"+keywords.AccountName+"%")
	}
	db.Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}
