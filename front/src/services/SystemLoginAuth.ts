import httpRequest from './http'
import { LoginAuthEditResp, LoginAuthListResp } from '../types/loginAuthType'
import Base from './Base'

const systemLoginAuthUrl = {
  loginAuthSave: '/system/login_auth/save',
  loginAuthEdit: '/system/login_auth/edit',
  loginAuthModify: '/system/login_auth/modify',
  loginAuthList: '/system/login_auth/list',
  loginAuthDelete: '/system/login_auth/delete',
  loginAuthUsed: '/system/login_auth/used'
}

class SystemLoginAuth extends Base {
  public constructor() {
    super()
  }

  public saveLoginAuth(loginAuthInfo: {}): Promise<any> {
    const saveUrl = this.getProxyUrl(systemLoginAuthUrl.loginAuthSave)
    return httpRequest.post<any>(saveUrl, {}, loginAuthInfo)
  }

  public getEditLoginAuthInfo(loginAuthId: number): Promise<LoginAuthEditResp> {
    const editUrl = this.getProxyUrl(systemLoginAuthUrl.loginAuthEdit)
    return httpRequest.get<LoginAuthEditResp>(editUrl, {
      login_auth_id: loginAuthId
    })
  }

  public modifyLoginAuth(editLoginAuthInfo: {}): Promise<any> {
    const modifyUrl = this.getProxyUrl(systemLoginAuthUrl.loginAuthModify)
    return httpRequest.post<any>(modifyUrl, {}, editLoginAuthInfo)
  }

  public getLoginAuthList(
    pageSize?: number,
    pageNum?: number,
    keywords?: {}
  ): Promise<LoginAuthListResp> {
    const listUrl = this.getProxyUrl(systemLoginAuthUrl.loginAuthList)
    return httpRequest.get<LoginAuthListResp>(listUrl, {
      page_size: pageSize,
      page_num: pageNum,
      keywords: JSON.stringify(keywords)
    })
  }

  public deleteLoginAuth(loginAuthId: number): Promise<any> {
    const deleteUrl = this.getProxyUrl(systemLoginAuthUrl.loginAuthDelete)
    return httpRequest.post<any>(deleteUrl, {}, { login_auth_id: loginAuthId })
  }

  public setLoginAuthUsed(loginAuthId: number): Promise<any> {
    const usedUrl = this.getProxyUrl(systemLoginAuthUrl.loginAuthUsed)
    return httpRequest.post<any>(usedUrl, {}, { login_auth_id: loginAuthId })
  }
}

export const SystemLoginAuthService = new SystemLoginAuth()
