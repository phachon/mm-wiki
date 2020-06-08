package controllers

import (
	"fmt"
	"github.com/phachon/mm-wiki/app/work"
	"strings"

	"github.com/phachon/mm-wiki/app/models"
)

type ConfigController struct {
	BaseController
}

func (this *ConfigController) Global() {

	configs, err := models.ConfigModel.GetConfigs()
	if err != nil {
		this.ErrorLog("获取全局配置失败: " + err.Error())
		this.ViewError("邮件服务器不存在", "/system/main/index")
	}

	var configValue = map[string]string{}
	for _, config := range configs {
		if config["key"] == models.ConfigKeyAutoFollowdoc && config["value"] != "1" {
			config["value"] = "0"
		}
		if config["key"] == models.ConfigKeySendEmail && config["value"] != "1" {
			config["value"] = "0"
		}
		if config["key"] == models.ConfigKeyAuthLogin && config["value"] != "1" {
			config["value"] = "0"
		}
		configValue[config["key"]] = config["value"]
	}

	this.Data["configValue"] = configValue
	this.viewLayout("config/form", "config")
}

func (this *ConfigController) Modify() {

	if !this.IsPost() {
		this.ViewError("请求方式有误！", "/system/email/list")
	}
	mainTitle := this.GetString(models.ConfigKeyMainTitle, "")
	mainDescription := strings.TrimSpace(this.GetString(models.ConfigKeyMainDescription, ""))
	autoFollowDocOpen := strings.TrimSpace(this.GetString(models.ConfigKeyAutoFollowdoc, "0"))
	sendEmailOpen := strings.TrimSpace(this.GetString(models.ConfigKeySendEmail, "0"))
	ssoOpen := strings.TrimSpace(this.GetString(models.ConfigKeyAuthLogin, "0"))
	fulltextSearch := strings.TrimSpace(this.GetString(models.ConfigKeyFulltextSearch, "0"))
	docSearchTimer := strings.TrimSpace(this.GetString(models.ConfigKeyDocSearchTimer, "3600"))
	systemName := strings.TrimSpace(this.GetString(models.ConfigKeySystemName, "Markdown Mini Wiki"))

	if sendEmailOpen == "1" {
		email, err := models.EmailModel.GetUsedEmail()
		if err != nil {
			this.ErrorLog("获取可用的邮箱配置失败: " + err.Error())
			this.jsonError("配置出错！")
		}
		if len(email) == 0 {
			this.jsonError("开启邮件通知必须先启用一个邮件服务器配置！")
		}
	}

	if ssoOpen == "1" {
		auth, err := models.AuthModel.GetUsedAuth()
		if err != nil {
			this.ErrorLog("获取可用的登录认证失败: " + err.Error())
			this.jsonError("配置出错！")
		}
		if len(auth) == 0 {
			this.jsonError("开启统一登录必须先添加并启用一个登录认证！")
		}
	}
	updateValues := map[string]string{
		models.ConfigKeyMainTitle:       mainTitle,
		models.ConfigKeyMainDescription: mainDescription,
		models.ConfigKeyAutoFollowdoc:   autoFollowDocOpen,
		models.ConfigKeySendEmail:       sendEmailOpen,
		models.ConfigKeyAuthLogin:       ssoOpen,
		models.ConfigKeyFulltextSearch:  fulltextSearch,
		models.ConfigKeyDocSearchTimer:  docSearchTimer,
		models.ConfigKeySystemName:      systemName,
	}
	// 有修改再更新
	configs, err := models.ConfigModel.GetConfigs()
	if err != nil {
		this.ErrorLog("获取配置信息失败: " + err.Error())
		this.jsonError("获取配置出错！")
	}
	updateKeys := make(map[string]string)
	for _, config := range configs {
		if len(config) == 0 {
			continue
		}
		name := config["name"]
		key := config["key"]
		value := config["value"]
		updateValue, ok := updateValues[key]
		if !ok {
			continue
		}
		// 没有修改不更新
		if value == updateValue {
			continue
		}
		_, err := models.ConfigModel.UpdateByKey(key, updateValue)
		if err != nil {
			this.ErrorLog(fmt.Sprintf("修改配置 %s 失败: %s", name, err.Error()))
			this.jsonError(fmt.Sprintf("修改配置 %s 失败", name))
		}
		updateKeys[key] = updateValue
	}

	// 更新后的回调
	this.configUpdateCallback(updateKeys)
	this.InfoLog("修改全局配置成功")
	this.jsonSuccess("修改全局配置成功", nil, "/system/config/global")
}

// 配置更新通知回调
func (this *ConfigController) configUpdateCallback(updateKeyMaps map[string]string) {
	fullTextOpenUpdate := false
	updateValue, ok := updateKeyMaps[models.ConfigKeyFulltextSearch]
	if ok {
		fullTextOpenUpdate = true
	}
	docSearchTimerUpdate := false
	_, ok = updateKeyMaps[models.ConfigKeyDocSearchTimer]
	if ok {
		docSearchTimerUpdate = true
	}
	// 索引时间更新，开关没有更新，重启一下 worker
	if docSearchTimerUpdate && !fullTextOpenUpdate {
		work.DocSearchWorker.Restart()
		return
	}
	// 开关更新
	if fullTextOpenUpdate {
		if updateValue == "1" {
			work.DocSearchWorker.Start()
			return
		}
		work.DocSearchWorker.Stop()
	}
}
