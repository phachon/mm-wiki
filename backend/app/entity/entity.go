// Package entity 实体定义
package entity

import (
	"math"
)

// Response 返回结果定义
type Response struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// PageInfo 列表分页结构
type PageInfo struct {
	TotalNum  int64 `json:"total_num"`  // 数据总条数
	TotalPage int64 `json:"total_page"` // 总页数
	PageSize  int64 `json:"page_size"`  // 每一页条数
	PageNum   int64 `json:"page_num"`   // 当前页码
	HasNext   int   `json:"has_next"`   // 是否下一页有数据 1默认是 0 否
}

// GetPageInfo 获取分页信息
func GetPageInfo(totalNum int64, pageSize int, pageNum int) *PageInfo {
	pageInfo := &PageInfo{
		TotalNum:  totalNum,
		TotalPage: int64(math.Ceil(float64(totalNum) / float64(pageSize))),
		PageSize:  int64(pageSize),
		PageNum:   int64(pageNum),
		HasNext:   1,
	}
	// 判断是否有下一页
	if int64(pageSize) >= totalNum || pageInfo.PageNum >= pageInfo.TotalPage {
		pageInfo.HasNext = 0
	}
	return pageInfo
}
