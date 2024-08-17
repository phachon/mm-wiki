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

// AccountAdd 添加账号页面
func AccountAdd(ctx *gin.Context) error {

	// 获取所有的角色
	allRoles, err := service.NewRole(ctx).GetAllRoles()
	if err != nil {
		sysLogErrorf(ctx, "[AccountAdd] 获取所有角色失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	// 重新排序，默认角色在最上面
	accountDefRoles := []*entity.RoleEntity{}
	otherRoles := []*entity.RoleEntity{}
	for _, role := range allRoles {
		if role.RoleType == entity.RoleTypeAccountDefaultRole {
			accountDefRoles = append(accountDefRoles, role)
			continue
		}
		otherRoles = append(otherRoles, role)
	}
	// 获取所有的部门列表
	departments, err := service.NewDepartment(ctx).GetAllDepartmentList()
	if err != nil {
		sysLogErrorf(ctx, "[AccountAdd] 获取部门失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	return RespJsonSuccess(ctx, map[string]interface{}{
		"roles":       append(accountDefRoles, otherRoles...),
		"departments": departments,
	})
}

// AccountSave 添加账号保存
func AccountSave(ctx *gin.Context) error {
	name := GetParamString(ctx, "name")
	givenName := GetParamString(ctx, "given_name")
	roleIds := GetParamString(ctx, "role_ids")
	mobile := GetParamString(ctx, "mobile")
	phone := GetParamString(ctx, "phone")
	email := GetParamString(ctx, "email")
	departmentId := GetParamInt64(ctx, "department_id")
	position := GetParamString(ctx, "position")
	location := GetParamString(ctx, "location")

	// 判断参数合法性
	if name == "" {
		logger.WithContext(ctx).Warnf("[AccountSave] 账号名不能为空")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "账号名不能为空")
	}
	if givenName == "" {
		logger.WithContext(ctx).Warnf("[AccountSave] 昵称不能为空")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "昵称不能为空")
	}
	if len(roleIds) == 0 {
		logger.WithContext(ctx).Warnf("[AccountSave] 角色不能为空")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "角色不能为空")
	}
	roleList := strings.Split(roleIds, ",")
	var roleIdsInt []int64
	for _, roleIdStr := range roleList {
		roleIdsInt = append(roleIdsInt, utils.Convert.StringToInt64(roleIdStr))
	}
	// 校验角色的合法性
	roles, err := service.NewRole(ctx).GetRolesByRoleIds(roleIdsInt)
	if err != nil {
		sysLogErrorf(ctx, "[AccountSave] 添加账号获取角色失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if len(roleIdsInt) != len(roles) {
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "角色不合法")
	}

	// account 账号实体
	accountEntity := &entity.AccountEntity{
		Name:         name,
		GivenName:    givenName,
		Mobile:       mobile,
		Phone:        phone,
		Email:        email,
		DepartmentId: departmentId,
		Position:     position,
		Location:     location,
	}
	err = service.NewAccount(ctx).Create(accountEntity)
	if err != nil {
		sysLogErrorf(ctx, "[AccountSave] 添加账号失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	// 添加账号和角色对应关系
	err = service.NewAccountRole(ctx).BatchCreate(accountEntity.AccountId, roleIdsInt)
	if err != nil {
		sysLogErrorf(ctx, "[AccountSave] 添加账号角色对应关系失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[AccountSave] 添加账号 %+v 成功", accountEntity.AccountId)

	return RespJsonSuccess(ctx, nil)
}

// AccountEdit 修改账号页面
func AccountEdit(ctx *gin.Context) error {

	accountId := GetParamInt64(ctx, "account_id")
	// 判断参数合法性
	if accountId == 0 {
		logger.WithContext(ctx).Warnf("[AccountEdit] 账号id不能为空")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "账号id不能为空")
	}

	// 获取账号信息
	accountInfo, err := service.NewAccount(ctx).GetAccountByAccountId(accountId)
	if err != nil {
		sysLogErrorf(ctx, "[AccountEdit] 获取账号 %d 信息失败: err=%+v", accountId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if accountInfo == nil {
		sysLogWarnf(ctx, "[AccountEdit] 账号id %d 不合法", accountId)
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "账号id不存在")
	}
	// 获取账号角色
	accountRoles, err := service.NewAccountRole(ctx).GetRolesByAccountId(accountInfo.AccountId)
	if err != nil {
		sysLogErrorf(ctx, "[AccountEdit] 获取账号 %d 下角色失败: err=%+v", accountId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取所有的角色
	allRoles, err := service.NewRole(ctx).GetAllRoles()
	if err != nil {
		sysLogErrorf(ctx, "[AccountEdit] 获取所有的角色失败 err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取所有的部门列表
	departments, err := service.NewDepartment(ctx).GetAllDepartmentList()
	if err != nil {
		sysLogErrorf(ctx, "[AccountAdd] 获取部门失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	return RespJsonSuccess(ctx, &entity.AccountEditResp{
		Departments:  departments,
		AccountInfo:  accountInfo,
		AccountRoles: accountRoles,
		RoleList:     allRoles,
	})
}

// AccountModify 修改账号保存
func AccountModify(ctx *gin.Context) error {

	accountId := GetParamInt64(ctx, "account_id")
	name := GetParamString(ctx, "name")
	roleIds := GetParamString(ctx, "role_ids")
	givenName := GetParamString(ctx, "given_name")
	mobile := GetParamString(ctx, "mobile")
	phone := GetParamString(ctx, "phone")
	email := GetParamString(ctx, "email")
	departmentId := GetParamInt64(ctx, "department_id")
	position := GetParamString(ctx, "position")
	location := GetParamString(ctx, "location")

	// 判断参数合法性
	if accountId == 0 {
		logger.WithContext(ctx).Warnf("[AccountModify] account_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "账号id不存在")
	}
	if name == "" {
		logger.WithContext(ctx).Warnf("[AccountModify] name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "账号名不能为空")
	}
	if givenName == "" {
		logger.WithContext(ctx).Warnf("[AccountModify] given_name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "昵称不能为空")
	}
	if len(roleIds) == 0 {
		logger.WithContext(ctx).Warnf("[AccountModify] role_ids empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "角色不能为空")
	}
	roleList := strings.Split(roleIds, ",")
	var roleIdsInt []int64
	for _, roleIdStr := range roleList {
		roleIdsInt = append(roleIdsInt, utils.Convert.StringToInt64(roleIdStr))
	}
	// 校验角色的合法性
	roles, err := service.NewRole(ctx).GetRolesByRoleIds(roleIdsInt)
	if err != nil {
		sysLogErrorf(ctx, "[AccountModify] 获取角色信息失败: err=%+v", err)
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "角色不合法")
	}
	if len(roleIdsInt) != len(roles) {
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "角色不合法")
	}

	// account 账号实体
	accountEntity := entity.AccountEntity{
		AccountId:    accountId,
		GivenName:    givenName,
		Mobile:       mobile,
		Phone:        phone,
		Email:        email,
		DepartmentId: departmentId,
		Position:     position,
		Location:     location,
	}
	err = service.NewAccount(ctx).Update(accountEntity)
	if err != nil {
		sysLogErrorf(ctx, "[AccountModify] 更新账号 %d 失败 err=%+v", accountId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	// 更新账号和角色的对应关系
	err = service.NewAccountRole(ctx).UpdateByAccountId(accountId, roleIdsInt)
	if err != nil {
		sysLogErrorf(ctx, "[AccountModify] 更新账号 %d 角色 %+v 失败: err=%+v",
			accountId, roleIdsInt, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[AccountModify] 更新账号 %d 成功", accountId)

	return RespJsonSuccess(ctx, nil)
}

// AccountDetail 账号详情
func AccountDetail(ctx *gin.Context) error {

	accountId := GetParamInt64(ctx, "account_id")

	// 判断参数合法性
	if accountId == 0 {
		logger.WithContext(ctx).Warnf("[AccountDetail] account_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "账号id不存在")
	}

	// 获取账号信息
	accountInfo, err := service.NewAccount(ctx).GetAccountByAccountId(accountId)
	if err != nil {
		sysLogErrorf(ctx, "[AccountDetail] 获取账号 %d 信息失败: err=%+v", accountId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if accountInfo == nil {
		sysLogWarnf(ctx, "[AccountDetail] 账号id %d 不存在", accountId)
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "账号id不存在")
	}
	// 密码置空
	accountInfo.Password = ""

	// 获取账号角色
	accountRoles, err := service.NewAccountRole(ctx).GetRolesByAccountId(accountInfo.AccountId)
	if err != nil {
		sysLogErrorf(ctx, "[AccountDetail] 获取账号 %d 下角色失败: err=%+v",
			accountId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取部门全称
	departmentNames, err := service.NewDepartment(ctx).GetDepartmentFullNames(accountInfo.DepartmentId)
	if err != nil {
		logger.WithContext(ctx).Warnf("[AccountDetail] 获取账号 %d 部门全称失败: err=%+v", accountId, err)
	}
	return RespJsonSuccess(ctx, &entity.AccountDetailResp{
		AccountEntity:   accountInfo,
		Roles:           accountRoles,
		DepartmentNames: departmentNames,
	})
}

// AccountList 账号列表
func AccountList(ctx *gin.Context) error {

	pageSize := GetParamIntDef(ctx, "page_size", 20)
	pageNum := GetParamIntDef(ctx, "page_num", 1)
	keywordsStr := GetParamStringDef(ctx, "keywords", "")

	var keywords *entity.AccountKeywords
	if len(keywordsStr) > 0 {
		jErr := json.Unmarshal([]byte(keywordsStr), &keywords)
		if jErr != nil {
			logger.WithContext(ctx).Errorf("[AccountList] GetAccountsByLimit err=%s", jErr.Error())
		}
	}

	serviceAccount := service.NewAccount(ctx)
	// 获取账号列表
	accounts, err := serviceAccount.GetAccountsByLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[AccountList] 获取账号列表失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取分页信息
	pageInfo, err := serviceAccount.GetPageInfoLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[AccountList] 获取账号分页信息失败: err=%+v", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 格式化账号列表
	accountList, err := serviceAccount.FormatAccountList(accounts)
	if err != nil {
		sysLogErrorf(ctx, "[AccountList] 格式化账号列表失败: err=%+v", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	data := map[string]interface{}{
		"list":      accountList,
		"page_info": pageInfo,
	}
	return RespJsonSuccess(ctx, data)
}

// AccountUpdateStatus 账号更新状态
func AccountUpdateStatus(ctx *gin.Context) error {

	accountId := GetParamInt64(ctx, "account_id")
	status := GetParamIntDef(ctx, "status", -100)

	// 判断参数合法性
	if accountId == 0 {
		logger.WithContext(ctx).Warnf("[AccountUpdateStatus] account_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "账号id不存在")
	}
	if status != entity.AccountStatusDefault && status != entity.AccountStatusForbid {
		logger.WithContext(ctx).Warnf("[AccountUpdateStatus] status err")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "状态参数错误")
	}

	// 更新账号状态
	err := service.NewAccount(ctx).UpdateStatus(accountId, status)
	if err != nil {
		sysLogErrorf(ctx, "[AccountUpdateStatus] 更新账号 %d 状态失败: err=%+v", accountId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[AccountUpdateStatus] 更新账号 %d 状态成功", accountId)

	return RespJsonSuccess(ctx, nil)
}
