import Doc from '@/pages/space/Doc'
import { IRouter } from '../type'
import Space from '@/pages/space/Space'
import { Navigate } from 'react-router-dom'

const routers: IRouter[] = [
  {
    path: '/spaces',
    key: 'spaces',
    component: <Space.SpaceIndex />,
    children: [
      {
        path: '',
        key: 'space_redirect',
        component: <Navigate to="/spaces/all" replace />,
        auth: false,
        permission: false
      },
      {
        path: 'all',
        key: 'spaces_all',
        component: <Space.SpaceAll />,
        auth: false,
        permission: false
      },
      {
        path: 'hot',
        key: 'spaces_hot',
        component: <Space.SpaceAll />,
        auth: false,
        permission: false
      },
      {
        path: '*',
        key: 'spaces_default',
        component: <Navigate to="/spaces/all" replace />,
        auth: true,
        permission: false
      }
    ],
    permission: false,
    auth: true
  },
  {
    path: '/space',
    key: 'space',
    component: <Doc.DocIndex />,
    permission: false,
    auth: true,
    children: [
      {
        path: '/space/:space_key',
        key: 'space_home',
        component: <Doc.DocView />,
        permission: false,
        auth: true
      },
      {
        path: '/space/setting/:space_key',
        key: 'space_setting',
        component: <Space.SpaceSetting />,
        permission: false,
        auth: true
      }
    ]
  }
]

export default routers
