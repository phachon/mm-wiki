import { PageInfoType } from './baseType'
import dayjs from 'dayjs'

/**
 * NoticeInfoType 公告信息结构
 */
export type NoticeInfoType = {
  notice_id: number // 公告ID
  title: string // 公告标题
  content: string // 公告内容
  account_id: bigint // 账号ID
  account_name: string // 账号名
  status: number // 公告状态
  publish_status: number // 发布状态
  create_time: string // 创建时间
  update_time: string // 修改时间
  start_time: string // 开始时间
  end_time: string // 结束时间
  range_time: dayjs.Dayjs[] // 时间范围
}

/**
 * NoticeListItemType 公告列表结构
 */
export type NoticeListItemType = NoticeInfoType & {
  action?: {
    is_edit: number // 是否可修改
    is_delete: number // 是否可删除
  }
}

/**
 * NoticeListResp 公告列表结构
 */
export type NoticeListResp = {
  list: NoticeListItemType[]
  page_info: PageInfoType
}

/**
 * NoticeEditResp 公告修改信息返回结构
 */
export type NoticeEditResp = {
  notice_info: NoticeInfoType
}
