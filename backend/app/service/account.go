package service

import (
	"context"

	"github.com/phachon/mm-wiki/app/dao"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
	"github.com/phachon/mm-wiki/utils"
)

const (
	AccountListDefaultPageSize = 20 // 账号列表默认一页 20 条
)

// Account 账号业务逻辑
type Account struct {
	ctx        context.Context
	daoAccount *dao.Account
}

// NewAccount 创建账号业务逻辑对象
func NewAccount(ctx context.Context) *Account {
	return &Account{
		ctx:        ctx,
		daoAccount: dao.NewAccount(ctx),
	}
}

// Create 创建账号
func (a *Account) Create(accountEntity *entity.AccountEntity) errors.BizError {
	// 查找 name 是否存在
	account, err := a.daoAccount.GetAccountByName(accountEntity.Name)
	if err != nil {
		return err
	}
	if account != nil {
		return errors.Errorf(errors.BusinessRecordExistError, "账号名 %s 已经存在", accountEntity.Name)
	}
	if accountEntity.Password == "" {
		accountEntity.Password = global.DefAccountPass // 默认的密码
	}
	// 密码加密
	accountEntity.Password = PasswordEncode(accountEntity.Password)
	// 插入一条记录
	err = a.daoAccount.Insert(accountEntity)
	if err != nil {
		return err
	}
	return nil
}

// Update 修改账号
func (a *Account) Update(accountEntity entity.AccountEntity) errors.BizError {
	// 查找账号是否存在
	updateAccount, err := a.daoAccount.GetAccountByAccountId(accountEntity.AccountId)
	if err != nil {
		return err
	}
	if updateAccount == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "账号不存在")
	}
	if updateAccount.Status == entity.AccountStatusForbid {
		return errors.Errorf(errors.BusinessForbiddenError, "被禁用账号无法修改")
	}
	if accountEntity.Name != "" {
		// 查找 name 是否存在
		hasName, err := a.daoAccount.HasSameName(accountEntity.AccountId, accountEntity.Name)
		if err != nil {
			logger.WithContext(a.ctx).Errorf("[service.Account] Update HasSameName accountId=%d name=%s err=%s",
				accountEntity.AccountId, accountEntity.Name, err.Error())
		}
		if hasName {
			return errors.Errorf(errors.BusinessRecordExistError, "账号名 %s 已经存在", accountEntity.Name)
		}
	}
	// 更新字段
	err = a.daoAccount.Update(accountEntity)
	if err != nil {
		return err
	}
	return nil
}

// GetAccountByAccountId 根据账号ID获取账号详情
func (a *Account) GetAccountByAccountId(accountID int64) (account *entity.AccountEntity, err errors.BizError) {
	account, err = a.daoAccount.GetAccountByAccountId(accountID)
	if err != nil {
		return account, err
	}
	return account, nil
}

// UpdatePassword 修改账号密码
func (a *Account) UpdatePassword(accountId int64, password string, newPassword string) errors.BizError {
	// 查找账号是否存在
	updateAccount, err := a.daoAccount.GetAccountByAccountId(accountId)
	if err != nil {
		return err
	}
	if updateAccount == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "账号id %d 不存在", accountId)
	}
	if updateAccount.Status == entity.AccountStatusForbid {
		return errors.Errorf(errors.BusinessForbiddenError, "账号id %d 被禁用", accountId)
	}
	// 判断密码是否相同
	if PasswordEncode(password) != updateAccount.Password {
		return errors.Errorf(errors.BusinessPasswordError, "旧密码错误")
	}
	newPasswordEncode := PasswordEncode(newPassword)
	// 更新字段
	err = a.daoAccount.UpdatePassword(accountId, newPasswordEncode)
	if err != nil {
		return err
	}
	return nil
}

// UpdateStatus 更新账号状态
func (a *Account) UpdateStatus(accountId int64, status int) errors.BizError {
	// 查找账号是否存在
	updateAccount, err := a.daoAccount.GetAccountByAccountId(accountId)
	if err != nil {
		return err
	}
	if updateAccount == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "账号id %d 不存在", accountId)
	}
	// 更新状态字段
	err = a.daoAccount.UpdateStatus(accountId, status)
	if err != nil {
		return err
	}
	return nil
}

