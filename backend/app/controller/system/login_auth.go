package system

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
)

// LoginAuthSave 添加登录认证保存
func LoginAuthSave(ctx *gin.Context) error {

	name := GetParamString(ctx, "name")
	accountPrefix := GetParamString(ctx, "account_prefix")
	url := GetParamString(ctx, "url")
	extData := GetParamStringDef(ctx, "ext_data", "")
	// 判断参数合法性
	if name == "" {
		logger.WithContext(ctx).Warnf("[LoginAuthSave] name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "认证名称不能为空")
	}
	if url == "" {
		logger.WithContext(ctx).Warnf("[LoginAuthSave] url empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "认证接口URL不能为空")
	}

	// loginAuth 认证实体
	loginAuthEntity := &entity.LoginAuthEntity{
		Name:          name,
		AccountPrefix: accountPrefix,
		URL:           url,
		ExtData:       extData,
		IsUsed:        entity.LoginAuthIsUsedNo,
		Status:        entity.LoginAuthStatusDefault,
	}
	// 创建认证
	err := service.NewLoginAuth(ctx).Create(loginAuthEntity)
	if err != nil {
		sysLogErrorf(ctx, "[LoginAuthSave] 添加认证失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[LoginAuthSave] 添加认证 %d 成功", loginAuthEntity.LoginAuthId)

	return RespJsonSuccess(ctx, nil)
}

// LoginAuthEdit 认证修改页面
func LoginAuthEdit(ctx *gin.Context) error {

	loginAuthId := GetParamInt64(ctx, "login_auth_id")
	// 判断参数合法性
	if loginAuthId == 0 {
		logger.WithContext(ctx).Warnf("[LoginAuthEdit] 认证id不存在")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "认证id不存在")
	}

	loginAuth, err := service.NewLoginAuth(ctx).GetLoginAuthByLoginAuthId(loginAuthId)
	if err != nil {
		sysLogErrorf(ctx, "[LoginAuthEdit] 获取认证 %d 信息失败: err=%+v", loginAuthId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if loginAuth == nil {
		logger.WithContext(ctx).Warnf("[LoginAuthEdit] 认证 %d 不存在", loginAuthId)
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "认证id错误")
	}

	return RespJsonSuccess(ctx, map[string]interface{}{
		"login_auth_info": loginAuth,
	})
}

// LoginAuthModify 修改认证保存
func LoginAuthModify(ctx *gin.Context) error {

	loginAuthId := GetParamInt64(ctx, "login_auth_id")
	name := GetParamString(ctx, "name")
	accountPrefix := GetParamString(ctx, "account_prefix")
	url := GetParamString(ctx, "url")
	extData := GetParamStringDef(ctx, "ext_data", "")
	// 判断参数合法性
	if loginAuthId <= 0 {
		logger.WithContext(ctx).Warnf("[LoginAuthModify] login_auth_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "认证ID不合法")
	}
	if name == "" {
		logger.WithContext(ctx).Warnf("[LoginAuthModify] name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "认证名称不能为空")
	}
	if url == "" {
		logger.WithContext(ctx).Warnf("[LoginAuthModify] url empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "认证接口URL不能为空")
	}

	// loginAuth 认证实体
	loginAuthEntity := entity.LoginAuthEntity{
		LoginAuthId:   loginAuthId,
		Name:          name,
		AccountPrefix: accountPrefix,
		URL:           url,
		ExtData:       extData,
	}

	// 更新认证
	err := service.NewLoginAuth(ctx).Update(loginAuthEntity)
	if err != nil {
		sysLogErrorf(ctx, "[LoginAuthModify] 更新认证 %d 失败: err=%+v", loginAuthId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[LoginAuthModify] 更新认证 %d 成功", loginAuthId)

	return RespJsonSuccess(ctx, nil)
}

// LoginAuthList 认证列表
func LoginAuthList(ctx *gin.Context) error {

	pageSize := GetParamIntDef(ctx, "page_size", 20)
	pageNum := GetParamIntDef(ctx, "page_num", 1)
	keywordsStr := GetParamStringDef(ctx, "keywords", "")

	var keywords *entity.LoginAuthKeywords
	if len(keywordsStr) > 0 {
		jErr := json.Unmarshal([]byte(keywordsStr), &keywords)
		if jErr != nil {
			logger.WithContext(ctx).Errorf("[LoginAuthList] keywords unmarshal err=%s", jErr.Error())
		}
	}

	loginAuthService := service.NewLoginAuth(ctx)

	// 获取认证列表
	loginAuths, err := loginAuthService.GetLoginAuthsByLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[LoginAuthList] 获取认证列表失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取分页信息
	pageInfo, err := loginAuthService.GetPageInfoLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[LoginAuthList] 获取认证分页信息失败: err=%s", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 格式化列表结构
	loginAuthList, err := loginAuthService.FormatLoginAuthList(loginAuths)
	if err != nil {
		sysLogErrorf(ctx, "[LoginAuthList] 格式化认证列表失败: err=%s", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	data := map[string]interface{}{
		"list":      loginAuthList,
		"page_info": pageInfo,
	}
	return RespJsonSuccess(ctx, data)
}

// LoginAuthDelete 认证删除
func LoginAuthDelete(ctx *gin.Context) error {

	loginAuthId := GetParamInt64(ctx, "login_auth_id")

	// 判断参数合法性
	if loginAuthId <= 0 {
		logger.WithContext(ctx).Warnf("[LoginAuthDelete] login_auth_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "认证id不存在")
	}

	// 删除认证
	err := service.NewLoginAuth(ctx).DeleteLoginAuth(loginAuthId)
	if err != nil {
		sysLogErrorf(ctx, "[LoginAuthDelete] 删除认证 %d 失败: err=%+v", loginAuthId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[LoginAuthDelete] 删除认证 %d 成功", loginAuthId)

	return RespJsonSuccess(ctx, nil)
}

// LoginAuthUsed 设置认证为使用中
func LoginAuthUsed(ctx *gin.Context) error {

	loginAuthId := GetParamInt64(ctx, "login_auth_id")

	// 判断参数合法性
	if loginAuthId <= 0 {
		logger.WithContext(ctx).Warnf("[LoginAuthUsed] login_auth_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "认证id不存在")
	}

	// 设置使用
	err := service.NewLoginAuth(ctx).SetUsed(loginAuthId)
	if err != nil {
		sysLogErrorf(ctx, "[LoginAuthUsed] 设置认证 %d 使用失败: err=%+v", loginAuthId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[LoginAuthUsed] 设置认证 %d 使用成功", loginAuthId)

	return RespJsonSuccess(ctx, nil)
}
