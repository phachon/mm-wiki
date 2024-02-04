package entity

import "github.com/phachon/mm-wiki/utils"

// LogEntity 日志表结构定义
type LogEntity struct {
	LogId       int64          `json:"log_id" gorm:"primary_key"` // 日志ID
	Uri         string         `json:"uri"`                       // 接口uri
	Get         string         `json:"get"`                       // get参数
	Post        string         `json:"post"`                      // post参数
	Message     string         `json:"message"`                   // 信息
	Level       int            `json:"level"`                     // 日志级别
	File        string         `json:"file"`                      // 文件
	Line        int            `json:"line"`                      // 行数
	Ip          string         `json:"ip"`                        // IP
	AccountId   int64          `json:"account_id"`                // 账号ID
	AccountName string         `json:"account_name"`              // 账号名
	CreateTime  utils.JsonTime `json:"create_time"`               // 创建时间
}

// LogSearchKeywords 日志搜索词
type LogSearchKeywords struct {
	Level     int    `json:"level"`      // 日志级别
	Message   string `json:"message"`    // 日志信息
	AccountId int64  `json:"account_id"` // 账号ID
}
