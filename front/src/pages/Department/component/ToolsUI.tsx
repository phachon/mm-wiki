import { DepartmentListItemType } from '@/types/departmentType'
import type { TreeDataNode } from 'antd'

export const getTreeData = (list: DepartmentListItemType[]): TreeDataNode[] => {
  let treeData: TreeDataNode[] = []
  list.forEach((item) => {
    let node: TreeDataNode = {
      title: item.name,
      key: item.department_id,
      isLeaf: true,
      children: []
    }
    if (item.children && item.children.length > 0) {
      node.children = getTreeData(item.children)
      node.isLeaf = false
    }
    treeData.push(node)
  })
  return treeData
}

export const getSelectTreeData = (departments: DepartmentListItemType[]): any[] => {
  return departments.map((department) => ({
    id: department.department_id,
    pid: department.parent_id,
    title: department.name,
    value: department.department_id.toString(),
    children: department.children ? getSelectTreeData(department.children) : []
  }))
}
