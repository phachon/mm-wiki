import { useEffect } from 'react'
import { useLocation, useNavigate } from 'react-router-dom'
import { routerPaths } from './paths'
import { useGlobalStore } from '@/stores'
import { LoginTokenStore } from '@/stores/local'
import { NO_ACCESS_PATH, NO_EXIST_PATH, AUTH_LOGIN_PATH, HOME_ROOT_PATH } from './type'
import { message } from 'antd'

/**
 * 路由拦截器
 */
export const RouterInterceptor = ({ children, router }: any) => {
  const location = useLocation()
  const navigate = useNavigate()
  const { iPrivilegeData, onMenuItemsClick } = useGlobalStore()

  // 监听 location 改变
  useEffect(() => {
    if (location.pathname == '/') {
      navigate(HOME_ROOT_PATH)
      return
    }
    console.log('router:', router)
    console.log('location:', location)
    // // 路由合法性校验
    const isExist = routerPaths.indexOf(location.pathname) != -1
    if (!isExist) {
      navigate(NO_EXIST_PATH)
      return
    }
    // 路由登录校验
    if (router.auth && !LoginTokenStore.checkTokenExpire()) {
      // todo 这里会执行两次
      navigate(AUTH_LOGIN_PATH)
      return
    }
    // 路由权限校验
    // const iMenuItem = iPrivilegeData.iMenuItemsKeyMap?.get(router.path)
    // if (router.permission && !iMenuItem) {
    //   navigate(NO_ACCESS_PATH)
    //   return
    // }
    // 路由改变，改变菜单的选中态
    // onMenuItemsClick(location.pathname)
  }, [location])

  return children
}

export default RouterInterceptor
