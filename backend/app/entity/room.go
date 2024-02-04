package entity

import "github.com/phachon/mm-wiki/utils"

const (
	RoomStatusDefault = 0  // 房间状态 0 正常
	RoomStatusDelete  = -1 // 房间状态 -1 删除

	RoomLeaseStatusNotRent = 0 // 房间出租状态 0 未出租
	RoomLeaseStatusRenting = 1 // 房间出租状态 1 已出租

	RoomDirectionTypeUnknown = 0 // 房间朝向 0 未知
	RoomDirectionTypeSouth   = 1 // 房间朝向 1 朝南
	RoomDirectionTypeEast    = 2 // 房间朝向 2 朝东
	RoomDirectionTypeWest    = 3 // 房间朝向 3 朝西
	RoomDirectionTypeNorth   = 4 // 房间朝向 4 朝北

	RoomTypeUnknown    = 0 // 房间类型 0 未知
	RoomTypeMasterRoom = 1 // 房间类型 1 主卧
	RoomTypeBedRoom    = 2 // 房间类型 2 次卧
	RoomTypeLivingRoom = 3 // 房间类型 3 客厅

	RoomHasNotToilet = 0 // 是否有卫生间 0 无
	RoomHasToilet    = 1 // 是否有卫生间 1 有

	RoomHasNotBalcony = 0 // 是否有阳台 0 无
	RoomHasBalcony    = 1 // 是否有阳台 1 有
)

// RoomEntity 房间表结构
type RoomEntity struct {
	RoomId        int64          `json:"room_id" gorm:"primary_key"` // 房间ID
	HouseId       int64          `json:"house_id"`                   // 所属房产ID
	Name          string         `json:"name"`                       // 房间名
	Area          int            `json:"area"`                       // 房间面积
	DirectionType int            `json:"direction_type"`             // 房间朝向
	RoomType      int            `json:"room_type"`                  // 房间类型
	HasToilet     int            `json:"has_toilet"`                 // 是否有卫生间
	HasBalcony    int            `json:"has_balcony"`                // 是否有阳台
	MouthRent     int            `json:"mouth_rent"`                 // 月租金
	Status        int            `json:"status"`                     // 删除状态
	LeaseStatus   int            `json:"lease_status"`               // 出租状态
	AllowLease    int            `json:"allow_lease"`                // 是否可出租 0 可以 -1 不可以
	CreateTime    utils.JsonTime `json:"create_time"`                // 创建时间
	UpdateTime    utils.JsonTime `json:"update_time"`                // 更新时间
}

// RoomListItem 房间列表 item 结构
type RoomListItem struct {
	*RoomEntity
	HouseInfo *HouseEntity `json:"house_info"` // 房产信息
}

// RoomKeywords 房间搜索词
type RoomKeywords struct {
	Name          string `json:"name"`           // 房间名
	DirectionType *int   `json:"direction_type"` // 房间朝向
	RoomType      *int   `json:"room_type"`      // 房间类型
	LeaseStatus   *int   `json:"lease_status"`   // 出租状态
	AllowLease    *int   `json:"allow_lease"`    // 允许出租
	HouseId       string `json:"house_id"`       // 房间ID
}

// RoomEditResp 房间编辑返回结果
type RoomEditResp struct {
	RoomInfo *RoomEntity    `json:"room_info"` // 房间信息
	Houses   []*HouseEntity `json:"houses"`    // 房产列表
}

// RoomDetailResp 房间详情返回结果
type RoomDetailResp struct {
	RoomInfo *RoomEntity `json:"room_info"` // 账号信息
}
