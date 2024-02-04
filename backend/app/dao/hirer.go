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
	// TableNameHirer 房屋租客表
	TableNameHirer = "hms_hirer"
	// HirerPrimaryKey 租客表主键ID
	HirerPrimaryKey = "hirer_id"
)

// Hirer 租客表数据
type Hirer struct {
	ctx context.Context
}

// NewHirer 创建房屋租客表数据对象
func NewHirer(ctx context.Context) *Hirer {
	return &Hirer{
		ctx: ctx,
	}
}

// Insert 创建租客插入一条租客记录
func (h *Hirer) Insert(hirerEntity *entity.HirerEntity) errors.BizError {
	hirerEntity.CreateTime = utils.NewJsonTime(time.Now())
	hirerEntity.UpdateTime = utils.NewJsonTime(time.Now())
	hirerEntity.Status = entity.HirerStatusDefault
	db := GetDB(dbNameKms).WithContext(h.ctx).
		Table(TableNameHirer).
		Save(hirerEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetHirerByName 根据租客姓名查找租客
func (h *Hirer) GetHirerByName(name string) (hirer *entity.HirerEntity, err errors.BizError) {
	hirer = &entity.HirerEntity{}
	db := GetDB(dbNameKms).WithContext(h.ctx).Table(TableNameHirer).
		Where(map[string]interface{}{
			"status": entity.HirerStatusDefault,
			"name":   name,
		}).
		First(&hirer)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return hirer, nil
}

// GetHirerByHirerId 根据租客ID获取租客信息
func (h *Hirer) GetHirerByHirerId(hirerId int64) (hirer *entity.HirerEntity, err errors.BizError) {

	hirer = &entity.HirerEntity{}
	db := GetDB(dbNameKms).WithContext(h.ctx).Table(TableNameHirer).
		Where(map[string]interface{}{
			HirerPrimaryKey: hirerId,
		}).
		First(&hirer)

	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return hirer, nil
}

// GetHirersByHirerIds 根据多个租客ID批量获取租客ID
func (h *Hirer) GetHirersByHirerIds(hirerIds []int64) (hirers []*entity.HirerEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(h.ctx).Table(TableNameHirer).
		Where("hirer_id IN (?)", hirerIds).
		Find(&hirers)
	if db.Error != nil {
		return hirers, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return hirers, nil
}

// CountByNickName 根据租客昵称获取租客数
func (h *Hirer) CountByName(name string) (count int64, err error) {
	if name == "" {
		return 0, nil
	}
	db := GetDB(dbNameKms).WithContext(h.ctx).Table(TableNameHirer).
		Where(map[string]interface{}{
			"status": entity.HirerStatusDefault,
			"name":   name,
		}).Count(&count)
	if db.Error != nil {
		return 0, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// HasSameName 租客ID和租客昵称是否存在
func (h *Hirer) HasSameName(hirerId int64, nickName string) (has bool, err errors.BizError) {
	var count int64
	db := GetDB(dbNameKms).WithContext(h.ctx).Table(TableNameHirer).
		Where("name = ?", nickName).
		Where("hirer_id <> ?", hirerId).
		Count(&count)
	if db.Error != nil {
		return false, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count > 0, nil
}

// Update 更新租客，只会更新如下字段
func (h *Hirer) Update(hirer entity.HirerEntity) errors.BizError {
	hirer.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(h.ctx).Table(TableNameHirer).
		Select(
			"Name", "Sex", "Email", "IdCardNumber", "Mobile",
			"EmergencyContact", "EmergencyMobile", "UpdateTime",
		).
		Where(map[string]interface{}{
			HirerPrimaryKey: hirer.HirerId,
		}).
		Updates(hirer)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// UpdateStatus 更新租客状态
func (h *Hirer) UpdateStatus(hirerId int64, status int) errors.BizError {
	updateTime := utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(h.ctx).Table(TableNameHirer).
		Where(map[string]interface{}{
			HirerPrimaryKey: hirerId,
		}).
		Update("status", status).
		Update("update_time", updateTime)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// GetAllHirers 获取所有的租客
func (h *Hirer) GetAllHirers() (hirer []*entity.HirerEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(h.ctx).
		Table(TableNameHirer).
		Where("status = ?", entity.HirerStatusDefault).
		Find(&hirer)
	if db.Error != nil {
		return hirer, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return hirer, nil
}

// GetHirersByKeywords 根据租客名模糊匹配租客
func (h *Hirer) GetHirersByKeywords(keywords *entity.HirerKeywords) (hirers []*entity.HirerEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(h.ctx).
		Table(TableNameHirer).
		Where("status = ?", entity.HirerStatusDefault)
	if keywords.Name != "" {
		db = db.Where("name LIKE ?", "%"+keywords.Name+"%")
	}
	db.Find(&hirers)
	if db.Error != nil {
		return hirers, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return hirers, nil
}

// GetHirersByKeywordAndLimit 根据关键字分页获取租客
func (h *Hirer) GetHirersByKeywordAndLimit(limit int, offset int, keywords *entity.HirerKeywords) (hirers []*entity.HirerEntity, err errors.BizError) {
	if keywords == nil {
		return h.GetHirersByLimit(limit, offset)
	}
	db := GetDB(dbNameKms).WithContext(h.ctx).
		Table(TableNameHirer).
		Where("status = ?", entity.HirerStatusDefault)
	if keywords.Name != "" {
		db = db.Where("name LIKE ?", "%"+keywords.Name+"%")
	}
	db = db.Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", HirerPrimaryKey)).
		Find(&hirers)
	if db.Error != nil {
		return hirers, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return hirers, nil
}

// GetHirersByLimit 分页获取租客列表
func (h *Hirer) GetHirersByLimit(limit int, offset int) (hirers []*entity.HirerEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(h.ctx).
		Table(TableNameHirer).
		Where("status = ?", entity.HirerStatusDefault).
		Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", HirerPrimaryKey)).
		Find(&hirers)
	if db.Error != nil {
		return hirers, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return hirers, nil
}

// CountHirers 租客总数
func (h *Hirer) CountHirers() (count int64, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(h.ctx).
		Table(TableNameHirer).
		Where("status = ?", entity.HirerStatusDefault).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// CountHirers 租客总数
func (h *Hirer) CountHirersByKeywords(keywords *entity.HirerKeywords) (count int64, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(h.ctx).
		Table(TableNameHirer).
		Where("status = ?", entity.HirerStatusDefault)
	if keywords.Name != "" {
		db = db.Where("name LIKE ?", "%"+keywords.Name+"%")
	}
	db.Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}
