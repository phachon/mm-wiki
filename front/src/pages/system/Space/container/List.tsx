import React, { useEffect, useState } from 'react'
import { message, Modal, TablePaginationConfig } from 'antd'
import SpaceListUI from '../component/ListUI'
import SpaceSearchUI from '../component/SearchUI'
import SpaceFormUI from '../component/FormUI'
import { AccountInfoType } from '@/types/accountType'
import { initPagination } from '@/types/adminType'
import {
  SpaceAdminListResp,
  SpaceEditResp,
  SpaceInfoType,
  SpaceListItemType,
  SpaceListResp
} from '@/types/spaceType'
import { SystemSpaceService } from '@/services/SystemSpace'
import SpaceAdminListUI from '../component/AdminListUI'

let searchKeyWords = {}
let adminListSpaceInfo: SpaceInfoType // 账号列表空间信息

const SpaceList: React.FC = () => {
  // 空间列表相关 state
  const [spaceList, setSpaceList] = useState<SpaceListItemType[]>([])
  const [pagination, setPagination] = useState(initPagination)
  const [editModalOpen, setEditModalOpen] = useState<boolean>(false)
  // 空间修改相关 state
  const [editSpaceInfo, setEditSpaceInfo] = useState<SpaceInfoType>()
  // 空间账号相关 state
  const [adminModalOpen, setAdminModalOpen] = useState<boolean>(false)
  const [spaceAdminList, setSpaceAdminList] = useState<AccountInfoType[]>([])
  const [selectedAccountList, setSelectedAccountList] = useState<AccountInfoType[]>([])

  useEffect(() => {
    getSpaceList(initPagination, {})
  }, [])

  /**
   * 列表分页处理
   * @param pageConfig
   * @param filters
   * @param sorter
   */
  const onSpaceListChange = (pageConfig: TablePaginationConfig, filters: any, sorter: any) => {
    getSpaceList(pageConfig, searchKeyWords)
  }

  /**
   * 修改点击操作
   * @param spaceInfo
   */
  const onEditClick = (spaceInfo: SpaceInfoType) => {
    SystemSpaceService.getEditSpaceInfo(spaceInfo.space_id)
      .then((editSpaceInfo: SpaceEditResp) => {
        setEditSpaceInfo(editSpaceInfo.space_info)
        setEditModalOpen(true)
      })
      .catch((e) => {
        console.log('修改空间失败err:', e)
      })
  }

  /**
   * 管理员列表点击操作
   * @param spaceInfo 空间信息
   */
  const onAdminListClick = (spaceInfo: SpaceInfoType) => {
    adminListSpaceInfo = spaceInfo
    getSpaceAdminList(spaceInfo.space_id)
  }

  /**
   * 修改保存操作
   * @param spaceInfo
   */
  const onEditSaveSubmit = (spaceInfo: SpaceInfoType) => {
    SystemSpaceService.modifySpace(spaceInfo)
      .then(() => {
        message.success('修改成功', 2, () => {
          setEditModalOpen(false)
          getSpaceList(pagination, searchKeyWords)
        })
      })
      .catch((e) => {
        console.log('修改失败err:', e)
      })
  }

  /**
   * 确认删除操作
   * @param spaceInfo
   */
  const onDeleteConfirm = (spaceInfo: SpaceInfoType) => {
    SystemSpaceService.deleteSpace(spaceInfo.space_id)
      .then(() => {
        message.success('删除成功', 2, () => {
          getSpaceList(pagination, searchKeyWords)
        })
      })
      .catch((e) => {
        console.log('删除空间失败:', e)
      })
  }

  /**
   * 搜索查询操作
   * @param values
   */
  const onSearchChange = (values: any) => {
    getSpaceList(initPagination, values)
  }

  /**
   * 搜索重置操作
   */
  const onSearchReset = () => {
    getSpaceList(initPagination, {})
  }

  /**
   * 获取空间列表
   * @param pagination
   * @param searchValues
   */
  const getSpaceList = (pagination: TablePaginationConfig, searchValues: {}) => {
    const pageSize = pagination.pageSize
    const current = pagination.current
    searchKeyWords = searchValues
    SystemSpaceService.getSpaceList(pageSize, current, searchValues)
      .then((resp: SpaceListResp) => {
        setSpaceList(resp.list)
        setPagination({
          ...initPagination,
          current: resp.page_info?.page_num,
          pageSize: resp.page_info?.page_size,
          total: resp.page_info?.total_num
        })
      })
      .catch((e) => {
        console.log(e)
      })
  }

  /**
   * 获取空间管理员账号列表
   * @param spaceId
   */
  const getSpaceAdminList = (spaceId: number) => {
    SystemSpaceService.getAdminList(spaceId)
      .then((resp: SpaceAdminListResp) => {
        setSpaceAdminList(resp.admin_list)
        setSelectedAccountList(resp.selected_list)
        if (!adminModalOpen) {
          setAdminModalOpen(true)
        }
      })
      .catch((e) => {
        console.log('获得空间管理员 err:', e)
      })
  }

  /**
   * 管理员账号移除操作
   * @param accountInfo 账号信息
   */
  const onAdminRemoveChange = (accountInfo: AccountInfoType) => {
    SystemSpaceService.removeSpaceAdmin(adminListSpaceInfo?.space_id, accountInfo.account_id)
      .then(() => {
        message.success('移除成功', 2, () => {
          getSpaceAdminList(adminListSpaceInfo?.space_id)
        })
      })
      .catch((e) => {
        console.log('移除账号失败err:', e)
      })
  }

  const onAdminAddChange = (values: any) => {
    console.log('onAdminAddChange values:', values)
    const accountIds = values.admin_account_ids
    SystemSpaceService.addSpaceAdmin(adminListSpaceInfo?.space_id, accountIds)
      .then(() => {
        message.success('添加成功', 2, () => {
          getSpaceAdminList(adminListSpaceInfo?.space_id)
        })
      })
      .catch((e) => {
        console.log('添加账号失败err:', e)
      })
  }

  return (
    <div className="panel">
      <SpaceSearchUI onSearchChange={onSearchChange} onSearchReset={onSearchReset} />
      <SpaceListUI
        listLoading={false}
        pagination={pagination}
        spaceList={spaceList}
        onListChange={onSpaceListChange}
        onEditClick={onEditClick}
        onDeleteConfirm={onDeleteConfirm}
        onAdminListClick={onAdminListClick}
      />
      <Modal
        title="空间修改"
        width={550}
        open={editModalOpen}
        onCancel={() => setEditModalOpen(false)}
        footer={null}
      >
        <SpaceFormUI spaceInfo={editSpaceInfo} onSaveSubmit={onEditSaveSubmit} />
      </Modal>
      <Modal
        title="管理员账号"
        width={900}
        open={adminModalOpen}
        onCancel={() => setAdminModalOpen(false)}
        footer={null}
      >
        <SpaceAdminListUI
          adminList={spaceAdminList}
          selectedAccountList={selectedAccountList}
          onRemoveAdminChange={onAdminRemoveChange}
          onAddAdminChange={onAdminAddChange}
        />
      </Modal>
    </div>
  )
}

export default SpaceList
