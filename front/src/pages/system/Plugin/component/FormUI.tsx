import { Button, Form, Input } from 'antd'
import { EditLayoutForm, LayoutForm } from '@/config/layout'
import { PluginInfoType } from '@/types/pluginType'
import { useEffect } from 'react'

interface PluginFormUIProps {
  pluginInfo?: PluginInfoType
  onSaveSubmit: (values: any) => void
  formLayout?: {
    labelCol: { span: number }
    wrapperCol: { span: number }
  }
}

const PluginFormUI = (props: PluginFormUIProps) => {
  const [form] = Form.useForm()
  const isEdit = props.pluginInfo ? true : false
  let layoutForm = isEdit ? EditLayoutForm : LayoutForm
  if (props.formLayout) {
    layoutForm = props.formLayout
  }

  useEffect(() => {
    if (props.pluginInfo) {
      form.setFieldsValue(props.pluginInfo)
    }
  }, [props.pluginInfo])

  return (
    <div className="panel-body">
      <Form name="plugin-form" {...layoutForm} form={form} onFinish={props.onSaveSubmit}>
        {isEdit && (
          <Form.Item label="插件ID" name="plugin_id" rules={[{ required: true }]}>
            <Input disabled placeholder="请输入插件ID" />
          </Form.Item>
        )}

        <Form.Item
          label="插件名称"
          name="name"
          rules={[{ required: true, message: '请输入插件名称!' }]}
        >
          <Input placeholder="请输入插件名称" />
        </Form.Item>

        <Form.Item
          label="插件标识"
          name="key"
          rules={[{ required: true, message: '请输入插件标识!' }]}
        >
          <Input disabled={isEdit} placeholder="请输入插件标识" />
        </Form.Item>

        <Form.Item label="插件描述" name="description">
          <Input.TextArea placeholder="请输入插件描述" />
        </Form.Item>

        <Form.Item label="版本号" name="version">
          <Input placeholder="请输入版本号" />
        </Form.Item>

        <Form.Item label="作者" name="author">
          <Input placeholder="请输入作者" />
        </Form.Item>

        <Form.Item label="配置JSON" name="config_json">
          <Input.TextArea placeholder="请输入配置JSON" />
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

export default PluginFormUI
