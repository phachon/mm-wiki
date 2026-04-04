package system

import (
	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/controller"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
)

// AuthCaptcha 获取验证码
func AuthCaptcha(ctx *gin.Context) error {
	captchaId := global.ContextValueRequestID(ctx)
	if captchaId == "" {
		captchaId = ctx.ClientIP()
	}
	captchaStore := service.GetCaptchaStore()
	code := captchaStore.Generate(captchaId)

	return RespJsonSuccess(ctx, map[string]interface{}{
		"captcha_id":   captchaId,
		"captcha_code": code,
	})
}

// AuthLogin 系统登录
func AuthLogin(ctx *gin.Context) error {

	accountName := ctx.PostForm("account_name")
	password := ctx.PostForm("password")
	verifyCode := ctx.PostForm("verify_code")
	captchaId := ctx.PostForm("captcha_id")

	// 删除密码
	ctx.Request.PostForm.Del("password")

	// 判断参数合法性
	if accountName == "" {
		logger.WithContext(ctx).Warnf("[AuthLogin] account_name empty")
		return controller.RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "账号名不能为空")
	}
	if password == "" {
		logger.WithContext(ctx).Warnf("[AuthLogin] password empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "密码不能为空")
	}
	if verifyCode == "" {
		logger.WithContext(ctx).Warnf("[AuthLogin] verify_code empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "验证码不能为空")
	}
	// 校验验证码
	if captchaId == "" {
		captchaId = global.ContextValueRequestID(ctx)
		if captchaId == "" {
			captchaId = ctx.ClientIP()
		}
	}
	captchaStore := service.GetCaptchaStore()
	if !captchaStore.Verify(captchaId, verifyCode) {
		logger.WithContext(ctx).Warnf("[AuthLogin] verify_code invalid, captchaId=%s", captchaId)
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "验证码错误或已过期")
	}

	authService := service.NewAuth(ctx)
	// 判断账号和密码是否正确
	loginToken, accountInfo, err := authService.Login(accountName, password)
	if err != nil {
		sysLogErrorf(ctx, "[AuthLogin] 账号 %s 登录失败 err=%+v", accountName, err)
		return RespJsonError(ctx, err.GetErrCode(), "登录失败：账号或密码错误！")
	}

	// 登录成功后设置到 Context 中
	global.ContextWithLoginAccountID(ctx, accountInfo.AccountId)
	global.ContextWithLoginAccountName(ctx, accountInfo.Name)

	sysLogInfof(ctx, "[AuthLogin] 账号 %s 登录成功", accountName)

	return RespJsonSuccess(ctx, map[string]interface{}{
		"login_token":  loginToken,
		"account_info": accountInfo,
	})
}
