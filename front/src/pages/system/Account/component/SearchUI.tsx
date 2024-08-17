import { Button, Form, Input, Select } from 'antd'
import { AccountSearchType } from '@/types/accountType'
import { AccountStatusSelectOptions } from './ToolsUI'

interface AccountSearchUIProps {
  onSearchChange: (values: AccountSearchType) => void
  onSearchReset: () => void
}

/**
 * 账号搜索 UI 组件
 * @param props 账号搜索 UI 组件依赖 props
 * @returns UI 组件
 */
const AccountSearchUI = (props: AccountSearchUIProps) => {
  const [form] = Form.useForm()

  const statusOptions = AccountStatusSelectOptions()
  statusOptions.unshift({ label: '全部', value: '' })

  return (
    <div className="panel-body">
      <div className="search-container">
        <Form
          layout={'inline'}
          style={{ justifyContent: 'end' }}
          onFinish={props.onSearchChange}
          form={form}
        >
          <Form.Item name="status" label="状态" initialValue={''}>
            <Select options={statusOptions} />
          </Form.Item>

          <Form.Item name="account_name" label="账号名" style={{ width: 260 }} initialValue={''}>
            <Input placeholder="请输入账号名" />
          </Form.Item>

          <Form.Item name="given_name" label="昵称" style={{ width: 260 }} initialValue={''}>
            <Input placeholder="请输入昵称" />
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

export default AccountSearchUI
