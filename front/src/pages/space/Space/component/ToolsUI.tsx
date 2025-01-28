import { SpacePermissionTypes } from '@/types/spaceType'
import { SelectProps, Space, Tag } from 'antd'

type TagRender = SelectProps['tagRender']

// 空间权限展示标签
export const SpacePermissionShowTagUI = (props: { premissionTypes: number[] }) => {
  return (
    <div>
      {props.premissionTypes.map((premissionType) => {
        const premission = SpacePermissionTypes.find((item) => item.type === premissionType)
        return (
          <Tag key={premissionType} color={premission?.color}>
            {premission?.name}
          </Tag>
        )
      })}
    </div>
  )
}

// 空间权限选择标签渲染
export const SpacePermissionSelectTagRender: TagRender = (props) => {
  const { label, value, closable, onClose } = props
  const onPreventMouseDown = (event: React.MouseEvent<HTMLSpanElement>) => {
    event.preventDefault()
    event.stopPropagation()
  }
  return (
    <Tag
      // color={'blue'}
      onMouseDown={onPreventMouseDown}
      closable={closable}
      onClose={onClose}
      style={{ marginInlineEnd: 4 }}
    >
      {label}
    </Tag>
  )
}
