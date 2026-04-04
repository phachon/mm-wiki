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
import Email from '@/pages/system/Email'
import Link from '@/pages/system/Link'
import Contact from '@/pages/system/Contact'
import LoginAuth from '@/pages/system/LoginAuth'
import { IRouter, SYSTEM_ROOT_PATH } from '../type'
import Space from '@/pages/system/Space'

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
      },
      {
        path: '/system/email/add',
        key: 'email_add',
        component: <Email.EmailAdd />,
        auth: true,
        permission: true
      },
      {
        path: '/system/email/list',
        key: 'email_list',
        component: <Email.EmailList />,
        auth: true,
        permission: true
      },
      {
        path: '/system/link/add',
        key: 'link_add',
        component: <Link.LinkAdd />,
        auth: true,
        permission: true
      },
      {
        path: '/system/link/list',
        key: 'link_list',
        component: <Link.LinkList />,
        auth: true,
        permission: true
      },
      {
        path: '/system/contact/add',
        key: 'contact_add',
        component: <Contact.ContactAdd />,
        auth: true,
        permission: true
      },
      {
        path: '/system/contact/list',
        key: 'contact_list',
        component: <Contact.ContactList />,
        auth: true,
        permission: true
      },
      {
        path: '/system/login_auth/add',
        key: 'login_auth_add',
        component: <LoginAuth.LoginAuthAdd />,
        auth: true,
        permission: true
      },
      {
        path: '/system/login_auth/list',
        key: 'login_auth_list',
        component: <LoginAuth.LoginAuthList />,
        auth: true,
        permission: true
      },
      {
        path: '/system/space/add',
        key: 'space_add',
        component: <Space.SpaceAdd />,
        auth: true,
        permission: true
      },
      {
        path: '/system/space/list',
        key: 'space_list',
        component: <Space.SpaceList />,
        auth: true,
        permission: true
      }
    ],
    permission: false,
    auth: true
  }
]

export default routers
