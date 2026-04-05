import { Button, Form, Input, InputNumber, Switch } from 'antd'
import { EditLayoutForm, LayoutForm } from '@/config/layout'
import { EmailInfoType } from '@/types/emailType'
import { useEffect } from 'react'

interface EmailFormUIProps {
  emailInfo?: EmailInfoType
  onSaveSubmit: (values: any) => void
  formLayout?: {
    labelCol: { span: number }
    wrapperCol: { span: number }
  }
}

const EmailFormUI = (props: EmailFormUIProps) => {
  const [form] = Form.useForm()
  const isEdit = props.emailInfo ? true : false
  let layoutForm = isEdit ? EditLayoutForm : LayoutForm
  if (props.formLayout) {
    layoutForm = props.formLayout
  }

  useEffect(() => {
    if (props.emailInfo) {
      form.setFieldsValue(props.emailInfo)
    }
  }, [props.emailInfo])

  return (
    <div className="panel-body">
      <Form name="email-form" {...layoutForm} form={form} onFinish={props.onSaveSubmit}>
        {isEdit && (
          <Form.Item label="邮箱ID" name="email_id" rules={[{ required: true }]}>
            <Input disabled placeholder="请输入邮箱ID" />
          </Form.Item>
        )}

        <Form.Item
          label="服务器名称"
          name="name"
          rules={[{ required: true, message: '请输入邮箱服务器名称!' }]}
        >
          <Input placeholder="请输入邮箱服务器名称" />
        </Form.Item>

        <Form.Item
          label="发件人地址"
          name="sender_address"
          rules={[{ required: true, message: '请输入发件人邮箱地址!' }]}
        >
          <Input placeholder="请输入发件人邮箱地址" />
        </Form.Item>

        <Form.Item label="发件人名称" name="sender_name">
          <Input placeholder="请输入发件人显示名" />
        </Form.Item>

        <Form.Item label="标题前缀" name="sender_title_prefix">
          <Input placeholder="请输入发送邮件标题前缀" />
        </Form.Item>

        <Form.Item
          label="服务器主机"
          name="host"
          rules={[{ required: true, message: '请输入服务器主机名!' }]}
        >
          <Input placeholder="请输入服务器主机名" />
        </Form.Item>

        <Form.Item
          label="服务器端口"
          name="port"
          rules={[{ required: true, message: '请输入服务器端口!' }]}
        >
          <InputNumber placeholder="请输入服务器端口" style={{ width: '100%' }} />
        </Form.Item>

        <Form.Item
          label="用户名"
          name="username"
          rules={[{ required: true, message: '请输入用户名!' }]}
        >
          <Input placeholder="请输入用户名" />
        </Form.Item>

        <Form.Item
          label="密码"
          name="password"
          rules={[{ required: true, message: '请输入密码!' }]}
        >
          <Input.Password placeholder="请输入密码" />
        </Form.Item>

        <Form.Item
          label="是否SSL"
          name="is_ssl"
          valuePropName="checked"
          initialValue={0}
          getValueProps={(value) => ({ checked: value === 1 ? true : false })}
          getValueFromEvent={(value) => {
            return value ? 1 : 0
          }}
        >
          <Switch checkedChildren="是" unCheckedChildren="否" />
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

export default EmailFormUI
