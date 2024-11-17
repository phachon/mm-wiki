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
	// TableNameSpace 系统空间表
	TableNameSpace = "mk_space"
	// SpacePrimaryKey 空间表主键ID
	SpacePrimaryKey = "space_id"
)

// Space 系统空间表数据
type Space struct {
	ctx context.Context
}

// NewSpace 创建系统空间表数据对象
func NewSpace(ctx context.Context) *Space {
	return &Space{
		ctx: ctx,
	}
}

// Insert 创建空间插入一条空间记录
func (r *Space) Insert(spaceEntity *entity.SpaceEntity) errors.BizError {
	spaceEntity.CreateTime = utils.NewJsonTime(time.Now())
	spaceEntity.UpdateTime = utils.NewJsonTime(time.Now())
	spaceEntity.Status = entity.SpaceStatusDefault
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameSpace).Save(spaceEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetSpaceByName 根据空间名查找正常的空间
func (r *Space) GetSpaceByName(spaceName string) (space *entity.SpaceEntity, err errors.BizError) {
	space = &entity.SpaceEntity{}
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameSpace).
		Where(map[string]interface{}{
			"status": entity.SpaceStatusDefault,
			"name":   spaceName,
		}).
		First(&space)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return space, nil
}

// GetSpaceBySpaceKey 根据空间key查找正常的空间
func (r *Space) GetSpaceBySpaceKey(spaceKey string) (space *entity.SpaceEntity, err errors.BizError) {
	space = &entity.SpaceEntity{}
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameSpace).
		Where(map[string]interface{}{
			"status":    entity.SpaceStatusDefault,
			"space_key": spaceKey,
		}).
		First(&space)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return space, nil
}

// GetSpaceBySpaceId 根据空间ID获取空间信息
func (r *Space) GetSpaceBySpaceId(spaceId int64) (space *entity.SpaceEntity, err errors.BizError) {

	space = &entity.SpaceEntity{}
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameSpace).
		Where(map[string]interface{}{
			SpacePrimaryKey: spaceId,
			"status":        entity.SpaceStatusDefault,
		}).
		First(&space)

	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return space, nil
}

// GetSpacesBySpaceIds 根据多个空间ID批量获取空间ID
func (r *Space) GetSpacesBySpaceIds(spaceIds []int64) (spaces []*entity.SpaceEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameSpace).
		Where("status=? and space_id IN (?)", entity.SpaceStatusDefault, spaceIds).
		Find(&spaces)
	if db.Error != nil {
		return spaces, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return spaces, nil
}

