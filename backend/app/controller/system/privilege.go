package system

import (
	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
)

// PrivilegeAdd 添加权限页面
func PrivilegeAdd(ctx *gin.Context) error {
	// 获取权限列表（非控制器的权限）
	servicePrivilege := service.NewPrivilege(ctx)
	privilegeList, err := servicePrivilege.GetNavAndMenuPrivileges()
	if err != nil {
		sysLogErrorf(ctx, "[PrivilegeAdd] 获取权限列表失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	data := map[string]interface{}{
		"parent_privileges": privilegeList,
		"api_marks":         service.GetAllApiMarks(), // 获取所有的接口标识
	}
	return RespJsonSuccess(ctx, data)
}

// PrivilegeSave 添加权限保存
func PrivilegeSave(ctx *gin.Context) error {

	identify := GetParamString(ctx, "identify")
	name := GetParamString(ctx, "name")
	parentIds := GetParamStringDef(ctx, "parent_ids", "")
	privilegeType := GetParamInt(ctx, "privilege_type")
	pageRouter := GetParamString(ctx, "page_router")
	apiMarks := GetParamString(ctx, "api_marks")
	icon := GetParamString(ctx, "icon")
	isDisplay := GetParamInt(ctx, "is_display")
	sequence := GetParamIntDef(ctx, "sequence", 1)

	// 判断参数合法性
	if identify == "" {
		logger.WithContext(ctx).Warnf("[PrivilegeSave] identify empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "权限标识不能为空")
	}
	if name == "" {
		logger.WithContext(ctx).Warnf("[PrivilegeSave] name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "权限名不能为空")
	}
	if privilegeType == 0 {
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "权限类型错误")
	}

	// privilege 权限实体
	privilegeEntity := &entity.PrivilegeEntity{
		Identify:      identify,
		Name:          name,
		ParentIds:     parentIds,
		PrivilegeType: privilegeType,
		PageRouter:    pageRouter,
		ApiMarks:      apiMarks,
		Icon:          icon,
		IsDisplay:     isDisplay,
		Sequence:      sequence,
	}
	// 添加权限
	err := service.NewPrivilege(ctx).Create(privilegeEntity)
	if err != nil {
		sysLogErrorf(ctx, "[PrivilegeSave] 添加权限失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[PrivilegeSave] 添加权限 %+v 成功", privilegeEntity.PrivilegeId)

	return RespJsonSuccess(ctx, nil)
}

// PrivilegeEdit 修改编辑权限页面
func PrivilegeEdit(ctx *gin.Context) error {

	privilegeId := GetParamInt64(ctx, "privilege_id")
	// 判断参数合法性
	if privilegeId <= 0 {
		logger.WithContext(ctx).Warnf("[PrivilegeEdit] privilege_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "参数错误")
	}
	servicePrivilege := service.NewPrivilege(ctx)

	// 查找权限是否存在
	privilege, err := servicePrivilege.GetPrivilegeByPrivilegeId(privilegeId)
	if err != nil {
		return err
	}
	if privilege == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "权限不存在")
	}

	// 获取权限列表（非控制器的权限）
	privilegeList, err := servicePrivilege.GetNavAndMenuPrivileges()
	if err != nil {
		sysLogErrorf(ctx, "[PrivilegeEdit] 获取权限列表失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	data := map[string]interface{}{
		"privilege_info":    privilege,
		"parent_privileges": privilegeList,
		"api_marks":         service.GetAllApiMarks(), // 获取所有的接口标识
	}
	return RespJsonSuccess(ctx, data)
}

// PrivilegeModify 修改权限保存
func PrivilegeModify(ctx *gin.Context) error {

	privilegeId := GetParamInt(ctx, "privilege_id")
	name := GetParamString(ctx, "name")
	parentIds := GetParamStringDef(ctx, "parent_ids", "")
	privilegeType := GetParamInt(ctx, "privilege_type")
	pageRouter := GetParamString(ctx, "page_router")
	apiMarks := GetParamString(ctx, "api_marks")
	icon := GetParamString(ctx, "icon")
	isDisplay := GetParamInt(ctx, "is_display")
	sequence := GetParamIntDef(ctx, "sequence", 1)

	// 判断参数合法性
	if privilegeId == 0 {
		logger.WithContext(ctx).Warnf("[PrivilegeModify] privilege_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "参数错误")
	}
	if name == "" {
		logger.WithContext(ctx).Warnf("[PrivilegeModify] name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "权限名不能为空")
	}
	if privilegeType == 0 {
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "权限类型错误")
	}

	// privilege 权限实体
	privilegeEntity := entity.PrivilegeEntity{
		PrivilegeId:   int64(privilegeId),
		Name:          name,
		ParentIds:     parentIds,
		PrivilegeType: privilegeType,
		PageRouter:    pageRouter,
		ApiMarks:      apiMarks,
		Icon:          icon,
		IsDisplay:     isDisplay,
		Sequence:      sequence,
	}

	err := service.NewPrivilege(ctx).Update(privilegeEntity)
	if err != nil {
		sysLogErrorf(ctx, "[PrivilegeModify] 修改权限 %d 失败: err=%+v", privilegeId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[PrivilegeModify] 修改权限 %d 成功", privilegeId)

	return RespJsonSuccess(ctx, nil)
}

// PrivilegeList 权限列表
func PrivilegeList(ctx *gin.Context) error {

	servicePrivilege := service.NewPrivilege(ctx)
	// 获取权限列表
	privilegeList, err := servicePrivilege.GetAllPrivilegeList()
	if err != nil {
		sysLogErrorf(ctx, "[PrivilegeList] 获取权限列表失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	data := map[string]interface{}{
		"list": privilegeList,
	}
	return RespJsonSuccess(ctx, data)
}

// PrivilegeDelete 权限删除
func PrivilegeDelete(ctx *gin.Context) error {

	privilegeId := GetParamInt(ctx, "privilege_id")

	// 判断参数合法性
	if privilegeId == 0 {
		logger.WithContext(ctx).Warnf("[PrivilegeDelete] privilege_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "参数错误")
	}
	// 删除权限
	err := service.NewPrivilege(ctx).DeleteByPrivilegeId(int64(privilegeId))
	if err != nil {
		sysLogErrorf(ctx, "[PrivilegeDelete] 删除权限 %d 失败 err=%+v", privilegeId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	// 删除角色下的权限
	err = service.NewRole(ctx).DeletePrivilegeIdsByPrivilegeId(int64(privilegeId))
	if err != nil {
		sysLogErrorf(ctx, "[PrivilegeDelete] 删除权限 %d 失败 err=%+v", privilegeId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[PrivilegeDelete] 删除权限 %d 成功", privilegeId)

	return RespJsonSuccess(ctx, nil)
}
