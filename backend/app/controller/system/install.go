package system

import (
	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/dao"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
)

// InstallStatus 获取安装状态
func InstallStatus(ctx *gin.Context) error {
	configService := service.NewConfig(ctx)
	version := configService.GetConfigValueByKey(entity.ConfigKeySystemVersion, "")

	isInstalled := version != ""
	return RespJsonSuccess(ctx, map[string]interface{}{
		"is_installed":   isInstalled,
		"system_version": version,
	})
}

// InstallCheckDB 检查数据库连接
func InstallCheckDB(ctx *gin.Context) error {
	// 测试数据库连接
	db := dao.GetDB("mm_wiki")
	if db == nil {
		return RespJsonError(ctx, int32(errors.DalMysqlErr), "数据库连接失败")
	}
	sqlDB, sqlErr := db.DB()
	if sqlErr != nil {
		return RespJsonError(ctx, int32(errors.DalMysqlErr), "获取数据库连接失败: "+sqlErr.Error())
	}
	pingErr := sqlDB.Ping()
	if pingErr != nil {
		return RespJsonError(ctx, int32(errors.DalMysqlErr), "数据库连接测试失败: "+pingErr.Error())
	}

	return RespJsonSuccess(ctx, map[string]interface{}{
		"status":  "connected",
		"message": "数据库连接成功",
	})
}

// InstallInitData 初始化系统数据
func InstallInitData(ctx *gin.Context) error {
	systemName := GetParamString(ctx, "system_name")
	mainTitle := GetParamString(ctx, "main_title")
	mainDescription := GetParamString(ctx, "main_description")

	if systemName == "" {
		systemName = "MM-Wiki"
	}
	if mainTitle == "" {
		mainTitle = "MM-Wiki 文档管理系统"
	}

	configService := service.NewConfig(ctx)

	// 检查是否已安装
	existingVersion := configService.GetConfigValueByKey(entity.ConfigKeySystemVersion, "")
	if existingVersion != "" {
		return RespJsonError(ctx, int32(errors.BusinessForbiddenError), "系统已完成安装，不能重复初始化")
	}

	// 初始化系统配置
	configUpdates := map[string]string{
		entity.ConfigKeySystemName:      systemName,
		entity.ConfigKeyMainTitle:       mainTitle,
		entity.ConfigKeyMainDescription: mainDescription,
		entity.ConfigKeySendEmail:       "0",
		entity.ConfigKeyAutoFollowDoc:   "0",
		entity.ConfigKeyAuthLogin:       "0",
		entity.ConfigKeyFulltextSearch:  "0",
		entity.ConfigKeyDocSearchTimer:  "3600",
	}

	for key, value := range configUpdates {
		err := configService.UpdateByKey(key, value)
		if err != nil {
			logger.WithContext(ctx).Warnf("[InstallInitData] 更新配置 %s 失败: err=%+v", key, err)
			// 配置不存在时跳过
			continue
		}
	}

	sysLogInfof(ctx, "[InstallInitData] 系统基本配置初始化完成")
	return RespJsonSuccess(ctx, nil)
}

// InstallCreateAdmin 创建管理员账号
func InstallCreateAdmin(ctx *gin.Context) error {
	accountName := GetParamString(ctx, "account_name")
	password := GetParamString(ctx, "password")
	email := GetParamString(ctx, "email")
	givenName := GetParamString(ctx, "given_name")

	if accountName == "" {
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "账号名称不能为空")
	}
	if password == "" {
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "密码不能为空")
	}
	if len(password) < 6 {
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "密码长度不能少于6位")
	}

	// 创建管理员账号
	accountEntity := &entity.AccountEntity{
		Name:      accountName,
		Password:  password,
		Email:     email,
		GivenName: givenName,
		Status:    entity.AccountStatusDefault,
	}

	err := service.NewAccount(ctx).Create(accountEntity)
	if err != nil {
		sysLogErrorf(ctx, "[InstallCreateAdmin] 创建管理员账号失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	sysLogInfof(ctx, "[InstallCreateAdmin] 创建管理员账号 %s 成功", accountName)
	return RespJsonSuccess(ctx, nil)
}

// InstallComplete 完成安装
func InstallComplete(ctx *gin.Context) error {
	configService := service.NewConfig(ctx)

	// 检查是否已安装
	existingVersion := configService.GetConfigValueByKey(entity.ConfigKeySystemVersion, "")
	if existingVersion != "" {
		return RespJsonError(ctx, int32(errors.BusinessForbiddenError), "系统已完成安装")
	}

	// 设置系统版本号标记安装完成
	err := configService.UpdateByKey(entity.ConfigKeySystemVersion, "2.0.0")
	if err != nil {
		logger.WithContext(ctx).Warnf("[InstallComplete] 设置系统版本失败: err=%+v", err)
		// 如果配置键不存在，忽略错误
	}

	sysLogInfof(ctx, "[InstallComplete] 系统安装完成")
	return RespJsonSuccess(ctx, map[string]interface{}{
		"message": "安装完成",
	})
}
