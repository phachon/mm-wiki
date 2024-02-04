import { Button, Form, Input, Select } from 'antd'
import { HouseSearchType } from '@/types/houseType'
import {
  HouseDecorationTypeSelectOptions,
  HouseLeaseStatusSelectOptions,
  HouseSizeTypeSelectOptions
} from './ToolsUI'

interface HouseSearchUIProps {
  /**
   * 房产列表搜索方法
   * @param values 房产搜索关键词
   */
  onHouseSearchChange: (values: HouseSearchType) => void

  /**
   * 房产列表搜索重置方法
   */
  onHouseSearchReset: () => void
}

/**
 * 房产搜索 UI 组件
 * @param props 房产搜索 UI 组件依赖 props
 * @returns UI 组件
 */
const HouseSearchUI = (props: HouseSearchUIProps) => {
  const [form] = Form.useForm()

  const sizeTypeOptions = HouseSizeTypeSelectOptions()
  sizeTypeOptions.unshift({ label: '全部', value: -1 })

  const decorationTypeOptions = HouseDecorationTypeSelectOptions()
  decorationTypeOptions.unshift({ label: '全部', value: -1 })

  const leaseStatusOptions = HouseLeaseStatusSelectOptions()
  leaseStatusOptions.unshift({ label: '全部', value: -1 })

  return (
    <div className="panel-body">
      <div className="search-container">
        <Form
          layout={'inline'}
          style={{ justifyContent: 'end' }}
          onFinish={props.onHouseSearchChange}
          form={form}
        >
          <Form.Item name="decoration_type" label="装修类型" initialValue={-1}>
            <Select style={{ width: 100 }} options={decorationTypeOptions} />
          </Form.Item>

          <Form.Item name="size_type" label="户型" initialValue={-1}>
            <Select style={{ width: 100 }} options={sizeTypeOptions} />
          </Form.Item>

          <Form.Item name="lease_status" label="出租状态" initialValue={-1}>
            <Select style={{ width: 100 }} options={leaseStatusOptions} />
          </Form.Item>

          <Form.Item name="address" label="地址" style={{ width: 300 }} initialValue={''}>
            <Input placeholder="请输入房产地址" />
          </Form.Item>

          <Form.Item style={{ justifyContent: 'end' }}>
            <Button type="default" htmlType="reset" onClick={props.onHouseSearchReset}>
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

export default HouseSearchUI
