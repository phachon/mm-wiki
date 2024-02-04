import { PageInfoType } from './baseType'
import { HouseInfoType } from './houseType'

// RoomDirectionTypes 房间朝向类型定义
export const RoomDirectionTypes = [
  {
    value: 0,
    name: '未知',
    color: 'error'
  },
  {
    value: 1,
    name: '朝南',
    color: 'blue'
  },
  {
    value: 2,
    name: '朝东',
    color: 'green'
  },
  {
    value: 3,
    name: '朝西',
    color: 'green'
  },
  {
    value: 4,
    name: '朝北',
    color: 'green'
  }
]

// RoomLeaseStatusTypes 房间出租状态定义
export const RoomLeaseStatusTypes = [
  {
    value: 0,
    name: '未出租',
    color: 'error'
  },
  {
    value: 1,
    name: '已出租',
    color: 'green'
  }
]

// RoomAllowLeaseTypes 房间允许出租状态定义
export const RoomAllowLeaseTypes = [
  {
    value: 0,
    name: '否',
    color: 'error'
  },
  {
    value: 1,
    name: '是',
    color: 'green'
  }
]

// RoomRoomTypes 房间类型定义
export const RoomRoomTypes = [
  {
    value: 0,
    name: '次卧',
    color: 'error'
  },
  {
    value: 1,
    name: '主卧',
    color: 'green'
  },
  {
    value: 2,
    name: '次卧',
    color: 'green'
  },
  {
    value: 3,
    name: '客厅',
    color: 'green'
  }
]

// RoomToiletTypes 房间是否有卫生间状态定义
export const RoomToiletTypes = [
  {
    value: 0,
    name: '无',
    color: 'error'
  },
  {
    value: 1,
    name: '有',
    color: 'green'
  }
]

// RoomBalconyTypes 房间是否有阳台状态定义
export const RoomBalconyTypes = [
  {
    value: 0,
    name: '无',
    color: 'error'
  },
  {
    value: 1,
    name: '有',
    color: 'green'
  }
]

/**
 * RoomInfoType 房间基础结构
 */
export type RoomInfoType = {
  room_id: bigint // 房间ID
  house_id: bigint // 房产id
  name: string // 房间名
  area: number // 面积
  direction_type: number // 朝向
  room_type: number // 房间类型
  has_toilet: number // 是否有卫生间
  has_balcony: number // 是否有阳台
  mouth_rent: number // 月租金
  status: number // 状态
  lease_status: number // 出租状态
  allow_lease: number // 是否允许出租
  create_time: string // 创建时间
  update_time: string // 修改时间
}

/**
 * RoomListItemType 房间列表结构
 */
export type RoomListItemType = RoomInfoType & {
  house_info: HouseInfoType // 房产信息
  action?: {
    is_edit: number // 是否可以修改
    is_delete: number // 是否可以删除
  }
}

/**
 * RoomAddResp 添加房间返回结构
 */
export type RoomAddResp = {
  houses: HouseInfoType[] // 房产列表
}

/**
 * RoomEditResp 编辑房间返回结构
 */
export type RoomEditResp = {
  room_info: RoomInfoType // 房间信息
  houses: HouseInfoType[] // 房产列表
}

/**
 * RoomDetailResp 房间详情返回结构
 */
export type RoomDetailResp = {
  room_info: RoomInfoType // 房间信息
}

/**
 * RoomListResp 房间列表返回
 */
export type RoomListResp = {
  list: RoomListItemType[]
  page_info: PageInfoType
}

/**
 * RoomSearchType 房间搜索结构定义
 */
export type RoomSearchType = {
  name: string // 房间名
  direction_type: number // 房间朝向
  room_type: number // 房间类型
  lease_status: number // 出租状态
  allow_lease: number // 允许出租
  house_id: string // 房间ID
}
