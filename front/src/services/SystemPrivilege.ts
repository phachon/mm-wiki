import httpRequest from './http'
import Base from './Base'
import {
  PrivilegeAddResp,
  PrivilegeEditResp,
  PrivilegeInfoType,
  PrivilegeListResp
} from '../types/privilegeType'

const systemPrivilegeUrl = {
  privilegeAdd: '/system/privilege/add',
  privilegeSave: '/system/privilege/save',
  privilegeList: '/system/privilege/list',
  privilegeEdit: '/system/privilege/edit',
  privilegeModify: '/system/privilege/modify',
  privilegeDelete: '/system/privilege/delete'
}

/**
 * SystemPrivilege 系统 - 权限服务
 */
class SystemPrivilege extends Base {
  public constructor() {
    super()
  }

  /**
   * getAddPrivilegeInfo 获取添加权限信息
   * @returns
   */
  public getAddPrivilegeInfo(): Promise<PrivilegeAddResp> {
    let privilegeAddUrl = this.getProxyUrl(systemPrivilegeUrl.privilegeAdd)
    return httpRequest.get<PrivilegeAddResp>(privilegeAddUrl)
  }

  /**
   * savePrivilege 保存权限
   * @param privilegeInfo 权限信息
   * @returns
   */
  public savePrivilege(privilegeInfo: PrivilegeInfoType): Promise<any> {
    let privilegeSaveUrl = this.getProxyUrl(systemPrivilegeUrl.privilegeSave)
    return httpRequest.post<any>(privilegeSaveUrl, {}, privilegeInfo)
  }

  /**
   * 获取编辑权限信息
   * @returns
   */
  public getEditPrivilegeInfo(privilegeId: bigint): Promise<PrivilegeEditResp> {
    let privilegeEditUrl = this.getProxyUrl(systemPrivilegeUrl.privilegeEdit)
    return httpRequest.get<PrivilegeEditResp>(privilegeEditUrl, {
      privilege_id: privilegeId
    })
  }

  /**
   * modifyPrivilege 权限修改保存
   * @param privilegeInfo 权限信息
   * @returns
   */
  public modifyPrivilege(privilegeInfo: PrivilegeInfoType): Promise<any> {
    let privilegeModifyUrl = this.getProxyUrl(systemPrivilegeUrl.privilegeModify)
    return httpRequest.post<any>(privilegeModifyUrl, {}, privilegeInfo)
  }

  /**
   * deletePrivilege 删除权限
   * @param privilegeInfo 权限信息
   * @returns
   */
  public deletePrivilege(privilegeInfo: PrivilegeInfoType): Promise<any> {
    let privilegeDeleteUrl = this.getProxyUrl(systemPrivilegeUrl.privilegeDelete)
    return httpRequest.post<any>(privilegeDeleteUrl, {}, privilegeInfo)
  }

  /**
   * privilegeList 权限列表
   */
  public privilegeList(): Promise<PrivilegeListResp> {
    let privilegeListUrl = this.getProxyUrl(systemPrivilegeUrl.privilegeList)
    return httpRequest.get<PrivilegeListResp>(privilegeListUrl, {})
  }
}

export const SystemPrivilegeService = new SystemPrivilege()
