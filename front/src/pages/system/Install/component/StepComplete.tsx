import React from 'react'
import { Button, Result } from 'antd'
import { useNavigate } from 'react-router-dom'
import { AUTH_LOGIN_PATH } from '@/router/type'

const StepComplete: React.FC = () => {
  const navigate = useNavigate()

  return (
    <Result
      status="success"
      title="安装完成"
      subTitle="系统已成功安装，您可以进入登录页开始使用。"
      extra={
        <Button type="primary" onClick={() => navigate(AUTH_LOGIN_PATH)}>
          进入登录页
        </Button>
      }
    />
  )
}

export default StepComplete
