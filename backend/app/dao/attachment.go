package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/utils"
	"gorm.io/gorm"
)

const (
	// TableNameAttachment 附件信息表
	TableNameAttachment = "mk_attachment"
	// AttachmentPrimaryKey 附件表主键ID
	AttachmentPrimaryKey = "attachment_id"
)

// Attachment 附件信息表数据对象
type Attachment struct {
	ctx context.Context
}

// NewAttachment 创建附件信息表数据对象
func NewAttachment(ctx context.Context) *Attachment {
	return &Attachment{
		ctx: ctx,
	}
}

// Insert 创建附件记录
func (a *Attachment) Insert(attachmentEntity *entity.AttachmentEntity) errors.BizError {
	attachmentEntity.CreateTime = utils.NewJsonTime(time.Now())
	attachmentEntity.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameMK).WithContext(a.ctx).Table(TableNameAttachment).Save(attachmentEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetAttachmentByAttachmentId 根据附件ID获取附件信息
func (a *Attachment) GetAttachmentByAttachmentId(attachmentId int64) (attachment *entity.AttachmentEntity, err errors.BizError) {
	attachment = &entity.AttachmentEntity{}
	db := GetDB(dbNameMK).WithContext(a.ctx).Table(TableNameAttachment).
		Where(map[string]interface{}{
			AttachmentPrimaryKey: attachmentId,
		}).
		First(&attachment)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return attachment, nil
}

// GetAttachmentsByDocId 根据文档ID获取附件列表
func (a *Attachment) GetAttachmentsByDocId(docId string) (attachments []*entity.AttachmentEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(a.ctx).Table(TableNameAttachment).
		Where("doc_id = ?", docId).
		Order(fmt.Sprintf("%s DESC", AttachmentPrimaryKey)).
		Find(&attachments)
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return attachments, nil
}

// GetAttachmentsByDocIdAndSource 根据文档ID和来源获取附件列表
func (a *Attachment) GetAttachmentsByDocIdAndSource(docId string, source int) (attachments []*entity.AttachmentEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(a.ctx).Table(TableNameAttachment).
		Where("doc_id = ? AND source = ?", docId, source).
		Order(fmt.Sprintf("%s DESC", AttachmentPrimaryKey)).
		Find(&attachments)
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return attachments, nil
}

// GetAttachmentsByLimit 分页获取附件列表
func (a *Attachment) GetAttachmentsByLimit(limit int, offset int) (attachments []*entity.AttachmentEntity, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(a.ctx).Table(TableNameAttachment).
		Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", AttachmentPrimaryKey)).
		Find(&attachments)
	if db.Error != nil {
		return attachments, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return attachments, nil
}

// CountAttachments 获取附件总数
func (a *Attachment) CountAttachments() (count int64, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(a.ctx).Table(TableNameAttachment).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// GetAttachmentsByKeywordsAndLimit 根据关键字分页获取附件列表
func (a *Attachment) GetAttachmentsByKeywordsAndLimit(limit int, offset int, keywords *entity.AttachmentKeywords) (attachments []*entity.AttachmentEntity, err errors.BizError) {
	if keywords == nil {
		return a.GetAttachmentsByLimit(limit, offset)
	}
	db := GetDB(dbNameMK).WithContext(a.ctx).Table(TableNameAttachment)
	if keywords.DocId != "" {
		db = db.Where("doc_id = ?", keywords.DocId)
	}
	if keywords.Name != "" {
		db = db.Where("name LIKE ?", "%"+keywords.Name+"%")
	}
	db = db.Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", AttachmentPrimaryKey)).
		Find(&attachments)
	if db.Error != nil {
		return attachments, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return attachments, nil
}

// CountAttachmentsByKeywords 根据关键字获取附件总数
func (a *Attachment) CountAttachmentsByKeywords(keywords *entity.AttachmentKeywords) (count int64, err errors.BizError) {
	db := GetDB(dbNameMK).WithContext(a.ctx).Table(TableNameAttachment)
	if keywords.DocId != "" {
		db = db.Where("doc_id = ?", keywords.DocId)
	}
	if keywords.Name != "" {
		db = db.Where("name LIKE ?", "%"+keywords.Name+"%")
	}
	db.Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// DeleteAttachment 删除附件
func (a *Attachment) DeleteAttachment(attachmentId int64) errors.BizError {
	db := GetDB(dbNameMK).WithContext(a.ctx).Table(TableNameAttachment).
		Where(map[string]interface{}{
			AttachmentPrimaryKey: attachmentId,
		}).
		Delete(&entity.AttachmentEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}

// DeleteAttachmentsByDocId 根据文档ID删除所有附件
func (a *Attachment) DeleteAttachmentsByDocId(docId string) errors.BizError {
	db := GetDB(dbNameMK).WithContext(a.ctx).Table(TableNameAttachment).
		Where("doc_id = ?", docId).
		Delete(&entity.AttachmentEntity{})
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlDeleteErr, db.Error.Error())
	}
	return nil
}
