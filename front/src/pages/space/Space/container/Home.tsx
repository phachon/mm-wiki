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
import DocAddUI from '../../Doc/component/AddUI'
import { SpaceDocService } from '@/services/SpaceDoc'
import DocViewUI from '../../Doc/component/ViewUI'
import { useNavigate, useLocation } from 'react-router-dom'

const SpaceHome: React.FC = () => {
  const store = useGlobalStore()
  const navigate = useNavigate()
  const { space_key } = useParams<{ space_key: string }>()

  // 空间主页拉取数据 /space/:space_key
  useEffect(() => {
    if (space_key) {
      store.initDocsBySpaceKey(space_key)
    }
  }, [space_key])

  const onClickDocSelect = (docId: string) => {
    console.log('选中文档:', docId)
    navigate(`/doc/${docId}`)
  }

  return (
    <Layout>
      <LayoutHeader
        {...store}
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
          <DocViewUI
            loading={store.viewDocLoading}
            docInfo={store.viewDocInfo}
            content={store.content}
            parentPath={store.parentPath}
          />
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

export default SpaceHome
