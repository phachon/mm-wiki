import { getUrlConfig } from '../config/url'

class Base {
  private _proxyPrefix: string // proxy 前缀

  constructor() {
    this._proxyPrefix = getUrlConfig().proxyUrl
  }

  /**
   * 获取 proxy 的 url
   * @param url 接口地址
   * @returns string
   */
  public getProxyUrl(url: string): string {
    return this._proxyPrefix + url
  }
}

export default Base
