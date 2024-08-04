import React, { useEffect, useState } from 'react'
import { Button, Layout, Spin } from 'antd'
import { Allotment } from 'allotment'
import 'allotment/dist/style.css'
import '../component/home.css'
import HomeSidebarUI from '../component/SidebarUI'
import LayoutHeader from '@/components/Layout/Header'
import { useGlobalStore } from '@/stores'
import { useNavigate } from 'react-router-dom'
import { LayoutHeaderHomeKey } from '@/components/Layout/types'

const { Content } = Layout

const HomeIndex: React.FC = () => {
  const [collapsed, setCollapsed] = useState(false)
  const navigate = useNavigate()
  const { getAccountInfo } = useGlobalStore()

  useEffect(() => {
    // initProfileInfo(location.pathname)
  }, [])

  return (
    <Layout>
      <LayoutHeader
        {...useGlobalStore()}
        accountInfo={getAccountInfo()}
        navSelectedKeys={[LayoutHeaderHomeKey]}
      />
      <Layout>
        <Allotment defaultSizes={collapsed ? [0, 48] : [0, 208]} separator={true}>
          <Allotment.Pane minSize={208}>
            <HomeSidebarUI />
          </Allotment.Pane>
          <Allotment.Pane snap>
            <Layout>
              <Content className="home-content">
                <div>正文我啊啊 啊啊啊</div>
              </Content>
              {/* <SpaceFooterUI text={SettingConfig.footerShowText} /> */}
            </Layout>
          </Allotment.Pane>
        </Allotment>
      </Layout>
    </Layout>
  )
}

export default HomeIndex
