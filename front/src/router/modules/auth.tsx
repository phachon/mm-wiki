import { IRouter, AUTH_LOGIN_PATH } from '../type'
import Login from '@/pages/system/Login'
import Test from '@/pages/system/Error/Test'
import Install from '@/pages/system/Install'

const routers: IRouter[] = [
  {
    path: AUTH_LOGIN_PATH,
    key: 'auth_login',
    auth: false,
    title: '登录',
    component: <Login />
  },
  {
    path: '/install',
    key: 'install_wizard',
    auth: false,
    permission: false,
    title: '安装向导',
    component: <Install.InstallWizard />
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
