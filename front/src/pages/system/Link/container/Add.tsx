import React from 'react'
import { message } from 'antd'
import { SystemLinkService } from '@/services/SystemLink'
import LinkFormUI from '../component/FormUI'

const LinkAdd: React.FC = () => {
  const onSaveSubmit = (values: any) => {
    SystemLinkService.saveLink(values).then(() => {
      message.success('添加成功', 2, () => {
        window.location.href = '/system/link/list'
      })
    })
  }

  return (
    <div className="pdt24">
      <LinkFormUI onSaveSubmit={onSaveSubmit} />
    </div>
  )
}

export default LinkAdd
