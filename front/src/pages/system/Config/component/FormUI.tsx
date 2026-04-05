import { useEffect } from 'react'
import { Button, Form, Input, InputNumber, Switch } from 'antd'
import { LayoutForm } from '@/config/layout'

interface ConfigFormUIProps {
  configMap?: { [key: string]: string }
  onSaveSubmit: (values: any) => void
}

/**
 * 系统配置表单 UI 组件
 * @param props 配置表单 UI 组件数据
 * @returns 配置表单 UI 组件
 */
const ConfigFormUI = (props: ConfigFormUIProps) => {
  const [form] = Form.useForm()
  const layoutForm = LayoutForm

  useEffect(() => {
    if (props.configMap) {
      form.setFieldsValue({
        main_title: props.configMap.main_title || '',
        main_description: props.configMap.main_description || '',
        system_name: props.configMap.system_name || '',
        auto_follow_doc_open: props.configMap.auto_follow_doc_open === '1',
        send_email_open: props.configMap.send_email_open === '1',
        sso_open: props.configMap.sso_open === '1',
        fulltext_search_open: props.configMap.fulltext_search_open === '1',
        doc_search_timer: props.configMap.doc_search_timer
          ? Number(props.configMap.doc_search_timer)
          : undefined
      })
    }
  }, [props.configMap])

  const onSubmit = (values: any) => {
    props.onSaveSubmit({
      ...values,
      auto_follow_doc_open: values.auto_follow_doc_open ? '1' : '0',
      send_email_open: values.send_email_open ? '1' : '0',
      sso_open: values.sso_open ? '1' : '0',
      fulltext_search_open: values.fulltext_search_open ? '1' : '0',
      doc_search_timer: values.doc_search_timer ? String(values.doc_search_timer) : '0'
    })
  }

  return (
    <div className="panel-body">
      <Form {...layoutForm} name="config-form" form={form} onFinish={onSubmit}>
        <Form.Item
          label="系统标题"
          name="main_title"
          rules={[{ required: true, message: '请输入系统标题!' }]}
        >
          <Input placeholder="请输入系统标题" />
        </Form.Item>

        <Form.Item
          label="系统描述"
          name="main_description"
          rules={[{ required: true, message: '请输入系统描述!' }]}
        >
          <Input placeholder="请输入系统描述" />
        </Form.Item>

        <Form.Item
          label="系统名称"
          name="system_name"
          rules={[{ required: true, message: '请输入系统名称!' }]}
        >
          <Input placeholder="请输入系统名称" />
        </Form.Item>

        <Form.Item
          label="自动关注文档"
          name="auto_follow_doc_open"
          valuePropName="checked"
        >
          <Switch checkedChildren="开" unCheckedChildren="关" />
        </Form.Item>

        <Form.Item
          label="邮件发送"
          name="send_email_open"
          valuePropName="checked"
        >
          <Switch checkedChildren="开" unCheckedChildren="关" />
        </Form.Item>

        <Form.Item
          label="SSO登录认证"
          name="sso_open"
          valuePropName="checked"
        >
          <Switch checkedChildren="开" unCheckedChildren="关" />
        </Form.Item>

        <Form.Item
          label="全文搜索"
          name="fulltext_search_open"
          valuePropName="checked"
        >
          <Switch checkedChildren="开" unCheckedChildren="关" />
        </Form.Item>

        <Form.Item
          label="文档搜索间隔(秒)"
          name="doc_search_timer"
          rules={[{ required: true, message: '请输入文档搜索间隔!' }]}
        >
          <InputNumber min={1} placeholder="请输入文档搜索间隔" style={{ width: '100%' }} />
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

export default ConfigFormUI
