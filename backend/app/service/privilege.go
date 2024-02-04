package service

import (
	"context"
	"strings"

	"github.com/phachon/mm-wiki/app/dao"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
	"github.com/phachon/mm-wiki/utils"
)

// Privilege 权限业务逻辑
type Privilege struct {
	ctx          context.Context
	daoPrivilege *dao.Privilege
}

// NewPrivilege 创建权限业务逻辑对象
func NewPrivilege(ctx context.Context) *Privilege {
	return &Privilege{
		ctx:          ctx,
		daoPrivilege: dao.NewPrivilege(ctx),
	}
}

// CheckParentIds 检查 parentIds 是否合法
func (p *Privilege) CheckParentIds(privilegeType int, parentIdsStr string) ([]int64, errors.BizError) {
	if privilegeType == entity.PrivilegeTypeNav {
		if parentIdsStr != "" {
			return []int64{}, errors.Errorf(errors.ClientReqParamWrongful, "导航不能有上级权限")
		}
		return []int64{}, nil
	}
	parentIdList := strings.Split(parentIdsStr, ",")
	var parentIds []int64
	for _, parentIdStr := range parentIdList {
		parentId := utils.Convert.StringToInt64(parentIdStr)
		parentIds = append(parentIds, parentId)
	}
	privileges, err := p.daoPrivilege.GetPrivilegesByPrivilegeIds(parentIds)
	if err != nil {
		return parentIds, errors.Errorf(errors.ClientReqParamWrongful, err.GetErrMsg())
	}
	if len(privileges) != len(parentIds) {
		return parentIds, errors.Errorf(errors.ClientReqParamWrongful, "上级权限不合法")
	}
	parentPrivilege := privileges[len(privileges)-1]
	// 菜单的上级只能是导航或者菜单
	if privilegeType == entity.PrivilegeTypeMenu {
		if parentPrivilege.PrivilegeType != entity.PrivilegeTypeNav &&
			parentPrivilege.PrivilegeType != entity.PrivilegeTypeMenu {
			return parentIds, errors.Errorf(errors.ClientReqParamWrongful, "上级权限不合法")
		}
		return parentIds, nil
	}
	// 操作类型的上级只能是菜单
	if privilegeType == entity.PrivilegeTypeOperation {
		if parentPrivilege.PrivilegeType != entity.PrivilegeTypeMenu {
			return parentIds, errors.Errorf(errors.ClientReqParamWrongful, "操作上级权限不合法")
		}
		return parentIds, nil
	}
	return parentIds, errors.Errorf(errors.ClientReqParamWrongful, "权限类型错误")
}

// Create 创建权限
func (p *Privilege) Create(privilegeEntity *entity.PrivilegeEntity) errors.BizError {
	// 检查 parentIds 是否合法
	parentIds, err := p.CheckParentIds(privilegeEntity.PrivilegeType, privilegeEntity.ParentIds)
	if err != nil {
		return err
	}
	if len(parentIds) > 0 {
		privilegeEntity.ParentId = parentIds[len(parentIds)-1]
	}
	// 查找标识是否存在
	privilege, err := p.daoPrivilege.GetPrivilegeByIdentify(privilegeEntity.Identify)
	if err != nil {
		return err
	}
	if privilege != nil {
		return errors.Errorf(errors.BusinessRecordExistError, "权限标识 %s 已经存在", privilegeEntity.Name)
	}
	// 查找 name 是否存在
	privilege, err = p.daoPrivilege.GetPrivilegeByName(privilegeEntity.Name)
	if err != nil {
		return err
	}
	if privilege != nil {
		return errors.Errorf(errors.BusinessRecordExistError, "权限名 %s 已经存在", privilegeEntity.Name)
	}
	// 插入一条记录
	err = p.daoPrivilege.Insert(privilegeEntity)
	if err != nil {
		return err
	}
	return nil
}

