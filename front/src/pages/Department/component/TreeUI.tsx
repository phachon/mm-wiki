import { Button, Card, Tree } from 'antd'
import { useEffect, useState } from 'react'
import { DeleteTwoTone, PlusSquareTwoTone, EditTwoTone, ApartmentOutlined } from '@ant-design/icons'
import { DataNode } from 'antd/lib/tree'
import { DepartmentListItemType } from '@/types/departmentType'
import { getTreeData } from './ToolsUI'

interface DepartmentTreeUIProps {
  departments: DepartmentListItemType[]
  onDeleteClick: (departmentId: string) => void
  onEditClick: (departmentId: string) => void
  onAddClick: (parentId: string) => void
  onSaveSubmit: (values: any) => void
}

/**
 * 部门树 UI 组件
 * @param props 部门树 UI 组件数据
 * @returns 部门树 UI 组件
 */
const DepartmentTreeUI = (props: DepartmentTreeUIProps) => {
  const treeData = getTreeData(props.departments)
  const [expandedKeys, setExpandedKeys] = useState<React.Key[]>([]) // 展开 key 列表

  useEffect(() => {
    const allKeys = getAllKeys(treeData)
    setExpandedKeys(allKeys)
  }, [props.departments])

  const getAllKeys = (treeData: DataNode[]) => {
    const keys: React.Key[] = []
    treeData.forEach((item) => {
      keys.push(item.key)
      if (item.children) {
        keys.push(...getAllKeys(item.children))
      }
    })
    return keys
  }

  const renderTreeNode = (node: DataNode) => (
    <span>
      {node.title?.toString()}
      <span style={{ marginLeft: 8 }}>
        <PlusSquareTwoTone
          style={{ marginRight: 6 }}
          onClick={() => {
            props.onAddClick(node.key.toString())
          }}
        />
        <EditTwoTone
          style={{ marginRight: 6 }}
          onClick={() => {
            props.onEditClick(node.key.toString())
          }}
        />
        <DeleteTwoTone onClick={() => props.onDeleteClick(node.key.toString())} />
      </span>
    </span>
  )

  return (
    <div className="panel-body" style={{ height: '100%' }}>
      <Card
        title={
          <span>
            <ApartmentOutlined style={{ marginRight: 8 }} />
            部门树
          </span>
        }
        style={{ height: '100%' }}
      >
        {treeData.length === 0 ? (
          <Button type="link" onClick={() => props.onAddClick('0')} icon={<PlusSquareTwoTone />}>
            添加部门
          </Button>
        ) : (
          <Tree
            treeData={treeData}
            showLine={true}
            showIcon={false}
            defaultExpandAll={true}
            expandedKeys={expandedKeys}
            onExpand={setExpandedKeys}
            titleRender={renderTreeNode}
          />
        )}
      </Card>
    </div>
  )
}

export default DepartmentTreeUI
