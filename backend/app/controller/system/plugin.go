package system

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
)

// PluginSave 添加插件保存
func PluginSave(ctx *gin.Context) error {
	name := GetParamString(ctx, "name")
	key := GetParamString(ctx, "key")
	description := GetParamString(ctx, "description")
	version := GetParamString(ctx, "version")
	author := GetParamString(ctx, "author")
	configJSON := GetParamString(ctx, "config_json")

	if name == "" {
		logger.WithContext(ctx).Warnf("[PluginSave] name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "插件名称不能为空")
	}
	if key == "" {
		logger.WithContext(ctx).Warnf("[PluginSave] key empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "插件标识不能为空")
	}

	pluginEntity := &entity.PluginEntity{
		Name:        name,
		Key:         key,
		Description: description,
		Version:     version,
		Author:      author,
		ConfigJSON:  configJSON,
		Status:      entity.PluginStatusDisabled,
	}

	err := service.NewPluginService(ctx).Create(pluginEntity)
	if err != nil {
		sysLogErrorf(ctx, "[PluginSave] 添加插件失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[PluginSave] 添加插件 %d 成功", pluginEntity.PluginId)
	return RespJsonSuccess(ctx, nil)
}

// PluginEdit 插件编辑页面
func PluginEdit(ctx *gin.Context) error {
	pluginId := GetParamInt64(ctx, "plugin_id")
	if pluginId == 0 {
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "插件id不存在")
	}

	plugin, err := service.NewPluginService(ctx).GetPluginByPluginId(pluginId)
	if err != nil {
		sysLogErrorf(ctx, "[PluginEdit] 获取插件 %d 信息失败: err=%+v", pluginId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if plugin == nil {
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "插件id错误")
	}

	return RespJsonSuccess(ctx, map[string]interface{}{
		"plugin_info": plugin,
	})
}

// PluginModify 修改插件保存
func PluginModify(ctx *gin.Context) error {
	pluginId := GetParamInt64(ctx, "plugin_id")
	name := GetParamString(ctx, "name")
	description := GetParamString(ctx, "description")
	version := GetParamString(ctx, "version")
	author := GetParamString(ctx, "author")
	configJSON := GetParamString(ctx, "config_json")

	if pluginId <= 0 {
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "插件ID不合法")
	}
	if name == "" {
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "插件名称不能为空")
	}

	pluginEntity := entity.PluginEntity{
		PluginId:    pluginId,
		Name:        name,
		Description: description,
		Version:     version,
		Author:      author,
		ConfigJSON:  configJSON,
	}

	err := service.NewPluginService(ctx).Update(pluginEntity)
	if err != nil {
		sysLogErrorf(ctx, "[PluginModify] 更新插件 %d 失败: err=%+v", pluginId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[PluginModify] 更新插件 %d 成功", pluginId)
	return RespJsonSuccess(ctx, nil)
}

// PluginList 插件列表
func PluginList(ctx *gin.Context) error {
	pageSize := GetParamIntDef(ctx, "page_size", 20)
	pageNum := GetParamIntDef(ctx, "page_num", 1)
	keywordsStr := GetParamStringDef(ctx, "keywords", "")

	var keywords *entity.PluginKeywords
	if len(keywordsStr) > 0 {
		jErr := json.Unmarshal([]byte(keywordsStr), &keywords)
		if jErr != nil {
			logger.WithContext(ctx).Errorf("[PluginList] parse keywords err=%s", jErr.Error())
		}
	}

	pluginService := service.NewPluginService(ctx)

	plugins, err := pluginService.GetPluginsByLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[PluginList] 获取插件列表失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	pageInfo, err := pluginService.GetPageInfoLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[PluginList] 获取插件分页信息失败: err=%s", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	pluginList := pluginService.FormatPluginList(plugins)

	data := map[string]interface{}{
		"list":      pluginList,
		"page_info": pageInfo,
	}
	return RespJsonSuccess(ctx, data)
}

// PluginDelete 删除插件
func PluginDelete(ctx *gin.Context) error {
	pluginId := GetParamInt64(ctx, "plugin_id")
	if pluginId <= 0 {
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "插件id不存在")
	}

	err := service.NewPluginService(ctx).DeletePlugin(pluginId)
	if err != nil {
		sysLogErrorf(ctx, "[PluginDelete] 删除插件 %d 失败: err=%+v", pluginId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[PluginDelete] 删除插件 %d 成功", pluginId)
	return RespJsonSuccess(ctx, nil)
}

// PluginUpdateStatus 更新插件状态（启用/禁用）
func PluginUpdateStatus(ctx *gin.Context) error {
	pluginId := GetParamInt64(ctx, "plugin_id")
	status := GetParamInt(ctx, "status")

	if pluginId <= 0 {
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "插件id不存在")
	}
	if status != entity.PluginStatusEnabled && status != entity.PluginStatusDisabled {
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "状态值不合法")
	}

	err := service.NewPluginService(ctx).UpdateStatus(pluginId, status)
	if err != nil {
		sysLogErrorf(ctx, "[PluginUpdateStatus] 更新插件 %d 状态失败: err=%+v", pluginId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[PluginUpdateStatus] 更新插件 %d 状态为 %d 成功", pluginId, status)
	return RespJsonSuccess(ctx, nil)
}
