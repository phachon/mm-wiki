import { PageInfoType } from './baseType'

// 租客性别定义
export const HirerSexTypes = [
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
 * HirerInfoType 租客基础结构
 */
export type HirerInfoType = {
  hirer_id: number // 租客ID
  name: string // 姓名
  sex: number
  email: string // 邮箱
  id_card_number: string // 电话
  mobile: string // 手机号码
  emergency_contact: string // 紧急联系人
  emergency_mobile: string // 紧急联系电话
  status: number // 状态
  create_time: string // 创建时间
  update_time: string // 修改时间
}

/**
 * HirerListItemType 业主列表结构
 */
export type HirerListItemType = HirerInfoType & {
  action?: {
    is_edit: number // 是否可以修改
    is_delete: number // 是否可以删除
  }
}

/**
 * HirerSaveResp 添加保存租客返回结构
 */
export type HirerSaveResp = {
  hirer_id: number // 租客ID
}

/**
 * HirerEditResp 编辑租客返回结构
 */
export type HirerEditResp = {
  hirer_info: HirerInfoType // 租客信息
}

/**
 * HirerListResp 租客列表返回
 */
export type HirerListResp = {
  list: HirerListItemType[]
  page_info: PageInfoType
}

/**
 * HirerSearchType 租客搜索结构定义
 */
export type HirerSearchType = {
  name: string // 姓名
}

/**
 * HirerDetailResp 租客详情返回结构
 */
export type HirerDetailResp = {
  hirer_info: HirerInfoType // 租客信息
}
