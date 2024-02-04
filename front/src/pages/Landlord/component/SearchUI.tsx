import { Button, Form, Input, Select } from 'antd'
import { LandlordSearchType } from '@/types/landlordType'

interface LandlordSearchUIProps {
  /**
   * 业主列表搜索方法
   * @param values 业主搜索关键词
   */
  onLandlordSearchChange: (values: LandlordSearchType) => void

  /**
   * 业主列表搜索重置方法
   */
  onLandlordSearchReset: () => void
}

/**
 * 业主搜索 UI 组件
 * @param props 业主搜索 UI 组件依赖 props
 * @returns UI 组件
 */
const LandlordSearchUI = (props: LandlordSearchUIProps) => {
  const [form] = Form.useForm()

  return (
    <div className="panel-body">
      <div className="search-container">
        <Form
          layout={'inline'}
          style={{ justifyContent: 'end' }}
          onFinish={props.onLandlordSearchChange}
          form={form}
        >
          <Form.Item name="given_name" label="姓名" style={{ width: 300 }} initialValue={''}>
            <Input placeholder="请输入业主姓名" />
          </Form.Item>

          <Form.Item>
            <Button type="default" htmlType="reset" onClick={props.onLandlordSearchReset}>
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

export default LandlordSearchUI
