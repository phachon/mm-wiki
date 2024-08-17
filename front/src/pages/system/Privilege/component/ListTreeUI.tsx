import { Card, Empty, List, Popconfirm, Space, Tabs, Tree } from 'antd'
import DynamicIcon from '@/components/DynamicIcon/DynamicIcon'
import { PrivilegeInfoType, PrivilegeListItemType } from '@/types/privilegeType'
import { EditOutlined, DeleteOutlined, CaretDownOutlined } from '@ant-design/icons'
import { DataNode } from 'antd/lib/tree'

const menuGrid = {
  gutter: 16,
  xs: 1,
  sm: 2,
  md: 4,
  lg: 4,
  xl: 6,
  xxl: 4
}

// PrivilegeListTreeUIProps 树形UI
interface PrivilegeListTreeUIProps {
  /**
   * 权限列表结构
   */
  privilegeList: PrivilegeListItemType[]

  /**
   * 修改点击操作
   * @param privilegeInfo
   * @returns
   */
  onEditClick: (privilegeInfo: PrivilegeInfoType) => void

  /**
   * 删除确认操作
   * @param privilegeInfo
   * @returns
   */
  onDeleteConfirm: (privilegeInfo: PrivilegeInfoType) => void
}

/**
 * 权限列表 tree UI 组件
 * @param props
 * @returns 组件
 */
const PrivilegeListTreeUI = (props: PrivilegeListTreeUIProps) => {
  const privilegeList = props.privilegeList
  if (privilegeList.length === 0) {
    return <Empty></Empty>
  }

  /**
   * 操作方法
   * @param privilegeItem
   * @returns
   */
  const getPrivilegeItemAction = (privilegeItem: PrivilegeListItemType) => {
    return (
      <span>
        {privilegeItem.action?.is_edit && (
          <a onClick={() => props.onEditClick(privilegeItem)}>
            <EditOutlined style={{ marginLeft: 6, marginRight: 0 }} />
          </a>
        )}
        {privilegeItem.action?.is_delete && (
          <Popconfirm
            title="确定要删除吗?"
            onConfirm={() => {
              props.onDeleteConfirm(privilegeItem)
            }}
            okText="确定"
            cancelText="取消"
          >
            <a>
              <DeleteOutlined style={{ marginLeft: 6, marginRight: 0 }} />
            </a>
          </Popconfirm>
        )}
      </span>
    )
  }

  /**
   * 获取 treeData 数据
   * @param privilegeList 权限列表
   * @param listTreeData 菜单树数据
   * @returns treeData 数据
   */
  const getListTreeData = (
    listTreeData: DataNode[],
    privilegeList?: PrivilegeListItemType[]
  ): DataNode[] => {
    privilegeList?.forEach((privilegeItem) => {
      let treeData: DataNode = {
        title: (
          <>
            {privilegeItem.name}
            {getPrivilegeItemAction(privilegeItem)}
          </>
        ),
        key: String(privilegeItem.parent_id) + '-' + String(privilegeItem.privilege_id),
        icon: <DynamicIcon name={privilegeItem.icon} />,
        selectable: false,
        children: getListTreeData([], privilegeItem.child_privileges)
      }
      listTreeData.push(treeData)
    })
    return listTreeData
  }

  /**
   * 获取 tabs 下的内容
   * @param privilegeListItem
   * @returns
   */
  const getTabsContent = (privilegeListItem: PrivilegeListItemType) => {
    return (
      <List
        grid={menuGrid}
        dataSource={privilegeListItem.child_privileges}
        renderItem={(menuPrivilegeItem) => (
          <List.Item key={menuPrivilegeItem?.privilege_id.toString()} style={{ padding: 0 }}>
            <Card
              size={'small'}
              title={
                <Space>
                  <DynamicIcon name={menuPrivilegeItem?.icon} />
                  {menuPrivilegeItem.name}
                </Space>
              }
              type="inner"
              bodyStyle={{ padding: 0 }}
              extra={getPrivilegeItemAction(menuPrivilegeItem)}
            >
              <Tree
                autoExpandParent
                showIcon
                showLine={{ showLeafIcon: false }}
                switcherIcon={<CaretDownOutlined />}
                rootStyle={{ padding: 8 }}
                checkStrictly={true}
                defaultExpandAll={true}
                treeData={getListTreeData([], menuPrivilegeItem.child_privileges)}
              />
            </Card>
          </List.Item>
        )}
      />
    )
  }

  /**
   * 获取每个 tab item 信息
   * @param privilegeListItem
   * @returns
   */
  const getTabsItem = (privilegeListItem: PrivilegeListItemType) => {
    return {
      key: privilegeListItem.privilege_id.toString(),
      label: (
        <>
          <DynamicIcon name={privilegeListItem.icon} style={{ marginRight: 6 }} />
          {privilegeListItem.name}
          {getPrivilegeItemAction(privilegeListItem)}
        </>
      ),
      children: getTabsContent(privilegeListItem)
    }
  }

  return (
    <div className="panel-body">
      <Tabs
        type="card"
        style={{ marginTop: 10 }}
        items={privilegeList.map((privilegeListItem) => getTabsItem(privilegeListItem))}
      />
    </div>
  )
}

export default PrivilegeListTreeUI
