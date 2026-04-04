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

// SpaceBasicSettingModify 空间基本设置修改
func SpaceBasicSettingModify(ctx *gin.Context) error {
	spaceId := controller.GetParamInt64(ctx, "space_id")
	name := controller.GetParamString(ctx, "name")
	description := controller.GetParamString(ctx, "description")
	visitLevel := controller.GetParamIntDef(ctx, "visit_level", 0)

	// 判断参数合法性
	if spaceId == 0 {
		logger.WithContext(ctx).Warnf("[SpaceBasicSettingModify] space_id empty")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "空间id不存在")
	}
	if name == "" {
		logger.WithContext(ctx).Warnf("[SpaceBasicSettingModify] name empty")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "空间名不能为空")
	}
	if description == "" {
		logger.WithContext(ctx).Warnf("[SpaceBasicSettingModify] description empty")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "空间描述不能为空")
	}
	if visitLevel != entity.SpaceVisitLevelDefaultPublic && visitLevel != entity.SpaceVisitLevelDefaultPrivate {
		logger.WithContext(ctx).Warnf("[SpaceBasicSettingModify] visit_level error")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "访问级别错误")
	}
	// space 空间实体
	spaceEntity := entity.SpaceEntity{
		SpaceId:     spaceId,
		Name:        name,
		Description: &description,
		VisitLevel:  &visitLevel,
	}
	err := service.NewSpace(ctx).Update(spaceEntity)
	if err != nil {
		return controller.RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	return controller.RespJsonSuccess(ctx, nil)
}

// SpacePermissionList 空间权限列表
func SpacePermissionList(ctx *gin.Context) error {
	spaceId := controller.GetParamInt64(ctx, "space_id")
	if spaceId <= 0 {
		logger.WithContext(ctx).Warnf("[SpacePermissionList] space_id empty")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "空间 id 不能为空")
	}
	// 获取空间信息
	serviceSpace := service.NewSpace(ctx)
	space, err := serviceSpace.GetSpaceBySpaceId(spaceId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[SpacePermissionList] GetSpaceBySpaceId err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "获取空间信息失败")
	}
	// 获取空间权限列表
	permissions, err := service.NewSpacePermission(ctx).GetPermissionsBySpaceId(space.SpaceId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[SpacePermissionList] GetPermissionsBySpaceId err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "获取空间权限列表失败")
	}

	var (
		departmentList []*entity.SpacePermissionEntity
		accountList    []*entity.SpacePermissionEntity
		adminList      []*entity.SpacePermissionEntity
	)
	for _, permission := range permissions {
		if permission.PermissionType == entity.SpacePermissionRelationTypeDepartment {
			departmentList = append(departmentList, permission)
		}
		if permission.PermissionType == entity.SpacePermissionRelationTypeAccount {
			accountList = append(accountList, permission)
		}
		if permission.PermissionType == entity.SpacePermissionRelationTypeAdmin {
			adminList = append(adminList, permission)
		}
	}
	data := map[string]interface{}{
		"admin_list":      adminList,
		"department_list": departmentList,
		"account_list":    accountList,
	}
	return controller.RespJsonSuccess(ctx, data)
}

