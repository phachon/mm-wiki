import React, { useEffect, useState } from 'react'
import { Button, Layout, Spin } from 'antd'
import 'allotment/dist/style.css'
import '../component/home.css'
import HomeSidebarUI from '../component/SidebarUI'
import LayoutHeader from '@/components/Layout/Header'
import { useGlobalStore } from '@/stores'
import { useNavigate } from 'react-router-dom'
import { LayoutHeaderHomeKey } from '@/components/Layout/types'
import LayoutSider from '@/components/Layout/Sider'

const { Content } = Layout

const HomeIndex: React.FC = () => {
  const [collapsed, setCollapsed] = useState(false)
  const navigate = useNavigate()
  const { getAccountInfo } = useGlobalStore()

  return (
    <Layout>
      <LayoutHeader
        {...useGlobalStore()}
        accountInfo={getAccountInfo()}
        navSelectedKeys={[LayoutHeaderHomeKey]}
      />
      <Layout>
        <LayoutSider content={<HomeSidebarUI />} />
        <Layout>
          <Content className="home-content">
            <div>正文我啊啊 啊啊啊</div>
          </Content>
        </Layout>
      </Layout>
    </Layout>
  )
}

export default HomeIndex
