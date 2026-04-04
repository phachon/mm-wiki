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
	// TableNameLink 快捷链接表
	TableNameLink = "mk_link"
	// LinkPrimaryKey 链接表主键ID
	LinkPrimaryKey = "link_id"
)

// Link 快捷链接表数据对象
type Link struct {
	ctx context.Context
}

// NewLink 创建快捷链接表数据对象
func NewLink(ctx context.Context) *Link {
	return &Link{
		ctx: ctx,
	}
}

// Insert 创建链接插入一条链接记录
func (r *Link) Insert(linkEntity *entity.LinkEntity) errors.BizError {
	linkEntity.CreateTime = utils.NewJsonTime(time.Now())
	linkEntity.UpdateTime = utils.NewJsonTime(time.Now())
	linkEntity.Status = entity.LinkStatusDefault
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameLink).Save(linkEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetLinkByLinkId 根据链接ID获取链接信息
func (r *Link) GetLinkByLinkId(linkId int64) (link *entity.LinkEntity, err errors.BizError) {

	link = &entity.LinkEntity{}
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameLink).
		Where(map[string]interface{}{
			LinkPrimaryKey: linkId,
			"status":       entity.LinkStatusDefault,
		}).
		First(&link)

	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return link, nil
}

// GetLinkByName 根据链接名称获取链接信息
func (r *Link) GetLinkByName(name string) (link *entity.LinkEntity, err errors.BizError) {

	link = &entity.LinkEntity{}
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameLink).
		Where(map[string]interface{}{
			"name":   name,
			"status": entity.LinkStatusDefault,
		}).
		First(&link)

	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return link, nil
}

// GetLinksByLimit 分页获取链接列表
func (r *Link) GetLinksByLimit(limit int, offset int) (links []*entity.LinkEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameLink).
		Where("status", entity.LinkStatusDefault).
		Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("sequence ASC, %s DESC", LinkPrimaryKey)).
		Find(&links)
	if db.Error != nil {
		return links, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return links, nil
}

// CountLinks 获取链接总数
func (r *Link) CountLinks() (count int64, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameLink).
		Where("status", entity.LinkStatusDefault).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// GetLinksByKeywordsAndLimit 根据关键字分页获取链接列表
func (r *Link) GetLinksByKeywordsAndLimit(limit int, offset int, keywords *entity.LinkKeywords) (links []*entity.LinkEntity, err errors.BizError) {

	if keywords == nil {
		return r.GetLinksByLimit(limit, offset)
	}

	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameLink)
	db.Where("status", entity.LinkStatusDefault)
	if keywords.Name != "" {
		db = db.Where("name LIKE ?", "%"+keywords.Name+"%")
	}
	db = db.Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("sequence ASC, %s DESC", LinkPrimaryKey)).
		Find(&links)
	if db.Error != nil {
		return links, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return links, nil
}

// CountLinksByKeywords 根据关键字获取链接总数
func (r *Link) CountLinksByKeywords(keywords *entity.LinkKeywords) (count int64, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameLink)
	db = db.Where("status = ?", entity.LinkStatusDefault)
	if keywords.Name != "" {
		db = db.Where("name LIKE ?", "%"+keywords.Name+"%")
	}
	db.Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// Update 更新链接只会更新如下字段
func (r *Link) Update(link entity.LinkEntity) errors.BizError {
	link.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameLink).
		Select("Name", "URL", "Sequence", "UpdateTime").
		Where(map[string]interface{}{
			LinkPrimaryKey: link.LinkId,
		}).
		Updates(link)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// DeleteLink 删除链接
func (r *Link) DeleteLink(linkId int64) errors.BizError {
	updateTime := utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameLink).
		Where(map[string]interface{}{
			LinkPrimaryKey: linkId,
			"status":       entity.LinkStatusDefault,
		}).
		Update("status", entity.LinkStatusDelete).
		Update("update_time", updateTime)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// GetAllLinks 获取所有的链接
func (r *Link) GetAllLinks() (links []*entity.LinkEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameLink).
		Where(map[string]interface{}{
			"status": entity.LinkStatusDefault,
		}).
		Order("sequence ASC").
		Find(&links)
	if db.Error != nil {
		return links, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return links, nil
}
