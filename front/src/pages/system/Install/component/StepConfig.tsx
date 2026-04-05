import React from 'react'
import { Button, Form, Input, Space } from 'antd'

interface StepConfigProps {
  onNext: (values: any) => void
  onPrev: () => void
}

const StepConfig: React.FC<StepConfigProps> = ({ onNext, onPrev }) => {
  const [form] = Form.useForm()

  const onFinish = (values: any) => {
    onNext(values)
  }

  return (
    <Form
      form={form}
      layout="vertical"
      onFinish={onFinish}
      style={{ maxWidth: 460, margin: '0 auto', padding: '24px 0' }}
    >
      <Form.Item
        label="系统名称"
        name="system_name"
        rules={[{ required: true, message: '请输入系统名称' }]}
      >
        <Input placeholder="请输入系统名称" />
      </Form.Item>
      <Form.Item
        label="系统标题"
        name="main_title"
        rules={[{ required: true, message: '请输入系统标题' }]}
      >
        <Input placeholder="请输入系统标题" />
      </Form.Item>
      <Form.Item label="系统描述" name="main_description">
        <Input.TextArea rows={4} placeholder="请输入系统描述" />
      </Form.Item>
      <Form.Item>
        <Space>
          <Button onClick={onPrev}>上一步</Button>
          <Button type="primary" htmlType="submit">
            下一步
          </Button>
        </Space>
      </Form.Item>
    </Form>
  )
}

export default StepConfig
