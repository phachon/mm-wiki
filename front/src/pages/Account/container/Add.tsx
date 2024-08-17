import React, { useEffect, useState } from 'react'
import { AccountService } from '@/services/Account'
import { Card, message } from 'antd'
import AccountFormUI from '../component/FormUI'
import { AccountAddResp, AccountInfoType } from '@/types/accountType'
import { RoleInfoType } from '@/types/roleType'
import { useNavigate } from 'react-router-dom'

const AccountAdd: React.FC = () => {
  const [roleList, setRoleList] = useState<RoleInfoType[]>([])
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
      })
      .catch((e) => {
        console.log('获取添加账号信息 err:', e)
      })
  }

  /**
   * 添加账号保存
   * @param accountInfo
   */
  const onFinishSubmit = (accountInfo: AccountInfoType) => {
    AccountService.saveAccount(accountInfo)
      .then(() => {
        message.success('保存成功', 2, () => {
          navigate('/account/list')
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
      <AccountFormUI roleList={roleList} onFinishSubmit={onFinishSubmit} />
    </div>
  )
}

export default AccountAdd
