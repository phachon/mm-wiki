import React from 'react'
import { message } from 'antd'
import { SystemLoginAuthService } from '@/services/SystemLoginAuth'
import LoginAuthFormUI from '../component/FormUI'

const LoginAuthAdd: React.FC = () => {
  const onSaveSubmit = (values: any) => {
    SystemLoginAuthService.saveLoginAuth(values).then(() => {
      message.success('添加成功', 2, () => {
        window.location.href = '/system/login_auth/list'
      })
    })
  }

  return (
    <div className="pdt24">
      <LoginAuthFormUI onSaveSubmit={onSaveSubmit} />
    </div>
  )
}

export default LoginAuthAdd
