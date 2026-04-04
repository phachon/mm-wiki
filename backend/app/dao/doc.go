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

// UpdateDocSequence 更新文档排序
func (d *Doc) UpdateDocSequence(docId int64, sequence int) errors.BizError {
	db := GetDB(dbNameMK).WithContext(d.ctx).Table(TableNameDoc).
		Where(map[string]interface{}{
			DocPrimaryKey: docId,
		}).
		Updates(map[string]interface{}{
			"sequence":    sequence,
			"update_time": utils.NewJsonTime(time.Now()),
		})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// SearchDocs 搜索文档（按文档名模糊匹配）
func (d *Doc) SearchDocs(keyword string, spaceKey string, limit int, offset int) ([]*entity.DocEntity, errors.BizError) {
	var docs []*entity.DocEntity
	query := GetDB(dbNameMK).WithContext(d.ctx).
		Table(TableNameDoc).
		Where("name LIKE ? AND status = ?", "%"+keyword+"%", entity.DocEntityStatusNormal)
	if spaceKey != "" {
		query = query.Where("space_key = ?", spaceKey)
	}
	db := query.Order("update_time DESC").Limit(limit).Offset(offset).Find(&docs)
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return docs, nil
}

// CountSearchDocs 统计搜索文档数量
func (d *Doc) CountSearchDocs(keyword string, spaceKey string) (int64, errors.BizError) {
	var count int64
	query := GetDB(dbNameMK).WithContext(d.ctx).
		Table(TableNameDoc).
		Where("name LIKE ? AND status = ?", "%"+keyword+"%", entity.DocEntityStatusNormal)
	if spaceKey != "" {
		query = query.Where("space_key = ?", spaceKey)
	}
	db := query.Count(&count)
	if db.Error != nil {
		return 0, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
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
