import React, { useEffect, useState } from 'react'
import {
  LoginAuthEditResp,
  LoginAuthInfoType,
  LoginAuthListItemType,
  LoginAuthListResp
} from '@/types/loginAuthType'
import { SystemLoginAuthService } from '@/services/SystemLoginAuth'
import { message, Modal, TablePaginationConfig } from 'antd'
import LoginAuthListUI from '../component/ListUI'
import LoginAuthSearchUI from '../component/SearchUI'
import LoginAuthFormUI from '../component/FormUI'
import { initPagination } from '@/types/adminType'

let searchKeyWords = {}

const LoginAuthList: React.FC = () => {
  const [loginAuthList, setLoginAuthList] = useState<LoginAuthListItemType[]>([])
  const [pagination, setPagination] = useState(initPagination)
  const [editModalOpen, setEditModalOpen] = useState<boolean>(false)
  const [editLoginAuthInfo, setEditLoginAuthInfo] = useState<LoginAuthInfoType>()

  useEffect(() => {
    getLoginAuthList(initPagination, {})
  }, [])

  const onListChange = (pageConfig: TablePaginationConfig, filters: any, sorter: any) => {
    getLoginAuthList(pageConfig, searchKeyWords)
  }

  const onEditClick = (loginAuthInfo: LoginAuthInfoType) => {
    SystemLoginAuthService.getEditLoginAuthInfo(loginAuthInfo.login_auth_id)
      .then((editLoginAuthInfo: LoginAuthEditResp) => {
        setEditLoginAuthInfo(editLoginAuthInfo.login_auth_info)
        setEditModalOpen(true)
      })
      .catch((e) => {
        console.log('修改认证 err:', e)
      })
  }

  const onEditSaveSubmit = (loginAuthInfo: LoginAuthInfoType) => {
    SystemLoginAuthService.modifyLoginAuth(loginAuthInfo)
      .then(() => {
        message.success('修改成功', 2, () => {
          setEditModalOpen(false)
          getLoginAuthList(pagination, searchKeyWords)
        })
      })
      .catch(() => {
        message.error('修改失败', 2)
      })
  }

  const onDeleteConfirm = (loginAuthInfo: LoginAuthInfoType) => {
    SystemLoginAuthService.deleteLoginAuth(loginAuthInfo.login_auth_id)
      .then(() => {
        message.success('删除成功', 2, () => {
          getLoginAuthList(pagination, searchKeyWords)
        })
      })
      .catch((e) => {
        console.log('删除认证失败:', e)
      })
  }

  const onUsedClick = (loginAuthInfo: LoginAuthInfoType) => {
    SystemLoginAuthService.setLoginAuthUsed(loginAuthInfo.login_auth_id)
      .then(() => {
        message.success('设置成功', 2, () => {
          getLoginAuthList(pagination, searchKeyWords)
        })
      })
      .catch((e) => {
        console.log('设置使用失败:', e)
      })
  }

  const onSearchChange = (values: any) => {
    getLoginAuthList(initPagination, values)
  }

  const onSearchReset = () => {
    getLoginAuthList(initPagination, {})
  }

  const getLoginAuthList = (pagination: TablePaginationConfig, searchValues: {}) => {
    const pageSize = pagination.pageSize
    const current = pagination.current
    searchKeyWords = searchValues
    SystemLoginAuthService.getLoginAuthList(pageSize, current, searchValues)
      .then((loginAuthList: LoginAuthListResp) => {
        setLoginAuthList(loginAuthList.list)
        setPagination({
          ...initPagination,
          current: loginAuthList.page_info?.page_num,
          pageSize: loginAuthList.page_info?.page_size,
          total: loginAuthList.page_info?.total_num
        })
      })
      .catch((e) => {
        console.log(e)
      })
  }

  return (
    <div className="panel">
      <LoginAuthSearchUI onSearchChange={onSearchChange} onSearchReset={onSearchReset} />
      <LoginAuthListUI
        listLoading={false}
        pagination={pagination}
        loginAuthList={loginAuthList}
        onListChange={onListChange}
        onEditClick={onEditClick}
        onDeleteConfirm={onDeleteConfirm}
        onUsedClick={onUsedClick}
      />
      <Modal
        title="认证修改"
        width={570}
        open={editModalOpen}
        onCancel={() => setEditModalOpen(false)}
        footer={null}
      >
        <LoginAuthFormUI loginAuthInfo={editLoginAuthInfo} onSaveSubmit={onEditSaveSubmit} />
      </Modal>
    </div>
  )
}

export default LoginAuthList
