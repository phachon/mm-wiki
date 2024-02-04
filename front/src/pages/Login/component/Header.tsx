import React from 'react'
import { Menu, Space, Layout } from 'antd'
import { UserOutlined, SoundOutlined, QuestionCircleOutlined } from '@ant-design/icons'
import { SettingConfig } from '@/config/setting'

const LoginHeaderUI = () => {
  return (
    <Layout.Header className="login-header">
      <div className="login-header-logo">
        <a className="login-header-link" href="#">
          {SettingConfig.frameHeaderIconUrl ? (
            <img src={SettingConfig.frameHeaderIconUrl} alt="logo"></img>
          ) : (
            ''
          )}
          <span className="login-header-text">{SettingConfig.frameHeaderTitle}</span>
        </a>
      </div>
      <div className="login-header-nav">
        <div className="login-header-left">
          <Menu theme="dark" mode="horizontal" selectedKeys={[]} items={[]} />
        </div>
        <div className="login-header-right">
          <span className="login-header-action">
            <Space>
              <QuestionCircleOutlined />
              帮助
            </Space>
          </span>
          <span className="login-header-action">
            <Space>
              <SoundOutlined />
              反馈
            </Space>
          </span>
        </div>
      </div>
    </Layout.Header>
  )
}

export default LoginHeaderUI
