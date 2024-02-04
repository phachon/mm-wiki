import { AccountInfoType } from '@/types/accountType'
import { Button, Form, Input } from 'antd'
import { LayoutForm, LayoutFormButton } from '../../../config/layout'
import { useEffect } from 'react'

interface ProfileInfoUIProps {
  accountInfo?: AccountInfoType
  onFinishCallback: (values: any) => void
}

/**
 * 个人信息 UI 组件
 * @param props
 * @returns
 */
const ProfileInfoUI = (props: ProfileInfoUIProps) => {
  const [form] = Form.useForm()

  useEffect(() => {
    form.setFieldsValue(props.accountInfo)
  }, [props.accountInfo])

  return (
    <div className="panel-body">
      <Form {...LayoutForm} name="basic" form={form} onFinish={props.onFinishCallback}>
        <Form.Item label="账号ID" name="account_id" rules={[{ required: true }]}>
          <Input disabled />
        </Form.Item>

        <Form.Item
          label="账号名"
          name="name"
          rules={[{ required: true, message: '请输入账号名!' }]}
        >
          <Input disabled placeholder="请输入账号名" />
        </Form.Item>

        <Form.Item
          label="昵称"
          name="given_name"
          rules={[{ required: true, message: '请输入昵称!' }]}
        >
          <Input placeholder="请输入昵称" />
        </Form.Item>

        <Form.Item label="邮箱" name="email">
          <Input placeholder="请输入邮箱地址：xxx@xxx.com" />
        </Form.Item>

        <Form.Item label="电话号" name="phone">
          <Input placeholder="请输入电话号码" />
        </Form.Item>

        <Form.Item label="手机号" name="mobile">
          <Input placeholder="请输入手机号码" />
        </Form.Item>

        <Form.Item label="创建时间" name="create_time">
          <Input disabled />
        </Form.Item>

        <Form.Item label="修改时间" name="update_time">
          <Input disabled />
        </Form.Item>

        <Form.Item {...LayoutFormButton}>
          <Button type="primary" htmlType="submit">
            保存
          </Button>
        </Form.Item>
      </Form>
    </div>
  )
}

export default ProfileInfoUI
