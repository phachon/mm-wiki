package service

import (
	"context"

	"github.com/phachon/mm-wiki/app/dao"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
)

// Config 配置业务逻辑
type Config struct {
	ctx       context.Context
	daoConfig *dao.Config
}

// NewConfig 创建配置业务逻辑对象
func NewConfig(ctx context.Context) *Config {
	return &Config{
		ctx:       ctx,
		daoConfig: dao.NewConfig(ctx),
	}
}

// GetConfigByConfigId 根据配置ID获取配置详情
func (c *Config) GetConfigByConfigId(configId int64) (*entity.ConfigEntity, errors.BizError) {
	return c.daoConfig.GetConfigByConfigId(configId)
}

// GetConfigByKey 根据配置键获取配置详情
func (c *Config) GetConfigByKey(key string) (*entity.ConfigEntity, errors.BizError) {
	return c.daoConfig.GetConfigByKey(key)
}

// GetAllConfigs 获取所有的配置
func (c *Config) GetAllConfigs() ([]*entity.ConfigEntity, errors.BizError) {
	return c.daoConfig.GetAllConfigs()
}

// GetConfigsMap 获取所有配置的 key=>value 映射
func (c *Config) GetConfigsMap() (map[string]string, errors.BizError) {
	return c.daoConfig.GetConfigsMap()
}

// UpdateByKey 根据配置键更新配置值
func (c *Config) UpdateByKey(key string, value string) errors.BizError {
	// 查找配置是否存在
	config, err := c.daoConfig.GetConfigByKey(key)
	if err != nil {
		return err
	}
	if config == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "配置 %s 不存在", key)
	}
	return c.daoConfig.UpdateByKey(key, value)
}

// GetConfigValueByKey 根据配置键获取配置值，不存在则返回默认值
func (c *Config) GetConfigValueByKey(key string, defaultValue string) string {
	config, err := c.daoConfig.GetConfigByKey(key)
	if err != nil || config == nil {
		return defaultValue
	}
	return config.Value
}
