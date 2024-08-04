import React, { useEffect, useState } from 'react'
import { message } from 'antd'
import { ProfileService } from '@/services/Profile'
import ProfileInfoUI from '../component/InfoUI'
import { useGlobalStore } from '@/stores/index'
import { AccountInfoType } from '@/types/accountType'

const ProfileInfo: React.FC = () => {
  const [profileAccountInfo, setProfileAccountInfo] = useState<AccountInfoType>()
  const { getAccountInfo, setAccountInfo } = useGlobalStore()

  useEffect(() => {
    // setProfileAccountInfo(accountInfo)
  }, [])

  /**
   * 修改个人资料操作
   * @param values
   */
  const onFinishCallback = (values: any) => {
    ProfileService.profileUpdate(values)
      .then(() => {
        message.success('保存成功', 1)
        setAccountInfo(values) // 更新 accountInfo store
      })
      .catch((e) => {
        console.log('保存失败：', e)
      })
  }

  return (
    <div className="pdt24">
      <ProfileInfoUI accountInfo={profileAccountInfo} onFinishCallback={onFinishCallback} />
    </div>
  )
}

export default ProfileInfo
