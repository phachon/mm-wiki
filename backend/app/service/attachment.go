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
	AttachmentListDefaultPageSize = 20 // 附件列表默认一页 20 条
)

// Attachment 附件业务逻辑
type Attachment struct {
	ctx           context.Context
	daoAttachment *dao.Attachment
}

// NewAttachment 创建附件业务逻辑对象
func NewAttachment(ctx context.Context) *Attachment {
	return &Attachment{
		ctx:           ctx,
		daoAttachment: dao.NewAttachment(ctx),
	}
}

// Create 创建附件
func (a *Attachment) Create(attachmentEntity *entity.AttachmentEntity) errors.BizError {
	err := a.daoAttachment.Insert(attachmentEntity)
	if err != nil {
		return err
	}
	return nil
}

// GetAttachmentByAttachmentId 根据附件ID获取附件详情
func (a *Attachment) GetAttachmentByAttachmentId(attachmentId int64) (attachment *entity.AttachmentEntity, err errors.BizError) {
	attachment, err = a.daoAttachment.GetAttachmentByAttachmentId(attachmentId)
	if err != nil {
		return attachment, err
	}
	return attachment, nil
}

// GetAttachmentsByDocId 根据文档ID获取附件列表
func (a *Attachment) GetAttachmentsByDocId(docId string) (attachments []*entity.AttachmentEntity, err errors.BizError) {
	return a.daoAttachment.GetAttachmentsByDocId(docId)
}

// GetAttachmentsByDocIdAndSource 根据文档ID和来源获取附件列表
func (a *Attachment) GetAttachmentsByDocIdAndSource(docId string, source int) (attachments []*entity.AttachmentEntity, err errors.BizError) {
	return a.daoAttachment.GetAttachmentsByDocIdAndSource(docId, source)
}

// DeleteAttachment 删除附件
func (a *Attachment) DeleteAttachment(attachmentId int64) errors.BizError {
	// 查找附件是否存在
	attachment, err := a.daoAttachment.GetAttachmentByAttachmentId(attachmentId)
	if err != nil {
		return err
	}
	if attachment == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "附件id %d 不存在", attachmentId)
	}
	// 删除附件
	err = a.daoAttachment.DeleteAttachment(attachmentId)
	if err != nil {
		return err
	}
	return nil
}

// DeleteAttachmentsByDocId 根据文档ID删除所有附件
func (a *Attachment) DeleteAttachmentsByDocId(docId string) errors.BizError {
	return a.daoAttachment.DeleteAttachmentsByDocId(docId)
}

// GetAttachmentsByLimit 分页获取附件列表
func (a *Attachment) GetAttachmentsByLimit(pageSize int, pageNum int, keywords *entity.AttachmentKeywords) (attachments []*entity.AttachmentEntity, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, AttachmentListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	offset := (pageNum - 1) * pageSize
	if keywords == nil {
		return a.daoAttachment.GetAttachmentsByLimit(pageSize, offset)
	}
	return a.daoAttachment.GetAttachmentsByKeywordsAndLimit(pageSize, offset, keywords)
}

// GetPageInfoLimit 获取分页信息
func (a *Attachment) GetPageInfoLimit(pageSize int, pageNum int, keywords *entity.AttachmentKeywords) (pagination *entity.PageInfo, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, AttachmentListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	pageInfo := new(entity.PageInfo)
	var totalNum int64
	if keywords == nil {
		totalNum, err = a.daoAttachment.CountAttachments()
	} else {
		totalNum, err = a.daoAttachment.CountAttachmentsByKeywords(keywords)
	}
	if err != nil {
		logger.WithContext(a.ctx).Warnf("[service.Attachment] GetPageInfoLimit pageSize=%d pageNum=%d err=%s",
			pageSize, pageNum, err.Error())
		return pageInfo, err
	}
	return entity.GetPageInfo(totalNum, pageSize, pageNum), nil
}
