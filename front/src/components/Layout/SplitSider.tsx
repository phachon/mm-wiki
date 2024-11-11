import React, { useState } from 'react'
import { Layout, Splitter } from 'antd'
import { RightOutlined, LeftOutlined } from '@ant-design/icons'
import './splitSider.css'
import { SiderWidth } from '@/config/layout'

type LayoutSiderProps = {
  left?: React.ReactNode
  right?: React.ReactNode
}

const LayoutSplitSider = (props: LayoutSiderProps) => {
  const [collapsed, setCollapsed] = useState(false)
  const [width, setWidth] = useState(SiderWidth)

  const toggleCollapsed = () => {
    setCollapsed(!collapsed)
    setWidth(collapsed ? SiderWidth : 0)
  }

  return (
    <Splitter className="split-sider" style={{ height: '100%' }}>
      <Splitter.Panel
        className="split-sider-panel"
        collapsible
        defaultSize={SiderWidth}
        min={64}
        max={'50%'}
      >
        <div className="split-sider-content">
          <div className="split-sider-inner">{props?.left}</div>
        </div>
      </Splitter.Panel>
      <Splitter.Panel>{props?.right}</Splitter.Panel>
    </Splitter>
  )
}

export default LayoutSplitSider
