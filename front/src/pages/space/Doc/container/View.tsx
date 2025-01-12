import { useGlobalStore } from '@/stores'
import DocViewUI from '../component/ViewUI'
import DocHistoryUI from '../component/HistoryUI'
import DiffViewUI from '../component/DiffViewUI' // 引入 DiffViewUI 组件
import { Modal, TablePaginationConfig, message } from 'antd' // 引入 Button 组件
import { useEffect, useState } from 'react'
import { SpaceDocService } from '@/services/SpaceDoc'
import {
  DocContentHistortyResp,
  DocContentVersionResp,
  DocVersionEntity
} from '@/types/contentType'
import { initPagination } from '@/types/adminType'
import { on } from 'events'
import { UserInteractionService } from '@/services/UserInteraction'

const DocView: React.FC = () => {
  const store = useGlobalStore()
  const [historyModalOpen, setHistoryModalOpen] = useState(false)
  const [showDiff, setShowDiff] = useState(false)
  const [historyList, setHistoryList] = useState<DocVersionEntity[]>([])
  const [historyPagination, setHistoryPagination] = useState(initPagination)
  const [historyDocVersion, setHistoryDocVersion] = useState<DocVersionEntity>()
  const [onlineContent, setOnlineContent] = useState('')
  const [docIsCollected, setDocIsCollected] = useState(false)

  useEffect(() => {
    initDocCollectionStatus(store.viewDocInfo?.doc_id || 0)
  }, [store.viewDocInfo])

  // initDocCollectionStatus 初始化文档收藏状态
  const initDocCollectionStatus = (docId: number) => {
    if (docId <= 0) {
      return
    }
    UserInteractionService.collectionDocStatus(docId)
      .then((res) => {
        setDocIsCollected(res.collection_status == 1)
      })
      .catch((e) => {
        console.error('获取文档收藏状态失败', e)
      })
  }

  // onDocCollectionChange 文档收藏操作
  const onDocCollectionChange = (docId: number, collected: boolean) => {
    if (collected) {
      UserInteractionService.collectionDoc(docId)
        .then(() => {
          message.success('收藏成功', 1, () => {
            initDocCollectionStatus(docId)
          })
        })
        .catch(() => {
          console.error('收藏失败')
        })
    } else {
      UserInteractionService.collectionDocCancel(docId)
        .then(() => {
          message.success('收藏已取消', 1, () => {
            initDocCollectionStatus(docId)
          })
        })
        .catch(() => {
          console.error('取消收藏失败')
        })
    }
  }

  // onHistoryClick 查看修改历史
  const onHistoryClick = (docId?: number) => {
    if (!docId) {
      return
    }
    setHistoryModalOpen(true)
    getDocHistory(docId, { ...initPagination })
  }

  // onViewClick 查看历史版本
  const onViewClick = (docVersion: DocVersionEntity) => {
    SpaceDocService.getDocContentVersion(docVersion.doc_id, docVersion.content_version_id)
      .then((resp: DocContentVersionResp) => {
        setHistoryDocVersion(resp.content_version)
        setOnlineContent(store.content?.content || '')
        setShowDiff(true) // 显示 diff 页面
      })
      .catch((e) => {
        console.error('获取历史版本失败', e)
      })
  }

  // onRecoverClick 恢复操作
  const onRecoverClick = (contentVersionId?: number, docId?: number) => {
    if (!contentVersionId || !docId) {
      return
    }
    SpaceDocService.recoverDoc(contentVersionId, docId)
      .then(() => {
        message.success('恢复成功', 1).then(() => {
          setHistoryModalOpen(false)
          setShowDiff(false)
          store.initDocsByDocId(docId)
        })
      })
      .catch((e) => {
        console.error('恢复失败', e)
      })
  }

  // onHistoryListChange 历史记录分页
  const onHistoryListChange = (pageConfig: TablePaginationConfig) => {
    if (!store.viewDocInfo) {
      return
    }
    getDocHistory(store.viewDocInfo.doc_id, pageConfig)
  }

  // onHistoryDelete 删除历史记录
  const onHistoryDelete = (contentVersionId: number, docId: number) => {
    SpaceDocService.delDocContentVersion(contentVersionId, docId)
      .then(() => {
        message.success('删除成功', 1).then(() => {
          getDocHistory(docId, historyPagination)
        })
      })
      .catch((e) => {
        console.error('删除失败', e)
      })
  }

  // getDocHistory 获取历史记录
  const getDocHistory = (docId: number, pageConfig: TablePaginationConfig) => {
    SpaceDocService.getDocHistory(docId, pageConfig.pageSize, pageConfig.current)
      .then((resp: DocContentHistortyResp) => {
        setHistoryList(resp.version_list)
        setHistoryPagination({
          ...initPagination,
          current: resp.page_info.page_num,
          pageSize: resp.page_info.page_size,
          total: resp.page_info.total_num
        })
      })
      .catch((e) => {
        console.error('获取历史记录失败', e)
      })
  }

  return (
    <>
      <DocViewUI
        loading={store.viewDocLoading}
        docInfo={store.viewDocInfo}
        isCollected={docIsCollected}
        content={store.content}
        parentPath={store.parentPath}
        onHistoryClick={onHistoryClick}
        onCollectionChange={onDocCollectionChange}
      />
      <Modal
        title={null}
        width={1150}
        open={historyModalOpen}
        onCancel={() => {
          setHistoryModalOpen(false)
          setShowDiff(false) // 关闭弹框时重置状态
        }}
        footer={null}
        closable={showDiff ? false : true}
      >
        {showDiff ? (
          <DiffViewUI
            historyDocVersion={historyDocVersion}
            onlineContent={onlineContent}
            onRecoverClick={onRecoverClick}
            onReturnClick={() => {
              setShowDiff(false)
            }}
          />
        ) : (
          <DocHistoryUI
            historyList={historyList}
            pagination={historyPagination}
            onViewClick={onViewClick}
            onRecoverClick={onRecoverClick}
            onListChange={onHistoryListChange}
            onDeleteConfirm={onHistoryDelete}
          />
        )}
      </Modal>
    </>
  )
}

export default DocView
