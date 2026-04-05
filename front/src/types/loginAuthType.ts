import { PageInfoType } from './baseType'

export type LoginAuthInfoType = {
  login_auth_id: number // 认证ID
  name: string // 登录认证名称
  account_prefix: string // 账号登录前缀
  url: string // 认证接口 url
  ext_data: string // 额外数据
  is_used: number // 是否被使用 0 未使用 1 使用
  status: number // 状态 0 正常 -1 删除
  create_time: number // 创建时间
  update_time: number // 更新时间
}

export type LoginAuthListItemType = LoginAuthInfoType & {
  action?: {
    is_edit: number // 是否可修改
    is_delete: number // 是否可删除
  }
}

export type LoginAuthListResp = {
  list: LoginAuthListItemType[]
  page_info: PageInfoType
}

export type LoginAuthEditResp = {
  login_auth_info: LoginAuthInfoType
}
