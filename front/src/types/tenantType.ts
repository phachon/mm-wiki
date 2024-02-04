import { PageInfoType } from './baseType'
import { PrivilegeListItemType } from './privilegeType'

export const TenantStatusDefault = 0 // 租户状态 0 正常
export const TenantStatusForbid = -1 // 租户状态 -1 禁用

export const TenantTypeDefault = 0 // 租户类型：普通租户
export const TenantTypeRoot = 1 // 租户类型：超级租户

// 租户状态定义
export const TenantStatusTypes = [
  {
    status: TenantStatusDefault,
    name: '正常',
    color: 'green'
  },
  {
    status: TenantStatusForbid,
    name: '禁用',
    color: 'red'
  }
]

// 租户类型定义
export const TenantTypes = [
  {
    type: TenantTypeDefault,
    name: '普通租户',
    color: 'blue'
  },
  {
    type: TenantTypeRoot,
    name: '超级租户',
    color: 'red'
  }
]

/**
 * TenantInfoType 租户信息
 */
export type TenantInfoType = {
  tenant_id: number // 租户ID
  name: string // 租户名
  description: string // 租户描述
  tenant_type: number // 租户类型
  phone: string // 联系电话
  start_time: string // 开始时间
  end_time: string // 结束时间
  total_price: number // 总价格
  remark: string // 备注
  status: number // 状态
  create_time: string // 创建时间
  update_time: string // 修改时间
}

/**
 * TenantInfoType 租户列表结构
 */
export type TenantListItemType = TenantInfoType & {
  action?: {
    is_edit: number // 是否可以修改
    is_update_status: number // 是否修改状态
    is_privilege_edit: number // 是否可以修改权限
  } // 操作权限
}

/**
 * TenantListResp 租户列表结构
 */
export type TenantListResp = {
  list: TenantListItemType[]
  page_info: PageInfoType
}

/**
 * TenantEditResp 租户修改信息返回结构
 */
export type TenantEditResp = {
  tenant_info: TenantInfoType
}

/**
 * TenantPrivilegeEditResp 租户权限修改返回结构
 */
export type TenantPrivilegeEditResp = {
  all_privilege: PrivilegeListItemType[]
  privilege_ids: BigInt[]
}

/**
 * TenantAddResp 添加租户信息返回结构
 */
export type TenantAddResp = {
  all_privilege: PrivilegeListItemType[]
}
