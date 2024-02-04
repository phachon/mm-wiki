import { ReactNode } from 'react'

export const HOME_ROOT_PATH = '/main/index'
export const AUTH_LOGIN_PATH = '/login'
export const NO_EXIST_PATH = '/error/404'
export const NO_ACCESS_PATH = '/error/403'

/**
 * IRouter 路由定义
 */
export interface IRouter {
  path: string
  key: string
  component?: ReactNode | null
  children?: IRouter[]
  index?: boolean
  redirect?: string
  auth?: boolean // 是否需要登录校验
  permission?: boolean // 是否需要权限校验
  title?: string
}
