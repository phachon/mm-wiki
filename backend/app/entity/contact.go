package entity

import "github.com/phachon/mm-wiki/utils"

const (
	ContactStatusDefault = 0  // 联系人状态 0 正常
	ContactStatusDelete  = -1 // 联系人状态 -1 删除
)

// ContactEntity 联系人数据结构
type ContactEntity struct {
	ContactId  int64          `json:"contact_id" gorm:"primary_key"` // 联系人ID
	Name       string         `json:"name"`                          // 联系人名称
	Mobile     string         `json:"mobile"`                        // 联系电话
	Email      string         `json:"email"`                         // 邮箱
	Position   string         `json:"position"`                      // 联系人职位
	Status     int            `json:"status"`                        // 状态 0 正常 -1 删除
	CreateTime utils.JsonTime `json:"create_time"`                   // 创建时间
	UpdateTime utils.JsonTime `json:"update_time"`                   // 修改时间
}

// ContactListItem 联系人列表结构
type ContactListItem struct {
	*ContactEntity
	Action *ContactListAction `json:"action"` // 数据操作
}

// ContactListAction 联系人列表操作权限
type ContactListAction struct {
	IsEdit   int `json:"is_edit,omitempty"`   // 联系人修改
	IsDelete int `json:"is_delete,omitempty"` // 联系人删除
}

// ContactKeywords 联系人搜索关键字
type ContactKeywords struct {
	Name string // 搜索名称
}
