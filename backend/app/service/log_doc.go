package service

import (
	"context"

	"github.com/phachon/mm-wiki/app/dao"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
	"github.com/phachon/mm-wiki/utils"
)

const (
	LogDocListDefaultPageSize = 20 // 文档日志列表默认一页 20 条
)

// LogDoc 文档操作日志业务逻辑
type LogDoc struct {
	ctx       context.Context
	daoLogDoc *dao.LogDoc
}

// NewLogDoc 创建文档操作日志业务逻辑对象
func NewLogDoc(ctx context.Context) *LogDoc {
	return &LogDoc{
		ctx:       ctx,
		daoLogDoc: dao.NewLogDoc(ctx),
	}
}

// Create 创建文档日志
func (ld *LogDoc) Create(logDocEntity *entity.LogDocEntity) errors.BizError {
	err := ld.daoLogDoc.Insert(logDocEntity)
	if err != nil {
		return err
	}
	return nil
}

// GetLogDocsByDocId 根据文档ID获取日志列表
func (ld *LogDoc) GetLogDocsByDocId(docId string) (logDocs []*entity.LogDocEntity, err errors.BizError) {
	return ld.daoLogDoc.GetLogDocsByDocId(docId)
}

// GetLogDocsBySpaceId 根据空间ID获取日志列表
func (ld *LogDoc) GetLogDocsBySpaceId(spaceId int64) (logDocs []*entity.LogDocEntity, err errors.BizError) {
	return ld.daoLogDoc.GetLogDocsBySpaceId(spaceId)
}

// GetLogDocsByLimit 分页获取文档日志列表
func (ld *LogDoc) GetLogDocsByLimit(pageSize int, pageNum int, keywords *entity.LogDocKeywords) (logDocs []*entity.LogDocEntity, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, LogDocListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	offset := (pageNum - 1) * pageSize
	if keywords == nil {
		return ld.daoLogDoc.GetLogDocsByLimit(pageSize, offset)
	}
	return ld.daoLogDoc.GetLogDocsByKeywordsAndLimit(pageSize, offset, keywords)
}

// GetPageInfoLimit 获取分页信息
func (ld *LogDoc) GetPageInfoLimit(pageSize int, pageNum int, keywords *entity.LogDocKeywords) (pagination *entity.PageInfo, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, LogDocListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	pageInfo := new(entity.PageInfo)
	var totalNum int64
	if keywords == nil {
		totalNum, err = ld.daoLogDoc.CountLogDocs()
	} else {
		totalNum, err = ld.daoLogDoc.CountLogDocsByKeywords(keywords)
	}
	if err != nil {
		logger.WithContext(ld.ctx).Warnf("[service.LogDoc] GetPageInfoLimit pageSize=%d pageNum=%d err=%s",
			pageSize, pageNum, err.Error())
		return pageInfo, err
	}
	return entity.GetPageInfo(totalNum, pageSize, pageNum), nil
}
