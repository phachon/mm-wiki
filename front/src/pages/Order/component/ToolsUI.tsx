import { HouseShowWholeAddress } from '@/pages/House/component/ToolsUI'
import { RoomShowWholeAddress } from '@/pages/Room/component/ToolsUI'
import { HouseInfoType } from '@/types/houseType'
import {
  OrderPayTypes,
  OrderStatusCanceled,
  OrderStatusFinished,
  OrderStatusTypes,
  OrderStatusVacated,
  OrderTypes
} from '@/types/orderType'
import { RoomInfoType } from '@/types/roomType'
import { CheckboxOptionType, MenuProps, Tag, Typography } from 'antd'
import { DefaultOptionType } from 'antd/es/select'
const { Text } = Typography

/**
 * 订单类型显示标签
 */
export const OrderTypeTag = (orderType?: number) => {
  let tagLabel = <Tag color="warning">未知</Tag>
  OrderTypes.forEach((orderTypeItem) => {
    if (orderTypeItem.value === orderType) {
      tagLabel = <Tag color={orderTypeItem.color}>{orderTypeItem.name}</Tag>
      return
    }
  })
  return tagLabel
}

/**
 * 订单选择组件
 * @returns
 */
export const OrderTypeRadioOptions = (): CheckboxOptionType[] => {
  let options: CheckboxOptionType[] = []
  OrderTypes.forEach((orderType) => {
    options.push({
      label: orderType.name,
      value: orderType.value
    })
  })
  return options
}

/**
 * 支付方式选择组件
 * @returns
 */
export const OrderPayTypeRadioOptions = (): CheckboxOptionType[] => {
  let options: CheckboxOptionType[] = []
  OrderPayTypes.forEach((orderPayType) => {
    options.push({
      label: orderPayType.name,
      value: orderPayType.value
    })
  })
  return options
}

/**
 * 订单类型选择组件
 * @returns
 */
export const OrderTypeSelectOptions = (): DefaultOptionType[] => {
  let options: DefaultOptionType[] = []
  OrderTypes.forEach((orderType) => {
    options.push({
      label: orderType.name,
      value: orderType.value
    })
  })
  return options
}

/**
 * 订单状态选择组件
 * @returns
 */
export const OrderStatusSelectOptions = (): DefaultOptionType[] => {
  let options: DefaultOptionType[] = []
  OrderStatusTypes.forEach((orderStatusType) => {
    options.push({
      label: orderStatusType.name,
      value: orderStatusType.value
    })
  })
  return options
}

/**
 * 订单状态选择组件
 * @returns
 */
export const OrderStatusRadioOptions = (orderStatus?: number): CheckboxOptionType[] => {
  if (orderStatus == undefined) {
    return []
  }
  const statusNextMap = getNextStatusMap()
  const nextMap = statusNextMap.get(orderStatus)
  if (nextMap?.size == 0) {
    return []
  }
  let options: CheckboxOptionType[] = []
  OrderStatusTypes.forEach((orderStatusType) => {
    let disabled = false
    // 自身状态可修改
    if (orderStatusType.value == orderStatus) {
      disabled = false
    } else {
      disabled = nextMap?.has(orderStatusType.value) ? false : true // 下个可点击的状态
    }
    options.push({
      label: orderStatusType.name,
      value: orderStatusType.value,
      disabled: disabled
    })
  })
  return options
}

/**
 * 支付方式选择组件
 * @returns
 */
export const OrderPayTypeText = (payType?: number): string => {
  let payTypeText = '未知'
  OrderPayTypes.forEach((orderPayType) => {
    if (orderPayType.value === payType) {
      payTypeText = orderPayType.name
      return
    }
  })
  return payTypeText
}

// OrderShowHouseAddress 显示订单的房产地址
export const OrderShowHouseAddress = (houseInfo: HouseInfoType, roomInfo: RoomInfoType): string => {
  // 合租有房间显示房产 + 房间
  if (roomInfo) {
    return RoomShowWholeAddress(houseInfo, roomInfo)
  }
  // 只显示房产的地址
  return HouseShowWholeAddress(houseInfo)
}

