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
        path: 'home',
        key: 'spaces_home',
        component: <Space.SpaceHome />,
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
    path: '/space/:key',
    key: 'space',
    component: <Space.SpaceHome />,
    permission: false,
    auth: true
  }
]

export default routers
