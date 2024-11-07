import React, { useEffect, useState } from 'react'
import { Button, Layout, Menu, Spin } from 'antd'
import LayoutHeader from '@/components/Layout/Header'
import { useGlobalStore } from '@/stores'
import { Link, Outlet, useNavigate, useLocation } from 'react-router-dom'
import { LayoutHeaderSpaceKey } from '@/components/Layout/types'
import LayoutSider from '@/components/Layout/Sider'
import {
  RightOutlined,
  LeftOutlined,
  BookOutlined,
  HomeOutlined,
  FireOutlined,
  FolderOutlined,
  DashboardOutlined,
  AppstoreAddOutlined,
  ProductOutlined,
  FolderOpenOutlined,
  StarOutlined
} from '@ant-design/icons'

const { Content } = Layout

const SpaceIndexMenus: React.FC = () => {
  const location = useLocation()
  const selectedKey = location.pathname === 'spaces_hot' ? 'spaces_hot' : 'spaces_all'

  return (
    <Menu mode="inline" style={{ marginTop: 5 }} selectedKeys={[selectedKey]}>
      <Menu.Item key="spaces_all" icon={<DashboardOutlined />}>
        <Link to={'/spaces/all'}>全部空间</Link>
      </Menu.Item>
      <Menu.Item key="spaces_hot" icon={<FireOutlined />}>
        <Link to={'/spaces/hot'}>热门空间</Link>
      </Menu.Item>
    </Menu>
  )
}

const SpaceIndex: React.FC = () => {
  const { getAccountInfo } = useGlobalStore()

  return (
    <Layout>
      <LayoutHeader
        {...useGlobalStore()}
        accountInfo={getAccountInfo()}
        navSelectedKeys={[LayoutHeaderSpaceKey]}
      />
      <Layout>
        <LayoutSider content={<SpaceIndexMenus />} />
        <Layout.Content className="home-content" style={{ padding: 16 }}>
          <Outlet />
        </Layout.Content>
      </Layout>
    </Layout>
  )
}

export default SpaceIndex
