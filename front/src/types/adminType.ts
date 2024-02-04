import { TablePaginationConfig } from 'antd'

export const AdminType = {}

// 默认的分页配置
export const initPagination: TablePaginationConfig = {
  current: 1,
  pageSize: 10,
  total: 0,
  showQuickJumper: true,
  showSizeChanger: true,
  showTotal: (total: number) => {
    return `总共 ${total} 条`
  }
}

/* 面包屑导航数据结构 */
export type FrameBreadcrumbItem = {
  key: string // 唯一key
  name: string // 导航名称
  link: string // 点击跳转陆游
}
