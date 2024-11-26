import { message, Modal } from 'antd'
import React, { useEffect, useState } from 'react'
import { SystemPrivilegeService } from '@/services/SystemPrivilege'
import { PrivilegeInfoType, PrivilegeListItemType } from '@/types/privilegeType'
import PrivilegeFormUI from '../component/FormUI'
import PrivilegeListTreeUI from '../component/ListTreeUI'

const PrivilegeList: React.FC = () => {
  const [privilegeList, setPrivilegeList] = useState<PrivilegeListItemType[]>([])
  const [parentPrivileges, setParentPrivileges] = useState<PrivilegeListItemType[]>([])
  const [editModalOpen, setEditModalOpen] = useState(false)
  const [editPrivilegeInfo, setEditPrivilegeInfo] = useState<PrivilegeInfoType>()
  const [apiMarks, setApiMarks] = useState<string[]>([])

  useEffect(() => {
    getPrivilegeList()
  }, [])

  /**
   * 获取权限列表
   */
  const getPrivilegeList = () => {
    SystemPrivilegeService.privilegeList().then((privilegeList) => {
      setPrivilegeList(privilegeList.list)
    })
  }

  /**
   * 修改点击操作
   * @param privilegeInfo 权限信息
   */
  const onEditClick = (privilegeInfo: PrivilegeInfoType) => {
    SystemPrivilegeService.getEditPrivilegeInfo(privilegeInfo.privilege_id)
      .then((resp) => {
        setParentPrivileges(resp.parent_privileges)
        setEditPrivilegeInfo(resp.privilege_info)
        setApiMarks(resp.api_marks)
        setEditModalOpen(true)
      })
      .catch((e) => {
        console.log('修改权限失败：', e)
      })
  }

  /**
   * 删除点击操作
   * @param privilegeInfo 权限信息
   */
  const onDeleteConfirm = (privilegeInfo: PrivilegeInfoType) => {
    SystemPrivilegeService.deletePrivilege(privilegeInfo)
      .then(() => {
        message.success('删除成功', 2, () => {
          getPrivilegeList()
        })
      })
      .catch((e) => {
        console.log('删除权限失败：', e)
      })
  }

  /**
   * 修改保存操作
   * @param privilegeInfo 权限信息
   */
  const onEditSaveSubmit = (privilegeInfo: PrivilegeInfoType) => {
    SystemPrivilegeService.modifyPrivilege(privilegeInfo)
      .then(() => {
        message.success('修改成功', 2, () => {
          setEditModalOpen(false)
          getPrivilegeList()
        })
      })
      .catch((e) => {
        console.log(e)
        message.error('修改失败', 2)
      })
  }

  /**
   * 修改取消弹窗
   */
  const editModalCancel = () => {
    setEditModalOpen(false)
  }

  return (
    <>
      <PrivilegeListTreeUI
        privilegeList={privilegeList}
        onEditClick={onEditClick}
        onDeleteConfirm={onDeleteConfirm}
      />
      <Modal
        title="权限修改"
        width={670}
        open={editModalOpen}
        onCancel={editModalCancel}
        footer={null}
      >
        <PrivilegeFormUI
          privilegeInfo={editPrivilegeInfo}
          onSaveSubmit={onEditSaveSubmit}
          parentPrivileges={parentPrivileges}
          allApiMarks={apiMarks}
        />
      </Modal>
    </>
  )
}

export default PrivilegeList
