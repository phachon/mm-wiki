import { Layout, Menu } from 'antd'
import { getMenuItems } from './ToolsUI'
import { useState } from 'react'
import { IMenuItem } from '@/types/frame'
import { RightOutlined, LeftOutlined } from '@ant-design/icons'
import { SiderWidth } from '@/config/layout'

interface SystemSidebarUIProps {
  menuItemOpenKeys: string[]
  menuItemSelectedKeys: string[]
  menuItems: IMenuItem[]
  onMenuItemClick: (event: any) => void // 菜单选中回调处理
  onMenuOpenChange: (menuKeys: string[]) => void // 菜单打开回调处理
}

const SystemSidebarUI = (props: SystemSidebarUIProps) => {
  const [collapsed, setCollapsed] = useState(false)

  const toggleCollapsed = () => {
    setCollapsed(!collapsed)
  }

  return (
    <Layout.Sider
      collapsed={collapsed}
      onCollapse={toggleCollapsed}
      width={SiderWidth}
      collapsedWidth="64px"
      className="admin-sidebar"
      theme="light"
      trigger={null} // 自定义触发器
    >
      <div className="admin-sidebar-trigger" onClick={toggleCollapsed}>
        {collapsed ? <RightOutlined /> : <LeftOutlined />}
      </div>
      <Menu
        mode="inline"
        openKeys={props.menuItemOpenKeys}
        selectedKeys={props.menuItemSelectedKeys}
        className="admin-sidebar-menu"
        style={{
          height: `${document.body.offsetHeight - 96}px`
        }}
        items={getMenuItems(props.menuItems)}
        onClick={props.onMenuItemClick}
        onOpenChange={props.onMenuOpenChange}
      />
    </Layout.Sider>
  )
}

export default SystemSidebarUI
