import { PageInfoType } from './baseType'
import { HirerInfoType } from './hirerType'
import { HouseInfoType } from './houseType'
import { RoomInfoType } from './roomType'

export const OrderStatusDefaultPending = 0 // 订单状态 0 待入住
export const OrderStatusOccupied = 1 // 订单状态 1 已入住
export const OrderStatusRenewed = 2 // 订单状态 2 已续租
export const OrderStatusVacated = 3 // 订单状态 3 已退租
export const OrderStatusFinished = 4 // 订单状态 4 已结束
export const OrderStatusCanceled = 5 // 订单状态 5 已取消

// OrderStatusTypes 订单状态定义
export const OrderStatusTypes = [
  {
    name: '待入住',
    actionName: '默认',
    value: OrderStatusDefaultPending,
    color: 'blue'
  },
  {
    name: '已入住',
    actionName: '入住',
    value: OrderStatusOccupied,
    color: 'green'
  },
  {
    name: '已续租',
    actionName: '续租',
    value: OrderStatusRenewed,
    color: 'purple'
  },
  {
    name: '已退租',
    actionName: '退租',
    value: OrderStatusVacated,
    color: 'red'
  },
  {
    name: '已结束',
    actionName: '结束',
    value: OrderStatusFinished,
    color: 'orange'
  },
  {
    name: '已取消',
    actionName: '取消',
    value: OrderStatusCanceled,
    color: 'gray'
  }
]

// OrderTypes 订单类型定义
export const OrderTypes = [
  {
    name: '整租',
    value: 1,
    color: 'green'
  },
  {
    name: '合租',
    value: 2,
    color: 'orange'
  }
]

// OrderPayTypes 订单支付类型
export const OrderPayTypes = [
  {
    name: '押一付一',
    value: 1
  },
  {
    name: '押一付三',
    value: 2
  },
  {
    name: '押一付六',
    value: 3
  },
  {
    name: '押一付年',
    value: 4
  }
]

// OrderInfoType 订单结构定义
export type OrderInfoType = {
  order_id: number // 订单ID
  house_id: number // 房屋ID
  room_id: number
  order_type: number
  hirer_id: number
  start_time: string
  end_time: string
  unit_rent: number
  pay_type: number
  account_id: number
  account_name: string
  status: number
  create_time: number
  update_time: number
}

// OrderListItemType 订单列表结构定义
export type OrderListItemType = OrderInfoType & {
  house_info: HouseInfoType // 房屋信息
  room_info: RoomInfoType // 房间信息
  hirer_info: HirerInfoType
  action?: {
    is_edit: number // 是否可以修改
  }
}

// OrderAddResp 订单添加页面数据结构
export type OrderAddResp = {
  select_houses: HouseInfoType[] // 选择房产列表
  select_house_rooms: Record<number, RoomInfoType[]> // 选择房产房间列表
}

/**
 * OrderListResp 订单列表返回
 */
export type OrderListResp = {
  list: OrderListItemType[]
  page_info: PageInfoType
}

/**
 * 订单搜索
 */
export type OrderSearchType = {
  order_id: string
  house_id: string
  hirer_id: string
  pay_type: number
  order_type: number
  order_status: number
}

/**
 * 订单详情信息
 */
export type OrderDetailResp = {
  order_info: OrderInfoType
}
