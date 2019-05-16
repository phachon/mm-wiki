// Copyright 2013 wetalk authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package utils

import (
	"math"
	"net/http"
	"net/url"
	"strconv"
)

// 默认的每页条数的选择范围
var defaultPerPageNumsSelect = []int{10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 100}

type Paginator struct {
	Request             *http.Request
	PerPageNums         int
	PerPageNumsSelect   []int
	MaxPages            int
	nums                int64
	pageRange           []int
	pageNums            int
	page                int
	pageParamName       string
	perPageNumParamName string
}

func (p *Paginator) PageNums() int {
	if p.pageNums != 0 {
		return p.pageNums
	}
	pageNums := math.Ceil(float64(p.nums) / float64(p.PerPageNums))
	if p.MaxPages > 0 {
		pageNums = math.Min(pageNums, float64(p.MaxPages))
	}
	p.pageNums = int(pageNums)
	return p.pageNums
}

func (p *Paginator) Nums() int64 {
	return p.nums
}

func (p *Paginator) SetNums(nums interface{}) {
	p.nums, _ = NewConvert().ToInt64(nums)
}

func (p *Paginator) SetPrePageNumsSelect(selectNums []int) {
	p.PerPageNumsSelect = selectNums
}

func (p *Paginator) SetPerPageNums(perPageNums int) {

	if perPageNums < p.PerPageNumsSelect[0] {
		perPageNums = p.PerPageNumsSelect[0]
	}
	if perPageNums > p.PerPageNumsSelect[len(p.PerPageNumsSelect)-1] {
		perPageNums = p.PerPageNumsSelect[len(p.PerPageNumsSelect)-1]
	}
	p.PerPageNums = perPageNums
}

func (p *Paginator) Page() int {
	if p.page != 0 {
		return p.page
	}
	if p.Request.Form == nil {
		p.Request.ParseForm()
	}
	p.page, _ = strconv.Atoi(p.Request.Form.Get(p.pageParamName))
	if p.page > p.PageNums() {
		p.page = p.PageNums()
	}
	if p.page <= 0 {
		p.page = 1
	}
	return p.page
}

func (p *Paginator) Pages() []int {
	if p.pageRange == nil && p.nums > 0 {
		var pages []int
		pageNums := p.PageNums()
		page := p.Page()
		switch {
		case page >= pageNums-4 && pageNums > 9:
			start := pageNums - 9 + 1
			pages = make([]int, 9)
			for i, _ := range pages {
				pages[i] = start + i
			}
		case page >= 5 && pageNums > 9:
			start := page - 5 + 1
			pages = make([]int, int(math.Min(9, float64(page+4+1))))
			for i, _ := range pages {
				pages[i] = start + i
			}
		default:
			pages = make([]int, int(math.Min(9, float64(pageNums))))
			for i, _ := range pages {
				pages[i] = i + 1
			}
		}
		p.pageRange = pages
	}
	return p.pageRange
}

func (p *Paginator) PageLink(page int) string {
	link, _ := url.ParseRequestURI(p.Request.RequestURI)
	values := link.Query()
	if page == 1 {
		values.Del(p.pageParamName)
	} else {
		values.Set(p.pageParamName, strconv.Itoa(page))
	}

	if p.PerPageNums < p.PerPageNumsSelect[0] {
		p.PerPageNums = p.PerPageNumsSelect[0]
	}
	if p.PerPageNums > p.PerPageNumsSelect[len(p.PerPageNumsSelect)-1] {
		p.PerPageNums = p.PerPageNumsSelect[len(p.PerPageNumsSelect)-1]
	}
	values.Set(p.perPageNumParamName, strconv.Itoa(p.PerPageNums))
	link.RawQuery = values.Encode()
	return link.String()
}

func (p *Paginator) PrePageNumLink(perPageNum int) string {
	link, _ := url.ParseRequestURI(p.Request.RequestURI)
	values := link.Query()

	if perPageNum < p.PerPageNumsSelect[0] {
		perPageNum = p.PerPageNumsSelect[0]
	}
	if perPageNum > p.PerPageNumsSelect[len(p.PerPageNumsSelect)-1] {
		perPageNum = p.PerPageNumsSelect[len(p.PerPageNumsSelect)-1]
	}
	values.Set(p.perPageNumParamName, strconv.Itoa(perPageNum))
	link.RawQuery = values.Encode()
	return link.String()
}

func (p *Paginator) PageLinkPrev() (link string) {
	if p.HasPrev() {
		link = p.PageLink(p.Page() - 1)
	}
	return
}

func (p *Paginator) PageLinkNext() (link string) {
	if p.HasNext() {
		link = p.PageLink(p.Page() + 1)
	}
	return
}

func (p *Paginator) PageLinkFirst() (link string) {
	return p.PageLink(1)
}

func (p *Paginator) PageLinkLast() (link string) {
	return p.PageLink(p.PageNums())
}

func (p *Paginator) HasPrev() bool {
	return p.Page() > 1
}

func (p *Paginator) HasNext() bool {
	return p.Page() < p.PageNums()
}

func (p *Paginator) IsActive(page int) bool {
	return p.Page() == page
}

func (p *Paginator) Offset() int {
	return (p.Page() - 1) * p.PerPageNums
}

func (p *Paginator) HasPages() bool {
	return p.PageNums() > 1
}

func NewPaginator(req *http.Request, per int, nums interface{}) *Paginator {
	p := Paginator{}
	p.Request = req
	// 翻页参数名，默认为 page
	p.pageParamName = "page"
	// 每页条数参数名，默认为 "number"
	p.perPageNumParamName = "number"
	if per <= 0 {
		per = 10
	}
	p.SetPrePageNumsSelect(defaultPerPageNumsSelect)
	p.SetPerPageNums(per)
	p.SetNums(nums)
	return &p
}
