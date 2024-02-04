package service

import (
	"context"

	"github.com/phachon/mm-wiki/app/dao"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/utils"
)

// RolePrivilege 角色权限业务逻辑
type RolePrivilege struct {
	ctx              context.Context
	daoRolePrivilege *dao.RolePrivilege
	daoPrivilege     *dao.Privilege
}

// NewRolePrivilege 创建角色权限业务逻辑对象
func NewRolePrivilege(ctx context.Context) *RolePrivilege {
	return &RolePrivilege{
		ctx:              ctx,
		daoRolePrivilege: dao.NewRolePrivilege(ctx),
		daoPrivilege:     dao.NewPrivilege(ctx),
	}
}

// Create 创建角色权限关系
func (a *RolePrivilege) Create(roleId int64, privilegeId int64) errors.BizError {
	if roleId <= 0 || privilegeId <= 0 {
		return errors.Errorf(errors.BusinessRecordNotExistError, "角色权限不合法")
	}
	rolePrivilegeEntity := &entity.RolePrivilegeEntity{
		RoleId:      roleId,
		PrivilegeId: privilegeId,
	}
	err := a.daoRolePrivilege.Insert(rolePrivilegeEntity)
	if err != nil {
		return err
	}
	return nil
}

// BatchCreate 批量创建角色权限关系
func (a *RolePrivilege) BatchCreate(accountId int64, roleIds []int64) errors.BizError {
	if accountId <= 0 || len(roleIds) == 0 {
		return errors.Errorf(errors.BusinessRecordNotExistError, "角色权限不合法")
	}
	for _, roleId := range roleIds {
		if err := a.Create(accountId, roleId); err != nil {
			return err
		}
	}
	return nil
}

// UpdateByRoleId 修改角色权限对应关系
func (a *RolePrivilege) UpdateByRoleId(roleId int64, privilegeIds []int64) errors.BizError {

	// privilegeIds 先去重
	var privilegeIdsMap = make(map[int64]uint8)
	var realPrivilegeIds []int64
	for _, privilegeId := range privilegeIds {
		if _, ok := privilegeIdsMap[privilegeId]; ok {
			continue
		}
		realPrivilegeIds = append(realPrivilegeIds, privilegeId)
		privilegeIdsMap[privilegeId] = 1
	}

	// 查找角色下的所有权限
	rolePrivileges, err := a.daoRolePrivilege.GetRolePrivilegesByRoleId(roleId)
	if err != nil {
		return err
	}
	var resPrivilegeIds []int64
	for _, rolePrivilege := range rolePrivileges {
		resPrivilegeIds = append(resPrivilegeIds, rolePrivilege.PrivilegeId)
	}
	addPrivilegeIds, deletePrivilegeIds := utils.GetIdsDiffRes(resPrivilegeIds, realPrivilegeIds)
	// 先删除
	if len(deletePrivilegeIds) > 0 {
		err = a.daoRolePrivilege.DeletePrivilegeIdsByRoleId(roleId, deletePrivilegeIds)
		if err != nil {
			return err
		}
	}
	// 再创建
	if len(addPrivilegeIds) > 0 {
		err = a.BatchCreate(roleId, privilegeIds)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetPrivilegeIdsByRoleId 获取角色的权限列表
func (a *RolePrivilege) GetPrivilegeIdsByRoleId(roleId int64) (privileges []*entity.PrivilegeEntity, err errors.BizError) {
	rolePrivileges, err := a.daoRolePrivilege.GetRolePrivilegesByRoleId(roleId)
	if err != nil {
		return privileges, err
	}
	var privilegeIds []int64
	for _, rolePrivilege := range rolePrivileges {
		privilegeIds = append(privilegeIds, rolePrivilege.RoleId)
	}
	return dao.NewPrivilege(a.ctx).GetPrivilegesByPrivilegeIds(privilegeIds)
}

// GetRolePrivilegesByRoleId 根据角色ID查找角色账号关系
func (a *RolePrivilege) GetRolePrivilegesByRoleId(roleId int64) (rolePrivileges []*entity.RolePrivilegeEntity,
	err errors.BizError) {
	return a.daoRolePrivilege.GetRolePrivilegesByRoleId(roleId)
}

// GetRolePrivilegesByRoleIds 根据多个角色ID查找角色账号关系
func (a *RolePrivilege) GetRolePrivilegesByRoleIds(roleIds []int64) (rolePrivileges []*entity.RolePrivilegeEntity,
	err errors.BizError) {
	return a.daoRolePrivilege.GetRolePrivilegesByRoleIds(roleIds)
}

// GetPrivilegesByRoleIds 根据多个角色ID查找角色账号关系
func (a *RolePrivilege) GetPrivilegesByRoleIds(roleIds []int64) (privileges []*entity.PrivilegeEntity,
	err errors.BizError) {
	rolePrivileges, err := a.daoRolePrivilege.GetRolePrivilegesByRoleIds(roleIds)
	if err != nil {
		return []*entity.PrivilegeEntity{}, err
	}
	var privilegeIds []int64
	for _, rolePrivilege := range rolePrivileges {
		privilegeIds = append(privilegeIds, rolePrivilege.PrivilegeId)
	}
	return a.daoPrivilege.GetPrivilegesByPrivilegeIds(privilegeIds)
}
