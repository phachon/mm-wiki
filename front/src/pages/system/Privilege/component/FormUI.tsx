import React, { useEffect } from 'react'
import {
  Button,
  Form,
  Input,
  Select,
  Radio,
  InputNumber,
  Tooltip,
  Typography,
  Switch,
  Cascader
} from 'antd'
import { EditLayoutForm, LayoutForm } from '@/config/layout'
import { QuestionCircleOutlined, InfoCircleOutlined } from '@ant-design/icons'
import * as icons from '@ant-design/icons'
import Icon from '@ant-design/icons'
import { PrivilegeListItemType, PrivilegeInfoType } from '@/types/privilegeType'
import { DefaultOptionType } from 'antd/lib/select'
import { PrivilegeRadioOptions } from './ToolsUI'

// icon 列表
const iconList = Object.keys(icons).filter((iconItem) => {
  const k = iconItem as keyof typeof icons
  if (k.indexOf('Outlined') > 0 && typeof icons[k] === 'object') {
    return true
  }
  return false
})

/**
 * 根据菜单权限下拉框 Options
 * @param privilegeList 权限列表
 * @return Option 下拉数据
 */
const getParentTypeOptions = (
  privilegeType: number,
  privilegeList?: PrivilegeListItemType[]
): DefaultOptionType[] => {
  const parentOptions: DefaultOptionType[] = []
  privilegeList?.forEach((privilege: PrivilegeListItemType) => {
    let parentOption: DefaultOptionType = {
      label: privilege.name,
      value: String(privilege.privilege_id),
      children: getParentTypeOptions(privilegeType, privilege.child_privileges)
    }
    parentOptions.push(parentOption)
  })
  return parentOptions
}

/**
 * 获取上级权限 Option 数据
 * @param privilegeType 权限类型
 * @returns
 */
const getParentOptions = (
  privilegeList: PrivilegeListItemType[],
  privilegeType: number
): DefaultOptionType[] => {
  let parentOptions: DefaultOptionType[] = []
  // 菜单的权限的上级只能是导航或菜单
  if (privilegeType === 2 || privilegeType === 3) {
    parentOptions = getParentTypeOptions(privilegeType, privilegeList)
  }
  return parentOptions
}

// PrivilegeFormUIProps 权限表单UI组件依赖
interface PrivilegeFormUIProps {
  /**
   * 接口标识列表
   */
  allApiMarks?: string[]

  /**
   * 上级权限列表
   */
  parentPrivileges: PrivilegeListItemType[]

  /**
   * 权限信息（修改）
   */
  privilegeInfo?: PrivilegeInfoType

  /**
   * 提交
   * @param values
   * @returns
   */
  onSaveSubmit: (values: any) => void

  formLayout?: {
    labelCol: { span: number }
    wrapperCol: { span: number }
  }
}

/**
 * 获取接口标识 Option 数据
 * @param allApiMarks 所有的接口标识列表
 * @returns
 */
const getApiMarksOptions = (allApiMarks: string[] | undefined): DefaultOptionType[] => {
  let apiMarksOptions: DefaultOptionType[] = []
  if (!allApiMarks || allApiMarks.length == 0) {
    return apiMarksOptions
  }
  allApiMarks.forEach((apiMark) => {
    apiMarksOptions.push({
      label: apiMark,
      value: apiMark
    })
  })
  return apiMarksOptions
}

/**
 * 权限表单 UI 组件
 * @param props 组件依赖数据
 * @returns 组件
 */
