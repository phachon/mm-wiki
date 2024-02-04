import { LandlordSexTypes } from '@/types/landlordType'
import { CheckboxOptionType, Tag } from 'antd'

/**
 * 性别单选项 Options
 * @returns options
 */
export const LandlordSexRadioOptions = (): CheckboxOptionType[] => {
  let options: CheckboxOptionType[] = []
  LandlordSexTypes.forEach((landlordSexItem) => {
    options.push({
      label: landlordSexItem.name,
      value: landlordSexItem.type
    })
  })
  return options
}

/**
 * LandlordSexTag 业主性别 tag 组件
 * @param sex 性别
 */
export const LandlordSexTag = (sex?: number) => {
  let tagLabel = <Tag color="warning">未知</Tag>
  LandlordSexTypes.forEach((landlordSexItem) => {
    if (landlordSexItem.type === sex) {
      tagLabel = <Tag color="cyan">{landlordSexItem.name}</Tag>
      return
    }
  })
  return tagLabel
}
