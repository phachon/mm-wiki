import { Button, Form, Input, Modal, Radio, Select } from 'antd'
import { DocTypeRadioOptions } from './ToolsUI'
import { DocType } from '@/types/docType'
import { useEffect } from 'react'

type DocAddUIProps = {
  onSaveSubmit?: (values: any) => void
  parentDoc?: {
    parent_id?: number
    parent_name?: string
  }
}

const DocAddUI = (props: DocAddUIProps) => {
  const [form] = Form.useForm()
  const docTypeOptions = DocTypeRadioOptions()

  useEffect(() => {
    if (props.parentDoc) {
      form.setFieldValue('parent_name', props.parentDoc.parent_name)
      form.setFieldValue('parent_id', props.parentDoc.parent_id)
    }
  }, [props.parentDoc])

  return (
    <div className="panel-body">
      <Form
        name="add-doc-form"
        form={form}
        onFinish={(values) => {
          if (props.onSaveSubmit) {
            props.onSaveSubmit(values)
          }
          // 延迟清空表单
          setTimeout(() => {
            form.resetFields()
          }, 2000)
        }}
      >
        <Form.Item name="parent_id" hidden>
          <Input />
        </Form.Item>
        <Form.Item
          label="父级目录"
          name="parent_name"
          rules={[{ required: true, message: '请选择父级目录!' }]}
        >
          <Input placeholder="请选择父级目录" disabled />
        </Form.Item>
        <Form.Item
          label="文档类型"
          name="doc_type"
          rules={[{ required: true, message: '请选择文档类型!' }]}
          initialValue={DocType.DOC}
        >
          <Radio.Group options={docTypeOptions} />
        </Form.Item>
        <Form.Item
          label="文档名称"
          name="name"
          rules={[{ required: true, message: '请输入文档名称!' }]}
        >
          <Input placeholder="请输入文档名称" />
        </Form.Item>
        <Form.Item label="" wrapperCol={{ offset: 4 }}>
          <Button type="primary" htmlType="submit" style={{ marginLeft: 8 }}>
            保存
          </Button>
        </Form.Item>
      </Form>
    </div>
  )
}

export default DocAddUI
