import LayoutHeader from '@/components/Layout/Header'
import LayoutSider from '@/components/Layout/Sider'
import { LayoutHeaderSpaceKey } from '@/components/Layout/types'
import { useGlobalStore } from '@/stores'
import { Layout, message, Modal } from 'antd'
import React, { useState } from 'react'
import { Outlet, useParams } from 'react-router-dom'
import SpaceSidebarUI from '../component/SidebarUI'
import { DocSaveResp, DocTreeEntity } from '@/types/docType'
import { SpaceDocsResp, SpaceInfoType } from '@/types/spaceType'
import { SpaceService } from '@/services/Space'
import { AddDocUI } from '../component/AddDocUI'
import { DocService } from '@/services/Doc'

const SpaceHome: React.FC = () => {
  const { getAccountInfo } = useGlobalStore()
  const { key } = useParams<{ key: string }>()
  const [dirTree, setDirTree] = useState<DocTreeEntity[]>([])
  const [homeDoc, setHomeDoc] = useState<DocTreeEntity>()
  const [spaceInfo, setSpaceInfo] = useState<SpaceInfoType>()
  const [addDocModal, setAddDocModal] = useState<boolean>(false)
  const [parentDoc, setParentDoc] = useState<DocTreeEntity>()

  React.useEffect(() => {
    if (key) {
      getSpaceDocs()
    }
  }, [key])

  if (!key) {
    return <div>空间 Key 不能为空</div>
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

  const onClickAddDoc = (docInfo?: DocTreeEntity) => {
    if (!docInfo) {
      message.error('数据异常，主页文档不存在！')
      return
    }
    setParentDoc(docInfo)
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
              onClickAddDoc={onClickAddDoc}
            />
          }
        />
        <Layout.Content className="home-content" style={{ padding: 16 }}>
          <Outlet />
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
