import { AccountInfoType } from './accountType'
import { PrivilegeListItemType } from './privilegeType'

/**
 * ProfileInfoType 个人资料数据结构定义
 */
export type ProfileInfoType = {
  account_info: AccountInfoType // 账号信息
  privilege_list?: PrivilegeListItemType[] // 权限列表
}
