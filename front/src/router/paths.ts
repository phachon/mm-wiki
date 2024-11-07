import { IRouter } from '@/router/type'

const basicRoutes: IRouter[] = []
const modules = require.context('@/router/modules', true, /\.tsx$/)
const routerPathArr: string[] = []

function setPathsList(routes: IRouter[], parentPath: string = '') {
  routes.forEach((item: IRouter) => {
    const fullPath =
      item.path === '' ? parentPath : `${parentPath}/${item.path}`.replace(/\/+/g, '/')

    routerPathArr.push(fullPath)

    if (item.children) {
      setPathsList(item.children, item.path === '*' ? parentPath : fullPath)
    }
  })
}

let files = modules.keys()
files.forEach((key) => {
  const mod = modules(key).default
  const modList = Array.isArray(mod) ? [...mod] : [mod]
  setPathsList(modList)
  basicRoutes.push(...modList)
})

export const routerPaths = routerPathArr
export default basicRoutes
