package global

const (
	PrivilegeIndentifyAccountEdit         = "system:account:edit"          // 账号修改
	PrivilegeIndentifyAccountDetail       = "system:account:detail"        // 账号删除
	PrivilegeIndentifyAccountUpdateStatus = "system:account:update_status" // 账号更新状态

	PrivilegeIndentifyRoleEdit          = "system:role:edit"           // 角色修改
	PrivilegeIndentifyRoleDelete        = "system:role:delete"         // 角色删除
	PrivilegeIndentifyRoleAccountList   = "system:role:account_list"   // 角色查看账号
	PrivilegeIndentifyRoleAccountRemove = "system:role:account_remove" // 角色账号移除
	PrivilegeIndentifyRolePrivilegeEdit = "system:role:privilege_edit" // 角色权限修改

	PrivilegeIndentifyPrivilegeEdit   = "system:privilege:edit"   // 权限修改
	PrivilegeIndentifyPrivilegeDelete = "system:privilege:delete" // 权限删除

	PrivilegeIndentifyNoticeEdit   = "system:notice:edit"   // 公告修改
	PrivilegeIndentifyNoticeDelete = "system:notice:delete" // 公告删除
)
