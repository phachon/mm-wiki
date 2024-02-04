import { Button, Form, Input, Select } from 'antd'
import { HirerSearchType } from '@/types/hirerType'

interface HirerSearchUIProps {
  /**
   * 租客列表搜索方法
   * @param values 租客搜索关键词
   */
  onHirerSearchChange: (values: HirerSearchType) => void

  /**
   * 租客列表搜索重置方法
   */
  onHirerSearchReset: () => void
}

/**
 * 租客搜索 UI 组件
 * @param props 租客搜索 UI 组件依赖 props
 * @returns UI 组件
 */
const HirerSearchUI = (props: HirerSearchUIProps) => {
  const [form] = Form.useForm()

  return (
    <div className="panel-body">
      <div className="search-container">
        <Form
          layout={'inline'}
          style={{ justifyContent: 'end' }}
          onFinish={props.onHirerSearchChange}
          form={form}
        >
          <Form.Item name="given_name" label="姓名" style={{ width: 300 }} initialValue={''}>
            <Input placeholder="请输入租客姓名" />
          </Form.Item>

          <Form.Item>
            <Button type="default" htmlType="reset" onClick={props.onHirerSearchReset}>
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

export default HirerSearchUI
