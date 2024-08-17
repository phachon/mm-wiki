import { Button, DatePicker, Form, Input, Switch } from 'antd'
import { EditLayoutForm, LayoutForm } from '@/config/layout'
import { NoticeInfoType } from '@/types/noticeType'
import TextArea from 'antd/lib/input/TextArea'
import moment, { Moment } from 'moment'
import dayjs from 'dayjs'
import type { Dayjs } from 'dayjs'

import type { RangePickerProps } from 'antd/es/date-picker'
import { useEffect } from 'react'

interface NoticeFormUIProps {
  noticeInfo?: NoticeInfoType
  onSaveSubmit: (values: any) => void
  formLayout?: {
    labelCol: { span: number }
    wrapperCol: { span: number }
  }
}

const { RangePicker } = DatePicker
const dateFormat = 'YYYY-MM-DD HH:mm'

const disabledDate = (current: Dayjs) => {
  return current < moment().startOf('day')
}

/**
 * 公告表单 UI 组件
 * @param props
 * @returns
 */
const NoticeFormUI = (props: NoticeFormUIProps) => {
  const [form] = Form.useForm()
  const isEdit = props.noticeInfo ? true : false
  let layoutForm = isEdit ? EditLayoutForm : LayoutForm
  if (props.formLayout) {
    layoutForm = props.formLayout
  }

  useEffect(() => {
    if (props.noticeInfo) {
      props.noticeInfo.range_time = [
        dayjs(props.noticeInfo.start_time, dateFormat),
        dayjs(props.noticeInfo.end_time, dateFormat)
      ]
      form.setFieldsValue(props.noticeInfo)
    }
  }, [props.noticeInfo])

  /**
   * 公共保存提交
   * @param values 数据
   */
  const noticeOnSubmit = (values: any) => {
    const rangeTime = values.range_time
    const startTime = rangeTime[0].format(dateFormat)
    const endTime = rangeTime[1].format(dateFormat)
    values.range_time = []
    props.onSaveSubmit({
      ...values,
      start_time: startTime,
      end_time: endTime
    })
  }

  return (
    <div className="panel-body">
      <Form name="notice-form" {...layoutForm} form={form} onFinish={noticeOnSubmit}>
        {isEdit && (
          <Form.Item label="公告ID" name="notice_id" rules={[{ required: true }]}>
            <Input disabled placeholder="请输入公告ID" />
          </Form.Item>
        )}

        <Form.Item
          label="公告标题"
          name="title"
          rules={[{ required: true, message: '请输入公告标题!' }]}
        >
          <Input placeholder="请输入公告标题" />
        </Form.Item>

        <Form.Item
          label="公告内容"
          name="content"
          rules={[{ required: true, message: '请输入公告内容!' }]}
        >
          <TextArea rows={5} placeholder="请输入公告内容" />
        </Form.Item>

        <Form.Item
          name="range_time"
          label="公告时间"
          rules={[{ required: true, message: '请选择公告时间!' }]}
        >
          <RangePicker
            showTime
            format={dateFormat}
            // disabledDate={disabledDate}
            // disabledTime={disabledRangeTime}
            style={{ width: '100%' }}
          />
        </Form.Item>

        <Form.Item
          label="是否发布"
          name="publish_status"
          valuePropName="checked"
          initialValue={1}
          rules={[{ required: true, message: '是否发布!' }]}
          getValueProps={(value) => ({ checked: value === 1 ? true : false })}
          getValueFromEvent={(value) => {
            return value ? 1 : 0
          }}
        >
          <Switch checkedChildren="是" unCheckedChildren="否" defaultChecked />
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

export default NoticeFormUI
