import { Button, Form, Input, InputNumber, Select, TreeSelect } from 'antd'
import { EditLayoutForm } from '@/config/layout'
import { useEffect } from 'react'
const { SHOW_PARENT } = TreeSelect
import { DepartmentInfoType, DepartmentListItemType } from '@/types/departmentType'
import { getSelectTreeData, getTreeData } from './ToolsUI'

interface DepartmentFormUIProps {
  department?: DepartmentInfoType // 修改时传入的部门信息，添加时传上级部门信息
  departments: DepartmentListItemType[] // 上级部门列表
  onSaveSubmit: (values: any) => void // 保存操作方法
}

/**
 * 部门表单 UI 组件
 * @param props 部门表单 UI 组件数据
 * @returns 部门表单 UI 组件
 */
const DepartmentFormUI = (props: DepartmentFormUIProps) => {
  const [form] = Form.useForm()
  // 初始化表单数据
  useEffect(() => {
    if (!props.department) {
      return
    }
    form.setFieldsValue({
      department_id: props.department.department_id,
      parent_id: props.department.parent_id,
      name: props.department.name,
      sequence: Number(props.department.sequence)
    })
  }, [props.department])

  const checkIsShowParent = () => {
    return !(props.department?.parent_id == BigInt(0))
  }

  return (
    <div className="panel-body">
      <Form name="department-form" {...EditLayoutForm} onFinish={props.onSaveSubmit} form={form}>
        <Form.Item label="部门ID" name="department_id" hidden>
          <Input />
        </Form.Item>

        {
          // 上级部门ID不存在，不显示上级部门选择框
          checkIsShowParent() && (
            <Form.Item
              label="上级部门"
              name="parent_id"
              rules={[{ required: true, message: '请选择上级部门!' }]}
            >
              <TreeSelect
                showSearch
                style={{ width: '100%' }}
                dropdownStyle={{ maxHeight: 400, overflow: 'auto' }}
                placeholder="请选择上级部门"
                allowClear
                treeDefaultExpandAll
                treeNodeFilterProp="title"
                showCheckedStrategy={SHOW_PARENT}
                treeLine={true}
                treeData={getSelectTreeData(props.departments)}
                treeDataSimpleMode={true}
              />
            </Form.Item>
          )
        }

        <Form.Item
          label="部门名称"
          name="name"
          rules={[{ required: true, message: '请输入部门名称!' }]}
        >
          <Input placeholder="请输入昵称" />
        </Form.Item>

        <Form.Item label="排序号" name="sequence">
          <InputNumber
            style={{ width: '100%' }}
            placeholder="请输入排序号（越小越靠前）"
            min={0}
            max={10000}
          />
        </Form.Item>

        <Form.Item wrapperCol={{ offset: 4 }}>
          <Button type="primary" htmlType="submit">
            保存
          </Button>
        </Form.Item>
      </Form>
    </div>
  )
}

export default DepartmentFormUI
