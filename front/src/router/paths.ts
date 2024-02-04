import { IRouter } from '@/router/type'

const basicRoutes: IRouter[] = []
const modules = require.context('@/router/modules', true, /\.tsx$/)
const routerPathArr: string[] = []

function setPathsList(routes: IRouter[]) {
  routes.map((item: IRouter) => {
    routerPathArr.push(item.path)
    item.children && setPathsList(item.children)
  })
}

// 读取模块下的文件，获取每个文件里的 routers
let files = modules.keys()
files.forEach((key) => {
  const mod = modules(key).default
  const modList = Array.isArray(mod) ? [...mod] : [mod]
  setPathsList(modList)
  basicRoutes.push(...modList)
})

export const routerPaths = routerPathArr
export default basicRoutes
