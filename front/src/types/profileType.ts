import { AccountInfoType } from './accountType'
import { PrivilegeListItemType } from './privilegeType'

// ProfileInfoResp /profile/info 返回结构
export type ProfileInfoResp = {
  account_info: AccountInfoType // 账号信息
}

// ProfilePrivilegesResp /profile/privileges 返回结构
export type ProfilePrivilegesResp = {
  privileges?: PrivilegeListItemType[] // 权限列表
}