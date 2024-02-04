import React, { useEffect, useState } from 'react'
import HirerFormUI from '../component/FormUI'
import { HirerInfoType, HirerSaveResp } from '@/types/hirerType'
import { HirerService } from '@/services/Hirer'
import { message } from 'antd'
import { useNavigate } from 'react-router-dom'

const HirerAdd: React.FC = () => {
  const navigate = useNavigate()

  /**
   * 新增租客保存
   * @param values 保存结果
   */
  const onAddLandloadSave = (values: HirerInfoType) => {
    HirerService.saveHirer(values)
      .then((resp: HirerSaveResp) => {
        message.success('保存成功', 2, () => {
          navigate('/hirer/list')
        })
      })
      .catch((e) => {
        console.log('新增租客保存失败：', e)
      })
  }

  /**
   * 返回组件
   */
  return (
    <div className="pdt24">
      <HirerFormUI onHirerSave={onAddLandloadSave} />
    </div>
  )
}

export default HirerAdd
