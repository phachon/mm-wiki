package space

import (
	"encoding/json"

	"github.com/gin-gonic/gin"

	"github.com/phachon/mm-wiki/app/controller"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
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
	data := map[string]interface{}{
		"list":      spaces,
		"page_info": pageInfo,
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

	docTree := serviceDoc.DocsToTree(docs, 0)

	data := map[string]interface{}{
		"docs": docTree,
		"info": space,
	}

	return controller.RespJsonSuccess(ctx, data)
}
