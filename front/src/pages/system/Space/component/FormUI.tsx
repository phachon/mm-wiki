import { Button, Form, Input, Modal, Radio, Select, Space, Tooltip, Typography } from 'antd'
import { EditLayoutForm, LayoutForm } from '@/config/layout'
import { SpaceInfoType } from '@/types/spaceType'
import { QuestionCircleOutlined } from '@ant-design/icons'
import { useEffect, useState } from 'react'
import { PrivilegeListItemType } from '@/types/privilegeType'
import { AccountInfoType } from '@/types/accountType'
import { SpaceIsExportYes, SpaceIsShareYes, SpaceTypeTeam } from '@/types/spaceType'
import { DepartmentInfoType } from '@/types/departmentType'
import TextArea from 'antd/es/input/TextArea'
import {
  SpaceIsExportRadioOptions,
  SpaceIsShareRadioOptions,
  SpaceTypeRadioOptions,
  SpaceVisitLevelRadioOptions
} from './ToolsUI'

interface SpaceFormUIProps {
  departmentList?: DepartmentInfoType[] // 部门列表
  accountList?: AccountInfoType[] // 账号列表
  spaceInfo?: SpaceInfoType // 空间信息
  formLayout?: {
    labelCol: { span: number }
    wrapperCol: { span: number }
  }
  privilegeList?: PrivilegeListItemType[] // 权限列表

  /**
   * 空间保存方法
   * @param values 空间保存方法
   */
  onSaveSubmit: (values: any) => void
}

/**
 * 空间表单 UI 组件
 * @param props
 * @returns
 */
const SpaceFormUI = (props: SpaceFormUIProps) => {
  const [form] = Form.useForm()

  const spaceVistLevelOptions = SpaceVisitLevelRadioOptions()
  const spaceIsShareOptions = SpaceIsShareRadioOptions()
  const spaceIsExportOptions = SpaceIsExportRadioOptions()
  const spaceTypeOptions = SpaceTypeRadioOptions()

  const isEdit = props.spaceInfo ? true : false
  let layoutForm = isEdit ? EditLayoutForm : LayoutForm
  if (props.formLayout) {
    layoutForm = props.formLayout
  }

  useEffect(() => {
    if (props.spaceInfo) {
      form.setFieldsValue(props.spaceInfo)
    }
  }, [props.spaceInfo])

  return (
    <div className="panel-body">
      <Form name="space-form" {...layoutForm} form={form} onFinish={props.onSaveSubmit}>
        {isEdit && (
          <Form.Item label="空间ID" name="space_id" hidden>
            <Input placeholder="请输入空间ID" disabled />
          </Form.Item>
        )}

        <Form.Item
          label="空间 Key"
          name="space_key"
          rules={[{ required: true, message: '请输入空间Key' }]}
        >
          <Input placeholder="请输入空间唯一key（数字、英文组成）" disabled={isEdit} />
        </Form.Item>

        <Form.Item
          label="空间名称"
          name="name"
          rules={[{ required: true, message: '请输入空间名称' }]}
        >
          <Input placeholder="请输入空间名称" />
        </Form.Item>

        <Form.Item
          label="空间描述"
          name="description"
          rules={[{ required: true, message: '请输入空间描述' }]}
        >
          <TextArea rows={5} placeholder="请输入空间描述" />
        </Form.Item>

        {!isEdit && (
          <Form.Item
            label="空间类型"
            name="space_type"
            rules={[{ required: true, message: '请选择空间类型' }]}
            initialValue={SpaceTypeTeam}
          >
            <Form.Item name="space_type" noStyle>
              <Radio.Group options={spaceTypeOptions} disabled />
            </Form.Item>
            <Tooltip title="个人空间只能创建者为管理员，团队空间可以添加多个管理员">
              <Typography.Link>
                <QuestionCircleOutlined />
              </Typography.Link>
            </Tooltip>
          </Form.Item>
        )}

        {!isEdit && (
          <Form.Item
            label="管 理 员"
            name="admin_account_ids"
            rules={[{ required: true, message: '请选择空间管理员' }]}
          >
            <Select
              mode="multiple"
              placeholder="请选择空间管理员"
              showSearch={true}
              filterOption={(input: any, option: any) => {
                const optionText = option.children.toLowerCase()
                const optionKey = option.key.toLowerCase()
                return (
                  optionText.indexOf(input.toLowerCase()) >= 0 ||
                  optionKey.indexOf(input.toLowerCase()) >= 0
                )
              }}
            >
              {props.accountList &&
                props.accountList.map((accountInfo) => (
                  <Select.Option
                    id={accountInfo.account_id}
                    key={accountInfo.name}
                    value={accountInfo.account_id}
                  >
                    {accountInfo.name + '(' + accountInfo.given_name + ')'}
                  </Select.Option>
                ))}
            </Select>
          </Form.Item>
        )}

        <Form.Item
          label="访问级别"
          name="visit_level"
          rules={[{ required: true }]}
          initialValue={0}
        >
          <Form.Item name="visit_level" noStyle>
            <Radio.Group options={spaceVistLevelOptions} />
          </Form.Item>
          <Tooltip title="公开空间默认所有账号可访问，私有空间只允许有权限账号访问">
            <Typography.Link>
              <QuestionCircleOutlined />
            </Typography.Link>
          </Tooltip>
        </Form.Item>
        <Form.Item
          label="是否分享"
          name="is_share"
          rules={[{ required: true }]}
          initialValue={SpaceIsShareYes}
        >
          <Form.Item name="is_share" noStyle>
            <Radio.Group options={spaceIsShareOptions} />
          </Form.Item>
          <Tooltip title="该空间下的所有文档是否可分享">
            <Typography.Link>
              <QuestionCircleOutlined />
            </Typography.Link>
          </Tooltip>
        </Form.Item>
        <Form.Item
          label="是否导出"
          name="is_export"
          rules={[{ required: true }]}
          initialValue={SpaceIsExportYes}
        >
          <Form.Item name="is_export" noStyle>
            <Radio.Group options={spaceIsExportOptions} />
          </Form.Item>
          <Tooltip title="该空间下的所有文档是否可导出">
            <Typography.Link>
              <QuestionCircleOutlined />
            </Typography.Link>
          </Tooltip>
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

export default SpaceFormUI
