import { LogInfoType } from '@/types/logType'
import { Descriptions } from 'antd'
import { LogLevelTagUI } from './ToolsUI'

/**
 * 日志详情 UI 组件
 */
const LogDetailUI = (props: { logInfo?: LogInfoType }) => {
  const logInfo = props.logInfo
  return (
    <div key={logInfo?.account_id.toString()}>
      <Descriptions bordered size="small" column={1} labelStyle={{ width: 100 }}>
        <Descriptions.Item label="日志ID" key="log_id">
          {logInfo?.log_id.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="接口URI" key="uri">
          {logInfo?.uri.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="GET参数" key="get">
          {logInfo?.get}
        </Descriptions.Item>
        <Descriptions.Item label="POST参数" key="post">
          {logInfo?.post}
        </Descriptions.Item>
        <Descriptions.Item label="日志信息" key="message">
          {logInfo?.message}
        </Descriptions.Item>
        <Descriptions.Item label="日志级别" key="level">
          {LogLevelTagUI(logInfo?.level ? logInfo.level : 0)}
        </Descriptions.Item>
        <Descriptions.Item label="IP地址" key="ip">
          {logInfo?.ip}
        </Descriptions.Item>
        <Descriptions.Item label="操作账号" key="account">
          <div>
            {logInfo?.account_name}（{logInfo?.account_id.toString()}）
          </div>
        </Descriptions.Item>
        <Descriptions.Item label="创建时间" key="create_time">
          {logInfo?.create_time.toString()}
        </Descriptions.Item>
      </Descriptions>
    </div>
  )
}

export default LogDetailUI
