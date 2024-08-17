import { Button, Form, Input, Select } from 'antd'
import { LogLevelSelectOptions } from './ToolsUI'

interface LogSearchUIProps {
  onSearchChange: (values: {}) => void
  onResetChange: () => void
}

/**
 * 日志搜索 UI 组件
 * @param props
 * @returns
 */
const LogSearchUI = (props: LogSearchUIProps) => {
  const [form] = Form.useForm()
  const logLevelOptions = LogLevelSelectOptions()

  return (
    <div className="panel-body">
      <div className="search-container">
        <Form
          layout={'inline'}
          style={{ justifyContent: 'end' }}
          onFinish={props.onSearchChange}
          form={form}
        >
          <Form.Item name="level" label="级别" initialValue={0}>
            <Select style={{ width: 80 }} options={logLevelOptions} />
          </Form.Item>

          <Form.Item
            name="account_id"
            label="账号ID"
            rules={[
              {
                type: 'integer',
                transform: (value) => {
                  return value ? Number(value) : 0
                }
              }
            ]}
          >
            <Input placeholder="请输入账号ID" />
          </Form.Item>

          <Form.Item name="message" label="日志内容" style={{ width: 300 }}>
            <Input placeholder="请输入日志内容" />
          </Form.Item>

          <Form.Item>
            <Button type="default" htmlType="reset" onClick={props.onResetChange}>
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

export default LogSearchUI
