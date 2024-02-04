import { EditLayoutForm, LayoutForm } from '@/config/layout'
import {
  Button,
  DatePicker,
  Form,
  Input,
  InputNumber,
  Radio,
  RadioChangeEvent,
  Select,
  Space
} from 'antd'
import { OrderInfoType } from '@/types/orderType'
import { OrderPayTypeRadioOptions, OrderTypeRadioOptions } from './ToolsUI'
import { DefaultOptionType } from 'antd/es/select'
import 'dayjs/locale/zh-cn'
import locale from 'antd/es/date-picker/locale/zh_CN'
import { HirerInfoType } from '@/types/hirerType'
import { useEffect, useState } from 'react'
import dayjs from 'dayjs'
import { HouseInfoType } from '@/types/houseType'
import { RoomInfoType } from '@/types/roomType'
import { HouseShowWholeAddress } from '@/pages/House/component/ToolsUI'
import type { Dayjs } from 'dayjs'
import { now } from 'moment'
import { getNowStartTimestamp } from '@/utils/utils'
import TextArea from 'antd/es/input/TextArea'

const dateFormat = 'YYYY-MM-DD'

interface OrderFormUIProps {
  /**
   * 订单信息
   */
  orderInfo?: OrderInfoType

  /**
   * 选择租客列表
   */
  selectHirerList?: HirerInfoType[]

  /**
   * 默认选中的租客
   */
  defaultSelectHirerValue?: string

  /**
   * 选择房产列表
   */
  selectHouseList?: HouseInfoType[]

  /**
   * 选择房产下房间列表
   */
  houseRooms?: Record<number, RoomInfoType[]>

  /**
   * 表单的布局配置
   */
  formLayout?: {
    labelCol: { span: number }
    wrapperCol: { span: number }
  }

  /**
   * 新增租客点击方法
   * @param e 点击事件对象
   */
  onAddHirerClick?: (e: any) => void

  /**
   * 订单保存方法
   */
  onSaveFinish: (values: OrderInfoType) => void

  /**
   * 订单类型改变
   * @param e 点击对象
   */
  onOrderTypeChange(e: RadioChangeEvent): void
}

/**
 * 获取租客 Option 数据
 * @param hirerList 租客列表
 * @returns
 */
const getHirerOptions = (hirerList?: HirerInfoType[]): DefaultOptionType[] => {
  let options: DefaultOptionType[] = []
  if (!hirerList) {
    return options
  }
  hirerList.forEach((hirerInfo) => {
    options.push({
      label: hirerInfo.name + '(' + hirerInfo.mobile + ')',
      value: String(hirerInfo.hirer_id)
    })
  })
  return options
}

const getSelectHousesOptions = (houseList?: HouseInfoType[]): DefaultOptionType[] => {
  let options: DefaultOptionType[] = []
  if (!houseList) {
    return options
  }
  houseList.forEach((houseInfo) => {
    options.push({
      label: HouseShowWholeAddress(houseInfo),
      value: String(houseInfo.house_id)
    })
  })
  return options
}

const getSelectRoomsOptions = (roomList?: RoomInfoType[]): DefaultOptionType[] => {
  let options: DefaultOptionType[] = []
  if (!roomList) {
    return options
  }
  roomList.forEach((roomInfo) => {
    options.push({
      label: roomInfo.name,
      value: String(roomInfo.room_id)
    })
  })
  return options
}

/**
 * 订单表单 UI
 * @param props
 */
