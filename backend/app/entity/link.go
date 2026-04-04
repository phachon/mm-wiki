package entity

import "github.com/phachon/mm-wiki/utils"

const (
	LinkStatusDefault = 0  // 链接状态 0 正常
	LinkStatusDelete  = -1 // 链接状态 -1 删除
)

// LinkEntity 快捷链接数据结构
type LinkEntity struct {
	LinkId     int64          `json:"link_id" gorm:"primary_key"` // 链接ID
	Name       string         `json:"name"`                       // 链接名称
	URL        string         `json:"url"`                        // 链接地址
	Sequence   int            `json:"sequence"`                   // 排序号(越小越靠前)
	Status     int            `json:"status"`                     // 状态 0 正常 -1 删除
	CreateTime utils.JsonTime `json:"create_time"`                // 创建时间
	UpdateTime utils.JsonTime `json:"update_time"`                // 修改时间
}

// LinkListItem 链接列表结构
type LinkListItem struct {
	*LinkEntity
	Action *LinkListAction `json:"action"` // 数据操作
}

// LinkListAction 链接列表操作权限
type LinkListAction struct {
	IsEdit   int `json:"is_edit,omitempty"`   // 链接修改
	IsDelete int `json:"is_delete,omitempty"` // 链接删除
}

// LinkKeywords 链接搜索关键字
type LinkKeywords struct {
	Name string // 搜索名称
}
