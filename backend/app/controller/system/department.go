package system

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
)

// DepartmentAdd 添加部门页面
func DepartmentAdd(ctx *gin.Context) error {
	// serviceDepartment := service.NewDepartment(ctx)
	// departmentList, err := serviceDepartment.GetAllDepartmentList()
	// if err != nil {
	// 	sysLogErrorf(ctx, "[DepartmentAdd] 获取部门列表失败: err=%+v", err)
	// 	return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	// }
	// data := map[string]interface{}{
	// 	"departments": departmentList,
	// }
	// return RespJsonSuccess(ctx, data)
	return nil
}

// DepartmentSave 添加部门保存
func DepartmentSave(ctx *gin.Context) error {

	name := GetParamString(ctx, "name")
	parentId := GetParamIntDef(ctx, "parent_id", 0)
	sequence := GetParamIntDef(ctx, "sequence", 0)

	// 判断参数合法性
	if name == "" {
		logger.WithContext(ctx).Warnf("[DepartmentSave] name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "部门名不能为空")
	}

	// 获取 parent_ids
	serviceDepartment := service.NewDepartment(ctx)
	parentIds := "0"
	if parentId != 0 {
		parentDepartment, err := serviceDepartment.GetDepartmentByDepartmentId(int64(parentId))
		if err != nil {
			sysLogErrorf(ctx, "[DepartmentSave] 添加部门失败: err=%+v", err)
			return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
		}
		if parentDepartment != nil {
			parentIds = parentDepartment.ParentIds + "," + fmt.Sprintf("%d", parentId)
		}
	}

	// department 部门实体
	departmentEntity := &entity.DepartmentEntity{
		Name:      name,
		ParentId:  int64(parentId),
		ParentIds: parentIds,
		Sequence:  sequence,
	}
	// 添加部门
	err := serviceDepartment.Create(departmentEntity)
	if err != nil {
		sysLogErrorf(ctx, "[DepartmentSave] 添加部门失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[DepartmentSave] 添加部门 %+v 成功", departmentEntity.DepartmentId)

	return RespJsonSuccess(ctx, nil)
}

// DepartmentEdit 修改编辑部门页面
func DepartmentEdit(ctx *gin.Context) error {

	departmentId := GetParamInt64(ctx, "department_id")
	// 判断参数合法性
	if departmentId <= 0 {
		logger.WithContext(ctx).Warnf("[DepartmentEdit] department_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "参数错误")
	}
	serviceDepartment := service.NewDepartment(ctx)

	// 查找部门是否存在
	department, err := serviceDepartment.GetDepartmentByDepartmentId(departmentId)
	if err != nil {
		return err
	}
	if department == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "部门不存在")
	}

	data := map[string]interface{}{
		"department": department,
	}
	return RespJsonSuccess(ctx, data)
}

// DepartmentModify 修改部门保存
func DepartmentModify(ctx *gin.Context) error {

	departmentId := GetParamInt(ctx, "department_id")
	name := GetParamString(ctx, "name")
	parentId := GetParamIntDef(ctx, "parent_id", 0)
	sequence := GetParamIntDef(ctx, "sequence", 0)

	// 判断参数合法性
	if departmentId == 0 {
		logger.WithContext(ctx).Warnf("[DepartmentModify] department_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "参数错误")
	}
	if name == "" {
		logger.WithContext(ctx).Warnf("[DepartmentModify] name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "部门名不能为空")
	}
	if parentId == departmentId {
		logger.WithContext(ctx).Warnf("[DepartmentModify] parent_id equal department_id")
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "上级部门不能是自己")
	}

	// 获取 parent_ids
	parentIds := "0"
	if parentId != 0 {
		serviceDepartment := service.NewDepartment(ctx)
		parentDepartment, err := serviceDepartment.GetDepartmentByDepartmentId(int64(parentId))
		if err != nil {
			sysLogErrorf(ctx, "[DepartmentModify] 修改部门失败: err=%+v", err)
			return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
		}
		if parentDepartment != nil {
			parentIds = parentDepartment.ParentIds + "," + fmt.Sprintf("%d", parentId)
		}
	}

	// department 部门实体
	departmentEntity := entity.DepartmentEntity{
		DepartmentId: int64(departmentId),
		Name:         name,
		ParentId:     int64(parentId),
		ParentIds:    parentIds,
		Sequence:     sequence,
	}
	err := service.NewDepartment(ctx).Update(departmentEntity)
	if err != nil {
		sysLogErrorf(ctx, "[DepartmentModify] 修改部门 %d 失败: err=%+v", departmentId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[DepartmentModify] 修改部门 %d 成功", departmentId)

	return RespJsonSuccess(ctx, nil)
}

// DepartmentList 部门列表
func DepartmentList(ctx *gin.Context) error {

	serviceDepartment := service.NewDepartment(ctx)
	// 获取部门列表
	departmentList, err := serviceDepartment.GetAllDepartmentList()
	if err != nil {
		sysLogErrorf(ctx, "[DepartmentList] 获取部门列表失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	data := map[string]interface{}{
		"list": departmentList,
	}
	return RespJsonSuccess(ctx, data)
}

// DepartmentDelete 部门删除
func DepartmentDelete(ctx *gin.Context) error {

	departmentId := GetParamInt(ctx, "department_id")

	// 判断参数合法性
	if departmentId == 0 {
		logger.WithContext(ctx).Warnf("[DepartmentDelete] department_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "参数错误")
	}
	// 删除部门
	err := service.NewDepartment(ctx).DeleteByDepartmentId(int64(departmentId))
	if err != nil {
		sysLogErrorf(ctx, "[DepartmentDelete] 删除部门 %d 失败 err=%+v", departmentId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[DepartmentDelete] 删除部门 %d 成功", departmentId)

	return RespJsonSuccess(ctx, nil)
}
