/**
 * DepartmentInfoType 部门返回结构
 */
export type DepartmentInfoType = {
  department_id: bigint // 部门ID
  name: string // 部门名
  parent_id: bigint // 上级部门ID
  parent_ids: string // 全部的上级部门ID逗号隔开
  sequence: number // 排序数字
  create_time: string // 创建时间
  update_time: string // 修改时间
}

/**
 * DepartmentListItemType 部门列表结构
 */
export type DepartmentListItemType = DepartmentInfoType & {
  children?: DepartmentListItemType[] // 子部门
  action?: {
    is_add: number
    is_edit: number
    is_delete: number
  }
}

/**
 * DepartmentListResp 部门列表返回
 */
export type DepartmentListResp = {
  list: DepartmentListItemType[]
}

/**
 * DepartmentEditResp 编辑部门返回结构
 */
export type DepartmentEditResp = {
  department: DepartmentInfoType // 部门信息
}

/**
 * DepartmentAddResp 添加部门返回结构
 */
export type DepartmentAddResp = {
  departments: DepartmentListItemType[] // 上级部门
}
