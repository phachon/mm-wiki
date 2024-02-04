import React, { useEffect, useState } from 'react'
import LandlordFormUI from '../component/FormUI'
import { LandlordInfoType, LandlordSaveResp } from '@/types/landlordType'
import { LandlordService } from '@/services/Landlord'
import { message } from 'antd'
import { useNavigate } from 'react-router-dom'

const LandlordAdd: React.FC = () => {
  const navigate = useNavigate()

  /**
   * 新增业主保存
   * @param values 保存结果
   */
  const onAddLandloadSave = (values: LandlordInfoType) => {
    LandlordService.saveLandlord(values)
      .then((resp: LandlordSaveResp) => {
        message.success('保存成功', 2, () => {
          navigate('/landlord/list')
        })
      })
      .catch((e) => {
        console.log('新增业主保存失败：', e)
      })
  }

  /**
   * 返回组件
   */
  return (
    <div className="pdt24">
      <LandlordFormUI onLandlordSave={onAddLandloadSave} />
    </div>
  )
}

export default LandlordAdd
