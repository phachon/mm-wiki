package entity

import "github.com/phachon/mm-wiki/utils"

const (
	AccountStatusDefault = 0  // 账号状态 0 正常
	AccountStatusForbid  = -1 // 账号状态 -1 禁用
)

// AccountEntity account 账号表结构
type AccountEntity struct {
	AccountId  int64          `json:"account_id" gorm:"primary_key"` // 账号ID
	Name       string         `json:"name"`                          // 账号名
	Password   string         `json:"-"`                             // 密码
	GivenName  string         `json:"given_name"`                    // 昵称
	Mobile     string         `json:"mobile"`                        // 电话
	Phone      string         `json:"phone"`                         // 手机号
	Email      string         `json:"email"`                         // 邮箱
	Department string         `json:"department"`                    // 部门
	Position   string         `json:"position"`                      // 职位
	Location   string         `json:"location"`                      // 办公位
	LastIP     string         `json:"last_ip"`                       // 最后登录 IP
	LastTime   int64          `json:"last_time"`                     // 最后登录时间
	Status     int            `json:"status"`                        // 状态 0 正常 -1 禁用
	CreateTime utils.JsonTime `json:"create_time"`                   // 创建时间
	UpdateTime utils.JsonTime `json:"update_time"`                   // 更新时间
}

// AccountListItem account 账号列表
type AccountListItem struct {
	*AccountEntity
	Roles  []*RoleEntity      `json:"roles"`  // 角色列表
	Action *AccountListAction `json:"action"` // 数据操作
}

// AccountListAction 账号列表操作权限
type AccountListAction struct {
	IsEdit         int `json:"is_edit,omitempty"`          // 修改
	IsDetail       int `json:"is_detail,omitempty"`        // 详情
	IsUpdateStatus int `json:"is_update_status,omitempty"` // 更新状态
}

// AccountKeywords 账号搜索词
type AccountKeywords struct {
	Status      string `json:"status"`       // 状态
	AccountName string `json:"account_name"` // 账号名称
	GivenName   string `json:"given_name"`   // 昵称
}

// AccountRoleEntity 账号角色表结构
type AccountRoleEntity struct {
	AccountRoleId int64          `json:"account_role_id" gorm:"primary_key"` // 账号角色ID
	AccountId     int64          `json:"account_id"`                         // 账号ID
	RoleId        int64          `json:"role_id"`                            // 角色ID
	CreateTime    utils.JsonTime `json:"create_time"`                        // 创建时间
}

// AccountEditResp 账号编辑返回结果
type AccountEditResp struct {
	AccountInfo  *AccountEntity `json:"account_info"`  // 账号信息
	AccountRoles []*RoleEntity  `json:"account_roles"` // 账号角色信息
	RoleList     []*RoleEntity  `json:"role_list"`     // 角色列表
}

// AccountDetailResp 账号详情返回结果
type AccountDetailResp struct {
	AccountInfo  *AccountEntity `json:"account_info"`  // 账号信息
	AccountRoles []*RoleEntity  `json:"account_roles"` // 账号角色信息
}
