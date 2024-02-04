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
	NoticeListDefaultPageSize = 20 // 公告列表默认一页 20 条
)

// Notice 公告业务逻辑
type Notice struct {
	ctx       context.Context
	daoNotice *dao.Notice
}

// NewNotice 创建公告业务逻辑对象
func NewNotice(ctx context.Context) *Notice {
	return &Notice{
		ctx:       ctx,
		daoNotice: dao.NewNotice(ctx),
	}
}

// Create 创建公告
func (n *Notice) Create(noticeEntity *entity.NoticeEntity) errors.BizError {
	// 插入一条记录
	err := n.daoNotice.Insert(noticeEntity)
	if err != nil {
		return err
	}
	return nil
}

// Update 修改公告
func (n *Notice) Update(noticeEntity entity.NoticeEntity) errors.BizError {
	// 查找公告是否存在
	updateNotice, err := n.daoNotice.GetNoticeByNoticeId(noticeEntity.NoticeId)
	if err != nil {
		return err
	}
	if updateNotice == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "公告不存在")
	}
	// 更新字段
	err = n.daoNotice.Update(noticeEntity)
	if err != nil {
		return err
	}
	return nil
}

// GetNoticeByNoticeId 根据公告ID获取公告详情
func (n *Notice) GetNoticeByNoticeId(noticeID int64) (notice *entity.NoticeEntity, err errors.BizError) {
	notice, err = n.daoNotice.GetNoticeByNoticeId(noticeID)
	if err != nil {
		return notice, err
	}
	return notice, nil
}

// DeleteNotice 删除公告
func (n *Notice) DeleteNotice(noticeId int64) errors.BizError {
	// 查找公告是否存在
	updateNotice, err := n.daoNotice.GetNoticeByNoticeId(noticeId)
	if err != nil {
		return err
	}
	if updateNotice == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "公告id %d 不存在", noticeId)
	}
	// 删除公告
	err = n.daoNotice.DeleteNotice(noticeId)
	if err != nil {
		return err
	}
	return nil
}

// GetNoticesByLimit 分页获取公告列表
func (n *Notice) GetNoticesByLimit(pageSize int, pageNum int, keywords *entity.NoticeKeywords) (notices []*entity.NoticeEntity, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, NoticeListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	offset := (pageNum - 1) * pageSize
	// 无搜索词
	if keywords == nil {
		return n.daoNotice.GetNoticesByLimit(pageSize, offset)
	}
	return n.daoNotice.GetNoticesByKeywordsAndLimit(pageSize, offset, keywords)
}

// GetPageInfoLimit 获取分页信息
func (n *Notice) GetPageInfoLimit(pageSize int, pageNum int, keywords *entity.NoticeKeywords) (pagination *entity.PageInfo, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, NoticeListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	pageInfo := new(entity.PageInfo)
	var totalNum int64
	// 无搜索词
	if keywords == nil {
		totalNum, err = n.daoNotice.CountNotices()
	} else {
		totalNum, err = n.daoNotice.CountNoticesByKeywords(keywords)
	}
	if err != nil {
		logger.WithContext(n.ctx).Warnf("[service.Notice] GetPageInfoLimit pageSize=%d pageNum=%d err=%s",
			pageSize, pageNum, err.Error())
		return pageInfo, err
	}
	return entity.GetPageInfo(totalNum, pageSize, pageNum), nil
}

// GetAllNotices 获取所有的公告
func (n *Notice) GetAllNotices() (notices []*entity.NoticeEntity, err errors.BizError) {
	return n.daoNotice.GetAllNotices()
}

// GetNoticesByNoticeIds 根据公告ID批量获取公告
func (n *Notice) GetNoticesByNoticeIds(noticeIds []int64) (notices []*entity.NoticeEntity, err errors.BizError) {
	return n.daoNotice.GetNoticesByNoticeIds(noticeIds)
}

// GetNoticesByLimit 分页获取公告列表
func (n *Notice) GetPublishNoticesByLimit(pageSize int, pageNum int) (notices []*entity.NoticeEntity, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, NoticeListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	offset := (pageNum - 1) * pageSize
	return n.daoNotice.GetPublishNoticesByLimit(pageSize, offset)
}

// CountPublishNotices 已发布的公告数
func (n *Notice) CountPublishNotices() (int64, errors.BizError) {
	return n.daoNotice.CountPublishNotices()
}

// FormatNoticeList 格式化账号列表根据账号信息
func (n *Notice) FormatNoticeList(notices []*entity.NoticeEntity) (noticeList []*entity.NoticeListItem, err errors.BizError) {
	if len(notices) == 0 {
		return noticeList, nil
	}
	for _, notice := range notices {
		action := n.GetListItemAction(notice)
		var noticeListItem = &entity.NoticeListItem{
			NoticeEntity: notice,
			Action:       action,
		}
		noticeList = append(noticeList, noticeListItem)
	}
	return noticeList, nil
}

// GetListItemAction 获取列表 action 权限
func (n *Notice) GetListItemAction(noticeItem *entity.NoticeEntity) *entity.NoticeListAction {
	indentifys := global.ContextValueLoginIdentifys(n.ctx)
	action := &entity.NoticeListAction{
		IsEdit:   indentifys[global.PrivilegeIndentifyNoticeEdit],
		IsDelete: indentifys[global.PrivilegeIndentifyNoticeDelete],
	}
	return action
}
