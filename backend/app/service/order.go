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
	OrderListDefaultPageSize = 20 // 订单列表默认一页 20 条
)

// Order 订单业务逻辑
type Order struct {
	ctx      context.Context
	daoOrder *dao.Order
	daoHouse *dao.House
	daoRoom  *dao.Room
}

// NewOrder 创建订单业务逻辑对象
func NewOrder(ctx context.Context) *Order {
	return &Order{
		ctx:      ctx,
		daoOrder: dao.NewOrder(ctx),
		daoHouse: dao.NewHouse(ctx),
		daoRoom:  dao.NewRoom(ctx),
	}
}

// Create 创建订单
func (o *Order) Create(orderEntity *entity.OrderEntity) errors.BizError {

	// 获取房产信息
	house, err := o.daoHouse.GetHouseByHouseId(orderEntity.HouseId)
	if err != nil {
		return err
	}
	if house == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "房产不存在")
	}
	// 房产是否已出租
	if house.LeaseStatus == entity.HouseLeaseStatusRenting {
		return errors.Errorf(errors.BusinessPermissionError, "房产已被出租")
	}
	// 整租是否允许
	if orderEntity.OrderType == entity.OrderTypeWhole && house.AllowLease == entity.HouseAllowLeaseNO {
		return errors.Errorf(errors.BusinessPermissionError, "房产不允许出租")
	}
	// 合租是否允许
	if orderEntity.OrderType == entity.OrderTypeSplit && house.AllowSplitLease == entity.HouseAllowSplitLeaseNO {
		return errors.Errorf(errors.BusinessPermissionError, "房产不允许合租")
	}

	// 合租获取房间信息
	if orderEntity.RoomId > 0 {
		room, err := o.daoRoom.GetRoomByRoomId(orderEntity.RoomId)
		if err != nil {
			return err
		}
		if room == nil {
			return errors.Errorf(errors.BusinessRecordNotExistError, "房间不存在")
		}
		// 房间是否已出租
		if room.LeaseStatus == entity.RoomLeaseStatusRenting {
			return errors.Errorf(errors.BusinessPermissionError, "房间已被出租")
		}
	}

	// 插入一条订单记录
	err = o.daoOrder.Insert(orderEntity)
	if err != nil {
		return err
	}

	// 无论是合租还是整租都需要更新房产状态
	err = o.daoHouse.UpdateLeaseStatus(orderEntity.HouseId, entity.HouseLeaseStatusRenting)
	if err != nil {
		return err
	}

	// 合租更新房间状态
	if orderEntity.OrderType == entity.OrderTypeSplit {
		err := o.daoRoom.UpdateLeaseStatus(orderEntity.RoomId, entity.RoomLeaseStatusRenting)
		if err != nil {
			return err
		}
	}
	return nil
}

// Update 修改订单
func (o *Order) Update(orderEntity entity.OrderEntity) errors.BizError {
	// 查找订单是否存在
	updateOrder, err := o.daoOrder.GetOrderByOrderId(orderEntity.OrderId)
	if err != nil {
		return err
	}
	if updateOrder == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "订单不存在")
	}

	// 订单上已结束/已退租/已取消的无法再修改
	if updateOrder.Status == entity.OrderStatusFinished ||
		updateOrder.Status == entity.OrderStatusVacated ||
		updateOrder.Status == entity.OrderStatusCanceled {
		return errors.Errorf(errors.BusinessForbiddenError, "已结束订单不允许修改")
	}

	// 更新字段
	err = o.daoOrder.Update(orderEntity)
	if err != nil {
		return err
	}
	return nil
}

// GetOrderByOrderId 根据订单ID获取订单详情
func (o *Order) GetOrderByOrderId(orderID int64) (order *entity.OrderEntity, err errors.BizError) {
	order, err = o.daoOrder.GetOrderByOrderId(orderID)
	if err != nil {
		return order, err
	}
	return order, nil
}

// UpdateStatus 根据订单ID更新状态
func (o *Order) UpdateStatus(orderId int64, status int) errors.BizError {
	// 查找订单是否存在
	updateOrder, err := o.daoOrder.GetOrderByOrderId(orderId)
	if err != nil {
		return err
	}
	if updateOrder == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "订单id %d 不存在", orderId)
	}

	// 这里需要判断状态合法性
	// 更新状态字段
	err = o.daoOrder.UpdateStatus(orderId, status)
	if err != nil {
		return err
	}
	return nil
}

// GetOrdersByLimit 分页获取订单列表
func (o *Order) GetOrdersByLimit(pageSize int, pageNum int, keywords *entity.OrderKeywords) (orders []*entity.OrderEntity, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, OrderListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	offset := (pageNum - 1) * pageSize
	// 无搜索词
	if keywords == nil {
		return o.daoOrder.GetOrdersByLimit(pageSize, offset)
	}
	return o.daoOrder.GetOrdersByKeywordAndLimit(pageSize, offset, keywords)
}

// GetPageInfoLimit 获取分页信息
func (a *Order) GetPageInfoLimit(pageSize int, pageNum int, keywords *entity.OrderKeywords) (pagination *entity.PageInfo, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, OrderListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	pageInfo := new(entity.PageInfo)
	var totalNum int64

	// 无搜索词
	if keywords == nil {
		totalNum, err = a.daoOrder.CountOrders()
	} else {
		totalNum, err = a.daoOrder.CountOrdersByKeywords(keywords)
	}
	if err != nil {
		logger.WithContext(a.ctx).Warnf("[service.Order] GetPageInfoLimit pageSize=%d pageNum=%d err=%s",
			pageSize, pageNum, err.Error())
		return pageInfo, err
	}
	return entity.GetPageInfo(totalNum, pageSize, pageNum), nil
}

// FormatOrderList 获取订单列表
func (a *Order) FormatOrderList(orders []*entity.OrderEntity) ([]*entity.OrderListItem, errors.BizError) {
	var (
		houseIds  []int64
		roomIds   []int64
		hirerIds  []int64
		orderList []*entity.OrderListItem
	)
	for _, order := range orders {
		houseIds = append(houseIds, order.HouseId)
		roomIds = append(roomIds, order.RoomId)
		hirerIds = append(hirerIds, order.HirerId)
	}
	// 获取房产信息
	housesMap, err := NewHouse(a.ctx).GetHousesMapByHouseIds(houseIds)
	if err != nil {
		return orderList, err
	}
	// 获取房间信息
	roomsMap, err := NewRoom(a.ctx).GetRoomsMapByRoomIds(roomIds)
	if err != nil {
		return orderList, err
	}
	// 获取租客信息
	hirersMap, err := NewHirer(a.ctx).GetHirersMapByHirerIds(hirerIds)
	if err != nil {
		return orderList, err
	}
	for _, order := range orders {
		var orderItem = &entity.OrderListItem{
			OrderEntity: order,
		}
		if house, ok := housesMap[order.HouseId]; ok {
			orderItem.HouseInfo = house
		}
		if room, ok := roomsMap[order.RoomId]; ok {
			orderItem.RoomInfo = room
		}
		if hirer, ok := hirersMap[order.HirerId]; ok {
			orderItem.HirerInfo = hirer
		}
		orderList = append(orderList, orderItem)
	}
	return orderList, nil
}
