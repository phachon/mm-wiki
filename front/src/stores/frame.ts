import { StateCreator } from 'zustand'
import { IAccount } from './account'
import { PrivilegeListItemType, PrivilegeTypeNav } from '@/types/privilegeType'
import { INavItem, IMenuItem, IFrameBreadcrumbItem } from '@/types/frame'
import { LoginTokenStore, removeLocalAccountInfo } from './local'
import { SystemNoticeService } from '@/services/SystemNotice'
import { SystemProfileService } from '@/services/SystemProfile'
import { NoticeInfoType } from '@/types/noticeType'
import { MenuProps, message } from 'antd'

type MenuItem = Required<MenuProps>['items'][number]

/**
 * IFarme 主框架权限数据
 */
interface IPrivilegeData {
  // 导航列表数据
  iNavItems?: INavItem[]
  // 导航ID对应导航信息数据
  iNavItemsMap?: Map<string, INavItem>
  // 导航ID对应的菜单列表数据
  iNavMenuItemsMap?: Map<string, IMenuItem[]>
  // 菜单key对应的菜单信息
  iMenuItemsKeyMap?: Map<string, IMenuItem>
}

// IFarme 主框架 store 定义
export interface IFrame {
  /** 框架全局 **/
  iPrivilegeData: IPrivilegeData // 处理后的权限数据
  isLoading?: boolean // 页面全局 loading
  initProfileInfo: (pathName?: string) => void // 初始化个人信息

  /** 顶部导航 **/
  navItems: MenuItem[]
  navSelectedKeys: string[] // 导航当前选中 key 列表
  noticeList: NoticeInfoType[] // 公告列表
  noticeTotal: number // 公告总数
  onNavSelectChange: (navKey: string) => void // 导航选择操作
  onNoticeListClick: (pageSize: number, pageNum: number) => void // 点击公告列表
  onNoticeChange: (page: number) => void // 公告列表翻页
  onLogoutClick: () => void // 退出操作处理

  /** 左侧菜单 **/
  iCurrentMenuItems: IMenuItem[] // 当前菜单列表数据
  iMenuItemSelectedKeys?: string[] // 菜单当前选中 key
  iMenuItemOpenKeys?: string[] // 菜单当前展开的 SubMenu 菜单项 key 数组
  onMenuItemsClick: (menuKey: string) => void // 菜单选中回调处理
  onMenuOpenChange: (menuKeys: string[]) => void // 菜单打开回调处理

  /** 面包屑 **/
  iBreadcrumbItems: IFrameBreadcrumbItem[]
}

