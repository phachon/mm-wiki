import { PageInfoType } from './baseType'

// ContentEntity 文档内容实体
export type ContentEntity = {
  doc_id: number // 文档ID
  content: string // 文档内容
  create_time: string // 创建时间
  update_time: string // 修改时间
}

// DocVersionEntity 文档版本实体
export type DocVersionEntity = {
  content_version_id: number // 版本ID
  content: string // 版本内容
  doc_id: number // 文档ID
  create_time: string // 创建时间
  update_time: string // 修改时间
  edit_account_id: number // 账号ID
  edit_account_name: string // 账号名
}

// DocContentSaveResp 文档操作响应
export type DocContentSaveResp = {}

// DocContentHistortyReq 文档内容历史请求
export type DocContentHistortyResp = {
  version_list: DocVersionEntity[] // 版本列表
  page_info: PageInfoType // 分页信息
}

// DocContentVersionResp 文档内容版本响应
export type DocContentVersionResp = {
  content_version: DocVersionEntity // 版本信息
}