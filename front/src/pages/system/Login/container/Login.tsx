import React from 'react'
import { useNavigate } from 'react-router-dom'
import { HOME_ROOT_PATH } from '@/router/type'
import { LoginResp } from '@/types/loginType'
import { useGlobalStore } from '@/stores/index'
import { SystemLoginService } from '@/services/SystemLogin'
import LoginHeaderUI from '../component/Header'
import LoginContentUI from '../component/Content'
import LoginFooterUI from '../component/Footer'
import { message } from 'antd'

const Login: React.FC = () => {
  const { setToken, setAccountInfo } = useGlobalStore()
  const navigate = useNavigate()

  /**
   * 账号登录操作
   * @param values
   */
  const onSystemLogin = async (values: { account_name: string; password: string }) => {
    const loginInfo = await SystemLoginService.systemLogin({
      account_name: values.account_name,
      password: values.password,
      verify_code: 'mock'
    })
    if (loginInfo) {
      setToken(loginInfo.login_token)
      setAccountInfo(loginInfo.account_info)
      message.success('登录成功！', 2, () => {
        navigate(HOME_ROOT_PATH)
      })
    } else {
      message.error('登录失败')
    }
    return
  }

  /**
   * 手机号登录操作
   * @param values
   */
  const onPhoneLogin = (values: {
    account_name: string
    password: string
  }): Promise<boolean | void> => {
    return Promise.resolve()
  }

  return (
    <div className="login-body">
      <LoginHeaderUI />
      <LoginContentUI onSystemLogin={onSystemLogin} onPhoneLogin={onPhoneLogin} />
      <LoginFooterUI />
    </div>
  )
}

export default Login
