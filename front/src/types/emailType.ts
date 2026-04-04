import { PageInfoType } from './baseType'

export type EmailInfoType = {
  email_id: number // 邮箱ID
  name: string // 邮箱服务器名称
  sender_address: string // 发件人邮件地址
  sender_name: string // 发件人显示名
  sender_title_prefix: string // 发送邮件标题前缀
  host: string // 服务器主机名
  port: number // 服务器端口
  username: string // 用户名
  password: string // 密码
  is_ssl: number // 是否使用ssl 0 否 1 是
  is_used: number // 是否被使用 0 否 1 是
  status: number // 状态 0 正常 -1 删除
  create_time: string // 创建时间
  update_time: string // 修改时间
}

export type EmailListItemType = EmailInfoType & {
  action?: {
    is_edit: number // 是否可修改
    is_delete: number // 是否可删除
  }
}

export type EmailListResp = {
  list: EmailListItemType[]
  page_info: PageInfoType
}

export type EmailEditResp = {
  email_info: EmailInfoType
}
