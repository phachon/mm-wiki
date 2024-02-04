// package event 事件分发相关逻辑
package event

const (
	ActionTypeAuthLogin = "auth:login" // 账号登录

	ActionTypeProfileRepass = "profile:repass" // 修改密码
	ActionTypeProfileUpdate = "profile:update" // 修改资料

	ActionTypeAccountAdd          = "account:add"           // 账号添加
	ActionTypeAccountEdit         = "account:edit"          // 账号修改
	ActionTypeAccountUpdateStatus = "account:update_status" // 账号状态更新

	ActionTypeRoleAdd             = "role:add"              // 角色添加
	ActionTypeRoleEdit            = "role:edit"             // 角色修改
	ActionTypeRoleAccountRemove   = "role:account_remove"   // 移除角色账号
	ActionTypeRoleDelete          = "role:delete"           // 角色删除
	ActionTypeRolePrivilegeUpdate = "role:privilege_update" // 角色权限更新

	ActionTypePrivilegeAdd    = "privilege:add"    // 权限添加
	ActionTypePrivilegeEdit   = "privilege:edit"   // 权限修改
	ActionTypePrivilegeDelete = "privilege:delete" // 权限删除

)

// ActionTypeMessage 操作类型对应文案
var ActionTypeMessage = map[string]string{
	ActionTypeAuthLogin:           "账号登录成功",
	ActionTypeProfileRepass:       "修改密码成功",
	ActionTypeProfileUpdate:       "修改个人信息成功",
	ActionTypeAccountAdd:          "添加账号成功",
	ActionTypeAccountEdit:         "修改账号成功",
	ActionTypeAccountUpdateStatus: "更新账号状态成功",
	ActionTypeRoleAdd:             "添加角色成功",
	ActionTypeRoleEdit:            "修改角色成功",
	ActionTypeRoleAccountRemove:   "移除角色下账号成功",
	ActionTypeRoleDelete:          "删除角色成功",
	ActionTypeRolePrivilegeUpdate: "更新角色下权限成功",
	ActionTypePrivilegeAdd:        "添加权限成功",
	ActionTypePrivilegeEdit:       "修改权限成功",
	ActionTypePrivilegeDelete:     "删除权限成功",
}
