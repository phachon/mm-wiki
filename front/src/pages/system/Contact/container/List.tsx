import React, { useEffect, useState } from 'react'
import {
  ContactEditResp,
  ContactInfoType,
  ContactListItemType,
  ContactListResp
} from '@/types/contactType'
import { SystemContactService } from '@/services/SystemContact'
import { message, Modal, TablePaginationConfig } from 'antd'
import ContactListUI from '../component/ListUI'
import ContactSearchUI from '../component/SearchUI'
import ContactFormUI from '../component/FormUI'
import { initPagination } from '@/types/adminType'

let searchKeyWords = {}

const ContactList: React.FC = () => {
  const [contactList, setContactList] = useState<ContactListItemType[]>([])
  const [pagination, setPagination] = useState(initPagination)
  const [editModalOpen, setEditModalOpen] = useState<boolean>(false)
  const [editContactInfo, setEditContactInfo] = useState<ContactInfoType>()

  useEffect(() => {
    getContactList(initPagination, {})
  }, [])

  const onListChange = (pageConfig: TablePaginationConfig, filters: any, sorter: any) => {
    getContactList(pageConfig, searchKeyWords)
  }

  const onEditClick = (contactInfo: ContactInfoType) => {
    SystemContactService.getEditContactInfo(contactInfo.contact_id)
      .then((editContactInfo: ContactEditResp) => {
        setEditContactInfo(editContactInfo.contact_info)
        setEditModalOpen(true)
      })
      .catch((e) => {
        console.log('修改联系人 err:', e)
      })
  }

  const onEditSaveSubmit = (contactInfo: ContactInfoType) => {
    SystemContactService.modifyContact(contactInfo)
      .then(() => {
        message.success('修改成功', 2, () => {
          setEditModalOpen(false)
          getContactList(pagination, searchKeyWords)
        })
      })
      .catch(() => {
        message.error('修改失败', 2)
      })
  }

  const onDeleteConfirm = (contactInfo: ContactInfoType) => {
    SystemContactService.deleteContact(contactInfo.contact_id)
      .then(() => {
        message.success('删除成功', 2, () => {
          getContactList(pagination, searchKeyWords)
        })
      })
      .catch((e) => {
        console.log('删除联系人失败:', e)
      })
  }

  const onSearchChange = (values: any) => {
    getContactList(initPagination, values)
  }

  const onSearchReset = () => {
    getContactList(initPagination, {})
  }

  const getContactList = (pagination: TablePaginationConfig, searchValues: {}) => {
    const pageSize = pagination.pageSize
    const current = pagination.current
    searchKeyWords = searchValues
    SystemContactService.getContactList(pageSize, current, searchValues)
      .then((contactList: ContactListResp) => {
        setContactList(contactList.list)
        setPagination({
          ...initPagination,
          current: contactList.page_info?.page_num,
          pageSize: contactList.page_info?.page_size,
          total: contactList.page_info?.total_num
        })
      })
      .catch((e) => {
        console.log(e)
      })
  }

  return (
    <div className="panel">
      <ContactSearchUI onSearchChange={onSearchChange} onSearchReset={onSearchReset} />
      <ContactListUI
        listLoading={false}
        pagination={pagination}
        contactList={contactList}
        onListChange={onListChange}
        onEditClick={onEditClick}
        onDeleteConfirm={onDeleteConfirm}
      />
      <Modal
        title="联系人修改"
        width={570}
        open={editModalOpen}
        onCancel={() => setEditModalOpen(false)}
        footer={null}
      >
        <ContactFormUI contactInfo={editContactInfo} onSaveSubmit={onEditSaveSubmit} />
      </Modal>
    </div>
  )
}

export default ContactList
