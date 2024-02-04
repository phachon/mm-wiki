import { Menu, Dropdown, Layout, Space, MenuProps, Drawer, message, List, Button } from 'antd'
import {
  UserOutlined,
  LoginOutlined,
  ProfileOutlined,
  LockOutlined,
  SoundOutlined,
  FieldTimeOutlined,
  CloseOutlined
} from '@ant-design/icons'
import { Link, useNavigate } from 'react-router-dom'
import { useState } from 'react'
import { useGlobalStore } from '@/stores/index'
import { SettingConfig } from '@/config/setting'
import React from 'react'
import { NoticeInfoType } from '@/types/noticeType'
import { getNavItems } from './ToolsUI'
import { HOME_ROOT_PATH } from '@/router/type'

const IconText = ({ icon, text }: { icon: React.FC; text: string }) => (
  <Space>
    {React.createElement(icon)}
    {text}
  </Space>
)

const noticePageSize: number = 4

const FrameHeaderUI = () => {
  const {
    accountInfo,
    logoutCallback,
    iPrivilegeData,
    iNavSelectedKey,
    onNavSelectChange,
    noticeList,
    noticeListCallback,
    noticeTotal
  } = useGlobalStore()

  const [noticeOpen, setNoticeOpen] = useState(false)
  const [noticeCurrentPage, setNoticeCurrentPage] = useState(1)
  const navigate = useNavigate()

  const navItems = getNavItems(iPrivilegeData?.iNavItems)

  // 退出登录
  const menuItems: MenuProps['items'] = [
    {
      label: <Link to="/profile/info">个人信息</Link>,
      key: 'profile_info',
      icon: <ProfileOutlined />
    },
    {
      label: <Link to="/profile/repass">修改密码</Link>,
      key: 'profile_repass',
      icon: <LockOutlined />
    },
    {
      type: 'divider'
    },
    {
      label: (
        <span
          onClick={() => {
            {
              message.success('退出登录成功', 1.5, () => {
                logoutCallback()
                navigate(HOME_ROOT_PATH)
              })
            }
          }}
        >
          退出登录
        </span>
      ),
      key: 'profile_logout',
      icon: <LoginOutlined />
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
        bodyStyle={{ padding: 12, paddingTop: 0 }}
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
            onChange: (page) => {
              setNoticeCurrentPage(page)
              noticeListCallback(noticePageSize, page)
            },
            size: 'default',
            current: noticeCurrentPage,
            pageSize: noticePageSize,
            total: noticeTotal,
            showSizeChanger: false,
            showLessItems: true,
            hideOnSinglePage: true
          }}
          dataSource={noticeList}
          renderItem={(item: NoticeInfoType) => (
            <List.Item
              style={{ padding: 16 }}
              key={item.title}
              actions={[
                <IconText icon={UserOutlined} text={item.account_name} />,
                <IconText icon={FieldTimeOutlined} text={item.update_time} />
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
            selectedKeys={[iNavSelectedKey]}
            items={navItems}
            onSelect={({ key, keyPath, selectedKeys, domEvent }) => {
              onNavSelectChange(key)
            }}
          />
          {/* {navItems &&
              navItems.map((item: INavItem) => {
                return (
                  <Menu.SubMenu
                    icon={<DynamicIcon name={item.icon} />}
                    key={item.key}
                    title={item.label}
                    onTitleClick={() => navSelectCallback(item)}
                  ></Menu.SubMenu>
                )
              })} */}
          {/* </Menu> */}
        </div>
        <div className="admin-header-right">
          <span className="admin-header-action">
            <a
              onClick={() => {
                noticeListCallback(noticePageSize, 1)
                setNoticeOpen(true)
              }}
            >
              <SoundOutlined /> 公告
            </a>
            {noticeDrawer()}
          </span>
          <Dropdown menu={{ items: menuItems }}>
            <span className="admin-header-action">
              <Space>
                <UserOutlined />
                {accountInfo?.name}
              </Space>
            </span>
          </Dropdown>
        </div>
      </div>
    </Layout.Header>
  )
}

export default FrameHeaderUI
