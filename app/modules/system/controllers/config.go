package controllers

import (
	"mm-wiki/app/models"
)

type ConfigController struct {
	BaseController
}

func (this *ConfigController) Global() {

	configs, err := models.ConfigModel.GetConfigs()
	if err != nil {
		this.ErrorLog("获取全局配置失败: "+err.Error())
		this.ViewError("邮件服务器不存在", "/system/main/index")
	}

	this.Data["configs"] = configs
	this.viewLayout("config/form", "default")
}

func (this *ConfigController) Modify() {

	//if !this.IsPost() {
	//	this.ViewError("请求方式有误！", "/system/email/list")
	//}
	//emailId := this.GetString("email_id", "")
	//name := strings.TrimSpace(this.GetString("name", ""))
	//senderAddress := strings.TrimSpace(this.GetString("sender_address", ""))
	//senderName := strings.TrimSpace(this.GetString("sender_name", ""))
	//senderTitlePrefix := strings.TrimSpace(this.GetString("sender_title_prefix", ""))
	//host := strings.TrimSpace(this.GetString("host", ""))
	//username := strings.TrimSpace(this.GetString("username", ""))
	//password := strings.TrimSpace(this.GetString("password", ""))
	//
	//if emailId == "" {
	//	this.jsonError("邮件服务器不存在！")
	//}
	//if name == "" {
	//	this.jsonError("邮件服务器名称不能为空！")
	//}
	//if host == "" {
	//	this.jsonError("邮件服务器主机不能为空！")
	//}
	//if senderAddress == "" {
	//	this.jsonError("发件人邮箱不能为空！")
	//}
	//if username == "" {
	//	this.jsonError("发件用户名不能为空！")
	//}
	//if password == "" {
	//	this.jsonError("发件人密码不能为空！")
	//}
	//
	//email, err := models.ConfigModel.GetConfigByConfigId(emailId)
	//if err != nil {
	//	this.ErrorLog("修改邮件服务器 "+emailId+" 失败: "+err.Error())
	//	this.jsonError("修改邮件服务器失败！")
	//}
	//if len(email) == 0 {
	//	this.jsonError("邮件服务器不存在！")
	//}
	//
	//ok , _ := models.ConfigModel.HasSameName(emailId, name)
	//if ok {
	//	this.jsonError("邮件服务器名称已经存在！")
	//}
	//_, err = models.ConfigModel.Update(emailId, map[string]interface{}{
	//	"name": name,
	//	"sender_address": senderAddress,
	//	"sender_name": senderName,
	//	"sender_title_prefix": senderTitlePrefix,
	//	"host": host,
	//	"username": username,
	//	"password": password,
	//})
	//
	//if err != nil {
	//	this.ErrorLog("修改邮件服务器 "+emailId+" 失败：" + err.Error())
	//	this.jsonError("修改邮件服务器"+emailId+"失败")
	//}
	//this.InfoLog("修改邮件服务器 "+emailId+" 成功")
	//this.jsonSuccess("修改邮件服务器成功", nil, "/system/email/list")
}