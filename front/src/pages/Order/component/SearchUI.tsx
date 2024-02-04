import { Button, Form, Input, Row, Select, Col, theme } from 'antd'
import { OrderSearchType } from '@/types/orderType'
import {
  OrderPayTypeSelectOptions,
  OrderStatusSelectOptions,
  OrderTypeSelectOptions
} from './ToolsUI'

interface OrderSearchUIProps {
  /**
   * 搜索改变方法
   * @param values 搜索结果
   */
  searchChangeCallback: (values: OrderSearchType) => void

  /**
   * 搜索重置操作方法
   */
  searchResetCallback: () => void
}

/**
 * 订单搜索 UI 组件
 * @param props 订单搜索 UI 组件依赖 props
 * @returns UI 组件
 */
const OrderSearchUI = (props: OrderSearchUIProps) => {
  const [form] = Form.useForm()

  const orderPayTypeOptions = OrderPayTypeSelectOptions()
  orderPayTypeOptions.unshift({ label: '全部', value: -1 })

  const orderTypeOptions = OrderTypeSelectOptions()
  orderTypeOptions.unshift({ label: '全部', value: -1 })

  const orderStatusOptions = OrderStatusSelectOptions()
  orderStatusOptions.unshift({ label: '全部', value: -1 })

  return (
    <div className="panel-body">
      <div className="search-container">
        <Form onFinish={props.searchChangeCallback} form={form}>
          <Row gutter={24}>
            <Col span={8}>
              <Form.Item name="order_id" label="订单 ID" initialValue={''} className="search-form">
                <Input placeholder="请输入订单ID" />
              </Form.Item>
            </Col>
            <Col span={8}>
              <Form.Item name="house_id" label="房产 ID" initialValue={''} className="search-form">
                <Input placeholder="请输入房产ID" />
              </Form.Item>
            </Col>
            <Col span={6}>
              <Form.Item name="hirer_id" label="租客 ID" initialValue={''} className="search-form">
                <Input placeholder="请输入租客ID" />
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
              <Form.Item name="pay_type" label="支付类型" initialValue={-1} className="search-form">
                <Select options={orderPayTypeOptions} />
              </Form.Item>
            </Col>
            <Col span={8}>
              <Form.Item
                name="order_type"
                label="订单类型"
                initialValue={-1}
                className="search-form"
              >
                <Select options={orderTypeOptions} />
              </Form.Item>
            </Col>
            <Col span={6}>
              <Form.Item
                name="order_status"
                label="订单状态"
                initialValue={-1}
                className="search-form"
              >
                <Select options={orderStatusOptions} />
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

export default OrderSearchUI
