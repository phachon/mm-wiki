package service

import (
	"context"

	"github.com/phachon/mm-wiki/app/dao"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/utils"
)

const (
	ContentVersionDefaultPageSize = 10 // 文档内容版本默认一页 10 条
)

// ContentVersion 文档内容版本服务
type ContentVersion struct {
	ctx               context.Context
	daoContentVersion *dao.ContentVersion
}

// NewContentVersion 创建文档内容版本服务
func NewContentVersion(ctx context.Context) *ContentVersion {
	return &ContentVersion{
		ctx:               ctx,
		daoContentVersion: dao.NewContentVersion(ctx),
	}
}

// CreateContentVersion 创建文档内容版本
func (c *ContentVersion) CreateContentVersion(contentVersionEntity *entity.ContentVersionEntity) errors.BizError {
	if contentVersionEntity.DocId <= 0 {
		return nil
	}
	err := c.daoContentVersion.Insert(contentVersionEntity)
	if err != nil {
		return err
	}
	// 删除多余版本
	return c.DeleteContentVersionsByDocIdKeepN(contentVersionEntity.DocId, 10)
}

// GetContentVersionsByDocIdLimit 获取文档多个版本信息(不含正文)
func (c *ContentVersion) GetContentVersionsByDocIdLimit(docId int64, pageSize int, pageNum int) (
	[]*entity.ContentVersionEntity, errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, ContentVersionDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	offset := (pageNum - 1) * pageSize
	return c.daoContentVersion.GetContentVersionsByDocIdLimit(docId, pageSize, offset)
}

// GetPageInfoLimit 获取分页信息
func (c *ContentVersion) GetPageInfoLimit(
	docId int64,
	pageSize int,
	pageNum int,
) (pagination *entity.PageInfo, err errors.BizError) {

	pageSize = utils.VerifyUint(pageSize, ContentVersionDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	pageInfo := new(entity.PageInfo)
	var totalNum int64
	totalNum, err = c.daoContentVersion.CountContentVersions(docId)
	if err != nil {
		return pageInfo, err
	}
	return entity.GetPageInfo(totalNum, pageSize, pageNum), nil
}

// GetContentVersionByVersionId 获取文档内容版本信息
func (c *ContentVersion) GetContentVersionByVersionId(contentVersionId int64) (*entity.ContentVersionEntity, errors.BizError) {
	return c.daoContentVersion.GetContentVersionByVersionId(contentVersionId)
}

// DeleteContentVersion 删除文档内容版本
func (c *ContentVersion) DeleteContentVersion(contentVersionId int64) errors.BizError {
	if contentVersionId <= 0 {
		return nil
	}
	return c.daoContentVersion.DeleteContentVersion(contentVersionId)
}

// DeleteContentVersionsByDocIdKeepN
func (c *ContentVersion) DeleteContentVersionsByDocIdKeepN(docId int64, keepNum int) errors.BizError {
	if docId <= 0 || keepNum <= 0 {
		return nil
	}
	// 获取要保留的版本ID
	contentVersions, err := c.daoContentVersion.GetContentVersionsByDocIdLimit(docId, keepNum, 0)
	if err != nil {
		return err
	}
	var contentVersionIds []int64
	for _, contentVersion := range contentVersions {
		contentVersionIds = append(contentVersionIds, contentVersion.ContentVersionId)
	}
	return c.daoContentVersion.DeleteContentVersionsNotInVersionIds(contentVersionIds)
}
