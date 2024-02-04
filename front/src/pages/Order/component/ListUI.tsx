import { OrderInfoType, OrderListItemType } from '@/types/orderType'
import {
  Dropdown,
  MenuProps,
  Modal,
  Popconfirm,
  Space,
  Table,
  TablePaginationConfig,
  message
} from 'antd'
import {
  SmileOutlined,
  StepForwardOutlined,
  FormOutlined,
  SelectOutlined,
  DownOutlined,
  RightCircleOutlined
} from '@ant-design/icons'
import {
  OrderIsEdit,
  OrderPayTypeText,
  OrderStatusActionItems,
  OrderStatusText,
  OrderTypeTag
} from './ToolsUI'
import { HirerInfoType } from '@/types/hirerType'
import { HouseShowWholeAddress } from '@/pages/House/component/ToolsUI'
import { HouseInfoType } from '@/types/houseType'
import { RoomInfoType } from '@/types/roomType'
import { useState } from 'react'
import ActionButton from '@/components/Action/ActionButton'

interface OrderListUIProps {
  /**
   * loading
   */
  listLoading: boolean
  /**
   * 订单列表
   */
  orderList: OrderListItemType[]

  /**
   * 订单翻页
   */
  pagination: TablePaginationConfig

  /**
   * 订单列表回调
   * @param pageConfig 翻页配置
   */
  listChangeCallback: (pageConfig: TablePaginationConfig, filters: any, sorter: any) => void

  /**
   * 订单详情点击
   * @param orderInfo
   */
  onOrderDetailClick: (orderListItem: OrderListItemType) => void

  /**
   * 房产详情点击
   * @param orderInfo
   */
  onHouseDetailClick: (houseInfo: HouseInfoType) => void

  /**
   * 房间详情点击
   * @param roomInfo 房间信息
   */
  onRoomDetailClick: (roomInfo: RoomInfoType) => void

  /**
   * 租客详情点击
   * @param hirerInfo 租客详情
   */
  onHirerDetailClick: (hirerInfo: HirerInfoType) => void

  /**
   * 订单修改点击
   * @param orderInfo 订单详情
   */
  onOrderEditClick: (orderListItem: OrderListItemType) => void
}

const OrderListUI = (props: OrderListUIProps) => {
  return (
    <div className="panel-body">
      <Table
        rowKey={'order_id'}
        bordered={true}
        dataSource={props.orderList}
        loading={props.listLoading}
        pagination={props.pagination}
        onChange={props.listChangeCallback}
        footer={() => ''}
      >
        <Table.Column
          title={'订单ID'}
          dataIndex="order_id"
          width={80}
          key={'order_id'}
          align={'center'}
        />
        <Table.Column
          title={'类型'}
          key={'order_type'}
          dataIndex="order_type"
          width={90}
          align={'center'}
          render={OrderTypeTag}
        />
        <Table.Column
          title={'房产地址'}
          key={'house_address'}
          render={(orderListItem: OrderListItemType) => (
            <a onClick={() => props.onHouseDetailClick(orderListItem.house_info)}>
              <Space>
                {HouseShowWholeAddress(orderListItem.house_info)}
                <SelectOutlined />
              </Space>
            </a>
          )}
        />
        <Table.Column
          title={'房间名'}
          key={'room_name'}
          align={'center'}
          render={(orderListItem: OrderListItemType) =>
            orderListItem.room_info ? (
              <a onClick={() => props.onRoomDetailClick(orderListItem.room_info)}>
                <Space>
                  {orderListItem.room_info?.name}
                  <SelectOutlined />
                </Space>
              </a>
            ) : (
              '——'
            )
          }
        />
        <Table.Column
          title={'租客'}
          key={'hirer_id'}
          width={100}
          align={'center'}
          render={(orderListItem: OrderListItemType) => (
            <a onClick={() => props.onHirerDetailClick(orderListItem.hirer_info)}>
              <Space>
                {orderListItem.hirer_info.name}
                <SelectOutlined />
              </Space>
            </a>
          )}
        />
        <Table.Column
          title={'出租日期'}
          width={230}
          key={'lease_time'}
          align={'center'}
          render={(orderListItem: OrderListItemType) =>
            orderListItem.start_time + '～' + orderListItem.end_time
          }
        />
        <Table.Column
          title={'支付类型'}
          dataIndex={'pay_type'}
          key={'pay_type'}
          width={120}
          align={'center'}
          render={OrderPayTypeText}
        />
        <Table.Column
          title={'月租金'}
          dataIndex={'unit_rent'}
          key={'unit_rent'}
          width={120}
          align={'center'}
          render={(unitRent: number) => unitRent + ' ¥/月'}
        />
        <Table.Column
          title={'状态'}
          dataIndex={'status'}
          key={'status'}
          width={90}
          align={'center'}
          render={OrderStatusText}
        />

        <Table.Column
          title={'操作'}
          width={150}
          key={'action'}
          align={'center'}
          render={(orderListItem: OrderListItemType) => (
            <Space>
              <ActionButton
                text="详情"
                icon={<SelectOutlined />}
                onClick={() => props.onOrderDetailClick(orderListItem)}
                havePermission={true}
              />
              <ActionButton
                text="修改"
                icon={<FormOutlined />}
                onClick={() => props.onOrderEditClick(orderListItem)}
                havePermission={orderListItem.action?.is_edit == 1}
              />
            </Space>
          )}
        />
      </Table>
    </div>
  )
}

export default OrderListUI
