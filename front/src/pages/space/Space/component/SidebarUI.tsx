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
  RetweetOutlined
} from '@ant-design/icons'
import './space.css'
import { SpaceInfoType } from '@/types/spaceType'
import { DocTreeEntity, DocType } from '@/types/docType'

const { DirectoryTree } = Tree

export const ActionType = {
  ADD: 'add',
  EDIT: 'edit',
  COPY: 'copy',
  MOVE: 'move',
  DELETE: 'delete'
}

const items: MenuProps['items'] = [
  {
    label: '添加',
    key: ActionType.ADD,
    icon: <PlusOutlined />
  },
  {
    label: '编辑',
    key: ActionType.EDIT,
    icon: <FormOutlined />
  },
  {
    label: '复制',
    key: ActionType.COPY,
    icon: <CopyOutlined />
  },
  {
    label: '移动',
    key: ActionType.MOVE,
    icon: <RetweetOutlined />
  },
  {
    label: '删除',
    key: ActionType.DELETE,
    icon: <DeleteOutlined />,
    danger: true
  }
]

type SpaceSidebarUIProps = {
  spaceInfo?: SpaceInfoType
  dirTree?: DocTreeEntity[]
  homeDoc?: DocTreeEntity
  onClickDocSelect?: (docId: string) => void
  onClickDocAction?: (action: string, node: TreeDataNode) => void
}

const SpaceSidebarUI = (props: SpaceSidebarUIProps) => {
  /**
   * 自定义树节点标题
   * @param node 树节点
   * @returns 自定义标题
   */
  const treeCustomTitle = (node: TreeDataNode) => {
    return (
      <div className="custom-title-wrapper">
        <div className="custom-title">
          <span className="title-text">{node.title as string}</span>
        </div>
        <div className="custom-icon">
          <Dropdown
            menu={{
              items,
              onClick: (e) => {
                props.onClickDocAction && props.onClickDocAction(e.key, node)
              }
            }}
            trigger={['click']}
          >
            <HolderOutlined style={{ fontSize: 16 }} />
          </Dropdown>
        </div>
      </div>
    )
  }

  const onSelect: GetProps<typeof Tree.DirectoryTree>['onSelect'] = (keys, info) => {
    console.log('Trigger Select', keys, info)
    if (keys.length > 0) {
      props.onClickDocSelect && props.onClickDocSelect(keys[0].toString())
    }
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
            onClick={() =>
              props.onClickDocAction &&
              props.onClickDocAction(ActionType.ADD, {
                key: props.homeDoc?.doc_id.toString() || '',
                title: props.homeDoc?.name || ''
              })
            }
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
        titleRender={treeCustomTitle}
      />
    </div>
  )
}

export default SpaceSidebarUI
