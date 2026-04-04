package entity

import "github.com/phachon/mm-wiki/utils"

const (
	AttachmentSourceDefault = 0 // 附件来源 0 默认是附件
	AttachmentSourceImage   = 1 // 附件来源 1 图片
)

// AttachmentEntity 附件信息数据结构
type AttachmentEntity struct {
	AttachmentId int64          `json:"attachment_id" gorm:"primary_key"` // 附件ID
	AccountId    int64          `json:"account_id"`                       // 创建用户id
	DocId        string         `json:"doc_id"`                           // 所属文档id
	Name         string         `json:"name"`                             // 附件名称
	Path         string         `json:"path"`                             // 附件路径
	Source       int            `json:"source"`                           // 附件来源 0 默认是附件 1 图片
	CreateTime   utils.JsonTime `json:"create_time"`                      // 创建时间
	UpdateTime   utils.JsonTime `json:"update_time"`                      // 更新时间
}

// AttachmentListItem 附件列表结构
type AttachmentListItem struct {
	*AttachmentEntity
	AccountName string `json:"account_name"` // 上传者名称
}

// AttachmentKeywords 附件搜索关键字
type AttachmentKeywords struct {
	DocId string `json:"doc_id"` // 文档ID
	Name  string `json:"name"`   // 附件名称
}
