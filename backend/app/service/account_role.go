package service

import (
	"context"

	"github.com/phachon/mm-wiki/app/dao"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
	"github.com/phachon/mm-wiki/utils"
)

const (
	RoleAccountDefaultPageSize = 10 // 角色账号列表默认一页 10 条
)

// AccountRole 账号角色业务逻辑
type AccountRole struct {
	ctx            context.Context
	daoAccountRole *dao.AccountRole
}

// NewAccountRole 创建账号角色业务逻辑对象
func NewAccountRole(ctx context.Context) *AccountRole {
	return &AccountRole{
		ctx:            ctx,
		daoAccountRole: dao.NewAccountRole(ctx),
	}
}

// Create 创建账号角色关系
func (a *AccountRole) Create(accountId int64, roleId int64) errors.BizError {
	if accountId <= 0 || roleId <= 0 {
		return errors.Errorf(errors.BusinessRecordNotExistError, "账号或角色不合法")
	}
	accountRoleEntity := &entity.AccountRoleEntity{
		AccountId: accountId,
		RoleId:    roleId,
	}
	err := a.daoAccountRole.Insert(accountRoleEntity)
	if err != nil {
		return err
	}
	return nil
}

// BatchCreate 批量创建账号角色关系
func (a *AccountRole) BatchCreate(accountId int64, roleIds []int64) errors.BizError {
	if accountId <= 0 || len(roleIds) == 0 {
		return errors.Errorf(errors.BusinessRecordNotExistError, "账号或角色不合法")
	}
	for _, roleId := range roleIds {
		if err := a.Create(accountId, roleId); err != nil {
			return err
		}
	}
	return nil
}

// UpdateByAccountId 修改账号角色对应关系
func (a *AccountRole) UpdateByAccountId(accountId int64, roleIds []int64) errors.BizError {
	// 先删除
	err := a.daoAccountRole.DeleteByAccountId(accountId)
	if err != nil {
		return err
	}
	// 再创建
	return a.BatchCreate(accountId, roleIds)
}

// GetRoleIdsByAccountId 获取账号的角色列表
func (a *AccountRole) GetRoleIdsByAccountId(accountId int64) (roleIds []int64, err errors.BizError) {
	accountRoles, err := a.daoAccountRole.GetAccountRolesByAccountId(accountId)
	if err != nil {
		return roleIds, err
	}
	for _, accountRole := range accountRoles {
		roleIds = append(roleIds, accountRole.RoleId)
	}
	// 查找所有的角色
	return roleIds, nil
}

// GetRolesByAccountId 获取账号的角色列表
func (a *AccountRole) GetRolesByAccountId(accountId int64) (roles []*entity.RoleEntity, err errors.BizError) {
	accountRoles, err := a.daoAccountRole.GetAccountRolesByAccountId(accountId)
	if err != nil {
		return roles, err
	}
	var roleIds []int64
	for _, accountRole := range accountRoles {
		roleIds = append(roleIds, accountRole.RoleId)
	}
	// 查找所有的角色
	return dao.NewRole(a.ctx).GetRolesByRoleIds(roleIds)
}

// GetRoleAccountsByRoleId 根据角色ID查找角色账号关系
func (a *AccountRole) GetRoleAccountsByRoleId(roleId int64) (accountRoles []*entity.AccountRoleEntity,
	err errors.BizError) {
	return a.daoAccountRole.GetRoleAccountsByRoleId(roleId)
}

// GetRoleAccountsByRoleIdLimit 根据分页获取账号角色关系
func (a *AccountRole) GetRoleAccountsByRoleIdLimit(limit int, offset int, roleId int64) (accountRoles []*entity.AccountRoleEntity,
	err errors.BizError) {
	if roleId <= 0 {
		return accountRoles, nil
	}
	return a.daoAccountRole.GetRoleAccountsByRoleIdLimit(limit, offset, roleId)
}

// CountRoleAccountsByRoleId 根据角色ID获取账号总数
func (a *AccountRole) CountRoleAccountsByRoleId(roleId int64) (count int64, err errors.BizError) {
	if roleId <= 0 {
		return 0, nil
	}
	return a.daoAccountRole.CountRoleAccountsByRoleId(roleId)
}

// GetAccountsByRoleIdLimit 根据分页获取账号列表
func (a *AccountRole) GetAccountsByRoleIdLimit(pageSize int, pageNum int, roleId int64) (accounts []*entity.AccountEntity,
	err errors.BizError) {
	if roleId <= 0 {
		return accounts, nil
	}
	pageSize = utils.VerifyUint(pageSize, RoleListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	offset := (pageNum - 1) * pageSize
	roleAccounts, err := a.daoAccountRole.GetRoleAccountsByRoleIdLimit(pageSize, offset, roleId)
	if err != nil {
		return accounts, nil
	}
	var accountIds []int64
	for _, roleAccount := range roleAccounts {
		accountIds = append(accountIds, roleAccount.AccountId)
	}
	return dao.NewAccount(a.ctx).GetAccountsByAccountIds(accountIds)
}

// GetPageInfoLimitByRoleId 根据角色ID获取分页信息
func (a *AccountRole) GetPageInfoLimitByRoleId(pageSize int, pageNum int, roleId int64) (pagination *entity.PageInfo, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, RoleAccountDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	pageInfo := new(entity.PageInfo)
	var totalNum int64

	totalNum, err = a.daoAccountRole.CountRoleAccountsByRoleId(roleId)
	if err != nil {
		logger.WithContext(a.ctx).Warnf("[AccountRole] CountRoleAccountsByRoleId pageSize=%d pageNum=%d err=%s",
			pageSize, pageNum, err.Error())
		return pageInfo, err
	}
	return entity.GetPageInfo(totalNum, pageSize, pageNum), nil
}
