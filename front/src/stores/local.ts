import LocalStorage from '../utils/LocalStorage'
import Token from '../utils/Token'
import { AccountInfoType } from '../types/accountType'

// LocalLoginTokenKey 登录 token 本地存储的 key
const LocalLoginTokenKey = 'MM_WIKI_LOGIN_TOKEN'
// LocalProfileAccountKey 本地存储的 key
const LocalProfileAccountKey = 'MM_WIKI_LOGIN_ACCOUNT'

export const LoginTokenStore = new Token(LocalLoginTokenKey, 10 * 60 * 60 * 1000)

// setLocalAccountInfo 存储账号信息到local
export const setLocalAccountInfo = (accountInfo: AccountInfoType) => {
  LocalStorage.setValue(LocalProfileAccountKey, accountInfo)
}

// getProfileAccountInfo 获取账号信息
export const getLocalAccountInfo = (): AccountInfoType | undefined => {
  const profileAccountInfo = LocalStorage.getValue<AccountInfoType>(LocalProfileAccountKey)
  if (profileAccountInfo == null) {
    return undefined
  }
  return profileAccountInfo
}

// removeProfileAccountInfo 清除 local 账号信息
export const removeLocalAccountInfo = (): void => {
  LocalStorage.removeValue(LocalProfileAccountKey)
}
