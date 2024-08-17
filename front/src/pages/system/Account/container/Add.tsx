import React, { useEffect, useState } from 'react'
import { AccountService } from '@/services/Account'
import { Card, message } from 'antd'
import AccountFormUI from '../component/FormUI'
import { AccountAddResp, AccountInfoType } from '@/types/accountType'
import { RoleInfoType } from '@/types/roleType'
import { useNavigate } from 'react-router-dom'
import { DepartmentInfoType } from '@/types/departmentType'

const AccountAdd: React.FC = () => {
  const [roleList, setRoleList] = useState<RoleInfoType[]>([])
  const [departmentList, setDepartmentList] = useState<DepartmentInfoType[]>([])
  const navigate = useNavigate()

  useEffect(() => {
    getAddAcountInfo()
  }, [])

  /**
   * 获取添加账号信息
   */
  const getAddAcountInfo = () => {
    AccountService.getAddAccountInfo()
      .then((accountAddResp: AccountAddResp) => {
        setRoleList(accountAddResp?.roles)
        setDepartmentList(accountAddResp?.departments)
      })
      .catch((e) => {
        console.log('获取添加账号信息 err:', e)
      })
  }

  /**
   * 添加账号保存
   * @param accountInfo
   */
  const onSaveSubmit = (accountInfo: AccountInfoType) => {
    AccountService.saveAccount(accountInfo)
      .then(() => {
        message.success('保存成功', 2, () => {
          window.location.href = '/system/account/list'
        })
      })
      .catch((e) => {
        console.log(e)
      })
  }

  /**
   * 返回组件
   */
  return (
    <div className="pdt24">
      <AccountFormUI roleList={roleList} departments={departmentList} onSaveSubmit={onSaveSubmit} />
    </div>
  )
}

export default AccountAdd
