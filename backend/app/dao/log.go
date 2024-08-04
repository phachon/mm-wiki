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
	// TableNameLog 系统日志表
	TableNameLog = "mk_log"
	// LogPrimaryKey 日志表主键ID
	LogPrimaryKey = "log_id"
)

// Log 系统日志表数据
type Log struct {
	ctx context.Context
}

// NewLog 创建系统日志表数据对象
func NewLog(ctx context.Context) *Log {
	return &Log{
		ctx: ctx,
	}
}

// Insert 创建日志插入一条日志记录
func (l *Log) Insert(logEntity *entity.LogEntity) errors.BizError {
	logEntity.CreateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(l.ctx).Table(TableNameLog).Save(logEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetLogByMessage 根据日志信息查找日志
func (l *Log) GetLogByMessage(message string) (log *entity.LogEntity, err errors.BizError) {
	log = &entity.LogEntity{}
	db := GetDB(dbNameMK).WithContext(l.ctx).Table(TableNameLog).
		Where(map[string]interface{}{
			"message LIKE ?": "%" + message + "%",
		}).
		First(&log)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return log, nil
}

// GetLogByLogId 根据日志ID获取日志信息
func (l *Log) GetLogByLogId(logId int64) (log *entity.LogEntity, err errors.BizError) {

	log = &entity.LogEntity{}
	db := GetDB(dbNameMK).WithContext(l.ctx).Table(TableNameLog).
		Where(map[string]interface{}{
			LogPrimaryKey: logId,
		}).
		First(&log)

	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return log, nil
}

// GetLogsByLogIds 根据多个日志ID批量获取日志ID
func (l *Log) GetLogsByLogIds(logIds []int64) (logs []*entity.LogEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(l.ctx).Table(TableNameLog).
		Where("log_id IN (?)", logIds).
		Find(&logs)
	if db.Error != nil {
		return logs, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return logs, nil
}

// CountByMessage 根据日志信息获取日志数量
func (l *Log) CountByMessage(message string) (count int64, err error) {
	if message == "" {
		return 0, nil
	}
	db := GetDB(dbNameMK).WithContext(l.ctx).Table(TableNameLog).
		Where(map[string]interface{}{
			"message LIKE ?": "%" + message + "%",
		}).Count(&count)
	if db.Error != nil {
		return 0, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// GetAllLogs 获取所有的日志
func (l *Log) GetAllLogs() (log []*entity.LogEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(l.ctx).Table(TableNameLog).
		Find(&log)
	if db.Error != nil {
		return log, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return log, nil
}

// GetLogsByLimit 分页获取日志列表
func (l *Log) GetLogsByLimit(limit int, offset int) (logs []*entity.LogEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(l.ctx).Table(TableNameLog).
		Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", LogPrimaryKey)).
		Find(&logs)
	if db.Error != nil {
		return logs, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return logs, nil
}

// CountLogs 获取日志总数
func (l *Log) CountLogs() (count int64, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(l.ctx).Table(TableNameLog).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// GetLogsByMessageAndLimit 根据日志信息分页获取日志列表
func (l *Log) GetLogsByMessageAndLimit(limit int, offset int, message string) (logs []*entity.LogEntity, err errors.BizError) {

	if message == "" {
		return l.GetLogsByLimit(limit, offset)
	}

	db := GetDB(dbNameMK).WithContext(l.ctx).
		Table(TableNameLog).
		Where("message LIKE ?", "%"+message+"%").
		Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", LogPrimaryKey)).
		Find(&logs)
	if db.Error != nil {
		return logs, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return logs, nil
}

// CountLogsByMessage 根据日志关键字获取日志总数
func (l *Log) CountLogsByMessage(message string) (count int64, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(l.ctx).
		Table(TableNameLog).
		Where("name LIKE ?", "%"+message+"%").
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// GetLogsByAccountId 根据账号ID获取日志列表
func (l *Log) GetLogsByAccountId(accountId int64) (logs []*entity.LogEntity, err errors.BizError) {

	db := GetDB(dbNameMK).WithContext(l.ctx).
		Table(TableNameLog).
		Where("account_id = ?", accountId).
		Order(fmt.Sprintf("%s DESC", LogPrimaryKey)).
		Find(&logs)
	if db.Error != nil {
		return logs, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return logs, nil
}

// GetLogsByLevel 根据日志类型获取日志列表
func (l *Log) GetLogsByLevel(level int) (logs []*entity.LogEntity, err errors.BizError) {

	db := GetDB(dbNameMK).WithContext(l.ctx).
		Table(TableNameLog).
		Where("level = ?", level).
		Order(fmt.Sprintf("%s DESC", LogPrimaryKey)).
		Find(&logs)
	if db.Error != nil {
		return logs, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return logs, nil
}

// GetLogsByAccountIdAndLimit 根据账号ID分页获取日志列表
func (l *Log) GetLogsByAccountIdAndLimit(accountId int64, limit int, offset int) (logs []*entity.LogEntity, err errors.BizError) {

	db := GetDB(dbNameMK).WithContext(l.ctx).
		Table(TableNameLog).
		Where("account_id = ?", accountId).
		Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", LogPrimaryKey)).
		Find(&logs)
	if db.Error != nil {
		return logs, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return logs, nil
}

// GetLogsByLevelAndLimit 根据日志类型分页获取日志列表
func (l *Log) GetLogsByLevelAndLimit(level int, limit int, offset int) (logs []*entity.LogEntity, err errors.BizError) {

	db := GetDB(dbNameMK).WithContext(l.ctx).
		Table(TableNameLog).
		Where("level = ?", level).
		Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", LogPrimaryKey)).
		Find(&logs)
	if db.Error != nil {
		return logs, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return logs, nil
}

// CountLogsByLevel 根据日志类型获取日志总数
func (l *Log) CountLogsByLevel(level int) (count int64, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(l.ctx).
		Table(TableNameLog).
		Where("level = ?", level).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// CountLogsByAccountId 根据账号ID获取日志总数
func (l *Log) CountLogsByAccountId(accountId int64) (count int64, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(l.ctx).
		Table(TableNameLog).
		Where("account_id = ?", accountId).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// GetLogsByKeywordsAndLimit 根据搜索关键词分页获取日志列表
func (l *Log) GetLogsByKeywordsAndLimit(keywords *entity.LogSearchKeywords, limit int, offset int) (logs []*entity.LogEntity, err errors.BizError) {

	db := GetDB(dbNameMK).WithContext(l.ctx).
		Table(TableNameLog)
	if keywords != nil && keywords.Level > 0 {
		db = db.Where("level = ?", keywords.Level)
	}
	if keywords != nil && keywords.AccountId != 0 {
		db = db.Where("account_id = ?", keywords.AccountId)
	}
	if keywords != nil && keywords.Message != "" {
		db = db.Where("message LIKE ?", "%"+keywords.Message+"%")
	}
	db.Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", LogPrimaryKey)).
		Find(&logs)
	if db.Error != nil {
		return logs, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return logs, nil
}

// CountLogsByKeywords 根据日志搜索词获取日志总数
func (l *Log) CountLogsByKeywords(keywords *entity.LogSearchKeywords) (count int64, err errors.BizError) {
	if keywords == nil {
		return
	}
	db := GetDB(dbNameMK).WithContext(l.ctx).
		Table(TableNameLog)
	if keywords != nil && keywords.Level > 0 {
		db = db.Where("level = ?", keywords.Level)
	}
	if keywords != nil && keywords.AccountId != 0 {
		db = db.Where("account_id = ?", keywords.AccountId)
	}
	if keywords != nil && keywords.Message != "" {
		db = db.Where("message LIKE ?", "%"+keywords.Message+"%")
	}
	if db.Count(&count).Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}
