package system

import (
	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/gopkg/mail"
	"github.com/phachon/mm-wiki/logger"
)

// EmailSendTest 发送测试邮件
func EmailSendTest(ctx *gin.Context) error {
	emailId := GetParamInt64(ctx, "email_id")
	toAddress := GetParamString(ctx, "to_address")

	if emailId <= 0 {
		logger.WithContext(ctx).Warnf("[EmailSendTest] email_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "邮箱id不能为空")
	}
	if toAddress == "" {
		logger.WithContext(ctx).Warnf("[EmailSendTest] to_address empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "收件人地址不能为空")
	}

	// 获取邮箱配置
	emailEntity, err := service.NewEmail(ctx).GetEmailByEmailId(emailId)
	if err != nil {
		sysLogErrorf(ctx, "[EmailSendTest] 获取邮箱 %d 失败: err=%+v", emailId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if emailEntity == nil {
		return RespJsonError(ctx, int32(errors.BusinessRecordNotExistError), "邮箱不存在")
	}

	// 检查邮箱发送功能是否开启
	configService := service.NewConfig(ctx)
	sendEmailOpen := configService.GetConfigValueByKey(entity.ConfigKeySendEmail, "0")
	if sendEmailOpen != "1" {
		return RespJsonError(ctx, int32(errors.BusinessPermissionError), "邮件发送功能未开启，请在系统配置中开启")
	}

	// 构建邮件配置
	mailConfig := &mail.MailConfig{
		Host:              emailEntity.Host,
		Port:              emailEntity.Port,
		Username:          emailEntity.Username,
		Password:          emailEntity.Password,
		SenderAddress:     emailEntity.SenderAddress,
		SenderName:        emailEntity.SenderName,
		SenderTitlePrefix: emailEntity.SenderTitlePrefix,
		IsSSL:             emailEntity.IsSSL == 1,
	}

	// 发送测试邮件
	message := &mail.MailMessage{
		To:      []string{toAddress},
		Subject: "MM-Wiki 测试邮件",
		Body:    "<h3>MM-Wiki 测试邮件</h3><p>这是一封来自 MM-Wiki 系统的测试邮件，如果您收到此邮件，说明邮箱配置正确。</p>",
		IsHTML:  true,
	}

	sendErr := mail.SendMail(mailConfig, message)
	if sendErr != nil {
		sysLogErrorf(ctx, "[EmailSendTest] 发送测试邮件失败: err=%s", sendErr.Error())
		return RespJsonError(ctx, int32(errors.ServerUnknownError), "发送测试邮件失败: "+sendErr.Error())
	}

	sysLogInfof(ctx, "[EmailSendTest] 发送测试邮件到 %s 成功", toAddress)
	return RespJsonSuccess(ctx, nil)
}
