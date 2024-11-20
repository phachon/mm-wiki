import { PageInfoType } from './baseType'
import { AccountInfoType } from './accountType'
import { DocTreeEntity } from './docType'

export const SpaceVisitLevelDefaultPublic = 0 // 访问级别：默认公开
export const SpaceVisitLevelPrivate = 1 // 访问级别：私有

export const SpaceTypeTeam = 0 // 空间类型：团队空间
export const SpaceTypePersonal = 1 // 空间类型：个人空间

export const SpaceIsShareNo = 0 // 不可分享
export const SpaceIsShareYes = 1 // 可分享

export const SpaceIsExportNo = 0 // 不可导出
export const SpaceIsExportYes = 1 // 可导出

// 空间访问级别定义
export const SpaceVisitLevelTypes = [
  {
    type: SpaceVisitLevelDefaultPublic,
    name: '公开',
    color: 'green'
  },
  {
    type: SpaceVisitLevelPrivate,
    name: '私有',
    color: 'red'
  }
]

// 空间是否分享定义
export const SpaceIsShareTypes = [
  {
    type: SpaceIsShareNo,
    name: '否'
  },
  {
    type: SpaceIsShareYes,
    name: '是'
  }
]

// 空间是否导出定义
export const SpaceIsExportTypes = [
  {
    type: SpaceIsExportNo,
    name: '否'
  },
  {
    type: SpaceIsExportYes,
    name: '是'
  }
]

// 空间类型定义
export const SpaceTypeTypes = [
  {
    type: SpaceTypeTeam,
    name: '团队空间',
    color: 'blue'
  },
  {
    type: SpaceTypePersonal,
    name: '个人空间',
    color: 'green'
  }
]

// SpaceInfoType 空间信息结构
export type SpaceInfoType = {
  space_id: number // 空间ID
  space_key: string // 空间Key
  name: string // 空间名
  description: string // 空间描述
  space_type: number // 空间类型
  visit_level: number // 访问级别
  is_share: number // 是否分享
  is_export: number // 是否导出
  creator_name: string // 创建者账号名
  creator_account_id: bigint // 创建者账号ID
  status: number // 状态 0:正常 -1:删除
  create_time: string // 创建时间
  update_time: string // 修改时间
}

/**
 * SpaceInfoType 空间列表结构
 */
export type SpaceListItemType = SpaceInfoType & {
  action?: {
    is_edit: number // 是否可以修改
    is_delete: number // 是否可以删除
  } // 操作权限
}

/**
 * SpaceListResp 空间列表结构
 */
export type SpaceListResp = {
  list: SpaceListItemType[]
  page_info: PageInfoType
}

/**
 * SpaceEditResp 空间修改信息返回结构
 */
export type SpaceEditResp = {
  space_info: SpaceInfoType
}

/**
 * SpaceAddResp 添加空间信息返回结构
 */
export type SpaceAddResp = {
  account_list: AccountInfoType[]
}

// SpaceAdminListResp 空间管理员列表返回结构
export type SpaceAdminListResp = {
  admin_list: AccountInfoType[]
  selected_list: AccountInfoType[]
}

// SpaceDocsResp 空间文档返回结构
export type SpaceDocsResp = {
  home_doc: DocTreeEntity // 主页文档
  dir_tree: DocTreeEntity[] // 目录树
  space_info: SpaceInfoType // 空间信息
}
