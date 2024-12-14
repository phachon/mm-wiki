import LayoutHeader from '@/components/Layout/Header'
import LayoutSider from '@/components/Layout/Sider'
import { LayoutHeaderSpaceKey } from '@/components/Layout/types'
import { useGlobalStore } from '@/stores'
import { Empty, Layout, Modal } from 'antd'
import React, { useEffect } from 'react'
import { Outlet, useParams } from 'react-router-dom'
import DocAddUI from '../../Doc/component/AddUI'
import { useNavigate } from 'react-router-dom'
import DocTreeUI from '../component/TreeUI'

const DocIndex: React.FC = () => {
  const store = useGlobalStore()
  const navigate = useNavigate()
  const { doc_id, space_key } = useParams<{ doc_id: string; space_key: string }>()

  useEffect(() => {
    store.initDocsByDocId(Number(doc_id))
  }, [doc_id])

  useEffect(() => {
    if (space_key) {
      store.initDocsBySpaceKey(space_key)
    }
  }, [space_key])

  const onClickDocSelect = (docId: string) => {
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
            store.siderLoading ? (
              <Empty />
            ) : (
              <DocTreeUI
                loading={store.siderLoading}
                spaceInfo={store.spaceInfo}
                dirTree={store.dirTree}
                homeDoc={store.homeDoc}
                onClickDocAction={store.onClickDocAction}
                onClickDocSelect={onClickDocSelect}
                selectDocId={store.selectDocId}
              />
            )
          }
        />
        <Layout.Content className="space-content">
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
