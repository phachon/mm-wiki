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

// Department 部门业务逻辑
type Department struct {
	ctx           context.Context
	daoDepartment *dao.Department
}

// NewDepartment 创建部门业务逻辑对象
func NewDepartment(ctx context.Context) *Department {
	return &Department{
		ctx:           ctx,
		daoDepartment: dao.NewDepartment(ctx),
	}
}

// CheckParentIds 检查 parentIds 是否合法
func (p *Department) CheckParentIds(parentIdsStr string) ([]int64, errors.BizError) {
	parentIdList := strings.Split(parentIdsStr, ",")
	var parentIds []int64
	for _, parentIdStr := range parentIdList {
		parentId := utils.Convert.StringToInt64(parentIdStr)
		parentIds = append(parentIds, parentId)
	}
	departments, err := p.daoDepartment.GetDepartmentsByDepartmentIds(parentIds)
	if err != nil {
		return parentIds, errors.Errorf(errors.ClientReqParamWrongful, err.GetErrMsg())
	}
	if len(departments) != len(parentIds) {
		return parentIds, errors.Errorf(errors.ClientReqParamWrongful, "上级部门不合法")
	}
	return parentIds, nil
}

// Create 创建部门
func (p *Department) Create(departmentEntity *entity.DepartmentEntity) errors.BizError {
	// 查找 name 是否存在
	department, err := p.daoDepartment.GetDepartmentByNameAndParentId(departmentEntity.Name, departmentEntity.ParentId)
	if err != nil {
		return err
	}
	if department != nil {
		return errors.Errorf(errors.BusinessRecordExistError, "该部门下部门名 %s 已经存在", departmentEntity.Name)
	}
	// 插入一条记录
	err = p.daoDepartment.Insert(departmentEntity)
	if err != nil {
		return err
	}
	return nil
}

// Update 修改部门
func (p *Department) Update(departmentEntity entity.DepartmentEntity) errors.BizError {
	// 查找部门是否存在
	updateDepartment, err := p.daoDepartment.GetDepartmentByDepartmentId(departmentEntity.DepartmentId)
	if err != nil {
		return err
	}
	if updateDepartment == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "部门不存在")
	}
	if departmentEntity.Name != "" {
		// 查找 name 是否存在
		hasName, pErr := p.daoDepartment.HasSameName(
			departmentEntity.ParentId,
			departmentEntity.Name,
			departmentEntity.DepartmentId,
		)
		if pErr != nil {
			logger.WithContext(p.ctx).Errorf("[service.Department] Update HasSameName departmentId=%d name=%s err=%s",
				departmentEntity.DepartmentId, departmentEntity.Name, pErr.Error())
		}
		if hasName {
			return errors.Errorf(errors.BusinessRecordExistError, "部门名 %s 已经存在", departmentEntity.Name)
		}
	}
	// 更新部门字段
	err = p.daoDepartment.Update(departmentEntity)
	if err != nil {
		return err
	}
	// 如果上级ID变更，需要更新所有的子部门下的上级ID
	if updateDepartment.ParentIds != departmentEntity.ParentIds {
		p.UpdateAllParentIdsContainParentId(departmentEntity.DepartmentId, departmentEntity.ParentIds)
	}

	return nil
}

// UpdateAllParentIdsContainParentId 更新所有的包含 parent_id 的 parent_ids
func (p *Department) UpdateAllParentIdsContainParentId(parentId int64, prefixParentIds string) errors.BizError {
	// 先获取该父节点下的所有的子节点
	allDepartments, err := p.daoDepartment.GetDepartmentsContainParentId(parentId)
	if err != nil {
		return err
	}
	// eg: parentId=5
	for _, department := range allDepartments {
		originParentIds := department.ParentIds
		newParentIds := utils.GetNewParentIds(originParentIds, parentId, prefixParentIds)
		err := p.daoDepartment.UpdateParentIdsByDepartmentId(department.DepartmentId, newParentIds)
		if err != nil {
			logger.WithContext(p.ctx).Errorf("[Department] department_id=%d update parentIds=%s",
				department.DepartmentId, newParentIds)
		}
	}
	return nil
}

// GetDepartmentsContainParentId 获取父节点下所有的子节点
func (p *Department) GetDepartmentsContainParentId(parentId int64) (department []*entity.DepartmentEntity, err errors.BizError) {
	return p.daoDepartment.GetDepartmentsContainParentId(parentId)
}

