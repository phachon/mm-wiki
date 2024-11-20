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
import { DocTreeEntity, DocType } from '@/types/docType'

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

type SpaceSidebarUIProps = {
  spaceInfo?: SpaceInfoType
  dirTree?: DocTreeEntity[]
  homeDoc?: DocTreeEntity
  onClickAddDoc?: (docInfo?: DocTreeEntity) => void
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

  const convertTreeData = (treeData?: DocTreeEntity[]): TreeDataNode[] => {
    if (!treeData) {
      return []
    }
    return treeData.map((docItem: DocTreeEntity) => ({
      title: docItem.name,
      key: docItem.doc_id.toString(),
      isLeaf: docItem.type === DocType.DOC,
      children: docItem.children ? convertTreeData(docItem.children) : undefined
    }))
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
            onClick={() => props.onClickAddDoc && props.onClickAddDoc(props.homeDoc)}
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
        treeData={convertTreeData(props.dirTree)}
        titleRender={TreeCustomTitle}
      />
    </div>
  )
}

export default SpaceSidebarUI
