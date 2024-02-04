import { PageInfoType } from './baseType'
import { LandlordInfoType } from './landlordType'

// 房产装修类型列表
export const HouseDecorationTypes = [
  {
    type: 0,
    name: '未知'
  },
  {
    type: 1,
    name: '毛胚'
  },
  {
    type: 2,
    name: '简装'
  },
  {
    type: 3,
    name: '精装'
  },
  {
    type: 4,
    name: '豪华'
  }
]

// 房产户型列表
export const HouseSizeTypes = [
  {
    type: 0,
    name: '未知'
  },
  {
    type: 1,
    name: '一室一厅'
  },
  {
    type: 2,
    name: '两室一厅'
  },
  {
    type: 3,
    name: '三室一厅'
  },
  {
    type: 4,
    name: '四室一厅'
  },
  {
    type: 5,
    name: '四室二厅'
  },
  {
    type: 6,
    name: '开间'
  }
]

// 房产出租状态列表
export const HouseLeaseStatuses = [
  {
    status: 0,
    name: '未出租',
    color_type: 'error'
  },
  {
    status: 1,
    name: '已出租',
    color_type: 'green'
  }
]

// 房产允许整租状态列表
export const HouseAllowLeaseStatuses = [
  {
    status: -1,
    name: '否',
    color_type: 'error'
  },
  {
    status: 0,
    name: '是',
    color_type: 'green'
  }
]

// 房产允许合租状态列表
export const HouseAllowSplitLeaseStatuses = [
  {
    status: -1,
    name: '否',
    color_type: 'error'
  },
  {
    status: 0,
    name: '是',
    color_type: 'green'
  }
]

// 房产区域列表
export const HouseRegions = [
  {
    name: '朝阳区',
    value: '朝阳区'
  },
  {
    name: '海淀区',
    value: '海淀区'
  },
  {
    name: '昌平区',
    value: '昌平区'
  }
]

/**
 * HouseInfoType 房产基础结构
 */
export type HouseInfoType = {
  house_id: number // 房产ID
  landlord_id: number // 业主ID
  idx: string // 唯一编号
  region: string // 区域
  address: string // 地址
  house_number: string // 门牌号
  decoration_type: number // 装修类型
  size_type: number // 户型
  area: string // 房产面积
  build_time: string // 建成日期
  status: number // 状态
  allow_lease: number // 允许整租
  allow_split_lease: number // 允许合租
  lease_status: number // 整租状态
  mouth_rent: number // 整租月租金
  create_time: string // 创建时间
  update_time: string // 修改时间
}

/**
 * HouseListItemType 房产列表结构
 */
export type HouseListItemType = HouseInfoType & {
  action?: {
    is_edit: number // 是否可以修改
    is_delete: number // 是否可以删除
  }
}

/**
 * HouseEditResp 编辑房产返回结构
 */
export type HouseEditResp = {
  house_info: HouseInfoType // 房产信息
  landlords: LandlordInfoType[] // 业主列表
}

/**
 * HouseDetailResp 房产详情返回结构
 */
export type HouseDetailResp = {
  house_info: HouseInfoType // 房产信息
}

/**
 * HouseListResp 房产列表返回
 */
export type HouseListResp = {
  list: HouseListItemType[]
  page_info: PageInfoType
}

/**
 * HouseSearchType 房产搜索结构定义
 */
export type HouseSearchType = {
  region: string // 区域
  address: string // 地址
  decoration_type: number // 装修类型
  size_type: number // 户型
  lease_status: number // 出租状态
}

/**
 * HouseTrustInfoType 房产托管数据结构
 */
export type HouseTrustInfoType = {
  house_trust_id: number // 房产托管ID
  house_id: number // 房产ID
  idx: string // 唯一编号
  start_time: string // 开始时间
  end_time: string // 结束时间
  mouth_rent: number // 月租金
  season_rent: number // 季租金
  half_year_rent: number // 半年租金
  year_rent: number // 年租金
  account_id: bigint // 操作账号
  account_name: string // 操作人
  status: number // 状态
  lease_status: number // 出租状态
  allow_lease: number // 允许出租
  create_time: string // 创建时间
  update_time: string // 修改时间
}
