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
  onSaveSubmit: (privilegeIds?: string[]) => void
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
const PrivilegeUI = (props: PrivilegeUIProps) => {
  const privilegeList = props.privilegeList
  const [checkedPrivilegeIds, setCheckedPrivilegeIds] = useState<string[]>([])

  useEffect(() => {
    const defaultPrivilegeIds = props.defaultPrivilegeIds ? props.defaultPrivilegeIds : []
    setCheckedPrivilegeIds(defaultPrivilegeIds)
  }, [props.defaultPrivilegeIds])

  if (privilegeList.length === 0) {
    return <Empty></Empty>
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
   * 菜单权限树点击回调
   * @param checkedKeysValue
   * @param e
   */
  const menuTreeOnCheckCallback = (
    checkedKeysValue: React.Key[] | any,
    info: any,
    menuPrivilegeItem: PrivilegeListItemType
  ) => {
    const parentIds = menuPrivilegeItem.parent_ids.split(',')
    const currCheckedPrivilegeIds = [
      ...checkedKeysValue, // 所有的子权限
      ...info.halfCheckedKeys,
      menuPrivilegeItem.privilege_id.toString(), // 自己的权限
      ...parentIds // 父级导航的权限
    ]
    console.log(info)
    // 选中状态: 已有的权限ID + 新增的权限ID + 自身的权限ID
    if (info.checked) {
      const newPrivilegeIds = [...checkedPrivilegeIds, ...currCheckedPrivilegeIds]
      setCheckedPrivilegeIds(Array.from(new Set(newPrivilegeIds))) // 去除重复
      return
    }
    // 取消状态：需要找到删除的子节点
    // 获取所有的子权限 ID
    const allChildPrivilegeIds = getAllChildPrivilegeIds(menuPrivilegeItem.child_privileges)
    const allChlidIds = arrayToString(allChildPrivilegeIds)
    // 最终的选中的权限ID
    const checkIdsMap = arrayToMapBool(checkedKeysValue)
    // 所有的权限 diff 最终的权限 = 删除的权限
    const deleteIds = allChlidIds.filter((chlidId) => {
      return checkIdsMap.has(chlidId) ? false : true
    })
    if (checkedKeysValue.length === 0) {
      deleteIds.push(menuPrivilegeItem.privilege_id.toString())
    }
    const deleteCheckedPrivilegeIdsMap = arrayToMapBool(deleteIds)
    setCheckedPrivilegeIds((prePrivilegeIds) => {
      return prePrivilegeIds.filter((prePrivilegeId) => {
        return deleteCheckedPrivilegeIdsMap.has(prePrivilegeId) ? false : true
      })
    })
  }

  /**
   * 菜单的选择操作
   * @param menuPrivilegeItem 菜单权限信息
   * @param e 事件
   * @returns
   */
  const menuCheckOnChange = (menuPrivilegeItem: PrivilegeListItemType, e: CheckboxChangeEvent) => {
    const allChildPrivilegeIds = getAllChildPrivilegeIds(menuPrivilegeItem.child_privileges)
    const currCheckedPrivilegeIds = [
      ...arrayToString(allChildPrivilegeIds), // 所有的子权限
      menuPrivilegeItem.privilege_id.toString(), // 自己的权限
      menuPrivilegeItem.parent_id.toString() // 父级导航的权限
    ]
    // 选中状态: 已有的权限ID + 新增的权限ID + 自身的权限ID
    if (e.target.checked) {
      const newPrivilegeIds = [...checkedPrivilegeIds, ...currCheckedPrivilegeIds]
      setCheckedPrivilegeIds(Array.from(new Set(newPrivilegeIds))) // 去除重复
      return
    }
    // 取消状态：只删除子节点，父节点就不删除了
    const deleteCheckedPrivilegeIds = [
      ...arrayToString(allChildPrivilegeIds), // 所有的子权限
      menuPrivilegeItem.privilege_id.toString() // 自己的权限
    ]
    const deleteCheckedPrivilegeIdsMap = arrayToMapBool(deleteCheckedPrivilegeIds)
    setCheckedPrivilegeIds((prePrivilegeIds) => {
      return prePrivilegeIds.filter((prePrivilegeId) => {
        return deleteCheckedPrivilegeIdsMap.has(prePrivilegeId) ? false : true
      })
    })
  }

  /**
   * 导航的选择操作
   * @param navPrivilegeItem 导航的权限信息
   * @param e 事件
   * @returns
   */
  const navCheckOnChange = (navPrivilegeItem: PrivilegeListItemType, e: CheckboxChangeEvent) => {
    const allChildPrivilegeIds = getAllChildPrivilegeIds(navPrivilegeItem.child_privileges)
    const currCheckedPrivilegeIds = [
      ...arrayToString(allChildPrivilegeIds),
      navPrivilegeItem.privilege_id.toString()
    ]
    // 选中状态: 已有的权限ID + 新增的权限ID + 自身的权限ID
    if (e.target.checked) {
      const newPrivilegeIds = [...checkedPrivilegeIds, ...currCheckedPrivilegeIds]
      setCheckedPrivilegeIds(Array.from(new Set(newPrivilegeIds))) // 去除重复
      return
    }
    // 取消状态：从列表删除所有的
    const currCheckedPrivilegeIdsMap = arrayToMapBool(currCheckedPrivilegeIds)
    setCheckedPrivilegeIds((prePrivilegeIds) => {
      return prePrivilegeIds.filter((prePrivilegeId) => {
        return currCheckedPrivilegeIdsMap.has(prePrivilegeId) ? false : true
      })
    })
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
                    onChange={(e) => menuCheckOnChange(menuPrivilegeItem, e)}
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
                  menuTreeOnCheckCallback(checkedKeysValue, info, menuPrivilegeItem)
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
            onChange={(e) => navCheckOnChange(privilegeListItem, e)}
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
      <Button type="primary" onClick={() => props.onSaveSubmit(checkedPrivilegeIds)}>
        保存
      </Button>
    </>
  )
}

export default PrivilegeUI
