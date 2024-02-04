import { Button, Card, Checkbox, Empty, List, Tabs, Tree } from 'antd'
import DynamicIcon from '../../../components/DynamicIcon/DynamicIcon'
import { PrivilegeInfoType, PrivilegeListItemType } from '@/types/privilegeType'
import { CaretDownOutlined } from '@ant-design/icons'
import { DataNode } from 'antd/lib/tree'
import React, { useState } from 'react'
import { CheckboxChangeEvent } from 'antd/es/checkbox'
import { arrayToMapBool, arrayToString } from '../../../utils/utils'
import { useEffect } from 'react'

const menuGrid = {
  gutter: 16,
  xs: 1,
  sm: 2,
  md: 4,
  lg: 4,
  xl: 6,
  xxl: 4
}

interface PrivilegeUIProps {
  /**
   * 默认选中的权限ID
   */
  defaultPrivilegeIds?: string[]
  /**
   * 权限列表数据
   */
  privilegeList: PrivilegeListItemType[]
  /**
   * 保存操作方法
   * @param privilegeIds
   */
  onFinishSubmit: (privilegeIds?: string[]) => void
}

/**
 * 获取 treeData 的key
 * @param privilegeInfo
 * @returns
 */
const getTreeDataKey = (privilegeInfo: PrivilegeInfoType): string => {
  return privilegeInfo.privilege_id.toString()
}

/**
 * 获取多个权限的 keys
 * @param privilegeList
 * @returns
 */
const getTreeDataKeys = (privilegeList?: PrivilegeListItemType[]): React.Key[] => {
  let keys: React.Key[] = []
  if (privilegeList?.length === 0) {
    return keys
  }
  privilegeList?.forEach((privilegeItem) => {
    keys.push(getTreeDataKey(privilegeItem))
    let childKeys = getTreeDataKeys(privilegeItem.child_privileges)
    keys.push(...childKeys)
  })
  return keys
}

/**
 * 获取 treeData 数据
 * @param privilegeList 权限列表
 * @returns treeData 数据
 */
const getListTreeData = (privilegeList?: PrivilegeListItemType[]): DataNode[] => {
  let listTreeData: DataNode[] = []
  privilegeList?.forEach((privilegeItem) => {
    let treeData: DataNode = {
      title: <span>{privilegeItem.name}</span>,
      key: getTreeDataKey(privilegeItem),
      icon: <DynamicIcon name={privilegeItem.icon} />,
      children: getListTreeData(privilegeItem.child_privileges)
    }
    listTreeData.push(treeData)
  })
  return listTreeData
}

/**
 * 获取多个权限的权限ID + 子权限ID
 * @param privilegeList
 * @returns 所有的权限 + 子权限的ID
 */
const getAllChildPrivilegeIds = (privilegeList?: PrivilegeListItemType[]): bigint[] => {
  let privilegeIds: bigint[] = []
  if (privilegeList?.length === 0) {
    return privilegeIds
  }
  privilegeList?.forEach((privilegeItem) => {
    privilegeIds.push(privilegeItem.privilege_id)
    let childPrivilegeIds = getAllChildPrivilegeIds(privilegeItem.child_privileges)
    privilegeIds.push(...childPrivilegeIds)
  })
  return privilegeIds
}

/**
 * 权限列表 UI 组件
 * @param props
 * @returns 组件
 */
const PrivilegeUI2 = (props: PrivilegeUIProps) => {
  const privilegeList = props.privilegeList
  const [checkedPrivilegeIds, setCheckedPrivilegeIds] = useState<string[]>([])

  useEffect(() => {
    const defaultPrivilegeIds = props.defaultPrivilegeIds ? props.defaultPrivilegeIds : []
    setCheckedPrivilegeIds(defaultPrivilegeIds)
  }, [props.defaultPrivilegeIds])

  if (privilegeList.length === 0) {
    return <Empty></Empty>
  }

  const onCheckedHandle = (isChecked: boolean, privilegeItem: PrivilegeListItemType) => {
    console.log('isChecked:', isChecked)
    console.log('privilegeItem:', privilegeItem)
  }

  /**
   * 获取菜单树选中的 key 列表
   */
  const getMenuTreeCheckedKeys = (menuTreePrivileges?: PrivilegeListItemType[]): React.Key[] => {
    let treeDatKeys: React.Key[] = []
    if (checkedPrivilegeIds.length === 0 || menuTreePrivileges?.length === 0) {
      return treeDatKeys
    }
    // 获取当前子菜单的所有的 key
    treeDatKeys = getTreeDataKeys(menuTreePrivileges)
    const checkedPrivilegeIdsMap = arrayToMapBool(checkedPrivilegeIds)
    return treeDatKeys.filter((treeDatKey) => {
      return checkedPrivilegeIdsMap.has(treeDatKey.toString())
    })
  }

  /**
   * 菜单和导航是否被选中
   * @param privilegeId 权限ID
   * @returns
   */
  const menuNavIsChecked = (privilegeId: bigint): boolean => {
    const checkedPrivilegeIdsMap = arrayToMapBool(checkedPrivilegeIds)
    return checkedPrivilegeIdsMap.has(privilegeId.toString())
  }

  /**
   * 获取 tabs 下的内容
   * @param privilegeListItem
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
                <span>
                  <Checkbox
                    onChange={(e: CheckboxChangeEvent) =>
                      onCheckedHandle(e.target.checked, menuPrivilegeItem)
                    }
                    checked={menuNavIsChecked(menuPrivilegeItem?.privilege_id)}
                  >
                    <DynamicIcon name={menuPrivilegeItem.icon} style={{ marginRight: '6px' }} />
                    {menuPrivilegeItem.name}
                  </Checkbox>
                </span>
              }
              type="inner"
              bodyStyle={{ padding: 0 }}
            >
              <Tree
                key={menuPrivilegeItem.privilege_id.toString()}
                checkable
                autoExpandParent
                showIcon
                showLine={{ showLeafIcon: false }}
                switcherIcon={<CaretDownOutlined />}
                rootStyle={{ padding: 8 }}
                treeData={getListTreeData(menuPrivilegeItem.child_privileges)}
                onCheck={(checkedKeysValue: React.Key[] | any, info: any) => {
                  // console.log(info.halfCheckedKeys + checkedKeysValue)
                  onCheckedHandle(info.checked, menuPrivilegeItem)
                }}
                checkedKeys={getMenuTreeCheckedKeys(menuPrivilegeItem.child_privileges)}
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
        <span>
          <Checkbox
            style={{ marginRight: '6px' }}
            onChange={(e) => onCheckedHandle(e.target.checked, privilegeListItem)}
            checked={menuNavIsChecked(privilegeListItem.privilege_id)}
          ></Checkbox>
          <DynamicIcon name={privilegeListItem.icon} style={{ marginRight: '6px' }} />
          {privilegeListItem.name}
        </span>
      ),
      children: getTabsContent(privilegeListItem)
    }
  }

  return (
    <>
      <Tabs
        type="card"
        items={privilegeList.map((privilegeListItem) => getTabsItem(privilegeListItem))}
      />
      <Button type="primary" onClick={() => props.onFinishSubmit(checkedPrivilegeIds)}>
        保存
      </Button>
    </>
  )
}

export default PrivilegeUI2
