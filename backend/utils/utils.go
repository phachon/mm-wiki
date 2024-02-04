// Package utils 公共工具包
package utils

import (
	"fmt"
	"strings"
)

// GetConfPaths 获取配置文件 path
func GetConfPaths(rootPath string) []string {
	return []string{
		rootPath,
		rootPath + "/conf/",
		rootPath + "/../conf/",
		rootPath + "/../../conf/",
	}
}

// GetNewParentIds 获取新的 parentIds
func GetNewParentIds(originParentIds string, parentId int64, prefixParentIds string) string {
	parentList := strings.Split(originParentIds, ",")
	var pos = -1
	for index, tmpParentId := range parentList {
		if tmpParentId != fmt.Sprintf("%d", parentId) {
			continue
		}
		pos = index
	}
	if pos == -1 {
		return originParentIds
	}
	newParentIds := prefixParentIds
	var suffixParentIds = parentList[pos:]
	if prefixParentIds == "" {
		return strings.Join(suffixParentIds, ",")
	}
	if len(suffixParentIds) > 0 {
		newParentIds = prefixParentIds + "," + strings.Join(suffixParentIds, ",")
	}
	return newParentIds
}

// GetIdsDiffRes 获取 id 的 diff 结果
func GetIdsDiffRes(resIds []int64, targetIds []int64) (addIds []int64, deleteIds []int64) {
	if len(resIds) == 0 {
		return targetIds, []int64{}
	}
	if len(targetIds) == 0 {
		return []int64{}, resIds
	}
	var resIdsMap = make(map[int64]uint8)
	var targetIdsMap = make(map[int64]uint8)

	for _, resId := range resIds {
		resIdsMap[resId] = 1
	}
	for _, targetId := range targetIds {
		targetIdsMap[targetId] = 1
		// 目标ID是否本身就包含已有的ID
		if resIdsMap[targetId] == 1 {
			continue
		}
		addIds = append(addIds, targetId) // 不存在的ID就是为新增的
	}
	for _, resId := range resIds {
		// 原ID是否在目标ID里
		if targetIdsMap[resId] == 1 {
			continue
		}
		deleteIds = append(deleteIds, resId) // 不存在的ID就是需要删除的
	}
	return addIds, deleteIds
}

// FilterStringIds 过滤掉 id
func FilterStringIds(resIds []string, filterId string) (resp []string) {
	if len(resIds) == 0 {
		return resIds
	}
	for _, resId := range resIds {
		if resId == filterId {
			continue
		}
		resp = append(resp, resId)
	}
	return
}
