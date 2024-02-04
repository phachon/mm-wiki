import { PageInfoType } from './baseType'

// 业主性别定义
export const LandlordSexTypes = [
  {
    type: 0,
    name: '保密'
  },
  {
    type: 1,
    name: '男'
  },
  {
    type: 2,
    name: '女'
  }
]

/**
 * LandlordInfoType 业主基础结构
 */
export type LandlordInfoType = {
  landlord_id: number // 业主ID
  nick_name: string // 昵称
  given_name: string // 姓名
  sex: number
  email: string // 邮箱
  id_card_number: string // 电话
  mobile: string // 手机号码
  address: string // 现住址
  status: number // 状态
  create_time: string // 创建时间
  update_time: string // 修改时间
}

/**
 * LandlordListItemType 业主列表结构
 */
export type LandlordListItemType = LandlordInfoType & {
  action?: {
    is_edit: number // 是否可以修改
    is_delete: number // 是否可以删除
  }
}

/**
 * LandlordSaveResp 添加保存业主返回结构
 */
export type LandlordSaveResp = {
  landlord_id: number // 业主ID
}

/**
 * LandlordEditResp 编辑业主返回结构
 */
export type LandlordEditResp = {
  landlord_info: LandlordInfoType // 业主信息
}

/**
 * LandlordListResp 业主列表返回
 */
export type LandlordListResp = {
  list: LandlordListItemType[]
  page_info: PageInfoType
}

/**
 * LandlordSearchType 业主搜索结构定义
 */
export type LandlordSearchType = {
  given_name: string // 姓名
}

/**
 * LandlordDetailResp 业主详情返回结构
 */
export type LandlordDetailResp = {
  landlord_info: LandlordInfoType // 业主信息
}
