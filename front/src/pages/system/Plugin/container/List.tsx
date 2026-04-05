import React, { useEffect, useState } from 'react'
import {
  PluginEditResp,
  PluginInfoType,
  PluginListItemType,
  PluginListResp
} from '@/types/pluginType'
import { SystemPluginService } from '@/services/SystemPlugin'
import { message, Modal, TablePaginationConfig } from 'antd'
import PluginListUI from '../component/ListUI'
import PluginSearchUI from '../component/SearchUI'
import PluginFormUI from '../component/FormUI'
import { initPagination } from '@/types/adminType'

let searchKeyWords = {}

const PluginList: React.FC = () => {
  const [pluginList, setPluginList] = useState<PluginListItemType[]>([])
  const [pagination, setPagination] = useState(initPagination)
  const [editModalOpen, setEditModalOpen] = useState<boolean>(false)
  const [editPluginInfo, setEditPluginInfo] = useState<PluginInfoType>()

  useEffect(() => {
    getPluginList(initPagination, {})
  }, [])

  const onListChange = (pageConfig: TablePaginationConfig, filters: any, sorter: any) => {
    getPluginList(pageConfig, searchKeyWords)
  }

  const onEditClick = (pluginInfo: PluginInfoType) => {
    SystemPluginService.getEditPluginInfo(pluginInfo.plugin_id)
      .then((editPluginInfo: PluginEditResp) => {
        setEditPluginInfo(editPluginInfo.plugin_info)
        setEditModalOpen(true)
      })
      .catch((e) => {
        console.log('修改插件 err:', e)
      })
  }

  const onEditSaveSubmit = (pluginInfo: PluginInfoType) => {
    SystemPluginService.modifyPlugin(pluginInfo)
      .then(() => {
        message.success('修改成功', 2, () => {
          setEditModalOpen(false)
          getPluginList(pagination, searchKeyWords)
        })
      })
      .catch(() => {
        message.error('修改失败', 2)
      })
  }

  const onDeleteConfirm = (pluginInfo: PluginInfoType) => {
    SystemPluginService.deletePlugin(pluginInfo.plugin_id)
      .then(() => {
        message.success('删除成功', 2, () => {
          getPluginList(pagination, searchKeyWords)
        })
      })
      .catch((e) => {
        console.log('删除插件失败:', e)
      })
  }

  const onStatusClick = (pluginInfo: PluginInfoType) => {
    const newStatus = pluginInfo.status === 1 ? 0 : 1
    SystemPluginService.updatePluginStatus(pluginInfo.plugin_id, newStatus)
      .then(() => {
        message.success('设置成功', 2, () => {
          getPluginList(pagination, searchKeyWords)
        })
      })
      .catch((e) => {
        console.log('设置状态失败:', e)
      })
  }

  const onSearchChange = (values: any) => {
    getPluginList(initPagination, values)
  }

  const onSearchReset = () => {
    getPluginList(initPagination, {})
  }

  const getPluginList = (pagination: TablePaginationConfig, searchValues: {}) => {
    const pageSize = pagination.pageSize
    const current = pagination.current
    searchKeyWords = searchValues
    SystemPluginService.getPluginList(pageSize, current, searchValues)
      .then((pluginList: PluginListResp) => {
        setPluginList(pluginList.list)
        setPagination({
          ...initPagination,
          current: pluginList.page_info?.page_num,
          pageSize: pluginList.page_info?.page_size,
          total: pluginList.page_info?.total_num
        })
      })
      .catch((e) => {
        console.log(e)
      })
  }

  return (
    <div className="panel">
      <PluginSearchUI onSearchChange={onSearchChange} onSearchReset={onSearchReset} />
      <PluginListUI
        listLoading={false}
        pagination={pagination}
        pluginList={pluginList}
        onListChange={onListChange}
        onEditClick={onEditClick}
        onDeleteConfirm={onDeleteConfirm}
        onStatusClick={onStatusClick}
      />
      <Modal
        title="插件修改"
        width={570}
        open={editModalOpen}
        onCancel={() => setEditModalOpen(false)}
        footer={null}
      >
        <PluginFormUI pluginInfo={editPluginInfo} onSaveSubmit={onEditSaveSubmit} />
      </Modal>
    </div>
  )
}

export default PluginList
