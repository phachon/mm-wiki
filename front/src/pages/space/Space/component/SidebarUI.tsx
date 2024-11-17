import React, { ReactNode } from 'react'
import { Button, Divider, Dropdown, Input, Space, Tree } from 'antd'
import type { GetProps, MenuProps, TreeDataNode } from 'antd'
import {
  DownOutlined,
  SearchOutlined,
  FileTextOutlined,
  FolderOpenOutlined,
  PlusOutlined,
  FormOutlined,
  CopyOutlined,
  DeleteOutlined,
  HolderOutlined,
  PlusSquareOutlined,
  RetweetOutlined
} from '@ant-design/icons'
import './space.css'
import { SpaceInfoType } from '@/types/spaceType'
import { DocTreeEntity } from '@/types/docType'

const { DirectoryTree } = Tree

const handleMenuClick: MenuProps['onClick'] = (e) => {
  console.log('click', e)
}

const items: MenuProps['items'] = [
  {
    label: '添加',
    key: '1',
    icon: <PlusOutlined />
  },
  {
    label: '编辑',
    key: '2',
    icon: <FormOutlined />
  },
  {
    label: '复制',
    key: '3',
    icon: <CopyOutlined />
  },
  {
    label: '移动',
    key: '4',
    icon: <RetweetOutlined />
  },
  {
    label: '删除',
    key: '5',
    icon: <DeleteOutlined />,
    danger: true
  }
]

const menuProps = {
  items,
  onClick: handleMenuClick
}

const TreeCustomTitle = (node: TreeDataNode) => {
  return (
    <div className="custom-title-wrapper">
      <div className="custom-title">
        <span className="title-text">{node.title as string}</span>
      </div>
      <div className="custom-icon">
        <Dropdown menu={menuProps} trigger={['click']}>
          <HolderOutlined style={{ fontSize: 16 }} />
        </Dropdown>
      </div>
    </div>
  )
}

const treeData: TreeDataNode[] = [
  {
    title: '新人指南',
    key: '0-0-0',
    children: [
      {
        title: '入职指引',
        key: '0-0-0-0',
        children: [
          {
            title: '入职指引1',
            key: '0-0-0-0-0',
            isLeaf: true
          },
          {
            title: '入职指引2',
            key: '0-0-0-0-1',
            isLeaf: true
          },
          {
            title: '入职指引3',
            key: '0-0-0-0-2',
            isLeaf: true
          }
        ]
      },
      {
        title: '工作流程',
        key: '0-0-0-1',
        isLeaf: true
      },
      {
        title: '研发环境',
        key: '0-0-0-2',
        isLeaf: true
      }
    ]
  },
  {
    title: '项目信息',
    key: '0-0-1',
    children: [
      {
        title: '专项项目1',
        key: '0-0-1-0',
        isLeaf: true
      }
    ]
  },
  {
    title: '团队规划',
    key: '0-0-2',
    children: [
      {
        title: '2021团队年度规划',
        key: '0-0-2-0',
        isLeaf: true
      },
      {
        title: '2022团队年度规划',
        key: '0-0-2-1',
        isLeaf: true
      }
    ]
  }
]

type SpaceSidebarUIProps = {
  spaceInfo?: SpaceInfoType
  docs?: DocTreeEntity[]
}

const SpaceSidebarUI = (props: SpaceSidebarUIProps) => {
  const onSelect: GetProps<typeof Tree.DirectoryTree>['onSelect'] = (keys, info) => {
    console.log('Trigger Select', keys, info)
  }

  const onExpand: GetProps<typeof Tree.DirectoryTree>['onExpand'] = (keys, info) => {
    console.log('Trigger Expand', keys, info)
  }

  const onSearch = (value: string) => {
    console.log('搜索:', value)
  }

  const customIcon = (props: any) => {
    if (props.isLeaf) {
      return <FileTextOutlined />
    }
    return null
  }

  return (
    <div className="doc-sider">
      <div className="doc-sider-header">
        <h2 className="space-title" style={{ display: 'flex', alignItems: 'center' }}>
          <a href={`/space/${props.spaceInfo?.space_key}`}>
            <Space>
              <FolderOpenOutlined />
              {props.spaceInfo?.name}
            </Space>
          </a>
          <Button
            type="default"
            size="small"
            style={{ marginLeft: 10, width: 18, height: 18, lineHeight: '24px' }}
            icon={<PlusOutlined />}
          ></Button>
        </h2>
      </div>
      <Divider className="doc-sider-divider" />
      <div className="doc-sider-search">
        <Input
          placeholder="搜索文档"
          allowClear
          prefix={<SearchOutlined style={{ color: '#bfbfbf' }} />}
          onChange={(e) => console.log('搜索:', e.target.value)}
        />
      </div>
      <DirectoryTree
        className="doc-tree"
        showIcon={true}
        icon={customIcon}
        switcherIcon={<DownOutlined />}
        defaultExpandAll
        onSelect={onSelect}
        onExpand={onExpand}
        treeData={treeData}
        titleRender={TreeCustomTitle}
      />
    </div>
  )
}

export default SpaceSidebarUI
