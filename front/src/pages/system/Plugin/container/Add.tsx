import React from 'react'
import { message } from 'antd'
import { SystemPluginService } from '@/services/SystemPlugin'
import PluginFormUI from '../component/FormUI'

const PluginAdd: React.FC = () => {
  const onSaveSubmit = (values: any) => {
    SystemPluginService.savePlugin(values).then(() => {
      message.success('添加成功', 2, () => {
        window.location.href = '/system/plugin/list'
      })
    })
  }

  return (
    <div className="pdt24">
      <PluginFormUI onSaveSubmit={onSaveSubmit} />
    </div>
  )
}

export default PluginAdd
