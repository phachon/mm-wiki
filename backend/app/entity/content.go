package entity

import "github.com/phachon/mm-wiki/utils"

// ContentEntity 文档内容表结构
type ContentEntity struct {
	DocId            int64          `json:"doc_id" gorm:"primary_key"` // 文档ID
	Content          string         `json:"content"`                   // 文档内容
	CurrentVersionID int64          `json:"current_version_id"`        // 当前版本ID
	CreateTime       utils.JsonTime `json:"create_time"`               // 创建时间
	UpdateTime       utils.JsonTime `json:"update_time"`               // 更新时间
}

// ContentVersionEntity 文档内容版本表结构
type ContentVersionEntity struct {
	VersionId     int64          `json:"version_id"`     // 版本ID
	DocId         int64          `json:"doc_id"`         // 文档ID
	VersionNumber int            `json:"version_number"` // 版本号
	Content       string         `json:"content"`        // 文档内容
	CreateTime    utils.JsonTime `json:"create_time"`    // 创建时间
}
