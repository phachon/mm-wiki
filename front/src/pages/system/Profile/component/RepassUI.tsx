import { Button, Divider, Form, Input } from 'antd'
import { LayoutForm, LayoutFormButton } from '@/config/layout'

interface ProfileRepassUIProps {
  onSaveSubmit: (values: any) => void
}

/**
 * 修改密码 UI 组件
 * @param props
 * @returns
 */
const ProfileRepassUI = (props: ProfileRepassUIProps) => {
  const [form] = Form.useForm()
  return (
    <div className="panel-body">
      <Form
        {...LayoutForm}
        labelCol={{ span: 2 }}
        name="repass-form"
        form={form}
        onFinish={props.onSaveSubmit}
      >
        <Form.Item
          label="旧密码"
          name="old_pwd"
          rules={[{ required: true, message: '请输入旧密码!' }]}
        >
          <Input.Password placeholder="请输入旧密码" />
        </Form.Item>

        <Form.Item
          label="新密码"
          name="new_pwd"
          rules={[{ required: true, message: '请输入新密码!' }]}
        >
          <Input.Password placeholder="请输入新密码" />
        </Form.Item>

        <Form.Item
          label="确认密码"
          name="confirm_pwd"
          rules={[{ required: true, message: '请再次输入新密码!' }]}
        >
          <Input.Password placeholder="请再次输入新密码" />
        </Form.Item>
        <Divider />
        <Form.Item {...LayoutFormButton} wrapperCol={{ offset: 2 }}>
          <Button type="primary" htmlType="submit">
            保存
          </Button>
        </Form.Item>
      </Form>
    </div>
  )
}

export default ProfileRepassUI