// CountByName 根据空间名获取空间数量
func (r *Space) CountByName(name string) (count int64, err error) {
	if name == "" {
		return 0, nil
	}
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameSpace).
		Where(map[string]interface{}{
			"name":   name,
			"status": entity.SpaceStatusDefault,
		}).Count(&count)
	if db.Error != nil {
		return 0, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// CheckNameExists 检查空间名是否存在
func (r *Space) CheckNameExists(name string) (bool, error) {
	count, err := r.CountByName(name)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// HasSameName 空间ID和空间名是否存在
func (r *Space) HasSameName(spaceId int64, name string) (has bool, err errors.BizError) {
	var count int64
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameSpace).
		Where("name = ?", name).
		Where("status = ?", entity.SpaceStatusDefault).
		Where("space_id <> ?", spaceId).
		Count(&count)
	if db.Error != nil {
		return false, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count > 0, nil
}

// Update 更新空间只会更新如下字段
func (r *Space) Update(space entity.SpaceEntity) errors.BizError {
	space.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameSpace).
		Select("Name", "Description", "VisitLevel", "IsShare", "IsExport").
		Where(map[string]interface{}{
			SpacePrimaryKey: space.SpaceId,
			"status":        entity.SpaceStatusDefault,
		}).
		Updates(space)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// UpdateStatus 更新空间状态
func (r *Space) DeleteSpace(spaceId int64) errors.BizError {
	updateTime := utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameSpace).
		Where(map[string]interface{}{
			SpacePrimaryKey: spaceId,
			"status":        entity.SpaceStatusDefault,
		}).
		Update("status", entity.SpaceStatusDelete).
		Update("update_time", updateTime)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// GetAllSpaces 获取所有的空间
func (r *Space) GetAllSpaces() (space []*entity.SpaceEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameSpace).
		Where(map[string]interface{}{
			"status": entity.SpaceStatusDefault,
		}).
		Find(&space)
	if db.Error != nil {
		return space, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return space, nil
}

// GetSpacesByLimit 分页获取空间列表
func (r *Space) GetSpacesByLimit(limit int, offset int) (spaces []*entity.SpaceEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameSpace).
		Where("status", entity.SpaceStatusDefault).
		Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", SpacePrimaryKey)).
		Find(&spaces)
	if db.Error != nil {
		return spaces, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return spaces, nil
}

// CountSpaces 获取空间总数
func (r *Space) CountSpaces() (count int64, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameSpace).
		Where("status", entity.SpaceStatusDefault).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// GetSpacesByKeywordsAndLimit 根据关键字分页获取空间列表
func (r *Space) GetSpacesByKeywordsAndLimit(limit int, offset int, keywords *entity.SpaceKeywords) (spaces []*entity.SpaceEntity, err errors.BizError) {

	if keywords == nil {
		return r.GetSpacesByLimit(limit, offset)
	}

	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameSpace)
	db.Where("status", entity.SpaceStatusDefault)
	if keywords.SpaceType != nil {
		db = db.Where("space_type = ?", keywords.SpaceType)
	}
	if keywords.SpaceName != "" {
		db = db.Where("name LIKE ?", "%"+keywords.SpaceName+"%")
	}
	db = db.Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", SpacePrimaryKey)).
		Find(&spaces)
	if db.Error != nil {
		return spaces, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return spaces, nil
}

// CountSpaces 根据关键字获取空间总数
func (r *Space) CountSpacesByKeywords(keywords *entity.SpaceKeywords) (count int64, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameSpace)
	db = db.Where("status = ?", entity.SpaceStatusDefault)
	if keywords.SpaceType != nil {
		db = db.Where("space_type = ?", keywords.SpaceType)
	}
	if keywords.SpaceName != "" {
		db = db.Where("name LIKE ?", "%"+keywords.SpaceName+"%")
	}
	db.Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// GetPublicSpacesByLimit 获取公开空间列表
func (s *Space) GetPublicSpacesByLimit(limit, offset int, keywords *entity.SpaceKeywords) ([]*entity.SpaceEntity, errors.BizError) {
	spaces := make([]*entity.SpaceEntity, 0)
	db := GetDB(dbNameMK).WithContext(s.ctx).Table(TableNameSpace)

	db = db.Where("visit_level = ?", entity.SpaceVisitLevelDefaultPublic)
	db = db.Where("status = ?", entity.SpaceStatusDefault)
	if keywords != nil {
		if keywords.SpaceName != "" {
			db = db.Where("name LIKE ?", "%"+keywords.SpaceName+"%")
		}
	}
	if err := db.Offset(offset).Limit(limit).Find(&spaces).Error; err != nil {
		return spaces, errors.Errorf(errors.DalMysqlSelectErr, err.Error())
	}

	return spaces, nil
}

// CountPublicSpaces 统计公开空间数量
func (d *Space) CountPublicSpaces(keywords *entity.SpaceKeywords) (int64, errors.BizError) {
	var count int64
	db := GetDB(dbNameMK).WithContext(d.ctx).Table(TableNameSpace)

	db = db.Where("visit_level = ?", entity.SpaceVisitLevelDefaultPublic)
	db = db.Where("status = ?", entity.SpaceStatusDefault)
	if keywords != nil {
		if keywords.SpaceName != "" {
			db = db.Where("name LIKE ?", "%"+keywords.SpaceName+"%")
		}
	}

	if err := db.Count(&count).Error; err != nil {
		return 0, errors.Errorf(errors.DalMysqlSelectErr, err.Error())
	}

	return count, nil
}
