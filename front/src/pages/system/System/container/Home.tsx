import React, { useEffect, useState } from 'react'
import { Layout, Spin } from 'antd'
import { useGlobalStore } from '@/stores/index'
import { Outlet, useLocation } from 'react-router-dom'
import FrameSidebarUI from '../component/SidebarUI'
import '../component/home.css'
import LayoutHeader from '@/components/Layout/Header'
import { useNavigate } from 'react-router-dom'
import { LayoutHeaderSystemKey } from '@/components/Layout/types'
import { SystemProfileService } from '@/services/SystemProfile'
import { ProfilePrivilegesResp } from '@/types/profileType'
import { IMenuItem } from '@/types/frame'
import { PrivilegeListItemType } from '@/types/privilegeType'
import { SYSTEM_ROOT_PATH } from '@/router/type'

const SystemHome: React.FC = () => {
  const location = useLocation()
  const navigate = useNavigate()
  const accountInfo = useGlobalStore().getAccountInfo()
  const [isLoading, setIsLoading] = useState(true)
  const [iMenuItmes, setIMenuItmes] = useState<IMenuItem[]>([])
  const [menuItemOpenKeys, setMenuItemOpenKeys] = useState<string[]>([])
  const [menuItemSelectedKeys, setMenuItemSelectedKeys] = useState<string[]>([])

  useEffect(() => {
    initMenus()
  }, [])

  const initMenus = async () => {
    const resp: ProfilePrivilegesResp = await SystemProfileService.getProfilePrivileges(
      LayoutHeaderSystemKey
    )
    if (resp.privileges) {
      const menuItems = getIMenuItems(LayoutHeaderSystemKey, resp.privileges)
      setIMenuItmes(menuItems)
      const pathName = location.pathname
      if (pathName == SYSTEM_ROOT_PATH) {
        setIsLoading(false)
        return
      }
      // 根据当前访问的 url，找到对应的 IMenuItem 对象，定位菜单
      const findMenuItem = findMenuItemByPageRouter(menuItems, pathName)
      console.log('findMenuItem:', findMenuItem)
      if (findMenuItem) {
        setMenuItemOpenKeys([findMenuItem.parentId])
        setMenuItemSelectedKeys([findMenuItem.key])
      } else {
        navigate(SYSTEM_ROOT_PATH)
      }
      setIsLoading(false)
    }
  }

  /**
   * 菜单点击操作
   * @param event
   */
  const onMenuItemClick = (event: any) => {
    const { key, item } = event
    setMenuItemSelectedKeys([key])
  }

  /**
   * 获取菜单列表数据
   * @param parentId 父权限ID
   * @param privilegeList 权限列表
   */
  const getIMenuItems = (
    parentId: string,
    privilegeList?: PrivilegeListItemType[]
  ): IMenuItem[] => {
    let menuItems: IMenuItem[] = []
    if (privilegeList && privilegeList.length !== 0) {
      privilegeList.forEach((privilegeItem) => {
        const subMenuItems = getIMenuItems(
          String(privilegeItem.privilege_id),
          privilegeItem.child_privileges
        )
        const iMenuItem: IMenuItem = {
          name: privilegeItem.name,
          pageRouter: privilegeItem.page_router,
          key: getMenuItemKey(privilegeItem),
          icon: privilegeItem.icon,
          children: subMenuItems,
          privilegeId: String(privilegeItem.privilege_id),
          parentId: parentId,
          navId: ''
        }
        menuItems?.push(iMenuItem)
      })
    }
    return menuItems
  }

  /**
   * 递归函数，用于查找特定 pageRouter 的 IMenuItem 对象
   * @param menuItems
   * @param pageRouter
   * @returns
   */
  const findMenuItemByPageRouter = (
    menuItems: IMenuItem[],
    pageRouter: string
  ): IMenuItem | undefined => {
    for (const item of menuItems) {
      if (item.pageRouter === pageRouter) {
        return item // 找到匹配的项，返回该项
      }
      if (item.children) {
        const foundItem = findMenuItemByPageRouter(item.children, pageRouter)
        if (foundItem) {
          return foundItem
        }
      }
    }
    return undefined
  }

  const getLevelKeys = (items: IMenuItem[]) => {
    const key: Record<string, number> = {}
    const func = (items: IMenuItem[], level = 1) => {
      items.forEach((item) => {
        if (item.key) {
          key[item.key] = level
        }
        if (item.children) {
          func(item.children, level + 1)
        }
      })
    }
    func(items)
    return key
  }

  // 获取菜单 item 的 key
  const getMenuItemKey = (privilegeItem: PrivilegeListItemType): string => {
    return String(privilegeItem.privilege_id)
  }

  /**
   * 菜单打开操作
   * @param openKeys
   */
  const onMenuOpenChange = (openKeys: string[]) => {
    console.log('openKeys:', openKeys)
    const currentOpenKey = openKeys.find((key) => menuItemOpenKeys.indexOf(key) === -1)
    // open
    if (currentOpenKey !== undefined) {
      const levelKeys = getLevelKeys(iMenuItmes as IMenuItem[])
      const repeatIndex = openKeys
        .filter((key) => key !== currentOpenKey)
        .findIndex((key) => levelKeys[key] === levelKeys[currentOpenKey])
      setMenuItemOpenKeys(
        openKeys
          // remove repeat key
          .filter((_, index) => index !== repeatIndex)
          // remove current level all child
          .filter((key) => levelKeys[key] <= levelKeys[currentOpenKey])
      )
    } else {
      setMenuItemOpenKeys(openKeys) // close
    }
  }

  return (
    <Layout>
      <LayoutHeader
        {...useGlobalStore()}
        navSelectedKeys={[LayoutHeaderSystemKey]}
        accountInfo={accountInfo}
      />
      <Layout>
        <FrameSidebarUI
          menuItemSelectedKeys={menuItemSelectedKeys}
          onMenuItemClick={onMenuItemClick}
          menuItemOpenKeys={menuItemOpenKeys}
          menuItems={iMenuItmes}
          onMenuOpenChange={onMenuOpenChange}
        />
        <Layout className="admin-main">
          {/* <FrameBreadcrumbUI /> */}
          <Layout.Content className="admin-content">
            {isLoading ? (
              <Spin spinning={true} size="large" tip="Loading..." className="loading"></Spin>
            ) : (
              <Outlet />
            )}
          </Layout.Content>
          {/* <FrameFooterUI text={SettingConfig.footerShowText} /> */}
        </Layout>
      </Layout>
    </Layout>
  )
}

export default SystemHome
