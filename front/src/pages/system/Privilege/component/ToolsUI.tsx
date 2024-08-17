import { PrivilegeTypes } from '@/types/privilegeType'
import { CheckboxOptionType } from 'antd'

/**
 * 添加权限选择类型 options 组件
 */
export const PrivilegeRadioOptions = () => {
  let options: CheckboxOptionType[] = []
  PrivilegeTypes.forEach((privilegeTypeItem) => {
    options.push({
      label: privilegeTypeItem.name,
      value: privilegeTypeItem.type
    })
  })
  return options
}
