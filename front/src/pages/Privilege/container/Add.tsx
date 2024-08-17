import { message } from 'antd'
import React, { useEffect, useState } from 'react'
import { PrivilegeService } from '@/services/Privilege'
import { PrivilegeInfoType, PrivilegeListItemType } from '@/types/privilegeType'
import PrivilegeFormUI from '../component/FormUI'
import { useNavigate } from 'react-router-dom'

const PrivilegeAdd: React.FC = () => {
  const [parentPrivileges, setParentPrivileges] = useState<PrivilegeListItemType[]>([])
  const [apiMarks, setApiMarks] = useState<string[]>([])
  const navigate = useNavigate()

  useEffect(() => {
    getAddPrivilegeInfo()
  }, [])

  /**
   * 获取添加权限需要信息
   */
  const getAddPrivilegeInfo = () => {
    PrivilegeService.getAddPrivilegeInfo().then((resp) => {
      setParentPrivileges(resp.parent_privileges)
      setApiMarks(resp.api_marks)
    })
  }

  /**
   * 保存权限
   */
  const onSaveSubmit = (values: PrivilegeInfoType) => {
    console.log('onSaveSubmit', values)
    PrivilegeService.savePrivilege(values)
      .then(() => {
        message.success('保存成功', 2, () => {
          navigate('/privilege/list')
        })
      })
      .catch((e) => {
        console.log(e)
      })
  }

  return (
    <div className="pdt24">
      <PrivilegeFormUI
        onSaveSubmit={onSaveSubmit}
        parentPrivileges={parentPrivileges}
        allApiMarks={apiMarks}
      />
    </div>
  )
}

export default PrivilegeAdd
