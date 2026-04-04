import httpRequest from './http'
import { getUrlConfig } from '../config/url'
import { LoginResp } from '../types/loginType'

const systemAuthUrl = {
  systemLogin: '/system/auth/login',
  domainLogin: '/system/auth/domain_login',
  captcha: '/system/auth/captcha'
}

/**
 * CaptchaResp 验证码返回结构
 */
export type CaptchaResp = {
  captcha_id: string
  captcha_code: string
}

/**
 * SystemLogin 系统 - 登录服务
 */
class SystemLogin {
  getSystemLoginUrl(): string {
    return getUrlConfig().proxyUrl + systemAuthUrl.systemLogin
  }

  /**
   * 系统登录
   */
  systemLogin(loginInfo: {
    account_name: string
    password: string
    verify_code: string
    captcha_id: string
  }): Promise<LoginResp> {
    return httpRequest.post<LoginResp>(this.getSystemLoginUrl(), {}, loginInfo)
  }

  /**
   * 获取验证码
   */
  getCaptcha(): Promise<CaptchaResp> {
    const captchaUrl = getUrlConfig().proxyUrl + systemAuthUrl.captcha
    return httpRequest.get<CaptchaResp>(captchaUrl, {})
  }

  /**
   * 域账号登录
   */
  domainLogin() {}
}

export const SystemLoginService = new SystemLogin()
