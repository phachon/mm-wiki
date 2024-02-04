import { PageInfoType } from './baseType'

// 日志级别定义
export const LogLevelTypes = [
  {
    level: 0,
    name: '全部',
    color: 'cyan'
  },
  {
    level: 1,
    name: 'Trace',
    color: 'success'
  },
  {
    level: 2,
    name: 'Debug',
    color: 'purple'
  },
  {
    level: 3,
    name: 'Info',
    color: 'blue'
  },
  {
    level: 4,
    name: 'Warn',
    color: 'warning'
  },
  {
    level: 5,
    name: 'Error',
    color: 'red'
  },
  {
    level: 6,
    name: 'Fatal',
    color: 'magenta'
  }
]

// LogInfoType 日志信息
export type LogInfoType = {
  log_id: BigInt // 日志ID
  uri: string // 接口uri
  get: string // get参数
  post: string // post参数
  message: string // 信息
  level: number // 日志级别
  file: string // 文件
  line: number // 行数
  ip: string // IP
  account_id: bigint // 账号ID
  account_name: string // 账号名
  create_time: string // 创建时间
}

/**
 * LogListResp 日志列表返回结构
 */
export type LogListResp = {
  list: LogInfoType[]
  page_info: PageInfoType
}
