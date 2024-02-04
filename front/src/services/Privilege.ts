import httpRequest from './http'
import Base from './Base'
import {
  PrivilegeAddResp,
  PrivilegeEditResp,
  PrivilegeInfoType,
  PrivilegeListResp
} from '../types/privilegeType'

const privilegeUrl = {
  privilegeAdd: '/system/privilege/add',
  privilegeSave: '/system/privilege/save',
  privilegeList: '/system/privilege/list',
  privilegeEdit: '/system/privilege/edit',
  privilegeModify: '/system/privilege/modify',
  privilegeDelete: '/system/privilege/delete'
}

/**
 * Privilege 权限服务
 */
class Privilege extends Base {
  public constructor() {
    super()
  }

  /**
   * getAddPrivilegeInfo 获取添加权限信息
   * @returns
   */
  public getAddPrivilegeInfo(): Promise<PrivilegeAddResp> {
    let privilegeAddUrl = this.getProxyUrl(privilegeUrl.privilegeAdd)
    return httpRequest.get<PrivilegeAddResp>(privilegeAddUrl)
  }

  /**
   * savePrivilege 保存权限
   * @param privilegeInfo 权限信息
   * @returns
   */
  public savePrivilege(privilegeInfo: PrivilegeInfoType): Promise<any> {
    let privilegeSaveUrl = this.getProxyUrl(privilegeUrl.privilegeSave)
    return httpRequest.post<any>(privilegeSaveUrl, {}, privilegeInfo)
  }

  /**
   * 获取编辑权限信息
   * @returns
   */
  public getEditPrivilegeInfo(privilegeId: bigint): Promise<PrivilegeEditResp> {
    let privilegeEditUrl = this.getProxyUrl(privilegeUrl.privilegeEdit)
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
    let privilegeModifyUrl = this.getProxyUrl(privilegeUrl.privilegeModify)
    return httpRequest.post<any>(privilegeModifyUrl, {}, privilegeInfo)
  }

  /**
   * deletePrivilege 删除权限
   * @param privilegeInfo 权限信息
   * @returns
   */
  public deletePrivilege(privilegeInfo: PrivilegeInfoType): Promise<any> {
    let privilegeDeleteUrl = this.getProxyUrl(privilegeUrl.privilegeDelete)
    return httpRequest.post<any>(privilegeDeleteUrl, {}, privilegeInfo)
  }

  /**
   * privilegeList 权限列表
   */
  public privilegeList(): Promise<PrivilegeListResp> {
    let privilegeListUrl = this.getProxyUrl(privilegeUrl.privilegeList)
    return httpRequest.get<PrivilegeListResp>(privilegeListUrl, {})
  }
}

export const PrivilegeService = new Privilege()
