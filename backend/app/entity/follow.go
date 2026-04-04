package entity

import "github.com/phachon/mm-wiki/utils"

const (
	FollowTypeUnknown  = 0 // 关注类型 0 未知
	FollowTypeDocument = 1 // 关注类型 1 文档
	FollowTypeUser     = 2 // 关注类型 2 用户
)

// FollowEntity 关注表结构
type FollowEntity struct {
	FollowId   int64          `json:"follow_id" gorm:"primary_key"` // 关注ID
	AccountId  int64          `json:"account_id"`                   // 账号ID
	FollowType int            `json:"follow_type"`                  // 关注类型 1 文档 2 用户
	ObjectId   string         `json:"object_id"`                    // 关注对象ID
	CreateTime utils.JsonTime `json:"create_time"`                  // 创建时间
}
