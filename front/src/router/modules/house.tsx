import FrameHome from '@/pages/Frame'
import { IRouter, NO_ACCESS_PATH, NO_EXIST_PATH } from '../type'
import Main from '@/pages/Main'
import House from '@/pages/House'
import Landlord from '@/pages/Landlord'
import Room from '@/pages/Room'
import Order from '@/pages/Order'
import Hirer from '@/pages/Hirer'

const routers: IRouter[] = [
  {
    path: '/',
    key: 'home',
    component: <FrameHome />,
    children: [
      {
        path: '/house/index',
        key: 'house_index',
        component: <Main.MainIndex />,
        auth: true,
        permission: true
      },
      {
        path: '/house/add',
        key: 'house_add',
        component: <House.HouseAdd />,
        auth: true,
        permission: true
      },
      {
        path: '/house/list',
        key: 'house_list',
        component: <House.HouseList />,
        auth: true,
        permission: true
      },
      {
        path: '/landlord/add',
        key: 'landlord_add',
        component: <Landlord.LandlordAdd />,
        auth: true,
        permission: true
      },
      {
        path: '/landlord/list',
        key: 'landlord_list',
        component: <Landlord.LandlordList />,
        auth: true,
        permission: true
      },
      {
        path: '/room/add',
        key: 'room_add',
        component: <Room.RoomAdd />,
        auth: true,
        permission: true
      },
      {
        path: '/room/list',
        key: 'room_list',
        component: <Room.RoomList />,
        auth: true,
        permission: true
      },
      {
        path: '/hirer/add',
        key: 'hirer_add',
        component: <Hirer.HirerAdd />,
        auth: true,
        permission: true
      },
      {
        path: '/hirer/list',
        key: 'hirer_list',
        component: <Hirer.HirerList />,
        auth: true,
        permission: true
      },
      {
        path: '/order/add',
        key: 'order_add',
        component: <Order.OrderAdd />,
        auth: true,
        permission: true
      },
      {
        path: '/order/list',
        key: 'order_list',
        component: <Order.OrderList />,
        auth: true,
        permission: true
      }
    ],
    permission: false,
    auth: true
  }
]

export default routers
