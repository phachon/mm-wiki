import httpRequest from './http'
import { InstallStatusResp, InstallCheckDBResp } from '../types/installType'
import Base from './Base'

const systemInstallUrl = {
  installStatus: '/system/install/status',
  installCheckDB: '/system/install/check_db',
  installInitData: '/system/install/init_data',
  installCreateAdmin: '/system/install/create_admin',
  installComplete: '/system/install/complete'
}

class SystemInstall extends Base {
  public constructor() {
    super()
  }

  public getInstallStatus(): Promise<InstallStatusResp> {
    const installStatusUrl = this.getProxyUrl(systemInstallUrl.installStatus)
    return httpRequest.get<InstallStatusResp>(installStatusUrl, {})
  }

  public checkDB(): Promise<InstallCheckDBResp> {
    const checkDBUrl = this.getProxyUrl(systemInstallUrl.installCheckDB)
    return httpRequest.get<InstallCheckDBResp>(checkDBUrl, {})
  }

  public initData(data: {}): Promise<any> {
    const initDataUrl = this.getProxyUrl(systemInstallUrl.installInitData)
    return httpRequest.post<any>(initDataUrl, {}, data)
  }

  public createAdmin(data: {}): Promise<any> {
    const createAdminUrl = this.getProxyUrl(systemInstallUrl.installCreateAdmin)
    return httpRequest.post<any>(createAdminUrl, {}, data)
  }

  public complete(): Promise<any> {
    const completeUrl = this.getProxyUrl(systemInstallUrl.installComplete)
    return httpRequest.post<any>(completeUrl, {}, {})
  }
}

export const SystemInstallService = new SystemInstall()
