import { AccountInfoType } from '@/types/accountType'
import { Avatar, Button, Card, Col, Divider, Form, Input, List, Row, Space, Tabs } from 'antd'
import { useEffect } from 'react'
import { LayoutForm, LayoutFormButton } from '@/config/layout'
import { FormOutlined, InfoCircleOutlined } from '@ant-design/icons'

interface ProfileSettingUIProps {
  accountInfo?: AccountInfoType
  onEditSubmit: (values: any) => void
}

/**
 * 基础设置 UI 组件
 * @param props
 * @returns
 */
const ProfileBasicSettingUI = (props: ProfileSettingUIProps) => {
  const [form] = Form.useForm()

  useEffect(() => {
    form.setFieldsValue(props.accountInfo)
  }, [props.accountInfo])

  return (
    <div className="panel-body">
      <Form layout="vertical" name="basic_setting" form={form} onFinish={props.onEditSubmit}>
        <Row gutter={24}>
          <Col span={6} key={'account_id'}>
            <Form.Item label="账号ID" name="account_id" rules={[{ required: true }]}>
              <Input disabled />
            </Form.Item>
          </Col>
          <Col span={6} key={'name'} offset={1}>
            <Form.Item
              label="账号名"
              name="name"
              rules={[{ required: true, message: '请输入账号名!' }]}
            >
              <Input disabled placeholder="请输入账号名" />
            </Form.Item>
          </Col>
          <Col span={7} key={'given_name'} offset={1}>
            <Form.Item
              label="昵称"
              name="given_name"
              rules={[{ required: true, message: '请输入昵称!' }]}
            >
              <Input placeholder="请输入昵称" />
            </Form.Item>
          </Col>

          <Col span={6} key={'email'}>
            <Form.Item label="邮箱" name="email">
              <Input placeholder="请输入邮箱地址：xxx@xxx.com" />
            </Form.Item>
          </Col>

          <Col span={6} key={'phone'} offset={1}>
            <Form.Item label="电话" name="phone">
              <Input placeholder="请输入电话号码" />
            </Form.Item>
          </Col>

          <Col span={7} key={'mobile'} offset={1}>
            <Form.Item label="手机" name="mobile">
              <Input placeholder="请输入手机号码" />
            </Form.Item>
          </Col>

          <Col span={6} key={'department'}>
            <Form.Item label="部门" name="department">
              <Input placeholder="请输入任职部门: 技术研发部/后台开发组" />
            </Form.Item>
          </Col>

          <Col span={6} key={'position'} offset={1}>
            <Form.Item label="职位" name="position">
              <Input placeholder="请输入职位: 开发工程师" />
            </Form.Item>
          </Col>

          <Col span={7} key={'location'} offset={1}>
            <Form.Item label="工位" name="location">
              <Input placeholder="请输入办公位: xx大楼SE-1231" />
            </Form.Item>
          </Col>
        </Row>
        <Divider />
        <Row gutter={24} justify={'center'}>
          <Form.Item>
            <Button type="primary" htmlType="submit">
              保存
            </Button>
          </Form.Item>
        </Row>
      </Form>
    </div>
  )
}

export default ProfileBasicSettingUI
