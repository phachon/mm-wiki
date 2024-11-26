import httpRequest from './http'
import Base from './Base'
import { SpaceDocsResp, SpaceListResp } from '@/types/spaceType'

const spaceSpaceUrl = {
  allSpaces: '/space/space/all',
  spaceDocs: '/space/space/docs',
  collect: '/space/collect',
  uncollect: '/space/uncollect'
}

/**
 * SpaceSpace 空间 - 空间服务
 */
class SpaceSpace extends Base {
  public constructor() {
    super()
  }

  /**
   * getSpaces 获取空间列表
   * @param pageSize 每一页条数
   * @param pageNum 页数
   * @param keywords 搜索值
   */
  public getSpaces(pageSize?: number, pageNum?: number, keywords?: {}): Promise<SpaceListResp> {
    const spacesUrl = this.getProxyUrl(spaceSpaceUrl.allSpaces)
    return httpRequest.get<SpaceListResp>(spacesUrl, {
      page_size: pageSize,
      page_num: pageNum,
      keywords: keywords ? JSON.stringify(keywords) : ''
    })
  }

  /**
   * getSpaceDocs 获取空间下文档
   * @param spaceKey 空间Key
   */
  public getSpaceDocs(spaceKey: string): Promise<SpaceDocsResp> {
    const spaceDocsUrl = this.getProxyUrl(spaceSpaceUrl.spaceDocs)
    return httpRequest.get<SpaceDocsResp>(spaceDocsUrl, {
      space_key: spaceKey
    })
  }

  /**
   * collectSpace 收藏空间
   * @param spaceKey 空间标识
   */
  public collectSpace(spaceKey: string): Promise<any> {
    const collectUrl = this.getProxyUrl(spaceSpaceUrl.collect)
    return httpRequest.post<any>(collectUrl, {}, { space_key: spaceKey })
  }

  /**
   * uncollectSpace 取消收藏空间
   * @param spaceKey 空间标识
   */
  public uncollectSpace(spaceKey: string): Promise<any> {
    const uncollectUrl = this.getProxyUrl(spaceSpaceUrl.uncollect)
    return httpRequest.post<any>(uncollectUrl, {}, { space_key: spaceKey })
  }
}

export const SpaceSpaceService = new SpaceSpace()
