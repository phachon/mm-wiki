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
  const { key } = useParams<{ key: string }>()

  const [dirTree, setDirTree] = useState<DocTreeEntity[]>([])
  const [homeDoc, setHomeDoc] = useState<DocTreeEntity>()
  const [spaceInfo, setSpaceInfo] = useState<SpaceInfoType>()
  const [addDocModal, setAddDocModal] = useState<boolean>(false)
  const [parentDoc, setParentDoc] = useState<{ parent_id: number; parent_name: string }>()
  const [viewDocInfo, setViewDocInfo] = useState<DocEntity>()
  const [content, setContent] = useState<ContentEntity>()
  const [selectedDocId, setSelectedDocId] = useState<string>()

  useEffect(() => {
    if (key) {
      getSpaceDocs()
    }
  }, [key])

  useEffect(() => {
    const params = new URLSearchParams(location.search)
    const docId = params.get('doc_id')
    if (docId) {
      getDocInfo(docId)
      setSelectedDocId(docId)
    }
  }, [location])

  if (!key) {
    return <div>空间 Key 不能为空</div>
  }

  const getSpaceDocs = () => {
    SpaceSpaceService.getSpaceDocs(key)
      .then((res: SpaceDocsResp) => {
        setSpaceInfo(res.space_info)
        setHomeDoc(res.home_doc)
        setDirTree(res.dir_tree)
      })
      .catch((err) => {
        console.error(err)
      })
  }

  const getDocInfo = (docId: string) => {
    SpaceDocService.getDocInfo(Number(docId))
      .then((res) => {
        setViewDocInfo(res.doc_info)
        setContent(res.content)
      })
      .catch((err) => {
        console.error(err)
      })
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
    values.space_key = key
    SpaceDocService.saveDoc(values)
      .then((res: DocSaveResp) => {
        message.success('文档保存成功！', 1).then(() => {
          setParentDoc(undefined)
          setAddDocModal(false)
          getSpaceDocs()
        })
      })
      .catch((err) => {
        console.error(err)
      })
  }

  const onClickDocSelect = (docId: string) => {
    console.log('选中文档:', docId)
    navigate(`?doc_id=${docId}`)
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
          <SpaceDocViewUI docInfo={viewDocInfo} content={content} />
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
