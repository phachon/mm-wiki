import { Descriptions, Tag } from 'antd'
import { OrderInfoType } from '@/types/orderType'
import { OrderPayTypeText, OrderStatusText, OrderTypeTag } from './ToolsUI'

interface OrderDetailUIProps {
  /**
   * 订单信息
   */
  orderInfo?: OrderInfoType
}

/**
 * 订单详情 UI 组件
 */
const OrderDetailUI = (props: OrderDetailUIProps) => {
  const orderInfo = props.orderInfo
  return (
    <div key={orderInfo?.order_id.toString()}>
      <Descriptions bordered size="small" column={12}>
        <Descriptions.Item label="订单ID" span={12} key="order_id">
          {orderInfo?.order_id.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="房产ID" span={12} key="house_id">
          {orderInfo?.house_id.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="房间ID" span={12} key="room_id">
          {orderInfo?.room_id.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="订单类型" span={12} key="order_type">
          {OrderTypeTag(orderInfo?.order_type)}
        </Descriptions.Item>
        <Descriptions.Item label="租客ID" span={12} key="hirer_id">
          {orderInfo?.hirer_id}
        </Descriptions.Item>
        <Descriptions.Item label="开始时间" span={12} key="start_time">
          {orderInfo?.start_time}
        </Descriptions.Item>
        <Descriptions.Item label="结束时间" span={12} key="end_time">
          {orderInfo?.end_time}
        </Descriptions.Item>
        <Descriptions.Item label="月租金" span={12} key="unit_rent">
          {orderInfo?.unit_rent}¥/月
        </Descriptions.Item>
        <Descriptions.Item label="支付类型" span={12} key="pay_type">
          {OrderPayTypeText(orderInfo?.pay_type)}
        </Descriptions.Item>
        <Descriptions.Item label="操作账号" span={12} key="account_id">
          {orderInfo?.account_name}（{orderInfo?.account_id}）
        </Descriptions.Item>
        <Descriptions.Item label="订单状态" span={12} key="status">
          {OrderStatusText(orderInfo?.status)}
        </Descriptions.Item>
        <Descriptions.Item label="创建时间" span={12} key="create_time">
          {orderInfo?.create_time.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="修改时间" span={12} key="update_time">
          {orderInfo?.update_time.toString()}
        </Descriptions.Item>
      </Descriptions>
    </div>
  )
}

export default OrderDetailUI
