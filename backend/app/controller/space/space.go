package space

import (
	"encoding/json"

	"github.com/gin-gonic/gin"

	"github.com/phachon/mm-wiki/app/controller"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
	"github.com/phachon/mm-wiki/utils"
)

// AllSpaces 所有的空间列表
func AllSpaces(ctx *gin.Context) error {
	pageSize := controller.GetParamIntDef(ctx, "page_size", 12)
	pageNum := controller.GetParamIntDef(ctx, "page_num", 1)
	keywordsStr := controller.GetParamStringDef(ctx, "keywords", "")

	var keywords *entity.SpaceKeywords
	if len(keywordsStr) > 0 {
		jErr := json.Unmarshal([]byte(keywordsStr), &keywords)
		if jErr != nil {
			logger.WithContext(ctx).Errorf("[SpaceList] GetSpacesByLimit err=%s", jErr.Error())
		}
	}

	serviceSpace := service.NewSpace(ctx)

	// 获取公开的空间列表
	spaces, err := serviceSpace.GetPublicSpacesByLimit(pageSize, pageNum, keywords)
	if err != nil {
		logger.WithContext(ctx).Errorf("[SpaceList] GetPublicSpacesByLimit err=%s", err.Error())
		return controller.RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取公开空间的分页信息
	pageInfo, err := serviceSpace.GetPublicSpacesPageInfo(pageSize, pageNum, keywords)
	if err != nil {
		logger.WithContext(ctx).Errorf("[SpaceList] GetPublicSpacesPageInfo err=%s", err.Error())
		return controller.RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	// 获取账号收藏的所有空间ID
	accountId := global.ContextValueLoginAccountID(ctx)
	serviceCollection := service.NewCollection(ctx)
	collections, err := serviceCollection.GetAccountCollectionAllSpace(accountId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[SpaceList] GetAccountCollectionAllSpace err=%s", err.Error())
		return controller.RespJsonError(ctx, err.GetErrCode(), "获取收藏空间失败")
	}
	// 获取空间ID
	spaceIds := make([]int64, 0)
	for _, collection := range collections {
		if collection == nil || collection.ResourceId == "" {
			continue
		}
		spaceId := utils.Convert.StringToInt64(collection.ResourceId)
		spaceIds = append(spaceIds, spaceId)
	}

	data := map[string]interface{}{
		"collection_ids": spaceIds,
		"list":           spaces,
		"page_info":      pageInfo,
	}
	return controller.RespJsonSuccess(ctx, data)
}

// SpaceDocs 空间文档
func SpaceDocs(ctx *gin.Context) error {
	spaceKey := controller.GetParamStringDef(ctx, "space_key", "")
	if len(spaceKey) == 0 {
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "空间 Key 不能为空")
	}

	// 获取空间信息
	serviceSpace := service.NewSpace(ctx)
	space, err := serviceSpace.GetSpaceByKey(spaceKey)
	if err != nil {
		logger.WithContext(ctx).Errorf("[SpaceInfo] GetSpaceInfo err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	// 获取空间下所有文档
	serviceDoc := service.NewDoc(ctx)

	docs, err := serviceDoc.GetDocsBySpaceKey(spaceKey)
	if err != nil {
		logger.WithContext(ctx).Errorf("[SpaceInfo] GetDocsBySpaceKey err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	logger.WithContext(ctx).Infof("[SpaceInfo] docs=%+v", docs)

	// 获取主页文档，循环 docs 获取 parent_id 为 0 的文档
	var homeDoc *entity.DocEntity   // 主页文档
	var dirDocs []*entity.DocEntity // 目录文档
	for _, doc := range docs {
		if doc.ParentId == 0 {
			homeDoc = doc
			continue
		}
		dirDocs = append(dirDocs, doc)
	}
	if homeDoc == nil {
		logger.WithContext(ctx).Errorf("[SpaceInfo] 空间主页数据异常 spaceKey=%s", spaceKey)
		return controller.RespJsonError(ctx, int32(errors.BusinessRecordNotExistError), "空间主页文档异常")
	}
	docTree := serviceDoc.DocsToTree(dirDocs, homeDoc.DocId)
	data := map[string]interface{}{
		"home_doc":   homeDoc,
		"doc_tree":   docTree,
		"space_info": space,
	}

	return controller.RespJsonSuccess(ctx, data)
}
