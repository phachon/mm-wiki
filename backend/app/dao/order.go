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
	// TableNameOrder 租赁订单关系表
	TableNameOrder = "hms_order"
	// OrderPrimaryKey 租赁订单关系表主键ID
	OrderPrimaryKey = "order_id"
)

// Order 租赁订单关系表数据
type Order struct {
	ctx context.Context
	db  *gorm.DB
}

// NewOrder 创建系统业主表数据对象
func NewOrder(ctx context.Context) *Order {
	return &Order{
		ctx: ctx,
		db:  GetDB(dbNameKms).WithContext(ctx),
	}
}

// Insert 创建业主插入一条业主记录
func (l *Order) Insert(orderEntity *entity.OrderEntity) errors.BizError {
	orderEntity.CreateTime = utils.NewJsonTime(time.Now())
	orderEntity.UpdateTime = utils.NewJsonTime(time.Now())
	db := l.db.Table(TableNameOrder).
		Save(orderEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// DeleteByhirerId 通过租客ID删除
func (l *Order) DeleteByhirerId(hirerId int64) errors.BizError {
	db := l.db.
		Table(TableNameOrder).
		Where("hirer_id = ?", hirerId).
		Delete(entity.OrderEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}

// DeleteByOrderId 通过订单ID删除
func (l *Order) DeleteByOrderId(orderId int64) errors.BizError {
	db := l.db.
		Table(TableNameOrder).
		Where("order_id = ?", orderId).
		Delete(entity.OrderEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}

// DeleteByRoomId 通过房间ID删除
func (l *Order) DeleteByRoomId(roomId int64) errors.BizError {
	db := l.db.
		Table(TableNameOrder).
		Where("room_id = ?", roomId).
		Delete(entity.OrderEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}

// GetOrdersByHirerId 根据租客ID获取租赁订单
func (l *Order) GetOrdersByHirerId(hirerId int64) (orders []*entity.OrderEntity,
	err errors.BizError) {
	db := l.db.Table(TableNameOrder).
		Where("hirer_id = ?", hirerId).
		Find(&orders)
	if db.Error != nil {
		return orders, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return orders, nil
}

// GetOrderByOrderId 根据订单ID获取租赁订单
func (l *Order) GetOrderByOrderId(orderId int64) (order *entity.OrderEntity,
	err errors.BizError) {
	order = &entity.OrderEntity{}
	db := l.db.Table(TableNameOrder).
		Where("order_id = ?", orderId).
		First(&order)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return order, nil
}

// GetOrdersByRoomId 根据房间ID获取租赁订单
func (l *Order) GetOrdersByRoomId(roomId int64) (orders []*entity.OrderEntity,
	err errors.BizError) {
	db := l.db.Table(TableNameOrder).
		Where("room_id = ?", roomId).
		Find(&orders)
	if db.Error != nil {
		return orders, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return orders, nil
}

// Update 更新订单，只会更新如下字段
func (l *Order) Update(order entity.OrderEntity) errors.BizError {
	order.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(l.ctx).Table(TableNameOrder).
		Select(
			"StartTime", "EndTime", "UnitRent",
			"PayType", "Status", "Remarks",
			"UpdateTime",
		).
		Where(map[string]interface{}{
			OrderPrimaryKey: order.OrderId,
		}).
		Updates(order)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// UpdateStatus 更新订单状态
func (l *Order) UpdateStatus(orderId int64, status int) errors.BizError {
	updateTime := utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(l.ctx).Table(TableNameOrder).
		Where(map[string]interface{}{
			OrderPrimaryKey: orderId,
		}).
		Update("status", status).
		Update("update_time", updateTime)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// GetOrdersByKeywords 根据订单名模糊匹配订单
func (l *Order) GetOrdersByKeywords(keywords *entity.OrderKeywords) (orders []*entity.OrderEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(l.ctx).Table(TableNameOrder)
	db = l.getOrderWhereByKeywords(keywords, db)
	db.Find(&orders)
	if db.Error != nil {
		return orders, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return orders, nil
}

func (l *Order) getOrderWhereByKeywords(keywords *entity.OrderKeywords, db *gorm.DB) *gorm.DB {
	if keywords.OrderId != "" {
		db = db.Where("order_id = ?", keywords.OrderId)
	}
	if keywords.HouseId != "" {
		db = db.Where("house_id = ?", keywords.HouseId)
	}
	if keywords.HirerId != "" {
		db = db.Where("hirer_id = ?", keywords.HirerId)
	}
	if keywords.PayType != nil && *keywords.PayType >= 0 {
		db = db.Where("pay_type = ?", *keywords.PayType)
	}
	if keywords.OrderType != nil && *keywords.OrderType >= 0 {
		db = db.Where("order_type = ?", *keywords.OrderType)
	}
	if keywords.OrderStatus != nil && *keywords.OrderStatus >= 0 {
		db = db.Where("status = ?", *keywords.OrderStatus)
	}
	return db
}

// GetOrdersByKeywordAndLimit 根据关键字分页获取订单
func (l *Order) GetOrdersByKeywordAndLimit(limit int, offset int, keywords *entity.OrderKeywords) (orders []*entity.OrderEntity, err errors.BizError) {
	if keywords == nil {
		return l.GetOrdersByLimit(limit, offset)
	}
	db := GetDB(dbNameKms).WithContext(l.ctx).Table(TableNameOrder)
	db = l.getOrderWhereByKeywords(keywords, db)
	db = db.Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", OrderPrimaryKey)).
		Find(&orders)
	if db.Error != nil {
		return orders, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return orders, nil
}

// GetOrdersByLimit 分页获取订单列表
func (l *Order) GetOrdersByLimit(limit int, offset int) (orders []*entity.OrderEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(l.ctx).
		Table(TableNameOrder).
		Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", OrderPrimaryKey)).
		Find(&orders)
	if db.Error != nil {
		return orders, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return orders, nil
}

// CountOrders 订单总数
func (l *Order) CountOrders() (count int64, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(l.ctx).
		Table(TableNameOrder).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// CountOrdersByKeywords 根据关键词搜索订单总数
func (l *Order) CountOrdersByKeywords(keywords *entity.OrderKeywords) (count int64, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(l.ctx).Table(TableNameOrder)
	db = l.getOrderWhereByKeywords(keywords, db)
	db.Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}
