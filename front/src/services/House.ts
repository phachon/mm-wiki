import { HouseDetailResp, HouseEditResp, HouseListResp } from '../types/houseType'
import httpRequest from './http'
import Base from './Base'

const houseUrl = {
  save: '/system/house/save',
  edit: '/system/house/edit',
  modify: '/system/house/modify',
  delete: '/system/house/delete',
  list: '/system/house/list',
  detail: '/system/house/detail'
}

/**
 * House 房产服务
 */
class House extends Base {
  public constructor() {
    super()
  }

  /**
   * saveHouse 添加保存房产
   */
  public saveHouse(houseInfo: {}): Promise<any> {
    const houseSaveUrl = this.getProxyUrl(houseUrl.save)
    return httpRequest.post<any>(houseSaveUrl, {}, houseInfo)
  }

  /**
   * getEditHouseInfo 获取编辑房产信息
   */
  public getEditHouseInfo(house_id: number): Promise<any> {
    const houseEditUrl = this.getProxyUrl(houseUrl.edit)
    return httpRequest.get<HouseEditResp>(houseEditUrl, {
      house_id: house_id
    })
  }

  /**
   * modifyHouse 修改保存房产
   */
  public modifyHouse(houseEditInfo: {}): Promise<any> {
    const houseModifyUrl = this.getProxyUrl(houseUrl.modify)
    return httpRequest.post<any>(houseModifyUrl, {}, houseEditInfo)
  }

  /**
   * houseList 房产列表
   */
  public houseList(
    pageSize: number | undefined,
    pageNum: number | undefined,
    keywords: {}
  ): Promise<any> {
    const houseListUrl = this.getProxyUrl(houseUrl.list)
    return httpRequest.get<HouseListResp>(houseListUrl, {
      page_size: pageSize,
      page_num: pageNum,
      keywords: JSON.stringify(keywords)
    })
  }

  /**
   * getAllHouse 获取所有的房产
   */
  public getAllHouse(): Promise<any> {
    const houseListUrl = this.getProxyUrl(houseUrl.list)
    return httpRequest.get<HouseListResp>(houseListUrl, {
      is_all: 1
    })
  }

  /**
   * deleteHouse 删除房产
   */
  public deleteHouse(houseId: number): Promise<any> {
    let houseDeleteUrl = this.getProxyUrl(houseUrl.delete)
    return httpRequest.post<any>(
      houseDeleteUrl,
      {},
      {
        house_id: houseId
      }
    )
  }

  /**
   * getHouseDetail 获取房产详情
   */
  public getHouseDetail(houseId: bigint): Promise<any> {
    let houseDetailUrl = this.getProxyUrl(houseUrl.detail)
    return httpRequest.get<HouseDetailResp>(houseDetailUrl, { house_id: houseId })
  }
}

export const HouseService = new House()
