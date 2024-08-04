import React, { ReactNode } from 'react'
import { Button, Dropdown, Popover, Space, Tree } from 'antd'
import type { GetProps, MenuProps, TreeDataNode } from 'antd'
import {
  DownOutlined,
  CarryOutOutlined,
  CaretDownOutlined,
  PlusOutlined,
  EllipsisOutlined,
  MoreOutlined,
  UserOutlined,
  FormOutlined,
  CopyOutlined,
  DeleteOutlined,
  HolderOutlined,
  RetweetOutlined
} from '@ant-design/icons'

type DirectoryTreeProps = GetProps<typeof Tree.DirectoryTree>

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

const TreeCustomTitle = (title: string) => {
  return (
    <div className="custom-title-wrapper">
      <span className="custom-title">{title}</span>
      <div className="custom-icon">
        <Dropdown menu={menuProps}>
          <HolderOutlined style={{ fontSize: 16 }} />
        </Dropdown>
      </div>
    </div>
  )
}

const treeData: TreeDataNode[] = [
  {
    title: TreeCustomTitle('xxx开发组'),
    key: '0-0',
    icon: <CarryOutOutlined />,
    children: [
      {
        title: TreeCustomTitle('新人指南'),
        key: '0-0-0',
        children: [
          {
            title: TreeCustomTitle('入职指引'),
            key: '0-0-0-0',
            isLeaf: true
          },
          {
            title: TreeCustomTitle('工作流程'),
            key: '0-0-0-1',
            isLeaf: true
          },
          {
            title: TreeCustomTitle('研发环境'),
            key: '0-0-0-2',
            isLeaf: true
          }
        ]
      },
      {
        title: TreeCustomTitle('项目信息'),
        key: '0-0-1',
        children: [
          {
            title: TreeCustomTitle('专项项目1'),
            key: '0-0-1-0',
            isLeaf: true
          }
        ]
      },
      {
        title: TreeCustomTitle('团队规划'),
        key: '0-0-2',
        children: [
          {
            title: TreeCustomTitle('2021团队年度规划'),
            key: '0-0-2-0',
            isLeaf: true
          },
          {
            title: TreeCustomTitle('2022团队年度规划'),
            key: '0-0-2-1',
            isLeaf: true
          }
        ]
      }
    ]
  }
]

const SpaceSidebarUI = () => {
  const onSelect: DirectoryTreeProps['onSelect'] = (keys, info) => {
    console.log('Trigger Select', keys, info)
  }

  const onExpand: DirectoryTreeProps['onExpand'] = (keys, info) => {
    console.log('Trigger Expand', keys, info)
  }

  return (
    <div className="doc-sider">
      <DirectoryTree
        className="doc-tree"
        showIcon={false}
        switcherIcon={<DownOutlined />}
        defaultExpandAll
        onSelect={onSelect}
        onExpand={onExpand}
        treeData={treeData}
      />
    </div>
  )
}

export default SpaceSidebarUI