// Update 修改权限
func (p *Privilege) Update(privilegeEntity entity.PrivilegeEntity) errors.BizError {
	// 查找权限是否存在
	updatePrivilege, err := p.daoPrivilege.GetPrivilegeByPrivilegeId(privilegeEntity.PrivilegeId)
	if err != nil {
		return err
	}
	if updatePrivilege == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "权限不存在")
	}
	if privilegeEntity.Name != "" {
		// 查找 name 是否存在
		hasName, pErr := p.daoPrivilege.HasSameName(privilegeEntity.PrivilegeId, privilegeEntity.Name)
		if pErr != nil {
			logger.WithContext(p.ctx).Errorf("[service.Privilege] Update HasSameName privilegeId=%d name=%s err=%s",
				privilegeEntity.PrivilegeId, privilegeEntity.Name, pErr.Error())
		}
		if hasName {
			return errors.Errorf(errors.BusinessRecordExistError, "权限名 %s 已经存在", privilegeEntity.Name)
		}
	}
	// 判断上级权限ID是否合法
	parentIds, err := p.CheckParentIds(privilegeEntity.PrivilegeType, privilegeEntity.ParentIds)
	if err != nil {
		return err
	}
	if len(parentIds) > 0 {
		privilegeEntity.ParentId = parentIds[len(parentIds)-1]
	}

	// 更新权限字段
	err = p.daoPrivilege.Update(privilegeEntity)
	if err != nil {
		return err
	}
	// 如果上级ID变更，需要更新所有的子权限下的上级ID
	if updatePrivilege.ParentIds != privilegeEntity.ParentIds {
		p.UpdateAllParentIdsContainParentId(privilegeEntity.PrivilegeId, privilegeEntity.ParentIds)
	}

	return nil
}

// UpdateAllParentIdsContainParentId 更新所有的包含 parent_id 的 parent_ids
func (p *Privilege) UpdateAllParentIdsContainParentId(parentId int64, prefixParentIds string) errors.BizError {
	// 先获取该父节点下的所有的子节点
	allPrivileges, err := p.daoPrivilege.GetPrivilegesContainParentId(parentId)
	if err != nil {
		return err
	}
	// eg: parentId=5
	for _, privilege := range allPrivileges {
		originParentIds := privilege.ParentIds
		newParentIds := utils.GetNewParentIds(originParentIds, parentId, prefixParentIds)
		err := p.daoPrivilege.UpdateParentIdsByPrivilegeId(privilege.PrivilegeId, newParentIds)
		if err != nil {
			logger.WithContext(p.ctx).Errorf("[Privilege] privilege_id=%d update parentIds=%s",
				privilege.PrivilegeId, newParentIds)
		}
	}
	return nil
}

// GetPrivilegesContainParentId 获取父节点下所有的子节点
func (p *Privilege) GetPrivilegesContainParentId(parentId int64) (privilege []*entity.PrivilegeEntity, err errors.BizError) {
	return p.daoPrivilege.GetPrivilegesContainParentId(parentId)
}

// GetPrivilegeByPrivilegeId 根据权限ID获取权限详情
func (p *Privilege) GetPrivilegeByPrivilegeId(privilegeID int64) (privilege *entity.PrivilegeEntity, err errors.BizError) {
	privilege, err = p.daoPrivilege.GetPrivilegeByPrivilegeId(privilegeID)
	if err != nil {
		return privilege, err
	}
	return privilege, nil
}

// GetPrivilegesByPrivilegeIds 根据权限ID获取权限详情
func (p *Privilege) GetPrivilegesByPrivilegeIds(privilegeIDs []int64) (privilege []*entity.PrivilegeEntity, err errors.BizError) {
	privilege, err = p.daoPrivilege.GetPrivilegesByPrivilegeIds(privilegeIDs)
	if err != nil {
		return privilege, err
	}
	return privilege, nil
}

// GetPrivilegeByParentId 根据权限ID获取权限详情
func (p *Privilege) GetPrivilegeByParentId(parentId int64) (privileges []*entity.PrivilegeEntity, err errors.BizError) {
	privileges, err = p.daoPrivilege.GetPrivilegesByParentId(parentId)
	if err != nil {
		return privileges, err
	}
	return privileges, nil
}

// GetPrivilegeByParentId 根据排序号正序获取所有的权限
func (p *Privilege) GetAllPrivilegesBySequence() (privileges []*entity.PrivilegeEntity, err errors.BizError) {
	privileges, err = p.daoPrivilege.GetAllPrivilegesBySequence()
	if err != nil {
		return privileges, err
	}
	return privileges, nil
}

// GetAllPrivilegeList 获取所有的权限列表
func (p *Privilege) GetAllPrivilegeList() (navPrivileges []*entity.PrivilegeListItem, err errors.BizError) {
	privileges, err := p.daoPrivilege.GetAllPrivilegesBySequence()
	if err != nil {
		return []*entity.PrivilegeListItem{}, err
	}
	privilegeList := p.GetPrivilegeChilds(privileges, 0)
	return privilegeList, nil
}

