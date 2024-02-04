import LocalStorage from '../utils/LocalStorage'
import Token from '../utils/Token'
import { AccountInfoType } from '../types/accountType'

// LocalLoginTokenKey 登录 token 本地存储的 key
const LocalLoginTokenKey = 'KMS_ADMIN_LOGIN_TOKEN'
// LocalProfileAccountKey 本地存储的 key
const LocalProfileAccountKey = 'KMS_ADMIN_PROFILE_ACCOUNT'

export const LoginTokenStore = new Token(LocalLoginTokenKey, 10 * 60 * 60 * 1000)

// setProfileAccountInfo 存储 profile 账号信息
export const setProfileAccountInfo = (accountInfo: AccountInfoType) => {
  LocalStorage.setValue(LocalProfileAccountKey, accountInfo)
}

// getProfileAccountInfo 获取 profile 账号信息
export const getProfileAccountInfo = (): AccountInfoType | undefined => {
  const profileAccountInfo = LocalStorage.getValue<AccountInfoType>(LocalProfileAccountKey)
  if (profileAccountInfo == null) {
    return undefined
  }
  return profileAccountInfo
}

// removeProfileAccountInfo 清除 profile 账号信息
export const removeProfileAccountInfo = (): void => {
  LocalStorage.removeValue(LocalProfileAccountKey)
}
