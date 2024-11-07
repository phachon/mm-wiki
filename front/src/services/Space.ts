import httpRequest from './http'
import { SpaceAddResp, SpaceAdminListResp, SpaceEditResp, SpaceListResp } from '../types/spaceType'
import Base from './Base'

const spaceUrl = {
  spaceAdd: '/system/space/add',
  spaceSave: '/system/space/save',
  spaceEdit: '/system/space/edit',
  spaceModify: '/system/space/modify',
  spaceList: '/system/space/list',
  spaceDelete: '/system/space/delete',
  adminList: '/system/space/admin_list',
  adminRemove: '/system/space/admin_remove',
  adminAdd: '/system/space/admin_add',
  // 获取公开空间列表
  spaces: '/space/spaces'
}

/**
 * Space 空间服务
 */
class Space extends Base {
  public constructor() {
    super()
  }

  /**
   * getAddSpaceInfo 获取添加空间信息
   */
  public getAddSpaceInfo(): Promise<SpaceAddResp> {
    const addSpaceUrl = this.getProxyUrl(spaceUrl.spaceAdd)
    return httpRequest.get<SpaceAddResp>(addSpaceUrl, {})
  }

  /**
   * saveSpace 添加空间保存
   * @param spaceInfo 添加空间信息
   */
  public saveSpace(spaceInfo: {}): Promise<any> {
    const saveSpaceUrl = this.getProxyUrl(spaceUrl.spaceSave)
    return httpRequest.post<any>(saveSpaceUrl, {}, spaceInfo)
  }

  /**
   * getEditSpaceInfo 获取编辑空间信息
   * @param spaceId 空间id
   */
  public getEditSpaceInfo(spaceId: number): Promise<SpaceEditResp> {
    const spaceEditUrl = this.getProxyUrl(spaceUrl.spaceEdit)
    return httpRequest.get<SpaceEditResp>(spaceEditUrl, {
      space_id: spaceId
    })
  }

  /**
   * modifySpace 更新空间保存
   * @param editSpaceInfo 修改的信息
   * @returns
   */
  public modifySpace(editSpaceInfo: {}): Promise<any> {
    const spaceModifyUrl = this.getProxyUrl(spaceUrl.spaceModify)
    return httpRequest.post<any>(spaceModifyUrl, {}, editSpaceInfo)
  }

  /**
   * getSpaceList 获取空间列表
   * @param pageSize 每一页条数
   * @param pageNum 页数
   * @param keywords 搜索值
   */
  public getSpaceList(pageSize?: number, pageNum?: number, keywords?: {}): Promise<SpaceListResp> {
    const spaceListUrl = this.getProxyUrl(spaceUrl.spaceList)
    return httpRequest.get<SpaceListResp>(spaceListUrl, {
      page_size: pageSize,
      page_num: pageNum,
      keywords: JSON.stringify(keywords)
    })
  }

  /**
   * getSpaces 获取空间列表
   * @param pageSize 每一页条数
   * @param pageNum 页数
   * @param keywords 搜索值
   */
  public getSpaces(pageSize?: number, pageNum?: number, keywords?: {}): Promise<SpaceListResp> {
    const spacesUrl = this.getProxyUrl(spaceUrl.spaces)
    return httpRequest.get<SpaceListResp>(spacesUrl, {
      page_size: pageSize,
      page_num: pageNum,
      keywords: keywords ? JSON.stringify(keywords) : ''
    })
  }

  /**
   * getAdminList 获取账号列表
   * @param pageSize 每一页条数
   * @param pageNum 页数
   * @param spaceId 空间ID
   */
  public getAdminList(spaceId?: number): Promise<SpaceAdminListResp> {
    const spaceAccountListUrl = this.getProxyUrl(spaceUrl.adminList)
    return httpRequest.get<SpaceAdminListResp>(spaceAccountListUrl, {
      space_id: spaceId
    })
  }

  /**
   * deleteSpace 删除空间
   * @param spaceId 空间id
   */
  public deleteSpace(spaceId: number): Promise<any> {
    const deleteSpaceUrl = this.getProxyUrl(spaceUrl.spaceDelete)
    return httpRequest.post<any>(deleteSpaceUrl, {}, { space_id: spaceId })
  }

  /**
   * removeSpaceAdmin 移除空间下管理员
   * @param spaceId 空间ID
   */
  public removeSpaceAdmin(spaceId?: number, accountId?: bigint): Promise<any> {
    const adminRemoveUrl = this.getProxyUrl(spaceUrl.adminRemove)
    const removeBody = {
      space_id: spaceId,
      account_id: accountId
    }
    return httpRequest.post<any>(adminRemoveUrl, {}, removeBody)
  }

  // addSpaceAdmin 添加空间管理员
  public addSpaceAdmin(spaceId?: number, accountIds?: bigint[]): Promise<any> {
    const adminAddUrl = this.getProxyUrl(spaceUrl.adminAdd)
    const adminAddBody = {
      space_id: spaceId,
      admin_account_ids: accountIds
    }
    return httpRequest.post<any>(adminAddUrl, {}, adminAddBody)
  }

  /**
   * collectSpace 收藏空间
   * @param spaceKey 空间标识
   */
  public collectSpace(spaceKey: string): Promise<any> {
    const collectUrl = this.getProxyUrl('/space/collect')
    return httpRequest.post<any>(collectUrl, {}, { space_key: spaceKey })
  }

  /**
   * uncollectSpace 取消收藏空间
   * @param spaceKey 空间标识
   */
  public uncollectSpace(spaceKey: string): Promise<any> {
    const uncollectUrl = this.getProxyUrl('/space/uncollect')
    return httpRequest.post<any>(uncollectUrl, {}, { space_key: spaceKey })
  }
}

export const SpaceService = new Space()
