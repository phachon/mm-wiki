import {
  UnorderedListOutlined,
  LinkOutlined,
  HddOutlined,
  ContactsOutlined,
  ReadOutlined,
  DesktopOutlined,
  ChromeOutlined,
  CodeOutlined
} from '@ant-design/icons'
import {
  Alert,
  Avatar,
  Card,
  Col,
  Divider,
  List,
  Row,
  Space,
  Table,
  Timeline,
  Typography
} from 'antd'
import Marquee from 'react-fast-marquee'
const { Text, Link } = Typography

interface MainIndexUIProps {
  noticeMsg?: string
}

const dataSource = [
  {
    key: '1',
    role: '技术负责',
    name: '张三',
    email: 'zhangsan@ailu-admin.com',
    phone: '123456789'
  },
  {
    key: '2',
    role: '权限审批',
    name: '李四',
    email: 'lisi@ailu-admin.com',
    phone: '123456789'
  },
  {
    key: '3',
    role: '销售管理',
    name: '老王',
    email: 'wang5@ailu-admin.com',
    phone: '123456789'
  },
  {
    key: '4',
    role: 'xxxxxx',
    name: 'xxxxxx',
    email: 'xxxxxx@ailu-admin.com',
    phone: '123456789'
  }
]

const columns = [
  {
    title: '角色',
    dataIndex: 'role',
    key: 'role'
  },
  {
    title: '姓名',
    dataIndex: 'name',
    key: 'name'
  },
  {
    title: '邮箱',
    dataIndex: 'email',
    key: 'email'
  },
  {
    title: '电话',
    dataIndex: 'phone',
    key: 'phone'
  }
]

const gridStyle: React.CSSProperties = {
  width: '25%',
  textAlign: 'center'
}

const cardStyle: React.CSSProperties = {
  marginBottom: 16
}

/**
 * 默认主页 UI 组件
 * @param props
 * @returns
 */
const MainIndexUI = (props: MainIndexUIProps) => {
  return (
    <div>
      <div className="panel-body">
        <Alert
          showIcon
          type="info"
          message={
            <Marquee pauseOnHover gradient={false} speed={60}>
              公告：xxxxxxxxxxxx，xxxxxxxxxxxx
            </Marquee>
          }
        />
      </div>
      <div className="panel-body">
        <Row gutter={16}>
          <Col span={16}>
            <Card
              title={
                <Space>
                  <UnorderedListOutlined />
                  最近动态
                </Space>
              }
              bordered={true}
              style={cardStyle}
              type="inner"
            >
              <Timeline style={{ marginTop: 16 }}>
                <Timeline.Item>
                  <Text strong style={{ color: '#1677ff' }}>
                    2015-09-01：
                  </Text>
                  创建用户 ID 成功
                </Timeline.Item>
                <Timeline.Item>
                  <Text strong style={{ color: '#1677ff' }}>
                    2015-09-01：
                  </Text>
                  ：创建用户 ID 成功
                </Timeline.Item>
                <Timeline.Item>
                  <Text strong style={{ color: '#1677ff' }}>
                    2015-09-01：
                  </Text>
                  ：创建用户 ID 成功
                </Timeline.Item>
                <Timeline.Item>
                  <Text strong style={{ color: '#1677ff' }}>
                    2015-09-01：
                  </Text>
                  ：创建用户 ID 成功
                </Timeline.Item>
                <Timeline.Item>
                  <Text strong style={{ color: '#1677ff' }}>
                    2015-09-01：
                  </Text>
                  ：创建用户 ID 成功
                </Timeline.Item>
                <Timeline.Item>
                  <Text strong style={{ color: '#1677ff' }}>
                    2015-09-01：
                  </Text>
                  ：创建用户 ID 成功
                </Timeline.Item>
                <Timeline.Item>
                  <Text strong style={{ color: '#1677ff' }}>
                    2015-09-01：
                  </Text>
                  ：创建用户 ID 成功
                </Timeline.Item>
                <Timeline.Item>
                  <Text strong style={{ color: '#1677ff' }}>
                    2015-09-01：
                  </Text>
                  ：创建用户 ID 成功
                </Timeline.Item>
                <Timeline.Item style={{ paddingBottom: 0 }}>
                  <Text strong style={{ color: '#1677ff' }}>
                    2015-09-01：
                  </Text>
                  ：创建用户 ID 成功xxx
                </Timeline.Item>
              </Timeline>
            </Card>
            <Card
              title={
                <Space>
                  <ContactsOutlined />
                  联系方式
                </Space>
              }
              bordered={true}
              type="inner"
              style={cardStyle}
              bodyStyle={{ padding: 16 }}
            >
              <Table dataSource={dataSource} columns={columns} pagination={false} />
            </Card>
          </Col>
          <Col span={8}>
            <Card
              title={
                <Space>
                  <LinkOutlined />
                  快捷导航
                </Space>
              }
              style={cardStyle}
              type="inner"
            >
              <Row>
                <Col span={6} className="text-center" style={{ marginBottom: 12 }}>
                  <a href="">百度</a>
                </Col>
                <Col span={6} className="text-center">
                  <a href="">腾讯网</a>
                </Col>
                <Col span={6} className="text-center">
                  <a href="">必应</a>
                </Col>
                <Col span={6} className="text-center">
                  <a href="">链接1</a>
                </Col>
                <Col span={6} className="text-center">
                  <a href="">链接1</a>
                </Col>
                <Col span={6} className="text-center">
                  <a href="">链接1</a>
                </Col>
                <Col span={6} className="text-center">
                  <a href="">链接1</a>
                </Col>
                <Col span={6} className="text-center">
                  <a href="">链接1</a>
                </Col>
              </Row>
            </Card>
            <Card
              title={
                <Space>
                  <ReadOutlined />
                  帮助文档
                </Space>
              }
              style={cardStyle}
              type="inner"
            >
              <Card.Grid style={gridStyle}>
                <a href="">文档1</a>
              </Card.Grid>
              <Card.Grid style={gridStyle}>
                <a href="">文档2</a>
              </Card.Grid>
              <Card.Grid style={gridStyle}>
                <a href="">文档3</a>
              </Card.Grid>
              <Card.Grid style={gridStyle}>
                <a href="">文档4</a>
              </Card.Grid>
              <Card.Grid style={gridStyle}>
                <a href="">文档5</a>
              </Card.Grid>
              <Card.Grid style={gridStyle}>
                <a href="">文档6</a>
              </Card.Grid>
              <Card.Grid style={gridStyle}>
                <a href="">文档7</a>
              </Card.Grid>
            </Card>
            <Card
              title={
                <Space>
                  <CodeOutlined />
                  系统信息
                </Space>
              }
              style={cardStyle}
              type="inner"
            >
              <List>
                <List.Item>
                  <Space>
                    <DesktopOutlined />
                    系统版本号：V1.10
                  </Space>
                </List.Item>
                <List.Item>
                  <Space>
                    <HddOutlined />
                    服务版本号：V1.11
                  </Space>
                </List.Item>
                <List.Item>
                  <Space>
                    <ChromeOutlined />
                    当前浏览器：Chrome
                  </Space>
                </List.Item>
              </List>
            </Card>
          </Col>
        </Row>
      </div>
    </div>
  )
}

export default MainIndexUI
