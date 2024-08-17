import { Descriptions, Modal, Table, TablePaginationConfig, Tag } from 'antd'
import { useState } from 'react'
import { LogInfoType } from '@/types/logType'
import { LogLevelTagUI } from './ToolsUI'
import LogDetailUI from './DetailUI'

interface LogListUIProps {
  listLoading: boolean
  logList: LogInfoType[]
  pagination: TablePaginationConfig
  onListChange: (pageConfig: TablePaginationConfig, filters: any, sorter: any) => void
}

/**
 * 日志列表 UI 组件
 * @param props
 * @returns
 */
const LogListUI = (props: LogListUIProps) => {
  const [detailModalOpen, setDetailModalOpen] = useState(false)
  const [detailLogInfo, setDetailLogInfo] = useState<LogInfoType>()

  return (
    <div className="panel-body">
      <Table
        rowKey={'log_id'}
        bordered={true}
        dataSource={props.logList}
        loading={props.listLoading}
        pagination={props.pagination}
        onChange={props.onListChange}
        footer={() => null}
      >
        <Table.Column
          title={'日志ID'}
          dataIndex="log_id"
          width={80}
          key={'log_id'}
          align={'center'}
        />
        <Table.Column
          title={'账号'}
          dataIndex="account_id"
          key={'account_id'}
          width={250}
          render={(accountId: bigint, logInfo: LogInfoType) => (
            <div>
              {logInfo.account_name}（{logInfo.account_id.toString()}）
            </div>
          )}
        />
        <Table.Column
          title={'日志内容'}
          dataIndex="message"
          key={'message'}
          render={(message: string, logInfo: LogInfoType) => (
            <a
              href="#!"
              onClick={() => {
                setDetailModalOpen(true)
                setDetailLogInfo(logInfo)
              }}
            >
              {message}
            </a>
          )}
        />
        <Table.Column
          title={'日志级别'}
          dataIndex="level"
          key={'level'}
          width={120}
          align={'center'}
          render={LogLevelTagUI}
        />
        <Table.Column
          title={'创建时间'}
          dataIndex="create_time"
          key={'create_time'}
          width={200}
          align={'center'}
        />
      </Table>
      <Modal
        title="账号详情"
        width={750}
        open={detailModalOpen}
        onCancel={() => {
          setDetailModalOpen(false)
        }}
        footer={null}
      >
        <LogDetailUI logInfo={detailLogInfo} />
      </Modal>
    </div>
  )
}

export default LogListUI
