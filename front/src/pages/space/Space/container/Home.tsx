import LayoutHeader from '@/components/Layout/Header'
import LayoutSider from '@/components/Layout/Sider'
import { LayoutHeaderSpaceKey } from '@/components/Layout/types'
import { useGlobalStore } from '@/stores'
import { Layout, message, Modal, TreeDataNode } from 'antd'
import React, { useState } from 'react'
import { Outlet, useParams } from 'react-router-dom'
import SpaceSidebarUI, { ActionType } from '../component/SidebarUI'
import { DocSaveResp, DocTreeEntity } from '@/types/docType'
import { SpaceDocsResp, SpaceInfoType } from '@/types/spaceType'
import { SpaceService } from '@/services/Space'
import { AddDocUI } from '../component/AddDocUI'
import { DocService } from '@/services/Doc'
import SpaceDocViewUI from '../component/DocViewUI'

const SpaceHome: React.FC = () => {
  const { getAccountInfo } = useGlobalStore()
  const { key } = useParams<{ key: string }>()
  const [dirTree, setDirTree] = useState<DocTreeEntity[]>([])
  const [homeDoc, setHomeDoc] = useState<DocTreeEntity>()
  const [spaceInfo, setSpaceInfo] = useState<SpaceInfoType>()
  const [addDocModal, setAddDocModal] = useState<boolean>(false)
  const [parentDoc, setParentDoc] = useState<{
    parent_id: number
    parent_name: string
  }>()
  const [demoMk, setDemoMk] = useState('')

  React.useEffect(() => {
    fetchData()
    if (key) {
      getSpaceDocs()
    }
  }, [key])

  if (!key) {
    return <div>空间 Key 不能为空</div>
  }

  const fetchData = async () => {
    fetch('/demo.md')
      .then((response) => response.text())
      .then((value) => {
        console.log('value:', value)
        setDemoMk(value)
      })
  }

  const getSpaceDocs = () => {
    SpaceService.getSpaceDocs(key)
      .then((res: SpaceDocsResp) => {
        setSpaceInfo(res.space_info)
        setHomeDoc(res.home_doc)
        setDirTree(res.dir_tree)
      })
      .catch((err) => {
        console.error(err)
      })
  }

  const onClickActionDoc = (action: string, node: TreeDataNode) => {
    console.log('action:', action, 'node:', node)
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
    DocService.saveDoc(values)
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
              onClickActionDoc={onClickActionDoc}
            />
          }
        />
        <Layout.Content className="space-content" style={{ padding: '20px 16px 0 24px' }}>
          <SpaceDocViewUI content={demoMk} />
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
