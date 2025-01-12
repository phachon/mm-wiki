import { useEffect, useState } from 'react'
import { Button, Divider, Dropdown, Input, Space, Tree } from 'antd'
import type { GetProps, MenuProps, TreeDataNode } from 'antd'
import {
  DownOutlined,
  SearchOutlined,
  FileTextOutlined,
  PlusOutlined,
  FormOutlined,
  CopyOutlined,
  DeleteOutlined,
  HolderOutlined,
  RetweetOutlined,
  FolderOutlined,
  StarOutlined,
  StarTwoTone,
  StarFilled,
  PlusCircleOutlined
} from '@ant-design/icons'
import { SpaceInfoType } from '@/types/spaceType'
import { ActionType, DocTreeEntity, DocType } from '@/types/docType'
import './tree.css'

const { DirectoryTree } = Tree

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

// DocTreeUIProps 文档目录树组件属性
type DocTreeUIProps = {
  loading?: boolean
  spaceInfo?: SpaceInfoType
  docTree?: DocTreeEntity[]
  homeDoc?: DocTreeEntity
  selectDocId?: string
  isCollected?: boolean
  onClickDocSelect?: (docId: string) => void
  onClickDocAction?: (action: string, node: TreeDataNode) => void
  onCollectionStatusChange?: (spaceId: number, collected: boolean) => void
}

// DocTreeUI 文档目录树组件
const DocTreeUI = (props: DocTreeUIProps) => {
  const [expandedKeys, setExpandedKeys] = useState<string[]>([])
  const [selectedKeys, setSelectedKeys] = useState<string[]>([])

  useEffect(() => {
    if (props.selectDocId) {
      setSelectedKeys([props.selectDocId])
      expandParentNodes(props.selectDocId, props.docTree)
    }
  }, [props.selectDocId, props.docTree])

  // 展开父节点
  const expandParentNodes = (docId: string, treeData?: DocTreeEntity[]) => {
    if (!treeData) return
    const findParentKeys = (
      nodes: DocTreeEntity[],
      targetKey: string,
      path: string[] = []
    ): string[] | null => {
      for (const node of nodes) {
        const currentPath = [...path, node.doc_id.toString()]
        if (node.doc_id.toString() === targetKey) {
          return currentPath
        }
        if (node.children) {
          const result = findParentKeys(node.children, targetKey, currentPath)
          if (result) {
            return result
          }
        }
      }
      return null
    }

    const parentKeys = findParentKeys(treeData, docId)
    if (parentKeys) {
      setExpandedKeys((prevKeys) => Array.from(new Set([...prevKeys, ...parentKeys])))
    }
  }

  const treeCustomTitle = (node: TreeDataNode) => {
    return (
      <div className="custom-title-wrapper">
        <div className="custom-title">
          <span
            className="title-text"
            onClick={(e) => {
              console.log('点击文档 treeCustomTitle:', node)
              e.stopPropagation()
              props.onClickDocSelect && props.onClickDocSelect(node.key as string)
            }}
          >
            {node.title as string}
          </span>
        </div>
        <div className="custom-icon" onClick={(e) => e.stopPropagation()}>
          <Dropdown
            menu={{
              items,
              onClick: (e) => {
                props.onClickDocAction && props.onClickDocAction(e.key, node)
              }
            }}
            trigger={['hover']}
          >
            <HolderOutlined style={{ fontSize: 16 }} />
          </Dropdown>
        </div>
      </div>
    )
  }

  const onExpand: GetProps<typeof Tree.DirectoryTree>['onExpand'] = (keys, info) => {
    console.log('Trigger Expand', keys, info)
    setExpandedKeys(keys.map((key) => key.toString()))
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
              <FolderOutlined />
              {props.spaceInfo?.name}
            </Space>
          </a>
          <a
            style={{ marginLeft: 10 }}
            onClick={() =>
              props.onClickDocAction &&
              props.onClickDocAction(ActionType.ADD, {
                key: props.homeDoc?.doc_id.toString() || '',
                title: props.homeDoc?.name || ''
              })
            }
          >
            <PlusCircleOutlined />
          </a>
          <a
            style={{ marginLeft: 10 }}
            onClick={() =>
              props.onCollectionStatusChange &&
              props.onCollectionStatusChange(props.spaceInfo?.space_id || 0, !props.isCollected)
            }
          >
            {props.isCollected ? <StarFilled className="star-collection" /> : <StarOutlined />}
          </a>
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
        onExpand={onExpand}
        expandedKeys={expandedKeys}
        selectedKeys={selectedKeys}
        treeData={convertTreeData(props.docTree)}
        titleRender={treeCustomTitle}
      />
    </div>
  )
}

export default DocTreeUI
