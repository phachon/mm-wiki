import { HirerDetailResp, HirerEditResp, HirerListResp } from '../types/hirerType'
import httpRequest from './http'
import Base from './Base'

const hirerUrl = {
  save: '/system/hirer/save',
  edit: '/system/hirer/edit',
  modify: '/system/hirer/modify',
  delete: '/system/hirer/delete',
  list: '/system/hirer/list',
  detail: '/system/hirer/detail'
}

/**
 * Hirer 租客服务
 */
class Hirer extends Base {
  public constructor() {
    super()
  }

  /**
   * saveHirer 添加保存租客
   */
  public saveHirer(hirerInfo: {}): Promise<any> {
    const hirerSaveUrl = this.getProxyUrl(hirerUrl.save)
    return httpRequest.post<any>(hirerSaveUrl, {}, hirerInfo)
  }

  /**
   * getEditHirerInfo 获取编辑租客信息
   */
  public getEditHirerInfo(hirer_id: number): Promise<any> {
    const hirerEditUrl = this.getProxyUrl(hirerUrl.edit)
    return httpRequest.get<HirerEditResp>(hirerEditUrl, {
      hirer_id: hirer_id
    })
  }

  /**
   * modifyHirer 修改保存租客
   */
  public modifyHirer(hirerEditInfo: {}): Promise<any> {
    const hirerModifyUrl = this.getProxyUrl(hirerUrl.modify)
    return httpRequest.post<any>(hirerModifyUrl, {}, hirerEditInfo)
  }

  /**
   * hirerList 租客列表
   */
  public hirerList(
    pageSize: number | undefined,
    pageNum: number | undefined,
    keywords: {}
  ): Promise<any> {
    const hirerListUrl = this.getProxyUrl(hirerUrl.list)
    return httpRequest.get<HirerListResp>(hirerListUrl, {
      page_size: pageSize,
      page_num: pageNum,
      keywords: JSON.stringify(keywords)
    })
  }

  /**
   * getAllHirers 获取所有的租客
   */
  public getAllHirers(): Promise<any> {
    const hirerAllListUrl = this.getProxyUrl(hirerUrl.list)
    return httpRequest.get<HirerListResp>(hirerAllListUrl, {
      is_all: 1
    })
  }

  /**
   * deleteHirer 删除租客
   */
  public deleteHirer(hirerId: number): Promise<any> {
    let hirerDeleteUrl = this.getProxyUrl(hirerUrl.delete)
    return httpRequest.post<any>(
      hirerDeleteUrl,
      {},
      {
        hirer_id: hirerId
      }
    )
  }

  /**
   * getHirerDetail 获取房产详情
   */
  public getHirerDetail(hirerId: number): Promise<any> {
    let hirerDetailUrl = this.getProxyUrl(hirerUrl.detail)
    return httpRequest.get<HirerDetailResp>(hirerDetailUrl, { hirer_id: hirerId })
  }
}

export const HirerService = new Hirer()
