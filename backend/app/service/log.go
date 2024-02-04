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
	LogListDefaultPageSize = 20 // 日志列表默认一页 20 条
)

// Log 日志业务逻辑
type Log struct {
	ctx    context.Context
	daoLog *dao.Log
}

// NewLog 创建日志业务逻辑对象
func NewLog(ctx context.Context) *Log {
	return &Log{
		ctx:    ctx,
		daoLog: dao.NewLog(ctx),
	}
}

// Create 创建日志
func (l *Log) Create(logEntity *entity.LogEntity) errors.BizError {
	return l.daoLog.Insert(logEntity)
}

// GetLogByLogId 根据日志ID获取日志详情
func (l *Log) GetLogByLogId(logID int64) (log *entity.LogEntity, err errors.BizError) {
	log, err = l.daoLog.GetLogByLogId(logID)
	if err != nil {
		return log, err
	}
	return log, nil
}

// GetLogsByKeywordsAndLimit 根据关键字分页获取日志列表
func (l *Log) GetLogsByKeywordsAndLimit(pageSize int, pageNum int, keywords *entity.LogSearchKeywords) (logs []*entity.LogEntity, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, LogListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	offset := (pageNum - 1) * pageSize
	// 无搜索词
	if keywords == nil {
		return l.daoLog.GetLogsByLimit(pageSize, offset)
	}
	return l.daoLog.GetLogsByKeywordsAndLimit(keywords, pageSize, offset)
}

// GetPageInfoLimit 获取分页信息
func (l *Log) GetPageInfoLimit(pageSize int, pageNum int, keywords *entity.LogSearchKeywords) (pagination *entity.PageInfo, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, LogListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	pageInfo := new(entity.PageInfo)
	var totalNum int64
	// 无搜索词
	if keywords == nil {
		totalNum, err = l.daoLog.CountLogs()
	} else {
		totalNum, err = l.daoLog.CountLogsByKeywords(keywords)
	}
	if err != nil {
		logger.WithContext(l.ctx).Warnf("[service.Log] GetPageInfoLimit pageSize=%d pageNum=%d err=%s",
			pageSize, pageNum, err.Error())
		return pageInfo, err
	}
	return entity.GetPageInfo(totalNum, pageSize, pageNum), nil
}

// GetAllLogs 获取所有的日志
func (l *Log) GetAllLogs() (logs []*entity.LogEntity, err errors.BizError) {
	return l.daoLog.GetAllLogs()
}

// GetLogsByLogIds 根据日志ID批量获取日志
func (l *Log) GetLogsByLogIds(logIds []int64) (logs []*entity.LogEntity, err errors.BizError) {
	return l.daoLog.GetLogsByLogIds(logIds)
}
