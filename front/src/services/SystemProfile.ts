import { getUrlConfig } from '../config/url'
import httpRequest from './http'
import { ProfileInfoResp, ProfilePrivilegesResp } from '../types/profileType'

const systemProfileUrl = {
  profilePrivileges: '/system/profile/privileges',
  profileInfo: '/system/profile/info',
  profileUpdate: '/system/profile/update',
  profileRepass: '/system/profile/repass'
}

/**
 * SystemProfile 系统 - 个人中心服务
 */
class SystemProfile {
  /**
   * getProfilePrivileges 获取个人权限列表
   */
  getProfilePrivileges(navKey: string): Promise<ProfilePrivilegesResp> {
    let profilePrivilegesUrl = getUrlConfig().proxyUrl + systemProfileUrl.profilePrivileges
    return httpRequest.get<ProfilePrivilegesResp>(profilePrivilegesUrl, {
      nav_key: navKey
    })
  }

  /**
   * getProfileInfo 获取个人信息
   */
  getProfileInfo(): Promise<ProfileInfoResp> {
    let profileInfoUrl = getUrlConfig().proxyUrl + systemProfileUrl.profileInfo
    return httpRequest.get<ProfileInfoResp>(profileInfoUrl, {})
  }

  /**
   * profileUpdate 个人信息更新
   */
  profileUpdate(profileInfo: {
    name: string
    given_name: string
    email: string
    phone: string
    mobile: string
  }): Promise<any> {
    let profileUpdateUrl = getUrlConfig().proxyUrl + systemProfileUrl.profileUpdate
    return httpRequest.post<any>(profileUpdateUrl, {}, profileInfo)
  }

  /**
   * profileRepass 更新个人中心密码
   */
  profileRepass(passInfo: { old_pwd: string; new_pwd: string; confirm_pwd: string }): Promise<any> {
    let profileUpdateUrl = getUrlConfig().proxyUrl + systemProfileUrl.profileRepass
    return httpRequest.post<any>(profileUpdateUrl, {}, passInfo)
  }
}

export const SystemProfileService = new SystemProfile()
