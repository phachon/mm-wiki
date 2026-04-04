package system

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/controller"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
)

// ProfilePrivileges 获取系统个人权限
func ProfilePrivileges(ctx *gin.Context) error {

	accountId := global.ContextValueLoginAccountID(ctx)
	navKey := controller.GetParamString(ctx, "nav_key")

	// 判断参数合法性
	if accountId == 0 {
		logger.WithContext(ctx).Warnf("[ProfileMenus] 账号id不存在")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "账号id不存在")
	}
	if navKey == "" {
		logger.WithContext(ctx).Warnf("[ProfileMenus] 导航key不存在")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "导航key不存在")
	}

	_, err := service.NewAccount(ctx).GetAccountByAccountId(accountId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[ProfileMenus] 获取账号 %d 信息失败: err=%+v", accountId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取账号所有的权限
	privileges, err := service.NewPermission(ctx).GetAccountPrivileges(accountId)
	if err != nil {
		sysLogErrorf(ctx, "[ProfileInfo] 获取角色权限失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取权限列表树结构
	privilegeItems, err := service.NewPrivilege(ctx).GetPrivilegeListItems(privileges)
	if err != nil {
		sysLogErrorf(ctx, "[ProfileInfo] 获取权限列表失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 根据 navKey 过滤次导航下的权限
	var profilePrivileges []*entity.PrivilegeListItem
	for _, privilegeItem := range privilegeItems {
		if privilegeItem.PrivilegeEntity.Identify == navKey {
			profilePrivileges = privilegeItem.ChildPrivileges
			break
		}
	}
	type Data struct {
		Privileges []*entity.PrivilegeListItem `json:"privileges"` // 权限列表
	}
	data := new(Data)
	data.Privileges = profilePrivileges

	return RespJsonSuccess(ctx, data)
}

// ProfileInfo 个人信息
func ProfileInfo(ctx *gin.Context) error {

	accountId := global.ContextValueLoginAccountID(ctx)
	// 判断参数合法性
	if accountId == 0 {
		logger.WithContext(ctx).Warnf("[ProfileInfo] 账号id不存在")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "账号id不存在")
	}

	account, err := service.NewAccount(ctx).GetAccountByAccountId(accountId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[ProfileInfo] 获取账号 %d 信息失败: err=%+v", accountId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 密码置空
	account.Password = ""

	// 获取部门全称
	departmentNames, err := service.NewDepartment(ctx).GetDepartmentFullNames(account.DepartmentId)
	if err != nil {
		logger.WithContext(ctx).Errorf("[ProfileInfo] 获取账号 %d 部门全称失败: err=%+v", accountId, err)
	}
	type ProfileData struct {
		AccountInfo     *entity.AccountEntity `json:"account_info"`     // 账号信息
		DepartmentNames []string              `json:"department_names"` // 部门全称
	}
	data := new(ProfileData)
	data.AccountInfo = account
	data.DepartmentNames = departmentNames

	return RespJsonSuccess(ctx, data)
}

// ProfileRePass 修改密码
func ProfileRePass(ctx *gin.Context) error {

	accountId := global.ContextValueLoginAccountID(ctx)
	oldPassword := GetParamString(ctx, "old_pwd")
	newPassword := GetParamString(ctx, "new_pwd")
	confirmPassword := GetParamString(ctx, "confirm_pwd")

	// 密码删除，防止后续上报误报
	ctx.Request.PostForm.Del("old_pwd")
	ctx.Request.PostForm.Del("new_pwd")
	ctx.Request.PostForm.Del("confirm_pwd")

	// 判断参数合法性
	if accountId == 0 {
		logger.WithContext(ctx).Warnf("[ProfileRePass] account not login")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "账号id不存在")
	}
	if oldPassword == "" || newPassword == "" || confirmPassword == "" {
		logger.WithContext(ctx).Warnf("[ProfileRePass] pwd empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "密码不能为空")
	}
	if oldPassword == newPassword {
		logger.WithContext(ctx).Warnf("[ProfileRePass] new password equal old password")
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "新密码不能与旧密码相同")
	}
	if newPassword != confirmPassword {
		logger.WithContext(ctx).Warnf("[ProfileRePass] new_pwd not equal confirmPwd")
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "确认密码不一致")
	}

	// 更新密码
	err := service.NewAccount(ctx).UpdatePassword(accountId, oldPassword, newPassword)
	if err != nil {
		sysLogErrorf(ctx, "[ProfileRePass] 更新账号 %d 密码失败: err=%s", accountId, err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[ProfileRePass] 更新账号 %d 密码成功", accountId)

	return RespJsonSuccess(ctx, nil)
}

// ProfileUpdate 修改资料
func ProfileUpdate(ctx *gin.Context) error {

	accountId := global.ContextValueLoginAccountID(ctx)
	givenName := GetParamString(ctx, "given_name")
	mobile := GetParamString(ctx, "mobile")
	phone := GetParamString(ctx, "phone")
	email := GetParamString(ctx, "email")
	departmentId := GetParamInt64(ctx, "department_id")
	position := GetParamString(ctx, "position")
	location := GetParamString(ctx, "location")

	// 判断参数合法性
	if accountId == 0 {
		logger.WithContext(ctx).Warnf("[ProfileUpdate] account not login")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "账号id不存在")
	}
	if givenName == "" {
		logger.WithContext(ctx).Warnf("[ProfileUpdate] given_name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "昵称不能为空")
	}
	// account 账号实体
	accountEntity := entity.AccountEntity{
		AccountId:    accountId,
		GivenName:    givenName,
		Email:        email,
		Phone:        phone,
		Mobile:       mobile,
		DepartmentId: departmentId,
		Position:     position,
		Location:     location,
	}
	err := service.NewAccount(ctx).Update(accountEntity)
	if err != nil {
		sysLogErrorf(ctx, "[ProfileUpdate] 更新个人信息失败: err=%s", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[ProfileUpdate] 更新个人信息成功")

	return RespJsonSuccess(ctx, nil)
}

// ProfileFollowDocs 获取用户关注的文档列表
func ProfileFollowDocs(ctx *gin.Context) error {

	accountId := global.ContextValueLoginAccountID(ctx)
	pageSize := controller.GetParamIntDef(ctx, "page_size", 10)
	pageNum := controller.GetParamIntDef(ctx, "page_num", 1)

	if accountId == 0 {
		logger.WithContext(ctx).Warnf("[ProfileFollowDocs] 账号id不存在")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "账号id不存在")
	}

	followService := service.NewFollow(ctx)

	// 获取关注的文档列表
	follows, err := followService.GetFollowsByAccountIdAndType(accountId, entity.FollowTypeDocument, pageSize, pageNum)
	if err != nil {
		logger.WithContext(ctx).Errorf("[ProfileFollowDocs] GetFollowsByAccountIdAndType err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), "获取关注文档列表失败")
	}

	// 获取文档详情
	docIds := make([]int64, 0, len(follows))
	for _, follow := range follows {
		docId, _ := strconv.ParseInt(follow.ObjectId, 10, 64)
		if docId > 0 {
			docIds = append(docIds, docId)
		}
	}

	var docs []*entity.DocEntity
	if len(docIds) > 0 {
		docService := service.NewDoc(ctx)
		docs, err = docService.GetDocByDocIds(docIds)
		if err != nil {
			logger.WithContext(ctx).Errorf("[ProfileFollowDocs] GetDocByDocIds err=%+v", err)
			return RespJsonError(ctx, err.GetErrCode(), "获取文档信息失败")
		}
	}

	// 获取总数
	total, err := followService.CountFollowsByAccountIdAndType(accountId, entity.FollowTypeDocument)
	if err != nil {
		logger.WithContext(ctx).Errorf("[ProfileFollowDocs] CountFollowsByAccountIdAndType err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), "获取关注文档总数失败")
	}
	pageInfo := entity.GetPageInfo(total, pageSize, pageNum)

	data := map[string]interface{}{
		"list":      docs,
		"page_info": pageInfo,
	}
	return RespJsonSuccess(ctx, data)
}

// ProfileFollowUsers 获取用户关注的用户列表
func ProfileFollowUsers(ctx *gin.Context) error {

	accountId := global.ContextValueLoginAccountID(ctx)
	pageSize := controller.GetParamIntDef(ctx, "page_size", 10)
	pageNum := controller.GetParamIntDef(ctx, "page_num", 1)

	if accountId == 0 {
		logger.WithContext(ctx).Warnf("[ProfileFollowUsers] 账号id不存在")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "账号id不存在")
	}

	followService := service.NewFollow(ctx)

	// 获取关注的用户列表
	follows, err := followService.GetFollowsByAccountIdAndType(accountId, entity.FollowTypeUser, pageSize, pageNum)
	if err != nil {
		logger.WithContext(ctx).Errorf("[ProfileFollowUsers] GetFollowsByAccountIdAndType err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), "获取关注用户列表失败")
	}

	// 获取用户详情
	userIds := make([]int64, 0, len(follows))
	for _, follow := range follows {
		userId, _ := strconv.ParseInt(follow.ObjectId, 10, 64)
		if userId > 0 {
			userIds = append(userIds, userId)
		}
	}

	var accounts []*entity.AccountEntity
	if len(userIds) > 0 {
		accountService := service.NewAccount(ctx)
		accounts, err = accountService.GetAccountsByAccountIds(userIds)
		if err != nil {
			logger.WithContext(ctx).Errorf("[ProfileFollowUsers] GetAccountsByAccountIds err=%+v", err)
			return RespJsonError(ctx, err.GetErrCode(), "获取用户信息失败")
		}
		// 清空密码
		for _, account := range accounts {
			account.Password = ""
		}
	}

	// 获取总数
	total, err := followService.CountFollowsByAccountIdAndType(accountId, entity.FollowTypeUser)
	if err != nil {
		logger.WithContext(ctx).Errorf("[ProfileFollowUsers] CountFollowsByAccountIdAndType err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), "获取关注用户总数失败")
	}
	pageInfo := entity.GetPageInfo(total, pageSize, pageNum)

	data := map[string]interface{}{
		"list":      accounts,
		"page_info": pageInfo,
	}
	return RespJsonSuccess(ctx, data)
}

// ProfileActivity 获取用户活动日志
func ProfileActivity(ctx *gin.Context) error {

	accountId := global.ContextValueLoginAccountID(ctx)
	pageSize := controller.GetParamIntDef(ctx, "page_size", 20)
	pageNum := controller.GetParamIntDef(ctx, "page_num", 1)

	if accountId == 0 {
		logger.WithContext(ctx).Warnf("[ProfileActivity] 账号id不存在")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "账号id不存在")
	}

	logService := service.NewLog(ctx)

	// 获取用户操作日志
	keywords := &entity.LogSearchKeywords{
		AccountId: accountId,
	}
	logs, err := logService.GetLogsByKeywordsAndLimit(pageSize, pageNum, keywords)
	if err != nil {
		logger.WithContext(ctx).Errorf("[ProfileActivity] GetLogsByKeywordsAndLimit err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), "获取活动日志失败")
	}

	// 获取分页信息
	pageInfo, err := logService.GetPageInfoLimit(pageSize, pageNum, keywords)
	if err != nil {
		logger.WithContext(ctx).Errorf("[ProfileActivity] GetPageInfoLimit err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), "获取活动日志分页失败")
	}

	data := map[string]interface{}{
		"list":      logs,
		"page_info": pageInfo,
	}
	return RespJsonSuccess(ctx, data)
}
