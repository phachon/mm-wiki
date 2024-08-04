import { Menu, Dropdown, Layout, Space, MenuProps, Drawer, message, List, Button } from 'antd'
import {
  UserOutlined,
  LoginOutlined,
  ProfileOutlined,
  LockOutlined,
  SoundOutlined,
  CloseOutlined,
  HomeOutlined,
  AppstoreOutlined,
  CaretDownOutlined,
  QuestionCircleOutlined,
  ReadOutlined,
  SendOutlined,
  SettingOutlined
} from '@ant-design/icons'
import { Link } from 'react-router-dom'
import { useState } from 'react'
import { SettingConfig } from '@/config/setting'
import { NoticeInfoType } from '@/types/noticeType'
import DynamicIcon from '../DynamicIcon/DynamicIcon'
import { AccountInfoType } from '@/types/accountType'
import {
  LayoutHeaderHomeKey,
  LayoutHeaderSpaceKey,
  LayoutHeaderSystemKey,
  LayoutHeaderUserKey
} from './types'

type MenuItem = Required<MenuProps>['items'][number]

const noticePageSize: number = 4

interface LayoutHeaderProps {
  accountInfo?: AccountInfoType | undefined
  navItems: MenuItem[]
  navSelectedKeys: string[]
  noticeList: NoticeInfoType[]
  noticeTotal: number
  onNavSelectChange: (navKey: string) => void
  onNoticeListClick: (pageSize: number, pageNum: number) => void
  onNoticeChange: (page: number) => void
  onLogoutClick: () => void
}

/**
 * 导航布局
 * @param props
 * @returns
 */
const LayoutHeader = (props: LayoutHeaderProps) => {
  const [noticeOpen, setNoticeOpen] = useState(false)
  const [noticeCurrentPage, setNoticeCurrentPage] = useState(1)
  const defNavItems: MenuItem[] = [
    {
      label: <a href="/">主页</a>,
      key: LayoutHeaderHomeKey,
      icon: <HomeOutlined />
    },
    {
      label: <a href="/space">空间</a>,
      key: LayoutHeaderSpaceKey,
      icon: <AppstoreOutlined />
    },
    {
      label: <a href="/user">用户</a>,
      key: LayoutHeaderUserKey,
      icon: <UserOutlined />
    },
    {
      label: <a href="/system">系统</a>,
      key: LayoutHeaderSystemKey,
      icon: <SettingOutlined />
    }
  ]

  // 退出登录
  const menuItems: MenuProps['items'] = [
    {
      label: <Link to="/system/profile/info">个人信息</Link>,
      key: 'profile_info',
      icon: <ProfileOutlined />
    },
    {
      label: <Link to="/system/profile/repass">修改密码</Link>,
      key: 'profile_repass',
      icon: <LockOutlined />
    },
    {
      type: 'divider'
    },
    {
      label: <span onClick={props.onLogoutClick}>退出登录</span>,
      key: 'profile_logout',
      icon: <LoginOutlined />
    }
  ]

  // 帮助
  const helpItems: MenuProps['items'] = [
    {
      label: <Link to="/system/profile/info">问题反馈</Link>,
      key: 'profile_info',
      icon: <QuestionCircleOutlined />
    },
    {
      label: <Link to="/system/profile/repass">在线文档</Link>,
      key: 'profile_repass',
      icon: <ReadOutlined />
    },
    {
      label: <Link to="/system/profile/repass">关于软件</Link>,
      key: 'profile_repass',
      icon: <SendOutlined />
    }
  ]

  // 公告
  const noticeDrawer = () => {
    return (
      <Drawer
        title="公告列表"
        placement="right"
        onClose={() => {
          setNoticeOpen(false)
        }}
        open={noticeOpen}
        style={{ padding: 12, paddingTop: 0 }}
        closable={false}
        extra={
          <Space>
            <Button
              type="text"
              size="small"
              onClick={() => {
                setNoticeOpen(false)
              }}
            >
              <CloseOutlined />
            </Button>
          </Space>
        }
      >
        <List
          itemLayout="vertical"
          pagination={{
            onChange: props.onNoticeChange,
            size: 'default',
            current: noticeCurrentPage,
            pageSize: noticePageSize,
            total: props.noticeTotal,
            showSizeChanger: false,
            showLessItems: true,
            hideOnSinglePage: true
          }}
          dataSource={props.noticeList}
          renderItem={(item: NoticeInfoType) => (
            <List.Item
              style={{ padding: 16 }}
              key={item.title}
              actions={[
                <Space>
                  <DynamicIcon name="UserOutlined" />
                  {item.account_name}
                </Space>,
                <Space>
                  <DynamicIcon name="FieldTimeOutlined" />
                  {item.update_time}
                </Space>
              ]}
            >
              <List.Item.Meta title={item.title} style={{ marginBottom: 0 }} />
              {item.content}
            </List.Item>
          )}
        />
      </Drawer>
    )
  }

  return (
    <Layout.Header className="admin-header">
      <div className="admin-header-logo">
        <a className="admin-header-link" href="/">
          {SettingConfig.frameHeaderIconUrl ? (
            <img src={SettingConfig.frameHeaderIconUrl} alt="logo"></img>
          ) : (
            ''
          )}
          <span className="admin-header-text">{SettingConfig.frameHeaderTitle}</span>
        </a>
      </div>
      <div className="admin-header-nav">
        <div className="admin-header-left">
          <Menu
            theme="dark"
            mode="horizontal"
            selectedKeys={props.navSelectedKeys}
            items={defNavItems}
            onSelect={({ key, keyPath, selectedKeys, domEvent }) => {
              props.onNavSelectChange(key)
            }}
          />
        </div>
        <div className="admin-header-right">
          <span className="admin-header-action">
            <a
              onClick={() => {
                props.onNoticeListClick(noticePageSize, 1)
                setNoticeOpen(true)
              }}
            >
              <SoundOutlined /> 公告 <CaretDownOutlined />
            </a>
            {noticeDrawer()}
          </span>
          <Dropdown menu={{ items: helpItems }}>
            <span className="admin-header-action">
              <Space>
                <QuestionCircleOutlined />
                帮助
                <CaretDownOutlined />
              </Space>
            </span>
          </Dropdown>
          <Dropdown menu={{ items: menuItems }}>
            <span className="admin-header-action">
              <Space>
                <UserOutlined />
                {props.accountInfo?.name}
                <CaretDownOutlined />
              </Space>
            </span>
          </Dropdown>
        </div>
      </div>
    </Layout.Header>
  )
}

export default LayoutHeader
