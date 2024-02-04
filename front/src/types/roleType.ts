import { PageInfoType } from './baseType'
import { PrivilegeListItemType } from './privilegeType'

export const RoleTypeCustomRole = 0 // 自定义角色
export const RoleTypeAccountDefaultRole = 1 // 账号默认角色

// 角色类型定义
export const RoleTypes = [
  {
    type: RoleTypeCustomRole,
    name: '自定义',
    color: 'blue'
  },
  {
    type: RoleTypeAccountDefaultRole,
    name: '默认',
    color: 'red'
  }
]

/**
 * RoleInfoType 角色信息
 */
export type RoleInfoType = {
  role_id: number // 角色ID
  name: string // 角色名
  remark: string // 角色备注
  status: number // 状态
  role_type: number // 角色类型
  create_time: string // 创建时间
  update_time: string // 修改时间
}

/**
 * RoleInfoType 角色列表结构
 */
export type RoleListItemType = RoleInfoType & {
  action?: {
    is_edit: number // 是否可以修改
    is_delete: number // 是否可以删除
    is_account_list: number // 是否查看查看账号列表
    is_privilege_edit: number // 是否可以修改权限
  } // 操作权限
}

/**
 * RoleListResp 角色列表结构
 */
export type RoleListResp = {
  list: RoleListItemType[]
  page_info: PageInfoType
}

/**
 * RoleEditResp 角色修改信息返回结构
 */
export type RoleEditResp = {
  role_info: RoleInfoType
}

/**
 * RolePrivilegeEditResp 角色权限修改返回结构
 */
export type RolePrivilegeEditResp = {
  all_privilege: PrivilegeListItemType[]
  privilege_ids: BigInt[]
}

/**
 * RoleAddResp 添加角色信息返回结构
 */
export type RoleAddResp = {
  all_privilege: PrivilegeListItemType[]
}
