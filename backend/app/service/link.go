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
	LinkListDefaultPageSize = 20 // 链接列表默认一页 20 条
)

// Link 链接业务逻辑
type Link struct {
	ctx     context.Context
	daoLink *dao.Link
}

// NewLink 创建链接业务逻辑对象
func NewLink(ctx context.Context) *Link {
	return &Link{
		ctx:     ctx,
		daoLink: dao.NewLink(ctx),
	}
}

// Create 创建链接
func (l *Link) Create(linkEntity *entity.LinkEntity) errors.BizError {
	// 检查链接名是否已存在
	existLink, err := l.daoLink.GetLinkByName(linkEntity.Name)
	if err != nil {
		return err
	}
	if existLink != nil {
		return errors.Errorf(errors.BusinessRecordExistError, "链接名称已存在")
	}
	// 插入一条记录
	err = l.daoLink.Insert(linkEntity)
	if err != nil {
		return err
	}
	return nil
}

// Update 修改链接
func (l *Link) Update(linkEntity entity.LinkEntity) errors.BizError {
	// 查找链接是否存在
	updateLink, err := l.daoLink.GetLinkByLinkId(linkEntity.LinkId)
	if err != nil {
		return err
	}
	if updateLink == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "链接不存在")
	}
	// 检查链接名是否已被其他链接使用
	existLink, err := l.daoLink.GetLinkByName(linkEntity.Name)
	if err != nil {
		return err
	}
	if existLink != nil && existLink.LinkId != linkEntity.LinkId {
		return errors.Errorf(errors.BusinessRecordExistError, "链接名称已存在")
	}
	// 更新字段
	err = l.daoLink.Update(linkEntity)
	if err != nil {
		return err
	}
	return nil
}

// GetLinkByLinkId 根据链接ID获取链接详情
func (l *Link) GetLinkByLinkId(linkId int64) (link *entity.LinkEntity, err errors.BizError) {
	link, err = l.daoLink.GetLinkByLinkId(linkId)
	if err != nil {
		return link, err
	}
	return link, nil
}

// DeleteLink 删除链接
func (l *Link) DeleteLink(linkId int64) errors.BizError {
	// 查找链接是否存在
	updateLink, err := l.daoLink.GetLinkByLinkId(linkId)
	if err != nil {
		return err
	}
	if updateLink == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "链接id %d 不存在", linkId)
	}
	// 删除链接
	err = l.daoLink.DeleteLink(linkId)
	if err != nil {
		return err
	}
	return nil
}

// GetLinksByLimit 分页获取链接列表
func (l *Link) GetLinksByLimit(pageSize int, pageNum int, keywords *entity.LinkKeywords) (links []*entity.LinkEntity, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, LinkListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	offset := (pageNum - 1) * pageSize
	// 无搜索词
	if keywords == nil {
		return l.daoLink.GetLinksByLimit(pageSize, offset)
	}
	return l.daoLink.GetLinksByKeywordsAndLimit(pageSize, offset, keywords)
}

// GetPageInfoLimit 获取分页信息
func (l *Link) GetPageInfoLimit(pageSize int, pageNum int, keywords *entity.LinkKeywords) (pagination *entity.PageInfo, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, LinkListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	pageInfo := new(entity.PageInfo)
	var totalNum int64
	// 无搜索词
	if keywords == nil {
		totalNum, err = l.daoLink.CountLinks()
	} else {
		totalNum, err = l.daoLink.CountLinksByKeywords(keywords)
	}
	if err != nil {
		logger.WithContext(l.ctx).Warnf("[service.Link] GetPageInfoLimit pageSize=%d pageNum=%d err=%s",
			pageSize, pageNum, err.Error())
		return pageInfo, err
	}
	return entity.GetPageInfo(totalNum, pageSize, pageNum), nil
}

// FormatLinkList 格式化链接列表
func (l *Link) FormatLinkList(links []*entity.LinkEntity) (linkList []*entity.LinkListItem, err errors.BizError) {
	if len(links) == 0 {
		return linkList, nil
	}
	for _, link := range links {
		action := l.GetListItemAction(link)
		var linkListItem = &entity.LinkListItem{
			LinkEntity: link,
			Action:     action,
		}
		linkList = append(linkList, linkListItem)
	}
	return linkList, nil
}

// GetListItemAction 获取列表 action 权限
func (l *Link) GetListItemAction(linkItem *entity.LinkEntity) *entity.LinkListAction {
	indentifys := global.ContextValueLoginIdentifys(l.ctx)
	action := &entity.LinkListAction{
		IsEdit:   indentifys[global.PrivilegeIndentifyLinkEdit],
		IsDelete: indentifys[global.PrivilegeIndentifyLinkDelete],
	}
	return action
}

// GetAllLinks 获取所有的链接
func (l *Link) GetAllLinks() (links []*entity.LinkEntity, err errors.BizError) {
	return l.daoLink.GetAllLinks()
}
