import React, { useEffect, useState } from 'react'
import {
  NoticeEditResp,
  NoticeInfoType,
  NoticeListItemType,
  NoticeListResp
} from '@/types/noticeType'
import { NoticeService } from '@/services/Notice'
import { message, Modal, TablePaginationConfig } from 'antd'
import NoticeListUI from '../component/ListUI'
import NoticeSearchUI from '../component/SearchUI'
import NoticeFormUI from '../component/FormUI'
import { initPagination } from '@/types/adminType'

let searchKeyWords = {}

const NoticeList: React.FC = () => {
  // 公告列表相关 state
  const [noticeList, setNoticeList] = useState<NoticeListItemType[]>([])
  const [pagination, setPagination] = useState(initPagination)
  const [editModalOpen, setEditModalOpen] = useState<boolean>(false)
  // 公告修改相关 state
  const [editNoticeInfo, setEditNoticeInfo] = useState<NoticeInfoType>()

  useEffect(() => {
    getNoticeList(initPagination, {})
  }, [])

  /**
   * 列表分页处理
   * @param pageConfig
   * @param filters
   * @param sorter
   */
  const onListChange = (pageConfig: TablePaginationConfig, filters: any, sorter: any) => {
    getNoticeList(pageConfig, searchKeyWords)
  }

  /**
   * 修改点击操作
   * @param noticeInfo
   */
  const onEditClick = (noticeInfo: NoticeInfoType) => {
    NoticeService.getEditNoticeInfo(noticeInfo.notice_id)
      .then((editNoticeInfo: NoticeEditResp) => {
        setEditNoticeInfo(editNoticeInfo.notice_info)
        setEditModalOpen(true)
      })
      .catch((e) => {
        console.log('修改公告 err:', e)
      })
  }

  /**
   * 修改保存操作
   * @param noticeInfo
   */
  const onEditFinishSubmit = (noticeInfo: NoticeInfoType) => {
    NoticeService.modifyNotice(noticeInfo)
      .then(() => {
        message.success('修改成功', 2, () => {
          setEditModalOpen(false)
          getNoticeList(pagination, searchKeyWords)
        })
      })
      .catch(() => {
        message.error('修改失败', 2)
      })
  }

  /**
   * 确认删除操作
   * @param noticeInfo
   */
  const onDeleteConfirm = (noticeInfo: NoticeInfoType) => {
    NoticeService.deleteNotice(noticeInfo.notice_id)
      .then(() => {
        message.success('删除成功', 2, () => {
          getNoticeList(pagination, searchKeyWords)
        })
      })
      .catch((e) => {
        console.log('删除公告失败:', e)
      })
  }

  /**
   * 搜索查询操作
   * @param values
   */
  const onSearchChange = (values: any) => {
    getNoticeList(initPagination, values)
  }

  /**
   * 搜索重置操作
   */
  const onSearchReset = () => {
    getNoticeList(initPagination, {})
  }

  /**
   * 获取公告列表
   * @param pagination
   * @param searchValues
   */
  const getNoticeList = (pagination: TablePaginationConfig, searchValues: {}) => {
    const pageSize = pagination.pageSize
    const current = pagination.current
    searchKeyWords = searchValues
    NoticeService.getNoticeList(pageSize, current, searchValues)
      .then((noticeList: NoticeListResp) => {
        setNoticeList(noticeList.list)
        setPagination({
          ...initPagination,
          current: noticeList.page_info?.page_num,
          pageSize: noticeList.page_info?.page_size,
          total: noticeList.page_info?.total_num
        })
      })
      .catch((e) => {
        console.log(e)
      })
  }

  return (
    <div className="panel">
      <NoticeSearchUI onSearchChange={onSearchChange} onSearchReset={onSearchReset} />
      <NoticeListUI
        listLoading={false}
        pagination={pagination}
        noticeList={noticeList}
        onListChange={onListChange}
        onEditClick={onEditClick}
        onDeleteConfirm={onDeleteConfirm}
      />
      <Modal
        title="公告修改"
        width={570}
        open={editModalOpen}
        onCancel={() => setEditModalOpen(false)}
        footer={null}
      >
        <NoticeFormUI noticeInfo={editNoticeInfo} onFinishSubmit={onEditFinishSubmit} />
      </Modal>
    </div>
  )
}

export default NoticeList
