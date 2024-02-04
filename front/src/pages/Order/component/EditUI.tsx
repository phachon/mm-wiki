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
import { OrderPayTypeRadioOptions, OrderStatusRadioOptions, OrderTypeRadioOptions } from './ToolsUI'
import 'dayjs/locale/zh-cn'
import locale from 'antd/es/date-picker/locale/zh_CN'
import { useEffect, useState } from 'react'
import dayjs from 'dayjs'
import type { Dayjs } from 'dayjs'
import { getNowStartTimestamp } from '@/utils/utils'
import TextArea from 'antd/es/input/TextArea'

const dateFormat = 'YYYY-MM-DD'

interface OrderEditUIProps {
  /**
   * 订单信息
   */
  orderInfo?: OrderInfoType

  /**
   * 修改保存
   */
  onEditSave: (values: OrderInfoType) => void
}

/**
 * 订单修改 UI
 * @param props
 */
const OrderEditUI = (props: OrderEditUIProps) => {
  const [form] = Form.useForm()
  const orderStatusOptions = OrderStatusRadioOptions(props.orderInfo?.status)
  const payTypeOptions = OrderPayTypeRadioOptions()
  const currentDate = getNowStartTimestamp()

  useEffect(() => {
    if (props.orderInfo) {
      form.setFieldsValue({
        ...props.orderInfo,
        start_time: dayjs(props.orderInfo?.start_time, dateFormat),
        end_time: dayjs(props.orderInfo?.end_time, dateFormat)
      })
      return
    }
    form.setFieldsValue({})
  }, [props.orderInfo])

  /**
   * 订单类型改变操作
   * @param e 点击事件
   */
  const onOrderStatusChange = (e: RadioChangeEvent): void => {}

  /**
   * 开始时间范围检查
   * @param current
   * @returns
   */
  const disabledStartDate = (current: Dayjs) => {
    return !current.isBefore(currentDate)
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
    props.onEditSave({
      ...values,
      start_time: values.start_time.format(dateFormat),
      end_time: values.end_time.format(dateFormat)
    })
  }

  return (
    <div className="panel-body">
      <Form name="basic" {...EditLayoutForm} form={form} onFinish={onSaveSubmit}>
        <Form.Item label="订单 ID" name="order_id" rules={[{ required: true }]}>
          <Input disabled placeholder="请输入订单ID" />
        </Form.Item>

        {orderStatusOptions.length > 0 ? (
          <Form.Item
            label="订单状态"
            name="status"
            rules={[{ required: true, message: '请选择订单类型!' }]}
            initialValue={1}
          >
            <Radio.Group options={orderStatusOptions} onChange={onOrderStatusChange} />
          </Form.Item>
        ) : null}

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

        <Form.Item
          name="remarks"
          label="修改备注"
          rules={[{ required: true, message: '请输入修改备注!' }]}
        >
          <TextArea style={{ width: '100%', height: 100 }} placeholder="请输入修改备注" />
        </Form.Item>

        <Form.Item wrapperCol={{ offset: EditLayoutForm.labelCol.span }}>
          <Button type="primary" htmlType="submit">
            确定
          </Button>
        </Form.Item>
      </Form>
    </div>
  )
}

export default OrderEditUI
