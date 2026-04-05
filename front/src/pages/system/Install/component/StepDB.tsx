import React, { useEffect, useState } from 'react'
import { Button, Result, Spin } from 'antd'
import { SystemInstallService } from '@/services/SystemInstall'

interface StepDBProps {
  onNext: () => void
}

const StepDB: React.FC<StepDBProps> = ({ onNext }) => {
  const [loading, setLoading] = useState<boolean>(true)
  const [success, setSuccess] = useState<boolean>(false)
  const [errorMsg, setErrorMsg] = useState<string>('')

  const checkDB = async () => {
    setLoading(true)
    try {
      const resp = await SystemInstallService.checkDB()
      if (resp.status === 'ok') {
        setSuccess(true)
      } else {
        setSuccess(false)
        setErrorMsg(resp.message || '数据库连接失败')
      }
    } catch (e: any) {
      setSuccess(false)
      setErrorMsg(e?.message || '数据库连接检测异常')
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    checkDB()
  }, [])

  if (loading) {
    return (
      <div style={{ textAlign: 'center', padding: '60px 0' }}>
        <Spin size="large" tip="正在检测数据库连接..." />
      </div>
    )
  }

  if (success) {
    return (
      <Result
        status="success"
        title="数据库连接成功"
        subTitle="数据库连接检测通过，可以进行下一步配置。"
        extra={
          <Button type="primary" onClick={onNext}>
            下一步
          </Button>
        }
      />
    )
  }

  return (
    <Result
      status="error"
      title="数据库连接失败"
      subTitle={errorMsg}
      extra={
        <Button type="primary" onClick={checkDB}>
          重新检测
        </Button>
      }
    />
  )
}

export default StepDB
