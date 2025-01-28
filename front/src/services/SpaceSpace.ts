import httpRequest from './http'
import Base from './Base'
import { SpaceDocsResp, SpaceListResp, SpacePermissionListResp } from '@/types/spaceType'

const spaceSpaceUrl = {
  allSpaces: '/space/space/all',
  spaceDocs: '/space/space/docs',
  spaceSettingBasicModify: '/space/setting/basic_modify',
  spaceSettingPermissionList: '/space/setting/permission_list'
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
   * 获取空间设置权限列表
   * @returns
   */
  public getSpaceSettingPermissionList(spaceId?: number): Promise<SpacePermissionListResp> {
    const spaceSettingPermissionListUrl = this.getProxyUrl(spaceSpaceUrl.spaceSettingPermissionList)
    return httpRequest.get<SpacePermissionListResp>(spaceSettingPermissionListUrl, {
      space_id: spaceId
    })
  }

  /**
   * modifySpaceBasicSetting 空间更新保存
   * @param spaceInfo 添加空间信息
   */
  public modifySpaceBasicSetting(spaceInfo: {}): Promise<any> {
    const modifySpaceUrl = this.getProxyUrl(spaceSpaceUrl.spaceSettingBasicModify)
    return httpRequest.post<any>(modifySpaceUrl, {}, spaceInfo)
  }
}

export const SpaceSpaceService = new SpaceSpace()
