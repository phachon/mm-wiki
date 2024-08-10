import { PageInfoType } from './baseType'
import { RoleInfoType } from './roleType'

// 账号状态定义
export const AccountStatusTypes = [
  {
    status: 0,
    name: '正常',
    color: 'green'
  },
  {
    status: -1,
    name: '禁用',
    color: 'error'
  }
]

/**
 * AccountInfoType 账号基础结构
 */
export type AccountInfoType = {
  account_id: bigint // 账号ID
  name: string // 账号名
  given_name: string // 昵称
  mobile: string // 手机号码
  phone: string // 电话
  email: string // 邮箱
  department: string // 部门
  position: string // 职位
  location: string // 办公位
  last_ip: string // 上次登录IP
  last_time: string // 上次登录时间
  status: number // 状态
  create_time: string // 创建时间
  update_time: string // 修改时间
  role_ids?: string // 角色ID，逗号隔开
}

/**
 * AccountListItemType 账号列表结构
 */
export type AccountListItemType = AccountInfoType & {
  roles?: RoleInfoType[] // 角色列表，逗号隔开
  action?: {
    is_edit: number // 是否可以修改
    is_detail: number // 是否查看详情
    is_update_status: number // 是否可以更新状态
  } // 列表操作
}

/**
 * AccountAddResp 添加账号返回结构
 */
export type AccountAddResp = {
  roles: RoleInfoType[] // 所有的角色
}

/**
 * AccountEditResp 编辑账号返回结构
 */
export type AccountEditResp = {
  account_info: AccountInfoType // 账号信息
  account_roles: RoleInfoType[] // 账号角色
  role_list: RoleInfoType[] // 所有的角色
}

/**
 * AccountDetailResp 账号详情返回结构
 */
export type AccountDetailResp = {
  account_info: AccountInfoType // 账号信息
  account_roles: RoleInfoType[] // 账号角色
}

/**
 * AccountListResp 账号列表返回
 */
export type AccountListResp = {
  list: AccountListItemType[]
  page_info: PageInfoType
}

/**
 * AccountSearchType 账号搜索结构定义
 */
export type AccountSearchType = {
  status: string // 状态 0 正常 -1 禁用
  account_name: string // 账号名
  given_name: string // 昵称
}
