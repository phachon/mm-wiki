import {
  HouseAllowLeaseStatuses,
  HouseAllowSplitLeaseStatuses,
  HouseDecorationTypes,
  HouseInfoType,
  HouseLeaseStatuses,
  HouseRegions,
  HouseSizeTypes
} from '@/types/houseType'
import { Tag } from 'antd'
import { DefaultOptionType } from 'antd/es/select'

/**
 * HouseLeaseStatusTag 出租状态标签UI组件
 * @param leaseStatus 状态
 */
export const HouseLeaseStatusTag = (leaseStatus?: number) => {
  let tagLabel = <Tag color="warning">未知</Tag>
  HouseLeaseStatuses.forEach((houseLeaseStatus) => {
    if (houseLeaseStatus.status === leaseStatus) {
      tagLabel = <Tag color={houseLeaseStatus.color_type}>{houseLeaseStatus.name}</Tag>
      return
    }
  })
  return tagLabel
}

/**
 * HouseAllowLeaseTag 房产允许整租状态标签UI组件
 * @param leaseStatus 状态
 */
export const HouseAllowLeaseTag = (allowLease?: number) => {
  let tagLabel = <Tag color="warning">未知</Tag>
  HouseAllowLeaseStatuses.forEach((houseAllowLeaseStatus) => {
    if (houseAllowLeaseStatus.status === allowLease) {
      tagLabel = <Tag color={houseAllowLeaseStatus.color_type}>{houseAllowLeaseStatus.name}</Tag>
      return
    }
  })
  return tagLabel
}

/**
 * HouseAllowSplitLeaseTag 房产允许合租状态标签UI组件
 * @param leaseStatus 状态
 */
export const HouseAllowSplitLeaseTag = (allowLease?: number) => {
  let tagLabel = <Tag color="warning">未知</Tag>
  HouseAllowSplitLeaseStatuses.forEach((houseAllowLeaseStatus) => {
    if (houseAllowLeaseStatus.status === allowLease) {
      tagLabel = <Tag color={houseAllowLeaseStatus.color_type}>{houseAllowLeaseStatus.name}</Tag>
      return
    }
  })
  return tagLabel
}

/**
 * 装修类型文字组件
 * @param decorationType 装修
 */
export const HouseDecorationTypeText = (decorationType?: number) => {
  let text = '未知'
  HouseDecorationTypes.forEach((decorationTypeItem) => {
    if (decorationTypeItem.type === decorationType) {
      text = decorationTypeItem.name
      return
    }
  })
  return text
}

/**
 * 房屋装修类型下拉选择 Options
 * @returns options
 */
export const HouseDecorationTypeSelectOptions = (): DefaultOptionType[] => {
  let options: DefaultOptionType[] = []
  HouseDecorationTypes.forEach((decorationTypeItem) => {
    options.push({
      label: decorationTypeItem.name,
      value: decorationTypeItem.type
    })
  })
  return options
}

/**
 * 房屋户型文字组件
 * @param sizeType 户型
 */
export const HouseSizeTypeText = (sizeType?: number) => {
  let text = '未知'
  HouseSizeTypes.forEach((sizeTypeItem) => {
    if (sizeTypeItem.type === sizeType) {
      text = sizeTypeItem.name
      return
    }
  })
  return text
}

/**
 * 房屋户型下拉选择 Options
 * @returns options
 */
export const HouseSizeTypeSelectOptions = (): DefaultOptionType[] => {
  let options: DefaultOptionType[] = []
  HouseSizeTypes.forEach((sizeTypeItem) => {
    options.push({
      label: sizeTypeItem.name,
      value: sizeTypeItem.type
    })
  })
  return options
}

/**
 * 房屋装修类型下拉选择 Options
 * @returns options
 */
export const HouseLeaseStatusSelectOptions = (): DefaultOptionType[] => {
  let options: DefaultOptionType[] = []
  HouseLeaseStatuses.forEach((leaseStatusItem) => {
    options.push({
      label: leaseStatusItem.name,
      value: leaseStatusItem.status
    })
  })
  return options
}

/**
 * 房屋区域下拉选择 Options
 * @returns options
 */
export const HouseRegionSelectOptions = (): DefaultOptionType[] => {
  let options: DefaultOptionType[] = []
  HouseRegions.forEach((houseRegionItem) => {
    options.push({
      label: houseRegionItem.name,
      value: houseRegionItem.value
    })
  })
  return options
}

// HouseShowWholeAddress 房产外显的整个地址
export const HouseShowWholeAddress = (houseInfo: HouseInfoType): string => {
  return houseInfo.region + '-' + houseInfo.address + '-' + houseInfo.house_number
}
