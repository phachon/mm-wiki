import React, { useEffect, useState } from 'react'
import { message } from 'antd'
import { RoleService } from '@/services/Role'
import SpaceFormUI from '../component/FormUI'
import { PrivilegeListItemType } from '@/types/privilegeType'
import { RoleAddResp } from '@/types/roleType'
import { useNavigate } from 'react-router-dom'
import { AccountInfoType } from '@/types/accountType'
import Space from '..'
import { SpaceService } from '@/services/Space'
import { SpaceAddResp } from '@/types/spaceType'

const SpaceAdd: React.FC = () => {
  // 账号列表
  const [accountList, setAccountList] = useState<AccountInfoType[]>([])

  useEffect(() => {
    getAddSpaceInfo()
  }, [])

  const getAddSpaceInfo = () => {
    SpaceService.getAddSpaceInfo()
      .then((addSpaceInfo: SpaceAddResp) => {
        setAccountList(addSpaceInfo.account_list)
      })
      .catch((e) => {
        console.log('获取添加空间信息失败err:', e)
      })
  }

  /**
   * 添加保存操作
   * @param values
   */
  const onSaveSubmit = (values: any) => {
    console.log('onSaveSubmit values:', values)
    // 保存操作
    SpaceService.saveSpace(values)
      .then(() => {
        message.success('添加空间成功', 2, () => {
          window.location.href = '/system/space/list'
        })
      })
      .catch((e) => {
        console.log('添加空间失败err:', e)
      })
  }

  return (
    <div className="pdt24">
      <SpaceFormUI onSaveSubmit={onSaveSubmit} accountList={accountList} />
    </div>
  )
}

export default SpaceAdd
