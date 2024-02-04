package entity

import "github.com/phachon/mm-wiki/utils"

const (
	RoleStatusDefault = 0   // 角色状态 0 正常
	RoleStatusDelete  = -1  // 角色状态 -1 删除
	RoleStatusUnknown = -10 // 角色状态 未知

	RoleTypeCustomRole         = 0 // 自定义角色
	RoleTypeAccountDefaultRole = 1 // 账号默认角色
)

// RoleEntity role 角色表结构
type RoleEntity struct {
	RoleId       int64          `json:"role_id" gorm:"primary_key"` // 角色ID
	Name         string         `json:"name"`                       // 角色名
	Remark       string         `json:"remark"`                     // 角色备注
	RoleType     int            `json:"role_type"`                  // 角色类型
	PrivilegeIds string         `json:"privilege_ids"`              // 权限ID
	Status       int            `json:"status"`                     // 角色状态
	CreateTime   utils.JsonTime `json:"create_time"`                // 创建时间
	UpdateTime   utils.JsonTime `json:"update_time"`                // 更新时间
}

// RoleListItem 角色列表结构
type RoleListItem struct {
	*RoleEntity
	Action *RoleListAction `json:"action"` // 数据操作
}

// RoleListAction 角色列表操作权限
type RoleListAction struct {
	IsEdit          int `json:"is_edit,omitempty"`           // 角色修改
	IsDelete        int `json:"is_delete,omitempty"`         // 角色删除
	IsAccountList   int `json:"is_account_list,omitempty"`   // 账号列表
	IsPrivilegeEdit int `json:"is_privilege_edit,omitempty"` // 权限修改
}

// RoleKeywords 角色搜索词
type RoleKeywords struct {
	RoleName string `json:"role_name"` // 角色名称
}

// RolePrivilegeEntity 角色权限表结构
type RolePrivilegeEntity struct {
	RolePrivilegeId int64          `json:"role_privilege_id" gorm:"primary_key"` // 角色权限ID
	RoleId          int64          `json:"role_id"`                              // 角色ID
	PrivilegeId     int64          `json:"privilege_id"`                         // 权限ID
	CreateTime      utils.JsonTime `json:"create_time"`                          // 创建时间
}
