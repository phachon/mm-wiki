import React, { ReactNode } from 'react'
import { Button, Dropdown, Layout, Popover, Space, Tree } from 'antd'
import type { GetProps, MenuProps, TreeDataNode } from 'antd'

type DirectoryTreeProps = GetProps<typeof Tree.DirectoryTree>

const HomeSidebarUI = () => {
  const onSelect: DirectoryTreeProps['onSelect'] = (keys, info) => {
    console.log('Trigger Select', keys, info)
  }

  const onExpand: DirectoryTreeProps['onExpand'] = (keys, info) => {
    console.log('Trigger Expand', keys, info)
  }

  return <div className="home-sider">收藏文档</div>
}

export default HomeSidebarUI
