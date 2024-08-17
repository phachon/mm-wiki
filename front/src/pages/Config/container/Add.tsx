import React, { useEffect, useState } from 'react'
import ConfigFormUI from '../component/FormUI'

const ConfigAdd: React.FC = () => {
  /**
   * 添加配置保存
   * @param accountInfo
   */
  const onSaveSubmit = (values: any) => {}

  /**
   * 返回组件
   */
  return (
    <div className="pdt24">
      <ConfigFormUI onSaveSubmit={onSaveSubmit} />
    </div>
  )
}

export default ConfigAdd
