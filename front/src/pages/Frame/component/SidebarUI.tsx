import { Layout, Menu } from 'antd'
import { useGlobalStore } from '@/stores/index'
import { getMenuItems } from './ToolsUI'

const FrameSidebarUI = () => {
  const { iCurrentMenuItems, iMenuItemSelectedKeys, iMenuItemOpenKeys, onMenuOpenChange } =
    useGlobalStore()

  /**
   * 菜单点击方法
   */
  const onMenuItemClick = (event: any) => {}

  return (
    <Layout.Sider
      collapsible
      collapsedWidth="48px"
      width="208px"
      className="admin-sidebar"
      theme="light"
    >
      <Menu
        mode="inline"
        openKeys={iMenuItemOpenKeys}
        selectedKeys={iMenuItemSelectedKeys}
        className="admin-sidebar-menu"
        style={{
          height: `${document.body.offsetHeight - 96}px`
        }}
        items={getMenuItems(iCurrentMenuItems)}
        onClick={onMenuItemClick}
        onOpenChange={onMenuOpenChange}
      />
    </Layout.Sider>
  )
}

export default FrameSidebarUI
