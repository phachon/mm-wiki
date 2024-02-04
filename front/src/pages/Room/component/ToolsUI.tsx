import { HouseInfoType } from '@/types/houseType'
import {
  RoomAllowLeaseTypes,
  RoomBalconyTypes,
  RoomDirectionTypes,
  RoomInfoType,
  RoomLeaseStatusTypes,
  RoomRoomTypes,
  RoomToiletTypes
} from '@/types/roomType'
import { CheckboxOptionType, Tag } from 'antd'
import { DefaultOptionType } from 'antd/es/select'

/**
 * 房间出租状态下拉选择 Options
 * @returns options
 */
export const RoomLeaseStatusSelectOptions = (): DefaultOptionType[] => {
  let options: DefaultOptionType[] = []
  RoomLeaseStatusTypes.forEach((roomLeaseStatus) => {
    options.push({
      label: roomLeaseStatus.name,
      value: roomLeaseStatus.value
    })
  })
  return options
}

/**
 * 房间允许出租状态下拉选择 Options
 * @returns options
 */
export const RoomAllowLeaseSelectOptions = (): DefaultOptionType[] => {
  let options: DefaultOptionType[] = []
  RoomAllowLeaseTypes.forEach((roomAllowLease) => {
    options.push({
      label: roomAllowLease.name,
      value: roomAllowLease.value
    })
  })
  return options
}

/**
 * 朝向单选项 Options
 * @returns options
 */
export const RoomDirectionTypeRadioOptions = (): CheckboxOptionType[] => {
  let options: CheckboxOptionType[] = []
  RoomDirectionTypes.forEach((directionTypeItem) => {
    options.push({
      label: directionTypeItem.name,
      value: directionTypeItem.value
    })
  })
  return options
}

/**
 * 朝向下拉选项 Options
 * @returns options
 */
export const RoomDirectionTypeSelectOptions = (): DefaultOptionType[] => {
  let options: DefaultOptionType[] = []
  RoomDirectionTypes.forEach((directionTypeItem) => {
    options.push({
      label: directionTypeItem.name,
      value: directionTypeItem.value
    })
  })
  return options
}

/**
 * 朝向显示标签
 */
export const RoomDirectionTypeTag = (directionType?: number) => {
  let tagLabel = <Tag color="warning">未知</Tag>
  RoomDirectionTypes.forEach((roomDirectionType) => {
    if (roomDirectionType.value === directionType) {
      tagLabel = <Tag color={roomDirectionType.color}>{roomDirectionType.name}</Tag>
      return
    }
  })
  return tagLabel
}

/**
 * 类型显示标签
 */
export const RoomTypeTag = (roomType?: number) => {
  let tagLabel = <Tag color="warning">未知</Tag>
  RoomRoomTypes.forEach((roomRoomType) => {
    if (roomRoomType.value === roomType) {
      tagLabel = <Tag color={roomRoomType.color}>{roomRoomType.name}</Tag>
      return
    }
  })
  return tagLabel
}

/**
 * 房间类型单选项 Options
 * @returns options
 */
export const RoomRoomTypeRadioOptions = (): CheckboxOptionType[] => {
  let options: CheckboxOptionType[] = []
  RoomRoomTypes.forEach((roomTypeItem) => {
    options.push({
      label: roomTypeItem.name,
      value: roomTypeItem.value
    })
  })
  return options
}

/**
 * 房间类型下拉选项 Options
 * @returns options
 */
export const RoomRoomTypeSelectOptions = (): DefaultOptionType[] => {
  let options: DefaultOptionType[] = []
  RoomRoomTypes.forEach((roomTypeItem) => {
    options.push({
      label: roomTypeItem.name,
      value: roomTypeItem.value
    })
  })
  return options
}

// RoomLeaseStatusTag 房间出租状态
export const RoomLeaseStatusTag = (leaseStatus?: number) => {
  let tagLabel = <Tag color="warning">未知</Tag>
  RoomLeaseStatusTypes.forEach((roomLeaseStatus) => {
    if (roomLeaseStatus.value === leaseStatus) {
      tagLabel = <Tag color={roomLeaseStatus.color}>{roomLeaseStatus.name}</Tag>
      return
    }
  })
  return tagLabel
}

// RoomAllowLeaseTag 房间是否允许出租状态
export const RoomAllowLeaseTag = (allowLease?: number) => {
  let tagLabel = <Tag color="warning">未知</Tag>
  RoomAllowLeaseTypes.forEach((roomAllowLease) => {
    if (roomAllowLease.value === allowLease) {
      tagLabel = <Tag color={roomAllowLease.color}>{roomAllowLease.name}</Tag>
      return
    }
  })
  return tagLabel
}

// RoomToiletStatusTag 房间是否有卫生间tag
export const RoomToiletStatusTag = (hasToilet?: number) => {
  let tagLabel = <Tag color="warning">未知</Tag>
  RoomToiletTypes.forEach((roomToiletType) => {
    if (roomToiletType.value === hasToilet) {
      tagLabel = <Tag color={roomToiletType.color}>{roomToiletType.name}</Tag>
      return
    }
  })
  return tagLabel
}

// RoomBalconyStatusTag 房间是否有阳台 tag
export const RoomBalconyStatusTag = (hasToilet?: number) => {
  let tagLabel = <Tag color="warning">未知</Tag>
  RoomBalconyTypes.forEach((roomBalconyType) => {
    if (roomBalconyType.value === hasToilet) {
      tagLabel = <Tag color={roomBalconyType.color}>{roomBalconyType.name}</Tag>
      return
    }
  })
  return tagLabel
}

// RoomShowAllTag 房间外显所有的配置 tag
export const RoomShowAllTag = (roomInfo: RoomInfoType) => {
  let roomTypeTag = RoomTypeTag(roomInfo.room_type)
  let roomDirectionTag = RoomDirectionTypeTag(roomInfo.direction_type)
  let roomHasToiletTag = <></>
  let roomHasBalconyTag = <></>
  if (roomInfo.has_toilet == 1) {
    roomHasToiletTag = <Tag color={'green'}>带独卫</Tag>
  }
  if (roomInfo.has_balcony == 1) {
    roomHasBalconyTag = <Tag color={'blue'}>带阳台</Tag>
  }
  return (
    <div>
      {roomTypeTag} {roomDirectionTag} {roomHasToiletTag} {roomHasBalconyTag}
    </div>
  )
}

// RoomShowWholeAddress 房间外显的整个地址
export const RoomShowWholeAddress = (houseInfo: HouseInfoType, roomInfo: RoomInfoType): string => {
  return (
    houseInfo.region + '-' + houseInfo.address + '-' + houseInfo.house_number + '-' + roomInfo.name
  )
}
