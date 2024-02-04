import { Button, Form, Input, Row, Select, Col } from 'antd'
import { RoomSearchType } from '@/types/roomType'
import {
  RoomAllowLeaseSelectOptions,
  RoomDirectionTypeSelectOptions,
  RoomLeaseStatusSelectOptions,
  RoomRoomTypeSelectOptions
} from './ToolsUI'

interface RoomSearchUIProps {
  /**
   * 搜索改变方法
   * @param values 搜索结果
   */
  searchChangeCallback: (values: RoomSearchType) => void

  /**
   * 搜索重置操作方法
   */
  searchResetCallback: () => void
}

/**
 * 房间搜索 UI 组件
 * @param props 房间搜索 UI 组件依赖 props
 * @returns UI 组件
 */
const RoomSearchUI = (props: RoomSearchUIProps) => {
  const [form] = Form.useForm()

  const leaseStatusOptions = RoomLeaseStatusSelectOptions()
  leaseStatusOptions.unshift({ label: '全部', value: -1 })

  const allowLeaseStatusOption = RoomAllowLeaseSelectOptions()
  allowLeaseStatusOption.unshift({ label: '全部', value: -1 })

  const directionTypeOptions = RoomDirectionTypeSelectOptions()
  directionTypeOptions.unshift({ label: '全部', value: -1 })

  const roomTypeOptions = RoomRoomTypeSelectOptions()
  roomTypeOptions.unshift({ label: '全部', value: -1 })

  return (
    <div className="panel-body">
      <div className="search-container">
        <Form style={{ justifyContent: 'end' }} onFinish={props.searchChangeCallback} form={form}>
          <Row gutter={24}>
            <Col span={8}>
              <Form.Item
                name="lease_status"
                label="出租状态"
                initialValue={-1}
                className="search-form"
              >
                <Select options={leaseStatusOptions} />
              </Form.Item>
            </Col>
            <Col span={8}>
              <Form.Item
                name="allow_lease"
                label="允许出租"
                initialValue={-1}
                className="search-form"
              >
                <Select options={allowLeaseStatusOption} />
              </Form.Item>
            </Col>
            <Col span={6}>
              <Form.Item
                name="direction_type"
                label="房间朝向"
                initialValue={-1}
                className="search-form"
              >
                <Select options={directionTypeOptions} />
              </Form.Item>
            </Col>
            <Col span={2} style={{ textAlign: 'center' }}>
              <Form.Item className="search-form">
                <Button type="default" htmlType="reset" onClick={props.searchResetCallback}>
                  重置
                </Button>
              </Form.Item>
            </Col>
          </Row>
          <Row gutter={24} style={{ marginTop: 16 }}>
            <Col span={8}>
              <Form.Item
                name="room_type"
                label="房间类型"
                initialValue={-1}
                className="search-form"
              >
                <Select options={roomTypeOptions} />
              </Form.Item>
            </Col>
            <Col span={8}>
              <Form.Item name="house_id" label="房间 ID" initialValue={''} className="search-form">
                <Input placeholder="请输入房间 ID" />
              </Form.Item>
            </Col>
            <Col span={6}>
              <Form.Item name="name" label="房间名称" initialValue={''} className="search-form">
                <Input placeholder="请输入房间名称" />
              </Form.Item>
            </Col>
            <Col span={2} style={{ textAlign: 'center' }}>
              <Form.Item className="search-form">
                <Button type="primary" htmlType="submit">
                  查询
                </Button>
              </Form.Item>
            </Col>
          </Row>
        </Form>
      </div>
    </div>
  )
}

export default RoomSearchUI
