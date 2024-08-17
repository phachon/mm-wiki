import Profile from '@/pages/system/Profile'
import Account from '@/pages/system/Account'
import Privilege from '@/pages/system/Privilege'
import Main from '@/pages/system/Main'
import Log from '@/pages/system/Log'
import Role from '@/pages/system/Role'
import SystemHome from '@/pages/system/System'
import Notice from '@/pages/system/Notice'
import Config from '@/pages/system/Config'
import Department from '@/pages/system/Department'
import { IRouter, SYSTEM_ROOT_PATH } from '../type'

const routers: IRouter[] = [
  {
    path: '/system',
    key: 'system',
    component: <SystemHome />,
    children: [
      {
        path: SYSTEM_ROOT_PATH,
        key: 'main_index',
        component: <Main.MainIndex />,
        auth: false,
        permission: false
      },
      {
        path: '/system/profile/info',
        key: 'profile_info',
        component: <Profile.ProfileInfo />,
        auth: true,
        permission: false
      },
      {
        path: '/system/profile/setting',
        key: 'profile_setting',
        component: <Profile.ProfileSetting />,
        auth: true,
        permission: true
      },
      {
        path: '/system/account/add',
        key: 'account_add',
        component: <Account.AccountAdd />,
        auth: true,
        permission: true
      },
      {
        path: '/system/account/list',
        key: 'account_list',
        component: <Account.AccountList />,
        auth: true,
        permission: true
      },
      {
        path: '/system/role/add',
        key: 'role_add',
        component: <Role.RoleAdd />,
        auth: true,
        permission: true
      },
      {
        path: '/system/role/list',
        key: 'role_list',
        component: <Role.RoleList />,
        auth: true,
        permission: true
      },
      {
        path: '/system/privilege/add',
        key: 'privilege_add',
        component: <Privilege.PrivilegeAdd />,
        auth: true,
        permission: true
      },
      {
        path: '/system/privilege/list',
        key: 'privilege_list',
        component: <Privilege.PrivilegeList />,
        auth: true,
        permission: true
      },
      {
        path: '/system/log/list',
        key: 'log_list',
        component: <Log.LogList />,
        auth: true,
        permission: true
      },
      {
        path: '/system/notice/add',
        key: 'notice_add',
        component: <Notice.NoticeAdd />,
        auth: true,
        permission: true
      },
      {
        path: '/system/notice/list',
        key: 'notice_list',
        component: <Notice.NoticeList />,
        auth: true,
        permission: true
      },
      {
        path: '/system/config/add',
        key: 'config_add',
        component: <Config.ConfigAdd />,
        auth: true,
        permission: true
      },
      {
        path: '/system/department/list',
        key: 'department_list',
        component: <Department.DepartmentList />,
        auth: true,
        permission: true
      }
    ],
    permission: false,
    auth: true
  }
]

export default routers
