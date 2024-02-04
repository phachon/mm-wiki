package entity

import "github.com/phachon/mm-wiki/utils"

const (
	OrderStatusDefaultPending = 0 // 订单状态 0 待入住
	OrderStatusOccupied       = 1 // 订单状态 1 已入住
	OrderStatusRenewed        = 2 // 订单状态 2 已续租
	OrderStatusVacated        = 3 // 订单状态 3 已退租
	OrderStatusFinished       = 4 // 订单状态 4 已结束
	OrderStatusCanceled       = 5 // 订单状态 5 已取消

	OrderTypeWhole = 1 // 订单类型 整租
	OrderTypeSplit = 2 // 订单类型 合租
)

// OrderEntity 订单表结构
type OrderEntity struct {
	OrderId     int64          `json:"order_id" gorm:"primary_key"` // 订单id
	HouseId     int64          `json:"house_id"`                    // 房产id
	RoomId      int64          `json:"room_id"`                     // 房间id
	OrderType   int            `json:"order_type"`                  // 订单类型
	HirerId     int64          `json:"hirer_id"`                    // 租客id
	StartTime   utils.JsonDate `json:"start_time"`                  // 开始时间
	EndTime     utils.JsonDate `json:"end_time"`                    // 结束时间
	UnitRent    int            `json:"unit_rent"`                   // 单位租金
	PayType     int            `json:"pay_type"`                    // 支付类型
	Remarks     string         `json:"remarks"`                     // 备注信息
	AccountId   int64          `json:"account_id"`                  // 操作账号id
	AccountName string         `json:"account_name"`                // 操作账号名
	Status      int            `json:"status"`                      // 状态
	CreateTime  utils.JsonTime `json:"create_time"`                 // 创建时间
	UpdateTime  utils.JsonTime `json:"update_time"`                 // 更新时间
}

// OrderKeywords 订单搜索词
type OrderKeywords struct {
	OrderId     string `json:"order_id"`     // 订单ID
	HouseId     string `json:"house_id"`     // 房产ID
	HirerId     string `json:"hirer_id"`     // 租客ID
	OrderType   *int   `json:"order_type"`   // 订单类型
	PayType     *int   `json:"pay_type"`     // 支付类型
	OrderStatus *int   `json:"order_status"` // 订单状态
}

// OrderEditResp 订单编辑返回结果
type OrderEditResp struct {
	OrderInfo *OrderEntity `json:"order_info"` // 订单信息
}

// OrderDetailResp 订单详情返回结果
type OrderDetailResp struct {
	OrderInfo *OrderEntity `json:"order_info"` // 订单信息
}

// OrderAddResp 订单添加页面返回
type OrderAddResp struct {
	SelectHouses     []*HouseEntity          `json:"select_houses"`      // 选择房产列表
	SelectHouseRooms map[int64][]*RoomEntity `json:"select_house_rooms"` // 选择房产房间列表
}

// OrderListItem 订单列表 item 结构
type OrderListItem struct {
	*OrderEntity
	HouseInfo *HouseEntity `json:"house_info"` // 房产信息
	RoomInfo  *RoomEntity  `json:"room_info"`  // 房间信息
	HirerInfo *HirerEntity `json:"hirer_info"` // 租客信息
}
