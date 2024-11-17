import LayoutHeader from '@/components/Layout/Header'
import LayoutSider from '@/components/Layout/Sider'
import { LayoutHeaderSpaceKey } from '@/components/Layout/types'
import { useGlobalStore } from '@/stores'
import { Layout } from 'antd'
import React, { useState } from 'react'
import { Outlet, useParams } from 'react-router-dom'
import SpaceSidebarUI from '../component/SidebarUI'
import { DocTreeEntity } from '@/types/docType'
import { SpaceDocsResp, SpaceInfoType } from '@/types/spaceType'
import { SpaceService } from '@/services/Space'

const SpaceHome: React.FC = () => {
  const { key } = useParams<{ key: string }>()
  const [spaceDocs, setSpaceDocs] = useState<DocTreeEntity[]>([])
  const [spaceInfo, setSpaceInfo] = useState<SpaceInfoType>()

  React.useEffect(() => {
    if (key) {
      SpaceService.getSpaceDocs(key)
        .then((res: SpaceDocsResp) => {
          setSpaceDocs(res.docs)
          setSpaceInfo(res.info)
        })
        .catch((err) => {
          console.error(err)
        })
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
        <LayoutSider content={<SpaceSidebarUI spaceInfo={spaceInfo} docs={spaceDocs} />} />
        <Layout.Content className="home-content" style={{ padding: 16 }}>
          <Outlet />
        </Layout.Content>
      </Layout>
    </Layout>
  )
}

export default SpaceHome
