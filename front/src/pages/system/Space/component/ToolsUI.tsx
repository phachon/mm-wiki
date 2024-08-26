import {
  SpaceIsExportTypes,
  SpaceIsShareTypes,
  SpaceTypeTypes,
  SpaceVisitLevelTypes
} from '@/types/spaceType'
import { CheckboxOptionType, Tag } from 'antd'

// 空间类型 tag 标签显示
export const SpaceTypeTagUI = (roleType?: number) => {
  let tagLabel = <Tag color="red">未知</Tag>
  SpaceTypeTypes.forEach((spaceTypeItem) => {
    if (spaceTypeItem.type === roleType) {
      tagLabel = <Tag color={spaceTypeItem.color}>{spaceTypeItem.name}</Tag>
      return
    }
  })
  return tagLabel
}

// 空间类型 Select Options 组件
export const SpaceTypeSelectOptions = () => {
  let options: CheckboxOptionType[] = []
  SpaceTypeTypes.forEach((spaceType) => {
    options.push({
      label: spaceType.name,
      value: spaceType.type
    })
  })
  return options
}

// 访问级别 tag 标签显示
export const SpaceVisitLevelTagUI = (visitLevel?: number) => {
  let tagLabel = <Tag color="red">未知</Tag>
  SpaceVisitLevelTypes.forEach((spaceVisitLevel) => {
    if (spaceVisitLevel.type === visitLevel) {
      tagLabel = <Tag color={spaceVisitLevel.color}>{spaceVisitLevel.name}</Tag>
      return
    }
  })
  return tagLabel
}

/**
 * 空间访问级别 Radio Options 组件
 */
export const SpaceVisitLevelRadioOptions = () => {
  let options: CheckboxOptionType[] = []
  SpaceVisitLevelTypes.forEach((spaceVisitLevel) => {
    options.push({
      label: spaceVisitLevel.name,
      value: spaceVisitLevel.type
    })
  })
  return options
}

/**
 * 空间是否可分享 Radio Options 组件
 */
export const SpaceIsShareRadioOptions = () => {
  let options: CheckboxOptionType[] = []
  SpaceIsShareTypes.forEach((spaceIsShare) => {
    options.push({
      label: spaceIsShare.name,
      value: spaceIsShare.type
    })
  })
  return options
}

// 空间是否可导出 Radio Options 组件
export const SpaceIsExportRadioOptions = () => {
  let options: CheckboxOptionType[] = []
  SpaceIsExportTypes.forEach((spaceIsExport) => {
    options.push({
      label: spaceIsExport.name,
      value: spaceIsExport.type
    })
  })
  return options
}

// 空间类型 Radio Options 组件
export const SpaceTypeRadioOptions = () => {
  let options: CheckboxOptionType[] = []
  SpaceTypeTypes.forEach((spaceType) => {
    options.push({
      label: spaceType.name,
      value: spaceType.type
    })
  })
  return options
}
