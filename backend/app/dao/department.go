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
	// TableNameDepartment 部门表
	TableNameDepartment = "mk_department"
	// DepartmentPrimaryKey 部门表主键ID
	DepartmentPrimaryKey = "department_id"
)

// Department 部门表数据
type Department struct {
	ctx context.Context
}

// NewDepartment 创建部门表数据对象
func NewDepartment(ctx context.Context) *Department {
	return &Department{
		ctx: ctx,
	}
}

// Insert 创建部门插入一条部门记录
func (p *Department) Insert(departmentEntity *entity.DepartmentEntity) errors.BizError {
	departmentEntity.CreateTime = utils.NewJsonTime(time.Now())
	departmentEntity.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(p.ctx).Table(TableNameDepartment).Save(departmentEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetDepartmentByName 根据部门名查找同一部门下部门
func (p *Department) GetDepartmentByNameAndParentId(name string, parentId int64) (
	department *entity.DepartmentEntity, err errors.BizError) {
	department = &entity.DepartmentEntity{}
	db := GetDB(dbNameMK).WithContext(p.ctx).Table(TableNameDepartment).
		Where(map[string]interface{}{
			"parent_id": parentId,
			"name":      name,
		}).
		First(&department)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return department, nil
}

// GetDepartmentByDepartmentId 根据部门ID获取部门信息
func (p *Department) GetDepartmentByDepartmentId(departmentId int64) (department *entity.DepartmentEntity, err errors.BizError) {

	department = &entity.DepartmentEntity{}
	db := GetDB(dbNameMK).WithContext(p.ctx).Table(TableNameDepartment).
		Where(map[string]interface{}{
			DepartmentPrimaryKey: departmentId,
		}).
		First(&department)

	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return department, nil
}

// GetDepartmentsByDepartmentIds 根据多个部门ID批量获取部门ID
func (p *Department) GetDepartmentsByDepartmentIds(departmentIds []int64) (departments []*entity.DepartmentEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(p.ctx).Table(TableNameDepartment).
		Where("department_id IN (?)", departmentIds).
		Order(fmt.Sprintf("%s ASC", "sequence")).
		Find(&departments)
	if db.Error != nil {
		return departments, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return departments, nil
}

// CountByNameAndDepartmentId 查找同一部门下的部门名数量
func (p *Department) CountByNameAndParentId(parentId int64, name string) (count int64, err error) {
	if name == "" {
		return 0, nil
	}
	db := GetDB(dbNameMK).WithContext(p.ctx).Table(TableNameDepartment).
		Where(map[string]interface{}{
			"parent_id": parentId,
			"name":      name,
		}).Count(&count)
	if db.Error != nil {
		return 0, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// HasSameName 部门ID和部门名是否存在
func (p *Department) HasSameName(parentId int64, name string, departmentId int64) (has bool, err errors.BizError) {
	var count int64
	db := GetDB(dbNameMK).WithContext(p.ctx).Table(TableNameDepartment).
		Where("parent_id = ?", parentId).
		Where("name = ?", name).
		Where("department_id <> ?", departmentId).
		Count(&count)
	if db.Error != nil {
		return false, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count > 0, nil
}

// Update 更新部门，只会更新如下字段
func (p *Department) Update(department entity.DepartmentEntity) errors.BizError {
	department.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(p.ctx).Table(TableNameDepartment).
		Select("Name", "ParentId", "ParentIds", "Sequence").
		Where(map[string]interface{}{
			DepartmentPrimaryKey: department.DepartmentId,
		}).
		Updates(department)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// UpdateParentIdsByDepartmentId 更新 parent_ids
func (p *Department) UpdateParentIdsByDepartmentId(departmentId int64, parentIds string) errors.BizError {
	db := GetDB(dbNameMK).WithContext(p.ctx).Table(TableNameDepartment).
		Where(DepartmentPrimaryKey, departmentId).
		Update("parent_ids", parentIds)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// GetAllDepartment 获取所有的部门
func (p *Department) GetAllDepartment() (department []*entity.DepartmentEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(p.ctx).Table(TableNameDepartment).
		Where(map[string]interface{}{}).
		Find(&department)
	if db.Error != nil {
		return department, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return department, nil
}

// GetAllDepartmentsBySequence 获取排序后的所有的部门
func (p *Department) GetAllDepartmentsBySequence() (department []*entity.DepartmentEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(p.ctx).Table(TableNameDepartment).
		Order(fmt.Sprintf("%s ASC", "sequence")).
		Find(&department)
	if db.Error != nil {
		return department, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return department, nil
}

// GetDepartmentsByParentId 查找上级部门
func (p *Department) GetDepartmentsByParentId(parentId int64) (departments []*entity.DepartmentEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(p.ctx).Table(TableNameDepartment).
		Where("parent_id = ?", parentId).
		Find(&departments)
	if db.Error != nil {
		return departments, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return departments, nil
}

// CountDepartments 部门总数
func (p *Department) CountDepartments() (count int64, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(p.ctx).Table(TableNameDepartment).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// GetDepartmentsContainParentId 获取所有的包含parent_id的子部门
func (p *Department) GetDepartmentsContainParentId(parentId int64) (department []*entity.DepartmentEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(p.ctx).Table(TableNameDepartment).
		Where("parent_ids LIKE ?", "%"+fmt.Sprintf("%d", parentId)+"%").
		Find(&department)
	if db.Error != nil {
		return department, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return department, nil
}

// DeleteDepartment 通过部门ID删除部门
func (p *Department) DeleteDepartment(departmentId int64) errors.BizError {
	db := GetDB(dbNameMK).WithContext(p.ctx).Table(TableNameDepartment).
		Where(DepartmentPrimaryKey, departmentId).
		Delete(entity.DepartmentEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}