// SpacePermissionAdd 添加空间权限
func SpacePermissionAdd(ctx *gin.Context) error {
	spaceId := controller.GetParamInt64(ctx, "space_id")
	permissionType := controller.GetParamIntDef(ctx, "permission_type", -1)
	accountId := controller.GetParamInt64Def(ctx, "account_id", 0)
	departmentId := controller.GetParamInt64Def(ctx, "department_id", 0)
	isView := controller.GetParamIntDef(ctx, "is_view", 1)
	isAdd := controller.GetParamIntDef(ctx, "is_add", 0)
	isEdit := controller.GetParamIntDef(ctx, "is_edit", 0)
	isDelete := controller.GetParamIntDef(ctx, "is_delete", 0)
	isExport := controller.GetParamIntDef(ctx, "is_export", 0)

	if spaceId <= 0 {
		logger.WithContext(ctx).Warnf("[SpacePermissionAdd] space_id empty")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "空间 id 不能为空")
	}
	if permissionType != entity.SpacePermissionRelationTypeAccount && permissionType != entity.SpacePermissionRelationTypeDepartment {
		logger.WithContext(ctx).Warnf("[SpacePermissionAdd] permission_type invalid")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "权限类型不合法")
	}

	// 验证空间存在
	serviceSpace := service.NewSpace(ctx)
	space, err := serviceSpace.GetSpaceBySpaceId(spaceId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[SpacePermissionAdd] GetSpaceBySpaceId err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "获取空间信息失败")
	}
	if space == nil {
		logger.WithContext(ctx).Warnf("[SpacePermissionAdd] 空间不存在")
		return controller.RespJsonError(ctx, int32(errors.BusinessRecordNotExistError), "空间不存在")
	}

	permission := &entity.SpacePermissionEntity{
		SpaceId:        spaceId,
		PermissionType: permissionType,
		AccountId:      accountId,
		DepartmentId:   departmentId,
		IsView:         &isView,
		IsAdd:          &isAdd,
		IsEdit:         &isEdit,
		IsDelete:       &isDelete,
		IsExport:       &isExport,
	}

	servicePermission := service.NewSpacePermission(ctx)
	err = servicePermission.Create(permission)
	if err != nil {
		logger.WithContext(ctx).Errorf("[SpacePermissionAdd] Create err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	return controller.RespJsonSuccess(ctx, map[string]interface{}{})
}

// SpacePermissionRemove 删除空间权限
func SpacePermissionRemove(ctx *gin.Context) error {
	spaceId := controller.GetParamInt64(ctx, "space_id")
	accountId := controller.GetParamInt64(ctx, "account_id")

	if spaceId <= 0 {
		logger.WithContext(ctx).Warnf("[SpacePermissionRemove] space_id empty")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "空间 id 不能为空")
	}
	if accountId <= 0 {
		logger.WithContext(ctx).Warnf("[SpacePermissionRemove] account_id empty")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "账号 id 不能为空")
	}

	servicePermission := service.NewSpacePermission(ctx)
	err := servicePermission.DeleteBySpaceIdAccountId(spaceId, accountId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[SpacePermissionRemove] DeleteBySpaceIdAccountId err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "删除空间权限失败")
	}

	return controller.RespJsonSuccess(ctx, map[string]interface{}{})
}

// SpacePermissionModify 修改空间权限
func SpacePermissionModify(ctx *gin.Context) error {
	spacePermissionId := controller.GetParamInt64(ctx, "space_permission_id")
	isView := controller.GetParamIntDef(ctx, "is_view", 0)
	isAdd := controller.GetParamIntDef(ctx, "is_add", 0)
	isEdit := controller.GetParamIntDef(ctx, "is_edit", 0)
	isDelete := controller.GetParamIntDef(ctx, "is_delete", 0)
	isExport := controller.GetParamIntDef(ctx, "is_export", 0)

	if spacePermissionId <= 0 {
		logger.WithContext(ctx).Warnf("[SpacePermissionModify] space_permission_id empty")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "权限 id 不能为空")
	}

	permission := &entity.SpacePermission{
		IsView:   isView,
		IsAdd:    isAdd,
		IsEdit:   isEdit,
		IsDelete: isDelete,
		IsExport: isExport,
	}

	servicePermission := service.NewSpacePermission(ctx)
	err := servicePermission.UpdateAccountPermission(spacePermissionId, permission)
	if err != nil {
		logger.WithContext(ctx).Errorf("[SpacePermissionModify] UpdateAccountPermission err=%+v", err)
		return controller.RespJsonError(ctx, err.GetErrCode(), "修改空间权限失败")
	}

	return controller.RespJsonSuccess(ctx, map[string]interface{}{})
}
