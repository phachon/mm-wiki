import React, { useEffect, useState } from 'react'
import HouseFormUI from '../component/FormUI'
import { Modal, message } from 'antd'
import LandlordFormUI from '@/pages/Landlord/component/FormUI'
import { EditLayoutForm } from '@/config/layout'
import { LandlordService } from '@/services/Landlord'
import { LandlordInfoType, LandlordListResp, LandlordSaveResp } from '@/types/landlordType'
import { HouseService } from '@/services/House'
import { HouseInfoType } from '@/types/houseType'
import { useNavigate } from 'react-router-dom'

const HouseAdd: React.FC = () => {
  const [landlorAddModalOpen, setLandlorAddModalOpen] = useState(false)
  const [landlorTrustModalOpen, setLandlorTrustModalOpen] = useState(false)
  const [landlords, setLandlords] = useState<LandlordInfoType[]>([])
  const [defaultSelectLandlordValue, setDefaultSelectLandlordValue] = useState<string>()
  const navigate = useNavigate()

  useEffect(() => {
    getAllLandlords()
  }, [])

  /**
   * 获取所有的业主列表
   */
  const getAllLandlords = () => {
    LandlordService.getAllLandlords()
      .then((resp: LandlordListResp) => {
        setLandlords(resp.list)
      })
      .catch((e) => {
        console.log('获取所有的业主失败：', e)
      })
  }

  /**
   * 添加房产保存
   * @param values 保存结果
   */
  const onAddHouseSave = (values: HouseInfoType) => {
    HouseService.saveHouse(values)
      .then(() => {
        message.success('保存成功', 2, () => {
          navigate('/house/list')
        })
      })
      .catch((e) => {
        console.log('添加房产失败：', e)
      })
  }

  /**
   * 新增业主点击
   * @param e 点击事件
   */
  const onAddLandlordClick = (e: any) => {
    setDefaultSelectLandlordValue(undefined) // 清除选中的业主
    setLandlorAddModalOpen(true) // 打开弹框
  }

  /**
   * 新增业主保存
   * @param values 保存结果
   */
  const onAddLandlordSave = (values: LandlordInfoType) => {
    LandlordService.saveLandlord(values)
      .then((resp: LandlordSaveResp) => {
        message.success('保存成功', 2, () => {
          getAllLandlords() // 重新拉取一次接口返回所有的业主
          setLandlorAddModalOpen(false) // 关闭弹框
          setDefaultSelectLandlordValue(resp.landlord_id.toString()) // 默认选中
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
    <div>
      <HouseFormUI
        onHouseInfoSave={onAddHouseSave}
        onAddLandlordClick={onAddLandlordClick}
        selectLandlords={landlords}
        defaultSelectLandlordValue={defaultSelectLandlordValue}
      />
      <Modal
        title="新增业主"
        width={600}
        open={landlorAddModalOpen}
        onCancel={() => setLandlorAddModalOpen(false)}
        footer={null}
      >
        <LandlordFormUI onLandlordSave={onAddLandlordSave} formLayout={EditLayoutForm} />
      </Modal>
      {/* <Modal
        title="房产托管"
        width={600}
        open={landlorTrustModalOpen}
        onCancel={() => setLandlorTrustModalOpen(false)}
        footer={null}
      >
        <LandlordFormUI onLandloadSave={onAddLandloadSave} formLayout={EditLayoutForm} />
      </Modal> */}
    </div>
  )
}

export default HouseAdd
