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
import { SpaceInfoType } from '@/types/spaceType'
import { DocEntity } from '@/types/docType'
import { HomeIndexService } from '@/services/Home'

const { Content } = Layout

const HomeIndex: React.FC = () => {
  const [collapsed, setCollapsed] = useState(false)
  const navigate = useNavigate()
  const { getAccountInfo } = useGlobalStore()
  const [mySpaces, setMySpaces] = useState<SpaceInfoType[]>([])
  const [collectionSpaces, setCollectionSpaces] = useState<SpaceInfoType[]>([])
  const [collectionDocs, setCollectionDocs] = useState<DocEntity[]>([])

  useEffect(() => {
    initSiderBar()
  }, [])

  const initSiderBar = async () => {
    // 获取我的空间
    HomeIndexService.getMySpaces().then((res) => {
      setMySpaces(res.list ? res.list : [])
    })
    // 获取收藏的空间
    HomeIndexService.getCollectionSpaces().then((res) => {
      setCollectionSpaces(res.list ? res.list : [])
    })
    // 获取收藏的文档
    HomeIndexService.getCollectionDocs().then((res) => {
      setCollectionDocs(res.list ? res.list : [])
    })
  }

  return (
    <Layout>
      <LayoutHeader
        {...useGlobalStore()}
        accountInfo={getAccountInfo()}
        navSelectedKeys={[LayoutHeaderHomeKey]}
      />
      <Layout>
        <LayoutSider
          content={
            <HomeSidebarUI
              mySpaces={mySpaces}
              collectionSpaces={collectionSpaces}
              collectionDocs={collectionDocs}
            />
          }
        />
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
