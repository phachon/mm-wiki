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
	// TableNameNotice 系统公告表
	TableNameNotice = "hms_notice"
	// NoticePrimaryKey 公告表主键ID
	NoticePrimaryKey = "notice_id"
)

// Notice 系统公告表数据对象
type Notice struct {
	ctx context.Context
}

// NewNotice 创建系统公告表数据对象
func NewNotice(ctx context.Context) *Notice {
	return &Notice{
		ctx: ctx,
	}
}

// Insert 创建公告插入一条公告记录
func (r *Notice) Insert(noticeEntity *entity.NoticeEntity) errors.BizError {
	noticeEntity.CreateTime = utils.NewJsonTime(time.Now())
	noticeEntity.UpdateTime = utils.NewJsonTime(time.Now())
	noticeEntity.Status = entity.NoticeStatusDefault
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameNotice).Save(noticeEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetNoticeByContent 根据公告内容查找正常的公告
func (r *Notice) GetNoticeByContent(content string) (notices []*entity.NoticeEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameNotice).
		Where(map[string]interface{}{
			"status":        entity.NoticeStatusDefault,
			"content LIKE ": "%" + content + "%",
		}).
		Find(&notices)
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return notices, nil
}

// GetNoticeByNoticeId 根据公告ID获取公告信息
func (r *Notice) GetNoticeByNoticeId(noticeId int64) (notice *entity.NoticeEntity, err errors.BizError) {

	notice = &entity.NoticeEntity{}
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameNotice).
		Where(map[string]interface{}{
			NoticePrimaryKey: noticeId,
			"status":         entity.NoticeStatusDefault,
		}).
		First(&notice)

	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return notice, nil
}

// GetNoticesByNoticeIds 根据多个公告ID批量获取公告ID
func (r *Notice) GetNoticesByNoticeIds(noticeIds []int64) (notices []*entity.NoticeEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameNotice).
		Where("status=? and notice_id IN (?)",
			entity.NoticeStatusDefault, noticeIds).
		Find(&notices)
	if db.Error != nil {
		return notices, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return notices, nil
}

// CountByName 根据公告名获取公告数量
func (r *Notice) CountByContent(content string) (count int64, err error) {
	if content == "" {
		return 0, nil
	}
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameNotice).
		Where(map[string]interface{}{
			"status":        entity.NoticeStatusDefault,
			"content LIKE ": "%" + content + "%",
		}).Count(&count)
	if db.Error != nil {
		return 0, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// Update 更新公告只会更新如下字段
func (r *Notice) Update(notice entity.NoticeEntity) errors.BizError {
	notice.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameNotice).
		Select("Title", "Content", "AccountId", "Status", "PublishStatus", "PublishTime", "StartTime", "EndTime").
		Where(map[string]interface{}{
			NoticePrimaryKey: notice.NoticeId,
		}).
		Updates(notice)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// UpdateStatus 删除公告
func (r *Notice) DeleteNotice(noticeId int64) errors.BizError {
	updateTime := utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameNotice).
		Where(map[string]interface{}{
			NoticePrimaryKey: noticeId,
			"status":         entity.NoticeStatusDefault,
		}).
		Update("status", entity.NoticeStatusDelete).
		Update("update_time", updateTime)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// GetAllNotices 获取所有的公告
func (r *Notice) GetAllNotices() (notice []*entity.NoticeEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameNotice).
		Where(map[string]interface{}{
			"status": entity.NoticeStatusDefault,
		}).
		Find(&notice)
	if db.Error != nil {
		return notice, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return notice, nil
}

// GetNoticesByLimit 分页获取公告列表
func (r *Notice) GetNoticesByLimit(limit int, offset int) (notices []*entity.NoticeEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameNotice).
		Where("status", entity.NoticeStatusDefault).
		Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", NoticePrimaryKey)).
		Find(&notices)
	if db.Error != nil {
		return notices, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return notices, nil
}

// GetPublishNoticesByLimit 分页获取已发布公告列表
func (r *Notice) GetPublishNoticesByLimit(limit int, offset int) (notices []*entity.NoticeEntity, err errors.BizError) {
	now := time.Now().Format("2006-01-02 15:04:05")
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameNotice).
		Where("status", entity.NoticeStatusDefault).
		Where("publish_status", entity.NoticePublishStatusPublishing).
		Where("start_time <= ? and end_time > ?", now, now).
		Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", NoticePrimaryKey)).
		Find(&notices)
	if db.Error != nil {
		return notices, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return notices, nil
}

// CountNotices 获取公告总数
func (r *Notice) CountNotices() (count int64, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameNotice).
		Where("status", entity.NoticeStatusDefault).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// CountPublishNotices 获取已发布公告总数
func (r *Notice) CountPublishNotices() (count int64, err errors.BizError) {
	now := time.Now().Format("2006-01-02 15:04:05")
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameNotice).
		Where("status", entity.NoticeStatusDefault).
		Where("status", entity.NoticeStatusDefault).
		Where("publish_status", entity.NoticePublishStatusPublishing).
		Where("start_time <= ? and end_time > ?", now, now).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// GetNoticesByKeywordsAndLimit 根据关键字分页获取公告列表
func (r *Notice) GetNoticesByKeywordsAndLimit(limit int, offset int, keywords *entity.NoticeKeywords) (notices []*entity.NoticeEntity, err errors.BizError) {

	if keywords == nil {
		return r.GetNoticesByLimit(limit, offset)
	}

	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameNotice)
	db.Where("status", entity.NoticeStatusDefault)
	if keywords.Content != "" {
		db = db.Where("content LIKE ?", "%"+keywords.Content+"%")
	}
	db = db.Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", NoticePrimaryKey)).
		Find(&notices)
	if db.Error != nil {
		return notices, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return notices, nil
}

// CountNotices 根据关键字获取公告总数
func (r *Notice) CountNoticesByKeywords(keywords *entity.NoticeKeywords) (count int64, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(r.ctx).Table(TableNameNotice)
	db = db.Where("status = ?", entity.NoticeStatusDefault)
	if keywords.Content != "" {
		db = db.Where("content LIKE ?", "%"+keywords.Content+"%")
	}
	db.Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}