export const createFrame: StateCreator<IAccount & IFrame, [], [], IFrame> = (set, get) => ({
  isLoading: true,
  iPrivilegeData: {},
  iNavItems: [],
  navItems: [],
  iBreadcrumbItems: [],
  iCurrentMenuItems: [],
  iNavSelectedKey: '',
  navSelectedKeys: [],
  iMenuItemSelectedKeys: [],
  noticeList: [],
  noticeTotal: 0,

  /**
   * 进入后台后获取化个人信息
   */
  initProfileInfo: async (pathName?: string) => {
    console.log('initProfileInfo start', pathName)
    // 获取个人信息
    try {
      const profileInfo = await SystemProfileService.getProfileInfo()
      get().setAccountInfo(profileInfo.account_info)
    } catch (e) {
      console.log('initProfileInfo getProfileInfo error', e)
    }
    set({
      isLoading: false
    })
  },

  /**
   * 导航选中处理
   * @param navKey string 导航key
   */
  onNavSelectChange: (navKey: string) => {
    // let privilegeId = navKey
    // let iNavMenuItemsMap = get().iPrivilegeData?.iNavMenuItemsMap
    // let menuItems = iNavMenuItemsMap?.get(privilegeId)
    // set({
    //   navSelectedKeys: [navKey],
    //   iCurrentMenuItems: menuItems,
    //   iMenuItemSelectedKeys: []
    // })
  },

  /**
   * 菜单选中回调
   * @param menuKey 菜单key；父菜单的 key 为 id，子菜单的 key 为 path
   */
  onMenuItemsClick: (menuKey: string) => {
    console.log('onMenuItemsClick:', menuKey)
    const iPrivilegeData = get().iPrivilegeData
    if (!iPrivilegeData) {
      return
    }
    console.log('onMenuItemsClick iPrivilegeData:', iPrivilegeData)

    let iNavSelectedKey = ''
    let iMenuItemSelectedKey = ''
    let iMenuItemOpenKey = ''
    // 先根据 pathName 查找对应的菜单和导航
    if (menuKey && iPrivilegeData.iMenuItemsKeyMap) {
      const selectMenuItem = iPrivilegeData.iMenuItemsKeyMap?.get(menuKey)
      iNavSelectedKey = selectMenuItem ? selectMenuItem.navId : ''
      iMenuItemSelectedKey = selectMenuItem ? selectMenuItem.key : ''
      iMenuItemOpenKey = selectMenuItem ? selectMenuItem.parentId : ''
    }
    // 没有找到对应的导航，默认为第一个
    if (iNavSelectedKey == '' && iPrivilegeData.iNavItems && iPrivilegeData.iNavItems?.length > 0) {
      iNavSelectedKey = iPrivilegeData?.iNavItems[0].key
    }
    // 获取选中的导航对应的菜单数据
    let iMenuItems = iPrivilegeData?.iNavMenuItemsMap?.get(iNavSelectedKey)

    // console.log('onMenuItemsClick===:', iNavSelectedKey, iMenuItemSelectedKey, iMenuItemOpenKey)
    // 获取面包屑导航
    const frameBreadcrumbItems = getIBreadcrumbItems(
      iNavSelectedKey,
      iMenuItemOpenKey,
      iMenuItemSelectedKey,
      iPrivilegeData
    )
    // 更新菜单
    set({
      iCurrentMenuItems: iMenuItems,
      navSelectedKeys: [iNavSelectedKey],
      iMenuItemSelectedKeys: iMenuItemSelectedKey ? [iMenuItemSelectedKey] : [],
      iMenuItemOpenKeys: iMenuItemOpenKey ? [iMenuItemOpenKey] : [],
      iBreadcrumbItems: frameBreadcrumbItems
    })
  },

  /**
   * 菜单选择回调
   * @param menuKeys 菜单打开的 open
   */
  onMenuOpenChange: (menuKeys: string[]) => {
    const currentOpenKeys = get().iMenuItemOpenKeys ? get().iMenuItemOpenKeys : []
    const latestOpenKey = menuKeys.find((key) => currentOpenKeys?.indexOf(key) === -1)
    let iMenuItemOpenKeys = []
    if (menuKeys.length <= 1) {
      iMenuItemOpenKeys = menuKeys
    } else {
      iMenuItemOpenKeys = latestOpenKey ? [latestOpenKey] : []
    }
    set({
      iMenuItemOpenKeys: iMenuItemOpenKeys
    })
  },

  /**
   * 退出登录回调处理
   */
  onLogoutClick: () => {
    LoginTokenStore.removeToken() // 删除本地 local storage 中 token
    removeLocalAccountInfo() // 删除本地 local storage 中账号信息
    message.success('退出成功！', 1, () => {
      window.location.href = `/`
    })
  },

  onNoticeListClick: (pageSize: number, pageNum: number) => {
    SystemNoticeService.getPublishNoticeList(pageSize, pageNum)
      .then((resp) => {
        set({
          noticeList: resp.list,
          noticeTotal: resp.page_info?.total_num
        })
        console.log(resp.list)
      })
      .catch((e) => {
        console.log(e)
      })
  },

  onNoticeChange: (page: number) => {
    get().onNoticeListClick(4, page)
  }
})

/**
 * 获取处理后的权限数据
 * @param privilegeList 权限列表
 */
const getIPrivilegeData = (privilegeList?: PrivilegeListItemType[]): IPrivilegeData => {
  let iPrivilegeData: IPrivilegeData = {
    iNavItems: [],
    iNavItemsMap: new Map(),
    iNavMenuItemsMap: new Map(),
    iMenuItemsKeyMap: new Map()
  }
  if (privilegeList && privilegeList.length === 0) {
    return iPrivilegeData
  }

  privilegeList?.forEach((privilegeItem: PrivilegeListItemType) => {
    let pId = String(privilegeItem.privilege_id)
    // 判断是导航
    if (privilegeItem.privilege_type == PrivilegeTypeNav) {
      // 导航列表
      const iNavItem = {
        label: privilegeItem.name,
        key: String(privilegeItem.privilege_id),
        icon: privilegeItem.icon,
        privilegeId: String(privilegeItem.privilege_id)
      }
      iPrivilegeData.iNavItems?.push(iNavItem)

      // 获取导航的map数据
      iPrivilegeData.iNavItemsMap?.set(pId, iNavItem)
      const menuItems = getIMenuItems(pId, privilegeItem.child_privileges)
      iPrivilegeData.iNavMenuItemsMap?.set(pId, menuItems)

      // 获取菜单map数据
      const iMenuItemsPathMap = getIMenuItemsKeyMap(
        pId,
        privilegeItem.child_privileges,
        iPrivilegeData.iMenuItemsKeyMap
      )
      iPrivilegeData.iMenuItemsKeyMap = iMenuItemsPathMap
    }
  })
  return iPrivilegeData
}

