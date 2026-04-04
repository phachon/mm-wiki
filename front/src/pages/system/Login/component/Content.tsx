import { SettingConfig } from '@/config/setting'
import { MobileOutlined, UserOutlined, LockOutlined, SafetyCertificateOutlined } from '@ant-design/icons'
import { LoginForm, ProFormCaptcha, ProFormCheckbox, ProFormText } from '@ant-design/pro-components'
import { Button, Divider, message, Tabs } from 'antd'
import { useState } from 'react'
import BannerImage from './BannerImage'
import './login.css'

type LoginType = 'phone_login' | 'system_login'

interface LoginUIProps {
  /**
   * 账号登录
   * @returns
   */
  onSystemLogin: (values: any) => void

  /**
   * 手机号登录
   * @returns
   */
  onPhoneLogin: (values: any) => Promise<boolean | void>

  /**
   * 验证码
   */
  captchaCode?: string

  /**
   * 刷新验证码
   */
  onRefreshCaptcha?: () => void
}

const LoginContentUI = (props: LoginUIProps) => {
  const [loginType, setLoginType] = useState<LoginType>('system_login')
  return (
    <div className="login-container">
      <div className="login-img">
        <BannerImage />
      </div>
      <div className="login-form">
        <LoginForm
          logo={<img src={SettingConfig.frameHeaderIconUrl} />}
          title="MM-Wiki"
          subTitle="轻量的企业知识分享与团队协同软件"
          actions={
            <div
              style={{
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                flexDirection: 'column'
              }}
            >
              <Divider plain>
                <span style={{ color: '#CCC', fontWeight: 'normal', fontSize: 12 }}>
                  推荐使用 Chrome 或 FireFox 浏览器访问
                </span>
              </Divider>
            </div>
          }
          isKeyPressSubmit={true}
          onFinish={loginType == 'system_login' ? props.onSystemLogin : props.onPhoneLogin}
        >
          <Tabs
            centered
            activeKey={loginType}
            onChange={(activeKey) => setLoginType(activeKey as LoginType)}
          >
            <Tabs.TabPane key={'system_login'} tab={'账号密码登录'} />
            <Tabs.TabPane key={'phone_login'} tab={'手机号登录'} />
          </Tabs>
          {loginType === 'system_login' && (
            <>
              <ProFormText
                name="account_name"
                fieldProps={{
                  size: 'large',
                  prefix: <UserOutlined className={'prefixIcon'} />
                }}
                placeholder={'请输入登录账号名'}
                rules={[
                  {
                    required: true,
                    message: '请输入登录账号名！'
                  }
                ]}
              />
              <ProFormText.Password
                name="password"
                fieldProps={{
                  size: 'large',
                  prefix: <LockOutlined className={'prefixIcon'} />
                }}
                placeholder={'请输入登录密码！'}
                rules={[
                  {
                    required: true,
                    message: '请输入登录密码！'
                  }
                ]}
              />
              <div style={{ display: 'flex', gap: 8, alignItems: 'flex-start' }}>
                <div style={{ flex: 1 }}>
                  <ProFormText
                    name="verify_code"
                    fieldProps={{
                      size: 'large',
                      prefix: <SafetyCertificateOutlined className={'prefixIcon'} />
                    }}
                    placeholder={'请输入验证码'}
                    rules={[
                      {
                        required: true,
                        message: '请输入验证码！'
                      }
                    ]}
                  />
                </div>
                <Button
                  size="large"
                  style={{
                    minWidth: 100,
                    height: 40,
                    fontSize: 16,
                    fontWeight: 'bold',
                    letterSpacing: 4
                  }}
                  onClick={props.onRefreshCaptcha}
                >
                  {props.captchaCode || '----'}
                </Button>
              </div>
            </>
          )}
          {loginType === 'phone_login' && (
            <>
              <ProFormText
                fieldProps={{
                  size: 'large',
                  prefix: <MobileOutlined className={'prefixIcon'} />
                }}
                name="mobile"
                placeholder={'请输入手机号'}
                rules={[
                  {
                    required: true,
                    message: '请输入手机号！'
                  },
                  {
                    pattern: /^1\d{10}$/,
                    message: '手机号格式错误！'
                  }
                ]}
              />
              <ProFormCaptcha
                fieldProps={{
                  size: 'large',
                  prefix: <LockOutlined className={'prefixIcon'} />
                }}
                captchaProps={{
                  size: 'large'
                }}
                placeholder={'请输入验证码'}
                captchaTextRender={(timing, count) => {
                  if (timing) {
                    return `${count} ${'获取验证码'}`
                  }
                  return '获取验证码'
                }}
                name="captcha"
                rules={[
                  {
                    required: true,
                    message: '请输入验证码！'
                  }
                ]}
                onGetCaptcha={async () => {
                  message.success('获取验证码成功！验证码为：1234')
                }}
              />
            </>
          )}
          <div
            style={{
              marginBlockEnd: 24
            }}
          >
            <ProFormCheckbox noStyle name="auto_login" initialValue={true}>
              自动登录
            </ProFormCheckbox>
            <a
              style={{
                float: 'right'
              }}
            >
              忘记密码
            </a>
          </div>
        </LoginForm>
      </div>
    </div>
  )
}

export default LoginContentUI
