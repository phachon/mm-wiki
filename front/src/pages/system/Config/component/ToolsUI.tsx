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

// export const AccountStatusTag = (status?: number) => {
//   if (status === 0) {
//     return <Tag color="green">正常</Tag>
//   } else if (status === -1) {
//     return <Tag color="error">禁用</Tag>
//   } else {
//     return <Tag color="warning">未知</Tag>
//   }
// }
