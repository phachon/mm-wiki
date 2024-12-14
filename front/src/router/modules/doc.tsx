import { IRouter } from '../type'
import Doc from '@/pages/space/Doc'
import { Navigate } from 'react-router-dom'

const routers: IRouter[] = [
  {
    path: '/doc',
    key: 'doc',
    component: <Doc.DocIndex />,
    permission: false,
    auth: true,
    children: [
      {
        path: '/doc/edit/:doc_id/',
        key: 'doc_edit',
        component: <Doc.DocEdit />,
        permission: false,
        auth: true
      },
      {
        path: '/doc/:doc_id',
        key: 'doc_view',
        component: <Doc.DocView />,
        permission: false,
        auth: true
      },
      {
        path: '*',
        key: 'spaces_default',
        component: <Navigate to="/spaces/all" replace />,
        auth: true,
        permission: false
      }
    ]
  }
]

export default routers
