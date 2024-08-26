import { Button, Form, Input, Select } from 'antd'
import { SpaceTypeSelectOptions } from './ToolsUI'

interface SpaceSearchUIProps {
  onSearchChange: (values: any) => void // 搜索操作方法
  onSearchReset: () => void // 搜索重置方法
}

/**
 * 空间搜索 UI 组件
 * @param props
 * @returns
 */
const SpaceSearchUI = (props: SpaceSearchUIProps) => {
  const [form] = Form.useForm()
  const spaceTypeOptions = SpaceTypeSelectOptions()

  return (
    <div className="panel-body">
      <div className="search-container">
        <Form
          layout={'inline'}
          style={{ justifyContent: 'end' }}
          onFinish={props.onSearchChange}
          form={form}
        >
          <Form.Item name="space_type" label="空间类型">
            <Select
              placeholder="请选择空间类型"
              style={{ width: 150 }}
              options={spaceTypeOptions}
              allowClear
              autoClearSearchValue
            />
          </Form.Item>
          <Form.Item name="space_name" label="空间名" style={{ width: 350 }}>
            <Input placeholder="请输入空间名" />
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

export default SpaceSearchUI
