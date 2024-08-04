import { StateCreator } from 'zustand'
import { AccountInfoType } from '../types/accountType'
import { IFrame } from './frame'
import { getLocalAccountInfo, LoginTokenStore, setLocalAccountInfo } from './local'

// 账号相关 store
export interface IAccount {
  // accountInfo: AccountInfoType | undefined
  setAccountInfo: (accountInfo: AccountInfoType) => void
  setToken: (loginToken: string) => void
  getAccountInfo: () => AccountInfoType | undefined
}

// 创建 Account store
export const createAccount: StateCreator<IAccount & IFrame, [], [], IAccount> = (set) => ({
  // accountInfo: undefined,
  setAccountInfo: (accountInfo: AccountInfoType) => {
    setLocalAccountInfo(accountInfo)
  },
  setToken: (loginToken: string) => {
    LoginTokenStore.storageToken(loginToken)
  },
  getAccountInfo: () => {
    return getLocalAccountInfo()
  }
})