/**
 * 获取处理后的菜单 map 数据
 * @param privilegeList 权限列表
 */
const getIMenuItemsKeyMap = (
  navId: string,
  menuPrivilegeList?: PrivilegeListItemType[],
  iMenuItemsKeyMap?: Map<string, IMenuItem>
): Map<string, IMenuItem> | undefined => {
  if (!menuPrivilegeList || !iMenuItemsKeyMap) {
    return iMenuItemsKeyMap
  }
  if (menuPrivilegeList.length === 0) {
    return iMenuItemsKeyMap
  }
  menuPrivilegeList?.forEach((privilegeItem) => {
    const iMenuItem: IMenuItem = {
      name: privilegeItem.name,
      pageRouter: privilegeItem.page_router,
      key: getMenuItemKey(privilegeItem),
      icon: privilegeItem.icon,
      privilegeId: String(privilegeItem.privilege_id),
      parentId: String(privilegeItem.parent_id),
      children: [],
      navId: navId
    }
    iMenuItemsKeyMap?.set(iMenuItem.key, iMenuItem)
    if (privilegeItem.child_privileges && privilegeItem.child_privileges.length > 0) {
      iMenuItemsKeyMap = getIMenuItemsKeyMap(
        navId,
        privilegeItem.child_privileges,
        iMenuItemsKeyMap
      )
    }
  })
  return iMenuItemsKeyMap
}

/**
 * 获取框架的菜单列表数据
 * @param parentId 父权限ID
 * @param privilegeList 权限列表
 */
const getIMenuItems = (parentId: string, privilegeList?: PrivilegeListItemType[]): IMenuItem[] => {
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
 * 获取面包屑导航数据
 * @param selectNavId 选中的导航ID
 * @param openMenuId 打开的菜单ID
 * @param selectSubMenuId 选中的子菜单ID
 */
const getIBreadcrumbItems = (
  selectNavId: string,
  openMenuId: string,
  selectSubMenuId: string,
  iPrivilegeData: IPrivilegeData
): IFrameBreadcrumbItem[] => {
  let frameBreadcrumbItems: IFrameBreadcrumbItem[] = []
  if (selectNavId == '') {
    return frameBreadcrumbItems
  }
  // 导航
  const iNavItems = iPrivilegeData.iNavItemsMap?.get(selectNavId)
  if (iNavItems) {
    frameBreadcrumbItems.push({
      title: iNavItems.label,
      key: iNavItems.key
    })
  }
  // 菜单
  const iMenuItem = iPrivilegeData.iMenuItemsKeyMap?.get(openMenuId)
  if (iMenuItem) {
    frameBreadcrumbItems.push({
      title: iMenuItem.name,
      key: iMenuItem.key
    })
  }
  // 子菜单
  const iSubMenuItem = iPrivilegeData.iMenuItemsKeyMap?.get(selectSubMenuId)
  if (iSubMenuItem) {
    frameBreadcrumbItems.push({
      title: iSubMenuItem.name,
      key: iSubMenuItem.key
    })
  }
  return frameBreadcrumbItems
}

// 获取菜单 item 的 key
// - 如果有子菜单，则key为权限ID
// - 如果没有子菜单，则 key 为 path
const getMenuItemKey = (privilegeItem: PrivilegeListItemType): string => {
  if (privilegeItem.child_privileges && privilegeItem.child_privileges.length > 0) {
    return String(privilegeItem.privilege_id)
  }
  if (privilegeItem.page_router) {
    return privilegeItem.page_router
  }
  return String(privilegeItem.privilege_id)
}
