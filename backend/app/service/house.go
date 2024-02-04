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
	HouseListDefaultPageSize = 20 // 房产列表默认一页 20 条
)

// House 房产业务逻辑
type House struct {
	ctx      context.Context
	daoHouse *dao.House
}

// NewHouse 创建房产业务逻辑对象
func NewHouse(ctx context.Context) *House {
	return &House{
		ctx:      ctx,
		daoHouse: dao.NewHouse(ctx),
	}
}

// Create 创建房产
func (a *House) Create(houseEntity *entity.HouseEntity) errors.BizError {

	// @todo 生成房产唯一编号

	// 插入一条记录
	err := a.daoHouse.Insert(houseEntity)
	if err != nil {
		return err
	}
	return nil
}

// Update 修改房产
func (a *House) Update(houseEntity entity.HouseEntity) errors.BizError {
	// 查找房产是否存在
	updateHouse, err := a.daoHouse.GetHouseByHouseId(houseEntity.HouseId)
	if err != nil {
		return err
	}
	if updateHouse == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "房产不存在")
	}
	// 更新字段
	err = a.daoHouse.Update(houseEntity)
	if err != nil {
		return err
	}
	return nil
}

// GetHouseByHouseId 根据房产ID获取房产详情
func (a *House) GetHouseByHouseId(houseID int64) (house *entity.HouseEntity, err errors.BizError) {
	house, err = a.daoHouse.GetHouseByHouseId(houseID)
	if err != nil {
		return house, err
	}
	return house, nil
}

// GetHousesByHouseIds 根据多个房产ID获取房产信息
func (a *House) GetHousesByHouseIds(houseIDs []int64) (houses []*entity.HouseEntity, err errors.BizError) {
	if len(houseIDs) == 0 {
		return houses, nil
	}
	houses, err = a.daoHouse.GetHousesByHouseIds(houseIDs)
	if err != nil {
		return houses, err
	}
	return houses, nil
}

// GetHousesMapByHouseIds 根据多个房产ID获取房产信息，返回map
func (a *House) GetHousesMapByHouseIds(houseIDs []int64) (housesMap map[int64]*entity.HouseEntity, err errors.BizError) {
	housesMap = make(map[int64]*entity.HouseEntity)
	if len(houseIDs) == 0 {
		return housesMap, nil
	}
	houses, err := a.GetHousesByHouseIds(houseIDs)
	if err != nil {
		return housesMap, err
	}
	for _, house := range houses {
		housesMap[house.HouseId] = house
	}
	return housesMap, nil
}

// DeleteByHouseId 根据房产ID删除
func (a *House) DeleteByHouseId(houseId int64) errors.BizError {
	// 查找房产是否存在
	updateHouse, err := a.daoHouse.GetHouseByHouseId(houseId)
	if err != nil {
		return err
	}
	if updateHouse == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "房产id %d 不存在", houseId)
	}
	// 更新状态字段
	err = a.daoHouse.UpdateStatus(houseId, entity.HouseStatusDelete)
	if err != nil {
		return err
	}
	return nil
}

// GetHousesByLimit 分页获取房产列表
func (a *House) GetHousesByLimit(pageSize int, pageNum int, keywords *entity.HouseKeywords) (houses []*entity.HouseEntity, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, HouseListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	offset := (pageNum - 1) * pageSize
	// 无搜索词
	if keywords == nil {
		return a.daoHouse.GetHousesByLimit(pageSize, offset)
	}
	return a.daoHouse.GetHousesByKeywordAndLimit(pageSize, offset, keywords)
}

// GetAllHouses 获取所有的房产
func (a *House) GetAllHouses() (houses []*entity.HouseEntity, err errors.BizError) {
	return a.daoHouse.GetAllHouses()
}

// GetAllowSplitLeaseHouses 获取允许合租的房产
func (a *House) GetAllowSplitLeaseHouses() (houses []*entity.HouseEntity, err errors.BizError) {
	return a.daoHouse.GetAllowSplitLeaseHouses()
}

// GetSplitLeaseAndNotRentHouses 获取允许合租但是未出租的房产
func (a *House) GetSplitLeaseAndNotRentHouses() (houses []*entity.HouseEntity, err errors.BizError) {
	return a.daoHouse.GetSplitLeaseAndNotRentHouses()
}

// GetLeaseAndNotRentHouses 获取允许整租但没有出租的房产
func (a *House) GetLeaseAndNotRentHouses() (houses []*entity.HouseEntity, err errors.BizError) {
	return a.daoHouse.GetLeaseAndNotRentHouses()
}

// GetPageInfoLimit 获取分页信息
func (a *House) GetPageInfoLimit(pageSize int, pageNum int, keywords *entity.HouseKeywords) (pagination *entity.PageInfo, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, HouseListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	pageInfo := new(entity.PageInfo)
	var totalNum int64

	// 无搜索词
	if keywords == nil {
		totalNum, err = a.daoHouse.CountHouses()
	} else {
		totalNum, err = a.daoHouse.CountHousesByKeywords(keywords)
	}
	if err != nil {
		logger.WithContext(a.ctx).Warnf("[service.House] GetPageInfoLimit pageSize=%d pageNum=%d err=%s",
			pageSize, pageNum, err.Error())
		return pageInfo, err
	}
	return entity.GetPageInfo(totalNum, pageSize, pageNum), nil
}
