
// DocTreeEntity 文档树形结构
export type DocTreeEntity = {
  id: number // 文档ID
  space_id: number // 空间ID
  space_key: string // 空间Key
  parent_id: number // 父级ID
  title: string // 标题
  sort: number // 排序
  type: number // 类型 1:文件夹 2:文档
  create_time: string // 创建时间
  update_time: string // 修改时间
  children: DocTreeEntity[] // 子文档
}
