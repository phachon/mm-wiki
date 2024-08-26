package entity

import "github.com/phachon/mm-wiki/utils"

const (
	SpacePermissionRelationTypeAccount    = 0 // 关联类型 账号
	SpacePermissionRelationTypeDepartment = 1 // 关联类型 部门
)

const (
	SpacePermissionIsNotView = 0 // 是否允许查看
	SpacePermissionIsView    = 1 // 是否允许查看

	SpacePermissionIsNotAdd = 0 // 是否允许添加
	SpacePermissionIsAdd    = 1 // 是否允许添加

	SpacePermissionIsNotEdit = 0 // 是否允许编辑
	SpacePermissionIsEdit    = 1 // 是否允许编辑

	SpacePermissionIsNotDelete = 0 // 是否允许删除
	SpacePermissionIsDelete    = 1 // 是否允许删除

	SpacePermissionIsNotAdmin = 0 // 是否是管理员
	SpacePermissionIsAdmin    = 1 // 是否是管理员
)

// SpacePermissionEntity 空间权限表结构
type SpacePermissionEntity struct {
	SpacePermissionId int64          `json:"space_permission_id" gorm:"primary_key"` // 空间权限ID
	SpaceId           int64          `json:"space_id"`                               // 空间ID
	RelationType      int            `json:"relation_type"`                          // 关联类型 0 个人 1 部门
	AccountId         int64          `json:"account_id"`                             // 账号ID
	DepartmentId      int64          `json:"department_id"`                          // 部门ID
	IsView            int            `json:"is_view"`                                // 是否允许查看
	IsAdd             int            `json:"is_add"`                                 // 是否允许添加
	IsEdit            int            `json:"is_edit"`                                // 是否允许编辑
	IsDelete          int            `json:"is_delete"`                              // 是否允许删除
	IsAdmin           int            `json:"is_admin"`                               // 是否是管理员
	CreateTime        utils.JsonTime `json:"create_time"`                            // 创建时间
	UpdateTime        utils.JsonTime `json:"update_time"`                            // 更新时间
}

// SpacePermission 空间权限
type SpacePermission struct {
	IsView   int // 是否允许查看
	IsAdd    int // 是否允许添加
	IsEdit   int // 是否允许编辑
	IsDelete int // 是否允许删除
	IsAdmin  int // 是否是管理员
}
