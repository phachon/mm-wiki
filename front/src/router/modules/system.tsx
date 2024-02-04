import Profile from '@/pages/Profile'
import Account from '@/pages/Account'
import Privilege from '@/pages/Privilege'
import Main from '@/pages/Main'
import Log from '@/pages/Log'
import Role from '@/pages/Role'
import FrameHome from '@/pages/Frame'
import Error404 from '@/pages/Error/Error404'
import Error403 from '@/pages/Error/Error403'
import Notice from '@/pages/Notice'
import { HOME_ROOT_PATH, IRouter, NO_ACCESS_PATH, NO_EXIST_PATH } from '../type'

const routers: IRouter[] = [
  {
    path: '/',
    key: 'home',
    component: <FrameHome />,
    children: [
      {
        path: HOME_ROOT_PATH,
        key: 'main_index',
        component: <Main.MainIndex />,
        auth: true,
        permission: false
      },
      {
        path: '/profile/info',
        key: 'profile_info',
        component: <Profile.ProfileInfo />,
        auth: true,
        permission: true
      },
      {
        path: '/profile/repass',
        key: 'profile_repass',
        component: <Profile.ProfileRepass />,
        auth: true,
        permission: true
      },
      {
        path: '/account/add',
        key: 'account_add',
        component: <Account.AccountAdd />,
        auth: true,
        permission: true
      },
      {
        path: '/account/list',
        key: 'account_list',
        component: <Account.AccountList />,
        auth: true,
        permission: true
      },
      {
        path: '/role/add',
        key: 'role_add',
        component: <Role.RoleAdd />,
        auth: true,
        permission: true
      },
      {
        path: '/role/list',
        key: 'role_list',
        component: <Role.RoleList />,
        auth: true,
        permission: true
      },
      {
        path: '/privilege/add',
        key: 'privilege_add',
        component: <Privilege.PrivilegeAdd />,
        auth: true,
        permission: true
      },
      {
        path: '/privilege/list',
        key: 'privilege_list',
        component: <Privilege.PrivilegeList />,
        auth: true,
        permission: true
      },
      {
        path: '/log/list',
        key: 'log_list',
        component: <Log.LogList />,
        auth: true,
        permission: true
      },
      {
        path: '/notice/add',
        key: 'notice_add',
        component: <Notice.NoticeAdd />,
        auth: true,
        permission: true
      },
      {
        path: '/notice/list',
        key: 'notice_list',
        component: <Notice.NoticeList />,
        auth: true,
        permission: true
      },
      {
        path: NO_ACCESS_PATH,
        key: 'error_403',
        component: <Error403 />,
        auth: false,
        permission: false
      },
      {
        path: NO_EXIST_PATH,
        key: 'error_404',
        component: <Error404 />,
        auth: false,
        permission: false
      }
    ],
    permission: false,
    auth: true
  }
]

export default routers
