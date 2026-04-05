import { PageInfoType } from './baseType'

export type LinkInfoType = {
  link_id: number // 链接ID
  name: string // 链接名称
  url: string // 链接地址
  sequence: number // 排序号(越小越靠前)
  status: number // 状态 0 正常 -1 删除
  create_time: string // 创建时间
  update_time: string // 修改时间
}

export type LinkListItemType = LinkInfoType & {
  action?: {
    is_edit: number // 是否可修改
    is_delete: number // 是否可删除
  }
}

export type LinkListResp = {
  list: LinkListItemType[]
  page_info: PageInfoType
}

export type LinkEditResp = {
  link_info: LinkInfoType
}
