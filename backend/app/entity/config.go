package entity

import "github.com/phachon/mm-wiki/utils"

const (
	ConfigKeyMainTitle       = "main_title"
	ConfigKeyMainDescription = "main_description"
	ConfigKeyAutoFollowDoc   = "auto_follow_doc_open"
	ConfigKeySendEmail       = "send_email_open"
	ConfigKeyAuthLogin       = "sso_open"
	ConfigKeySystemVersion   = "system_version"
	ConfigKeyFulltextSearch  = "fulltext_search_open"
	ConfigKeyDocSearchTimer  = "doc_search_timer"
	ConfigKeySystemName      = "system_name"
)

// ConfigEntity 全局配置数据结构
type ConfigEntity struct {
	ConfigId   int64          `json:"config_id" gorm:"primary_key"` // 配置ID
	Name       string         `json:"name"`                         // 配置名称
	Key        string         `json:"key"`                          // 配置键
	Value      string         `json:"value"`                        // 配置值
	CreateTime utils.JsonTime `json:"create_time"`                  // 创建时间
	UpdateTime utils.JsonTime `json:"update_time"`                  // 修改时间
}
