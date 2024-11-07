package space

import (
	"encoding/json"

	"github.com/gin-gonic/gin"

	"github.com/phachon/mm-wiki/app/controller"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/logger"
)

// GetSpaces 获取空间列表
func GetSpaces(ctx *gin.Context) error {
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
