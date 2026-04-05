import React, { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { Card, Steps, message } from 'antd'
import { AUTH_LOGIN_PATH } from '@/router/type'
import { SystemInstallService } from '@/services/SystemInstall'
import StepDB from '../component/StepDB'
import StepConfig from '../component/StepConfig'
import StepAdmin from '../component/StepAdmin'
import StepComplete from '../component/StepComplete'

const stepItems = [
  { title: '数据库检测' },
  { title: '系统配置' },
  { title: '管理员账号' },
  { title: '安装完成' }
]

const InstallWizard: React.FC = () => {
  const [current, setCurrent] = useState<number>(0)
  const navigate = useNavigate()

  useEffect(() => {
    const checkStatus = async () => {
      try {
        const resp = await SystemInstallService.getInstallStatus()
        if (resp.is_installed) {
          message.info('系统已安装，即将跳转到登录页')
          navigate(AUTH_LOGIN_PATH)
        }
      } catch (e) {
        console.log('检查安装状态失败', e)
      }
    }
    checkStatus()
  }, [navigate])

  const handleConfigNext = async (values: any) => {
    try {
      await SystemInstallService.initData(values)
      message.success('系统配置保存成功')
      setCurrent(2)
    } catch (e: any) {
      message.error(e?.message || '系统配置保存失败')
    }
  }

  const handleAdminNext = async (values: any) => {
    try {
      await SystemInstallService.createAdmin(values)
      await SystemInstallService.complete()
      message.success('管理员账号创建成功')
      setCurrent(3)
    } catch (e: any) {
      message.error(e?.message || '管理员账号创建失败')
    }
  }

  const renderStep = () => {
    switch (current) {
      case 0:
        return <StepDB onNext={() => setCurrent(1)} />
      case 1:
        return <StepConfig onNext={handleConfigNext} onPrev={() => setCurrent(0)} />
      case 2:
        return <StepAdmin onNext={handleAdminNext} onPrev={() => setCurrent(1)} />
      case 3:
        return <StepComplete />
      default:
        return null
    }
  }

  return (
    <div
      style={{
        minHeight: '100vh',
        background: '#f0f2f5',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center'
      }}
    >
      <Card
        title="MM-Wiki 安装向导"
        style={{ width: '100%', maxWidth: 600 }}
        styles={{ header: { textAlign: 'center', fontSize: 20 } }}
      >
        <Steps current={current} items={stepItems} style={{ marginBottom: 32 }} />
        {renderStep()}
      </Card>
    </div>
  )
}

export default InstallWizard
