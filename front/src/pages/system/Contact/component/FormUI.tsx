import { Button, Form, Input } from 'antd'
import { EditLayoutForm, LayoutForm } from '@/config/layout'
import { ContactInfoType } from '@/types/contactType'
import { useEffect } from 'react'

interface ContactFormUIProps {
  contactInfo?: ContactInfoType
  onSaveSubmit: (values: any) => void
  formLayout?: {
    labelCol: { span: number }
    wrapperCol: { span: number }
  }
}

const ContactFormUI = (props: ContactFormUIProps) => {
  const [form] = Form.useForm()
  const isEdit = props.contactInfo ? true : false
  let layoutForm = isEdit ? EditLayoutForm : LayoutForm
  if (props.formLayout) {
    layoutForm = props.formLayout
  }

  useEffect(() => {
    if (props.contactInfo) {
      form.setFieldsValue(props.contactInfo)
    }
  }, [props.contactInfo])

  return (
    <div className="panel-body">
      <Form name="contact-form" {...layoutForm} form={form} onFinish={props.onSaveSubmit}>
        {isEdit && (
          <Form.Item label="联系人ID" name="contact_id" rules={[{ required: true }]}>
            <Input disabled placeholder="请输入联系人ID" />
          </Form.Item>
        )}

        <Form.Item
          label="联系人名称"
          name="name"
          rules={[{ required: true, message: '请输入联系人名称!' }]}
        >
          <Input placeholder="请输入联系人名称" />
        </Form.Item>

        <Form.Item label="联系电话" name="mobile">
          <Input placeholder="请输入联系电话" />
        </Form.Item>

        <Form.Item label="邮箱" name="email">
          <Input placeholder="请输入邮箱" />
        </Form.Item>

        <Form.Item label="职位" name="position">
          <Input placeholder="请输入联系人职位" />
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

export default ContactFormUI
