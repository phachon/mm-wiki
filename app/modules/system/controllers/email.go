package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/phachon/mm-wiki/app/models"
	"github.com/phachon/mm-wiki/app/utils"

	"github.com/astaxie/beego"
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
		this.ErrorLog("获取邮件服务器列表失败: " + err.Error())
		this.ViewError("获取邮件服务器列表失败", "/system/main/index")
	}

	this.Data["emails"] = emails
	this.Data["keyword"] = keyword
	this.viewLayout("email/list", "email")
}

func (this *EmailController) Add() {
	this.viewLayout("email/form", "email")
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
	isSsl := strings.TrimSpace(this.GetString("is_ssl", "0"))

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
		this.ErrorLog("添加邮件服务器失败：" + err.Error())
		this.jsonError("添加邮件服务器失败！")
	}
	if ok {
		this.jsonError("邮件服务器名称已经存在！")
	}

	emailId, err := models.EmailModel.Insert(map[string]interface{}{
		"name":                name,
		"sender_address":      senderAddress,
		"sender_name":         senderName,
		"sender_title_prefix": senderTitlePrefix,
		"host":                host,
		"port":                port,
		"username":            username,
		"password":            password,
		"is_ssl":              isSsl,
	})

	if err != nil {
		this.ErrorLog("添加邮件服务器失败：" + err.Error())
		this.jsonError("添加邮件服务器失败")
	}
	this.InfoLog("添加邮件服务器 " + utils.Convert.IntToString(emailId, 10) + " 成功")
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
	this.viewLayout("email/form", "email")
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
	isSsl := strings.TrimSpace(this.GetString("is_ssl", "0"))

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
		this.ErrorLog("修改邮件服务器 " + emailId + " 失败: " + err.Error())
		this.jsonError("修改邮件服务器失败！")
	}
	if len(email) == 0 {
		this.jsonError("邮件服务器不存在！")
	}

	ok, _ := models.EmailModel.HasSameName(emailId, name)
	if ok {
		this.jsonError("邮件服务器名称已经存在！")
	}
	_, err = models.EmailModel.Update(emailId, map[string]interface{}{
		"name":                name,
		"sender_address":      senderAddress,
		"sender_name":         senderName,
		"sender_title_prefix": senderTitlePrefix,
		"host":                host,
		"port":                port,
		"username":            username,
		"password":            password,
		"is_ssl":              isSsl,
	})

	if err != nil {
		this.ErrorLog("修改邮件服务器 " + emailId + " 失败：" + err.Error())
		this.jsonError("修改邮件服务器失败")
	}
	this.InfoLog("修改邮件服务器 " + emailId + " 成功")
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
		this.ErrorLog("邮件服务器 " + emailId + " 启用失败: " + err.Error())
		this.jsonError("邮件服务器启用失败")
	}
	if len(email) == 0 {
		this.jsonError("邮件服务器不存在")
	}
	_, err = models.EmailModel.SetEmailUsed(emailId)
	if err != nil {
		this.ErrorLog("邮件服务器 " + emailId + " 启用失败: " + err.Error())
		this.jsonError("邮件服务器启用失败")
	}

	this.InfoLog("启用邮件服务器 " + emailId + " 成功")
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
		this.ErrorLog("删除邮件服务器 " + emailId + " 失败: " + err.Error())
		this.jsonError("删除邮件服务器失败")
	}
	if len(email) == 0 {
		this.jsonError("邮件服务器不存在")
	}
	err = models.EmailModel.Delete(emailId)
	if err != nil {
		this.ErrorLog("删除邮件服务器 " + emailId + " 失败: " + err.Error())
		this.jsonError("删除邮件服务器失败")
	}

	this.InfoLog("删除邮件服务器 " + emailId + " 成功")
	this.jsonSuccess("删除邮件服务器成功", nil, "/system/email/list")
}

func (this *EmailController) Test() {

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
	isSsl := strings.TrimSpace(this.GetString("is_ssl", "0"))
	emails := strings.TrimSpace(this.GetString("emails", ""))

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
	if emails == "" {
		this.jsonError("要发送的邮件地址不能为空！")
	}

	emailConfig := map[string]string{
		"sender_address":      senderAddress,
		"port":                port,
		"password":            password,
		"host":                host,
		"sender_name":         senderName,
		"username":            username,
		"sender_title_prefix": senderTitlePrefix,
		"is_ssl":              isSsl,
	}

	to := strings.Split(emails, ";")
	documentValue := map[string]string{
		"name":         "MM-Wiki测试邮件",
		"username":     this.User["username"],
		"update_time":  fmt.Sprintf("%d", time.Now().Unix()),
		"comment":      "",
		"document_url": "",
		"content":      "欢迎使用 <a href='https://github.com/phachon/mm-wiki'>MM-Wiki</a>，这是一封测试邮件，请勿回复!",
	}

	emailTemplate := beego.BConfig.WebConfig.ViewsPath + "/system/email/template_test.html"
	body, err := utils.Email.MakeDocumentHtmlBody(documentValue, emailTemplate)
	if err != nil {
		this.ErrorLog("发送测试邮件失败：" + err.Error())
		this.jsonError("发送测试邮件失败！")
	}
	// start send email
	err = utils.Email.Send(emailConfig, to, "测试邮件", body)
	if err != nil {
		this.ErrorLog("发送测试邮件失败：" + err.Error())
		this.jsonError("发送测试邮件失败！")
	}

	this.jsonSuccess("发送测试邮件成功", nil)
}
