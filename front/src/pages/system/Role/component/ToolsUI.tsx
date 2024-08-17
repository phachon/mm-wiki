import { RoleTypeAccountDefaultRole } from '@/types/roleType'
import { RoleInfoType, RoleTypeCustomRole, RoleTypes } from '@/types/roleType'
import { CheckboxOptionType, Space, Tag, Tooltip, Typography } from 'antd'
import React, { CSSProperties } from 'react'
import { QuestionCircleOutlined } from '@ant-design/icons'

/**
 * RoleTypeTagUI 角色类型 tag 标签显示
 * @param type 类型
 */
export const RoleTypeTagUI = (roleType?: number) => {
  let tagLabel = <Tag color="red">未知</Tag>
  RoleTypes.forEach((roleTypeItem) => {
    if (roleTypeItem.type === roleType) {
      tagLabel = <Tag color={roleTypeItem.color}>{roleTypeItem.name}</Tag>
      return
    }
  })
  return tagLabel
}

/**
 * 选择角色外显UI组件
 * @param roleInfo 角色信息
 * @returns
 */
export const SelectRoleLabelUI = (roleInfo: RoleInfoType, style?: CSSProperties | undefined) => {
  let tagLabel = '未知：' + roleInfo.name
  RoleTypes.forEach((roleTypeItem) => {
    if (roleTypeItem.type === roleInfo.role_type) {
      tagLabel = roleTypeItem.name + '：' + roleInfo.name
      return
    }
  })
  return tagLabel
}

/**
 * 添加角色选择类型 Radio Options 组件
 */
export const RoleRadioOptions = () => {
  let options: CheckboxOptionType[] = []
  RoleTypes.forEach((roleTypeItem) => {
    options.push({
      label: roleTypeItem.name,
      value: roleTypeItem.type
    })
  })
  return options
}

/**
 * 角色外显成 Tag
 * @param roles
 * @returns
 */
export const RoleTags = (roles?: RoleInfoType[]): React.ReactElement[] => {
  if (!roles) {
    return []
  }
  return roles.map((role: RoleInfoType) => <Tag color="blue">{role.name}</Tag>)
}
