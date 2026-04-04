package dao

import (
	"context"
	"time"

	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/utils"
	"gorm.io/gorm"
)

const (
	// TableNameConfig 系统配置表
	TableNameConfig = "mk_config"
	// ConfigPrimaryKey 配置表主键ID
	ConfigPrimaryKey = "config_id"
)

// Config 系统配置表数据对象
type Config struct {
	ctx context.Context
}

// NewConfig 创建系统配置表数据对象
func NewConfig(ctx context.Context) *Config {
	return &Config{
		ctx: ctx,
	}
}

// Insert 插入一条配置记录
func (r *Config) Insert(configEntity *entity.ConfigEntity) errors.BizError {
	configEntity.CreateTime = utils.NewJsonTime(time.Now())
	configEntity.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameConfig).Save(configEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetConfigByConfigId 根据配置ID获取配置信息
func (r *Config) GetConfigByConfigId(configId int64) (config *entity.ConfigEntity, err errors.BizError) {
	config = &entity.ConfigEntity{}
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameConfig).
		Where(map[string]interface{}{
			ConfigPrimaryKey: configId,
		}).
		First(&config)

	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return config, nil
}

// GetConfigByKey 根据配置键获取配置信息
func (r *Config) GetConfigByKey(key string) (config *entity.ConfigEntity, err errors.BizError) {
	config = &entity.ConfigEntity{}
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameConfig).
		Where(map[string]interface{}{
			"key": key,
		}).
		First(&config)

	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return config, nil
}

// GetAllConfigs 获取所有的配置
func (r *Config) GetAllConfigs() (configs []*entity.ConfigEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameConfig).
		Find(&configs)
	if db.Error != nil {
		return configs, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return configs, nil
}

// UpdateByKey 根据配置键更新配置值
func (r *Config) UpdateByKey(key string, value string) errors.BizError {
	updateTime := utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNameConfig).
		Where(map[string]interface{}{
			"key": key,
		}).
		Updates(map[string]interface{}{
			"value":       value,
			"update_time": updateTime,
		})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// GetConfigsMap 获取所有配置的 key=>value 映射
func (r *Config) GetConfigsMap() (map[string]string, errors.BizError) {
	configs, err := r.GetAllConfigs()
	if err != nil {
		return nil, err
	}
	configsMap := make(map[string]string)
	for _, config := range configs {
		configsMap[config.Key] = config.Value
	}
	return configsMap, nil
}
