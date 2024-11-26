import httpRequest from './http'
import Base from './Base'

const logUrl = {
  list: '/system/log/list'
}

/**
 * SystemLog 系统 - 日志服务
 */
class SystemLog extends Base {
  public constructor() {
    super()
  }

  /**
   * logList 日志列表
   */
  public logList(
    pageSize: number | undefined,
    pageNum: number | undefined,
    logKeywords: {}
  ): Promise<any> {
    const logListUrl = this.getProxyUrl(logUrl.list)
    return httpRequest.get<any>(logListUrl, {
      page_size: pageSize,
      page_num: pageNum,
      keywords: JSON.stringify(logKeywords, (key: string, val: any) => {
        if (key == 'account_id') {
          return val != '' ? Number(val) : 0
        }
        return val
      })
    })
  }
}

export const SystemLogService = new SystemLog()
