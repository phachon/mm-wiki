import { Button, Form, Input } from 'antd'

interface RoleSearchUIProps {
  /**
   * 搜索操作方法
   * @param values
   */
  onSearchChange: (values: any) => void
  /**
   * 搜索重置方法
   */
  onSearchReset: () => void
}

/**
 * 角色搜索 UI 组件
 * @param props
 * @returns
 */
const RoleSearchUI = (props: RoleSearchUIProps) => {
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
          <Form.Item name="role_name" label="角色名" style={{ width: 260 }}>
            <Input placeholder="请输入角色名" />
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

export default RoleSearchUI
