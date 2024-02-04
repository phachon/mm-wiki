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
	// TableNameRole 系统角色表
	TableNameRole = "hms_role"
	// RolePrimaryKey 角色表主键ID
	RolePrimaryKey = "role_id"
)

// Role 系统角色表数据
type Role struct {
	ctx context.Context
}

// NewRole 创建系统角色表数据对象
func NewRole(ctx context.Context) *Role {
	return &Role{
		ctx: ctx,
	}
}

// Insert 创建角色插入一条角色记录
func (r *Role) Insert(roleEntity *entity.RoleEntity) errors.BizError {
	roleEntity.CreateTime = utils.NewJsonTime(time.Now())
	roleEntity.UpdateTime = utils.NewJsonTime(time.Now())
	roleEntity.Status = entity.RoleStatusDefault
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameRole).Save(roleEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetRoleByName 根据角色名查找正常的角色
func (r *Role) GetRoleByName(roleName string) (role *entity.RoleEntity, err errors.BizError) {
	role = &entity.RoleEntity{}
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameRole).
		Where(map[string]interface{}{
			"name":   roleName,
			"status": entity.RoleStatusDefault,
		}).
		First(&role)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return role, nil
}

// GetRoleByRoleId 根据角色ID获取角色信息
func (r *Role) GetRoleByRoleId(roleId int64) (role *entity.RoleEntity, err errors.BizError) {

	role = &entity.RoleEntity{}
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameRole).
		Where(map[string]interface{}{
			RolePrimaryKey: roleId,
			"status":       entity.RoleStatusDefault,
		}).
		First(&role)

	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return role, nil
}

// GetRolesByRoleIds 根据多个角色ID批量获取角色ID
func (r *Role) GetRolesByRoleIds(roleIds []int64) (roles []*entity.RoleEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameRole).
		Where("role_id IN (?) and status=?",
			roleIds, entity.RoleStatusDefault).
		Find(&roles)
	if db.Error != nil {
		return roles, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return roles, nil
}

// CountByName 根据角色名获取角色数量
func (r *Role) CountByName(name string) (count int64, err error) {
	if name == "" {
		return 0, nil
	}
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameRole).
		Where(map[string]interface{}{
			"name":   name,
			"status": entity.RoleStatusDefault,
		}).Count(&count)
	if db.Error != nil {
		return 0, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// CheckNameExists 检查角色名是否存在
func (r *Role) CheckNameExists(name string) (bool, error) {
	count, err := r.CountByName(name)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// HasSameName 角色ID和角色名是否存在
func (r *Role) HasSameName(roleId int64, name string) (has bool, err errors.BizError) {
	var count int64
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameRole).
		Where("name = ?", name).
		Where("status = ?", entity.RoleStatusDefault).
		Where("role_id <> ?", roleId).
		Count(&count)
	if db.Error != nil {
		return false, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count > 0, nil
}

// Update 更新角色只会更新如下字段
func (r *Role) Update(role entity.RoleEntity) errors.BizError {
	role.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameRole).
		Select("Name", "Remark", "RoleType").
		Where(map[string]interface{}{
			RolePrimaryKey: role.RoleId,
			"status":       entity.RoleStatusDefault,
		}).
		Updates(role)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// UpdatePrivilegeIds 更新角色下权限
func (r *Role) UpdatePrivilegeIds(role entity.RoleEntity) errors.BizError {
	role.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameRole).
		Select("PrivilegeIds").
		Where(map[string]interface{}{
			RolePrimaryKey: role.RoleId,
			"status":       entity.RoleStatusDefault,
		}).
		Updates(role)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// UpdateStatus 更新角色状态
func (r *Role) DeleteRole(roleId int64) errors.BizError {
	updateTime := utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameRole).
		Where(map[string]interface{}{
			RolePrimaryKey: roleId,
			"status":       entity.RoleStatusDefault,
		}).
		Update("status", entity.RoleStatusDelete).
		Update("update_time", updateTime)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// GetAllRoles 获取所有的角色
func (r *Role) GetAllRoles() (role []*entity.RoleEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameRole).
		Where(map[string]interface{}{
			"status": entity.RoleStatusDefault,
		}).
		Find(&role)
	if db.Error != nil {
		return role, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return role, nil
}

// GetRolesByLimit 分页获取角色列表
func (r *Role) GetRolesByLimit(limit int, offset int) (roles []*entity.RoleEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameRole).
		Where("status", entity.RoleStatusDefault).
		Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", RolePrimaryKey)).
		Find(&roles)
	if db.Error != nil {
		return roles, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return roles, nil
}

// CountRoles 获取角色总数
func (r *Role) CountRoles() (count int64, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameRole).
		Where("status", entity.RoleStatusDefault).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// GetRolesByKeywordsAndLimit 根据关键字分页获取角色列表
func (r *Role) GetRolesByKeywordsAndLimit(limit int, offset int, keywords *entity.RoleKeywords) (roles []*entity.RoleEntity, err errors.BizError) {

	if keywords == nil {
		return r.GetRolesByLimit(limit, offset)
	}

	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameRole)
	db.Where("status", entity.RoleStatusDefault)
	if keywords.RoleName != "" {
		db = db.Where("name LIKE ?", "%"+keywords.RoleName+"%")
	}
	db = db.Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", RolePrimaryKey)).
		Find(&roles)
	if db.Error != nil {
		return roles, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return roles, nil
}

// CountRoles 根据关键字获取角色总数
func (r *Role) CountRolesByKeywords(keywords *entity.RoleKeywords) (count int64, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameRole)
	db = db.Where("status = ?", entity.RoleStatusDefault)
	if keywords.RoleName != "" {
		db = db.Where("name LIKE ?", "%"+keywords.RoleName+"%")
	}
	db.Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// GetRolesByPrivilegeId 根据权限ID获取角色列表
func (r *Role) GetRolesByPrivilegeId(privilegeId int64) (roles []*entity.RoleEntity, err errors.BizError) {
	privilegeIdStr := fmt.Sprintf("%d", privilegeId)
	db := GetDB(dbNameKms).WithContext(r.ctx).
		Table(TableNameRole).
		Where("status", entity.RoleStatusDefault).
		Where("privilege_ids LIKE ?", "%"+privilegeIdStr+"%").
		Order(fmt.Sprintf("%s DESC", RolePrimaryKey)).
		Find(&roles)
	if db.Error != nil {
		return roles, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return roles, nil
}
