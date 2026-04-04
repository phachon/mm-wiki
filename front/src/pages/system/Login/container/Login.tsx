import React, { useState, useEffect } from 'react'
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
  const [captchaId, setCaptchaId] = useState<string>('')
  const [captchaCode, setCaptchaCode] = useState<string>('')

  // 获取验证码
  const fetchCaptcha = async () => {
    try {
      const resp = await SystemLoginService.getCaptcha()
      setCaptchaId(resp.captcha_id)
      setCaptchaCode(resp.captcha_code)
    } catch (e) {
      console.log('获取验证码失败', e)
    }
  }

  useEffect(() => {
    fetchCaptcha()
  }, [])

  /**
   * 账号登录操作
   * @param values
   */
  const onSystemLogin = async (values: {
    account_name: string
    password: string
    verify_code: string
  }) => {
    const loginInfo = await SystemLoginService.systemLogin({
      account_name: values.account_name,
      password: values.password,
      verify_code: values.verify_code,
      captcha_id: captchaId
    })
    if (loginInfo) {
      setToken(loginInfo.login_token)
      setAccountInfo(loginInfo.account_info)
      message.success('登录成功！', 2, () => {
        navigate(HOME_ROOT_PATH)
      })
    } else {
      message.error('登录失败')
      fetchCaptcha()
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
      <LoginContentUI
        onSystemLogin={onSystemLogin}
        onPhoneLogin={onPhoneLogin}
        captchaCode={captchaCode}
        onRefreshCaptcha={fetchCaptcha}
      />
      <LoginFooterUI />
    </div>
  )
}

export default Login
