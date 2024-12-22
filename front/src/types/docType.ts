// DocType 文档类型
export enum DocType {
  DOC = 1,
  FOLDER = 2
}

// 文档类型定义
export const DocTypes = [
  {
    type: DocType.DOC,
    name: '文档',
    icon: 'FileTextOutlined'
  },
  {
    type: DocType.FOLDER,
    name: '文件夹',
    icon: 'FolderOutlined'
  }
]

// DocEntity 文档实体
export type DocEntity = {
  doc_id: number // 文档ID
  parent_id: number // 父级ID
  space_id: number // 空间ID
  space_key: string // 空间Key
  name: string // 文档名称
  type: DocType // 文档类型
  path: string // 路径
  sequence: number // 排序
  create_account_id: number // 创建账号ID
  create_account_name: string // 创建账号名
  edit_account_id: number // 最后修改账号ID
  edit_account_name: string // 最后修改账号名
  create_time: string // 创建时间
  update_time: string // 修改时间
}

// DocTreeEntity 文档树形结构
export type DocTreeEntity = DocEntity & {
  children: DocTreeEntity[] // 子文档
}

// ContentEntity 文档内容实体
export type ContentEntity = {
  doc_id: number // 文档ID
  content: string // 文档内容
  current_version_id: number // 当前版本ID
  create_time: string // 创建时间
  update_time: string // 修改时间
}

// DocAddSaveReq 文档添加保存请求
export type DocAddSaveReq = {
  space_key: string // 空间Key
  parent_id: number // 父级ID
  name: string // 文档名称
  doc_type: DocType // 文档类型
}

// DocAddSaveResp 文档添加保存响应
export type DocAddSaveResp = {
  doc_id: number // 文档ID
}

// DocInfoResp 文档信息响应
export type DocInfoResp = {
  doc_info: DocEntity // 文档信息
  content: ContentEntity // 文档内容
}

// ActionType 操作类型
export const ActionType = {
  ADD: 'add',
  EDIT: 'edit',
  COPY: 'copy',
  MOVE: 'move',
  DELETE: 'delete'
}

// DocActionReq 文档操作请求
export type DocContentSaveReq = {
  doc_id: number // 文档ID
  content: string // 文档内容
  name: string // 文档名称
}

// DocContentSaveResp 文档操作响应
export type DocContentSaveResp = {
  doc_id: number // 文档ID
  current_version_id: number // 当前版本ID
}
