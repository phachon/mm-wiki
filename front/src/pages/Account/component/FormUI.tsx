import { Button, Card, Form, Input, Select } from 'antd'
import { EditLayoutForm, LayoutForm } from '../../../config/layout'
import { AccountInfoType } from '@/types/accountType'
import { RoleInfoType, RoleTypeAccountDefaultRole } from '@/types/roleType'
import { DefaultOptionType } from 'antd/lib/select'
import { useEffect } from 'react'
import { SelectRoleLabelUI } from '@/pages/Role/component/ToolsUI'

interface AccountFormUIProps {
  /**
   * 角色列表
   */
  roleList?: RoleInfoType[]
  /**
   * 账号信息
   */
  accountInfo?: AccountInfoType
  /**
   * 表单布局
   */
  formLayout?: {
    labelCol: { span: number }
    wrapperCol: { span: number }
  }
  /**
   * 保存操作方法
   * @param values
   */
  onFinishSubmit: (values: any) => void
}

/**
 * 获取角色 Option 数据
 * @param roleList 角色列表
 * @returns
 */
const getRoleOptions = (roleList: RoleInfoType[] | undefined): DefaultOptionType[] => {
  let roleOptions: DefaultOptionType[] = []
  if (!roleList) {
    return roleOptions
  }
  roleList.forEach((roleInfo) => {
    roleOptions.push({
      label: SelectRoleLabelUI(roleInfo, { margin: 0 }),
      value: String(roleInfo.role_id)
    })
  })
  return roleOptions
}

/**
 * 从所有的角色中获取默认的账号角色
 * @param roleList
 * @returns 角色ID列表
 */
const getDefaultAccountRoleIds = (roleList: RoleInfoType[] | undefined): string => {
  if (!roleList) {
    return ''
  }
  let defaultAccountRoleIds: string[] = []
  roleList.forEach((roleInfo) => {
    if (roleInfo.role_type === RoleTypeAccountDefaultRole) {
      defaultAccountRoleIds.push(roleInfo.role_id.toString())
    }
  })
  return defaultAccountRoleIds.length > 0 ? defaultAccountRoleIds.toString() : ''
}

/**
 * 账号表单 UI 组件
 * @param props 账号表单 UI 组件数据
 * @returns 账号表单 UI 组件
 */
const AccountFormUI = (props: AccountFormUIProps) => {
  const [form] = Form.useForm()
  const isEdit = props.accountInfo ? true : false

  let layoutForm = isEdit ? EditLayoutForm : LayoutForm
  if (props.formLayout) {
    layoutForm = props.formLayout
  }

  useEffect(() => {
    const defaultRoleIds = getDefaultAccountRoleIds(props.roleList)
    if (!isEdit && defaultRoleIds !== '') {
      form.setFieldsValue({
        role_ids: defaultRoleIds
      })
    }
  }, [props.roleList, isEdit])

  useEffect(() => {
    if (props.accountInfo) {
      form.setFieldsValue(props.accountInfo)
    }
  }, [props.accountInfo])

  return (
    <div className="panel-body">
      <Form {...layoutForm} name="basic" onFinish={props.onFinishSubmit} form={form}>
        {isEdit && (
          <Form.Item label="账号id" name="account_id" rules={[{ required: true }]}>
            <Input disabled placeholder="请输入账号ID" />
          </Form.Item>
        )}

        <Form.Item
          label="账号名"
          name="name"
          rules={[{ required: true, message: '请输入账号名!' }]}
        >
          <Input placeholder="请输入账号名" />
        </Form.Item>

        <Form.Item
          label="昵称"
          name="given_name"
          rules={[{ required: true, message: '请输入昵称!' }]}
        >
          <Input placeholder="请输入昵称" />
        </Form.Item>

        <Form.Item
          label="角色"
          name="role_ids"
          rules={[{ required: true, message: '请选择角色!' }]}
          getValueProps={(val) => {
            let value = val ? val.split(',') : []
            return { value: value }
          }}
          getValueFromEvent={(values: string[]) => {
            return values.length > 0 ? values.toString() : ''
          }}
        >
          <Select
            mode="multiple"
            allowClear
            style={{ width: '100%' }}
            placeholder="请选择角色"
            options={getRoleOptions(props.roleList)}
          />
        </Form.Item>

        <Form.Item label="手机" name="mobile">
          <Input placeholder="请输入手机号码" />
        </Form.Item>

        <Form.Item label="电话" name="phone">
          <Input placeholder="请输入电话号码" />
        </Form.Item>

        <Form.Item label="邮箱" name="email">
          <Input placeholder="请输入邮箱地址: xxx@xxx.com" />
        </Form.Item>

        <Form.Item label="部门" name="department">
          <Input placeholder="请输入任职部门: 技术研发部/后台开发组" />
        </Form.Item>

        <Form.Item label="职位" name="position">
          <Input placeholder="请输入职位: 开发工程师" />
        </Form.Item>

        <Form.Item label="工位" name="location">
          <Input placeholder="请输入办公位: xx大楼SE-1231" />
        </Form.Item>

        <Form.Item wrapperCol={{ offset: layoutForm.labelCol.span }}>
          <Button type="primary" htmlType="submit">
            保存
          </Button>
        </Form.Item>
      </Form>
    </div>
  )
}

export default AccountFormUI
