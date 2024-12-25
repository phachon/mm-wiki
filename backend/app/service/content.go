package service

import (
	"context"

	"github.com/phachon/mm-wiki/app/dao"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
)

// Content 文档内容服务
type Content struct {
	ctx               context.Context
	daoContent        *dao.Content
	daoContentVersion *dao.ContentVersion
}

// NewContent 创建文档内容服务
func NewContent(ctx context.Context) *Content {
	return &Content{
		ctx:               ctx,
		daoContent:        dao.NewContent(ctx),
		daoContentVersion: dao.NewContentVersion(ctx),
	}
}

// SaveContent 保存文档内容
func (c *Content) Create(docId int64, content string) errors.BizError {
	contentEntity := &entity.ContentEntity{
		DocId:   docId,
		Content: content,
	}
	// 保存文档内容
	err := c.daoContent.Insert(contentEntity)
	if err != nil {
		return err
	}
	return nil
}

// GetContentByDocId 获取文档内容
func (c *Content) GetContentByDocId(docId int64) (*entity.ContentEntity, errors.BizError) {
	content, err := c.daoContent.GetContentByDocId(docId)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// UpdateContent 更新文档内容
func (c *Content) UpdateContent(docId int64, content string, oldContent string) errors.BizError {
	if oldContent == content {
		return nil
	}
	contentEntity := &entity.ContentEntity{
		DocId:   docId,
		Content: content,
	}
	// 更新文档内容
	err := c.daoContent.UpdateContent(contentEntity)
	if err != nil {
		return err
	}

	// 保存文档版本
	contentVersionEntity := &entity.ContentVersionEntity{
		DocId:   docId,
		Content: oldContent,
	}
	err = c.daoContentVersion.Insert(contentVersionEntity)
	if err != nil {
		return err
	}

	return nil
}
