import React, { useEffect, useState } from 'react'
import { Button, Layout, Spin } from 'antd'
import { useGlobalStore } from '@/stores/index'
import SpaceFooterUI from '../component/FooterUI'
import { Allotment } from 'allotment'
import 'allotment/dist/style.css'
import '../component/home.css'
import { SettingConfig } from '@/config/setting'
import { DoubleLeftOutlined, DoubleRightOutlined } from '@ant-design/icons'
import SpaceSidebarUI from '../component/SidebarUI'

const { Content } = Layout

const Demo1: React.FC = () => {
  // const { isLoading } = useGlobalStore()

  // const ref = React.useRef(ref)

  const [collapsed, setCollapsed] = useState(false)

  useEffect(() => {
    // initProfileInfo(location.pathname)
  }, [])

  const [siderWidth, setSiderWidth] = useState(208)

  const handleToggleCollapse = () => {
    setCollapsed(!collapsed)
  }

  return (
    <Layout>
      {/* <FrameHeaderUI /> */}
      <Layout>
        <Allotment
          defaultSizes={collapsed ? [0, 48] : [0, 208]}
          separator={true}
          className="allotment"
        >
          <Allotment.Pane minSize={208} className="space-sidebar">
            <SpaceSidebarUI />
          </Allotment.Pane>
          <Allotment.Pane snap>
            <div className="separator-container">
              <Button
                className="toggle-collapse-btn"
                onClick={handleToggleCollapse}
                icon={collapsed ? <DoubleRightOutlined /> : <DoubleLeftOutlined />}
              />
            </div>
            <Layout>
              <Content className="space-content">
                <div>正文我啊啊 啊啊啊</div>
              </Content>
              <SpaceFooterUI text={SettingConfig.footerShowText} />
            </Layout>
          </Allotment.Pane>
        </Allotment>
      </Layout>
    </Layout>
  )
}

export default Demo1
