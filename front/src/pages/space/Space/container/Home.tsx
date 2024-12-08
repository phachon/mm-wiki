import LayoutHeader from '@/components/Layout/Header'
import LayoutSider from '@/components/Layout/Sider'
import { LayoutHeaderSpaceKey } from '@/components/Layout/types'
import { useGlobalStore } from '@/stores'
import { Layout, message, Modal, TreeDataNode } from 'antd'
import React, { useEffect, useState } from 'react'
import { Outlet, useParams } from 'react-router-dom'
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
  const location = useLocation()
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

  // 空间主页拉取数据 /space/:space_key
  useEffect(() => {
    setDocsLoading(true)
    setViewDocLoading(true)
    if (space_key) {
      fetchDocsBySpaceKey(space_key)
    }
  }, [space_key])

  // 文档详情拉取数据 /doc/:doc_id
  useEffect(() => {
    setDocsLoading(true)
    setViewDocLoading(true)
    if (doc_id) {
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
      setViewDocLoading(false)
    }
  }

  // 获取文档详情
  const fetchDocsByDocId = async (docId: number) => {
    const docInfoRes = await SpaceDocService.getDocInfo(docId)
    // 获取空间下的所有文档
    if (spaceInfo?.space_id != docInfoRes.doc_info.space_id || !dirTree.length) {
      const spaceDocsRes = await SpaceSpaceService.getSpaceDocs(docInfoRes.doc_info.space_key)
      setHomeDoc(spaceDocsRes.home_doc)
      setSpaceInfo(spaceDocsRes.space_info)
      setDirTree(spaceDocsRes.dir_tree)
    }
    setDocsLoading(false)
    setViewDocInfo(docInfoRes.doc_info)
    setContent(docInfoRes.content)
    setViewDocLoading(false)
    setSelectedDocId(docId.toString()) // 设置选中的文档ID
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
              selectDocId={selectedDocId} // 传递选中的文档ID
            />
          }
        />
        <Layout.Content className="space-content" style={{ padding: '20px 16px 0 24px' }}>
          <SpaceDocViewUI loading={viewDocLoading} docInfo={viewDocInfo} content={content} />
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
