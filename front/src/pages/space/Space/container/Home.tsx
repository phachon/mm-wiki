import LayoutHeader from '@/components/Layout/Header'
import LayoutSider from '@/components/Layout/Sider'
import { LayoutHeaderSpaceKey } from '@/components/Layout/types'
import { useGlobalStore } from '@/stores'
import { Layout, message, Modal, TreeDataNode } from 'antd'
import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import SpaceSidebarUI, { ActionType } from '../component/SidebarUI'
import { ContentEntity, DocEntity, DocInfoResp, DocSaveResp, DocTreeEntity } from '@/types/docType'
import { SpaceDocsResp, SpaceInfoType } from '@/types/spaceType'
import { SpaceSpaceService } from '@/services/SpaceSpace'
import { AddDocUI } from '../component/AddDocUI'
import { SpaceDocService } from '@/services/SpaceDoc'
import SpaceDocViewUI from '../component/DocViewUI'
import { useNavigate, useLocation } from 'react-router-dom'

const SpaceHome: React.FC = () => {
  const { getAccountInfo } = useGlobalStore()
  const navigate = useNavigate()
  const { space_key, doc_id } = useParams<{ space_key: string; doc_id: string }>()

  const [docsLoading, setDocsLoading] = useState<boolean>(false)
  const [viewDocLoading, setViewDocLoading] = useState<boolean>(false)
  const [dirTree, setDirTree] = useState<DocTreeEntity[]>([])
  const [homeDoc, setHomeDoc] = useState<DocTreeEntity>()
  const [spaceInfo, setSpaceInfo] = useState<SpaceInfoType>()
  const [addDocModal, setAddDocModal] = useState<boolean>(false)
  const [parentDoc, setParentDoc] = useState<{ parent_id: number; parent_name: string }>()
  const [viewDocInfo, setViewDocInfo] = useState<DocEntity>()
  const [content, setContent] = useState<ContentEntity>()
  const [selectedDocId, setSelectedDocId] = useState<string>()
  const [parentPath, setParentPath] = useState<string[]>([])

  // 空间主页拉取数据 /space/:space_key
  useEffect(() => {
    console.log('space_key:', space_key)
    if (space_key) {
      setDocsLoading(true)
      setViewDocLoading(true)
      fetchDocsBySpaceKey(space_key)
    }
  }, [space_key])

  // 文档详情拉取数据 /doc/:doc_id
  useEffect(() => {
    console.log('doc_id:', doc_id)
    if (doc_id) {
      setDocsLoading(true)
      setViewDocLoading(true)
      fetchDocsByDocId(Number(doc_id))
    }
  }, [doc_id])

  // 获取空间下文档
  const fetchDocsBySpaceKey = async (spaceKey: string) => {
    const spaceDocsRes = await SpaceSpaceService.getSpaceDocs(spaceKey)
    setSpaceInfo(spaceDocsRes.space_info)
    setDirTree(spaceDocsRes.dir_tree)
    setHomeDoc(spaceDocsRes.home_doc)
    setDocsLoading(false)
    if (spaceDocsRes.home_doc.doc_id) {
      const homeDocRes = await SpaceDocService.getDocInfo(spaceDocsRes.home_doc.doc_id)
      setViewDocInfo(homeDocRes.doc_info)
      setContent(homeDocRes.content)
    }
    setViewDocLoading(false)
    setParentPath([spaceDocsRes.space_info.name])
  }

  // 获取文档详情
  const fetchDocsByDocId = async (docId: number) => {
    const docInfoRes = await SpaceDocService.getDocInfo(docId)
    // 获取空间下的所有文档
    let dirTreeList = dirTree
    let spaceName = spaceInfo?.name || ''
    if (spaceInfo?.space_id != docInfoRes.doc_info.space_id || !dirTree.length) {
      const spaceDocsRes = await SpaceSpaceService.getSpaceDocs(docInfoRes.doc_info.space_key)
      setHomeDoc(spaceDocsRes.home_doc)
      setSpaceInfo(spaceDocsRes.space_info)
      setDirTree(spaceDocsRes.dir_tree)
      dirTreeList = spaceDocsRes.dir_tree
      spaceName = spaceDocsRes.space_info.name
    }
    setDocsLoading(false)
    setViewDocInfo(docInfoRes.doc_info)
    setContent(docInfoRes.content)
    setViewDocLoading(false)
    setSelectedDocId(docId.toString()) // 设置选中的文档ID
    updateParentPath(docId.toString(), spaceName, dirTreeList)
  }

  const onClickDocAction = (action: string, node: TreeDataNode) => {
    if (action === ActionType.ADD) {
      onClickAddDoc(node)
    }
  }

  const onClickAddDoc = (node?: TreeDataNode) => {
    if (!node) {
      message.error('文档数据异常，请刷新页面重试！')
      return
    }
    setParentDoc({
      parent_id: Number(node.key),
      parent_name: node.title as string
    })
    setAddDocModal(true)
  }

  const onAddDocSubmit = (values: any) => {
    values.space_key = space_key
    SpaceDocService.saveDoc(values)
      .then((res: DocSaveResp) => {
        message.success('文档保存成功！', 1).then(() => {
          setParentDoc(undefined)
          setAddDocModal(false)
          setDirTree([]) //  清空文档树
          navigate(`/doc/${res.doc_id}`)
        })
      })
      .catch((err) => {
        console.error(err)
      })
  }

  const onClickDocSelect = (docId: string) => {
    console.log('选中文档:', docId)
    navigate(`/doc/${docId}`)
  }

  // 更新父级路径
  const updateParentPath = (docId: string, spaceName: string, dirTreeData: DocTreeEntity[]) => {
    const parentPath = getParentPath(docId, dirTreeData)
    parentPath.unshift(spaceName) // 空间名作为第一个目录
    setParentPath(parentPath)
  }

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

  return (
    <Layout>
      <LayoutHeader
        {...useGlobalStore()}
        accountInfo={getAccountInfo()}
        navSelectedKeys={[LayoutHeaderSpaceKey]}
      />
      <Layout>
        <LayoutSider
          content={
            <SpaceSidebarUI
              loading={docsLoading}
              spaceInfo={spaceInfo}
              dirTree={dirTree}
              homeDoc={homeDoc}
              onClickDocAction={onClickDocAction}
              onClickDocSelect={onClickDocSelect}
              selectDocId={selectedDocId}
            />
          }
        />
        <Layout.Content className="space-content" style={{ padding: '20px 16px 0 24px' }}>
          <SpaceDocViewUI
            loading={viewDocLoading}
            docInfo={viewDocInfo}
            content={content}
            parentPath={parentPath}
          />
        </Layout.Content>
      </Layout>
      <Modal
        title="添加文档"
        open={addDocModal}
        onCancel={() => setAddDocModal(false)}
        footer={null}
      >
        <AddDocUI onSaveSubmit={onAddDocSubmit} parentDoc={parentDoc} />
      </Modal>
    </Layout>
  )
}

export default SpaceHome
