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
	// TableNameDoc 文档表
	TableNameDoc = "mk_doc"
	// DocPrimaryKey 文档表主键ID
	DocPrimaryKey = "doc_id"
)

// Doc 文档表数据
type Doc struct {
	ctx context.Context
}

// NewDoc 创建文档表数据对象
func NewDoc(ctx context.Context) *Doc {
	return &Doc{
		ctx: ctx,
	}
}

// Insert 创建文档插入一条文档记录
func (d *Doc) Insert(docEntity *entity.DocEntity) errors.BizError {
	docEntity.CreateTime = utils.NewJsonTime(time.Now())
	docEntity.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(d.ctx).Table(TableNameDoc).Save(docEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetDocsBySpaceId 获取空间下所有文档
func (d *Doc) GetDocsBySpaceId(spaceId int64) ([]*entity.DocEntity, errors.BizError) {
	var docs []*entity.DocEntity
	db := GetDB(dbNameMK).WithContext(d.ctx).
		Table(TableNameDoc).
		Where("space_id = ?", spaceId).
		Find(&docs)
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return docs, nil
}

// GetDocByDocId 获取文档信息
func (d *Doc) GetDocByDocId(docId int64) (*entity.DocEntity, errors.BizError) {
	var doc = &entity.DocEntity{}
	db := GetDB(dbNameMK).WithContext(d.ctx).
		Table(TableNameDoc).
		Where("doc_id = ?", docId).
		First(doc)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return doc, nil
}

// GetDocByDocIds 获取多个文档信息
func (d *Doc) GetDocByDocIds(docIds []int64) ([]*entity.DocEntity, errors.BizError) {
	var docs []*entity.DocEntity
	db := GetDB(dbNameMK).WithContext(d.ctx).
		Table(TableNameDoc).
		Where("doc_id in (?)", docIds).
		Find(&docs)
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return docs, nil
}

// GetDocsBySpaceKey 获取空间下所有文档
func (d *Doc) GetDocsBySpaceKey(spaceKey string) ([]*entity.DocEntity, errors.BizError) {
	var docs []*entity.DocEntity
	db := GetDB(dbNameMK).WithContext(d.ctx).
		Table(TableNameDoc).
		Where("space_key = ?", spaceKey).
		Find(&docs)
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return docs, nil
}

// UpdateDoc 更新文档信息
func (d *Doc) UpdateDoc(docEntity *entity.DocEntity) errors.BizError {
	docEntity.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(d.ctx).Table(TableNameDoc).Save(docEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// DeleteDoc 删除文档
func (d *Doc) DeleteDoc(docId int64) errors.BizError {
	db := GetDB(dbNameMK).WithContext(d.ctx).Table(TableNameDoc).Delete(&entity.DocEntity{}, docId)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}

// UpdateNameAndEditAccount 更新文档名称和编辑人
func (d *Doc) UpdateNameAndEditAccount(docId int64, name string, accountId int64, accountName string) errors.BizError {
	updateDoc := &entity.DocEntity{
		Name:            name,
		UpdateTime:      utils.NewJsonTime(time.Now()),
		EditAccountId:   accountId,
		EditAccountName: accountName,
	}
	db := GetDB(dbNameMK).WithContext(d.ctx).Table(TableNameDoc).
		Select("UpdateTime", "Name", "EditAccountId", "EditAccountName").
		Where(map[string]interface{}{
			DocPrimaryKey: docId,
			"status":      entity.DocEntityStatusNormal,
		}).
		Updates(updateDoc)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}
