package dao

import (
	"context"
	"time"

	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/utils"
)

const (
	// TableNameContentVersion 文档内容版本表
	TableNameContentVersion = "mk_content_version"
)

// ContentVersion 文档内容版本表数据对象
type ContentVersion struct {
	ctx context.Context
}

// NewContentVersion 创建文档内容版本表数据对象
func NewContentVersion(ctx context.Context) *ContentVersion {
	return &ContentVersion{
		ctx: ctx,
	}
}

// Insert 创建文档内容版本插入一条文档内容版本记录
func (c *ContentVersion) Insert(contentVersionEntity *entity.ContentVersionEntity) errors.BizError {
	contentVersionEntity.CreateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(c.ctx).Table(TableNameContentVersion).Save(contentVersionEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetContentVersionsByDocId 获取文档多个版本信息(不含正文)
func (c *ContentVersion) GetContentVersionsByDocId(docId int64) ([]*entity.ContentVersionEntity, errors.BizError) {
	var contentVersions []*entity.ContentVersionEntity
	db := GetDB(dbNameMK).WithContext(c.ctx).
		Table(TableNameContentVersion).
		Where("doc_id = ?", docId).
		Select("version_id", "doc_id", "version_number", "create_time").
		Find(&contentVersions)
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return contentVersions, nil
}

// GetContentVersionByDocIdAndVersionId 获取文档版本信息
func (c *ContentVersion) GetContentVersionByDocIdAndVersionId(docId, versionId int64) (*entity.ContentVersionEntity, errors.BizError) {
	var contentVersion = &entity.ContentVersionEntity{}
	db := GetDB(dbNameMK).WithContext(c.ctx).
		Table(TableNameContentVersion).
		Where("doc_id = ? AND version_id = ?", docId, versionId).
		First(contentVersion)
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return contentVersion, nil
}

// GetContentVersionByDocIdAndVersionNumber 获取文档版本信息
func (c *ContentVersion) GetContentVersionByDocIdAndVersionNumber(docId int64, versionNumber int) (*entity.ContentVersionEntity, errors.BizError) {
	var contentVersion = &entity.ContentVersionEntity{}
	db := GetDB(dbNameMK).WithContext(c.ctx).
		Table(TableNameContentVersion).
		Where("doc_id = ? AND version_number = ?", docId, versionNumber).
		First(contentVersion)
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return contentVersion, nil
}
