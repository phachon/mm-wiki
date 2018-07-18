package controllers

import (
	"strings"
	"mm-wiki/app/models"
	"mm-wiki/app/utils"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type EmailController struct {
	BaseController
}

func (this *EmailController) List() {

	keyword := strings.TrimSpace(this.GetString("keyword", ""))
	var err error
	var emails []map[string]string
	if keyword != "" {
		emails, err = models.EmailModel.GetEmailsByLikeName(keyword)
	} else {
		emails, err = models.EmailModel.GetEmails()
	}
	if err != nil {
		this.ErrorLog("获取邮件服务器列表失败: "+err.Error())
		this.ViewError("获取邮件服务器列表失败", "/system/main/index")
	}

	this.Data["emails"] = emails
	this.Data["keyword"] = keyword
	this.viewLayout("email/list", "default")
}

func (this *EmailController) Add() {
	this.viewLayout("email/form", "default")
}

func (this *EmailController) Save() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/email/list")
	}
	name := strings.TrimSpace(this.GetString("name", ""))
	senderAddress := strings.TrimSpace(this.GetString("sender_address", ""))
	senderName := strings.TrimSpace(this.GetString("sender_name", ""))
	senderTitlePrefix := strings.TrimSpace(this.GetString("sender_title_prefix", ""))
	host := strings.TrimSpace(this.GetString("host", ""))
	port := strings.TrimSpace(this.GetString("port", "25"))
	username := strings.TrimSpace(this.GetString("username", ""))
	password := strings.TrimSpace(this.GetString("password", ""))

	if name == "" {
		this.jsonError("邮件服务器名称不能为空！")
	}
	if host == "" {
		this.jsonError("邮件服务器主机不能为空！")
	}
	if validation.Validate(host, is.Host) != nil {
		this.jsonError("邮件服务器主机格式不正确！")
	}
	if port == "" {
		this.jsonError("邮件服务器端口不能为空！")
	}
	if validation.Validate(port, is.Port) != nil {
		this.jsonError("邮件服务器端口格式不正确！")
	}
	if senderAddress == "" {
		this.jsonError("发件人邮箱不能为空！")
	}
	if username == "" {
		this.jsonError("发件用户名不能为空！")
	}
	if password == "" {
		this.jsonError("发件人密码不能为空！")
	}

	ok, err := models.EmailModel.HasEmailName(name)
	if err != nil {
		this.ErrorLog("添加邮件服务器失败："+err.Error())
		this.jsonError("添加邮件服务器失败！")
	}
	if ok {
		this.jsonError("邮件服务器名称已经存在！")
	}

	emailId, err := models.EmailModel.Insert(map[string]interface{}{
		"name": name,
		"sender_address": senderAddress,
		"sender_name": senderName,
		"sender_title_prefix": senderTitlePrefix,
		"host": host,
		"port": port,
		"username": username,
		"password": password,
	})

	if err != nil {
		this.ErrorLog("添加邮件服务器失败：" + err.Error())
		this.jsonError("添加邮件服务器失败")
	}
	this.InfoLog("添加邮件服务器 "+utils.Convert.IntToString(emailId, 10)+" 成功")
	this.jsonSuccess("添加邮件服务器成功", nil, "/system/email/list")
}

func (this *EmailController) Edit() {

	emailId := this.GetString("email_id", "")
	if emailId == "" {
		this.ViewError("邮件服务器不存在", "/system/email/list")
	}

	email, err := models.EmailModel.GetEmailByEmailId(emailId)
	if err != nil {
		this.ViewError("邮件服务器不存在", "/system/email/list")
	}

	this.Data["email"] = email
	this.viewLayout("email/form", "default")
}

