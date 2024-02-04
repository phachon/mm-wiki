package entity

import "github.com/phachon/mm-wiki/utils"

const (
	HirerStatusDefault = 0  // 租客状态 0 正常
	HirerStatusDelete  = -1 // 租客状态 -1 删除
)

// HirerEntity 租客表结构
type HirerEntity struct {
	HirerId          int64          `json:"hirer_id" gorm:"primary_key"` // 租客ID
	Name             string         `json:"name"`                        // 租客姓名
	Sex              int            `json:"sex"`                         // 性别
	Email            string         `json:"email"`                       // 邮箱
	IdCardNumber     string         `json:"id_card_number"`              // 身份证
	Mobile           string         `json:"mobile"`                      // 手机号
	EmergencyContact string         `json:"emergency_contact"`           // 紧急联系人
	EmergencyMobile  string         `json:"emergency_mobile"`            // 紧急联系人电话
	Status           int            `json:"status"`                      // 状态
	CreateTime       utils.JsonTime `json:"create_time"`                 // 创建时间
	UpdateTime       utils.JsonTime `json:"update_time"`                 // 更新时间
}

// HirerKeywords 租客搜索词
type HirerKeywords struct {
	Name string `json:"given_name"` // 租客姓名
}

// HirerEditResp 租客编辑返回结果
type HirerEditResp struct {
	HirerInfo *HirerEntity `json:"hirer_info"` // 租客信息
}

// HirerDetailResp 房产详情返回结果
type HirerDetailResp struct {
	HirerInfo *HirerEntity `json:"hirer_info"` // 租客信息
}
