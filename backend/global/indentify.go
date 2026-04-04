package global

// PrivilegeIndentify 操作权限标识
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

	PrivilegeIndentifyContactEdit   = "system:contact:edit"   // 联系人修改
	PrivilegeIndentifyContactDelete = "system:contact:delete" // 联系人删除

	PrivilegeIndentifyLinkEdit   = "system:link:edit"   // 链接修改
	PrivilegeIndentifyLinkDelete = "system:link:delete" // 链接删除

	PrivilegeIndentifyEmailEdit   = "system:email:edit"   // 邮箱修改
	PrivilegeIndentifyEmailDelete = "system:email:delete" // 邮箱删除

	PrivilegeIndentifyDepartmentAdd    = "system:department:add"    // 部门添加
	PrivilegeIndentifyDepartmentEdit   = "system:department:edit"   // 部门修改
	PrivilegeIndentifyDepartmentDelete = "system:department:delete" // 部门删除

	PrivilegeIndentifySpaceEdit   = "space:edit"   // 空间修改
	PrivilegeIndentifySpaceDelete = "space:delete" // 空间删除

	PrivilegeIndentifyLoginAuthEdit   = "system:login_auth:edit"   // 登录认证修改
	PrivilegeIndentifyLoginAuthDelete = "system:login_auth:delete" // 登录认证删除
)
