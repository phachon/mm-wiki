package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/phachon/mm-wiki/app/dao"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
	"github.com/phachon/mm-wiki/utils"
)

const (
	RoleListDefaultPageSize = 20 // 角色列表默认一页 20 条
)

// Role 角色业务逻辑
type Role struct {
	ctx            context.Context
	daoRole        *dao.Role
	daoAccountRole *dao.AccountRole
	daoPrivilege   *dao.Privilege
}

// NewRole 创建角色业务逻辑对象
func NewRole(ctx context.Context) *Role {
	return &Role{
		ctx:            ctx,
		daoRole:        dao.NewRole(ctx),
		daoAccountRole: dao.NewAccountRole(ctx),
		daoPrivilege:   dao.NewPrivilege(ctx),
	}
}

// Create 创建角色
func (r *Role) Create(roleEntity *entity.RoleEntity) errors.BizError {
	// 查找 name 是否存在
	role, err := r.daoRole.GetRoleByName(roleEntity.Name)
	if err != nil {
		return err
	}
	if role != nil {
		return errors.Errorf(errors.BusinessRecordExistError, "角色名 %s 已经存在", roleEntity.Name)
	}
	// 插入一条记录
	err = r.daoRole.Insert(roleEntity)
	if err != nil {
		return err
	}
	return nil
}

// Update 修改角色
func (r *Role) Update(roleEntity entity.RoleEntity) errors.BizError {
	// 查找角色是否存在
	updateRole, err := r.daoRole.GetRoleByRoleId(roleEntity.RoleId)
	if err != nil {
		return err
	}
	if updateRole == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "角色不存在")
	}
	// 查找 name 是否存在
	hasName, err := r.daoRole.HasSameName(roleEntity.RoleId, roleEntity.Name)
	if err != nil {
		logger.WithContext(r.ctx).Errorf("[service.Role] Update HasSameName roleId=%d name=%s err=%s",
			roleEntity.RoleId, roleEntity.Name, err.Error())
	}
	if hasName {
		return errors.Errorf(errors.BusinessRecordExistError, "角色名 %s 已经存在", roleEntity.Name)
	}
	// 更新字段
	err = r.daoRole.Update(roleEntity)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePrivilegeIds 更新角色权限
func (r *Role) UpdatePrivilegeIds(roleId int64, privilegeIds string) errors.BizError {
	// 查找角色是否存在
	updateRole, err := r.daoRole.GetRoleByRoleId(roleId)
	if err != nil {
		return err
	}
	if updateRole == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "角色不存在")
	}
	// 更新字段
	roleEntity := entity.RoleEntity{
		RoleId:       roleId,
		PrivilegeIds: privilegeIds,
	}
	err = r.daoRole.UpdatePrivilegeIds(roleEntity)
	if err != nil {
		return err
	}
	return nil
}

// GetRoleByRoleId 根据角色ID获取角色详情
func (r *Role) GetRoleByRoleId(roleID int64) (role *entity.RoleEntity, err errors.BizError) {
	role, err = r.daoRole.GetRoleByRoleId(roleID)
	if err != nil {
		return role, err
	}
	return role, nil
}

// DeleteRole 删除角色
func (r *Role) DeleteRole(roleId int64) errors.BizError {
	// 查找角色是否存在
	updateRole, err := r.daoRole.GetRoleByRoleId(roleId)
	if err != nil {
		return err
	}
	if updateRole == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "角色id %d 不存在", roleId)
	}
	// 角色下是否有账号
	accountRoles, err := r.daoAccountRole.GetRoleAccountsByRoleId(roleId)
	if err != nil {
		return err
	}
	if len(accountRoles) > 0 {
		return errors.Errorf(errors.BusinessRecordExistError, "角色已有账号，请先移除后再删除")
	}
	// 删除角色
	err = r.daoRole.DeleteRole(roleId)
	if err != nil {
		return err
	}
	// 删除账号角色对应关系
	err = r.daoAccountRole.DeleteByRoleId(roleId)
	if err != nil {
		return err
	}
	return nil
}

// RemoveRoleAccount 移除角色下的账号
func (r *Role) RemoveRoleAccount(roleId int64, accountId int64) errors.BizError {
	// 删除账号角色对应关系
	err := r.daoAccountRole.DeleteByRoleIdAccountId(roleId, accountId)
	if err != nil {
		return err
	}
	return nil
}

// GetRolesByLimit 分页获取角色列表
func (r *Role) GetRolesByLimit(pageSize int, pageNum int, keywords *entity.RoleKeywords) (roles []*entity.RoleEntity, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, RoleListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	offset := (pageNum - 1) * pageSize
	// 无搜索词
	if keywords == nil {
		return r.daoRole.GetRolesByLimit(pageSize, offset)
	}
	return r.daoRole.GetRolesByKeywordsAndLimit(pageSize, offset, keywords)
}

