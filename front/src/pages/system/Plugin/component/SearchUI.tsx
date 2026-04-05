import { Button, Form, Input } from 'antd'

interface PluginSearchUIProps {
  onSearchChange: (values: any) => void
  onSearchReset: () => void
}

const PluginSearchUI = (props: PluginSearchUIProps) => {
  const [form] = Form.useForm()
  return (
    <div className="panel-body">
      <div className="search-container">
        <Form
          layout={'inline'}
          style={{ justifyContent: 'end' }}
          onFinish={props.onSearchChange}
          form={form}
        >
          <Form.Item name="name" label="插件名称" style={{ width: 300 }}>
            <Input placeholder="请输入插件名称关键字" />
          </Form.Item>

          <Form.Item>
            <Button type="default" htmlType="reset" onClick={props.onSearchReset}>
              重置
            </Button>
          </Form.Item>

          <Form.Item style={{ margin: 0 }}>
            <Button type="primary" htmlType="submit">
              查询
            </Button>
          </Form.Item>
        </Form>
      </div>
    </div>
  )
}

export default PluginSearchUI
