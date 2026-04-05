import React from 'react'
import { message } from 'antd'
import { SystemEmailService } from '@/services/SystemEmail'
import EmailFormUI from '../component/FormUI'

const EmailAdd: React.FC = () => {
  const onSaveSubmit = (values: any) => {
    SystemEmailService.saveEmail(values).then(() => {
      message.success('添加成功', 2, () => {
        window.location.href = '/system/email/list'
      })
    })
  }

  return (
    <div className="pdt24">
      <EmailFormUI onSaveSubmit={onSaveSubmit} />
    </div>
  )
}

export default EmailAdd
