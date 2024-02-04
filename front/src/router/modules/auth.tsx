import { IRouter, AUTH_LOGIN_PATH } from '../type'
import Login from '@/pages/Login'
import Test from '@/pages/Error/Test'

const routers: IRouter[] = [
  {
    path: AUTH_LOGIN_PATH,
    key: 'auth_login',
    auth: false,
    title: '登录',
    component: <Login />
  },
  {
    path: '/test',
    key: '/test',
    component: <Test />,
    auth: false,
    permission: false
  }
]

export default routers
