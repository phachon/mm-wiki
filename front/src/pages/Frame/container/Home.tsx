import React, { useEffect } from 'react'
import { Layout, Spin } from 'antd'
import { useGlobalStore } from '@/stores/index'
import { Outlet, useLocation } from 'react-router-dom'
import FrameHeaderUI from '../component/HeaderUI'
import FrameBreadcrumbUI from '../component/BreadcrumbUI'
import FrameSidebarUI from '../component/SidebarUI'
import FrameFooterUI from '../component/FooterUI'
import '../component/home.css'
import { SettingConfig } from '@/config/setting'

const FrameHome: React.FC = () => {
  const initProfileInfo = useGlobalStore((state: any) => state.initProfileInfo)
  const { isLoading } = useGlobalStore()
  const location = useLocation()

  useEffect(() => {
    initProfileInfo(location.pathname)
  }, [])

  return (
    <Layout>
      <FrameHeaderUI />
      <Layout>
        <FrameSidebarUI />
        <Layout className="admin-main">
          <FrameBreadcrumbUI />
          <Layout.Content className="admin-content">
            {isLoading ? (
              <Spin spinning={true} size="large" tip="Loading..." className="loading"></Spin>
            ) : (
              <Outlet />
            )}
          </Layout.Content>
          <FrameFooterUI text={SettingConfig.footerShowText} />
        </Layout>
      </Layout>
    </Layout>
  )
}

export default FrameHome
