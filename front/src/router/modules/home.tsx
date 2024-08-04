import Error403 from '@/pages/Error/Error403'
import { IRouter, NO_ACCESS_PATH, NO_EXIST_PATH } from '../type'
import SpaceHome from '@/pages/Space'
import Error404 from '@/pages/Error/Error404'
import Home from '@/pages/Home'

const routers: IRouter[] = [
  {
    path: '/',
    key: 'home',
    component: <Home.HomeIndex />,
    children: [
      {
        path: '/home/index',
        key: 'home_index',
        component: <Home.HomeIndex />,
        auth: false,
        permission: false
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
      },
      {
        path: '*',
        key: 'home_default',
        component: <SpaceHome />,
        auth: true,
        permission: false
      }
    ],
    permission: false,
    auth: true
  }
]

export default routers
