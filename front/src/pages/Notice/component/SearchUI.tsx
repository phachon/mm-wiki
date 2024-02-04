import { Button, Form, Input } from 'antd'

interface NoticeSearchUIProps {
  onSearchChange: (values: any) => void
  onSearchReset: () => void
}

/**
 * 公告搜索 UI 组件
 * @param props
 * @returns
 */
const NoticeSearchUI = (props: NoticeSearchUIProps) => {
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
          <Form.Item name="content" label="公告内容" style={{ width: 300 }}>
            <Input placeholder="请输入公告内容关键字" />
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

export default NoticeSearchUI
