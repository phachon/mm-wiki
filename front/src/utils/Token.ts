/**
 * Token 管理
 */
class Token {
  // token 存储 key
  tokenKey: string
  // token 失效时间存储 Key
  tokenExpireKey: string
  // token 失效时间（秒）
  tokenExpireSecond: number

  constructor(tokenKey: string, expireSecond: number) {
    this.tokenKey = tokenKey || 'MM_ADMIN_TOKEN_KEY'
    this.tokenExpireKey = this.tokenKey + '_EXPIRE'
    this.tokenExpireSecond = expireSecond || 10 * 30 * 30 * 3600
  }

  // storageToken 存储 token 到本地
  storageToken(token: string): void {
    sessionStorage.setItem(this.tokenKey, token)
    let expireTime = (new Date().getTime() + this.tokenExpireSecond).toString()
    sessionStorage.setItem(this.tokenExpireKey, expireTime)
  }

  // getToken 获取 token
  getToken(): string {
    return sessionStorage.getItem(this.tokenKey) || ''
  }

  // removeToken 清除 token
  removeToken() {
    sessionStorage.removeItem(this.tokenKey)
    sessionStorage.removeItem(this.tokenExpireKey)
  }

  // 检查 token 是否失效
  checkTokenExpire(): boolean {
    let token = this.getToken()
    let limitTime = sessionStorage.getItem(this.tokenExpireKey)
    if (token && limitTime && token !== '' && parseInt(limitTime) > new Date().getTime()) {
      return true
    }
    return false
  }
}

export default Token
