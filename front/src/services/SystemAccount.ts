import {
  AccountAddResp,
  AccountDetailResp,
  AccountEditResp,
  AccountListResp
} from '../types/accountType'
import httpRequest from './http'
import Base from './Base'

const systemAccountUrl = {
  add: '/system/account/add',
  save: '/system/account/save',
  edit: '/system/account/edit',
  modify: '/system/account/modify',
  detail: '/system/account/detail',
  updateStatus: '/system/account/update_status',
  list: '/system/account/list'
}

/**
 * SystemAccount 系统 - 账号服务
 */
class SystemAccount extends Base {
  public constructor() {
    super()
  }

  /**
   * getAddAccountInfo 获取添加账号信息
   * @returns
   */
  public getAddAccountInfo(): Promise<AccountAddResp> {
    const accountAddUrl = this.getProxyUrl(systemAccountUrl.add)
    return httpRequest.get<AccountAddResp>(accountAddUrl)
  }

  /**
   * saveAccount 添加保存账号
   */
  public saveAccount(accountInfo: {}): Promise<any> {
    const accountSaveUrl = this.getProxyUrl(systemAccountUrl.save)
    return httpRequest.post<any>(accountSaveUrl, {}, accountInfo)
  }

  /**
   * getEditAccountInfo 获取编辑账号信息
   */
  public getEditAccountInfo(account_id: bigint): Promise<AccountEditResp> {
    const accountEditUrl = this.getProxyUrl(systemAccountUrl.edit)
    return httpRequest.get<AccountEditResp>(accountEditUrl, {
      account_id: account_id
    })
  }

  /**
   * modifyAccount 修改保存账号
   */
  public modifyAccount(accountEditInfo: {}): Promise<any> {
    const accountModifyUrl = this.getProxyUrl(systemAccountUrl.modify)
    return httpRequest.post<any>(accountModifyUrl, {}, accountEditInfo)
  }

  /**
   * accountList 账号列表
   */
  public accountList(
    pageSize: number | undefined,
    pageNum: number | undefined,
    keywords: {}
  ): Promise<any> {
    const accountListUrl = this.getProxyUrl(systemAccountUrl.list)
    return httpRequest.get<AccountListResp>(accountListUrl, {
      page_size: pageSize,
      page_num: pageNum,
      keywords: JSON.stringify(keywords)
    })
  }

  /**
   * updateAccountStatus 修改账号状态
   */
  public updateAccountStatus(accountId: bigint, status: number): Promise<any> {
    let accountUpdateStatueUrl = this.getProxyUrl(systemAccountUrl.updateStatus)
    return httpRequest.post<any>(
      accountUpdateStatueUrl,
      {},
      {
        account_id: accountId,
        status: status
      }
    )
  }

  /**
   * getAccountDetail 获取账号详情
   * @param accountId 账号ID
   */
  public getAccountDetail(accountId: bigint): Promise<AccountDetailResp> {
    const accountDetailUrl = this.getProxyUrl(systemAccountUrl.detail)
    return httpRequest.get<AccountDetailResp>(accountDetailUrl, {
      account_id: accountId
    })
  }
}

export const SystemAccountService = new SystemAccount()
