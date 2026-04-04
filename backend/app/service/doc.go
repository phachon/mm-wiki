package service

import (
	"context"
	"fmt"

	"github.com/phachon/mm-wiki/app/dao"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
)

// Doc 文档服务
type Doc struct {
	ctx               context.Context
	daoDoc            *dao.Doc
	daoContent        *dao.Content
	daoContentVersion *dao.ContentVersion
	daoCollection     *dao.Collection
	daoFollow         *dao.Follow
	daoLogDoc         *dao.LogDoc
}

// NewDoc 创建文档服务
func NewDoc(ctx context.Context) *Doc {
	return &Doc{
		ctx:               ctx,
		daoDoc:            dao.NewDoc(ctx),
		daoContent:        dao.NewContent(ctx),
		daoContentVersion: dao.NewContentVersion(ctx),
		daoCollection:     dao.NewCollection(ctx),
		daoFollow:         dao.NewFollow(ctx),
		daoLogDoc:         dao.NewLogDoc(ctx),
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
	err = d.daoContent.Insert(contentEntity)
	if err != nil {
		return err
	}

	// 记录文档创建日志
	d.logDocAction(docEntity.DocId, docEntity.SpaceId, entity.LogDocActionCreate, "创建文档: "+docEntity.Name)

	return nil
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

// DeleteDocWithRelated 删除文档及其关联数据（正文、版本、收藏、关注）
func (d *Doc) DeleteDocWithRelated(docId int64) errors.BizError {
	docIdStr := fmt.Sprintf("%d", docId)

	// 获取文档信息用于日志
	doc, getErr := d.daoDoc.GetDocByDocId(docId)
	if getErr != nil {
		logger.WithContext(d.ctx).Errorf("[DeleteDocWithRelated] GetDocByDocId docId=%d err=%+v", docId, getErr)
	}
	var spaceId int64
	var docName string
	if doc != nil {
		spaceId = doc.SpaceId
		docName = doc.Name
	}

	// 删除文档记录
	err := d.daoDoc.DeleteDoc(docId)
	if err != nil {
		return err
	}

	// 删除文档正文
	err = d.daoContent.DeleteContentByDocId(docId)
	if err != nil {
		logger.WithContext(d.ctx).Errorf("[DeleteDocWithRelated] DeleteContentByDocId docId=%d err=%+v", docId, err)
	}

	// 删除文档版本
	err = d.daoContentVersion.DeleteContentVersionsByDocId(docId)
	if err != nil {
		logger.WithContext(d.ctx).Errorf("[DeleteDocWithRelated] DeleteContentVersionsByDocId docId=%d err=%+v", docId, err)
	}

	// 删除文档收藏
	err = d.daoCollection.DeleteByResourceIdAndType(entity.CollectionTypeDocument, docIdStr)
	if err != nil {
		logger.WithContext(d.ctx).Errorf("[DeleteDocWithRelated] DeleteByResourceIdAndType docId=%d err=%+v", docId, err)
	}

	// 删除文档关注
	err = d.daoFollow.DeleteByObjectIdAndType(entity.FollowTypeDocument, docIdStr)
	if err != nil {
		logger.WithContext(d.ctx).Errorf("[DeleteDocWithRelated] DeleteByObjectIdAndType docId=%d err=%+v", docId, err)
	}

	// 记录文档删除日志
	d.logDocAction(docId, spaceId, entity.LogDocActionDelete, "删除文档: "+docName)

	return nil
}

// GetDocsByParentId 获取父文档下的子文档
func (d *Doc) GetDocsByParentId(parentId int64) ([]*entity.DocEntity, errors.BizError) {
	return d.daoDoc.GetDocsByParentId(parentId)
}

// MoveDoc 移动文档到目标目录
func (d *Doc) MoveDoc(docId int64, targetId int64) errors.BizError {
	doc, err := d.daoDoc.GetDocByDocId(docId)
	if err != nil {
		return err
	}
	if doc == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "文档不存在")
	}

	target, err := d.daoDoc.GetDocByDocId(targetId)
	if err != nil {
		return err
	}
	if target == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "目标文档不存在")
	}
	if target.Type != entity.DirEntityType {
		return errors.Errorf(errors.ClientReqParamWrongful, "目标文档必须是目录")
	}
	if doc.SpaceId != target.SpaceId {
		return errors.Errorf(errors.ClientReqParamWrongful, "文档和目标文档不在同一空间")
	}

	// 构建新路径
	newPath := target.Path
	if newPath == "" {
		newPath = fmt.Sprintf("%d", targetId)
	} else {
		newPath = fmt.Sprintf("%s,%d", target.Path, targetId)
	}

	accountId := global.ContextValueLoginAccountID(d.ctx)
	accountName := global.ContextValueLoginAccountName(d.ctx)

	return d.daoDoc.MoveDoc(docId, targetId, newPath, accountId, accountName)
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

// logDocAction 记录文档操作日志
func (d *Doc) logDocAction(docId int64, spaceId int64, action int, comment string) {
	logDocEntity := &entity.LogDocEntity{
		DocId:     fmt.Sprintf("%d", docId),
		SpaceId:   spaceId,
		AccountId: global.ContextValueLoginAccountID(d.ctx),
		Action:    action,
		Comment:   comment,
	}
	err := d.daoLogDoc.Insert(logDocEntity)
	if err != nil {
		logger.WithContext(d.ctx).Errorf("[logDocAction] Insert doc log err=%+v", err)
	}
}
