package entity

const (
	LoginAuthStatusDefault = 0  // 认证状态 0 正常
	LoginAuthStatusDelete  = -1 // 认证状态 -1 删除
	LoginAuthIsUsedNo      = 0  // 未使用
	LoginAuthIsUsedYes     = 1  // 使用中
)

// LoginAuthEntity 统一登录认证数据结构
type LoginAuthEntity struct {
	LoginAuthId   int64  `json:"login_auth_id" gorm:"primary_key"` // 认证表主键ID
	Name          string `json:"name"`                              // 登录认证名称
	AccountPrefix string `json:"account_prefix"`                    // 账号登录前缀
	URL           string `json:"url"`                               // 认证接口 url
	ExtData       string `json:"ext_data"`                          // 额外数据
	IsUsed        int    `json:"is_used"`                           // 是否被使用 0 未使用 1 使用
	Status        int    `json:"status"`                            // 状态 0 正常 -1 删除
	CreateTime    int64  `json:"create_time"`                       // 创建时间
	UpdateTime    int64  `json:"update_time"`                       // 更新时间
}

// LoginAuthListItem 登录认证列表结构
type LoginAuthListItem struct {
	*LoginAuthEntity
	Action *LoginAuthListAction `json:"action"` // 数据操作
}

// LoginAuthListAction 登录认证列表操作权限
type LoginAuthListAction struct {
	IsEdit   int `json:"is_edit,omitempty"`   // 修改
	IsDelete int `json:"is_delete,omitempty"` // 删除
}

// LoginAuthKeywords 登录认证搜索关键字
type LoginAuthKeywords struct {
	Name string `json:"name"` // 搜索名称
}
