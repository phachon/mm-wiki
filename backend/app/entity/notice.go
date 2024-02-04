package entity

import "github.com/phachon/mm-wiki/utils"

const (
	NoticeStatusDefault           = 0  // 角色状态 0 正常
	NoticeStatusDelete            = -1 // 角色状态 -1 删除
	NoticePublishStatusNotPublish = 0  // 未发布
	NoticePublishStatusPublishing = 1  // 已发布
)

// NoticeEntity 系统公告数据结构
type NoticeEntity struct {
	NoticeId      int64          `json:"notice_id" gorm:"primary_key"` // 公告ID
	Title         string         `json:"title"`                        // 公告标题
	Content       string         `json:"content"`                      // 公告信息
	AccountId     int64          `json:"account_id"`                   // 公告创建人
	AccountName   string         `json:"account_name"`                 // 公告账号名
	Status        int            `json:"status"`                       // 公告状态 0 正常 -1 删除
	PublishStatus int            `json:"publish_status"`               // 发布状态 0 未发布 1 已发布
	CreateTime    utils.JsonTime `json:"create_time"`                  // 创建时间
	UpdateTime    utils.JsonTime `json:"update_time"`                  // 修改时间
	StartTime     utils.JsonTime `json:"start_time"`                   // 开始时间
	EndTime       utils.JsonTime `json:"end_time"`                     // 结束时间
}

// NoticeListItem 系统公告列表结构
type NoticeListItem struct {
	*NoticeEntity
	Action *NoticeListAction `json:"action"` // 数据操作
}

// RoleListAction 角色列表操作权限
type NoticeListAction struct {
	IsEdit   int `json:"is_edit,omitempty"`   // 公告修改
	IsDelete int `json:"is_delete,omitempty"` // 公告删除
}

// NoticeKeywords 公告搜索关键字
type NoticeKeywords struct {
	Content string // 搜索内容
}