// GetAccountsByLimit 分页获取账号列表
func (a *Account) GetAccountsByLimit(pageSize int, pageNum int, keywords *entity.AccountKeywords) (accounts []*entity.AccountEntity, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, AccountListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	offset := (pageNum - 1) * pageSize
	// 无搜索词
	if keywords == nil {
		return a.daoAccount.GetAccountsByLimit(pageSize, offset)
	}
	return a.daoAccount.GetAccountsByKeywordAndLimit(pageSize, offset, keywords)
}

// GetPageInfoLimit 获取分页信息
func (a *Account) GetPageInfoLimit(pageSize int, pageNum int, keywords *entity.AccountKeywords) (pagination *entity.PageInfo, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, AccountListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	pageInfo := new(entity.PageInfo)
	var totalNum int64

	// 无搜索词
	if keywords == nil {
		totalNum, err = a.daoAccount.CountAccounts()
	} else {
		totalNum, err = a.daoAccount.CountAccountsByKeywords(keywords)
	}
	if err != nil {
		logger.WithContext(a.ctx).Warnf("[service.Account] GetPageInfoLimit pageSize=%d pageNum=%d err=%s",
			pageSize, pageNum, err.Error())
		return pageInfo, err
	}
	return entity.GetPageInfo(totalNum, pageSize, pageNum), nil
}

// FormatAccountList 格式化账号列表根据账号信息
func (a *Account) FormatAccountList(accounts []*entity.AccountEntity) (accountList []*entity.AccountListItem, err errors.BizError) {
	if len(accounts) == 0 {
		return accountList, nil
	}
	var accountIds []int64
	for _, account := range accounts {
		accountIds = append(accountIds, account.AccountId)
	}
	// 获取账号对应角色
	accountRoles, err := dao.NewAccountRole(a.ctx).GetAccountRolesByAccountIds(accountIds)
	if err != nil {
		return
	}
	var roleIds []int64
	var accountRoleIds = make(map[int64][]int64)
	for _, accountRole := range accountRoles {
		roleIds = append(roleIds, accountRole.RoleId)
		if roleIds, ok := accountRoleIds[accountRole.AccountId]; ok {
			accountRoleIds[accountRole.AccountId] = append(roleIds, accountRole.RoleId)
		} else {
			accountRoleIds[accountRole.AccountId] = []int64{accountRole.RoleId}
		}
	}
	// 查找所有的角色信息
	roles, err := dao.NewRole(a.ctx).GetRolesByRoleIds(roleIds)
	if err != nil {
		return
	}
	var rolesMap = make(map[int64]*entity.RoleEntity)
	for _, role := range roles {
		rolesMap[role.RoleId] = role
	}
	for _, account := range accounts {
		var accountRoles []*entity.RoleEntity
		roleIds := accountRoleIds[account.AccountId]
		for _, roleId := range roleIds {
			role := rolesMap[roleId]
			if role != nil {
				accountRoles = append(accountRoles, role)
			}
		}
		action := a.GetListItemAction(account)
		var accountListItem = &entity.AccountListItem{
			AccountEntity: account,
			Roles:         accountRoles,
			Action:        action,
		}
		accountList = append(accountList, accountListItem)
	}
	return accountList, nil
}

// SetAccountListItemAction 设置列表 action 权限
func (a *Account) GetListItemAction(accountItem *entity.AccountEntity) *entity.AccountListAction {
	indentifys := global.ContextValueLoginIdentifys(a.ctx)
	action := &entity.AccountListAction{
		IsEdit:         indentifys[global.PrivilegeIndentifyAccountEdit],
		IsDetail:       indentifys[global.PrivilegeIndentifyAccountDetail],
		IsUpdateStatus: indentifys[global.PrivilegeIndentifyAccountUpdateStatus],
	}
	return action
}

// GetAllNormalAccounts 获取所有正常的账号
func (a *Account) GetAllNormalAccounts() (accounts []*entity.AccountEntity, err errors.BizError) {
	return a.daoAccount.GetAllNormalAccounts()
}
