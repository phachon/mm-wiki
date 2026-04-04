import React from 'react'
import { message } from 'antd'
import { SystemContactService } from '@/services/SystemContact'
import ContactFormUI from '../component/FormUI'

const ContactAdd: React.FC = () => {
  const onSaveSubmit = (values: any) => {
    SystemContactService.saveContact(values).then(() => {
      message.success('添加成功', 2, () => {
        window.location.href = '/system/contact/list'
      })
    })
  }

  return (
    <div className="pdt24">
      <ContactFormUI onSaveSubmit={onSaveSubmit} />
    </div>
  )
}

export default ContactAdd
