package entity

import "github.com/phachon/mm-wiki/utils"

const (
	PluginStatusDisabled = 0  // 插件状态 0 禁用
	PluginStatusEnabled  = 1  // 插件状态 1 启用
	PluginStatusDelete   = -1 // 插件状态 -1 删除
)

// PluginEntity 插件数据结构
type PluginEntity struct {
	PluginId    int64          `json:"plugin_id" gorm:"primary_key"` // 插件ID
	Name        string         `json:"name"`                         // 插件名称
	Key         string         `json:"key"`                          // 插件标识
	Description string         `json:"description"`                  // 插件描述
	Version     string         `json:"version"`                      // 插件版本
	Author      string         `json:"author"`                       // 插件作者
	ConfigJSON  string         `json:"config_json"`                  // 插件配置 JSON
	Status      int            `json:"status"`                       // 状态 0 禁用 1 启用 -1 删除
	CreateTime  utils.JsonTime `json:"create_time"`                  // 创建时间
	UpdateTime  utils.JsonTime `json:"update_time"`                  // 修改时间
}

// TableName 表名
func (e *PluginEntity) TableName() string {
	return "mk_plugin"
}

// PluginListItem 插件列表结构
type PluginListItem struct {
	*PluginEntity
	Action *PluginListAction `json:"action"` // 数据操作
}

// PluginListAction 插件列表操作权限
type PluginListAction struct {
	IsEdit   int `json:"is_edit,omitempty"`   // 插件修改
	IsDelete int `json:"is_delete,omitempty"` // 插件删除
}

// PluginKeywords 插件搜索关键字
type PluginKeywords struct {
	Name string // 搜索名称
}
