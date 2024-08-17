import { Button, Form, Input, Select } from 'antd'
import { EditLayoutForm, LayoutForm } from '@/config/layout'

interface ConfigFormUIProps {
  onSaveSubmit: (values: any) => void
}

/**
 * 配置表单 UI 组件
 * @param props 配置表单 UI 组件数据
 * @returns 配置表单 UI 组件
 */
const ConfigFormUI = (props: ConfigFormUIProps) => {
  const [form] = Form.useForm()
  // const isEdit = props.accountInfo ? true : false

  let layoutForm = LayoutForm

  return (
    <div className="panel-body">
      <Form {...layoutForm} name="basic" onFinish={props.onSaveSubmit} form={form}>
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

export default ConfigFormUI
