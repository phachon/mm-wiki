package entity

import "github.com/phachon/mm-wiki/utils"

// ContentEntity 文档内容表结构
type ContentEntity struct {
	DocId      int64          `json:"doc_id" gorm:"primary_key"` // 文档ID
	Content    string         `json:"content"`                   // 文档内容
	CreateTime utils.JsonTime `json:"create_time"`               // 创建时间
	UpdateTime utils.JsonTime `json:"update_time"`               // 更新时间
}
