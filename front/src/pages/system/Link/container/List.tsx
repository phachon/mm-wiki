import React, { useEffect, useState } from 'react'
import { LinkEditResp, LinkInfoType, LinkListItemType, LinkListResp } from '@/types/linkType'
import { SystemLinkService } from '@/services/SystemLink'
import { message, Modal, TablePaginationConfig } from 'antd'
import LinkListUI from '../component/ListUI'
import LinkSearchUI from '../component/SearchUI'
import LinkFormUI from '../component/FormUI'
import { initPagination } from '@/types/adminType'

let searchKeyWords = {}

const LinkList: React.FC = () => {
  const [linkList, setLinkList] = useState<LinkListItemType[]>([])
  const [pagination, setPagination] = useState(initPagination)
  const [editModalOpen, setEditModalOpen] = useState<boolean>(false)
  const [editLinkInfo, setEditLinkInfo] = useState<LinkInfoType>()

  useEffect(() => {
    getLinkList(initPagination, {})
  }, [])

  const onListChange = (pageConfig: TablePaginationConfig, filters: any, sorter: any) => {
    getLinkList(pageConfig, searchKeyWords)
  }

  const onEditClick = (linkInfo: LinkInfoType) => {
    SystemLinkService.getEditLinkInfo(linkInfo.link_id)
      .then((editLinkInfo: LinkEditResp) => {
        setEditLinkInfo(editLinkInfo.link_info)
        setEditModalOpen(true)
      })
      .catch((e) => {
        console.log('修改链接 err:', e)
      })
  }

  const onEditSaveSubmit = (linkInfo: LinkInfoType) => {
    SystemLinkService.modifyLink(linkInfo)
      .then(() => {
        message.success('修改成功', 2, () => {
          setEditModalOpen(false)
          getLinkList(pagination, searchKeyWords)
        })
      })
      .catch(() => {
        message.error('修改失败', 2)
      })
  }

  const onDeleteConfirm = (linkInfo: LinkInfoType) => {
    SystemLinkService.deleteLink(linkInfo.link_id)
      .then(() => {
        message.success('删除成功', 2, () => {
          getLinkList(pagination, searchKeyWords)
        })
      })
      .catch((e) => {
        console.log('删除链接失败:', e)
      })
  }

  const onSearchChange = (values: any) => {
    getLinkList(initPagination, values)
  }

  const onSearchReset = () => {
    getLinkList(initPagination, {})
  }

  const getLinkList = (pagination: TablePaginationConfig, searchValues: {}) => {
    const pageSize = pagination.pageSize
    const current = pagination.current
    searchKeyWords = searchValues
    SystemLinkService.getLinkList(pageSize, current, searchValues)
      .then((linkList: LinkListResp) => {
        setLinkList(linkList.list)
        setPagination({
          ...initPagination,
          current: linkList.page_info?.page_num,
          pageSize: linkList.page_info?.page_size,
          total: linkList.page_info?.total_num
        })
      })
      .catch((e) => {
        console.log(e)
      })
  }

  return (
    <div className="panel">
      <LinkSearchUI onSearchChange={onSearchChange} onSearchReset={onSearchReset} />
      <LinkListUI
        listLoading={false}
        pagination={pagination}
        linkList={linkList}
        onListChange={onListChange}
        onEditClick={onEditClick}
        onDeleteConfirm={onDeleteConfirm}
      />
      <Modal
        title="链接修改"
        width={570}
        open={editModalOpen}
        onCancel={() => setEditModalOpen(false)}
        footer={null}
      >
        <LinkFormUI linkInfo={editLinkInfo} onSaveSubmit={onEditSaveSubmit} />
      </Modal>
    </div>
  )
}

export default LinkList
