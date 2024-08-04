import { getUrlConfig } from '../config/url'
import httpRequest from './http'
import { ProfileInfoType, ProfilePrivilegesResp } from '../types/profileType'

const profileUrl = {
  profilePrivileges: '/system/profile/privileges',
  profileInfo: '/system/profile/info',
  profileUpdate: '/system/profile/update',
  profileRepass: '/system/profile/repass'
}

/**
 * Profile 个人中心服务
 */
class Profile {
  /**
   * getProfilePrivileges 获取个人权限列表
   */
  getProfilePrivileges(navKey: string): Promise<ProfilePrivilegesResp> {
    let profilePrivilegesUrl = getUrlConfig().proxyUrl + profileUrl.profilePrivileges
    return httpRequest.get<ProfilePrivilegesResp>(profilePrivilegesUrl, {
      nav_key: navKey
    })
  }

  /**
   * getProfileInfo 获取个人信息
   */
  getProfileInfo(): Promise<ProfileInfoType> {
    let profileInfoUrl = getUrlConfig().proxyUrl + profileUrl.profileInfo
    return httpRequest.get<ProfileInfoType>(profileInfoUrl, {})
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
    let profileUpdateUrl = getUrlConfig().proxyUrl + profileUrl.profileUpdate
    return httpRequest.post<any>(profileUpdateUrl, {}, profileInfo)
  }

  /**
   * profileRepass 更新个人中心密码
   */
  profileRepass(passInfo: { old_pwd: string; new_pwd: string; confirm_pwd: string }): Promise<any> {
    let profileUpdateUrl = getUrlConfig().proxyUrl + profileUrl.profileRepass
    return httpRequest.post<any>(profileUpdateUrl, {}, passInfo)
  }
}

export const ProfileService = new Profile()