const OrderFormUI = (props: OrderFormUIProps) => {
  const [form] = Form.useForm()
  const payTypeOptions = OrderPayTypeRadioOptions()
  const orderTypeOptions = OrderTypeRadioOptions()
  const [rooms, setRooms] = useState<RoomInfoType[]>([])
  const currentDate = getNowStartTimestamp()

  let layoutForm = LayoutForm

  /**
   * 下拉选择房间事件
   * @param e 点击事件
   */
  const onSelectHouseChange = (value: string): void => {
    const houseId = parseInt(value)
    // 判断是否有房间 - 合租
    if (!props.houseRooms) {
      return
    }
    // 重新设置房间列表
    form.setFieldValue('room_id', null)
    const selectRooms = props.houseRooms[houseId]
    setRooms(selectRooms)
  }

  /**
   * 订单类型改变操作
   * @param e 点击事件
   */
  const onOrderTypeChange = (e: RadioChangeEvent): void => {
    // 重新设置房屋列表
    form.setFieldValue('house_id', null)
    form.setFieldValue('room_id', null)
    props.onOrderTypeChange(e)
  }

  /**
   * 开始时间范围检查
   * @param current
   * @returns
   */
  const disabledStartDate = (current: Dayjs) => {
    return current.isBefore(currentDate)
  }

  /**
   * 结束时间范围检查
   * @param current
   * @returns
   */
  const disabledEndDate = (current: Dayjs) => {
    return current.isBefore(currentDate)
  }

  /**
   * 表单提交
   * @param values
   */
  const onSaveSubmit = (values: any) => {
    props.onSaveFinish({
      ...values,
      house_id: parseInt(values.house_id),
      hirer_id: parseInt(values.hirer_id),
      start_time: values.start_time.format(dateFormat),
      end_time: values.end_time.format(dateFormat)
    })
  }

  return (
    <div className="panel-body">
      <Form name="basic" {...layoutForm} form={form} onFinish={onSaveSubmit}>
        <Form.Item
          label="订单类型"
          name="order_type"
          rules={[{ required: true, message: '请选择订单类型!' }]}
          initialValue={1}
        >
          <Radio.Group options={orderTypeOptions} onChange={onOrderTypeChange} />
        </Form.Item>

        <Form.Item label="选择房屋" name="house_id" rules={[{ required: true }]}>
          <Select
            allowClear
            style={{ width: '100%' }}
            placeholder="请选择房屋"
            options={getSelectHousesOptions(props.selectHouseList)}
            showSearch
            onChange={onSelectHouseChange}
          />
        </Form.Item>

        <Form.Item dependencies={['order_type']} noStyle>
          {({ getFieldValue }) =>
            getFieldValue('order_type') == 1 ? null : (
              <Form.Item label="选择房间" name="room_id" rules={[{ required: true }]}>
                <Select
                  allowClear
                  style={{ width: '100%' }}
                  placeholder="请选择房间"
                  options={getSelectRoomsOptions(rooms)}
                  showSearch
                />
              </Form.Item>
            )
          }
        </Form.Item>

        <Form.Item
          label="选择租客"
          name="hirer_id"
          rules={[{ required: true, message: '请选择租客!' }]}
        >
          <Space.Compact style={{ width: '100%' }}>
            <Form.Item name="hirer_id" noStyle>
              <Select
                placeholder="请选择租客"
                allowClear
                showSearch
                options={getHirerOptions(props.selectHirerList)}
              />
            </Form.Item>
            <Button onClick={props.onAddHirerClick}>新增租客</Button>
          </Space.Compact>
        </Form.Item>

        <Form.Item
          label="开始日期"
          name="start_time"
          rules={[{ required: true, message: '请选择开始日期!' }]}
          initialValue={dayjs()}
        >
          <DatePicker
            picker={'date'}
            locale={locale}
            placeholder="请选择开始日期"
            style={{ width: '100%' }}
            format={dateFormat}
            disabledDate={disabledStartDate}
            defaultValue={dayjs()}
          />
        </Form.Item>

        <Form.Item
          label="结束日期"
          name="end_time"
          rules={[{ required: true, message: '请选择结束日期!' }]}
          initialValue={dayjs()}
        >
          <DatePicker
            picker={'date'}
            locale={locale}
            placeholder="请选择结束日期"
            style={{ width: '100%' }}
            format={dateFormat}
            disabledDate={disabledEndDate}
            defaultValue={dayjs()}
          />
        </Form.Item>

        <Form.Item
          label="支付方式"
          name="pay_type"
          rules={[{ required: true, message: '请选择房间类型!' }]}
          initialValue={1}
        >
          <Radio.Group options={payTypeOptions} />
        </Form.Item>

        <Form.Item
          name="unit_rent"
          label="月付租金"
          rules={[{ required: true, message: '请输入单位租金!' }]}
        >
          <InputNumber style={{ width: '100%' }} placeholder="请输入月付租金" addonAfter={'CNY'} />
        </Form.Item>

        <Form.Item name="remarks" label="备注信息">
          <TextArea style={{ width: '100%', height: 100 }} placeholder="请输入备注" />
        </Form.Item>
        <Form.Item wrapperCol={{ offset: layoutForm.labelCol.span }}>
          <Button type="primary" htmlType="submit">
            保存
          </Button>
        </Form.Item>
      </Form>
    </div>
  )
}

export default OrderFormUI
