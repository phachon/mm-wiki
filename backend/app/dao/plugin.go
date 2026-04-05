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
	// TableNamePlugin 插件表
	TableNamePlugin = "mk_plugin"
	// PluginPrimaryKey 插件表主键ID
	PluginPrimaryKey = "plugin_id"
)

// Plugin 插件表数据对象
type Plugin struct {
	ctx context.Context
}

// NewPlugin 创建插件表数据对象
func NewPlugin(ctx context.Context) *Plugin {
	return &Plugin{
		ctx: ctx,
	}
}

// Insert 创建插件插入一条记录
func (r *Plugin) Insert(pluginEntity *entity.PluginEntity) errors.BizError {
	pluginEntity.CreateTime = utils.NewJsonTime(time.Now())
	pluginEntity.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNamePlugin).Save(pluginEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetPluginByPluginId 根据插件ID获取插件信息
func (r *Plugin) GetPluginByPluginId(pluginId int64) (plugin *entity.PluginEntity, err errors.BizError) {
	plugin = &entity.PluginEntity{}
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNamePlugin).
		Where(map[string]interface{}{
			PluginPrimaryKey: pluginId,
		}).
		Where("status != ?", entity.PluginStatusDelete).
		First(&plugin)

	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return plugin, nil
}

// GetPluginByKey 根据插件标识获取插件
func (r *Plugin) GetPluginByKey(key string) (plugin *entity.PluginEntity, err errors.BizError) {
	plugin = &entity.PluginEntity{}
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNamePlugin).
		Where(map[string]interface{}{
			"key": key,
		}).
		Where("status != ?", entity.PluginStatusDelete).
		First(&plugin)

	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return plugin, nil
}

// GetPluginsByLimit 分页获取插件列表
func (r *Plugin) GetPluginsByLimit(limit int, offset int) (plugins []*entity.PluginEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNamePlugin).
		Where("status != ?", entity.PluginStatusDelete).
		Limit(limit).
		Offset(offset).
		Order("plugin_id DESC").
		Find(&plugins)

	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return plugins, nil
}

// GetPluginsByKeywordsAndLimit 根据关键字分页获取插件列表
func (r *Plugin) GetPluginsByKeywordsAndLimit(limit int, offset int, keywords *entity.PluginKeywords) (plugins []*entity.PluginEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNamePlugin)
	db = db.Where("status != ?", entity.PluginStatusDelete)
	if keywords.Name != "" {
		db = db.Where("name LIKE ?", "%"+keywords.Name+"%")
	}
	db = db.Limit(limit).
		Offset(offset).
		Order("plugin_id DESC").
		Find(&plugins)

	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return plugins, nil
}

// CountPlugins 获取插件总数
func (r *Plugin) CountPlugins() (int64, errors.BizError) {
	var count int64
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNamePlugin).
		Where("status != ?", entity.PluginStatusDelete).
		Count(&count)

	if db.Error != nil {
		return 0, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// CountPluginsByKeywords 根据关键字获取插件总数
func (r *Plugin) CountPluginsByKeywords(keywords *entity.PluginKeywords) (int64, errors.BizError) {
	var count int64
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNamePlugin)
	db = db.Where("status != ?", entity.PluginStatusDelete)
	if keywords.Name != "" {
		db = db.Where("name LIKE ?", "%"+keywords.Name+"%")
	}
	db = db.Count(&count)

	if db.Error != nil {
		return 0, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// Update 更新插件
func (r *Plugin) Update(pluginEntity entity.PluginEntity) errors.BizError {
	pluginEntity.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNamePlugin).
		Where(PluginPrimaryKey, pluginEntity.PluginId).
		Updates(map[string]interface{}{
			"name":        pluginEntity.Name,
			"description": pluginEntity.Description,
			"version":     pluginEntity.Version,
			"author":      pluginEntity.Author,
			"config_json": pluginEntity.ConfigJSON,
			"update_time": pluginEntity.UpdateTime,
		})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// DeletePlugin 删除插件（软删除）
func (r *Plugin) DeletePlugin(pluginId int64) errors.BizError {
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNamePlugin).
		Where(PluginPrimaryKey, pluginId).
		Update("status", entity.PluginStatusDelete)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}

// UpdateStatus 更新插件状态
func (r *Plugin) UpdateStatus(pluginId int64, status int) errors.BizError {
	db := GetDB(dbNameMK).WithContext(r.ctx).Table(TableNamePlugin).
		Where(PluginPrimaryKey, pluginId).
		Updates(map[string]interface{}{
			"status":      status,
			"update_time": utils.NewJsonTime(time.Now()),
		})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}
