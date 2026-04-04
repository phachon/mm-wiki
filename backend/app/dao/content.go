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
	// TableNameContent 正文表
	TableNameContent = "mk_content"
)

// Content 正文表数据
type Content struct {
	ctx context.Context
}

// NewContent 创建正文表数据对象
func NewContent(ctx context.Context) *Content {
	return &Content{
		ctx: ctx,
	}
}

// Insert 创建正文插入一条正文记录
func (c *Content) Insert(contentEntity *entity.ContentEntity) errors.BizError {
	contentEntity.CreateTime = utils.NewJsonTime(time.Now())
	contentEntity.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(c.ctx).Table(TableNameContent).Save(contentEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetContentByDocId 获取文档正文信息
func (c *Content) GetContentByDocId(docId int64) (*entity.ContentEntity, errors.BizError) {
	var content = &entity.ContentEntity{}
	db := GetDB(dbNameMK).WithContext(c.ctx).
		Table(TableNameContent).
		Where("doc_id = ?", docId).
		First(content)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return content, nil
}

// UpdateContent 更新文档正文信息
func (c *Content) UpdateContent(content *entity.ContentEntity) errors.BizError {
	content.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(c.ctx).Table(TableNameContent).
		Select("Content", "CurrentVersionID", "UpdateTime").
		Where(map[string]interface{}{
			"doc_id": content.DocId,
		}).
		Updates(content)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// DeleteContentByDocId 根据文档ID删除正文
func (c *Content) DeleteContentByDocId(docId int64) errors.BizError {
	db := GetDB(dbNameMK).WithContext(c.ctx).Table(TableNameContent).
		Where("doc_id = ?", docId).
		Delete(&entity.ContentEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}

// UpdateCurrentVersion 更新文档当前版本
func (c *Content) UpdateCurrentVersion(docId, versionId int64) errors.BizError {
	db := GetDB(dbNameMK).WithContext(c.ctx).Table(TableNameContent).
		Select("CurrentVersionID").
		Where(map[string]interface{}{
			"doc_id": docId,
		}).
		Updates(map[string]interface{}{
			"CurrentVersionID": versionId,
		})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}
