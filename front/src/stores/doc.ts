import { SpaceDocService } from '@/services/SpaceDoc'
import { SpaceSpaceService } from '@/services/SpaceSpace'
import { ActionType, ContentEntity, DocEntity, DocSaveResp, DocTreeEntity } from '@/types/docType'
import { SpaceInfoType } from '@/types/spaceType'
import { message, TreeDataNode } from 'antd'
import { StateCreator } from 'zustand'
import { useNavigate } from 'react-router-dom'

// IDoc: interface for doc store
export interface IDoc {
  initDocsByDocId: (docId: number) => void
  initDocsBySpaceKey: (spaceKey: string) => void

  /** 文档左侧栏 **/
  siderLoading?: boolean
  spaceInfo?: SpaceInfoType
  dirTree?: DocTreeEntity[]
  homeDoc?: DocTreeEntity
  selectDocId?: string
  onClickDocAction?: (action: string, node: TreeDataNode) => void
  onAddDocSubmit?: (values: any) => void
  setAddDocModal: (visible: boolean) => void
  addDocInfo?: {
    modal?: boolean
    parent_id?: number
    parent_name?: string
  }

  /** 文档正文 **/
  viewDocLoading?: boolean
  viewDocInfo?: DocEntity
  content?: ContentEntity
  parentPath?: string[]
}

export const createDoc: StateCreator<IDoc> = (set, get) => ({
  // 左侧栏
  siderLoading: true,
  spaceInfo: undefined,
  dirTree: [],
  homeDoc: undefined,
  selectDocId: '',

  // 文档正文
  viewDocLoading: true,
  viewDocInfo: undefined,
  content: undefined,
  parentPath: [],

  /**
   * 初始化文档
   */
  initDocsByDocId: async (docId: number) => {
    if (!docId) {
      return
    }
    console.log('初始化文档:', docId)
    const docInfoRes = await SpaceDocService.getDocInfo(docId)
    // 获取空间下所有文档
    if (!get().dirTree?.length) {
      const spaceDocsRes = await SpaceSpaceService.getSpaceDocs(docInfoRes.doc_info.space_key)
      set({
        siderLoading: false,
        homeDoc: spaceDocsRes.home_doc,
        spaceInfo: spaceDocsRes.space_info,
        dirTree: spaceDocsRes.dir_tree
      })
    }
    const parentPath = getParentPath(docId.toString(), get().dirTree)
    if (get().spaceInfo?.name) {
      parentPath.unshift(get().spaceInfo?.name || '')
    }
    set({
      viewDocLoading: false,
      viewDocInfo: docInfoRes.doc_info,
      content: docInfoRes.content,
      selectDocId: docId.toString(),
      parentPath: parentPath
    })
  },

  /**
   * 初始化空间
   */
  initDocsBySpaceKey: async (spaceKey: string) => {
    console.log('初始化空间:', spaceKey)
    const spaceDocsRes = await SpaceSpaceService.getSpaceDocs(spaceKey)
    if (!spaceDocsRes) {
      return
    }
    set({
      siderLoading: false,
      spaceInfo: spaceDocsRes.space_info,
      dirTree: spaceDocsRes.dir_tree,
      homeDoc: spaceDocsRes.home_doc,
      parentPath: [spaceDocsRes.space_info.name]
    })
    // 获取空间下所有文档
    if (spaceDocsRes.home_doc.doc_id) {
      const homeDocRes = await SpaceDocService.getDocInfo(spaceDocsRes.home_doc.doc_id)
      if (!homeDocRes) {
        return
      }
      set({
        viewDocInfo: homeDocRes.doc_info,
        content: homeDocRes.content
      })
    }
    set({
      viewDocLoading: false
    })
  },

  /**
   * 文档操作
   */
  onClickDocAction: (action: string, node: TreeDataNode) => {
    console.log('点击文档操作:', action, node)
    if (action === ActionType.ADD) {
      set({
        addDocInfo: {
          modal: true,
          parent_id: Number(node.key),
          parent_name: node.title as string
        }
      })
    }
  },

  /**
   * 添加文档弹框操作
   */
  setAddDocModal: (visible: boolean) => {
    set({
      addDocInfo: {
        ...get().addDocInfo,
        modal: visible
      }
    })
  },

  /**
   * 添加文档保存
   */
  onAddDocSubmit: async (values: any) => {
    const resp = await SpaceDocService.saveDoc(values)
    if (resp.doc_id) {
      message.success('文档保存成功！', 1).then(() => {
        set({
          addDocInfo: {
            modal: false,
            parent_id: 0,
            parent_name: ''
          },
          dirTree: [] //  清空文档树
        })
        get().initDocsByDocId(resp.doc_id)
      })
    }
  }
})

// 获取父级路径
const getParentPath = (docId: string, treeData?: DocTreeEntity[]): string[] => {
  if (!treeData) return []
  const findParentKeys = (
    nodes: DocTreeEntity[],
    targetKey: string,
    path: string[] = []
  ): string[] | null => {
    for (const node of nodes) {
      const currentPath = [...path, node.name.toString()]
      if (node.doc_id.toString() === targetKey) {
        return path // 返回父级路径，不包括当前节点
      }
      if (node.children) {
        const result = findParentKeys(node.children, targetKey, currentPath)
        if (result) {
          return result
        }
      }
    }
    return null
  }
  const parentKeys = findParentKeys(treeData, docId)
  return parentKeys || []
}
