import React, { useEffect, useState } from 'react'
import { message } from 'antd'
import { RoleService } from '@/services/Role'
import RoleFormUI from '../component/FormUI'
import { PrivilegeListItemType } from '@/types/privilegeType'
import { RoleAddResp } from '@/types/roleType'
import { useNavigate } from 'react-router-dom'

const RoleAdd: React.FC = () => {
  const [allPrivileges, setAllPrivileges] = useState<PrivilegeListItemType[]>([])
  const navigate = useNavigate()

  useEffect(() => {
    getAddRoleInfo()
  }, [])

  const getAddRoleInfo = () => {
    RoleService.getAddRoleInfo().then((addRoleInfo: RoleAddResp) => {
      setAllPrivileges(addRoleInfo.all_privilege)
    })
  }

  /**
   * 添加保存操作
   * @param values
   */
  const onSaveSubmit = (values: any) => {
    RoleService.saveRole(values).then(() => {
      message.success('保存成功', 2, () => {
        window.location.href = '/system/role/list'
      })
    })
  }

  return (
    <div className="pdt24">
      <RoleFormUI onSaveSubmit={onSaveSubmit} privilegeList={allPrivileges} />
    </div>
  )
}

export default RoleAdd
