import httpRequest from './http'
import Base from './Base'
import { DocAddSaveReq, DocAddSaveResp, DocInfoResp, DocContentSaveReq } from '@/types/docType'
import {
  DocContentSaveResp,
  DocContentHistortyResp,
  DocContentVersionResp
} from '@/types/contentType'
import { CollectionType } from '@/types/collectionType'

const homeIndexUrl = {
  mySpaces: '/home/my_spaces',
  collectionSpaces: '/home/collection_spaces',
  collectionDocs: '/home/collection_docs'
}

/**
 * HomeIndex 主页服务
 */
class HomeIndex extends Base {
  public constructor() {
    super()
  }

  /**
   * 获取我的空间
   */
  public getMySpaces(): Promise<any> {
    return httpRequest.get<any>(this.getProxyUrl(homeIndexUrl.mySpaces), {})
  }

  /**
   * 获取收藏的空间
   */
  public getCollectionSpaces(): Promise<any> {
    return httpRequest.get<any>(this.getProxyUrl(homeIndexUrl.collectionSpaces), {})
  }

  /**
   * 获取收藏的文档
   */
  public getCollectionDocs(): Promise<any> {
    return httpRequest.get<any>(this.getProxyUrl(homeIndexUrl.collectionDocs), {})
  }
}

export const HomeIndexService = new HomeIndex()
