package service

import (
	"context"

	"github.com/phachon/mm-wiki/app/dao"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/gopkg/errors"
)

// Doc 文档服务
type Doc struct {
	ctx        context.Context
	daoDoc     *dao.Doc
	daoContent *dao.Content
}

// NewDoc 创建文档服务
func NewDoc(ctx context.Context) *Doc {
	return &Doc{
		ctx:        ctx,
		daoDoc:     dao.NewDoc(ctx),
		daoContent: dao.NewContent(ctx),
	}
}

// DocsToTree 递归实现文档转换为树形结构
func (d *Doc) DocsToTree(docs []*entity.DocEntity, parentId int64) []*entity.DocTreeEntity {
	docTree := make([]*entity.DocTreeEntity, 0)
	for _, doc := range docs {
		if doc.ParentId == parentId {
			docTree = append(docTree, &entity.DocTreeEntity{
				DocEntity: *doc,
				Children:  d.DocsToTree(docs, doc.DocId),
			})
		}
	}
	return docTree
}

// GetDocsBySpaceKey 获取空间下所有文档
func (d *Doc) GetDocsBySpaceKey(spaceKey string) ([]*entity.DocEntity, errors.BizError) {
	if spaceKey == "" {
		return nil, nil
	}
	return d.daoDoc.GetDocsBySpaceKey(spaceKey)
}

// GetDocByDocId 获取文档信息
func (d *Doc) GetDocByDocId(docId int64) (*entity.DocEntity, errors.BizError) {
	return d.daoDoc.GetDocByDocId(docId)
}

// GetDocByDocIds 获取多个文档信息
func (d *Doc) GetDocByDocIds(docIds []int64) ([]*entity.DocEntity, errors.BizError) {
	return d.daoDoc.GetDocByDocIds(docIds)
}

// CreateDoc 创建文档
func (d *Doc) CreateDoc(docEntity *entity.DocEntity) errors.BizError {
	if docEntity == nil {
		return nil
	}
	docEntity.CreateAccountId = global.ContextValueLoginAccountID(d.ctx)
	docEntity.CreateAccountName = global.ContextValueLoginAccountName(d.ctx)
	docEntity.EditAccountId = global.ContextValueLoginAccountID(d.ctx)
	docEntity.EditAccountName = global.ContextValueLoginAccountName(d.ctx)
	err := d.daoDoc.Insert(docEntity)
	if err != nil {
		return err
	}
	// 创建文档内容
	contentEntity := &entity.ContentEntity{
		DocId:   docEntity.DocId,
		Content: "",
	}
	return d.daoContent.Insert(contentEntity)
}

// UpdateDoc 更新文档
func (d *Doc) UpdateDoc(docEntity *entity.DocEntity) errors.BizError {
	return d.daoDoc.UpdateDoc(docEntity)
}

// UpdateNameAndEditAccount 更新文档名称和编辑人
func (d *Doc) UpdateNameAndEditAccount(docId int64, name string) errors.BizError {
	return d.daoDoc.UpdateNameAndEditAccount(
		docId,
		name,
		global.ContextValueLoginAccountID(d.ctx),
		global.ContextValueLoginAccountName(d.ctx),
	)
}

// DeleteDoc 删除文档
func (d *Doc) DeleteDoc(docId int64) errors.BizError {
	return d.daoDoc.DeleteDoc(docId)
}

// UpdateDocSequence 更新文档排序
func (d *Doc) UpdateDocSequence(docId int64, sequence int) errors.BizError {
	return d.daoDoc.UpdateDocSequence(docId, sequence)
}

// SearchDocs 搜索文档
func (d *Doc) SearchDocs(keyword string, spaceKey string, pageSize int, pageNum int) ([]*entity.DocEntity, errors.BizError) {
	offset := (pageNum - 1) * pageSize
	return d.daoDoc.SearchDocs(keyword, spaceKey, pageSize, offset)
}

// CountSearchDocs 统计搜索文档数量
func (d *Doc) CountSearchDocs(keyword string, spaceKey string) (int64, errors.BizError) {
	return d.daoDoc.CountSearchDocs(keyword, spaceKey)
}

// CreateSpaceHomeDoc 创建空间主页文档
func (d *Doc) CreateSpaceHomeDoc(spaceId int64, spaceKey string, title string) errors.BizError {
	docEntity := &entity.DocEntity{
		SpaceId:           spaceId,
		ParentId:          0,
		Name:              title,
		SpaceKey:          spaceKey,
		Type:              entity.DirEntityType,
		CreateAccountId:   global.ContextValueLoginAccountID(d.ctx),
		CreateAccountName: global.ContextValueLoginAccountName(d.ctx),
		EditAccountId:     global.ContextValueLoginAccountID(d.ctx),
		EditAccountName:   global.ContextValueLoginAccountName(d.ctx),
	}
	return d.CreateDoc(docEntity)
}
