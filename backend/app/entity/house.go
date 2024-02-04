package entity

import "github.com/phachon/mm-wiki/utils"

const (
	HouseStatusDefault = 0  // 房产状态 0 正常
	HouseStatusDelete  = -1 // 房产状态 -1 删除

	HouseLeaseStatusNotRent = 0 // 房产出租状态 0 未出租
	HouseLeaseStatusRenting = 1 // 房产出租状态 1 已出租

	HouseAllowLeaseYes = 0  // 房产允许整租 0 是
	HouseAllowLeaseNO  = -1 // 房产允许整租 -1 否

	HouseAllowSplitLeaseYes = 0  // 房产允许合租 0 是
	HouseAllowSplitLeaseNO  = -1 // 房产允许合租 -1 否

	HouseDecorationTypeDefault = 0 // 装修类型 0 未知
	HouseDecorationTypeNot     = 1 // 装修类型 1 毛胚
	HouseDecorationTypeLittle  = 2 // 装修类型 2 简装
	HouseDecorationTypeBetter  = 3 // 装修类型 3 精装
	HouseDecorationTypeBest    = 4 // 装修类型 4 豪华

	HouseSizeTypeDefault          = 0 // 户型 0 未知
	HouseSizeTypeOneRoomOneHall   = 1 // 户型 1 一室一厅
	HouseSizeTypeTwoRoomOneHall   = 2 // 户型 2 两室一厅
	HouseSizeTypeThreeRoomOneHall = 3 // 户型 3 三室一厅
	HouseSizeTypeFourRoomOneHall  = 4 // 户型 4 四室一厅
	HouseSizeTypeFourRoomTwoHall  = 5 // 户型 5 四室二厅
	HouseSizeTypeOneRoom          = 6 // 户型 6 开间
)

// HouseEntity 房产表结构
type HouseEntity struct {
	HouseId         int64          `json:"house_id" gorm:"primary_key"` // 房产ID
	LandlordId      int64          `json:"landlord_id"`                 // 业主ID
	Idx             string         `json:"idx"`                         // 唯一编号
	Region          string         `json:"region"`                      // 所在区域
	Address         string         `json:"address"`                     // 所在地址
	HouseNumber     string         `json:"house_number"`                // 门牌号
	DecorationType  int            `json:"decoration_type"`             // 装修类型
	SizeType        int            `json:"size_type"`                   // 房产户型
	Area            int            `json:"area"`                        // 房产面积
	BuildTime       string         `json:"build_time"`                  // 建成日期
	Status          int            `json:"status"`                      // 删除状态
	AllowLease      int            `json:"allow_lease"`                 // 是否可整租 0 是 -1 否
	AllowSplitLease int            `json:"allow_split_lease"`           // 是否可合租 0 是 -1 否
	LeaseStatus     int            `json:"lease_status"`                // 出租状态
	MouthRent       int            `json:"mouth_rent"`                  // 整租月租金
	CreateTime      utils.JsonTime `json:"create_time"`                 // 创建时间
	UpdateTime      utils.JsonTime `json:"update_time"`                 // 更新时间
}

// HouseKeywords 房产搜索词
type HouseKeywords struct {
	Idx            string `json:"idx"`             // 房产编号
	Region         string `json:"region"`          // 房产区域
	Address        string `json:"address"`         // 房产地址
	DecorationType *int   `json:"decoration_type"` // 装修类型
	SizeType       *int   `json:"size_type"`       // 户型
	LeaseStatus    *int   `json:"lease_status"`    // 出租状态
}

// HouseEditResp 房产编辑返回结果
type HouseEditResp struct {
	HouseInfo  *HouseEntity      `json:"house_info"`  // 房产信息
	LandlordId int64             `json:"landlord_id"` // 房产业主id
	Landlords  []*LandlordEntity `json:"landlords"`   // 业主列表
}

// HouseDetailResp 房产详情返回结果
type HouseDetailResp struct {
	HouseInfo *HouseEntity `json:"house_info"` // 房产信息
}

// HouseLandlordDetailResp 房产业主详情返回结果
type HouseLandlordDetailResp struct {
	LandlordInfo *LandlordEntity `json:"landlord_info"` // 业主信息
}
