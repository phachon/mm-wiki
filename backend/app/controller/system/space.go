package system

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
	"github.com/phachon/mm-wiki/utils"
)

// SpaceAdd 添加空间页面
func SpaceAdd(ctx *gin.Context) error {
	// 获取所有的账号
	accountList, err := service.NewAccount(ctx).GetAllNormalAccounts()
	if err != nil {
		sysLogErrorf(ctx, "[SpaceAdd] 获取所有账号失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	data := map[string]interface{}{
		"account_list": accountList,
	}
	return RespJsonSuccess(ctx, data)
}

// SpaceSave 添加空间保存
func SpaceSave(ctx *gin.Context) error {

	spaceKey := GetParamString(ctx, "space_key")
	name := GetParamString(ctx, "name")
	description := GetParamString(ctx, "description")
	spaceType := GetParamIntDef(ctx, "space_type", 0)
	visitLevel := GetParamIntDef(ctx, "visit_level", 0)
	isShare := GetParamIntDef(ctx, "is_share", 0)
	isExport := GetParamIntDef(ctx, "is_export", 0)
	adminAccountIds := GetParamArray(ctx, "admin_account_ids")

	// 判断参数合法性
	if name == "" {
		logger.WithContext(ctx).Warnf("[SpaceSave] name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "空间名不能为空")
	}
	if description == "" {
		logger.WithContext(ctx).Warnf("[SpaceSave] description empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "空间描述不能为空")
	}
	if spaceType != entity.SpaceTypeTeam && spaceType != entity.SpaceTypeDefaultPersonal {
		logger.WithContext(ctx).Warnf("[SpaceSave] space_type error")
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "空间类型错误")
	}
	if visitLevel != entity.SpaceVisitLevelDefaultPublic && visitLevel != entity.SpaceVisitLevelDefaultPrivate {
		logger.WithContext(ctx).Warnf("[SpaceSave] visit_level error")
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "访问级别错误")
	}
	if isShare != entity.SpaceIsShareDefault && isShare != entity.SpaceIsShareYes {
		logger.WithContext(ctx).Warnf("[SpaceSave] is_share error")
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "是否分享错误")
	}
	if isExport != entity.SpaceIsExportDefault && isExport != entity.SpaceIsExportYes {
		logger.WithContext(ctx).Warnf("[SpaceSave] is_export error")
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "是否导出错误")
	}
	if len(adminAccountIds) == 0 {
		logger.WithContext(ctx).Warnf("[SpaceSave] adminIds empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "请选择空间管理员")
	}

	// space 空间实体
	spaceEntity := &entity.SpaceEntity{
		SpaceKey:         spaceKey,
		Name:             name,
		Description:      description,
		SpaceType:        spaceType,
		VisitLevel:       visitLevel,
		IsShare:          isShare,
		IsExport:         isExport,
		CreatorAccountId: global.ContextValueLoginAccountID(ctx),
		CreatorName:      global.ContextValueLoginAccountName(ctx),
	}
	err := service.NewSpace(ctx).Create(spaceEntity)
	if err != nil {
		sysLogErrorf(ctx, "[SpaceAdd] 添加空间失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// adminAccountIds 转换为 int64 数组
	accountIds := utils.Convert.StringsToInt64(adminAccountIds)
	err = service.NewSpacePermission(ctx).CreateBatchAdminPerssions(spaceEntity.SpaceId, accountIds)
	if err != nil {
		sysLogErrorf(ctx, "[SpaceAdd] 添加空间管理员权限失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[SpaceAdd] 添加空间 %d 成功", spaceEntity.SpaceId)

	return RespJsonSuccess(ctx, nil)
}

// SpaceEdit 空间修改页面
func SpaceEdit(ctx *gin.Context) error {

	spaceId := GetParamInt64(ctx, "space_id")
	// 判断参数合法性
	if spaceId == 0 {
		logger.WithContext(ctx).Warnf("[SpaceEdit] 空间id不存在")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "空间id不存在")
	}

	space, err := service.NewSpace(ctx).GetSpaceBySpaceId(spaceId)
	if err != nil {
		sysLogErrorf(ctx, "[SpaceEdit] 获取空间 %d 信息失败: err=%+v", spaceId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if space == nil {
		logger.WithContext(ctx).Warnf("[SpaceEdit] 空间 %d 不存在", spaceId)
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "空间id错误")
	}

	return RespJsonSuccess(ctx, map[string]interface{}{
		"space_info": space,
	})
}

// SpaceModify 修改空间保存
func SpaceModify(ctx *gin.Context) error {

	spaceId := GetParamInt64(ctx, "space_id")
	name := GetParamString(ctx, "name")
	description := GetParamString(ctx, "description")
	visitLevel := GetParamIntDef(ctx, "visit_level", 0)
	isShare := GetParamIntDef(ctx, "is_share", 0)
	isExport := GetParamIntDef(ctx, "is_export", 0)

	// 判断参数合法性
	if spaceId == 0 {
		logger.WithContext(ctx).Warnf("[SpaceModify] space_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "空间id不存在")
	}
	if name == "" {
		logger.WithContext(ctx).Warnf("[SpaceModify] name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "空间名不能为空")
	}
	if description == "" {
		logger.WithContext(ctx).Warnf("[SpaceModify] description empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "空间描述不能为空")
	}
	if visitLevel != entity.SpaceVisitLevelDefaultPublic && visitLevel != entity.SpaceVisitLevelDefaultPrivate {
		logger.WithContext(ctx).Warnf("[SpaceModify] visit_level error")
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "访问级别错误")
	}
	if isShare != entity.SpaceIsShareDefault && isShare != entity.SpaceIsShareYes {
		logger.WithContext(ctx).Warnf("[SpaceModify] is_share error")
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "是否分享错误")
	}
	if isExport != entity.SpaceIsExportDefault && isExport != entity.SpaceIsExportYes {
		logger.WithContext(ctx).Warnf("[SpaceModify] is_export error")
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "是否导出错误")
	}

	// space 空间实体
	spaceEntity := entity.SpaceEntity{
		SpaceId:     spaceId,
		Name:        name,
		Description: description,
		VisitLevel:  visitLevel,
		IsShare:     isShare,
		IsExport:    isExport,
	}
	err := service.NewSpace(ctx).Update(spaceEntity)
	if err != nil {
		sysLogErrorf(ctx, "[SpaceModify] 更新空间 %d 失败: err=%+v", spaceId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[SpaceModify] 更新空间 %d 成功", spaceId)

	return RespJsonSuccess(ctx, nil)
}

// SpaceList 空间列表
func SpaceList(ctx *gin.Context) error {

	pageSize := GetParamIntDef(ctx, "page_size", 20)
	pageNum := GetParamIntDef(ctx, "page_num", 1)
	keywordsStr := GetParamStringDef(ctx, "keywords", "")

	var keywords *entity.SpaceKeywords
	if len(keywordsStr) > 0 {
		jErr := json.Unmarshal([]byte(keywordsStr), &keywords)
		if jErr != nil {
			logger.WithContext(ctx).Errorf("[SpaceList] GetSpacesByLimit err=%s", jErr.Error())
		}
	}

	serviceSpace := service.NewSpace(ctx)

	// 获取空间列表
	spaces, err := serviceSpace.GetSpacesByLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[SpaceList] 获取空间列表失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取分页信息
	pageInfo, err := serviceSpace.GetPageInfoLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[SpaceList] 获取空间分页信息失败: err=%s", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 格式化列表结构
	spaceList, err := serviceSpace.FormatSpaceList(spaces)
	if err != nil {
		sysLogErrorf(ctx, "[SpaceList] 格式化空间列表失败: err=%s", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	data := map[string]interface{}{
		"list":      spaceList,
		"page_info": pageInfo,
	}
	return RespJsonSuccess(ctx, data)
}

// SpaceDelete 空间删除
func SpaceDelete(ctx *gin.Context) error {

	spaceId := GetParamInt64(ctx, "space_id")

	// 判断参数合法性
	if spaceId <= 0 {
		logger.WithContext(ctx).Warnf("[SpaceDelete] space_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "空间id不存在")
	}

	// 删除空间
	err := service.NewSpace(ctx).DeleteSpace(spaceId)
	if err != nil {
		sysLogErrorf(ctx, "[SpaceDelete] 删除空间 %d 失败: err=%+v", spaceId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[SpaceDelete] 删除空间 %d 成功", spaceId)

	return RespJsonSuccess(ctx, nil)
}

// SpaceAdminList 空间管理员列表
func SpaceAdminList(ctx *gin.Context) error {

	spaceId := GetParamInt64(ctx, "space_id")

	// 判断参数合法性
	if spaceId <= 0 {
		logger.WithContext(ctx).Warnf("[SpaceAdminList] space_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "空间id不存在")
	}

	// 获取空间管理员
	admins, err := service.NewSpacePermission(ctx).GetAdminsBySpaceId(spaceId)
	if err != nil {
		sysLogErrorf(ctx, "[SpaceAdminList] 获取空间 %d 管理员失败: err=%+v", spaceId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	// 获取所有的账号
	accountList, err := service.NewAccount(ctx).GetAllNormalAccounts()
	if err != nil {
		sysLogErrorf(ctx, "[SpaceAdminList] 获取所有账号失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	// 找到可以添加的管理员
	var adminMap = make(map[int64]bool)
	for _, admin := range admins {
		adminMap[admin.AccountId] = true
	}
	var selectedList []*entity.AccountEntity
	for _, account := range accountList {
		if _, ok := adminMap[account.AccountId]; !ok {
			selectedList = append(selectedList, account)
		}
	}

	return RespJsonSuccess(ctx, map[string]interface{}{
		"admin_list":    admins,
		"selected_list": selectedList,
	})
}

// SpaceAdminRemove 移除空间管理员
func SpaceAdminRemove(ctx *gin.Context) error {

	spaceId := GetParamInt64(ctx, "space_id")
	accountId := GetParamInt64(ctx, "account_id")

	// 判断参数合法性
	if spaceId <= 0 {
		logger.WithContext(ctx).Warnf("[SpaceAdminRemove] space_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "空间id不存在")
	}
	if accountId <= 0 {
		logger.WithContext(ctx).Warnf("[SpaceAdminRemove] account_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "账号id不存在")
	}

	// 删除空间管理员
	err := service.NewSpacePermission(ctx).DeleteBySpaceIdAccountId(spaceId, accountId)
	if err != nil {
		sysLogErrorf(ctx, "[SpaceAdminRemove] 移除空间 %d 管理员 %d 失败: err=%+v", spaceId, accountId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[SpaceAdminRemove] 移除空间 %d 管理员 %d 成功", spaceId, accountId)

	return RespJsonSuccess(ctx, nil)
}

// SpaceAdminAdd 添加空间管理员
func SpaceAdminAdd(ctx *gin.Context) error {

	spaceId := GetParamInt64(ctx, "space_id")
	adminAccountIds := GetParamArray(ctx, "admin_account_ids")

	// 判断参数合法性
	if spaceId <= 0 {
		logger.WithContext(ctx).Warnf("[SpaceAdminAdd] space_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "空间id不存在")
	}
	if len(adminAccountIds) == 0 {
		logger.WithContext(ctx).Warnf("[SpaceAdminAdd] adminIds empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "请选择空间管理员")
	}

	// adminAccountIds 转换为 int64 数组
	accountIds := utils.Convert.StringsToInt64(adminAccountIds)
	err := service.NewSpacePermission(ctx).CreateBatchAdminPerssions(spaceId, accountIds)
	if err != nil {
		sysLogErrorf(ctx, "[SpaceAdminAdd] 添加空间管理员失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[SpaceAdminAdd] 添加空间 %d 管理员成功", spaceId)
	return RespJsonSuccess(ctx, nil)
}
