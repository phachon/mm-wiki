import React, { useState } from 'react'
import { Layout } from 'antd'
import { ResizableBox } from 'react-resizable'
import { RightOutlined, LeftOutlined } from '@ant-design/icons'

const HomeSidebarUI = () => {
  const [collapsed, setCollapsed] = useState(false)
  const [width, setWidth] = useState(208)

  const toggleCollapsed = () => {
    setCollapsed(!collapsed)
    setWidth(collapsed ? 208 : 0)
  }

  const handleResize = (e: React.SyntheticEvent, data: any) => {
    setWidth(data.size.width)
  }

  return (
    <ResizableBox
      width={width}
      height={Infinity}
      axis="x"
      minConstraints={[0, Infinity]}
      maxConstraints={[800, Infinity]}
      onResize={handleResize}
      className="resizable-sidebar"
    >
      <Layout.Sider
        collapsed={collapsed}
        width={width}
        collapsedWidth="0"
        className="home-sidebar"
        theme="light"
        trigger={null} // 自定义触发器
      />
      <div className="home-sidebar-trigger" onClick={toggleCollapsed}>
        {collapsed ? <RightOutlined /> : <LeftOutlined />}
      </div>
    </ResizableBox>
  )
}

export default HomeSidebarUI
