package entity

import "github.com/phachon/mm-wiki/utils"

const (
	DocEntityType = 1 // 文档
	DirEntityType = 2 // 目录
)

// DocEntity 文档表结构
type DocEntity struct {
	DocId             int64          `json:"doc_id" gorm:"primary_key"` // 文档ID
	ParentId          int64          `json:"parent_id"`                 // 父文档ID
	SpaceId           int64          `json:"space_id"`                  // 空间ID
	SpaceKey          string         `json:"space_key"`                 // 空间key
	Name              string         `json:"name"`                      // 文档名称
	Type              int            `json:"type"`                      // 文档类型 1 page 2 dir
	Path              string         `json:"path"`                      // 路径
	Sequence          int            `json:"sequence"`                  // 排序
	CreateAccountId   int64          `json:"create_account_id"`         // 创建账号ID
	CreateAccountName string         `json:"create_account_name"`       // 创建账号名
	EditAccountId     int64          `json:"edit_account_id"`           // 最后修改账号ID
	EditAccountName   string         `json:"edit_account_name"`         // 最后修改账号名
	Status            int            `json:"status"`                    // 状态 0 正常 -1 删除
	CreateTime        utils.JsonTime `json:"create_time"`               // 创建时间
	UpdateTime        utils.JsonTime `json:"update_time"`               // 更新时间
}

type DocTreeEntity struct {
	DocEntity
	Children []*DocTreeEntity `json:"children"` // 子文档
}
