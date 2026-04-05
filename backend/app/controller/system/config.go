package system

import (
	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
)

// ConfigList 获取系统配置列表
func ConfigList(ctx *gin.Context) error {
	configService := service.NewConfig(ctx)
	configsMap, err := configService.GetConfigsMap()
	if err != nil {
		sysLogErrorf(ctx, "[ConfigList] 获取系统配置失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	return RespJsonSuccess(ctx, configsMap)
}

// ConfigModify 修改系统配置
func ConfigModify(ctx *gin.Context) error {
	configKeys := []string{
		entity.ConfigKeyMainTitle,
		entity.ConfigKeyMainDescription,
		entity.ConfigKeyAutoFollowDoc,
		entity.ConfigKeySendEmail,
		entity.ConfigKeyAuthLogin,
		entity.ConfigKeyFulltextSearch,
		entity.ConfigKeyDocSearchTimer,
		entity.ConfigKeySystemName,
	}
	configService := service.NewConfig(ctx)
	for _, key := range configKeys {
		value := GetParamString(ctx, key)
		if value == "" {
			continue
		}
		err := configService.UpdateByKey(key, value)
		if err != nil {
			sysLogErrorf(ctx, "[ConfigModify] 更新配置 %s 失败: err=%+v", key, err)
			return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
		}
	}
	sysLogInfof(ctx, "[ConfigModify] 更新系统配置成功")
	return RespJsonSuccess(ctx, nil)
}
