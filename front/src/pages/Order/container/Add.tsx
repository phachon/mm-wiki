import { Modal, RadioChangeEvent, message } from 'antd'
import { UserOutlined, UsergroupAddOutlined } from '@ant-design/icons'
import OrderFormUI from '../component/FormUI'
import { useEffect, useState } from 'react'
import { HirerService } from '@/services/Hirer'
import { HirerInfoType, HirerListResp, HirerSaveResp } from '@/types/hirerType'
import HirerFormUI from '@/pages/Hirer/component/FormUI'
import { EditLayoutForm } from '@/config/layout'
import { OrderService } from '@/services/Order'
import { OrderAddResp, OrderInfoType } from '@/types/orderType'
import { HouseInfoType } from '@/types/houseType'
import { RoomInfoType } from '@/types/roomType'
import { useNavigate } from 'react-router-dom'

// OrderAdd 添加订单操作
const OrderAdd: React.FC = () => {
  const [selectHirerList, setSelectHirerList] = useState<HirerInfoType[]>([])
  const [hirerAddModalOpen, setHirerAddModalOpen] = useState<boolean>(false)
  const [defaultSelectHirerValue, setDefaultSelectHirerValue] = useState<string>()
  const [selectHouseList, setSelectHouseList] = useState<HouseInfoType[]>([])
  const [selectHouseRooms, setSelectHouseRooms] = useState<Record<number, RoomInfoType[]>>()
  const navigate = useNavigate()

  useEffect(() => {
    getHirerList()
    getOrderAddInfo(1) // 默认整租
  }, [])

  const getOrderAddInfo = (orderType: number) => {
    OrderService.getAddOrderInfo(orderType)
      .then((resp: OrderAddResp) => {
        setSelectHouseList(resp.select_houses)
        setSelectHouseRooms(resp.select_house_rooms)
        console.log(resp.select_house_rooms.constructor)
      })
      .catch((e) => {
        console.log('获取订单页添加信息失败:', e)
      })
  }

  const getHirerList = () => {
    HirerService.getAllHirers()
      .then((resp: HirerListResp) => {
        setSelectHirerList(resp.list)
      })
      .catch((e) => {
        console.log('获取租客列表失败:', e)
      })
  }

  /**
   * 新增租客点击
   * @param e 点击事件
   */
  const onAddHirerClick = (e: any) => {
    setDefaultSelectHirerValue(undefined) // 清除选中的租客
    setHirerAddModalOpen(true) // 打开弹框
  }

  /**
   * 订单类型选择改变
   * @param e 点击对象
   */
  const onOrderTypeChange = (e: RadioChangeEvent) => {
    const orderType = e.target.value
    getOrderAddInfo(orderType)
  }

  /**
   * 订单保存
   * @param values
   */
  const onSaveFinish = (values: OrderInfoType) => {
    OrderService.saveOrder(values)
      .then(() => {
        message.success('保存成功', 2, () => {
          navigate('/order/list')
        })
      })
      .catch((e) => {
        console.log('订单保存 err:', e)
      })
  }

  /**
   * 添加租客保存
   * @param values
   */
  function onAddHirerSave(values: any): void {
    HirerService.saveHirer(values)
      .then((resp: HirerSaveResp) => {
        message.success('保存成功', 2, () => {
          getHirerList() // 重新拉取一次接口返回所有的租客
          setHirerAddModalOpen(false)
          setDefaultSelectHirerValue(resp.hirer_id.toString()) // 默认选中
        })
      })
      .catch((e) => {
        console.log('新增租客保存失败：', e)
      })
  }

  return (
    <div className="pdt24">
      <OrderFormUI
        onSaveFinish={onSaveFinish}
        selectHirerList={selectHirerList}
        onAddHirerClick={onAddHirerClick}
        defaultSelectHirerValue={defaultSelectHirerValue}
        selectHouseList={selectHouseList}
        onOrderTypeChange={onOrderTypeChange}
        houseRooms={selectHouseRooms}
      />
      <Modal
        title="新增租客"
        width={600}
        open={hirerAddModalOpen}
        onCancel={() => setHirerAddModalOpen(false)}
        footer={null}
      >
        <HirerFormUI onHirerSave={onAddHirerSave} formLayout={EditLayoutForm} />
      </Modal>
    </div>
  )
}

export default OrderAdd
