import { AccountStatusTypes } from '@/types/accountType'
import { Tag } from 'antd'
import { DefaultOptionType } from 'antd/es/select'

/**
 * AccountStatusTag 账号状态的 tag 标签显示
 * @param status 状态
 */
export const AccountStatusTag = (status?: number) => {
  let tagLabel = <Tag color="warning">未知</Tag>
  AccountStatusTypes.forEach((accountStatusItem) => {
    if (accountStatusItem.status === status) {
      tagLabel = <Tag color={accountStatusItem.color}>{accountStatusItem.name}</Tag>
      return
    }
  })
  return tagLabel
}

/**
 * 账号状态下拉选择 Options
 * @returns options
 */
export const AccountStatusSelectOptions = (): DefaultOptionType[] => {
  let options: DefaultOptionType[] = []
  AccountStatusTypes.forEach((accountTypeItem) => {
    options.push({
      label: accountTypeItem.name,
      value: String(accountTypeItem.status)
    })
  })
  return options
}

// AccountDepartmentFullName 账号部门全称
export const AccountDepartmentFullName = (departmentNames?: string[]): string => {
  if (!departmentNames || departmentNames.length == 0) {
    return '未知部门'
  }
  let departmentFullName = ''
  departmentNames.forEach((departmentName, index) => {
    departmentFullName += departmentName
    if (index < departmentNames.length - 1) {
      departmentFullName += ' / '
    }
  })
  return departmentFullName
}
