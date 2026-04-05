import { PageInfoType } from './baseType'

export type PluginInfoType = {
  plugin_id: number
  name: string
  key: string
  description: string
  version: string
  author: string
  config_json: string
  status: number
  create_time: string
  update_time: string
}

export type PluginListItemType = PluginInfoType & {
  action?: {
    is_edit: number
    is_delete: number
  }
}

export type PluginListResp = {
  list: PluginListItemType[]
  page_info: PageInfoType
}

export type PluginEditResp = {
  plugin_info: PluginInfoType
}
