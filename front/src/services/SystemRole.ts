import httpRequest from './http'
import { RoleAddResp, RoleEditResp, RoleListResp, RolePrivilegeEditResp } from '../types/roleType'
import Base from './Base'
import { AccountListResp } from '../types/accountType'

const roleUrl = {
  roleAdd: '/system/role/add',
  roleSave: '/system/role/save',
  roleEdit: '/system/role/edit',
  roleModify: '/system/role/modify',
  roleList: '/system/role/list',
  roleDelete: '/system/role/delete',
  accountList: '/system/role/account_list',
  accountRemove: '/system/role/account_remove',
  privilegeEdit: '/system/role/privilege_edit',
  privilegeModify: '/system/role/privilege_modify'
}

/**
 * SystemRole 系统 - 角色服务
 */
class SystemRole extends Base {
  public constructor() {
    super()
  }

  /**
   * getAddRoleInfo 获取添加角色信息
   */
  public getAddRoleInfo(): Promise<RoleAddResp> {
    const addRoleUrl = this.getProxyUrl(roleUrl.roleAdd)
    return httpRequest.get<any>(addRoleUrl, {})
  }

  /**
   * saveRole 添加角色保存
   * @param roleInfo 添加角色信息
   */
  public saveRole(roleInfo: {}): Promise<any> {
    const saveRoleUrl = this.getProxyUrl(roleUrl.roleSave)
    return httpRequest.post<any>(saveRoleUrl, {}, roleInfo)
  }

  /**
   * getEditRoleInfo 获取编辑角色信息
   * @param roleId 角色id
   */
  public getEditRoleInfo(roleId: number): Promise<RoleEditResp> {
    const roleEditUrl = this.getProxyUrl(roleUrl.roleEdit)
    return httpRequest.get<RoleEditResp>(roleEditUrl, {
      role_id: roleId
    })
  }

  /**
   * modifyRole 更新角色保存
   * @param editRoleInfo 修改的信息
   * @returns
   */
  public modifyRole(editRoleInfo: {}): Promise<any> {
    const roleModifyUrl = this.getProxyUrl(roleUrl.roleModify)
    return httpRequest.post<any>(roleModifyUrl, {}, editRoleInfo)
  }

  /**
   * getRoleList 获取角色列表
   * @param pageSize 每一页条数
   * @param pageNum 页数
   * @param keywords 搜索值
   */
  public getRoleList(pageSize?: number, pageNum?: number, keywords?: {}): Promise<RoleListResp> {
    const roleListUrl = this.getProxyUrl(roleUrl.roleList)
    return httpRequest.get<RoleListResp>(roleListUrl, {
      page_size: pageSize,
      page_num: pageNum,
      keywords: JSON.stringify(keywords)
    })
  }

  /**
   * getAccountList 获取账号列表
   * @param pageSize 每一页条数
   * @param pageNum 页数
   * @param roleId 角色ID
   */
  public getAccountList(
    pageSize?: number,
    pageNum?: number,
    roleId?: number
  ): Promise<AccountListResp> {
    const roleAccountListUrl = this.getProxyUrl(roleUrl.accountList)
    return httpRequest.get<AccountListResp>(roleAccountListUrl, {
      page_size: pageSize,
      page_num: pageNum,
      role_id: roleId
    })
  }

  /**
   * deleteRole 删除角色
   * @param roleId 角色id
   */
  public deleteRole(roleId: number): Promise<any> {
    const deleteRoleUrl = this.getProxyUrl(roleUrl.roleDelete)
    return httpRequest.post<any>(deleteRoleUrl, {
      role_id: roleId
    })
  }

  /**
   * getPrivilegeEdit 获取权限修改信息
   * @param roleId 角色ID
   */
  public getPrivilegeEdit(roleId: number): Promise<RolePrivilegeEditResp> {
    const rolePrivilegeEditUrl = this.getProxyUrl(roleUrl.privilegeEdit)
    return httpRequest.get<RolePrivilegeEditResp>(rolePrivilegeEditUrl, {
      role_id: roleId
    })
  }

  /**
   * modifyRolePrivilege 角色权限保存
   * @param roleId 角色ID
   */
  public modifyRolePrivilege(roleId?: number, privilegeIds?: string[]): Promise<any> {
    const rolePrivilegeModifyUrl = this.getProxyUrl(roleUrl.privilegeModify)
    return httpRequest.post<any>(rolePrivilegeModifyUrl, {
      role_id: roleId,
      privilege_ids: privilegeIds ? privilegeIds.join(',') : '' // eg: "123,456,11,21"
    })
  }

  /**
   * removeRoleAccount 移除角色下账号
   * @param roleId 角色ID
   */
  public removeRoleAccount(roleId?: number, accountId?: bigint): Promise<any> {
    const accountRemoveUrl = this.getProxyUrl(roleUrl.accountRemove)
    return httpRequest.post<any>(accountRemoveUrl, {
      role_id: roleId,
      account_id: accountId
    })
  }
}

export const SystemRoleService = new SystemRole()
