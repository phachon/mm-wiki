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
	// TableNameContentVersion 文档内容版本表
	TableNameContentVersion = "mk_content_version"
	// ContentVersionPrimaryKey 内容版本表主键ID
	ContentVersionPrimaryKey = "content_version_id"
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

// GetContentVersionsByDocIdLimit 获取文档多个版本信息(不含正文)
func (c *ContentVersion) GetContentVersionsByDocIdLimit(docId int64, limit int, offset int) (
	[]*entity.ContentVersionEntity, errors.BizError) {

	var contentVersions []*entity.ContentVersionEntity
	db := GetDB(dbNameMK).WithContext(c.ctx).
		Table(TableNameContentVersion).
		Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", ContentVersionPrimaryKey)).
		Where("doc_id = ?", docId).
		Select(
			ContentVersionPrimaryKey,
			"doc_id",
			"create_time",
			"update_time",
			"edit_account_id",
			"edit_account_name",
		).
		Find(&contentVersions)
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return contentVersions, nil
}

// CountContentVersions 获取文档版本数量
func (c *ContentVersion) CountContentVersions(docId int64) (count int64, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(c.ctx).Table(TableNameContentVersion).
		Where("doc_id = ?", docId).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// GetContentVersionByVersionId 获取文档版本信息
func (c *ContentVersion) GetContentVersionByVersionId(versionId int64) (*entity.ContentVersionEntity, errors.BizError) {
	var contentVersion = &entity.ContentVersionEntity{}
	db := GetDB(dbNameMK).WithContext(c.ctx).
		Table(TableNameContentVersion).
		Where("content_version_id = ?", versionId).
		First(contentVersion)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return contentVersion, nil
}

// DeleteContentVersion 删除文档版本
func (c *ContentVersion) DeleteContentVersion(versionId int64) errors.BizError {
	db := GetDB(dbNameMK).WithContext(c.ctx).
		Table(TableNameContentVersion).
		Where("content_version_id = ?", versionId).
		Delete(&entity.ContentVersionEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}

// DeleteContentVersionsByDocId 删除文档所有版本
func (c *ContentVersion) DeleteContentVersionsByDocId(docId int64) errors.BizError {
	db := GetDB(dbNameMK).WithContext(c.ctx).
		Table(TableNameContentVersion).
		Where("doc_id = ?", docId).
		Delete(&entity.ContentVersionEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}

// DeleteContentVersionsNotInVersionIds 删除不在版本ID列表中的版本
func (c *ContentVersion) DeleteContentVersionsNotInVersionIds(docId int64, versionIds []int64) errors.BizError {
	db := GetDB(dbNameMK).WithContext(c.ctx).
		Table(TableNameContentVersion).
		Where("doc_id = ?", docId).
		Where("content_version_id NOT IN (?)", versionIds).
		Delete(&entity.ContentVersionEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}
