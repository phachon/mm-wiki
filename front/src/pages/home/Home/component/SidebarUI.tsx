import React, { useState } from 'react'
import { Collapse, CollapseProps, Divider, Layout, Menu, Space } from 'antd'
import { ResizableBox } from 'react-resizable'
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
import 'react-resizable/css/styles.css'
import SubMenu from 'antd/es/menu/SubMenu'
import { DocList, SpaceList, TextDeliver } from './ToolsUI'
import DynamicIcon from '@/components/DynamicIcon/DynamicIcon'
import { SpaceInfoType } from '@/types/spaceType'
import { DocEntity } from '@/types/docType'

const mySpaces: any[] = [
  {
    icon: 'FolderOpenOutlined',
    name: '产品技术中心产品',
    link: '/product'
  },
  {
    icon: 'FolderOpenOutlined',
    name: '工程研发部门工程研发部门工程研发部门工程研发部门',
    link: '/engineering'
  },
  {
    icon: 'FolderOpenOutlined',
    name: '算法研发部门',
    link: '/algorithm'
  },
  {
    icon: 'FolderOpenOutlined',
    name: '产品技术中心产品',
    link: '/product'
  },
  {
    icon: 'FolderOpenOutlined',
    name: '工程研发部门工程研发部门工程研发部门工程研发部门',
    link: '/engineering'
  },
  {
    icon: 'FolderOpenOutlined',
    name: '算法研发部门',
    link: '/algorithm'
  }
]

const myDocs: any[] = [
  {
    icon: 'FileTextOutlined',
    name: '登录功能流程梳理',
    link: '/doc/ada'
  },
  {
    icon: 'FileTextOutlined',
    name: '技术开发手册',
    link: '/doc/ada'
  },
  {
    icon: 'FileTextOutlined',
    name: '产品安全合审核流程',
    link: '/doc/ada'
  },
  {
    icon: 'FileTextOutlined',
    name: '登录功能流程梳理',
    link: '/doc/ada'
  },
  {
    icon: 'FileTextOutlined',
    name: '技术开发手册',
    link: '/doc/ada'
  },
  {
    icon: 'FileTextOutlined',
    name: '产品安全合审核流程',
    link: '/doc/ada'
  },
  {
    icon: 'FileTextOutlined',
    name: '登录功能流程梳理',
    link: '/doc/ada'
  },
  {
    icon: 'FileTextOutlined',
    name: '技术开发手册',
    link: '/doc/ada'
  },
  {
    icon: 'FileTextOutlined',
    name: '产品安全合审核流程',
    link: '/doc/ada'
  }
]

// HomeSidebarUIProps 主页侧边栏UI属性
type HomeSidebarUIProps = {
  mySpaces: SpaceInfoType[]
  collectionSpaces: SpaceInfoType[]
  collectionDocs: DocEntity[]
}

// HomeSidebarUI 主页侧边栏UI
const HomeSidebarUI = (props: HomeSidebarUIProps) => {
  const [collapsed, setCollapsed] = useState(false)
  const [width, setWidth] = useState(208)

  const items: CollapseProps['items'] = [
    {
      key: 'my_spaces',
      label: (
        <Space>
          <AppstoreAddOutlined />
          <span>我的空间</span>
        </Space>
      ),
      children: (
        <div>
          <SpaceList items={props.mySpaces} />
        </div>
      )
    },
    {
      key: 'collect_spaces',
      label: (
        <Space>
          <StarOutlined />
          <span>收藏空间</span>
        </Space>
      ),
      children: <SpaceList items={props.collectionSpaces} />
    },
    {
      key: 'collect_docs',
      label: (
        <Space>
          <BookOutlined />
          <span>收藏文档</span>
        </Space>
      ),
      children: <DocList items={props.collectionDocs} />
    }
  ]

  const toggleCollapsed = () => {
    setCollapsed(!collapsed)
    setWidth(collapsed ? 208 : 0)
  }

  const handleResize = (e: React.SyntheticEvent, data: any) => {
    setWidth(data.size.width)
  }

  return (
    <div>
      <Menu mode="inline" style={{ marginTop: 5 }}>
        <Menu.Item key="1" icon={<DashboardOutlined />}>
          主页面板
        </Menu.Item>
        <Menu.Item key="2" icon={<FireOutlined />}>
          探索发现
        </Menu.Item>
      </Menu>
      <div style={{ paddingLeft: 6, paddingRight: 6 }}>
        <Divider
          style={{
            marginTop: 0,
            marginBottom: 6,
            borderWidth: 1,
            borderColor: '#e9e9e9'
          }}
        />
      </div>
      <Collapse defaultActiveKey={[]} ghost items={items} />
    </div>
  )
}

export default HomeSidebarUI
