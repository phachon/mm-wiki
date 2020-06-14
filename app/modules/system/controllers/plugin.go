package controllers

import (
	"encoding/json"
	"github.com/phachon/mm-wiki/app/models"
	"strings"
)

type PluginController struct {
	BaseController
}

func (this *PluginController) List() {

	page, _ := this.GetInt("page", 1)
	keyword := strings.TrimSpace(this.GetString("keyword", ""))
	number, _ := this.GetRangeInt("number", 20, 10, 100)
	limit := (page - 1) * number

	var err error
	var count int64
	var plugins []map[string]string
	if keyword != "" {
		count, err = models.PluginModel.CountPluginsByKeyword(keyword)
		plugins, err = models.PluginModel.GetPluginsByKeywordAndLimit(keyword, limit, number)
	} else {
		count, err = models.PluginModel.CountPlugins()
		plugins, err = models.PluginModel.GetPluginsByLimit(limit, number)
	}
	if err != nil {
		this.ErrorLog("获取插件列表失败: " + err.Error())
		this.ViewError("获取插件列表失败", "/system/main/index")
	}

	this.Data["plugins"] = plugins
	this.Data["keyword"] = keyword
	this.SetPaginator(number, count)
	this.viewLayout("plugin/list", "plugin")
}

func (this *PluginController) Config() {

	pluginId := this.GetString("plugin_id", "")
	if pluginId == "" {
		this.ViewError("参数错误", "/system/plugin/list")
	}

	plugin, err := models.PluginModel.GetPluginByPluginId(pluginId)
	if err != nil {
		this.ErrorLog("查找插件错误：" + err.Error())
		this.ViewError("查找插件错误", "/system/plugin/list")
	}
	if len(plugin) == 0 {
		this.ViewError("插件不存在", "/system/plugin/list")
	}
	pluginKey, ok := plugin["plugin_key"]
	if !ok || pluginKey == "" {
		this.ViewError("插件数据不完整 Key 不存在", "/system/plugin/list")
	}
	configValue := make(map[string]string)
	configValueStr, ok := plugin["conf_value"]
	if ok && configValueStr != "" {
		json.Unmarshal([]byte(configValueStr), &configValue)
	}
	// 根据不同的插件，显示不同的配置页面
	this.Data["plugin_config"] = configValue
	this.Data["plugin"] = plugin
	this.viewLayout("plugin/"+pluginKey, "plugin")
}

func (this *PluginController) ConfigModify() {

	pluginId := this.GetString("plugin_id", "")
	confValue := this.GetString("conf_value", "")
	if pluginId == "" {
		this.jsonError("参数错误")
	}
	if confValue == "" {
		this.jsonError("配置参数为空")
	}

	plugin, err := models.PluginModel.GetPluginByPluginId(pluginId)
	if err != nil {
		this.ErrorLog("查找插件错误：" + err.Error())
		this.jsonError("查找插件错误")
	}
	if len(plugin) == 0 {
		this.jsonError("插件不存在")
	}

	// 更新插件
	_, err = models.PluginModel.UpdateConfValueByPluginId(pluginId, confValue)
	if err != nil {
		this.ErrorLog("修改插件配置错误：" + err.Error())
		this.jsonError("修改插件配置失败")
	}
	this.jsonSuccess("修改插件配置成功", nil, "/system/plugin/list")
}