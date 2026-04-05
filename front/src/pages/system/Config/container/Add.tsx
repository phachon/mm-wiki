import React, { useEffect, useState } from 'react'
import { message } from 'antd'
import { SystemConfigService } from '@/services/SystemConfig'
import ConfigFormUI from '../component/FormUI'

const ConfigAdd: React.FC = () => {
  const [configMap, setConfigMap] = useState<{ [key: string]: string }>()

  useEffect(() => {
    SystemConfigService.getConfigList().then((res) => {
      setConfigMap(res)
    })
  }, [])

  /**
   * 保存系统配置
   * @param values 配置数据
   */
  const onSaveSubmit = (values: any) => {
    SystemConfigService.modifyConfig(values).then(() => {
      message.success('保存成功')
    })
  }

  return (
    <div className="pdt24">
      <ConfigFormUI configMap={configMap} onSaveSubmit={onSaveSubmit} />
    </div>
  )
}

export default ConfigAdd