func (this *EmailController) Modify() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/email/list")
	}
	emailId := this.GetString("email_id", "")
	name := strings.TrimSpace(this.GetString("name", ""))
	senderAddress := strings.TrimSpace(this.GetString("sender_address", ""))
	senderName := strings.TrimSpace(this.GetString("sender_name", ""))
	senderTitlePrefix := strings.TrimSpace(this.GetString("sender_title_prefix", ""))
	host := strings.TrimSpace(this.GetString("host", ""))
	port := strings.TrimSpace(this.GetString("port", ""))
	username := strings.TrimSpace(this.GetString("username", ""))
	password := strings.TrimSpace(this.GetString("password", ""))

	if emailId == "" {
		this.jsonError("邮件服务器不存在！")
	}
	if name == "" {
		this.jsonError("邮件服务器名称不能为空！")
	}
	if host == "" {
		this.jsonError("邮件服务器主机不能为空！")
	}
	if validation.Validate(host, is.Host) != nil {
		this.jsonError("邮件服务器主机格式不正确！")
	}
	if port == "" {
		this.jsonError("邮件服务器端口不能为空！")
	}
	if validation.Validate(port, is.Port) != nil {
		this.jsonError("邮件服务器端口格式不正确！")
	}
	if senderAddress == "" {
		this.jsonError("发件人邮箱不能为空！")
	}
	if username == "" {
		this.jsonError("发件用户名不能为空！")
	}
	if password == "" {
		this.jsonError("发件人密码不能为空！")
	}

	email, err := models.EmailModel.GetEmailByEmailId(emailId)
	if err != nil {
		this.ErrorLog("修改邮件服务器 "+emailId+" 失败: "+err.Error())
		this.jsonError("修改邮件服务器失败！")
	}
	if len(email) == 0 {
		this.jsonError("邮件服务器不存在！")
	}

	ok , _ := models.EmailModel.HasSameName(emailId, name)
	if ok {
		this.jsonError("邮件服务器名称已经存在！")
	}
	_, err = models.EmailModel.Update(emailId, map[string]interface{}{
		"name": name,
		"sender_address": senderAddress,
		"sender_name": senderName,
		"sender_title_prefix": senderTitlePrefix,
		"host": host,
		"port": port,
		"username": username,
		"password": password,
	})

	if err != nil {
		this.ErrorLog("修改邮件服务器 "+emailId+" 失败：" + err.Error())
		this.jsonError("修改邮件服务器"+emailId+"失败")
	}
	this.InfoLog("修改邮件服务器 "+emailId+" 成功")
	this.jsonSuccess("修改邮件服务器成功", nil, "/system/email/list")
}

func (this *EmailController) Used() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/email/list")
	}
	emailId := this.GetString("email_id", "")
	if emailId == "" {
		this.jsonError("没有选择邮件服务器！")
	}

	email, err := models.EmailModel.GetEmailByEmailId(emailId)
	if err != nil {
		this.ErrorLog("邮件服务器 "+emailId+" 启用失败: "+err.Error())
		this.jsonError("邮件服务器启用失败")
	}
	if len(email) == 0 {
		this.jsonError("邮件服务器不存在")
	}
	_, err = models.EmailModel.SetEmailUsed(emailId)
	if err != nil {
		this.ErrorLog("邮件服务器 "+emailId+" 启用失败: "+err.Error())
		this.jsonError("邮件服务器启用失败")
	}

	this.InfoLog("启用邮件服务器 "+emailId+" 成功")
	this.jsonSuccess("启用邮件服务器成功", nil, "/system/email/list")
}

func (this *EmailController) Delete() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/email/list")
	}
	emailId := this.GetString("email_id", "")
	if emailId == "" {
		this.jsonError("没有选择邮件服务器！")
	}

	email, err := models.EmailModel.GetEmailByEmailId(emailId)
	if err != nil {
		this.ErrorLog("删除邮件服务器 "+emailId+" 失败: "+err.Error())
		this.jsonError("删除邮件服务器失败")
	}
	if len(email) == 0 {
		this.jsonError("邮件服务器不存在")
	}
	err = models.EmailModel.Delete(emailId)
	if err != nil {
		this.ErrorLog("删除邮件服务器 "+emailId+" 失败: "+err.Error())
		this.jsonError("删除邮件服务器失败")
	}

	this.InfoLog("删除邮件服务器 "+emailId+" 成功")
	this.jsonSuccess("删除邮件服务器成功", nil, "/system/email/list")
}