import React, { useEffect, useState } from 'react'
import { message } from 'antd'
import { ProfileService } from '@/services/Profile'
import { AccountInfoType } from '@/types/accountType'
import { ProfileInfoResp } from '@/types/profileType'
import ProfileInfoUI from '../component/InfoUI'

const ProfileInfo: React.FC = () => {
  const [profileInfo, setProfileInfo] = useState<ProfileInfoResp>()

  useEffect(() => {
    initProfileInfo()
  }, [])

  const initProfileInfo = () => {
    ProfileService.getProfileInfo()
      .then((resp: ProfileInfoResp) => {
        setProfileInfo(resp)
      })
      .catch((e) => {
        console.error('获取个人资料err:', e)
      })
  }

  return (
    <div className="panel">
      <ProfileInfoUI profileInfo={profileInfo} />
    </div>
  )
}

export default ProfileInfo
