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
	// TableNameHouse 房屋房产表
	TableNameHouse = "hms_house"
	// HousePrimaryKey 房产表主键ID
	HousePrimaryKey = "house_id"
)

// House 房屋房产表数据
type House struct {
	ctx context.Context
}

// NewHouse 创建房屋房产表数据对象
func NewHouse(ctx context.Context) *House {
	return &House{
		ctx: ctx,
	}
}

// Insert 创建房产插入一条房产记录
func (l *House) Insert(houseEntity *entity.HouseEntity) errors.BizError {
	houseEntity.CreateTime = utils.NewJsonTime(time.Now())
	houseEntity.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(l.ctx).
		Table(TableNameHouse).
		Save(houseEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetHouseByHouseId 根据房产ID获取房产信息
func (l *House) GetHouseByHouseId(houseId int64) (house *entity.HouseEntity, err errors.BizError) {

	house = &entity.HouseEntity{}
	db := GetDB(dbNameKms).WithContext(l.ctx).Table(TableNameHouse).
		Where(map[string]interface{}{
			HousePrimaryKey: houseId,
			"status":        entity.HouseStatusDefault,
		}).
		First(&house)

	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return house, nil
}

// GetHousesByHouseIds 根据多个房产ID批量获取房产ID
func (l *House) GetHousesByHouseIds(houseIds []int64) (houses []*entity.HouseEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(l.ctx).Table(TableNameHouse).
		Where("status = ?", entity.HouseStatusDefault).
		Where("house_id IN (?)", houseIds).
		Find(&houses)
	if db.Error != nil {
		return houses, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return houses, nil
}

// Update 更新房产，只会更新如下字段
func (l *House) Update(house entity.HouseEntity) errors.BizError {
	house.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(l.ctx).Table(TableNameHouse).
		Select(
			"LandlordId", "Region", "Address", "HouseNumber",
			"DecorationType", "SizeType", "Area",
			"AllowLease", "BuildTime", "UpdateTime",
			"AllowSplitLease", "MouthRent",
		).
		Where(map[string]interface{}{
			HousePrimaryKey: house.HouseId,
		}).
		Updates(house)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// UpdateStatus 更新房产状态
func (l *House) UpdateStatus(houseId int64, status int) errors.BizError {
	updateTime := utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(l.ctx).Table(TableNameHouse).
		Where(map[string]interface{}{
			HousePrimaryKey: houseId,
		}).
		Update("status", status).
		Update("update_time", updateTime)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// UpdateLeaseStatus 更新房产出租状态
func (l *House) UpdateLeaseStatus(houseId int64, leaseStatus int) errors.BizError {
	updateTime := utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(l.ctx).Table(TableNameHouse).
		Where(map[string]interface{}{
			HousePrimaryKey: houseId,
		}).
		Update("lease_status", leaseStatus).
		Update("update_time", updateTime)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// GetAllHouses 获取所有的房产
func (l *House) GetAllHouses() (house []*entity.HouseEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(l.ctx).
		Table(TableNameHouse).
		Where("status = ?", entity.HouseStatusDefault).
		Find(&house)
	if db.Error != nil {
		return house, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return house, nil
}

// GetAllowSplitLeaseHouses 获取允许合租的房产
func (l *House) GetAllowSplitLeaseHouses() (houses []*entity.HouseEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(l.ctx).
		Table(TableNameHouse).
		Where("status = ?", entity.HouseStatusDefault).
		Where("allow_split_lease = ?", entity.HouseAllowSplitLeaseYes).
		Find(&houses)
	if db.Error != nil {
		return houses, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return houses, nil
}

// GetAllowLeaseHouses 获取允许整租的房产
func (l *House) GetAllowLeaseHouses() (houses []*entity.HouseEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(l.ctx).
		Table(TableNameHouse).
		Where("status = ?", entity.HouseStatusDefault).
		Where("allow_lease = ?", entity.HouseAllowLeaseYes).
		Find(&houses)
	if db.Error != nil {
		return houses, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return houses, nil
}

// GetLeaseAndNotRentHouses 获取允许整租但未出租的房产
func (l *House) GetLeaseAndNotRentHouses() (houses []*entity.HouseEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(l.ctx).
		Table(TableNameHouse).
		Where("status = ?", entity.HouseStatusDefault).
		Where("allow_lease = ?", entity.HouseAllowLeaseYes).
		Where("lease_status = ?", entity.HouseLeaseStatusNotRent).
		Find(&houses)
	if db.Error != nil {
		return houses, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return houses, nil
}

// GetSplitLeaseAndNotRentHouses 获取允许合租但是未出租的房产
func (l *House) GetSplitLeaseAndNotRentHouses() (houses []*entity.HouseEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(l.ctx).
		Table(TableNameHouse).
		Where("status = ?", entity.HouseStatusDefault).
		Where("allow_split_lease = ?", entity.HouseAllowSplitLeaseYes).
		Where("lease_status = ?", entity.HouseLeaseStatusNotRent).
		Find(&houses)
	if db.Error != nil {
		return houses, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return houses, nil
}

// GetHousesByKeywords 根据房产名模糊匹配房产
func (l *House) GetHousesByKeywords(keywords *entity.HouseKeywords) (houses []*entity.HouseEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(l.ctx).
		Table(TableNameHouse).
		Where("status = ?", entity.HouseStatusDefault)
	db = l.getHouseWhereByKeywords(keywords, db)
	db.Find(&houses)
	if db.Error != nil {
		return houses, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return houses, nil
}

func (*House) getHouseWhereByKeywords(keywords *entity.HouseKeywords, db *gorm.DB) *gorm.DB {
	if keywords.Region != "" {
		db = db.Where("region LIKE ?", "%"+keywords.Region+"%")
	}
	if keywords.Address != "" {
		db = db.Where("address LIKE ?", "%"+keywords.Address+"%")
	}
	if keywords.DecorationType != nil && *keywords.DecorationType >= 0 {
		db = db.Where("decoration_type = ?", *keywords.DecorationType)
	}
	if keywords.SizeType != nil && *keywords.SizeType >= 0 {
		db = db.Where("size_type = ?", *keywords.SizeType)
	}
	if keywords.LeaseStatus != nil && *keywords.LeaseStatus >= 0 {
		db = db.Where("lease_status = ?", *keywords.LeaseStatus)
	}
	return db
}

// GetHousesByKeywordAndLimit 根据关键字分页获取房产
func (l *House) GetHousesByKeywordAndLimit(limit int, offset int, keywords *entity.HouseKeywords) (houses []*entity.HouseEntity, err errors.BizError) {
	if keywords == nil {
		return l.GetHousesByLimit(limit, offset)
	}
	db := GetDB(dbNameKms).WithContext(l.ctx).
		Table(TableNameHouse).
		Where("status = ?", entity.HouseStatusDefault)
	db = l.getHouseWhereByKeywords(keywords, db)
	db = db.Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", HousePrimaryKey)).
		Find(&houses)
	if db.Error != nil {
		return houses, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return houses, nil
}

// GetHousesByLimit 分页获取房产列表
func (l *House) GetHousesByLimit(limit int, offset int) (houses []*entity.HouseEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(l.ctx).
		Table(TableNameHouse).
		Where("status = ?", entity.HouseStatusDefault).
		Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", HousePrimaryKey)).
		Find(&houses)
	if db.Error != nil {
		return houses, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return houses, nil
}

// CountHouses 房产总数
func (l *House) CountHouses() (count int64, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(l.ctx).
		Table(TableNameHouse).
		Where("status = ?", entity.HouseStatusDefault).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// CountHousesByKeywords 根据关键词搜索房产总数
func (l *House) CountHousesByKeywords(keywords *entity.HouseKeywords) (count int64, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(l.ctx).
		Table(TableNameHouse).
		Where("status = ?", entity.HouseStatusDefault)
	db = l.getHouseWhereByKeywords(keywords, db)
	db.Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}
