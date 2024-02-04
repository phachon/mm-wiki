export const PrivilegeTypeNav = 1 // 权限类型：导航
export const PrivilegeTypeMenu = 2 // 权限类型：菜单
export const PrivilegeTypeOperation = 3 // 权限类型：操作

// 角色类型定义
export const PrivilegeTypes = [
  {
    type: PrivilegeTypeNav,
    name: '导航'
  },
  {
    type: PrivilegeTypeMenu,
    name: '菜单'
  },
  {
    type: PrivilegeTypeOperation,
    name: '操作'
  }
]

/**
 * PrivilegeInfoType 权限返回结构
 */
export type PrivilegeInfoType = {
  privilege_id: bigint // 权限ID
  name: string // 权限名
  parent_id: bigint // 上级权限ID
  parent_ids: string // 全部的上级权限ID逗号隔开
  privilege_type: number // 权限类型
  page_router: string // 页面路由 path
  api_marks: string // 接口标识（多个逗号隔开）
  icon: string // icon 图标
  is_display: number // 是否显示
  sequence: number // 排序数字
  create_time: string // 创建时间
  update_time: string // 修改时间
}

/**
 * PrivilegeListItemType 权限列表结构
 */
export type PrivilegeListItemType = PrivilegeInfoType & {
  child_privileges?: PrivilegeListItemType[] // 子权限
  action?: {
    is_edit: number
    is_delete: number
  }
}

/**
 * PrivilegeListResp 权限列表返回
 */
export type PrivilegeListResp = {
  list: PrivilegeListItemType[]
}

/**
 * PrivilegeEditResp 编辑权限返回结构
 */
export type PrivilegeEditResp = {
  privilege_info: PrivilegeInfoType // 权限信息
  parent_privileges: PrivilegeListItemType[] // 上级权限列表
  api_marks: string[] // 接口标识
}

/**
 * PrivilegeAddResp 添加权限返回结构
 */
export type PrivilegeAddResp = {
  parent_privileges: PrivilegeListItemType[] // 上级权限
  api_marks: string[] // 接口标识列表
}
