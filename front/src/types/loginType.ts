import { AccountInfoType } from "./accountType"

/**
 * LoginResp 登录成功返回结构
 */
export type LoginResp = {
  login_token: string // 登录 token
  account_info: AccountInfoType // 登录账号信息
}
