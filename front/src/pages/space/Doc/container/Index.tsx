import LayoutHeader from '@/components/Layout/Header'
import LayoutSider from '@/components/Layout/Sider'
import { LayoutHeaderSpaceKey } from '@/components/Layout/types'
import { useGlobalStore } from '@/stores'
import { Layout, message, Modal, TreeDataNode } from 'antd'
import React, { useEffect, useState } from 'react'
import { Outlet, useParams } from 'react-router-dom'
import SpaceSidebarUI, { ActionType } from '../../Space/component/SidebarUI'
import { ContentEntity, DocEntity, DocInfoResp, DocSaveResp, DocTreeEntity } from '@/types/docType'
import { SpaceDocsResp, SpaceInfoType } from '@/types/spaceType'
import { SpaceSpaceService } from '@/services/SpaceSpace'
import DocAddUI from '../../Doc/component/AddUI'
import { SpaceDocService } from '@/services/SpaceDoc'
import DocViewUI from '../../Doc/component/ViewUI'
import { useNavigate, useLocation } from 'react-router-dom'

const DocIndex: React.FC = () => {
  const store = useGlobalStore()
  const navigate = useNavigate()
  const { doc_id } = useParams<{ doc_id: string }>()

  // 文档详情拉取数据 /doc/:doc_id
  useEffect(() => {
    store.initDocsByDocId(Number(doc_id))
  }, [doc_id])

  const onClickDocSelect = (docId: string) => {
    navigate(`/doc/${docId}`)
  }

  return (
    <Layout>
      <LayoutHeader
        {...useGlobalStore()}
        accountInfo={store.getAccountInfo()}
        navSelectedKeys={[LayoutHeaderSpaceKey]}
      />
      <Layout>
        <LayoutSider
          content={
            <SpaceSidebarUI
              loading={store.siderLoading}
              spaceInfo={store.spaceInfo}
              dirTree={store.dirTree}
              homeDoc={store.homeDoc}
              onClickDocAction={store.onClickDocAction}
              onClickDocSelect={onClickDocSelect}
              selectDocId={store.selectDocId}
            />
          }
        />
        <Layout.Content className="space-content" style={{ padding: '20px 16px 0 24px' }}>
          <Outlet />
        </Layout.Content>
      </Layout>
      <Modal
        title="添加文档"
        open={store.addDocInfo?.modal}
        onCancel={() => store.setAddDocModal(false)}
        footer={null}
      >
        <DocAddUI onSaveSubmit={store.onAddDocSubmit} parentDoc={store.addDocInfo} />
      </Modal>
    </Layout>
  )
}

export default DocIndex
