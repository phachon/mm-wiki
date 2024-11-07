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

  // 监听 location 改变
  useEffect(() => {
    if (location.pathname === '/') {
      navigate(HOME_ROOT_PATH)
      return
    }

    // 登录校验
    if (router.auth && !LoginTokenStore.checkTokenExpire()) {
      navigate(AUTH_LOGIN_PATH)
      return
    }

    // 移除路由合法性校验，因为嵌套路由和重定向会导致误判
    // 或者使用更智能的路由匹配逻辑
    // const isExist = routerPaths.some(path => {
    //   if (path === '*') return false
    //   return location.pathname.startsWith(path)
    // })
    // if (!isExist) {
    //   navigate(NO_EXIST_PATH)
    //   return
    // }
  }, [location])

  return children
}

export default RouterInterceptor
