import { Route, Routes } from 'react-router-dom'
import { IRouter } from './type'
import RouterInterceptor from './interceptor'
import basicRoutes from './paths'

const routeView = (routers: IRouter[]) => {
  return routers.map((router: IRouter) => {
    return (
      <Route
        path={router.path}
        key={router.path}
        element={<RouterInterceptor router={router}>{router.component}</RouterInterceptor>}
      >
        {router?.children && routeView(router.children)}
      </Route>
    )
  })
}

const RouterView = () => {
  const routers = routeView(basicRoutes)
  return <Routes>{routers}</Routes>
}

export default RouterView
