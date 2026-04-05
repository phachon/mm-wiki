import { Button, Form, Input } from 'antd'
import { EditLayoutForm, LayoutForm } from '@/config/layout'
import { LoginAuthInfoType } from '@/types/loginAuthType'
import { useEffect } from 'react'
import TextArea from 'antd/lib/input/TextArea'

interface LoginAuthFormUIProps {
  loginAuthInfo?: LoginAuthInfoType
  onSaveSubmit: (values: any) => void
  formLayout?: {
    labelCol: { span: number }
    wrapperCol: { span: number }
  }
}

const LoginAuthFormUI = (props: LoginAuthFormUIProps) => {
  const [form] = Form.useForm()
  const isEdit = props.loginAuthInfo ? true : false
  let layoutForm = isEdit ? EditLayoutForm : LayoutForm
  if (props.formLayout) {
    layoutForm = props.formLayout
  }

  useEffect(() => {
    if (props.loginAuthInfo) {
      form.setFieldsValue(props.loginAuthInfo)
    }
  }, [props.loginAuthInfo])

  return (
    <div className="panel-body">
      <Form name="login-auth-form" {...layoutForm} form={form} onFinish={props.onSaveSubmit}>
        {isEdit && (
          <Form.Item label="认证ID" name="login_auth_id" rules={[{ required: true }]}>
            <Input disabled placeholder="请输入认证ID" />
          </Form.Item>
        )}

        <Form.Item
          label="认证名称"
          name="name"
          rules={[{ required: true, message: '请输入认证名称!' }]}
        >
          <Input placeholder="请输入登录认证名称" />
        </Form.Item>

        <Form.Item label="账号前缀" name="account_prefix">
          <Input placeholder="请输入账号登录前缀" />
        </Form.Item>

        <Form.Item
          label="认证接口"
          name="url"
          rules={[{ required: true, message: '请输入认证接口URL!' }]}
        >
          <Input placeholder="请输入认证接口 URL" />
        </Form.Item>

        <Form.Item label="额外数据" name="ext_data">
          <TextArea rows={4} placeholder="请输入额外数据（JSON格式）" />
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

export default LoginAuthFormUI