// GetDepartmentByDepartmentId 根据部门ID获取部门详情
func (p *Department) GetDepartmentByDepartmentId(departmentID int64) (department *entity.DepartmentEntity, err errors.BizError) {
	department, err = p.daoDepartment.GetDepartmentByDepartmentId(departmentID)
	if err != nil {
		return department, err
	}
	return department, nil
}

// GetDepartmentsByDepartmentIds 根据部门ID获取部门详情
func (p *Department) GetDepartmentsByDepartmentIds(departmentIDs []int64) (departments []*entity.DepartmentEntity, err errors.BizError) {
	departments, err = p.daoDepartment.GetDepartmentsByDepartmentIds(departmentIDs)
	if err != nil {
		return departments, err
	}
	return departments, nil
}

// GetDepartmentByParentId 根据上级部门ID获取部门详情
func (p *Department) GetDepartmentByParentId(parentId int64) (departments []*entity.DepartmentEntity, err errors.BizError) {
	departments, err = p.daoDepartment.GetDepartmentsByParentId(parentId)
	if err != nil {
		return departments, err
	}
	return departments, nil
}

// GetDepartmentByParentId 根据排序号正序获取所有的部门
func (p *Department) GetAllDepartmentsBySequence() (departments []*entity.DepartmentEntity, err errors.BizError) {
	departments, err = p.daoDepartment.GetAllDepartmentsBySequence()
	if err != nil {
		return departments, err
	}
	return departments, nil
}

// GetAllDepartmentList 获取所有的部门列表
func (p *Department) GetAllDepartmentList() (navDepartments []*entity.DepartmentListItem, err errors.BizError) {
	departments, err := p.daoDepartment.GetAllDepartmentsBySequence()
	if err != nil {
		return []*entity.DepartmentListItem{}, err
	}
	departmentList := p.GetDepartmentChilds(departments, 0)
	return departmentList, nil
}

// GetDepartmentChilds 递归获取所有的子部门
func (p *Department) GetDepartmentChilds(departments []*entity.DepartmentEntity, parantId int64) []*entity.DepartmentListItem {
	if len(departments) == 0 {
		return []*entity.DepartmentListItem{}
	}
	chlidDepartments := []*entity.DepartmentListItem{}
	for _, department := range departments {
		if department.ParentId == parantId {
			chlidDepartments = append(chlidDepartments, &entity.DepartmentListItem{
				DepartmentEntity: department,
				Children:         p.GetDepartmentChilds(departments, department.DepartmentId),
				Action:           p.GetDepartmentItemAction(department),
			})
		}
	}
	return chlidDepartments
}

// GetDepartmentItemAction 获取部门操作
func (p *Department) GetDepartmentItemAction(department *entity.DepartmentEntity) *entity.DepartmentListAction {
	indentifys := global.ContextValueLoginIdentifys(p.ctx)
	action := &entity.DepartmentListAction{
		IsAdd:    indentifys[global.PrivilegeIndentifyDepartmentAdd],
		IsEdit:   indentifys[global.PrivilegeIndentifyDepartmentEdit],
		IsDelete: indentifys[global.PrivilegeIndentifyDepartmentDelete],
	}
	return action
}

// GetDepartmentListItems 递归获取所有的子部门
func (p *Department) GetDepartmentListItems(departments []*entity.DepartmentEntity) (departmentListItems []*entity.DepartmentListItem,
	err errors.BizError) {
	departmentListItems = []*entity.DepartmentListItem{}
	if len(departments) == 0 {
		return departmentListItems, nil
	}
	departmentList := p.GetDepartmentChilds(departments, 0)
	return departmentList, nil
}

// DeleteByDepartmentId 通过部门ID删除部门
func (p *Department) DeleteByDepartmentId(departmentId int64) errors.BizError {
	// 查找部门是否存在
	deleteDepartment, err := p.daoDepartment.GetDepartmentByDepartmentId(departmentId)
	if err != nil {
		return err
	}
	if deleteDepartment == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "部门不存在")
	}
	// 查找是否有子部门
	departmentIds, err := p.daoDepartment.GetDepartmentsByParentId(deleteDepartment.DepartmentId)
	if err != nil {
		return err
	}
	if len(departmentIds) > 0 {
		return errors.Errorf(errors.BusinessRecordExistError, "该部门下存在子部门，不能直接删除！")
	}
	// 删除部门
	err = p.daoDepartment.DeleteDepartment(departmentId)
	if err != nil {
		return err
	}

	// todo 删除部门下的所有用户

	// todo 删除部门下的所有文档

	return nil
}
