import DynamicIcon from '@/components/DynamicIcon/DynamicIcon'
import { IMenuItem, INavItem } from '@/types/frame'
import { MenuProps } from 'antd'
import { Link } from 'react-router-dom'
type MenuItem = Required<MenuProps>['items'][number]

const defaultNavItems: MenuItem[] = [
  {
    label: '空间',
    key: 'space',
    icon: <DynamicIcon name={'AppstoreOutlined'} />
  }
]

/**
 * 获取导航的 items
 * @param iNavItems
 */
export const getNavItems = (iNavItems?: INavItem[] | null): MenuItem[] => {
  const navItems: MenuItem[] = []
  iNavItems?.forEach((iNavItem: INavItem) => {
    navItems.push({
      label: iNavItem.label,
      key: iNavItem.key,
      icon: <DynamicIcon name={iNavItem.icon} />
    })
  })
  return navItems
}

/**
 * 获取菜单的 items
 * @param iMenuItmes
 */
export const getMenuItems = (iMenuItmes?: IMenuItem[] | null): MenuItem[] => {
  const menuItems: MenuItem[] = []
  iMenuItmes?.forEach((iMenuItem: IMenuItem) => {
    let children = getMenuItems(iMenuItem.children)
    let isChildren = children.length > 0 ? children : undefined
    menuItems.push({
      label: !isChildren ? (
        <Link to={iMenuItem.pageRouter} key={iMenuItem.pageRouter}>
          {iMenuItem.name}
        </Link>
      ) : (
        iMenuItem.name
      ),
      key: iMenuItem.key,
      icon: <DynamicIcon name={iMenuItem.icon} />,
      children: children.length > 0 ? children : undefined
    })
  })
  return menuItems
}
