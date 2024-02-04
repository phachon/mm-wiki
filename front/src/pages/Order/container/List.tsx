import React, { useEffect, useState } from 'react'
import OrderListUI from '../component/ListUI'
import { OrderDetailResp, OrderInfoType, OrderListItemType, OrderListResp } from '@/types/orderType'
import { OrderService } from '@/services/Order'
import { message, Modal, TablePaginationConfig } from 'antd'
import OrderFormUI from '../component/FormUI'
import { initPagination } from '@/types/adminType'
import { HouseDetailResp, HouseInfoType } from '@/types/houseType'
import { HouseService } from '@/services/House'
import HouseDetailUI from '@/pages/House/component/DetailUI'
import OrderSearchUI from '../component/SearchUI'
import { HirerDetailResp, HirerInfoType } from '@/types/hirerType'
import HirerDetailUI from '@/pages/Hirer/component/DetailUI'
import { HirerService } from '@/services/Hirer'
import { RoomInfoType } from '@/types/roomType'
import RoomDetailUI from '@/pages/Room/component/DetailUI'
import OrderDetailUI from '../component/DetailUI'
import OrderEditUI from '../component/EditUI'

let searchValues = {}

const OrderList: React.FC = () => {
  // 订单列表
  const [orderList, setOrderList] = useState<OrderListItemType[]>([])
  const [pagination, setPagination] = useState(initPagination)

  // 租客详情
  const [detailHirerInfo, setDetailHirerInfo] = useState<HirerInfoType>()
  const [detailHirerModalOpen, setDetailHirerModalOpen] = useState<boolean>(false)

  // 房产详情
  const [detailHouseInfo, setDetailHouseInfo] = useState<HouseInfoType>()
  const [detailHouseModalOpen, setDetailHouseModalOpen] = useState<boolean>(false)

  // 房间详情
  const [detailRoomInfo, setDetailRoomInfo] = useState<RoomInfoType>()
  const [detailRoomModalOpen, setDetailRoomModalOpen] = useState<boolean>(false)

  // 订单修改
  const [editModalOpen, setEditModalOpen] = useState<boolean>(false)
  const [editOrderInfo, setEditOrderInfo] = useState<OrderInfoType>()

  // 订单详情
  const [detailOrderInfo, setDetailOrderInfo] = useState<OrderInfoType>()
  const [detailOrderModalOpen, setDetailOrderModalOpen] = useState<boolean>(false)

  useEffect(() => {
    getOrderList(initPagination, {})
  }, [])

  /**
   * 请求订单列表
   * @param pageConfig 翻页信息
   * @param searchKeywords 搜索信息
   */
  const getOrderList = (pageConfig: TablePaginationConfig, searchKeywords: {}) => {
    searchValues = searchKeywords
    OrderService.orderList(pageConfig.pageSize, pageConfig.current, searchKeywords)
      .then((res: OrderListResp) => {
        setOrderList(res.list)
        setPagination({
          ...initPagination,
          current: res.page_info?.page_num,
          pageSize: res.page_info?.page_size,
          total: res.page_info?.total_num
        })
      })
      .catch((e) => {
        console.log('获取房间列表失败 err:', e)
      })
  }
  /**
   * 搜索查询操作
   * @param values 搜索表单数据
   */
  const searchChangeCallback = (values: {}) => {
    getOrderList(initPagination, values)
  }
  /**
   * 搜索重置操作
   */
  const searchResetCallback = () => {
    getOrderList(initPagination, {})
  }
  /**
   * 列表分页请求
   * @param pageConfig
   * @param filters
   * @param sorter
   */
  const listChangeCallback = (pageConfig: TablePaginationConfig, filters: any, sorter: any) => {
    getOrderList(pageConfig, searchValues)
  }
  /**
   * 修改点击操作
   * @param orderInfo 订单信息
   */
  const onOrderEditClick = (orderListItem: OrderListItemType) => {
    OrderService.getOrderDetail(orderListItem.order_id)
      .then((resp: OrderDetailResp) => {
        setEditOrderInfo(resp.order_info)
        setEditModalOpen(true)
      })
      .catch((e) => {
        console.log('订单修改 err:', e)
      })
  }

  /**
   * 订单详情点击操作
   * @param orderInfo 订单信息
   */
  const onOrderDetailClick = (orderListItem: OrderListItemType) => {
    OrderService.getOrderDetail(orderListItem.order_id)
      .then((resp: OrderDetailResp) => {
        setDetailOrderInfo(resp.order_info)
        setDetailOrderModalOpen(true)
      })
      .catch((e) => {
        console.log('详情点击出错err:', e)
      })
  }

  /**
   * 房产详情点击操作
   * @param houseInfo 房产信息
   */
  const onHouseDetailClick = (houseInfo: HouseInfoType) => {
    setDetailHouseInfo(houseInfo)
    setDetailHouseModalOpen(true)
  }

  /**
   * 房间详情
   * @param roomInfo 房间信息
   */
  const onRoomDetailClick = (roomInfo: RoomInfoType) => {
    setDetailRoomInfo(roomInfo)
    setDetailRoomModalOpen(true)
  }

  /**
   * 租客详情点击操作
   * @param hirerInfo 租客信息
   */
  const onHirerDetailClick = (hirerInfo: HirerInfoType) => {
    setDetailHirerInfo(hirerInfo)
    setDetailHirerModalOpen(true)
  }

  /**
   * 修改完成操作
   * @param orderInfo OrderInfoType
   */
  const onEditSave = (orderInfo: OrderInfoType) => {
    OrderService.orderModify(orderInfo)
      .then((resp) => {
        message.success('修改成功', 2, () => {
          setEditModalOpen(false)
          getOrderList(pagination, searchValues)
        })
      })
      .catch((e) => {
        console.log('修改订单err:', e)
      })
  }

  return (
    <div className="panel">
      <OrderSearchUI
        searchChangeCallback={searchChangeCallback}
        searchResetCallback={searchResetCallback}
      />
      <OrderListUI
        listLoading={false}
        orderList={orderList}
        pagination={pagination}
        listChangeCallback={listChangeCallback}
        onOrderEditClick={onOrderEditClick}
        onOrderDetailClick={onOrderDetailClick}
        onHouseDetailClick={onHouseDetailClick}
        onRoomDetailClick={onRoomDetailClick}
        onHirerDetailClick={onHirerDetailClick}
      />
      <Modal
        title="租客详情"
        width={570}
        open={detailHirerModalOpen}
        onCancel={() => {
          setDetailHirerModalOpen(false)
        }}
        footer={null}
      >
        <HirerDetailUI hirerDetail={detailHirerInfo} />
      </Modal>
      <Modal
        title="房产详情"
        width={570}
        open={detailHouseModalOpen}
        onCancel={() => {
          setDetailHouseModalOpen(false)
        }}
        footer={null}
      >
        <HouseDetailUI detailHouseInfo={detailHouseInfo} />
      </Modal>
      <Modal
        title="房间详情"
        width={570}
        open={detailRoomModalOpen}
        onCancel={() => {
          setDetailRoomModalOpen(false)
        }}
        footer={null}
      >
        <RoomDetailUI roomInfo={detailRoomInfo} />
      </Modal>
      <Modal
        title="订单详情"
        width={570}
        open={detailOrderModalOpen}
        onCancel={() => {
          setDetailOrderModalOpen(false)
        }}
        footer={null}
      >
        <OrderDetailUI orderInfo={detailOrderInfo} />
      </Modal>
      <Modal
        title="订单修改"
        width={570}
        open={editModalOpen}
        onCancel={() => {
          setEditModalOpen(false)
        }}
        footer={null}
      >
        <OrderEditUI orderInfo={editOrderInfo} onEditSave={onEditSave} />
      </Modal>
    </div>
  )
}

export default OrderList
