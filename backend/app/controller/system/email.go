package system

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
)

// EmailSave 添加邮箱保存
func EmailSave(ctx *gin.Context) error {

	name := GetParamString(ctx, "name")
	senderAddress := GetParamString(ctx, "sender_address")
	senderName := GetParamString(ctx, "sender_name")
	senderTitlePrefix := GetParamString(ctx, "sender_title_prefix")
	host := GetParamString(ctx, "host")
	port := GetParamInt(ctx, "port")
	username := GetParamString(ctx, "username")
	password := GetParamString(ctx, "password")
	isSSL := GetParamInt(ctx, "is_ssl")
	// 判断参数合法性
	if name == "" {
		logger.WithContext(ctx).Warnf("[EmailSave] name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "邮箱名称不能为空")
	}

	// email 邮箱实体
	emailEntity := &entity.EmailEntity{
		Name:              name,
		SenderAddress:     senderAddress,
		SenderName:        senderName,
		SenderTitlePrefix: senderTitlePrefix,
		Host:              host,
		Port:              port,
		Username:          username,
		Password:          password,
		IsSSL:             isSSL,
		Status:            entity.EmailStatusDefault,
	}
	// 创建邮箱
	err := service.NewEmail(ctx).Create(emailEntity)
	if err != nil {
		sysLogErrorf(ctx, "[EmailSave] 添加邮箱失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[EmailSave] 添加邮箱 %d 成功", emailEntity.EmailId)

	return RespJsonSuccess(ctx, nil)
}

// EmailEdit 邮箱修改页面
func EmailEdit(ctx *gin.Context) error {

	emailId := GetParamInt64(ctx, "email_id")
	// 判断参数合法性
	if emailId == 0 {
		logger.WithContext(ctx).Warnf("[EmailEdit] 邮箱id不存在")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "邮箱id不存在")
	}

	email, err := service.NewEmail(ctx).GetEmailByEmailId(emailId)
	if err != nil {
		sysLogErrorf(ctx, "[EmailEdit] 获取邮箱 %d 信息失败: err=%+v", emailId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if email == nil {
		logger.WithContext(ctx).Warnf("[EmailEdit] 邮箱 %d 不存在", emailId)
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "邮箱id错误")
	}

	return RespJsonSuccess(ctx, map[string]interface{}{
		"email_info": email,
	})
}

// EmailModify 修改邮箱保存
func EmailModify(ctx *gin.Context) error {

	emailId := GetParamInt64(ctx, "email_id")
	name := GetParamString(ctx, "name")
	senderAddress := GetParamString(ctx, "sender_address")
	senderName := GetParamString(ctx, "sender_name")
	senderTitlePrefix := GetParamString(ctx, "sender_title_prefix")
	host := GetParamString(ctx, "host")
	port := GetParamInt(ctx, "port")
	username := GetParamString(ctx, "username")
	password := GetParamString(ctx, "password")
	isSSL := GetParamInt(ctx, "is_ssl")
	// 判断参数合法性
	if emailId <= 0 {
		logger.WithContext(ctx).Warnf("[EmailModify] email_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "邮箱ID不合法")
	}
	if name == "" {
		logger.WithContext(ctx).Warnf("[EmailModify] name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "邮箱名称不能为空")
	}

	// email 邮箱实体
	emailEntity := entity.EmailEntity{
		EmailId:           emailId,
		Name:              name,
		SenderAddress:     senderAddress,
		SenderName:        senderName,
		SenderTitlePrefix: senderTitlePrefix,
		Host:              host,
		Port:              port,
		Username:          username,
		Password:          password,
		IsSSL:             isSSL,
	}

	// 更新邮箱
	err := service.NewEmail(ctx).Update(emailEntity)
	if err != nil {
		sysLogErrorf(ctx, "[EmailModify] 更新邮箱 %d 失败: err=%+v", emailId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[EmailModify] 更新邮箱 %d 成功", emailId)

	return RespJsonSuccess(ctx, nil)
}

// EmailList 邮箱列表
func EmailList(ctx *gin.Context) error {

	pageSize := GetParamIntDef(ctx, "page_size", 20)
	pageNum := GetParamIntDef(ctx, "page_num", 1)
	keywordsStr := GetParamStringDef(ctx, "keywords", "")

	var keywords *entity.EmailKeywords
	if len(keywordsStr) > 0 {
		jErr := json.Unmarshal([]byte(keywordsStr), &keywords)
		if jErr != nil {
			logger.WithContext(ctx).Errorf("[EmailList] GetEmailsByLimit err=%s", jErr.Error())
		}
	}

	emailService := service.NewEmail(ctx)

	// 获取邮箱列表
	emails, err := emailService.GetEmailsByLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[EmailList] 获取邮箱列表失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取分页信息
	pageInfo, err := emailService.GetPageInfoLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[EmailList] 获取邮箱分页信息失败: err=%s", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 格式化列表结构
	emailList, err := emailService.FormatEmailList(emails)
	if err != nil {
		sysLogErrorf(ctx, "[EmailList] 格式化邮箱列表失败: err=%s", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	data := map[string]interface{}{
		"list":      emailList,
		"page_info": pageInfo,
	}
	return RespJsonSuccess(ctx, data)
}

// EmailDelete 邮箱删除
func EmailDelete(ctx *gin.Context) error {

	emailId := GetParamInt64(ctx, "email_id")

	// 判断参数合法性
	if emailId <= 0 {
		logger.WithContext(ctx).Warnf("[EmailDelete] email_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "邮箱id不存在")
	}

	// 删除邮箱
	err := service.NewEmail(ctx).DeleteEmail(emailId)
	if err != nil {
		sysLogErrorf(ctx, "[EmailDelete] 删除邮箱 %d 失败: err=%+v", emailId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[EmailDelete] 删除邮箱 %d 成功", emailId)

	return RespJsonSuccess(ctx, nil)
}

// EmailUsed 设置邮箱为使用中
func EmailUsed(ctx *gin.Context) error {

	emailId := GetParamInt64(ctx, "email_id")

	// 判断参数合法性
	if emailId <= 0 {
		logger.WithContext(ctx).Warnf("[EmailUsed] email_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "邮箱id不存在")
	}

	// 设置邮箱为使用中
	err := service.NewEmail(ctx).SetEmailUsed(emailId)
	if err != nil {
		sysLogErrorf(ctx, "[EmailUsed] 设置邮箱 %d 使用失败: err=%+v", emailId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[EmailUsed] 设置邮箱 %d 使用成功", emailId)

	return RespJsonSuccess(ctx, nil)
}
