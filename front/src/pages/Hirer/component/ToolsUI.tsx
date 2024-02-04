import { HirerSexTypes } from '@/types/hirerType'
import { CheckboxOptionType, Tag } from 'antd'

/**
 * 性别单选项 Options
 * @returns options
 */
export const HirerSexRadioOptions = (): CheckboxOptionType[] => {
  let options: CheckboxOptionType[] = []
  HirerSexTypes.forEach((hirerSexItem) => {
    options.push({
      label: hirerSexItem.name,
      value: hirerSexItem.type
    })
  })
  return options
}

/**
 * HirerSexTag 租客性别 tag 组件
 * @param sex 性别
 */
export const HirerSexTag = (sex?: number) => {
  let tagLabel = <Tag color="warning">未知</Tag>
  HirerSexTypes.forEach((hirerSexItem) => {
    if (hirerSexItem.type === sex) {
      tagLabel = <Tag color="cyan">{hirerSexItem.name}</Tag>
      return
    }
  })
  return tagLabel
}
