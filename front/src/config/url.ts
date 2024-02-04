interface UrlConfig {
  proxyUrl: string
}

const devUrlConfig: UrlConfig = {
  proxyUrl: '/api'
}

const proUrlConfig: UrlConfig = {
  proxyUrl: '/api'
}

/**
 * 获取 urlConfig
 * @constructor
 */
export function getUrlConfig(): UrlConfig {
  if (process.env.NODE_ENV === 'development') {
    return devUrlConfig
  }
  return proUrlConfig
}

interface httpDebug {
  /**
   * 跳过登录
   */
  skip_auth_login: boolean

  /**
   * 跳过权限
   */
  skip_auth_premission: boolean
}

/**
 * 获取 debug 参数
 * @returns string
 */
export function getHttpDebug(): string {
  if (process.env.NODE_ENV !== 'development') {
    return ''
  }
  const debugConf: httpDebug = {
    skip_auth_login: false,
    skip_auth_premission: true
  }
  return JSON.stringify(debugConf)
}
