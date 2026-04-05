package service

import (
	"context"

	"github.com/phachon/mm-wiki/app/dao"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
	"github.com/phachon/mm-wiki/utils"
)

const (
	PluginListDefaultPageSize = 20 // 插件列表默认一页 20 条
)

// PluginService 插件业务逻辑
type PluginService struct {
	ctx       context.Context
	daoPlugin *dao.Plugin
}

// NewPlugin 创建插件业务逻辑对象
func NewPluginService(ctx context.Context) *PluginService {
	return &PluginService{
		ctx:       ctx,
		daoPlugin: dao.NewPlugin(ctx),
	}
}

// Create 创建插件
func (p *PluginService) Create(pluginEntity *entity.PluginEntity) errors.BizError {
	// 检查插件标识是否已存在
	existPlugin, err := p.daoPlugin.GetPluginByKey(pluginEntity.Key)
	if err != nil {
		return err
	}
	if existPlugin != nil {
		return errors.Errorf(errors.BusinessRecordExistError, "插件标识 %s 已经存在", pluginEntity.Key)
	}
	// 插入一条记录
	err = p.daoPlugin.Insert(pluginEntity)
	if err != nil {
		return err
	}
	return nil
}

// Update 修改插件
func (p *PluginService) Update(pluginEntity entity.PluginEntity) errors.BizError {
	// 查找插件是否存在
	existPlugin, err := p.daoPlugin.GetPluginByPluginId(pluginEntity.PluginId)
	if err != nil {
		return err
	}
	if existPlugin == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "插件不存在")
	}
	// 更新字段
	err = p.daoPlugin.Update(pluginEntity)
	if err != nil {
		return err
	}
	return nil
}

// GetPluginByPluginId 根据插件ID获取插件详情
func (p *PluginService) GetPluginByPluginId(pluginId int64) (*entity.PluginEntity, errors.BizError) {
	return p.daoPlugin.GetPluginByPluginId(pluginId)
}

// DeletePlugin 删除插件
func (p *PluginService) DeletePlugin(pluginId int64) errors.BizError {
	existingPlugin, err := p.daoPlugin.GetPluginByPluginId(pluginId)
	if err != nil {
		return err
	}
	if existingPlugin == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "插件id %d 不存在", pluginId)
	}
	err = p.daoPlugin.DeletePlugin(pluginId)
	if err != nil {
		return err
	}
	return nil
}

// UpdateStatus 更新插件状态
func (p *PluginService) UpdateStatus(pluginId int64, status int) errors.BizError {
	existingPlugin, err := p.daoPlugin.GetPluginByPluginId(pluginId)
	if err != nil {
		return err
	}
	if existingPlugin == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "插件id %d 不存在", pluginId)
	}
	err = p.daoPlugin.UpdateStatus(pluginId, status)
	if err != nil {
		return err
	}
	return nil
}

// GetPluginsByLimit 分页获取插件列表
func (p *PluginService) GetPluginsByLimit(pageSize int, pageNum int, keywords *entity.PluginKeywords) ([]*entity.PluginEntity, errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, PluginListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	offset := (pageNum - 1) * pageSize
	if keywords == nil {
		return p.daoPlugin.GetPluginsByLimit(pageSize, offset)
	}
	return p.daoPlugin.GetPluginsByKeywordsAndLimit(pageSize, offset, keywords)
}

// GetPageInfoLimit 获取分页信息
func (p *PluginService) GetPageInfoLimit(pageSize int, pageNum int, keywords *entity.PluginKeywords) (*entity.PageInfo, errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, PluginListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	var totalNum int64
	var err errors.BizError
	if keywords == nil {
		totalNum, err = p.daoPlugin.CountPlugins()
	} else {
		totalNum, err = p.daoPlugin.CountPluginsByKeywords(keywords)
	}
	if err != nil {
		logger.WithContext(p.ctx).Warnf("[service.Plugin] GetPageInfoLimit err=%s", err.Error())
		return new(entity.PageInfo), err
	}
	return entity.GetPageInfo(totalNum, pageSize, pageNum), nil
}

// FormatPluginList 格式化插件列表
func (p *PluginService) FormatPluginList(plugins []*entity.PluginEntity) []*entity.PluginListItem {
	var pluginList []*entity.PluginListItem
	if len(plugins) == 0 {
		return pluginList
	}
	for _, plugin := range plugins {
		pluginListItem := &entity.PluginListItem{
			PluginEntity: plugin,
			Action: &entity.PluginListAction{
				IsEdit:   1,
				IsDelete: 1,
			},
		}
		pluginList = append(pluginList, pluginListItem)
	}
	return pluginList
}
