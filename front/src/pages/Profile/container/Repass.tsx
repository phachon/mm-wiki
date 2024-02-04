import React from 'react'
import { message } from 'antd'
import { ProfileService } from '@/services/Profile'
import ProfileRepassUI from '../component/RepassUI'

const ProfileRepass: React.FC = () => {
  /**
   * 修改密码操作
   * @param values
   */
  const onFinishCallback = (values: { old_pwd: string; new_pwd: string; confirm_pwd: string }) => {
    // 判断两次密码是否一致
    if (values.confirm_pwd !== values.new_pwd) {
      message.error('确认密码与新密码不一致')
      return
    }
    // 修改密码请求
    ProfileService.profileRepass(values)
      .then((res) => {
        message.success('保存成功', 1)
      })
      .catch((e) => {
        console.log('修改密码失败:', e)
        message.error('保存失败')
      })
  }

  return (
    <div className="pdt24">
      <ProfileRepassUI onFinishCallback={onFinishCallback} />
    </div>
  )
}

export default ProfileRepass
