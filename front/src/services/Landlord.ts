import { LandlordDetailResp, LandlordEditResp, LandlordListResp } from '../types/landlordType'
import httpRequest from './http'
import Base from './Base'

const landlordUrl = {
  save: '/system/landlord/save',
  edit: '/system/landlord/edit',
  modify: '/system/landlord/modify',
  delete: '/system/landlord/delete',
  list: '/system/landlord/list',
  detail: '/system/landlord/detail'
}

/**
 * Landlord 业主服务
 */
class Landlord extends Base {
  public constructor() {
    super()
  }

  /**
   * saveLandlord 添加保存业主
   */
  public saveLandlord(landlordInfo: {}): Promise<any> {
    const landlordSaveUrl = this.getProxyUrl(landlordUrl.save)
    return httpRequest.post<any>(landlordSaveUrl, {}, landlordInfo)
  }

  /**
   * getEditLandlordInfo 获取编辑业主信息
   */
  public getEditLandlordInfo(landlord_id: number): Promise<any> {
    const landlordEditUrl = this.getProxyUrl(landlordUrl.edit)
    return httpRequest.get<LandlordEditResp>(landlordEditUrl, {
      landlord_id: landlord_id
    })
  }

  /**
   * modifyLandlord 修改保存业主
   */
  public modifyLandlord(landlordEditInfo: {}): Promise<any> {
    const landlordModifyUrl = this.getProxyUrl(landlordUrl.modify)
    return httpRequest.post<any>(landlordModifyUrl, {}, landlordEditInfo)
  }

  /**
   * landlordList 业主列表
   */
  public landlordList(
    pageSize: number | undefined,
    pageNum: number | undefined,
    keywords: {}
  ): Promise<any> {
    const landlordListUrl = this.getProxyUrl(landlordUrl.list)
    return httpRequest.get<LandlordListResp>(landlordListUrl, {
      page_size: pageSize,
      page_num: pageNum,
      keywords: JSON.stringify(keywords)
    })
  }

  /**
   * getAllLandlords 获取所有的业主
   */
  public getAllLandlords(): Promise<any> {
    const landlordAllListUrl = this.getProxyUrl(landlordUrl.list)
    return httpRequest.get<LandlordListResp>(landlordAllListUrl, {
      is_all: 1
    })
  }

  /**
   * deleteLandlord 删除业主
   */
  public deleteLandlord(landlordId: number): Promise<any> {
    let landlordDeleteUrl = this.getProxyUrl(landlordUrl.delete)
    return httpRequest.post<any>(
      landlordDeleteUrl,
      {},
      {
        landlord_id: landlordId
      }
    )
  }

  /**
   * getLandlordDetail 获取房产详情
   */
  public getLandlordDetail(landlordId: number): Promise<any> {
    let landlordDetailUrl = this.getProxyUrl(landlordUrl.detail)
    return httpRequest.get<LandlordDetailResp>(landlordDetailUrl, { landlord_id: landlordId })
  }
}

export const LandlordService = new Landlord()
