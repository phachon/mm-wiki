import { Button, message, Tabs } from 'antd'
import {
  ArrowLeftOutlined,
  LockOutlined,
  PlusOutlined,
  SettingOutlined,
  ShareAltOutlined
} from '@ant-design/icons'
import { useNavigate } from 'react-router-dom'
import { useEffect, useState } from 'react'
import SpaceBasicSettingUI from '../component/BasicSettingUI'
import PermissionSettingUI from '../component/PermissionSettingUI'
import ShareSettingUI from '../component/ShareSettingUI'
import { useGlobalStore } from '@/stores'
import { SpaceSpaceService } from '@/services/SpaceSpace'
import { SpacePermissionInfoType, SpacePermissionListResp } from '@/types/spaceType'

const SpaceSetting: React.FC = () => {
  const navigate = useNavigate()
  const store = useGlobalStore()
  const [activeTab, setActiveTab] = useState('basic')
  const [adminPermissions, setAdminPermissions] = useState<SpacePermissionInfoType[]>([])
  const [departmentPermissions, setDepartmentPermissions] = useState<SpacePermissionInfoType[]>([])
  const [accountPermissions, setAccountPermissions] = useState<SpacePermissionInfoType[]>([])

  useEffect(() => {}, [])

  const handleTabChange = (key: string) => {
    if (key === 'permissions') {
      fetchPermissionList()
    }
    setActiveTab(key)
  }

  const fetchPermissionList = () => {
    SpaceSpaceService.getSpaceSettingPermissionList(store.spaceInfo?.space_id)
      .then((res: SpacePermissionListResp) => {
        setAdminPermissions(res.admin_list)
        setDepartmentPermissions(res.department_list)
        setAccountPermissions(res.account_list)
      })
      .catch((err) => {
        console.log('获取权限列表失败:', err)
      })
  }

  const onBasicSettingSave = (values: any) => {
    SpaceSpaceService.modifySpaceBasicSetting(values)
      .then((res) => {
        message.success('保存成功', 2, () => {
          window.location.reload()
        })
      })
      .catch((err) => {
        message.error('保存失败:', err)
      })
  }

  return (
    <div style={{ padding: '2px 20px 14px 20px' }}>
      <div>
        <h2>
          <Button
            type="default"
            icon={<ArrowLeftOutlined />}
            onClick={() => navigate(-1)}
            style={{ marginRight: 12 }}
          />
          空间设置
        </h2>
      </div>
      <Tabs
        activeKey={activeTab}
        onChange={handleTabChange}
        type="card"
        items={[
          {
            label: (
              <span>
                <SettingOutlined /> 基础设置
              </span>
            ),
            key: 'basic',
            children: (
              <SpaceBasicSettingUI spaceInfo={store.spaceInfo} onSaveSubmit={onBasicSettingSave} />
            )
          },
          {
            label: (
              <span>
                <LockOutlined /> 权限设置
              </span>
            ),
            key: 'permissions',
            children: (
              <PermissionSettingUI
                adminPermissions={adminPermissions}
                departmentPermissions={departmentPermissions}
                accountPermissions={accountPermissions}
              />
            )
          },
          {
            label: (
              <span>
                <ShareAltOutlined /> 分享设置
              </span>
            ),
            key: 'sharing',
            children: <ShareSettingUI />
          },
          {
            label: (
              <span>
                <PlusOutlined /> 其他设置
              </span>
            ),
            key: 'others',
            children: <ShareSettingUI />
          }
        ]}
      />
    </div>
  )
}

export default SpaceSetting
