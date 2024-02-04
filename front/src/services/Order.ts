import httpRequest from './http'
import { OrderAddResp, OrderDetailResp, OrderInfoType, OrderListResp } from '../types/orderType'
import Base from './Base'

const orderUrl = {
  orderAdd: '/system/order/add',
  orderSave: '/system/order/save',
  orderEdit: '/system/order/edit',
  orderModify: '/system/order/modify',
  orderList: '/system/order/list',
  orderDetail: '/system/order/detail'
}

/**
 * Order 订单服务
 */
class Order extends Base {
  public constructor() {
    super()
  }

  /**
   * getAddOrderInfo 获取添加订单信息
   */
  public getAddOrderInfo(orderType: number): Promise<OrderAddResp> {
    const addOrderUrl = this.getProxyUrl(orderUrl.orderAdd)
    return httpRequest.get<any>(addOrderUrl, {
      order_type: orderType
    })
  }

  /**
   * saveOrder 添加保存订单
   */
  public saveOrder(orderInfo: OrderInfoType): Promise<any> {
    const orderSaveUrl = this.getProxyUrl(orderUrl.orderSave)
    return httpRequest.post<any>(orderSaveUrl, {}, orderInfo)
  }

  /**
   * orderList 订单列表
   */
  public orderList(
    pageSize: number | undefined,
    pageNum: number | undefined,
    keywords: {}
  ): Promise<any> {
    const orderListUrl = this.getProxyUrl(orderUrl.orderList)
    return httpRequest.get<OrderListResp>(orderListUrl, {
      page_size: pageSize,
      page_num: pageNum,
      keywords: JSON.stringify(keywords)
    })
  }

  /**
   * getOrderDetail 获取订单详情
   */
  public getOrderDetail(orderId: number): Promise<any> {
    let roomDetailUrl = this.getProxyUrl(orderUrl.orderDetail)
    return httpRequest.get<OrderDetailResp>(roomDetailUrl, { order_id: orderId })
  }

  /**
   * orderModify 订单修改保存
   */
  public orderModify(orderInfo: OrderInfoType): Promise<any> {
    const modifyUrl = this.getProxyUrl(orderUrl.orderModify)
    return httpRequest.post<any>(modifyUrl, {}, orderInfo)
  }
}

export const OrderService = new Order()
