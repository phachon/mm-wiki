import React, { useEffect, useState } from 'react'
import { message } from 'antd'
import { ProfileService } from '@/services/Profile'
import { AccountInfoType } from '@/types/accountType'
import { ProfileInfoResp } from '@/types/profileType'
import ProfileInfoUI from '../component/InfoUI'

const ProfileInfo: React.FC = () => {
  const [profileAccountInfo, setProfileAccountInfo] = useState<AccountInfoType>()

  useEffect(() => {
    initProfileInfo()
  }, [])

  const initProfileInfo = () => {
    ProfileService.getProfileInfo()
      .then((resp: ProfileInfoResp) => {
        setProfileAccountInfo(resp.account_info)
      })
      .catch((e) => {
        console.error('获取个人资料err:', e)
      })
  }

  /**
   * 修改个人资料操作
   * @param values
   */
  const onFinishCallback = (values: any) => {
    ProfileService.profileUpdate(values)
      .then(() => {
        message.success('保存成功', 1)
        // setAccountInfo(values) // 更新 accountInfo store
      })
      .catch((e) => {
        console.log('保存失败：', e)
      })
  }

  return (
    <div className="panel">
      <ProfileInfoUI accountInfo={profileAccountInfo} onFinishCallback={onFinishCallback} />
    </div>
  )
}

export default ProfileInfo
