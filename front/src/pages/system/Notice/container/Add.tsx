import React from 'react'
import { message } from 'antd'
import { NoticeService } from '@/services/Notice'
import NoticeFormUI from '../component/FormUI'

const NoticeAdd: React.FC = () => {
  /**
   * 添加操作
   * @param values
   */
  const onSaveSubmit = (values: any) => {
    NoticeService.saveNotice(values).then(() => {
      message.success('添加成功', 2, () => {
        window.location.href = '/notice/list'
      })
    })
  }

  return (
    <div className="pdt24">
      <NoticeFormUI onSaveSubmit={onSaveSubmit} />
    </div>
  )
}

export default NoticeAdd
