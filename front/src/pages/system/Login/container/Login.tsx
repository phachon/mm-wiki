import React from 'react'
import { useNavigate } from 'react-router-dom'
import { HOME_ROOT_PATH } from '@/router/type'
import { LoginResp } from '@/types/loginType'
import { useGlobalStore } from '@/stores/index'
import { LoginService } from '@/services/Login'
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
  const onSystemLogin = (values: {
    account_name: string
    password: string
  }): Promise<boolean | void> => {
    LoginService.systemLogin({
      account_name: values.account_name,
      password: values.password,
      verify_code: 'mock'
    })
      .then((loginInfo: LoginResp) => {
        message.success('登录成功！', 2, () => {
          setToken(loginInfo.login_token)
          setAccountInfo(loginInfo.account_info)
          navigate(HOME_ROOT_PATH)
        })
      })
      .catch((e) => {
        console.log('登录失败err:', e)
      })
    return Promise.resolve()
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
