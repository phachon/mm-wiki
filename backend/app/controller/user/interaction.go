package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/controller"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
)

// CollectionList 收藏状态
func CollectionStatus(ctx *gin.Context) error {
	resourceId := controller.GetParamStringDef(ctx, "resource_id", "")
	collectionType := controller.GetParamIntDef(ctx, "collection_type", 0)

	if resourceId == "" {
		logger.WithContext(ctx).Warnf("[DocCollectionStatus] 资源ID不能为空")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "资源ID不能为空")
	}
	if collectionType != entity.CollectionTypeDocument && collectionType != entity.CollectionTypeSpace {
		logger.WithContext(ctx).Warnf("[DocCollectionStatus] 收藏类型不合法")
	}
	accountId := global.ContextValueLoginAccountID(ctx)

	serviceCollection := service.NewCollection(ctx)
	collectionInfo, err := serviceCollection.GetCollectionByAccountIdAndTypeAndResourceId(accountId, collectionType, resourceId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocCollectionStatus] GetCollectionByAccountIdResourceAndType err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "获取收藏状态失败")
	}
	status := 0 // 未收藏
	if collectionInfo != nil {
		status = 1 // 已收藏
	}

	return controller.RespJsonSuccess(ctx, map[string]interface{}{
		"collection_status": status, // 是否收藏
	})
}

// CollectionAdd 添加收藏操作
func CollectionAdd(ctx *gin.Context) error {
	resourceId := controller.GetParamInt64Def(ctx, "resource_id", 0)
	collectionType := controller.GetParamIntDef(ctx, "collection_type", 0)

	if resourceId <= 0 {
		logger.WithContext(ctx).Warnf("[DocCollectionAction] 资源ID不能为空")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "资源ID不能为空")
	}
	if collectionType != entity.CollectionTypeDocument && collectionType != entity.CollectionTypeSpace {
		logger.WithContext(ctx).Warnf("[DocCollectionAction] 收藏类型不合法")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "收藏类型不合法")
	}

	// 收藏文档
	if collectionType == entity.CollectionTypeDocument {
		serviceDoc := service.NewDoc(ctx)
		doc, err := serviceDoc.GetDocByDocId(resourceId)
		if err != nil {
			logger.WithContext(ctx).Errorf("[DocCollectionAction] GetDocById err=%+v", err)
			return controller.RespJsonError(ctx, err.GetErrCode(), "获取文档信息失败")
		}
		if doc == nil {
			logger.WithContext(ctx).Warnf("[DocCollectionAction] 文档不存在")
			return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "文档不存在")
		}
	}

	// 收藏空间
	if collectionType == entity.CollectionTypeSpace {
		serviceSpace := service.NewSpace(ctx)
		space, err := serviceSpace.GetSpaceBySpaceId(resourceId)
		if err != nil {
			logger.WithContext(ctx).Errorf("[DocCollectionAction] GetSpaceBySpaceId err=%+v", err)
			return controller.RespJsonError(ctx, err.GetErrCode(), "获取空间信息失败")
		}
		if space == nil {
			logger.WithContext(ctx).Warnf("[DocCollectionAction] 空间不存在")
			return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "空间不存在")
		}
	}

	serviceCollection := service.NewCollection(ctx)
	collectionInfo := &entity.CollectionEntity{
		AccountId:      global.ContextValueLoginAccountID(ctx),
		CollectionType: collectionType,
		ResourceId:     fmt.Sprintf("%d", resourceId),
	}
	err := serviceCollection.Create(collectionInfo)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocCollectionAction] CreateCollection err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "收藏失败")
	}
	return controller.RespJsonSuccess(ctx, map[string]interface{}{})
}

// CollectionCancel 取消收藏操作
func CollectionCancel(ctx *gin.Context) error {
	resourceId := controller.GetParamStringDef(ctx, "resource_id", "")
	collectionType := controller.GetParamIntDef(ctx, "collection_type", 0)

	if resourceId == "" {
		logger.WithContext(ctx).Warnf("[DocCollectionCancel] 资源id不能为空")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "资源ID不能为空")
	}
	if collectionType != entity.CollectionTypeDocument && collectionType != entity.CollectionTypeSpace {
		logger.WithContext(ctx).Warnf("[DocCollectionCancel] 收藏类型不合法")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "收藏类型不合法")
	}

	accountId := global.ContextValueLoginAccountID(ctx)

	serviceCollection := service.NewCollection(ctx)
	err := serviceCollection.DeleteByAccountIdResourceAndType(accountId, collectionType, resourceId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[DocCollectionCancel] DeleteCollection err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "取消收藏失败")
	}

	return controller.RespJsonSuccess(ctx, map[string]interface{}{})
}
