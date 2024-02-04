import { Button, Form, Input, Modal, Radio, Select, Space, Tooltip, Typography } from 'antd'
import { EditLayoutForm, LayoutForm } from '../../../config/layout'
import { RoleInfoType } from '@/types/roleType'
import { QuestionCircleOutlined } from '@ant-design/icons'
import PrivilegeUI from './PrivilegeUI'
import { useEffect, useState } from 'react'
import { PrivilegeListItemType } from '@/types/privilegeType'
import { RoleRadioOptions } from './ToolsUI'

interface RoleFormUIProps {
  /**
   * 角色信息
   */
  roleInfo?: RoleInfoType
  /**
   * 表单的布局配置
   */
  formLayout?: {
    labelCol: { span: number }
    wrapperCol: { span: number }
  }
  /**
   * 权限列表
   */
  privilegeList?: PrivilegeListItemType[]

  /**
   * 角色保存方法
   * @param values 角色保存方法
   */
  onFinishSubmit: (values: any) => void
}

/**
 * 角色表单 UI 组件
 * @param props
 * @returns
 */
const RoleFormUI = (props: RoleFormUIProps) => {
  const [form] = Form.useForm()
  const [privilegeModalOpen, setPrivilegeModalOpen] = useState<boolean>(false)
  const [selectPrivilegeIds, setSelectPrivilegeIds] = useState<string[]>()
  const roleTypeOptions = RoleRadioOptions()

  const isEdit = props.roleInfo ? true : false
  let layoutForm = isEdit ? EditLayoutForm : LayoutForm
  if (props.formLayout) {
    layoutForm = props.formLayout
  }

  useEffect(() => {
    if (props.roleInfo) {
      form.setFieldsValue(props.roleInfo)
    }
  }, [props.roleInfo])

  return (
    <div className="panel-body">
      <Form name="basic" {...layoutForm} form={form} onFinish={props.onFinishSubmit}>
        {isEdit && (
          <Form.Item label="角色ID" name="role_id" rules={[{ required: true }]}>
            <Input disabled placeholder="请输入角色ID" />
          </Form.Item>
        )}

        <Form.Item
          label="角色名称"
          name="name"
          rules={[{ required: true, message: '请输入角色名称' }]}
        >
          <Input placeholder="请输入角色名" />
        </Form.Item>

        <Form.Item label="角色类型" name="role_type" rules={[{ required: true }]} initialValue={0}>
          <Form.Item name="role_type" noStyle>
            <Radio.Group options={roleTypeOptions} defaultValue={0} />
          </Form.Item>
          <Tooltip title="添加账号会自动选中默认角色">
            <Typography.Link>
              <QuestionCircleOutlined />
            </Typography.Link>
          </Tooltip>
        </Form.Item>

        {!isEdit && (
          <Form.Item
            label="角色权限"
            name="privilege_ids"
            rules={[{ required: true, message: '请选择角色权限' }]}
          >
            <Space.Compact style={{ width: '100%' }}>
              <Input placeholder="请选择角色权限" value={selectPrivilegeIds} disabled />
              <Button onClick={() => setPrivilegeModalOpen(true)}>选择权限</Button>
            </Space.Compact>
          </Form.Item>
        )}
        <Form.Item label="角色备注" name="remark" initialValue={''}>
          <Input placeholder="请输入备注信息" />
        </Form.Item>

        <Form.Item wrapperCol={{ offset: layoutForm.labelCol.span }}>
          <Button type="primary" htmlType="submit">
            保存
          </Button>
        </Form.Item>
      </Form>
      <Modal
        title="角色权限"
        width={1100}
        open={privilegeModalOpen}
        onCancel={() => setPrivilegeModalOpen(false)}
        footer={null}
      >
        <PrivilegeUI
          privilegeList={props.privilegeList ? props.privilegeList : []}
          defaultPrivilegeIds={selectPrivilegeIds}
          onFinishSubmit={(privilegeIds?: string[]) => {
            setSelectPrivilegeIds(privilegeIds)
            form.setFieldValue('privilege_ids', privilegeIds?.join(','))
            setPrivilegeModalOpen(false)
          }}
        />
      </Modal>
    </div>
  )
}

export default RoleFormUI