// GetNavAndMenuPrivileges 获取导航和菜单类型的权限
func (p *Privilege) GetNavAndMenuPrivileges() (navPrivileges []*entity.PrivilegeListItem, err errors.BizError) {
	privileges, err := p.daoPrivilege.GetAllPrivilegesBySequence()
	if err != nil {
		return []*entity.PrivilegeListItem{}, err
	}
	// 过滤非导航和菜单权限
	var filterPrivileges []*entity.PrivilegeEntity
	for _, privilege := range privileges {
		if privilege.PrivilegeType != entity.PrivilegeTypeNav &&
			privilege.PrivilegeType != entity.PrivilegeTypeMenu {
			continue
		}
		filterPrivileges = append(filterPrivileges, privilege)
	}
	privilegeList := p.GetPrivilegeChilds(filterPrivileges, 0)
	return privilegeList, nil
}

// GetPrivilegeChilds 递归获取所有的子权限
func (p *Privilege) GetPrivilegeChilds(privileges []*entity.PrivilegeEntity, parantId int64) []*entity.PrivilegeListItem {
	if len(privileges) == 0 {
		return []*entity.PrivilegeListItem{}
	}
	chlidPrivileges := []*entity.PrivilegeListItem{}
	for _, privilege := range privileges {
		if privilege.ParentId == parantId {
			chlidPrivileges = append(chlidPrivileges, &entity.PrivilegeListItem{
				PrivilegeEntity: privilege,
				ChildPrivileges: p.GetPrivilegeChilds(privileges, privilege.PrivilegeId),
				Action:          p.GetPrivilegeItemAction(privilege),
			})
		}
	}
	return chlidPrivileges
}

// GetPrivilegeItemAction 获取权限操作
func (p *Privilege) GetPrivilegeItemAction(privilege *entity.PrivilegeEntity) *entity.PrivilegeListAction {
	indentifys := global.ContextValueLoginIdentifys(p.ctx)
	action := &entity.PrivilegeListAction{
		IsEdit:   indentifys[global.PrivilegeIndentifyPrivilegeEdit],
		IsDelete: indentifys[global.PrivilegeIndentifyPrivilegeDelete],
	}
	return action
}

// GetPrivilegeListItems 递归获取所有的子权限
func (p *Privilege) GetPrivilegeListItems(privileges []*entity.PrivilegeEntity) (privilegeListItems []*entity.PrivilegeListItem,
	err errors.BizError) {
	privilegeListItems = []*entity.PrivilegeListItem{}
	if len(privileges) == 0 {
		return privilegeListItems, nil
	}
	// 过滤操作权限
	var filterPrivileges []*entity.PrivilegeEntity
	for _, privilege := range privileges {
		if privilege.PrivilegeType == entity.PrivilegeTypeOperation {
			continue
		}
		filterPrivileges = append(filterPrivileges, privilege)
	}
	privilegeList := p.GetPrivilegeChilds(filterPrivileges, 0)
	return privilegeList, nil
}

// DeleteByPrivilegeId 通过权限ID删除权限
func (p *Privilege) DeleteByPrivilegeId(privilegeId int64) errors.BizError {
	// 查找权限是否存在
	deletePrivilege, err := p.daoPrivilege.GetPrivilegeByPrivilegeId(privilegeId)
	if err != nil {
		return err
	}
	if deletePrivilege == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "权限不存在")
	}
	// 查找是否有子权限
	privilegeIds, err := p.daoPrivilege.GetPrivilegesByParentId(deletePrivilege.PrivilegeId)
	if err != nil {
		return err
	}
	if len(privilegeIds) > 0 {
		return errors.Errorf(errors.BusinessRecordExistError, "权限存在子权限，不能删除")
	}
	// 删除权限记录
	err = p.daoPrivilege.DeletePrivilege(privilegeId)
	if err != nil {
		return err
	}
	return nil
}

// GetAccountPrivileges 获取账号的权限列表
func (p *Privilege) GetAccountPrivileges(accountId int64) ([]*entity.PrivilegeEntity, errors.BizError) {
	// 获取账号所有的角色ID
	roleIds, err := NewAccountRole(p.ctx).GetRoleIdsByAccountId(accountId)
	if err != nil {
		return []*entity.PrivilegeEntity{}, err
	}
	// 获取角色所有的权限
	privileges, err := NewRole(p.ctx).GetPrivilegesByRoleIds(roleIds)
	if err != nil {
		return []*entity.PrivilegeEntity{}, err
	}
	return privileges, nil
}
