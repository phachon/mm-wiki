import httpRequest from './http'
import { ConfigMapResp } from '../types/configType'
import Base from './Base'

const systemConfigUrl = {
  configList: '/system/config/list',
  configModify: '/system/config/modify',
  emailSendTest: '/system/email/send_test'
}

class SystemConfig extends Base {
  public constructor() {
    super()
  }

  public getConfigList(): Promise<ConfigMapResp> {
    const configListUrl = this.getProxyUrl(systemConfigUrl.configList)
    return httpRequest.get<ConfigMapResp>(configListUrl, {})
  }

  public modifyConfig(data: {}): Promise<any> {
    const configModifyUrl = this.getProxyUrl(systemConfigUrl.configModify)
    return httpRequest.post<any>(configModifyUrl, {}, data)
  }

  public sendTestEmail(emailId: number, toAddress: string): Promise<any> {
    const sendTestEmailUrl = this.getProxyUrl(systemConfigUrl.emailSendTest)
    return httpRequest.post<any>(sendTestEmailUrl, {}, { email_id: emailId, to_address: toAddress })
  }
}

export const SystemConfigService = new SystemConfig()
