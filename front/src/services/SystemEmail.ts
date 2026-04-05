import httpRequest from './http'
import { EmailEditResp, EmailListResp } from '../types/emailType'
import Base from './Base'

const systemEmailUrl = {
  emailSave: '/system/email/save',
  emailEdit: '/system/email/edit',
  emailModify: '/system/email/modify',
  emailList: '/system/email/list',
  emailDelete: '/system/email/delete',
  emailUsed: '/system/email/used'
}

class SystemEmail extends Base {
  public constructor() {
    super()
  }

  public saveEmail(emailInfo: {}): Promise<any> {
    const saveEmailUrl = this.getProxyUrl(systemEmailUrl.emailSave)
    return httpRequest.post<any>(saveEmailUrl, {}, emailInfo)
  }

  public getEditEmailInfo(emailId: number): Promise<EmailEditResp> {
    const emailEditUrl = this.getProxyUrl(systemEmailUrl.emailEdit)
    return httpRequest.get<EmailEditResp>(emailEditUrl, {
      email_id: emailId
    })
  }

  public modifyEmail(editEmailInfo: {}): Promise<any> {
    const emailModifyUrl = this.getProxyUrl(systemEmailUrl.emailModify)
    return httpRequest.post<any>(emailModifyUrl, {}, editEmailInfo)
  }

  public getEmailList(
    pageSize?: number,
    pageNum?: number,
    keywords?: {}
  ): Promise<EmailListResp> {
    const emailListUrl = this.getProxyUrl(systemEmailUrl.emailList)
    return httpRequest.get<EmailListResp>(emailListUrl, {
      page_size: pageSize,
      page_num: pageNum,
      keywords: JSON.stringify(keywords)
    })
  }

  public deleteEmail(emailId: number): Promise<any> {
    const deleteEmailUrl = this.getProxyUrl(systemEmailUrl.emailDelete)
    return httpRequest.post<any>(deleteEmailUrl, {}, { email_id: emailId })
  }

  public setEmailUsed(emailId: number): Promise<any> {
    const emailUsedUrl = this.getProxyUrl(systemEmailUrl.emailUsed)
    return httpRequest.post<any>(emailUsedUrl, {}, { email_id: emailId })
  }
}

export const SystemEmailService = new SystemEmail()