// GetPageInfoLimit 获取分页信息
func (r *Role) GetPageInfoLimit(pageSize int, pageNum int, keywords *entity.RoleKeywords) (pagination *entity.PageInfo, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, RoleListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	pageInfo := new(entity.PageInfo)
	var totalNum int64
	// 无搜索词
	if keywords == nil {
		totalNum, err = r.daoRole.CountRoles()
	} else {
		totalNum, err = r.daoRole.CountRolesByKeywords(keywords)
	}
	if err != nil {
		logger.WithContext(r.ctx).Warnf("[service.Role] GetPageInfoLimit pageSize=%d pageNum=%d err=%s",
			pageSize, pageNum, err.Error())
		return pageInfo, err
	}
	return entity.GetPageInfo(totalNum, pageSize, pageNum), nil
}

// GetAllRoles 获取所有的角色
func (r *Role) GetAllRoles() (roles []*entity.RoleEntity, err errors.BizError) {
	return r.daoRole.GetAllRoles()
}

// GetRolesByRoleIds 根据角色ID批量获取角色
func (r *Role) GetRolesByRoleIds(roleIds []int64) (roles []*entity.RoleEntity, err errors.BizError) {
	return r.daoRole.GetRolesByRoleIds(roleIds)
}

// GetPrivilegesByRoleIds 根据多个角色ID查找权限列表
func (a *Role) GetPrivilegesByRoleIds(roleIds []int64) (privileges []*entity.PrivilegeEntity,
	err errors.BizError) {
	roles, err := a.daoRole.GetRolesByRoleIds(roleIds)
	if err != nil {
		return []*entity.PrivilegeEntity{}, err
	}
	var allPrivilegeIds []int64
	for _, role := range roles {
		// 超级管理员角色默认返回所有权限
		if role.RoleType == entity.RoleTypeRootRole {
			return a.daoPrivilege.GetAllPrivilegesBySequence()
		}
		privilegeIdsStr := role.PrivilegeIds
		if len(privilegeIdsStr) == 0 {
			continue
		}
		privilegeIds := strings.Split(privilegeIdsStr, ",")
		for _, privilegeIdStr := range privilegeIds {
			privilegeIdInt := utils.Convert.StringToInt64(privilegeIdStr)
			allPrivilegeIds = append(allPrivilegeIds, privilegeIdInt)
		}
	}
	return a.daoPrivilege.GetPrivilegesByPrivilegeIds(allPrivilegeIds)
}

// DeletePrivilegeIdsByPrivilegeId 根据权限ID删除角色中的权限ID
func (a *Role) DeletePrivilegeIdsByPrivilegeId(privilegeId int64) (err errors.BizError) {
	roles, err := a.daoRole.GetRolesByPrivilegeId(privilegeId)
	if err != nil {
		return err
	}
	for _, role := range roles {
		privilegeIds := strings.Split(role.PrivilegeIds, ",")
		passPrivilegeIds := utils.FilterStringIds(privilegeIds, fmt.Sprintf("%d", privilegeId))
		// 去除掉id, 需要更新
		updatePrivilegeIds := strings.Join(passPrivilegeIds, ",")
		// 如果和原来ID一样，不需要更新
		if updatePrivilegeIds == role.PrivilegeIds {
			continue
		}
		// 更新
		updateErr := a.UpdatePrivilegeIds(role.RoleId, updatePrivilegeIds)
		if updateErr != nil {
			logger.WithContext(a.ctx).Errorf(
				"[servicel.Role] UpdatePrivilegeIds role=%d, updatePrivilegeIds=%+v err=%+v",
				role.RoleId, updatePrivilegeIds, updateErr)
		}
	}
	return nil
}

// FormatRoleList 格式化账号列表根据账号信息
func (a *Role) FormatRoleList(roles []*entity.RoleEntity) (roleList []*entity.RoleListItem, err errors.BizError) {
	if len(roles) == 0 {
		return roleList, nil
	}
	for _, role := range roles {
		action := a.GetListItemAction(role)
		var roleListItem = &entity.RoleListItem{
			RoleEntity: role,
			Action:     action,
		}
		roleList = append(roleList, roleListItem)
	}
	return roleList, nil
}

// GetListItemAction 获取列表 action 权限
func (a *Role) GetListItemAction(roleItem *entity.RoleEntity) *entity.RoleListAction {
	indentifys := global.ContextValueLoginIdentifys(a.ctx)
	action := &entity.RoleListAction{
		IsEdit:          indentifys[global.PrivilegeIndentifyRoleEdit],
		IsDelete:        indentifys[global.PrivilegeIndentifyRoleDelete],
		IsAccountList:   indentifys[global.PrivilegeIndentifyRoleAccountList],
		IsPrivilegeEdit: indentifys[global.PrivilegeIndentifyRolePrivilegeEdit],
	}
	return action
}
