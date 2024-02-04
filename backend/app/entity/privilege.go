package entity

import (
	"github.com/phachon/mm-wiki/utils"
)

const (
	PrivilegeIsDisplay  = 0 // 权限是否显示 0 不显示
	PrivilegeNotDisplay = 1 // 权限是否显示 1 显示

	PrivilegeTypeOperation = 3 // 权限类型：操作
	PrivilegeTypeMenu      = 2 // 权限类型：菜单
	PrivilegeTypeNav       = 1 // 权限类型：导航
)

// PrivilegeEntity privilege 权限表结构
type PrivilegeEntity struct {
	PrivilegeId   int64          `json:"privilege_id" gorm:"primary_key"` // 权限ID
	Identify      string         `json:"identify"`                        // 权限标识（唯一）
	Name          string         `json:"name"`                            // 权限名
	ParentId      int64          `json:"parent_id"`                       // 上级ID
	ParentIds     string         `json:"parent_ids"`                      // 所有的上级ID , 隔开
	PrivilegeType int            `json:"privilege_type"`                  // 权限类型：1'nav',2'menu',3'操作'
	PageRouter    string         `json:"page_router"`                     // 页面路由path
	ApiMarks      string         `json:"api_marks"`                       // 接口标识（多个接口逗号隔开）
	Icon          string         `json:"icon"`                            // icon 图标
	IsDisplay     int            `json:"is_display"`                      // 是否显示：0不显示 1显示
	Sequence      int            `json:"sequence"`                        // 排序(越小越靠前)
	CreateTime    utils.JsonTime `json:"create_time"`                     // 创建时间
	UpdateTime    utils.JsonTime `json:"update_time"`                     // 更新时间
}

// PrivilegeListItem 权限列表 item 结构
type PrivilegeListItem struct {
	*PrivilegeEntity                      // 权限信息
	ChildPrivileges  []*PrivilegeListItem `json:"child_privileges"` // 子权限信息
	Action           *PrivilegeListAction `json:"action"`           // 权限操作
}

// PrivilegeListAction 权限列表操作权限
type PrivilegeListAction struct {
	IsEdit   int `json:"is_edit,omitempty"`   // 权限修改
	IsDelete int `json:"is_delete,omitempty"` // 权限删除
}
