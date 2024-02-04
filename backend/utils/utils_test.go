package utils

import (
	"fmt"
	"testing"
)

func TestGetNewParentIds(t *testing.T) {
	caseList := []struct {
		originParentIds string
		parentId        int64
		prefixParentIds string
		newParentId     string
	}{
		{
			originParentIds: "1,3,5,6,10,21",
			parentId:        5,
			prefixParentIds: "2,8",
			newParentId:     "2,8,5,6,10,21",
		},
		{
			originParentIds: "1,3,5,6,10,21",
			parentId:        3,
			prefixParentIds: "2,8",
			newParentId:     "2,8,3,5,6,10,21",
		},
		{
			originParentIds: "0,1,3,5,6,10,21",
			parentId:        0,
			prefixParentIds: "",
			newParentId:     "0,1,3,5,6,10,21",
		},
		{
			originParentIds: "0,1,3,5,6,10,21",
			parentId:        21,
			prefixParentIds: "0,1,3,5,6",
			newParentId:     "0,1,3,5,6,21",
		},
	}

	for _, caseItem := range caseList {
		res := GetNewParentIds(caseItem.originParentIds, caseItem.parentId, caseItem.prefixParentIds)
		if res != caseItem.newParentId {
			t.Errorf("case=%+v res=%+v err", caseItem, res)
		}
	}

}

func TestGetIdsDiff(t *testing.T) {
	caseList := []struct {
		resIds    []int64
		targetIds []int64
		addIds    []int64
		deleteIds []int64
	}{
		{
			resIds:    []int64{},
			targetIds: []int64{1, 2, 3},
			addIds:    []int64{1, 2, 3},
			deleteIds: []int64{},
		},
		{
			resIds:    []int64{1, 2, 3},
			targetIds: []int64{},
			addIds:    []int64{},
			deleteIds: []int64{1, 2, 3},
		},
		{
			resIds:    []int64{2, 5, 1, 9},
			targetIds: []int64{1, 2, 3},
			addIds:    []int64{3},
			deleteIds: []int64{5, 9},
		},
		{
			resIds:    []int64{1, 2, 3},
			targetIds: []int64{4, 5, 6},
			addIds:    []int64{4, 5, 6},
			deleteIds: []int64{1, 2, 3},
		},
	}

	for _, caseItem := range caseList {
		addIds, deleteIds := GetIdsDiffRes(caseItem.resIds, caseItem.targetIds)
		if fmt.Sprintf("%+v", addIds) != fmt.Sprintf("%+v", caseItem.addIds) {
			t.Errorf("addIds=%+v, expect=%+v", addIds, caseItem.addIds)
		}
		if fmt.Sprintf("%+v", deleteIds) != fmt.Sprintf("%+v", caseItem.deleteIds) {
			t.Errorf("deleteIds=%+v, expect=%+v", deleteIds, caseItem.deleteIds)
		}
	}

}
