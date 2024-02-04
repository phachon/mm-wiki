package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
	"github.com/phachon/mm-wiki/utils"
	"gorm.io/gorm"
)

const (
	// TableNamePrivilege 权限动作表
	TableNamePrivilege = "hms_privilege"
	// PrivilegePrimaryKey 权限表主键ID
	PrivilegePrimaryKey = "privilege_id"
)

// Privilege 权限动作表数据
type Privilege struct {
	ctx context.Context
}

// NewPrivilege 创建权限动作表数据对象
func NewPrivilege(ctx context.Context) *Privilege {
	return &Privilege{
		ctx: ctx,
	}
}

// Insert 创建权限插入一条权限记录
func (p *Privilege) Insert(privilegeEntity *entity.PrivilegeEntity) errors.BizError {
	privilegeEntity.CreateTime = utils.NewJsonTime(time.Now())
	privilegeEntity.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(p.ctx).Table(TableNamePrivilege).Save(privilegeEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetPrivilegeByName 根据权限名查找正常的权限
func (p *Privilege) GetPrivilegeByName(privilegeName string) (privilege *entity.PrivilegeEntity, err errors.BizError) {
	privilege = &entity.PrivilegeEntity{}
	db := GetDB(dbNameKms).WithContext(p.ctx).Table(TableNamePrivilege).
		Where(map[string]interface{}{
			"name": privilegeName,
		}).
		First(&privilege)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return privilege, nil
}

// GetPrivilegeByIdentify 根据权限标识查找权限
func (p *Privilege) GetPrivilegeByIdentify(identify string) (privilege *entity.PrivilegeEntity, err errors.BizError) {
	privilege = &entity.PrivilegeEntity{}
	db := GetDB(dbNameKms).WithContext(p.ctx).Table(TableNamePrivilege).
		Where(map[string]interface{}{
			"identify": identify,
		}).
		First(&privilege)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return privilege, nil
}

// GetPrivilegeByType 根据权限类型查找权限
func (p *Privilege) GetPrivilegeByType(privilegeType int) (privileges []*entity.PrivilegeEntity, err errors.BizError) {
	privileges = []*entity.PrivilegeEntity{}
	db := GetDB(dbNameKms).WithContext(p.ctx).Table(TableNamePrivilege).
		Where(map[string]interface{}{
			"privilege_type": privilegeType,
		}).
		Order(fmt.Sprintf("%s ASC", "sequence")).
		Find(&privileges)
	if db.Error == gorm.ErrRecordNotFound {
		return privileges, nil
	}
	if db.Error != nil {
		return privileges, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return privileges, nil
}

// GetPrivilegeByPrivilegeId 根据权限ID获取权限信息
func (p *Privilege) GetPrivilegeByPrivilegeId(privilegeId int64) (privilege *entity.PrivilegeEntity, err errors.BizError) {

	privilege = &entity.PrivilegeEntity{}
	db := GetDB(dbNameKms).WithContext(p.ctx).Table(TableNamePrivilege).
		Where(map[string]interface{}{
			PrivilegePrimaryKey: privilegeId,
		}).
		First(&privilege)

	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return privilege, nil
}

// GetPrivilegesByPrivilegeIds 根据多个权限ID批量获取权限ID
func (p *Privilege) GetPrivilegesByPrivilegeIds(privilegeIds []int64) (privileges []*entity.PrivilegeEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(p.ctx).Table(TableNamePrivilege).
		Where("privilege_id IN (?)", privilegeIds).
		Order(fmt.Sprintf("%s ASC", "sequence")).
		Find(&privileges)
	if db.Error != nil {
		return privileges, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return privileges, nil
}

// CountByName 根据权限名获取权限数
func (p *Privilege) CountByName(name string) (count int64, err error) {
	if name == "" {
		return 0, nil
	}
	db := GetDB(dbNameKms).WithContext(p.ctx).Table(TableNamePrivilege).
		Where(map[string]interface{}{
			"name": name,
		}).Count(&count)
	if db.Error != nil {
		return 0, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// CheckNameExists 检查权限名是否存在
func (p *Privilege) CheckNameExists(name string) (bool, error) {
	count, err := p.CountByName(name)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// HasSameName 权限ID和权限名是否存在
func (p *Privilege) HasSameName(privilegeId int64, name string) (has bool, err errors.BizError) {
	var count int64
	db := GetDB(dbNameKms).WithContext(p.ctx).Table(TableNamePrivilege).
		Where("name = ?", name).
		Where("privilege_id <> ?", privilegeId).
		Count(&count)
	if db.Error != nil {
		return false, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count > 0, nil
}

// Update 更新权限，只会更新如下字段
func (p *Privilege) Update(privilege entity.PrivilegeEntity) errors.BizError {
	logger.WithContext(p.ctx).Infof("privilege:%v", privilege)
	privilege.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(p.ctx).Table(TableNamePrivilege).
		Select("Name", "ParentId", "ParentIds", "PrivilegeType", "PageRouter",
			"ApiMarks", "Icon", "IsDisplay", "Sequence").
		Where(map[string]interface{}{
			PrivilegePrimaryKey: privilege.PrivilegeId,
		}).
		Updates(privilege)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// UpdateParentIdsByPrivilegeId 更新 parent_ids
func (p *Privilege) UpdateParentIdsByPrivilegeId(privilegeId int64, parentIds string) errors.BizError {
	db := GetDB(dbNameKms).WithContext(p.ctx).Table(TableNamePrivilege).
		Where(PrivilegePrimaryKey, privilegeId).
		Update("parent_ids", parentIds)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// GetAllPrivilege 获取所有的权限
func (p *Privilege) GetAllPrivilege() (privilege []*entity.PrivilegeEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(p.ctx).Table(TableNamePrivilege).
		Where(map[string]interface{}{}).
		Find(&privilege)
	if db.Error != nil {
		return privilege, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return privilege, nil
}

// GetAllPrivilegesBySequence 获取排序后的所有的权限
func (p *Privilege) GetAllPrivilegesBySequence() (privilege []*entity.PrivilegeEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(p.ctx).Table(TableNamePrivilege).
		Order(fmt.Sprintf("%s ASC", "sequence")).
		Find(&privilege)
	if db.Error != nil {
		return privilege, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return privilege, nil
}

// GetPrivilegesByParentId 查找上级权限
func (p *Privilege) GetPrivilegesByParentId(parentId int64) (privileges []*entity.PrivilegeEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(p.ctx).Table(TableNamePrivilege).
		Where("parent_id = ?", parentId).
		Find(&privileges)
	if db.Error != nil {
		return privileges, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return privileges, nil
}

// CountPrivileges 权限总数
func (p *Privilege) CountPrivileges() (count int64, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(p.ctx).Table(TableNamePrivilege).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// GetPrivilegesContainParentId 获取所有的包含parent_id的子权限
func (p *Privilege) GetPrivilegesContainParentId(parentId int64) (privilege []*entity.PrivilegeEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(p.ctx).Table(TableNamePrivilege).
		Where("parent_ids LIKE ?", "%"+fmt.Sprintf("%d", parentId)+"%").
		Find(&privilege)
	if db.Error != nil {
		return privilege, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return privilege, nil
}

// DeletePrivilege 通过权限ID删除权限
func (p *Privilege) DeletePrivilege(privilegeId int64) errors.BizError {
	db := GetDB(dbNameKms).WithContext(p.ctx).Table(TableNamePrivilege).
		Where("privilege_id", privilegeId).
		Delete(entity.PrivilegeEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}
