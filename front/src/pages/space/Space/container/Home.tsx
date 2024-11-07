import LayoutHeader from '@/components/Layout/Header'
import LayoutSider from '@/components/Layout/Sider'
import { LayoutHeaderSpaceKey } from '@/components/Layout/types'
import { useGlobalStore } from '@/stores'
import { Layout } from 'antd'
import React from 'react'
import { Outlet, useParams } from 'react-router-dom'

const SpaceHome: React.FC = () => {
  const { key } = useParams<{ key: string }>()

  React.useEffect(() => {
    if (key) {
      console.log('加载空间数据:', key)
    }
  }, [key])

  if (!key) {
    return <div>空间 Key 不能为空</div>
  }

  const { getAccountInfo } = useGlobalStore()

  return (
    <Layout>
      <LayoutHeader
        {...useGlobalStore()}
        accountInfo={getAccountInfo()}
        navSelectedKeys={[LayoutHeaderSpaceKey]}
      />
      <Layout>
        <LayoutSider content={<p>菜单树</p>} />
        <Layout.Content className="home-content" style={{ padding: 16 }}>
          <Outlet />
        </Layout.Content>
      </Layout>
    </Layout>
  )
}

export default SpaceHome
