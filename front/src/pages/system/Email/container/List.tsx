import React, { useEffect, useState } from 'react'
import {
  EmailEditResp,
  EmailInfoType,
  EmailListItemType,
  EmailListResp
} from '@/types/emailType'
import { SystemEmailService } from '@/services/SystemEmail'
import { message, Modal, TablePaginationConfig } from 'antd'
import EmailListUI from '../component/ListUI'
import EmailSearchUI from '../component/SearchUI'
import EmailFormUI from '../component/FormUI'
import { initPagination } from '@/types/adminType'

let searchKeyWords = {}

const EmailList: React.FC = () => {
  const [emailList, setEmailList] = useState<EmailListItemType[]>([])
  const [pagination, setPagination] = useState(initPagination)
  const [editModalOpen, setEditModalOpen] = useState<boolean>(false)
  const [editEmailInfo, setEditEmailInfo] = useState<EmailInfoType>()

  useEffect(() => {
    getEmailList(initPagination, {})
  }, [])

  const onListChange = (pageConfig: TablePaginationConfig, filters: any, sorter: any) => {
    getEmailList(pageConfig, searchKeyWords)
  }

  const onEditClick = (emailInfo: EmailInfoType) => {
    SystemEmailService.getEditEmailInfo(emailInfo.email_id)
      .then((editEmailInfo: EmailEditResp) => {
        setEditEmailInfo(editEmailInfo.email_info)
        setEditModalOpen(true)
      })
      .catch((e) => {
        console.log('修改邮箱 err:', e)
      })
  }

  const onEditSaveSubmit = (emailInfo: EmailInfoType) => {
    SystemEmailService.modifyEmail(emailInfo)
      .then(() => {
        message.success('修改成功', 2, () => {
          setEditModalOpen(false)
          getEmailList(pagination, searchKeyWords)
        })
      })
      .catch(() => {
        message.error('修改失败', 2)
      })
  }

  const onDeleteConfirm = (emailInfo: EmailInfoType) => {
    SystemEmailService.deleteEmail(emailInfo.email_id)
      .then(() => {
        message.success('删除成功', 2, () => {
          getEmailList(pagination, searchKeyWords)
        })
      })
      .catch((e) => {
        console.log('删除邮箱失败:', e)
      })
  }

  const onUsedClick = (emailInfo: EmailInfoType) => {
    SystemEmailService.setEmailUsed(emailInfo.email_id)
      .then(() => {
        message.success('设置成功', 2, () => {
          getEmailList(pagination, searchKeyWords)
        })
      })
      .catch((e) => {
        console.log('设置使用失败:', e)
      })
  }

  const onSearchChange = (values: any) => {
    getEmailList(initPagination, values)
  }

  const onSearchReset = () => {
    getEmailList(initPagination, {})
  }

  const getEmailList = (pagination: TablePaginationConfig, searchValues: {}) => {
    const pageSize = pagination.pageSize
    const current = pagination.current
    searchKeyWords = searchValues
    SystemEmailService.getEmailList(pageSize, current, searchValues)
      .then((emailList: EmailListResp) => {
        setEmailList(emailList.list)
        setPagination({
          ...initPagination,
          current: emailList.page_info?.page_num,
          pageSize: emailList.page_info?.page_size,
          total: emailList.page_info?.total_num
        })
      })
      .catch((e) => {
        console.log(e)
      })
  }

  return (
    <div className="panel">
      <EmailSearchUI onSearchChange={onSearchChange} onSearchReset={onSearchReset} />
      <EmailListUI
        listLoading={false}
        pagination={pagination}
        emailList={emailList}
        onListChange={onListChange}
        onEditClick={onEditClick}
        onDeleteConfirm={onDeleteConfirm}
        onUsedClick={onUsedClick}
      />
      <Modal
        title="邮箱修改"
        width={570}
        open={editModalOpen}
        onCancel={() => setEditModalOpen(false)}
        footer={null}
      >
        <EmailFormUI emailInfo={editEmailInfo} onSaveSubmit={onEditSaveSubmit} />
      </Modal>
    </div>
  )
}

export default EmailList
