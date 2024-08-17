package system

import (
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
	type ProfileData struct {
		AccountInfo *entity.AccountEntity `json:"account_info"` // 账号信息
	}
	data := new(ProfileData)
	data.AccountInfo = account

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
