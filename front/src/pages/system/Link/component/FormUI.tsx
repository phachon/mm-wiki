import { Button, Form, Input, InputNumber } from 'antd'
import { EditLayoutForm, LayoutForm } from '@/config/layout'
import { LinkInfoType } from '@/types/linkType'
import { useEffect } from 'react'

interface LinkFormUIProps {
  linkInfo?: LinkInfoType
  onSaveSubmit: (values: any) => void
  formLayout?: {
    labelCol: { span: number }
    wrapperCol: { span: number }
  }
}

const LinkFormUI = (props: LinkFormUIProps) => {
  const [form] = Form.useForm()
  const isEdit = props.linkInfo ? true : false
  let layoutForm = isEdit ? EditLayoutForm : LayoutForm
  if (props.formLayout) {
    layoutForm = props.formLayout
  }

  useEffect(() => {
    if (props.linkInfo) {
      form.setFieldsValue(props.linkInfo)
    }
  }, [props.linkInfo])

  return (
    <div className="panel-body">
      <Form name="link-form" {...layoutForm} form={form} onFinish={props.onSaveSubmit}>
        {isEdit && (
          <Form.Item label="链接ID" name="link_id" rules={[{ required: true }]}>
            <Input disabled placeholder="请输入链接ID" />
          </Form.Item>
        )}

        <Form.Item
          label="链接名称"
          name="name"
          rules={[{ required: true, message: '请输入链接名称!' }]}
        >
          <Input placeholder="请输入链接名称" />
        </Form.Item>

        <Form.Item
          label="链接地址"
          name="url"
          rules={[{ required: true, message: '请输入链接地址!' }]}
        >
          <Input placeholder="请输入链接地址" />
        </Form.Item>

        <Form.Item label="排序号" name="sequence" initialValue={0}>
          <InputNumber placeholder="越小越靠前" style={{ width: '100%' }} />
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

export default LinkFormUI
