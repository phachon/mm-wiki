import httpRequest from './http'
import { PluginEditResp, PluginListResp } from '../types/pluginType'
import Base from './Base'

const systemPluginUrl = {
  pluginSave: '/system/plugin/save',
  pluginEdit: '/system/plugin/edit',
  pluginModify: '/system/plugin/modify',
  pluginList: '/system/plugin/list',
  pluginDelete: '/system/plugin/delete',
  pluginUpdateStatus: '/system/plugin/update_status'
}

class SystemPlugin extends Base {
  public constructor() {
    super()
  }

  public savePlugin(pluginInfo: {}): Promise<any> {
    const savePluginUrl = this.getProxyUrl(systemPluginUrl.pluginSave)
    return httpRequest.post<any>(savePluginUrl, {}, pluginInfo)
  }

  public getEditPluginInfo(pluginId: number): Promise<PluginEditResp> {
    const pluginEditUrl = this.getProxyUrl(systemPluginUrl.pluginEdit)
    return httpRequest.get<PluginEditResp>(pluginEditUrl, {
      plugin_id: pluginId
    })
  }

  public modifyPlugin(editPluginInfo: {}): Promise<any> {
    const pluginModifyUrl = this.getProxyUrl(systemPluginUrl.pluginModify)
    return httpRequest.post<any>(pluginModifyUrl, {}, editPluginInfo)
  }

  public getPluginList(
    pageSize?: number,
    pageNum?: number,
    keywords?: {}
  ): Promise<PluginListResp> {
    const pluginListUrl = this.getProxyUrl(systemPluginUrl.pluginList)
    return httpRequest.get<PluginListResp>(pluginListUrl, {
      page_size: pageSize,
      page_num: pageNum,
      keywords: JSON.stringify(keywords)
    })
  }

  public deletePlugin(pluginId: number): Promise<any> {
    const deletePluginUrl = this.getProxyUrl(systemPluginUrl.pluginDelete)
    return httpRequest.post<any>(deletePluginUrl, {}, { plugin_id: pluginId })
  }

  public updatePluginStatus(pluginId: number, status: number): Promise<any> {
    const updateStatusUrl = this.getProxyUrl(systemPluginUrl.pluginUpdateStatus)
    return httpRequest.post<any>(updateStatusUrl, {}, { plugin_id: pluginId, status: status })
  }
}

export const SystemPluginService = new SystemPlugin()
