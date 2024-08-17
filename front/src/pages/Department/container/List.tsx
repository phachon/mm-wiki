import React, { useEffect, useState } from 'react'
import DepartmentTreeUI from '../component/TreeUI'
import { message, Modal } from 'antd'
import DepartmentFormUI from '../component/FormUI'
import { DepartmentService } from '@/services/Department'
import {
  DepartmentEditResp,
  DepartmentInfoType,
  DepartmentListItemType,
  DepartmentListResp
} from '@/types/departmentType'

const DepartmentList: React.FC = () => {
  const [departments, setDepartments] = useState<DepartmentListItemType[]>([])
  const [formVisible, setFormVisible] = useState(false)
  const [department, setDepartment] = useState<DepartmentInfoType>()
  const [modalTitle, setModalTitle] = useState('')
  const [deleleVisible, setDeleteVisible] = useState(false)

  useEffect(() => {
    getDepartmentList()
  }, [])

  const getDepartmentList = () => {
    // 获取部门列表的逻辑
    DepartmentService.departmentList()
      .then((res: DepartmentListResp) => {
        setDepartments(res.list)
      })
      .catch((e) => {
        console.log('获取部门列表 err:', e)
      })
  }

  /**
   * 添加部门保存
   * @param accountInfo
   */
  const onSaveSubmit = (values: DepartmentInfoType) => {
    // 判断是添加还是编辑
    if (values.department_id) {
      modifyDepartment(values)
    } else {
      saveDepartment(values)
    }
  }

  const saveDepartment = (values: DepartmentInfoType) => {
    DepartmentService.saveDepartment(values)
      .then(() => {
        message.success('保存成功', 2, () => {
          getDepartmentList()
          setFormVisible(false)
        })
      })
      .catch((e) => {
        console.log('添加部门保存 err:', e)
      })
  }

  const modifyDepartment = (values: DepartmentInfoType) => {
    // 编辑
    DepartmentService.modifyDepartment(values)
      .then(() => {
        message.success('保存成功', 2, () => {
          getDepartmentList()
          setFormVisible(false)
        })
      })
      .catch((e) => {
        console.log('编辑部门保存 err:', e)
      })
    return
  }

  const onAddClick = (parentId: string) => {
    setModalTitle('添加部门')
    setDepartment({
      department_id: BigInt(0),
      name: '',
      parent_id: BigInt(parentId),
      parent_ids: '',
      sequence: 0,
      create_time: '',
      update_time: ''
    })
    setFormVisible(true)
  }

  const onEditClick = (departmentId: string) => {
    DepartmentService.getEditDepartmentInfo(BigInt(departmentId))
      .then((res: DepartmentEditResp) => {
        setModalTitle('编辑部门')
        setDepartment(res.department)
        setFormVisible(true)
      })
      .catch((e) => {
        console.log('获取编辑部门信息 err:', e)
      })
  }

  const onDeleteClick = (departmentId: string) => {
    console.log('Delete node', departmentId)
    // 删除节点的逻辑
  }

  /**
   * 返回组件
   */
  return (
    <div className="panel" style={{ height: '100%' }}>
      <DepartmentTreeUI
        departments={departments}
        onAddClick={onAddClick}
        onEditClick={onEditClick}
        onDeleteClick={onDeleteClick}
        onSaveSubmit={onSaveSubmit}
      />
      <Modal
        title={modalTitle ? modalTitle : '添加部门'}
        width={570}
        open={formVisible}
        onCancel={() => {
          setFormVisible(false)
        }}
        footer={null}
      >
        <DepartmentFormUI
          departments={departments}
          department={department}
          onSaveSubmit={onSaveSubmit}
        />
      </Modal>
    </div>
  )
}

export default DepartmentList
