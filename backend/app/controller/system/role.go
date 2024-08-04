package system

import (
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
	"github.com/phachon/mm-wiki/utils"
)

// RoleAdd 添加角色页面
func RoleAdd(ctx *gin.Context) error {
	// 获取所有的权限列表
	allPrivilege, err := service.NewPrivilege(ctx).GetAllPrivilegeList()
	if err != nil {
		logger.WithContext(ctx).Warnf("[RolePrivilegeEdit] GetAllPrivilegeList err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	data := map[string]interface{}{
		"all_privilege": allPrivilege,
	}
	return RespJsonSuccess(ctx, data)
}

// RoleSave 添加角色保存
func RoleSave(ctx *gin.Context) error {

	name := GetParamString(ctx, "name")
	remark := GetParamString(ctx, "remark")
	roleType := GetParamIntDef(ctx, "role_type", 0)
	privilegeIds := GetParamString(ctx, "privilege_ids")
	// 判断参数合法性
	if name == "" {
		logger.WithContext(ctx).Warnf("[RoleSave] name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "角色名不能为空")
	}
	if len(privilegeIds) == 0 {
		logger.WithContext(ctx).Warnf("[RoleSave] privilege_ids empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "权限不能为空")
	}

	// role 角色实体
	roleEntity := &entity.RoleEntity{
		Name:         name,
		Remark:       remark,
		RoleType:     roleType,
		PrivilegeIds: privilegeIds,
	}
	err := service.NewRole(ctx).Create(roleEntity)
	if err != nil {
		sysLogErrorf(ctx, "[RoleAdd] 添加角色失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[RoleAdd] 添加角色 %d 成功", roleEntity.RoleId)

	return RespJsonSuccess(ctx, nil)
}

// RoleEdit 角色修改页面
func RoleEdit(ctx *gin.Context) error {

	roleId := GetParamInt64(ctx, "role_id")
	// 判断参数合法性
	if roleId == 0 {
		logger.WithContext(ctx).Warnf("[RoleEdit] 角色id不存在")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "角色id不存在")
	}

	role, err := service.NewRole(ctx).GetRoleByRoleId(roleId)
	if err != nil {
		sysLogErrorf(ctx, "[RoleEdit] 获取角色 %d 信息失败: err=%+v", roleId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if role == nil {
		logger.WithContext(ctx).Warnf("[RoleEdit] 角色 %d 不存在", roleId)
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "角色id错误")
	}

	return RespJsonSuccess(ctx, map[string]interface{}{
		"role_info": role,
	})
}

// RoleModify 修改角色保存
func RoleModify(ctx *gin.Context) error {

	roleId := GetParamInt64(ctx, "role_id")
	name := GetParamString(ctx, "name")
	remark := GetParamString(ctx, "remark")
	roleType := GetParamIntDef(ctx, "role_type", 0)

	// 判断参数合法性
	if roleId == 0 {
		logger.WithContext(ctx).Warnf("[RoleModify] role_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "角色id不存在")
	}
	if name == "" {
		logger.WithContext(ctx).Warnf("[RoleModify] name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "角色名不能为空")
	}

	// role 角色实体
	roleEntity := entity.RoleEntity{
		RoleId:   roleId,
		Name:     name,
		Remark:   remark,
		RoleType: roleType,
	}
	err := service.NewRole(ctx).Update(roleEntity)
	if err != nil {
		sysLogErrorf(ctx, "[RoleModify] 更新角色 %d 失败: err=%+v", roleId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[RoleModify] 更新角色 %d 成功", roleId)

	return RespJsonSuccess(ctx, nil)
}

// RoleList 角色列表
func RoleList(ctx *gin.Context) error {

	pageSize := GetParamIntDef(ctx, "page_size", 20)
	pageNum := GetParamIntDef(ctx, "page_num", 1)
	keywordsStr := GetParamStringDef(ctx, "keywords", "")

	var keywords *entity.RoleKeywords
	if len(keywordsStr) > 0 {
		jErr := json.Unmarshal([]byte(keywordsStr), &keywords)
		if jErr != nil {
			logger.WithContext(ctx).Errorf("[RoleList] GetRolesByLimit err=%s", jErr.Error())
		}
	}

	serviceRole := service.NewRole(ctx)

	// 获取角色列表
	roles, err := serviceRole.GetRolesByLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[RoleList] 获取角色列表失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取分页信息
	pageInfo, err := serviceRole.GetPageInfoLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[RoleList] 获取角色分页信息失败: err=%s", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 格式化列表结构
	roleList, err := serviceRole.FormatRoleList(roles)
	if err != nil {
		sysLogErrorf(ctx, "[RoleList] 格式化角色列表失败: err=%s", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	data := map[string]interface{}{
		"list":      roleList,
		"page_info": pageInfo,
	}
	return RespJsonSuccess(ctx, data)
}

// RoleAccountList 角色账号列表
func RoleAccountList(ctx *gin.Context) error {

	pageSize := GetParamIntDef(ctx, "page_size", 20)
	pageNum := GetParamIntDef(ctx, "page_num", 1)
	roleId := GetParamInt64Def(ctx, "role_id", 0)
	if roleId <= 0 {
		logger.WithContext(ctx).Warnf("[RoleAccountList] role_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "角色id不存在")
	}

	serviceAccountRole := service.NewAccountRole(ctx)
	accounts, err := serviceAccountRole.GetAccountsByRoleIdLimit(pageSize, pageNum, roleId)
	if err != nil {
		sysLogErrorf(ctx, "[RoleAccountList] 获取角色 %d 下账号列表失败 err=%+v", roleId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取分页信息
	pageInfo, err := serviceAccountRole.GetPageInfoLimitByRoleId(pageSize, pageNum, roleId)
	if err != nil {
		sysLogErrorf(ctx, "[RoleAccountList] 获取角色 %d 下账号分页信息失败: err=%+v", roleId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	data := map[string]interface{}{
		"list":      accounts,
		"page_info": pageInfo,
	}
	return RespJsonSuccess(ctx, data)
}

// RoleAccountRemove 角色移除账号
func RoleAccountRemove(ctx *gin.Context) error {

	roleId := GetParamInt64(ctx, "role_id")
	accountId := GetParamInt64(ctx, "account_id")

	// 判断参数合法性
	if roleId <= 0 {
		logger.WithContext(ctx).Warnf("[RoleAccountRemove] role_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "角色id不存在")
	}
	if accountId == 0 {
		logger.WithContext(ctx).Warnf("[RoleAccountRemove] account_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "账号id不存在")
	}

	// 移除角色和账号的对应关系
	err := service.NewRole(ctx).RemoveRoleAccount(roleId, accountId)
	if err != nil {
		sysLogErrorf(ctx, "[RoleAccountRemove] 移除角色 %d 下账号 %d 失败: err=%+v", roleId, accountId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[RoleAccountRemove] 移除角色 %d 下账号 %d 成功", roleId, accountId)

	return RespJsonSuccess(ctx, nil)
}

// RolePrivilegeEdit 角色权限编辑
func RolePrivilegeEdit(ctx *gin.Context) error {

	roleId := GetParamInt64(ctx, "role_id")

	// 判断参数合法性
	if roleId <= 0 {
		logger.WithContext(ctx).Warnf("[RolePrivilegeEdit] role_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "角色id为空")
	}
	// 获取该角色信息
	role, err := service.NewRole(ctx).GetRoleByRoleId(roleId)
	if err != nil {
		sysLogErrorf(ctx, "[RolePrivilegeEdit] 获取角色 %d 权限失败: err=%+v", roleId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if role == nil {
		logger.WithContext(ctx).Warnf("[RolePrivilegeEdit] role_id=%d not exists", roleId)
		return RespJsonError(ctx, int32(errors.BusinessRecordNotExistError), "角色id不存在")
	}

	// 获取所有的权限列表
	allPrivilege, err := service.NewPrivilege(ctx).GetAllPrivilegeList()
	if err != nil {
		logger.WithContext(ctx).Warnf("[RolePrivilegeEdit] GetAllPrivilegeList err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	privilegeIdsStr := strings.Split(role.PrivilegeIds, ",")

	// 获取角色对应权限
	var privilegeIds = []int64{}
	for _, privilegeId := range privilegeIdsStr {
		privilegeIds = append(privilegeIds, utils.Convert.StringToInt64(privilegeId))
	}

	data := map[string]interface{}{
		"all_privilege": allPrivilege,
		"privilege_ids": privilegeIds,
	}
	return RespJsonSuccess(ctx, data)
}

// RolePrivilegeModify 角色权限更新保存
func RolePrivilegeModify(ctx *gin.Context) error {

	roleId := GetParamInt64(ctx, "role_id")
	privilegeIds := GetParamString(ctx, "privilege_ids")

	// 判断参数合法性
	if roleId <= 0 {
		logger.WithContext(ctx).Warnf("[RolePrivilegeModify] role_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "角色id不存在")
	}
	if len(privilegeIds) == 0 {
		logger.WithContext(ctx).Warnf("[RolePrivilegeModify] privilege_ids empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "权限id不存在")
	}

	// 修改保存角色权限
	err := service.NewRole(ctx).UpdatePrivilegeIds(roleId, privilegeIds)
	if err != nil {
		sysLogErrorf(ctx, "[RolePrivilegeModify] 更新角色 %d 权限失败: err=%+v", roleId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[RolePrivilegeModify] 更新角色 %d 权限成功", roleId)

	return RespJsonSuccess(ctx, nil)
}

// RoleDelete 角色删除
func RoleDelete(ctx *gin.Context) error {

	roleId := GetParamInt64(ctx, "role_id")

	// 判断参数合法性
	if roleId <= 0 {
		logger.WithContext(ctx).Warnf("[RoleDelete] role_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "角色id不存在")
	}

	// 删除角色
	err := service.NewRole(ctx).DeleteRole(roleId)
	if err != nil {
		sysLogErrorf(ctx, "[RoleDelete] 删除角色 %d 失败: err=%+v", roleId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[RoleDelete] 删除角色 %d 成功", roleId)

	return RespJsonSuccess(ctx, nil)
}
