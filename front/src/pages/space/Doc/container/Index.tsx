import LayoutHeader from '@/components/Layout/Header'
import LayoutSider from '@/components/Layout/Sider'
import { LayoutHeaderSpaceKey } from '@/components/Layout/types'
import { useGlobalStore } from '@/stores'
import { Empty, Layout, message, Modal } from 'antd'
import React, { useEffect } from 'react'
import { Outlet, useParams } from 'react-router-dom'
import DocAddUI from '../../Doc/component/AddUI'
import { useNavigate } from 'react-router-dom'
import DocTreeUI from '../component/TreeUI'
import { UserInteractionService } from '@/services/UserInteraction'

// 文档主页组件
const DocIndex: React.FC = () => {
  const store = useGlobalStore()
  const navigate = useNavigate()
  const { doc_id, space_key } = useParams<{ doc_id: string; space_key: string }>()
  const [spaceIsCollected, setSpaceIsCollected] = React.useState(false)

  useEffect(() => {
    store.initDocsByDocId(Number(doc_id), navigate)
  }, [doc_id])

  useEffect(() => {
    if (space_key) {
      store.initDocsBySpaceKey(space_key, navigate)
    }
  }, [space_key])

  useEffect(() => {
    if (store.spaceInfo) {
      initSpaceCollectedStatus(store.spaceInfo.space_id)
    }
  }, [store.spaceInfo])

  // 初始化空间收藏状态
  const initSpaceCollectedStatus = (spaceId: number) => {
    UserInteractionService.collectionSpaceStatus(spaceId)
      .then((res) => {
        setSpaceIsCollected(res.collection_status == 1)
      })
      .catch((e) => {
        console.log('获取空间收藏状态失败', e)
      })
  }

  // 空间收藏
  const onSpaceCollection = (spaceId: number) => {
    UserInteractionService.collectionSpace(spaceId)
      .then(() => {
        message.success('收藏成功', 1, () => {
          initSpaceCollectedStatus(spaceId)
        })
      })
      .catch(() => {
        console.error('收藏失败')
      })
  }

  // 空间取消收藏
  const onSpaceCollectionCancel = (spaceId: number) => {
    UserInteractionService.collectionSpaceCancel(spaceId)
      .then(() => {
        message.success('收藏取消成功', 1, () => {
          initSpaceCollectedStatus(spaceId)
        })
      })
      .catch(() => {
        console.error('取消收藏失败')
      })
  }

  // 空间收藏状态变化
  const onCollectionStatusChange = (spaceId: number, isCollected: boolean) => {
    if (isCollected) {
      onSpaceCollection(spaceId)
    } else {
      onSpaceCollectionCancel(spaceId)
    }
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
                isCollected={spaceIsCollected}
                docTree={store.docTree}
                homeDoc={store.homeDoc}
                onClickDocAction={store.onClickDocAction}
                onClickDocSelect={store.onClickDocSelect}
                selectDocId={store.selectDocId}
                onCollectionStatusChange={onCollectionStatusChange}
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
