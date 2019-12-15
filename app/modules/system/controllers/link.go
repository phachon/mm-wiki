package controllers

import (
	"strings"

	"github.com/phachon/mm-wiki/app/models"
	"github.com/phachon/mm-wiki/app/utils"

	valid "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type LinkController struct {
	BaseController
}

func (this *LinkController) Add() {
	this.viewLayout("link/form", "link")
}

func (this *LinkController) Save() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/link/list")
	}
	name := strings.TrimSpace(this.GetString("name", ""))
	url := strings.TrimSpace(this.GetString("url", ""))
	sequence := strings.TrimSpace(this.GetString("sequence", "0"))
	if name == "" {
		this.jsonError("链接名称不能为空！")
	}
	if url == "" {
		this.jsonError("链接地址不能为空！")
	}
	if valid.Validate(url, is.URL) != nil {
		this.jsonError("链接地址格式不正确！")
	}
	ok, err := models.LinkModel.HasLinkName(name)
	if err != nil {
		this.ErrorLog("添加链接失败：" + err.Error())
		this.jsonError("添加链接失败！")
	}
	if ok {
		this.jsonError("链接名已经存在！")
	}

	linkId, err := models.LinkModel.Insert(map[string]interface{}{
		"name":     name,
		"url":      url,
		"sequence": sequence,
	})

	if err != nil {
		this.ErrorLog("添加链接失败：" + err.Error())
		this.jsonError("添加链接失败")
	}
	this.InfoLog("添加链接 " + utils.Convert.IntToString(linkId, 10) + " 成功")
	this.jsonSuccess("添加链接成功", nil, "/system/link/list")
}

func (this *LinkController) List() {

	page, _ := this.GetInt("page", 1)
	keyword := strings.TrimSpace(this.GetString("keyword", ""))
	number, _ := this.GetRangeInt("number", 20, 10, 100)
	limit := (page - 1) * number

	var err error
	var count int64
	var links []map[string]string
	if keyword != "" {
		count, err = models.LinkModel.CountLinksByKeyword(keyword)
		links, err = models.LinkModel.GetLinksByKeywordAndLimit(keyword, limit, number)
	} else {
		count, err = models.LinkModel.CountLinks()
		links, err = models.LinkModel.GetLinksByLimit(limit, number)
	}
	if err != nil {
		this.ErrorLog("获取链接列表失败: " + err.Error())
		this.ViewError("获取链接列表失败", "/system/main/index")
	}

	this.Data["links"] = links
	this.Data["keyword"] = keyword
	this.SetPaginator(number, count)
	this.viewLayout("link/list", "link")
}

func (this *LinkController) Edit() {

	linkId := this.GetString("link_id", "")
	if linkId == "" {
		this.ViewError("链接不存在", "/system/link/list")
	}

	link, err := models.LinkModel.GetLinkByLinkId(linkId)
	if err != nil {
		this.ViewError("链接不存在", "/system/link/list")
	}

	this.Data["link"] = link
	this.viewLayout("link/form", "link")
}

func (this *LinkController) Modify() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/link/list")
	}
	linkId := this.GetString("link_id", "")
	name := strings.TrimSpace(this.GetString("name", ""))
	url := strings.TrimSpace(this.GetString("url", ""))
	sequence := strings.TrimSpace(this.GetString("sequence", ""))

	if linkId == "" {
		this.jsonError("链接不存在！")
	}
	if name == "" {
		this.jsonError("链接名称不能为空！")
	}
	if url == "" {
		this.jsonError("链接地址不能为空！")
	}
	if valid.Validate(url, is.URL) != nil {
		this.jsonError("链接地址格式不正确！")
	}

	link, err := models.LinkModel.GetLinkByLinkId(linkId)
	if err != nil {
		this.ErrorLog("修改链接 " + linkId + " 失败: " + err.Error())
		this.jsonError("修改链接失败！")
	}
	if len(link) == 0 {
		this.jsonError("链接不存在！")
	}

	ok, _ := models.LinkModel.HasSameName(linkId, name)
	if ok {
		this.jsonError("链接名已经存在！")
	}
	_, err = models.LinkModel.Update(linkId, map[string]interface{}{
		"name":     name,
		"url":      url,
		"sequence": sequence,
	})

	if err != nil {
		this.ErrorLog("修改链接 " + linkId + " 失败：" + err.Error())
		this.jsonError("修改链接失败")
	}
	this.InfoLog("修改链接 " + linkId + " 成功")
	this.jsonSuccess("修改链接成功", nil, "/system/link/list")
}

func (this *LinkController) Delete() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/link/list")
	}
	linkId := this.GetString("link_id", "")
	if linkId == "" {
		this.jsonError("没有选择链接！")
	}

	link, err := models.LinkModel.GetLinkByLinkId(linkId)
	if err != nil {
		this.ErrorLog("删除链接 " + linkId + " 失败: " + err.Error())
		this.jsonError("删除链接失败")
	}
	if len(link) == 0 {
		this.jsonError("链接不存在")
	}

	err = models.LinkModel.Delete(linkId)
	if err != nil {
		this.ErrorLog("删除链接 " + linkId + " 失败: " + err.Error())
		this.jsonError("删除链接失败")
	}

	this.InfoLog("删除链接 " + linkId + " 成功")
	this.jsonSuccess("删除链接成功", nil, "/system/link/list")
}
