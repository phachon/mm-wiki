package entity

import "github.com/phachon/mm-wiki/utils"

const (
	EmailStatusDefault = 0  // 邮箱状态 0 正常
	EmailStatusDelete  = -1 // 邮箱状态 -1 删除
	EmailUsedTrue      = 1  // 正在使用
	EmailUsedFalse     = 0  // 未使用
)

// EmailEntity 邮箱服务器数据结构
type EmailEntity struct {
	EmailId           int64          `json:"email_id" gorm:"primary_key"` // 邮箱ID
	Name              string         `json:"name"`                        // 邮箱服务器名称
	SenderAddress     string         `json:"sender_address"`              // 发件人邮件地址
	SenderName        string         `json:"sender_name"`                 // 发件人显示名
	SenderTitlePrefix string         `json:"sender_title_prefix"`         // 发送邮件标题前缀
	Host              string         `json:"host"`                        // 服务器主机名
	Port              int            `json:"port"`                        // 服务器端口
	Username          string         `json:"username"`                    // 用户名
	Password          string         `json:"password"`                    // 密码
	IsSSL             int            `json:"is_ssl"`                      // 是否使用ssl 0 否 1 是
	IsUsed            int            `json:"is_used"`                     // 是否被使用 0 否 1 是
	Status            int            `json:"status"`                      // 状态 0 正常 -1 删除
	CreateTime        utils.JsonTime `json:"create_time"`                 // 创建时间
	UpdateTime        utils.JsonTime `json:"update_time"`                 // 修改时间
}

// EmailListItem 邮箱列表结构
type EmailListItem struct {
	*EmailEntity
	Action *EmailListAction `json:"action"` // 数据操作
}

// EmailListAction 邮箱列表操作权限
type EmailListAction struct {
	IsEdit   int `json:"is_edit,omitempty"`   // 邮箱修改
	IsDelete int `json:"is_delete,omitempty"` // 邮箱删除
}

// EmailKeywords 邮箱搜索关键字
type EmailKeywords struct {
	Name string // 搜索名称
}
