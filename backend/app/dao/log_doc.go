package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
)

const (
	// TableNameLogDoc 文档操作日志表
	TableNameLogDoc = "mk_log_doc"
	// LogDocPrimaryKey 文档日志表主键ID
	LogDocPrimaryKey = "log_doc_id"
)

// LogDoc 文档操作日志表数据对象
type LogDoc struct {
	ctx context.Context
}

// NewLogDoc 创建文档操作日志表数据对象
func NewLogDoc(ctx context.Context) *LogDoc {
	return &LogDoc{
		ctx: ctx,
	}
}

// Insert 插入一条文档日志记录
func (l *LogDoc) Insert(logDocEntity *entity.LogDocEntity) errors.BizError {
	logDocEntity.CreateTime = time.Now().Unix()
	db := GetDB(dbNameMK).WithContext(l.ctx).Table(TableNameLogDoc).Save(logDocEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetLogDocsByDocId 根据文档ID获取日志列表
func (l *LogDoc) GetLogDocsByDocId(docId string) (logDocs []*entity.LogDocEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(l.ctx).Table(TableNameLogDoc).
		Where("doc_id = ?", docId).
		Order(fmt.Sprintf("%s DESC", LogDocPrimaryKey)).
		Find(&logDocs)
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return logDocs, nil
}

// GetLogDocsBySpaceId 根据空间ID获取日志列表
func (l *LogDoc) GetLogDocsBySpaceId(spaceId int64) (logDocs []*entity.LogDocEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(l.ctx).Table(TableNameLogDoc).
		Where("space_id = ?", spaceId).
		Order(fmt.Sprintf("%s DESC", LogDocPrimaryKey)).
		Find(&logDocs)
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return logDocs, nil
}

// GetLogDocsByLimit 分页获取文档日志列表
func (l *LogDoc) GetLogDocsByLimit(limit int, offset int) (logDocs []*entity.LogDocEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(l.ctx).Table(TableNameLogDoc).
		Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", LogDocPrimaryKey)).
		Find(&logDocs)
	if db.Error != nil {
		return logDocs, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return logDocs, nil
}

// CountLogDocs 获取文档日志总数
func (l *LogDoc) CountLogDocs() (count int64, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(l.ctx).Table(TableNameLogDoc).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// GetLogDocsByKeywordsAndLimit 根据关键字分页获取文档日志列表
func (l *LogDoc) GetLogDocsByKeywordsAndLimit(limit int, offset int, keywords *entity.LogDocKeywords) (logDocs []*entity.LogDocEntity, err errors.BizError) {
	if keywords == nil {
		return l.GetLogDocsByLimit(limit, offset)
	}
	db := GetDB(dbNameMK).WithContext(l.ctx).Table(TableNameLogDoc)
	if keywords.DocId != "" {
		db = db.Where("doc_id = ?", keywords.DocId)
	}
	if keywords.SpaceId > 0 {
		db = db.Where("space_id = ?", keywords.SpaceId)
	}
	db = db.Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", LogDocPrimaryKey)).
		Find(&logDocs)
	if db.Error != nil {
		return logDocs, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return logDocs, nil
}

// CountLogDocsByKeywords 根据关键字获取文档日志总数
func (l *LogDoc) CountLogDocsByKeywords(keywords *entity.LogDocKeywords) (count int64, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(l.ctx).Table(TableNameLogDoc)
	if keywords.DocId != "" {
		db = db.Where("doc_id = ?", keywords.DocId)
	}
	if keywords.SpaceId > 0 {
		db = db.Where("space_id = ?", keywords.SpaceId)
	}
	db.Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}
