import httpRequest from './http'
import { getUrlConfig } from '../config/url'
import { LoginResp } from '../types/loginType'

const systemAuthUrl = {
  systemLogin: '/system/auth/login',
  domainLogin: '/system/auth/domain_login'
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
  }): Promise<LoginResp> {
    return httpRequest.post<LoginResp>(this.getSystemLoginUrl(), {}, loginInfo)
  }

  /**
   * 域账号登录
   */
  domainLogin() {}
}

export const SystemLoginService = new SystemLogin()
