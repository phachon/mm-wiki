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
	// TableNameCollection 收藏表
	TableNameCollection = "mk_collection"
)

// Collection 收藏表数据
type Collection struct {
	ctx context.Context
}

// NewCollection 创建收藏表数据对象
func NewCollection(ctx context.Context) *Collection {
	return &Collection{
		ctx: ctx,
	}
}

// Insert 创建收藏插入一条收藏记录
func (kc *Collection) Insert(collectionEntity *entity.CollectionEntity) errors.BizError {
	collectionEntity.CreateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(kc.ctx).Table(TableNameCollection).Save(collectionEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetCollectionByAccountIdAndResourceId 根据账号ID和资源ID查找收藏
func (kc *Collection) GetCollectionByAccountIdAndResourceId(accountId int64, resourceId string) (collection *entity.CollectionEntity, err errors.BizError) {
	collection = &entity.CollectionEntity{}
	db := GetDB(dbNameMK).WithContext(kc.ctx).Table(TableNameCollection).
		Where(map[string]interface{}{
			"account_id":  accountId,
			"resource_id": resourceId,
		}).
		First(&collection)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return collection, nil
}

// DeleteCollectionByAccountIdAndResourceId 根据账号ID和资源ID删除收藏
func (kc *Collection) DeleteCollectionByAccountIdAndResourceId(accountId int64, resourceId string) errors.BizError {
	db := GetDB(dbNameMK).WithContext(kc.ctx).Table(TableNameCollection).
		Where(map[string]interface{}{
			"account_id":  accountId,
			"resource_id": resourceId,
		}).
		Delete(&entity.CollectionEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}

// GetCollectionsByAccountIdAndType 根据账号ID和类型查找收藏
func (kc *Collection) GetCollectionsByAccountIdAndType(accountId int64, collectionType int) (collections []*entity.CollectionEntity, err errors.BizError) {
	collections = make([]*entity.CollectionEntity, 0)
	db := GetDB(dbNameMK).WithContext(kc.ctx).Table(TableNameCollection).
		Where(map[string]interface{}{
			"account_id":      accountId,
			"collection_type": collectionType,
		}).
		Find(&collections)
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return collections, nil
}

// GetCollectionByAccountIdAndTypeAndResourceId 根据账号ID和类型和资源ID查找收藏
func (kc *Collection) GetCollectionByAccountIdAndTypeAndResourceId(accountId int64, collectionType int, resourceId string) (*entity.CollectionEntity, errors.BizError) {
	collection := new(entity.CollectionEntity)
	db := GetDB(dbNameMK).WithContext(kc.ctx).Table(TableNameCollection).
		Where(map[string]interface{}{
			"account_id":      accountId,
			"collection_type": collectionType,
			"resource_id":     resourceId,
		}).
		First(&collection)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return collection, nil
}

// DeleteCollectionByAccountIdAndCollectionId 根据账号ID和收藏ID删除收藏
func (kc *Collection) DeleteCollectionByAccountIdAndCollectionId(accountId int64, collectionId int64) errors.BizError {
	db := GetDB(dbNameMK).WithContext(kc.ctx).Table(TableNameCollection).
		Where(map[string]interface{}{
			"account_id":    accountId,
			"collection_id": collectionId,
		}).
		Delete(&entity.CollectionEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}

// DeleteByResourceIdAndType 根据资源ID和类型删除所有收藏
func (kc *Collection) DeleteByResourceIdAndType(collectionType int, resourceId string) errors.BizError {
	db := GetDB(dbNameMK).WithContext(kc.ctx).Table(TableNameCollection).
		Where(map[string]interface{}{
			"collection_type": collectionType,
			"resource_id":     resourceId,
		}).
		Delete(&entity.CollectionEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}

// DeleteByAccountIdResourceAndType 根据账号ID删除收藏资源
func (kc *Collection) DeleteByAccountIdResourceAndType(accountId int64, collectionType int, resourceId string) errors.BizError {
	db := GetDB(dbNameMK).WithContext(kc.ctx).Table(TableNameCollection).
		Where(map[string]interface{}{
			"account_id":      accountId,
			"collection_type": collectionType,
			"resource_id":     resourceId,
		}).
		Delete(&entity.CollectionEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}
