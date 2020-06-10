package controllers

import (
	"strings"
	"time"

	"github.com/phachon/mm-wiki/app/models"
	"github.com/phachon/mm-wiki/app/utils"

	"github.com/astaxie/beego/validation"
)

type ContactController struct {
	BaseController
}

func (this *ContactController) Add() {
	this.viewLayout("contact/form", "contact")
}

func (this *ContactController) Save() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/contact/list")
	}
	name := strings.Trim(this.GetString("name", ""), "")
	mobile := strings.Trim(this.GetString("mobile", ""), "")
	position := strings.Trim(this.GetString("position", ""), "")
	email := strings.Trim(this.GetString("email", ""), "")

	v := validation.Validation{}
	if name == "" {
		this.jsonError("联系人姓名不能为空！")
	}
	if email == "" {
		this.jsonError("邮箱不能为空！")
	}
	if !v.Email(email, "email").Ok {
		this.jsonError("邮箱格式不正确！")
	}

	contact := map[string]interface{}{
		"name":     name,
		"mobile":   mobile,
		"position": position,
		"email":    email,
	}

	contactId, err := models.ContactModel.Insert(contact)
	if err != nil {
		this.ErrorLog("添加联系人失败：" + err.Error())
		this.jsonError("添加联系人失败")
	}
	this.InfoLog("添加联系人 " + utils.Convert.IntToString(contactId, 10) + " 成功")
	this.jsonSuccess("添加联系人成功", nil, "/system/contact/list")
}

func (this *ContactController) List() {

	var err error
	var contacts []map[string]string
	contacts, err = models.ContactModel.GetAllContact()

	if err != nil {
		this.ErrorLog("获取联系人列表出错: " + err.Error())
		this.ViewError("获取联系人列表出错", "/system/main/index")
	}

	this.Data["contacts"] = contacts
	this.viewLayout("contact/list", "contact")
}

func (this *ContactController) Edit() {

	contactId := this.GetString("contact_id", "")
	if contactId == "" {
		this.ViewError("联系人不存在", "/system/contact/list")
	}

	contact, err := models.ContactModel.GetContactByContactId(contactId)
	if err != nil {
		this.ViewError("联系人不存在", "/system/contact/list")
	}

	this.Data["contact"] = contact
	this.viewLayout("contact/form", "contact")
}

func (this *ContactController) Modify() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/contact/list")
	}
	contactId := strings.Trim(this.GetString("contact_id", ""), "")
	name := strings.Trim(this.GetString("name", ""), "")
	mobile := strings.Trim(this.GetString("mobile", ""), "")
	position := strings.Trim(this.GetString("position", ""), "")
	email := strings.Trim(this.GetString("email", ""), "")

	v := validation.Validation{}
	if contactId == "" {
		this.jsonError("参数错误！")
	}
	if name == "" {
		this.jsonError("联系人姓名不能为空！")
	}
	if position == "" {
		this.jsonError("职位不能为空！")
	}
	if mobile == "" {
		this.jsonError("联系电话不能为空！")
	}
	if !v.Phone(mobile, "mobile").Ok {
		this.jsonError("联系电话格式不正确！")
	}
	if email == "" {
		this.jsonError("邮箱不能为空！")
	}
	if !v.Email(email, "email").Ok {
		this.jsonError("邮箱格式不正确！")
	}

	contact := map[string]interface{}{
		"name":        name,
		"mobile":      mobile,
		"position":    position,
		"email":       email,
		"update_time": time.Now().Unix(),
	}
	_, err := models.ContactModel.UpdateByContactId(contact, contactId)
	if err != nil {
		this.ErrorLog("修改联系人 " + contactId + " 失败：" + err.Error())
		this.jsonError("修改联系人失败")
	}
	this.InfoLog("修改联系人 " + contactId + " 成功")
	this.jsonSuccess("修改联系人成功", nil, "/system/contact/list")
}

func (this *ContactController) Delete() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/contact/list")
	}
	contactId := this.GetString("contact_id", "")
	if contactId == "" {
		this.jsonError("没有选择联系人！")
	}

	contact, err := models.ContactModel.GetContactByContactId(contactId)
	if err != nil {
		this.ErrorLog("删除联系人 " + contactId + " 失败: " + err.Error())
		this.jsonError("删除联系人失败")
	}
	if len(contact) == 0 {
		this.jsonError("联系人不存在")
	}

	_, err = models.ContactModel.DeleteByContactId(contactId)
	if err != nil {
		this.ErrorLog("删除联系人 " + contactId + " 失败: " + err.Error())
		this.jsonError("删除联系人失败")
	}

	this.InfoLog("删除联系人 " + contactId + " 成功")
	this.jsonSuccess("删除联系人成功", nil, "/system/contact/list")
}

// 从用户导入联系人
func (this *ContactController) Import() {

	keywords := map[string]string{}
	page, _ := this.GetInt("page", 1)
	username := strings.TrimSpace(this.GetString("username", ""))
	number, _ := this.GetRangeInt("number", 20, 10, 100)
	if username != "" {
		keywords["username"] = username
	}
	limit := (page - 1) * number

	var err error
	var count int64
	var users []map[string]string
	if len(keywords) != 0 {
		count, err = models.UserModel.CountUsersByKeywords(keywords)
		users, err = models.UserModel.GetUsersByKeywordsAndLimit(keywords, limit, number)
	} else {
		count, err = models.UserModel.CountUsers()
		users, err = models.UserModel.GetUsersByLimit(limit, number)
	}
	if err != nil {
		this.ErrorLog("获取用户列表失败: " + err.Error())
		this.ViewError("获取用户列表失败", "/system/main/index")
	}

	this.Data["users"] = users
	this.Data["username"] = username
	this.SetPaginator(number, count)
	this.viewLayout("contact/import", "contact")
}