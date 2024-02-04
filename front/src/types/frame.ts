export type INavItem = {
  label: string
  key: string
  icon: string | null
  privilegeId: string | ''
}

export type IMenuItem = {
  name: string
  pageRouter: string
  key: string
  icon: string
  children: IMenuItem[] | null
  privilegeId: string | ''
  parentId: string | ''
  navId: string
}

export type IFrameBreadcrumbItem = {
  key: string // 唯一key
  title: string // 导航名称
  link?: string // 点击跳转陆游
}
