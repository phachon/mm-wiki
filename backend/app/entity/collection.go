package entity

import "github.com/phachon/mm-wiki/utils"

const (
	CollectionTypeUnknown  = 0 // 收藏类型 0 未知
	CollectionTypeDocument = 1 // 收藏类型 1 文档
	CollectionTypeSpace    = 2 // 收藏类型 2 空间
)

const (
	CollectionActionAdd = 1 // 收藏操作 1 添加
	CollectionActionDel = 2 // 收藏操作 2 取消
)

// CollectionEntity 收藏表结构
type CollectionEntity struct {
	CollectionId   int64          `json:"collection_id" gorm:"primary_key"` // 文档集合ID
	AccountId      int64          `json:"account_id"`                       // 账号ID
	CollectionType int            `json:"collection_type"`                  // 收藏类型 1 文档 2 空间
	ResourceId     string         `json:"resource_id"`                      // 收藏资源ID
	CreateTime     utils.JsonTime `json:"create_time"`                      // 创建时间
}
