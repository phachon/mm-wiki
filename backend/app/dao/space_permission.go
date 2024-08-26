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
	// TableNameSpacePermission 空间权限关系表
	TableNameSpacePermission = "mk_space_permission"
	// SpacePermissionPrimaryKey 空间权限关系表主键ID
	SpacePermissionPrimaryKey = "space_permission_id"
)

// SpacePermission 空间权限关系表数据
type SpacePermission struct {
	ctx context.Context
	db  *gorm.DB
}

// NewSpacePermission 创建空间权限表数据对象
func NewSpacePermission(ctx context.Context) *SpacePermission {
	return &SpacePermission{
		ctx: ctx,
		db:  GetDB(dbNameMK).WithContext(ctx),
	}
}

// Insert 插入一条记录
func (ar *SpacePermission) Insert(spacePermissionEntity *entity.SpacePermissionEntity) errors.BizError {
	spacePermissionEntity.CreateTime = utils.NewJsonTime(time.Now())
	spacePermissionEntity.UpdateTime = utils.NewJsonTime(time.Now())
	db := ar.db.Table(TableNameSpacePermission).
		Save(spacePermissionEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// DeleteBySpaceIdAccountId 通过空间ID和账号ID删除
func (ar *SpacePermission) DeleteBySpaceIdAccountId(spaceId int64, accountId int64) errors.BizError {
	db := ar.db.
		Table(TableNameSpacePermission).
		Where("space_id", spaceId).
		Where("account_id", accountId).
		Delete(entity.SpacePermissionEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}

// DeleteByAccountId 通过账号ID删除
func (ar *SpacePermission) DeleteByAccountId(accountId int64) errors.BizError {
	db := ar.db.
		Table(TableNameSpacePermission).
		Where("account_id", accountId).
		Delete(entity.SpacePermissionEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}

// DeleteBySpaceId 通过空间ID删除
func (ar *SpacePermission) DeleteBySpaceId(spaceId int64) errors.BizError {
	db := ar.db.
		Table(TableNameSpacePermission).
		Where("space_id", spaceId).
		Delete(entity.SpacePermissionEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}

// GetPermissionsByAccountId 根据账号ID获取空间权限
func (ar *SpacePermission) GetPermissionsByAccountId(accountId int64) (
	spacePermissions []*entity.SpacePermissionEntity,
	err errors.BizError) {
	db := ar.db.Table(TableNameSpacePermission).
		Where("account_id = ?", accountId).
		Find(&spacePermissions)
	if db.Error != nil {
		return spacePermissions, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return spacePermissions, nil
}

// GetPermissionsByAccountIds 根据多个账号ID获取角色账号关系
func (ar *SpacePermission) GetPermissionsByAccountIds(accountIds []int64) (spacePermissions []*entity.SpacePermissionEntity,
	err errors.BizError) {
	db := ar.db.Table(TableNameSpacePermission).
		Where("account_id IN (?)", accountIds).
		Find(&spacePermissions)
	if db.Error != nil {
		return spacePermissions, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return spacePermissions, nil
}

// GetPermissionsBySpaceId 根据空间ID查找空间权限
func (ar *SpacePermission) GetPermissionsBySpaceId(spaceId int64) (spacePermissions []*entity.SpacePermissionEntity,
	err errors.BizError) {
	db := ar.db.Table(TableNameSpacePermission).
		Where("space_id = ?", spaceId).
		Find(&spacePermissions)
	if db.Error != nil {
		return spacePermissions, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return spacePermissions, nil
}

// GetPermissionsBySpaceIdLimit 分页获取空间权限
func (ar *SpacePermission) GetPermissionsBySpaceIdLimit(limit int, offset int, spaceId int64) (spacePermissions []*entity.SpacePermissionEntity,
	err errors.BizError) {

	db := ar.db.Table(TableNameSpacePermission).
		Where("space_id = ?", spaceId).
		Limit(limit).
		Offset(offset).
		Find(&spacePermissions)
	if db.Error != nil {
		return spacePermissions, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return spacePermissions, nil
}

// CountPermissionsBySpaceId 根据空间 ID 获取权限总数
func (ar *SpacePermission) CountPermissionsBySpaceId(spaceId int64) (count int64, err errors.BizError) {
	db := ar.db.Table(TableNameSpacePermission).
		Where("space_id = ?", spaceId).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// GetPermissionByAccountIdAndSpaceId 根据账号ID和空间ID获取权限
func (ar *SpacePermission) GetPermissionByAccountIdAndSpaceId(accountId int64, spaceId int64) (
	spacePermission *entity.SpacePermissionEntity, err errors.BizError) {
	db := ar.db.Table(TableNameSpacePermission).
		Where("account_id = ?", accountId).
		Where("space_id = ?", spaceId).
		First(&spacePermission)
	if db.Error != nil {
		return spacePermission, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return spacePermission, nil
}

// GetPermissionByDepartmentIdAndSpaceId 根据部门ID和空间ID获取权限
func (ar *SpacePermission) GetPermissionByDepartmentIdAndSpaceId(departmentId int64, spaceId int64) (
	spacePermission *entity.SpacePermissionEntity, err errors.BizError) {
	db := ar.db.Table(TableNameSpacePermission).
		Where("department_id = ?", departmentId).
		Where("space_id = ?", spaceId).
		First(&spacePermission)
	if db.Error != nil {
		return spacePermission, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return spacePermission, nil
}

// UpdatePermission 更新权限
func (ar *SpacePermission) UpdatePermission(spacePermissionEntity *entity.SpacePermissionEntity) errors.BizError {
	spacePermissionEntity.UpdateTime = utils.NewJsonTime(time.Now())
	db := ar.db.Table(TableNameSpacePermission).
		Select("IsView", "IsAdd", "IsEdit", "IsDelete", "IsAdmin").
		Where(SpacePermissionPrimaryKey, spacePermissionEntity.SpacePermissionId).
		Updates(spacePermissionEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}