const PrivilegeFormUI = (props: PrivilegeFormUIProps) => {
  const [form] = Form.useForm()
  const isEdit = props.privilegeInfo ? true : false
  const privilegeTypeOptions = PrivilegeRadioOptions()

  let layoutForm = isEdit ? EditLayoutForm : LayoutForm
  if (props.formLayout) {
    layoutForm = props.formLayout
  }
  const parentPrivilege = props.parentPrivileges

  useEffect(() => {
    if (props.privilegeInfo) {
      form.setFieldsValue(props.privilegeInfo)
    }
  }, [props.privilegeInfo])

  /**
   * 权限类型点击操作
   * @param privilegeType 权限类型
   */
  const privilegeTypeOnChange = (e: any) => {
    form.setFieldsValue({
      parent_ids: ''
    })
  }

  return (
    <div className="panel-body">
      <Form {...layoutForm} name="privilege-form" form={form} onFinish={props.onSaveSubmit}>
        {isEdit && (
          <Form.Item label="权限ID" name="privilege_id" rules={[{ required: true }]}>
            <Input disabled placeholder="请输入权限ID" />
          </Form.Item>
        )}
        <Form.Item
          label="权限标识"
          name="identify"
          rules={[{ required: true, message: '请输入权限标识!' }]}
        >
          <Input
            placeholder="请输入权限标识(英文字符)"
            disabled={props.privilegeInfo ? true : false}
          />
        </Form.Item>

        <Form.Item
          label="权限名称"
          name="name"
          rules={[{ required: true, message: '请输入权限名称!' }]}
        >
          <Input placeholder="请输入权限名称" />
        </Form.Item>

        <Form.Item
          label="权限类型"
          name="privilege_type"
          rules={[{ required: true }]}
          initialValue={1}
        >
          <Form.Item name="privilege_type" noStyle>
            <Radio.Group
              options={privilegeTypeOptions}
              defaultValue={1}
              onChange={privilegeTypeOnChange}
            />
          </Form.Item>
          <Tooltip title="操作权限：对资源有操作行为的权限（例如：编辑/删除）">
            <Typography.Link>
              <QuestionCircleOutlined />
            </Typography.Link>
          </Tooltip>
        </Form.Item>

        {/* 只有导航不需要上级权限 */}
        <Form.Item dependencies={['privilege_type']} noStyle>
          {({ getFieldValue }) => {
            const privilegeType = getFieldValue('privilege_type')
            if (privilegeType === 1) {
              return null
            }
            return (
              <Form.Item
                label="上级权限"
                name="parent_ids"
                rules={[{ required: true, message: '请选择上级权限!' }]}
                getValueProps={(val) => {
                  let value = val ? val.split(',') : []
                  return { value: value }
                }}
                getValueFromEvent={(values: string[]) => {
                  return values.length > 0 ? values.toString() : ''
                }}
                initialValue=""
              >
                <Cascader
                  key={'Cascader' + privilegeType}
                  options={getParentOptions(parentPrivilege, privilegeType)}
                  changeOnSelect={privilegeType === 2 ? true : false}
                  allowClear={true}
                  placeholder="请选择上级权限"
                />
              </Form.Item>
            )
          }}
        </Form.Item>

        {/* 操作类型权限不需要页面路由 */}
        <Form.Item dependencies={['privilege_type']} noStyle>
          {({ getFieldValue }) => {
            const privilegeType = getFieldValue('privilege_type')
            if (privilegeType === 3) {
              return null
            }
            return (
              <Form.Item label="页面路由" name="page_router" initialValue={''}>
                <Input placeholder="页面跳转路由：/user/info" />
              </Form.Item>
            )
          }}
        </Form.Item>

        <Form.Item
          label="接口标识"
          name="api_marks"
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
            placeholder="请选择关联的接口标识"
            options={getApiMarksOptions(props.allApiMarks)}
            showSearch
          />
        </Form.Item>

        {/* 操作类型权限不需要 icon 图标/是否外显/排序数字 */}
        <Form.Item dependencies={['privilege_type']} noStyle>
          {({ getFieldValue }) => {
            const privilegeType = getFieldValue('privilege_type')
            if (privilegeType === 3) {
              return null
            }
            return (
              <>
                <Form.Item label="Icon图标" name="icon">
                  <Select
                    showSearch
                    allowClear={true}
                    style={{ width: '100%' }}
                    placeholder="请选择 icon 图标"
                  >
                    {iconList.map((iconItem) => {
                      const k = iconItem as keyof typeof icons
                      return (
                        <Select.Option value={iconItem} key={iconItem}>
                          <Icon
                            component={icons[k] as React.ForwardRefExoticComponent<any>}
                            style={{ marginRight: '8px' }}
                          />
                          {iconItem}
                        </Select.Option>
                      )
                    })}
                  </Select>
                </Form.Item>

                <Form.Item
                  label="是否外显"
                  name="is_display"
                  valuePropName="checked"
                  initialValue={1}
                  getValueProps={(value) => ({ checked: value === 1 ? true : false })}
                  getValueFromEvent={(value) => {
                    return value ? 1 : 0
                  }}
                >
                  <Switch checkedChildren="是" unCheckedChildren="否" defaultChecked />
                </Form.Item>

                <Form.Item
                  label="排序数字"
                  name="sequence"
                  rules={[{ type: 'number', min: 1, max: 1000 }]}
                >
                  <InputNumber
                    width={1000}
                    defaultValue={1}
                    style={{ width: '100%' }}
                    placeholder="排序数字越小显示越靠前"
                    addonAfter={
                      <Tooltip title="排序数字越小显示越靠前">
                        <QuestionCircleOutlined style={{ color: 'rgba(0,0,0,.45)' }} />
                      </Tooltip>
                    }
                  />
                </Form.Item>
              </>
            )
          }}
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

export default PrivilegeFormUI
