import React, { useEffect, useState } from 'react'
import { TablePaginationConfig } from 'antd'
import LogListUI from '../component/ListUI'
import LogSearchUI from '../component/SearchUI'
import { LogInfoType, LogListResp } from '@/types/logType'
import { LogService } from '@/services/Log'
import { initPagination } from '@/types/adminType'

let searchKeyWords = {}

const LogList: React.FC = () => {
  const [logList, setLogList] = useState<LogInfoType[]>([])
  const [pagination, setPagination] = useState(initPagination)

  useEffect(() => {
    getLogList(initPagination, {})
  }, [])

  /**
   * 列表分页处理
   * @param pageConfig
   * @param filters
   * @param sorter
   */
  const onListChange = (pageConfig: TablePaginationConfig, filters: any, sorter: any) => {
    getLogList(pageConfig, searchKeyWords)
  }

  /**
   * 搜索查询操作
   * @param values
   */
  const onSearchChange = (values: {}) => {
    getLogList(initPagination, values)
  }

  /**
   * 搜索重置操作
   */
  const onResetChange = () => {
    getLogList(initPagination, {})
  }

  /**
   * 获取日志列表
   * @param pagination
   * @param searchValues
   */
  const getLogList = (pagination: TablePaginationConfig, searchValues: {}) => {
    const pageSize = pagination.pageSize
    const current = pagination.current
    searchKeyWords = searchValues
    LogService.logList(pageSize, current, searchValues)
      .then((logList: LogListResp) => {
        setLogList(logList.list)
        setPagination({
          ...initPagination,
          current: logList.page_info?.page_num,
          pageSize: logList.page_info?.page_size,
          total: logList.page_info?.total_num
        })
      })
      .catch((e) => {
        console.log('获取日志列表失败err:', e)
      })
  }

  return (
    <div className="panel">
      <LogSearchUI onSearchChange={onSearchChange} onResetChange={onResetChange} />
      <LogListUI
        listLoading={false}
        pagination={pagination}
        logList={logList}
        onListChange={onListChange}
      />
    </div>
  )
}

export default LogList
