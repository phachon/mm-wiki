package entity

import "github.com/phachon/mm-wiki/utils"

const (
	LandlordStatusDefault = 0  // 业主状态 0 正常
	LandlordStatusDelete  = -1 // 业主状态 -1 删除
)

// LandlordEntity 业主表结构
type LandlordEntity struct {
	LandlordId   int64          `json:"landlord_id" gorm:"primary_key"` // 业主ID
	NickName     string         `json:"nick_name"`                      // 业主昵称
	GivenName    string         `json:"given_name"`                     // 业主姓名
	Sex          int            `json:"sex"`                            // 性别
	Email        string         `json:"email"`                          // 邮箱
	IdCardNumber string         `json:"id_card_number"`                 // 身份证
	Mobile       string         `json:"mobile"`                         // 手机号
	Address      string         `json:"address"`                        // 现住址
	Status       int            `json:"status"`                         // 状态
	CreateTime   utils.JsonTime `json:"create_time"`                    // 创建时间
	UpdateTime   utils.JsonTime `json:"update_time"`                    // 更新时间
}

// LandlordKeywords 业主搜索词
type LandlordKeywords struct {
	GivenName string `json:"given_name"` // 业主姓名
}

// LandlordEditResp 业主编辑返回结果
type LandlordEditResp struct {
	LandlordInfo *LandlordEntity `json:"landlord_info"` // 业主信息
}

// LandlordDetailResp 房产详情返回结果
type LandlordDetailResp struct {
	LandlordInfo *LandlordEntity `json:"landlord_info"` // 业主信息
}
