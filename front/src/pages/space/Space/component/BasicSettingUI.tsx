import {
  SpaceTypeRadioOptions,
  SpaceVisitLevelRadioOptions
} from '@/pages/system/Space/component/ToolsUI'
import { SpaceInfoType, SpaceTypeTeam } from '@/types/spaceType'
import { QuestionCircleOutlined } from '@ant-design/icons'
import { Button, Input, Typography, Form, Radio, Tooltip } from 'antd'
import FormItem from 'antd/es/form/FormItem'

type SpaceBasicSettingUIProps = {
  spaceInfo?: SpaceInfoType
  onSaveSubmit?: (values: any) => void
}

const SpaceBasicSettingUI = (props: SpaceBasicSettingUIProps) => {
  const spaceTypeOptions = SpaceTypeRadioOptions()
  const spaceVistLevelOptions = SpaceVisitLevelRadioOptions()
  // spaceInfo 信息更新到表单
  const { spaceInfo } = props
  const [form] = Form.useForm()
  form.setFieldsValue(spaceInfo)

  return (
    <div style={{ paddingLeft: 6 }}>
      <Form layout="vertical" form={form} onFinish={props.onSaveSubmit}>
        <Form.Item label="空间ID" name="space_id" hidden>
          <Input placeholder="请输入空间ID" disabled />
        </Form.Item>
        <Form.Item
          label={<strong>空间Key</strong>}
          name="space_key"
          rules={[{ required: true, message: '请输入空间Key' }]}
          style={{ width: '50%' }}
        >
          <Input disabled />
        </Form.Item>
        <FormItem
          label={<strong>空间名称</strong>}
          name="name"
          rules={[{ required: true, message: '请输入空间名称' }]}
          style={{ width: '50%' }}
        >
          <Input placeholder="请输入空间名称" />
        </FormItem>
        <FormItem
          label={<strong>空间描述</strong>}
          name="description"
          style={{ width: '50%' }}
          rules={[{ required: true, message: '请输入空间描述' }]}
        >
          <Input.TextArea rows={3} placeholder="请输入空间描述" />
        </FormItem>
        <Form.Item
          label={<strong>空间类型</strong>}
          name="space_type"
          rules={[{ required: true, message: '请选择空间类型' }]}
          initialValue={1}
          style={{ width: '50%' }}
        >
          <Form.Item name="space_type" noStyle style={{ paddingBottom: 8 }}>
            <Radio.Group options={spaceTypeOptions} />
          </Form.Item>
          <Tooltip title="个人空间只能创建者为管理员，团队空间可以添加多个管理员">
            <Typography.Link>
              <QuestionCircleOutlined />
            </Typography.Link>
          </Tooltip>
        </Form.Item>
        <Form.Item
          label={<strong>访问级别</strong>}
          name="visit_level"
          rules={[{ required: true }]}
          initialValue={SpaceTypeTeam}
          style={{ width: '50%' }}
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
        <FormItem>
          <Button type="primary" htmlType="submit">
            保存
          </Button>
        </FormItem>
      </Form>
    </div>
  )
}

export default SpaceBasicSettingUI
