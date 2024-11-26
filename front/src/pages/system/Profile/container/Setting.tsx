import React, { useEffect, useState } from 'react'
import { message, Tabs, TabsProps } from 'antd'
import { SystemProfileService } from '@/services/SystemProfile'
import { AccountInfoType } from '@/types/accountType'
import { ProfileInfoResp } from '@/types/profileType'
import ProfileBasicSettingUI from '../component/BasicSettingUI'
import { LockOutlined, FormOutlined } from '@ant-design/icons'
import ProfileRepassUI from '../component/RepassUI'

const ProfileSetting: React.FC = () => {
  const [profileAccountInfo, setProfileAccountInfo] = useState<AccountInfoType>()

  useEffect(() => {
    initProfileInfo()
  }, [])

  const initProfileInfo = () => {
    SystemProfileService.getProfileInfo()
      .then((resp: ProfileInfoResp) => {
        setProfileAccountInfo(resp.account_info)
      })
      .catch((e) => {
        console.error('获取个人资料err:', e)
      })
  }

  /**
   * 修改密码操作
   * @param values
   */
  const onRepassSave = (values: { old_pwd: string; new_pwd: string; confirm_pwd: string }) => {
    // 判断两次密码是否一致
    if (values.confirm_pwd !== values.new_pwd) {
      message.error('确认密码与新密码不一致')
      return
    }
    // 修改密码请求
    SystemProfileService.profileRepass(values)
      .then((res) => {
        message.success('保存成功', 1)
      })
      .catch((e) => {
        console.log('修改密码失败:', e)
        message.error('保存失败')
      })
  }

  /**
   * 修改基础设置操作
   * @param values
   */
  const onBasicEditSubmit = (values: any) => {
    SystemProfileService.profileUpdate(values)
      .then(() => {
        message.success('保存成功', 1)
      })
      .catch((e) => {
        console.log('保存失败：', e)
      })
  }

  const items: TabsProps['items'] = [
    {
      key: '1',
      label: '基本设置',
      icon: <FormOutlined />,
      children: (
        <ProfileBasicSettingUI accountInfo={profileAccountInfo} onEditSubmit={onBasicEditSubmit} />
      )
    },
    {
      key: '2',
      label: '密码设置',
      icon: <LockOutlined />,
      children: <ProfileRepassUI onSaveSubmit={onRepassSave} />
    }
  ]

  return (
    <div className="panel" style={{ paddingLeft: 16 }}>
      <div className="panel-body">
        <Tabs defaultActiveKey="1" items={items} />
      </div>
    </div>
  )
}

export default ProfileSetting
