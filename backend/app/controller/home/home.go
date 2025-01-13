package home

import (
	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/controller"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/logger"
	"github.com/phachon/mm-wiki/utils"
)

// GetMySpaces 获取我的空间
func GetMySpaces(ctx *gin.Context) error {

	accountId := global.ContextValueLoginAccountID(ctx)

	spacePermissionService := service.NewSpacePermission(ctx)

	spaces, err := spacePermissionService.GetSpacesByAdminId(accountId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[GetMySpaces] GetAdminsBySpaceId err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "获取账号空间失败")
	}
	return controller.RespJsonSuccess(ctx, map[string]interface{}{
		"list": spaces,
	})
}

// GetCollectionSpaces 获取收藏空间
func GetCollectionSpaces(ctx *gin.Context) error {

	accountId := global.ContextValueLoginAccountID(ctx)

	collectionService := service.NewCollection(ctx)
	collections, err := collectionService.GetAccountCollectionAllSpace(accountId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[GetCollectionSpaces] GetAccountCollectionAllSpace err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "获取收藏空间失败")
	}
	var spaceIds []int64
	for _, collection := range collections {
		spaceIds = append(spaceIds, utils.Convert.StringToInt64(collection.ResourceId))
	}

	// 获取空间信息
	spaceService := service.NewSpace(ctx)
	spaces, err := spaceService.GetSpacesBySpaceIds(spaceIds)
	if err != nil {
		logger.WithContext(ctx).Errorf("[GetCollectionSpaces] GetSpacesBySpaceIds err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "获取收藏空间失败")
	}

	return controller.RespJsonSuccess(ctx, map[string]interface{}{
		"list": spaces,
	})
}

// GetCollectionDocs 获取收藏文档
func GetCollectionDocs(ctx *gin.Context) error {

	accountId := global.ContextValueLoginAccountID(ctx)

	collectionService := service.NewCollection(ctx)
	collections, err := collectionService.GetAccountCollectionAllDocs(accountId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[GetCollectionDocs] GetAccountCollectionAllDocument err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "获取收藏文档失败")
	}

	var docIds []int64
	for _, collection := range collections {
		docIds = append(docIds, utils.Convert.StringToInt64(collection.ResourceId))
	}

	// 获取文档信息
	docService := service.NewDoc(ctx)
	docs, err := docService.GetDocByDocIds(docIds)
	if err != nil {
		logger.WithContext(ctx).Errorf("[GetCollectionDocs] GetDocumentsByDocumentIds err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "获取收藏文档失败")
	}

	return controller.RespJsonSuccess(ctx, map[string]interface{}{
		"list": docs,
	})
}
