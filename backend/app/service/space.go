package service

import (
	"context"

	"github.com/phachon/mm-wiki/app/dao"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
	"github.com/phachon/mm-wiki/utils"
)

const (
	SpaceListDefaultPageSize = 20 // 空间列表默认一页 20 条
)

// Space 空间业务逻辑
type Space struct {
	ctx          context.Context
	daoSpace     *dao.Space
	daoPrivilege *dao.Privilege
}

// NewSpace 创建空间业务逻辑对象
func NewSpace(ctx context.Context) *Space {
	return &Space{
		ctx:          ctx,
		daoSpace:     dao.NewSpace(ctx),
		daoPrivilege: dao.NewPrivilege(ctx),
	}
}

// Create 创建空间
func (s *Space) Create(spaceEntity *entity.SpaceEntity) errors.BizError {
	// 查找 name 是否存在
	space, err := s.daoSpace.GetSpaceByName(spaceEntity.Name)
	if err != nil {
		return err
	}
	if space != nil {
		return errors.Errorf(errors.BusinessRecordExistError, "空间名 %s 已经存在", spaceEntity.Name)
	}
	// 判断空间 key 是否存在
	isExists, err := s.daoSpace.CheckSpaceKeyExists(spaceEntity.SpaceKey)
	if err != nil {
		return err
	}
	if isExists {
		return errors.Errorf(errors.BusinessRecordExistError, "空间 key %s 已经存在", spaceEntity.SpaceKey)
	}
	// 插入一条记录
	err = s.daoSpace.Insert(spaceEntity)
	if err != nil {
		return err
	}
	return nil
}

// Update 修改空间
func (s *Space) Update(spaceEntity entity.SpaceEntity) errors.BizError {
	// 查找空间是否存在
	updateSpace, err := s.daoSpace.GetSpaceBySpaceId(spaceEntity.SpaceId)
	if err != nil {
		return err
	}
	if updateSpace == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "空间不存在")
	}
	// 查找 name 是否存在
	hasName, err := s.daoSpace.HasSameName(spaceEntity.SpaceId, spaceEntity.Name)
	if err != nil {
		logger.WithContext(s.ctx).Errorf("[service.Space] Update HasSameName spaceId=%d name=%s err=%s",
			spaceEntity.SpaceId, spaceEntity.Name, err.Error())
	}
	if hasName {
		return errors.Errorf(errors.BusinessRecordExistError, "空间名 %s 已经存在", spaceEntity.Name)
	}
	// 更新字段
	err = s.daoSpace.Update(spaceEntity)
	if err != nil {
		return err
	}
	return nil
}

// GetSpaceBySpaceId 根据空间ID获取空间详情
func (s *Space) GetSpaceBySpaceId(spaceID int64) (space *entity.SpaceEntity, err errors.BizError) {
	space, err = s.daoSpace.GetSpaceBySpaceId(spaceID)
	if err != nil {
		return nil, err
	}
	return space, nil
}

// DeleteSpace 删除空间
func (s *Space) DeleteSpace(spaceId int64) errors.BizError {
	// 查找空间是否存在
	spaceInfo, err := s.daoSpace.GetSpaceBySpaceId(spaceId)
	if err != nil {
		return err
	}
	if spaceInfo == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "空间id %d 不存在", spaceId)
	}
	// todo 删除空间下的所有文档

	// 删除空间
	err = s.daoSpace.DeleteSpace(spaceId)
	if err != nil {
		return err
	}
	// 删除空间对应的权限
	err = dao.NewSpacePermission(s.ctx).DeleteBySpaceId(spaceId)
	if err != nil {
		return err
	}

	return nil
}

// GetSpacesByLimit 分页获取空间列表
func (s *Space) GetSpacesByLimit(pageSize int, pageNum int, keywords *entity.SpaceKeywords) (spaces []*entity.SpaceEntity, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, SpaceListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	offset := (pageNum - 1) * pageSize
	// 无搜索词
	if keywords == nil {
		return s.daoSpace.GetSpacesByLimit(pageSize, offset)
	}
	return s.daoSpace.GetSpacesByKeywordsAndLimit(pageSize, offset, keywords)
}

// GetPageInfoLimit 获取分页信息
func (s *Space) GetPageInfoLimit(pageSize int, pageNum int, keywords *entity.SpaceKeywords) (pagination *entity.PageInfo, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, SpaceListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	pageInfo := new(entity.PageInfo)
	var totalNum int64
	// 无搜索词
	if keywords == nil {
		totalNum, err = s.daoSpace.CountSpaces()
	} else {
		totalNum, err = s.daoSpace.CountSpacesByKeywords(keywords)
	}
	if err != nil {
		logger.WithContext(s.ctx).Warnf("[service.Space] GetPageInfoLimit pageSize=%d pageNum=%d err=%s",
			pageSize, pageNum, err.Error())
		return pageInfo, err
	}
	return entity.GetPageInfo(totalNum, pageSize, pageNum), nil
}

// GetAllSpaces 获取所有的空间
func (s *Space) GetAllSpaces() (spaces []*entity.SpaceEntity, err errors.BizError) {
	return s.daoSpace.GetAllSpaces()
}

// GetSpacesBySpaceIds 根据空间ID批量获取空间
func (s *Space) GetSpacesBySpaceIds(spaceIds []int64) (spaces []*entity.SpaceEntity, err errors.BizError) {
	return s.daoSpace.GetSpacesBySpaceIds(spaceIds)
}

// FormatSpaceList 格式化账号列表根据账号信息
func (s *Space) FormatSpaceList(spaces []*entity.SpaceEntity) (spaceList []*entity.SpaceListItem, err errors.BizError) {
	if len(spaces) == 0 {
		return spaceList, nil
	}
	for _, space := range spaces {
		action := s.GetListItemAction(space)
		var spaceListItem = &entity.SpaceListItem{
			SpaceEntity: space,
			Action:      action,
		}
		spaceList = append(spaceList, spaceListItem)
	}
	return spaceList, nil
}

// GetListItemAction 获取列表 action 权限
func (s *Space) GetListItemAction(spaceItem *entity.SpaceEntity) *entity.SpaceListAction {
	indentifys := global.ContextValueLoginIdentifys(s.ctx)
	action := &entity.SpaceListAction{
		IsEdit:   indentifys[global.PrivilegeIndentifySpaceEdit],
		IsDelete: indentifys[global.PrivilegeIndentifySpaceDelete],
	}
	return action
}

// GetPublicSpacesByLimit 获取公开的空间列表
func (s *Space) GetPublicSpacesByLimit(limit, page int, keywords *entity.SpaceKeywords) ([]*entity.SpaceEntity, errors.BizError) {
	offset := (page - 1) * limit
	spaces, err := s.daoSpace.GetPublicSpacesByLimit(limit, offset, keywords)
	if err != nil {
		return nil, err
	}
	return spaces, nil
}

// GetPublicSpacesPageInfo 获取公开空间的分页信息
func (s *Space) GetPublicSpacesPageInfo(limit, page int, keywords *entity.SpaceKeywords) (*entity.PageInfo, errors.BizError) {
	totalNum, err := s.daoSpace.CountPublicSpaces(keywords)
	if err != nil {
		return nil, err
	}
	return entity.GetPageInfo(totalNum, limit, page), nil
}

// GetSpaceByKey 根据空间 Key 获取空间信息
func (s *Space) GetSpaceByKey(spaceKey string) (space *entity.SpaceEntity, err errors.BizError) {
	if len(spaceKey) == 0 {
		return nil, nil
	}
	return s.daoSpace.GetSpaceBySpaceKey(spaceKey)
}
