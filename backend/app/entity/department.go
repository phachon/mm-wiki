package entity

import "github.com/phachon/mm-wiki/utils"

// DepartmentEntity 部门表结构
type DepartmentEntity struct {
	DepartmentId int64          `json:"department_id" gorm:"primary_key"` // 部门ID
	Name         string         `json:"name"`                             // 部门名
	ParentId     int64          `json:"parent_id"`                        // 上级ID
	ParentIds    string         `json:"parent_ids"`                       // 所有的上级ID , 隔开
	Sequence     int            `json:"sequence"`                         // 排序(越小越靠前)
	CreateTime   utils.JsonTime `json:"create_time"`                      // 创建时间
	UpdateTime   utils.JsonTime `json:"update_time"`                      // 更新时间
}

// DepartmentListItem 部门列表 item 结构
type DepartmentListItem struct {
	*DepartmentEntity                       // 部门信息
	Children          []*DepartmentListItem `json:"children"` // 子部门信息
	Action            *DepartmentListAction `json:"action"`   // 部门操作
}

// DepartmentListAction 部门列表操作部门
type DepartmentListAction struct {
	IsAdd    int `json:"is_add,omitempty"`    // 部门添加
	IsEdit   int `json:"is_edit,omitempty"`   // 部门修改
	IsDelete int `json:"is_delete,omitempty"` // 部门删除
}
