package service

import (
	"context"

	"github.com/phachon/mm-wiki/app/dao"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
)

// Collection 收藏业务逻辑
type Collection struct {
	ctx           context.Context
	daoCollection *dao.Collection
}

// NewCollection 创建收藏业务逻辑对象
func NewCollection(ctx context.Context) *Collection {
	return &Collection{
		ctx:           ctx,
		daoCollection: dao.NewCollection(ctx),
	}
}

// Create 创建收藏
func (c *Collection) Create(collectionEntity *entity.CollectionEntity) errors.BizError {
	// 查找是否收藏
	collection, err := c.daoCollection.GetCollectionByAccountIdAndTypeAndResourceId(
		collectionEntity.AccountId,
		collectionEntity.CollectionType,
		collectionEntity.ResourceId,
	)
	if err != nil {
		return err
	}
	if collection != nil {
		return errors.Errorf(errors.BusinessRecordExistError, "已经收藏过")
	}
	// 插入一条记录
	err = c.daoCollection.Insert(collectionEntity)
	if err != nil {
		return err
	}
	return nil
}

// CreateDocumentCollection 收藏文档
func (c *Collection) CreateDocumentCollection(accountId int64, documentId string) errors.BizError {
	collectionEntity := &entity.CollectionEntity{
		AccountId:      accountId,
		CollectionType: entity.CollectionTypeDocument,
		ResourceId:     documentId,
	}
	return c.Create(collectionEntity)
}

// CreateSpaceCollection 收藏空间
func (c *Collection) CreateSpaceCollection(accountId int64, spaceId string) errors.BizError {
	collectionEntity := &entity.CollectionEntity{
		AccountId:      accountId,
		CollectionType: entity.CollectionTypeSpace,
		ResourceId:     spaceId,
	}
	return c.Create(collectionEntity)
}

// Delete 删除收藏
func (c *Collection) Delete(accountId int64, collectionId int64) errors.BizError {
	if accountId <= 0 || collectionId <= 0 {
		return nil
	}
	err := c.daoCollection.DeleteCollectionByAccountIdAndCollectionId(accountId, collectionId)
	if err != nil {
		return err
	}
	return nil
}

// GetAccountCollectionAllSpace 获取用户收藏的所有空间
func (c *Collection) GetAccountCollectionAllSpace(accountId int64) ([]*entity.CollectionEntity, errors.BizError) {
	if accountId <= 0 {
		return nil, nil
	}
	collections, err := c.daoCollection.GetCollectionsByAccountIdAndType(accountId, entity.CollectionTypeSpace)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

// GetAccountCollectionAllDocs 获取用户收藏的所有文档
func (c *Collection) GetAccountCollectionAllDocs(accountId int64) ([]*entity.CollectionEntity, errors.BizError) {
	if accountId <= 0 {
		return nil, nil
	}
	collections, err := c.daoCollection.GetCollectionsByAccountIdAndType(accountId, entity.CollectionTypeDocument)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

// DeleteByAccountIdResourceAndType 删除账号下的资源ID
func (c *Collection) DeleteByAccountIdResourceAndType(accountId int64, collectionType int, resourceId string) errors.BizError {
	if accountId <= 0 || resourceId == "" {
		return nil
	}
	err := c.daoCollection.DeleteByAccountIdResourceAndType(accountId, collectionType, resourceId)
	if err != nil {
		return err
	}
	return nil
}

// GetCollectionByAccountIdAndTypeAndResourceId 获取收藏信息
func (c *Collection) GetCollectionByAccountIdAndTypeAndResourceId(accountId int64, collectionType int, resourceId string) (*entity.CollectionEntity, errors.BizError) {
	if accountId <= 0 || resourceId == "" {
		return nil, nil
	}
	collection, err := c.daoCollection.GetCollectionByAccountIdAndTypeAndResourceId(accountId, collectionType, resourceId)
	if err != nil {
		return nil, err
	}
	return collection, nil
}
