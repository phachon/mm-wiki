import { AccountInfoType } from '@/types/accountType'
import { Avatar, Button, Card, Col, Divider, Form, Input, List, Row, Tabs } from 'antd'
import { useEffect } from 'react'
import TabPane from 'antd/es/tabs/TabPane'
import {
  UserOutlined,
  EnvironmentOutlined,
  IdcardOutlined,
  PhoneOutlined,
  MailOutlined,
  ClusterOutlined,
  MobileOutlined
} from '@ant-design/icons'

interface ProfileInfoUIProps {
  accountInfo?: AccountInfoType
  onFinishCallback: (values: any) => void
}

/**
 * 个人信息 UI 组件
 * @param props
 * @returns
 */
const ProfileInfoUI = (props: ProfileInfoUIProps) => {
  const [form] = Form.useForm()

  useEffect(() => {
    form.setFieldsValue(props.accountInfo)
  }, [props.accountInfo])

  return (
    <div className="panel-body">
      <Row gutter={24}>
        <Col span={6}>
          <Card style={{ marginBottom: 24 }}>
            <div style={{ textAlign: 'center' }}>
              <Avatar
                size={104}
                src={
                  'https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png'
                }
              />
              <h2 style={{ marginTop: 16, marginBottom: 6 }}>{props.accountInfo?.name}</h2>
              <p style={{ margin: 0 }}>昵称：{props.accountInfo?.given_name}</p>
            </div>
            <Divider />
            <div style={{ textAlign: 'left' }}>
              <p>
                <ClusterOutlined /> 部门：{props.accountInfo?.department_id}
              </p>
              <p>
                <IdcardOutlined /> 职位：{props.accountInfo?.position}
              </p>
              <p>
                <EnvironmentOutlined /> 工位：{props.accountInfo?.location}
              </p>
            </div>
            <Divider />
            <div style={{ textAlign: 'left' }}>
              <p>
                <MobileOutlined /> 手机：{props.accountInfo?.mobile}
              </p>
              <p>
                <PhoneOutlined /> 电话：{props.accountInfo?.phone}
              </p>
              <p>
                <MailOutlined /> 邮箱：{props.accountInfo?.email}
              </p>
            </div>
          </Card>
        </Col>
        <Col span={18} style={{ paddingLeft: 4 }}>
          <Card>
            <Tabs defaultActiveKey="1">
              <TabPane tab="动态" key="1">
                <List
                  itemLayout="horizontal"
                  dataSource={[
                    {
                      title: '动态标题',
                      description: '动态描述'
                    },
                    {
                      title: '动态标题',
                      description: '动态描述'
                    }
                  ]}
                  renderItem={(item) => (
                    <List.Item>
                      <List.Item.Meta
                        avatar={<Avatar icon={<UserOutlined />} />}
                        title={item.title}
                        description={item.description}
                      />
                    </List.Item>
                  )}
                />
              </TabPane>
              <TabPane tab="空间" key="2">
                <List
                  itemLayout="horizontal"
                  dataSource={[
                    {
                      title: '项目名称',
                      description: '项目描述'
                    }
                  ]}
                  renderItem={(item) => (
                    <List.Item>
                      <List.Item.Meta
                        avatar={<Avatar icon={<UserOutlined />} />}
                        title={item.title}
                        description={item.description}
                      />
                    </List.Item>
                  )}
                />
              </TabPane>
              <TabPane tab="关注" key="3">
                <List
                  itemLayout="horizontal"
                  dataSource={[
                    {
                      title: '团队名称',
                      description: '团队描述'
                    }
                  ]}
                  renderItem={(item) => (
                    <List.Item>
                      <List.Item.Meta
                        avatar={<Avatar icon={<UserOutlined />} />}
                        title={item.title}
                        description={item.description}
                      />
                    </List.Item>
                  )}
                />
              </TabPane>
            </Tabs>
          </Card>
        </Col>
      </Row>
    </div>
  )
}

export default ProfileInfoUI
