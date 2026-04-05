import { PageInfoType } from './baseType'

export type ContactInfoType = {
  contact_id: number // 联系人ID
  name: string // 联系人名称
  mobile: string // 联系电话
  email: string // 邮箱
  position: string // 联系人职位
  status: number // 状态 0 正常 -1 删除
  create_time: string // 创建时间
  update_time: string // 修改时间
}

export type ContactListItemType = ContactInfoType & {
  action?: {
    is_edit: number // 是否可修改
    is_delete: number // 是否可删除
  }
}

export type ContactListResp = {
  list: ContactListItemType[]
  page_info: PageInfoType
}

export type ContactEditResp = {
  contact_info: ContactInfoType
}
