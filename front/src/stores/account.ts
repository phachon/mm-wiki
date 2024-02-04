import { StateCreator } from 'zustand'
import { AccountInfoType } from '../types/accountType'
import { LoginResp } from '../types/loginType'
import { IFrame } from './frame'
import { LoginTokenStore } from './local'

// 账号相关 store
export interface IAccount {
  accountInfo: AccountInfoType | undefined
  setAccountInfo: (accountInfo: AccountInfoType) => void
  setToken: (loginInfo: LoginResp) => void
}

// 创建 Account store
export const createAccount: StateCreator<IAccount & IFrame, [], [], IAccount> = (set) => ({
  accountInfo: undefined,
  setAccountInfo: (accountInfo: AccountInfoType) => set(() => ({ accountInfo: accountInfo })),
  setToken: (loginInfo: LoginResp) => {
    LoginTokenStore.storageToken(loginInfo.login_token)
  }
})
