import { TextDeliver } from '@/pages/home/Home/component/ToolsUI'
import {
  SpaceTypeRadioOptions,
  SpaceVisitLevelRadioOptions
} from '@/pages/system/Space/component/ToolsUI'
import { SpaceTypeTeam } from '@/types/spaceType'
import {
  PlusOutlined,
  SettingOutlined,
  ArrowLeftOutlined,
  LockOutlined,
  ShareAltOutlined,
  QuestionCircleOutlined
} from '@ant-design/icons'
import { Button, Input, Space, Upload, Tabs, Divider, Typography, Form, Radio, Tooltip } from 'antd'
import FormItem from 'antd/es/form/FormItem'
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'

const { TabPane } = Tabs

type SpaceBasicSettingUIProps = {}

const SpaceBasicSettingUI = (props: SpaceBasicSettingUIProps) => {
  return (
    <div style={{ padding: 16 }}>
      <h3>分享设置</h3>
      {/* 分享设置表单内容 */}
    </div>
  )
}

export default SpaceBasicSettingUI
