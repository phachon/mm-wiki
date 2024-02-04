package dao

import (
	"context"
	"time"

	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/utils"
	"gorm.io/gorm"
)

const (
	// TableNameRolePrivilege 角色权限关系表
	TableNameRolePrivilege = "hms_role_privilege"
	// RolePrivilegePrimaryKey 角色权限关系表主键ID
	RolePrivilegePrimaryKey = "role_privilege_id"
)

// RolePrivilege 角色权限关系表数据
type RolePrivilege struct {
	ctx context.Context
	db  *gorm.DB
}

// NewRolePrivilege 创建系统账号表数据对象
func NewRolePrivilege(ctx context.Context) *RolePrivilege {
	return &RolePrivilege{
		ctx: ctx,
		db:  GetDB(dbNameKms).WithContext(ctx),
	}
}

// Insert 创建账号插入一条账号记录
func (ar *RolePrivilege) Insert(rolePrivilegeEntity *entity.RolePrivilegeEntity) errors.BizError {
	rolePrivilegeEntity.CreateTime = utils.NewJsonTime(time.Now())
	db := ar.db.Table(TableNameRolePrivilege).
		Save(rolePrivilegeEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// DeleteByPrivilegeId 通过权限ID删除
func (ar *RolePrivilege) DeleteByPrivilegeId(privilegeId int64) errors.BizError {
	db := ar.db.
		Table(TableNameRolePrivilege).
		Where("privilege_id", privilegeId).
		Delete(entity.RolePrivilegeEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}

// DeletePrivilegeIdsByAndRoleId 通过角色ID删除部分权限
func (ar *RolePrivilege) DeletePrivilegeIdsByRoleId(roleId int64, privilegeIds []int64) errors.BizError {
	db := ar.db.
		Table(TableNameRolePrivilege).
		Where("role_id", roleId).
		Where("privilege_id IN ?", privilegeIds).
		Delete(entity.RolePrivilegeEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}

// DeleteByRoleId 通过角色ID删除
func (ar *RolePrivilege) DeleteByRoleId(roleId int64) errors.BizError {
	db := ar.db.
		Table(TableNameRolePrivilege).
		Where("role_id", roleId).
		Delete(entity.RolePrivilegeEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}

// GetRolePrivilegesByRoleId 根据角色ID获取角色账号关系
func (ar *RolePrivilege) GetRolePrivilegesByRoleId(roleId int64) (rolePrivileges []*entity.RolePrivilegeEntity,
	err errors.BizError) {
	db := ar.db.Table(TableNameRolePrivilege).
		Where("role_id = ?", roleId).
		Find(&rolePrivileges)
	if db.Error != nil {
		return rolePrivileges, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return rolePrivileges, nil
}

// GetRolePrivilegesByRoleIds 根据多个角色ID获取角色账号关系
func (ar *RolePrivilege) GetRolePrivilegesByRoleIds(roleIds []int64) (rolePrivileges []*entity.RolePrivilegeEntity,
	err errors.BizError) {
	db := ar.db.Table(TableNameRolePrivilege).
		Where("role_id IN ?", roleIds).
		Find(&rolePrivileges)
	if db.Error != nil {
		return rolePrivileges, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return rolePrivileges, nil
}
