import React from 'react'
import { Button, Form, Input, Space } from 'antd'

interface StepAdminProps {
  onNext: (values: any) => void
  onPrev: () => void
}

const StepAdmin: React.FC<StepAdminProps> = ({ onNext, onPrev }) => {
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
        label="管理员账号"
        name="account_name"
        rules={[{ required: true, message: '请输入管理员账号' }]}
      >
        <Input placeholder="请输入管理员账号" />
      </Form.Item>
      <Form.Item
        label="密码"
        name="password"
        rules={[
          { required: true, message: '请输入密码' },
          { min: 6, message: '密码长度不能少于6位' }
        ]}
      >
        <Input.Password placeholder="请输入密码" />
      </Form.Item>
      <Form.Item
        label="邮箱"
        name="email"
        rules={[{ type: 'email', message: '请输入有效的邮箱地址' }]}
      >
        <Input placeholder="请输入邮箱" />
      </Form.Item>
      <Form.Item label="昵称" name="given_name">
        <Input placeholder="请输入昵称" />
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

export default StepAdmin
