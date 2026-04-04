import { useEffect } from 'react'
import { useLocation, useNavigate } from 'react-router-dom'
import { routerPaths } from './paths'
import { useGlobalStore } from '@/stores'
import { LoginTokenStore } from '@/stores/local'
import { NO_ACCESS_PATH, NO_EXIST_PATH, AUTH_LOGIN_PATH, HOME_ROOT_PATH } from './type'
import { message } from 'antd'

/**
 * 检查路由路径是否合法
 * 支持嵌套路由和动态参数（:param 格式）
 */
const isValidPath = (pathname: string): boolean => {
  return routerPaths.some((path) => {
    if (path === '*') return false
    // 支持动态参数匹配: /doc/:doc_id 匹配 /doc/123
    const pathParts = path.split('/')
    const locationParts = pathname.split('/')
    if (pathParts.length !== locationParts.length) {
      // 也允许前缀匹配（嵌套路由）
      return pathname.startsWith(path) || path.startsWith(pathname)
    }
    return pathParts.every((part, index) => {
      if (part.startsWith(':')) return true
      return part === locationParts[index]
    })
  })
}

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

    // 路由合法性校验
    if (!isValidPath(location.pathname)) {
      navigate(NO_EXIST_PATH)
      return
    }
  }, [location])

  return children
}

export default RouterInterceptor
