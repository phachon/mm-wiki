package service

import (
	"context"

	"github.com/phachon/mm-wiki/app/dao"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
)

const (
	SpacePermissionDefaultPageSize = 10 // 角色账号列表默认一页 10 条
)

// SpacePermission 空间权限业务逻辑
type SpacePermission struct {
	ctx                context.Context
	daoSpacePermission *dao.SpacePermission
}

// NewSpacePermission 创建空间权限业务逻辑对象
func NewSpacePermission(ctx context.Context) *SpacePermission {
	return &SpacePermission{
		ctx:                ctx,
		daoSpacePermission: dao.NewSpacePermission(ctx),
	}
}

// Create 创建空间权限关系
func (a *SpacePermission) Create(permission *entity.SpacePermissionEntity) errors.BizError {
	if permission == nil {
		return nil
	}
	if permission.SpaceId <= 0 {
		return errors.Errorf(errors.BusinessRecordNotExistError, "空间不合法")
	}

	// 如果是账号关联，账号ID必须大于0
	if permission.RelationType == entity.SpacePermissionRelationTypeAccount {
		if permission.AccountId <= 0 {
			return errors.Errorf(errors.BusinessRecordNotExistError, "账号不存在")
		}
		// 检查记录是否存在
		_, err := a.daoSpacePermission.GetPermissionByAccountIdAndSpaceId(permission.AccountId, permission.SpaceId)
		if err == nil {
			return errors.Errorf(errors.BusinessRecordExistError, "该账号权限已存在")
		}
	}

	// 如果是部门关联，部门ID必须大于0
	if permission.RelationType == entity.SpacePermissionRelationTypeDepartment {
		if permission.DepartmentId <= 0 {
			return errors.Errorf(errors.BusinessRecordNotExistError, "部门不存在")
		}
		// 检查记录是否存在
		_, err := a.daoSpacePermission.GetPermissionByDepartmentIdAndSpaceId(permission.DepartmentId, permission.SpaceId)
		if err == nil {
			return errors.Errorf(errors.BusinessRecordExistError, "该部门权限已存在")
		}
	}
	// 插入记录
	err := a.daoSpacePermission.Insert(permission)
	if err != nil {
		return err
	}
	return nil
}

// 批量添加管理员权限
func (a *SpacePermission) CreateBatchAdminPerssions(spaceId int64, accountIds []int64) errors.BizError {
	if spaceId <= 0 {
		return errors.Errorf(errors.BusinessRecordNotExistError, "空间不合法")
	}
	if len(accountIds) == 0 {
		return errors.Errorf(errors.BusinessRecordNotExistError, "账号不存在")
	}
	// 循环插入
	for _, accountId := range accountIds {
		permission := &entity.SpacePermissionEntity{
			SpaceId:      spaceId,
			RelationType: entity.SpacePermissionRelationTypeAccount,
			AccountId:    accountId,
			IsAdmin:      entity.SpacePermissionIsAdmin,
		}
		err := a.Create(permission)
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdateAccountPermission 修改空间账号权限
func (a *SpacePermission) UpdateAccountPermission(spacePermissionId int64, permission *entity.SpacePermission) errors.BizError {

	spacePermission := &entity.SpacePermissionEntity{
		SpacePermissionId: spacePermissionId,
		IsView:            permission.IsView,
		IsAdd:             permission.IsAdd,
		IsEdit:            permission.IsEdit,
		IsDelete:          permission.IsDelete,
		IsAdmin:           permission.IsAdmin,
	}
	// 更新记录
	err := a.daoSpacePermission.UpdatePermission(spacePermission)
	return err
}

// GetAdminsBySpaceId 获取空间管理员list
func (a *SpacePermission) GetAdminsBySpaceId(spaceId int64) ([]*entity.AccountEntity, errors.BizError) {
	permissions, err := a.daoSpacePermission.GetPermissionsBySpaceId(spaceId)
	if err != nil {
		return nil, err
	}
	accountIds := make([]int64, 0)
	for _, permission := range permissions {
		if permission.IsAdmin == entity.SpacePermissionIsAdmin {
			accountIds = append(accountIds, permission.AccountId)
		}
	}
	accounts, err := dao.NewAccount(a.ctx).GetAccountsByAccountIds(accountIds)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// DeleteBySpaceIdAccountId 根据空间ID和账号ID删除权限
func (a *SpacePermission) DeleteBySpaceIdAccountId(spaceId int64, accountId int64) errors.BizError {
	return a.daoSpacePermission.DeleteBySpaceIdAccountId(spaceId, accountId)
}

// GetSpacesByAdminId 根据账号ID获取空间列表
func (a *SpacePermission) GetSpacesByAdminId(accountId int64) ([]*entity.SpaceEntity, errors.BizError) {
	permissions, err := a.daoSpacePermission.GetPermissionsByAccountId(accountId)
	if err != nil {
		return nil, err
	}
	spaceIds := make([]int64, 0)
	for _, permission := range permissions {
		spaceIds = append(spaceIds, permission.SpaceId)
	}
	spaces, err := dao.NewSpace(a.ctx).GetSpacesBySpaceIds(spaceIds)
	if err != nil {
		return nil, err
	}
	return spaces, nil
}
