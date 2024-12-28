import { useGlobalStore } from '@/stores'
import DocViewUI from '../component/ViewUI'
import DocHistoryUI from '../component/HistoryUI'
import { Modal } from 'antd'
import { useState } from 'react'
import { SpaceDocService } from '@/services/SpaceDoc'
import { DocContentHistortyResp, DocVersionEntity } from '@/types/contentType'
import { initPagination } from '@/types/adminType'

const DocView: React.FC = () => {
  const store = useGlobalStore()
  const [historyModalOpen, setHistoryModalOpen] = useState(false)
  const [historyList, setHistoryList] = useState<DocVersionEntity[]>([])
  const [historyPagination, setHistoryPagination] = useState(initPagination)

  // onHistoryClick 查看修改历史
  const onHistoryClick = (docId?: number) => {
    if (!docId) {
      return
    }
    console.log('查看修改历史', docId)
    setHistoryModalOpen(true)
    SpaceDocService.getDocHistory(docId)
      .then((resp: DocContentHistortyResp) => {
        setHistoryList(resp.version_list)
        setHistoryPagination({
          ...initPagination,
          current: resp.page_info.page_num,
          pageSize: resp.page_info.page_size,
          total: resp.page_info.total_num
        })
        console.log('历史记录', resp)
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
        content={store.content}
        parentPath={store.parentPath}
        onHistoryClick={onHistoryClick}
      />
      <Modal
        title="文档历史"
        width={950}
        open={historyModalOpen}
        onCancel={() => {
          setHistoryModalOpen(false)
        }}
        footer={null}
      >
        <DocHistoryUI historyList={historyList} pagination={historyPagination} />
      </Modal>
    </>
  )
}

export default DocView