/**
 * 订单支付类型下拉选项 Options
 * @returns options
 */
export const OrderPayTypeSelectOptions = (): DefaultOptionType[] => {
  let options: DefaultOptionType[] = []
  OrderPayTypes.forEach((orderPayTypeItem) => {
    options.push({
      label: orderPayTypeItem.name,
      value: orderPayTypeItem.value
    })
  })
  return options
}

/**
 * 订单状态文字组件
 * @returns
 */
export const OrderStatusText = (orderStatus?: number) => {
  let textLabel = (
    <Text style={{ color: 'red' }} strong={true}>
      未知
    </Text>
  )
  OrderStatusTypes.forEach((orderStatusType) => {
    if (orderStatusType.value === orderStatus) {
      textLabel = (
        <Text style={{ color: orderStatusType.color }} strong>
          {orderStatusType.name}
        </Text>
      )
      return
    }
  })
  return textLabel
}

/**
 * 订单状态文字组件
 * @returns
 */
export const OrderStatusActionItems = (orderStatus: number): MenuProps['items'] => {
  const items: MenuProps['items'] = []

  const statusNextMap = new Map<number, Map<number, boolean>>()
  // 待入住下一步只能是以结束或已取消
  const next0Map = new Map()
  next0Map.set(1, true)
  next0Map.set(5, true)
  statusNextMap.set(0, next0Map)
  // 已入住下一步只能是以结束或已取消
  const next1Map = new Map()
  next1Map.set(2, true)
  next1Map.set(3, true)
  next1Map.set(4, true)
  next1Map.set(5, true)
  statusNextMap.set(1, next1Map)
  // 已续租下一步只能是以结束或已取消
  const next2Map = new Map()
  next2Map.set(2, true)
  next2Map.set(3, true)
  next2Map.set(4, true)
  next2Map.set(5, true)
  statusNextMap.set(2, next2Map)
  // 已退租/已结束/已取消没有下一步
  const next3Map = new Map()
  statusNextMap.set(3, next3Map)
  statusNextMap.set(4, next3Map)
  statusNextMap.set(5, next3Map)

  const nextMap = statusNextMap.get(orderStatus)
  if (nextMap?.size == 0) {
    return undefined
  }

  OrderStatusTypes.forEach((orderStatusType) => {
    if (nextMap?.has(orderStatusType.value)) {
      items.push({
        label: orderStatusType.actionName,
        key: orderStatusType.value
      })
    }
  })
  return items
}

export const getNextStatusMap = (): Map<number, Map<number, boolean>> => {
  const statusNextMap = new Map<number, Map<number, boolean>>()
  // 待入住下一步只能是以结束或已取消
  const next0Map = new Map()
  next0Map.set(1, true)
  next0Map.set(5, true)
  statusNextMap.set(0, next0Map)
  // 已入住下一步只能是以结束或已取消
  const next1Map = new Map()
  next1Map.set(2, true)
  next1Map.set(3, true)
  next1Map.set(4, true)
  next1Map.set(5, true)
  statusNextMap.set(1, next1Map)
  // 已续租下一步只能是以结束或已取消
  const next2Map = new Map()
  next2Map.set(2, true)
  next2Map.set(3, true)
  next2Map.set(4, true)
  next2Map.set(5, true)
  statusNextMap.set(2, next2Map)
  // 已退租/已结束/已取消没有下一步
  const next3Map = new Map()
  statusNextMap.set(3, next3Map)
  statusNextMap.set(4, next3Map)
  statusNextMap.set(5, next3Map)

  return statusNextMap
}

export const OrderIsEdit = (status: number): boolean => {
  if (
    status == OrderStatusVacated ||
    status == OrderStatusFinished ||
    status == OrderStatusCanceled
  ) {
    return false
  }
  return true
}
