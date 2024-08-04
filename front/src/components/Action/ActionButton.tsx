import { Button, Tooltip } from 'antd'
import React from 'react'

interface ActionButtonProps {
  text: string // 显示文案
  icon?: React.ReactNode // icon
  havePermission?: boolean // 是否有权限
  onClick?:
    | (React.MouseEventHandler<HTMLAnchorElement> & React.MouseEventHandler<HTMLButtonElement>)
    | undefined // 点击操作
}

/**
 * 列表操作按钮组件
 * @param props
 * @returns
 */
const ActionButton = (props: ActionButtonProps) => {
  return props.havePermission ? (
    <Button
      type="link"
      onClick={props.onClick}
      style={{
        padding: '0',
        height: 'auto',
        display: 'flex',
        alignItems: 'center'
      }}
    >
      {props.icon}
      <span className="button-text">{props.text}</span>
    </Button>
  ) : (
    <Tooltip title="禁止操作">
      <Button
        type="link"
        style={{
          padding: 0,
          height: 0,
          color: '#1677ff',
          opacity: 0.5
        }}
        disabled
      >
        {props.icon}
        <span className="button-text">{props.text}</span>
      </Button>
    </Tooltip>
  )
}

export default ActionButton
