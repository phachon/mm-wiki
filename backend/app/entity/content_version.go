package entity

import "github.com/phachon/mm-wiki/utils"

// ContentVersionEntity 文档内容版本表结构
type ContentVersionEntity struct {
	ContentVersionId int64          `json:"content_version_id" gorm:"primary_key"` // 版本ID
	DocId            int64          `json:"doc_id"`                                // 文档ID
	Content          string         `json:"content"`                               // 文档内容
	CreateTime       utils.JsonTime `json:"create_time"`                           // 版本创建时间
	UpdateTime       utils.JsonTime `json:"update_time"`                           // 文档更新时间
	EditAccountId    int64          `json:"edit_account_id"`                       // 文档修改账号ID
	EditAccountName  string         `json:"edit_account_name"`                     // 文档修改账号名
}
