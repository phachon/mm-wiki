import { Layout, Menu } from 'antd'
import { getMenuItems } from './ToolsUI'
import { useState } from 'react'
import { IMenuItem } from '@/types/frame'

interface SystemSidebarUIProps {
  menuItemOpenKeys: string[]
  menuItemSelectedKeys: string[]
  menuItems: IMenuItem[]
  onMenuItemClick: (event: any) => void // 菜单选中回调处理
  onMenuOpenChange: (menuKeys: string[]) => void // 菜单打开回调处理
}

const SystemSidebarUI = (props: SystemSidebarUIProps) => {
  const [collapsed, setCollapsed] = useState(false)

  return (
    <Layout.Sider
      collapsible
      collapsed={collapsed}
      onCollapse={(value) => setCollapsed(value)}
      width="208px"
      className="admin-sidebar"
      theme="light"
    >
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
