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
	// TableNameFollow 关注表
	TableNameFollow = "mk_follow"
)

// Follow 关注表数据
type Follow struct {
	ctx context.Context
}

// NewFollow 创建关注表数据对象
func NewFollow(ctx context.Context) *Follow {
	return &Follow{
		ctx: ctx,
	}
}

// Insert 插入一条关注记录
func (f *Follow) Insert(followEntity *entity.FollowEntity) errors.BizError {
	followEntity.CreateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(f.ctx).Table(TableNameFollow).Save(followEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetFollowByAccountIdAndTypeAndObjectId 根据账号ID、类型和对象ID查找关注
func (f *Follow) GetFollowByAccountIdAndTypeAndObjectId(accountId int64, followType int, objectId string) (*entity.FollowEntity, errors.BizError) {
	follow := new(entity.FollowEntity)
	db := GetDB(dbNameMK).WithContext(f.ctx).Table(TableNameFollow).
		Where("account_id = ? AND follow_type = ? AND object_id = ?", accountId, followType, objectId).
		First(follow)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return follow, nil
}

// DeleteByObjectIdAndType 根据对象ID和类型删除所有关注
func (f *Follow) DeleteByObjectIdAndType(followType int, objectId string) errors.BizError {
	db := GetDB(dbNameMK).WithContext(f.ctx).Table(TableNameFollow).
		Where("follow_type = ? AND object_id = ?", followType, objectId).
		Delete(&entity.FollowEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}

// GetFollowsByAccountIdAndType 根据账号ID和类型分页获取关注列表
func (f *Follow) GetFollowsByAccountIdAndType(accountId int64, followType int, limit int, offset int) ([]*entity.FollowEntity, errors.BizError) {
	var follows []*entity.FollowEntity
	db := GetDB(dbNameMK).WithContext(f.ctx).Table(TableNameFollow).
		Where("account_id = ? AND follow_type = ?", accountId, followType).
		Limit(limit).
		Offset(offset).
		Order("follow_id DESC").
		Find(&follows)
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return follows, nil
}

// CountFollowsByAccountIdAndType 根据账号ID和类型获取关注总数
func (f *Follow) CountFollowsByAccountIdAndType(accountId int64, followType int) (int64, errors.BizError) {
	var count int64
	db := GetDB(dbNameMK).WithContext(f.ctx).Table(TableNameFollow).
		Where("account_id = ? AND follow_type = ?", accountId, followType).
		Count(&count)
	if db.Error != nil {
		return 0, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// DeleteByAccountIdTypeAndObjectId 根据账号ID、类型和对象ID删除关注
func (f *Follow) DeleteByAccountIdTypeAndObjectId(accountId int64, followType int, objectId string) errors.BizError {
	db := GetDB(dbNameMK).WithContext(f.ctx).Table(TableNameFollow).
		Where("account_id = ? AND follow_type = ? AND object_id = ?", accountId, followType, objectId).
		Delete(&entity.FollowEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}
