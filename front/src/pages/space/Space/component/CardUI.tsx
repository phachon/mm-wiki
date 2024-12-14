import React from 'react'
import { Avatar, Card, message } from 'antd'
import {
  ExportOutlined,
  StarOutlined,
  StarFilled,
  UserOutlined,
  ClockCircleOutlined
} from '@ant-design/icons'
// @ts-ignore
import spaceIcon from '@/assets/images/icons-space.png'
import { useNavigate } from 'react-router-dom'
import './space.css'

// SpaceCardProps 空间卡片组件属性
interface SpaceCardProps {
  title: string
  description: string
  creator: string
  createTime: string
  spaceKey: string
  isCollected?: boolean
  onCollectionChange?: (spaceKey: string, collected: boolean) => void
  avatarSrc?: string
}

// SpaceCardUI 空间卡片UI组件
const SpaceCardUI: React.FC<SpaceCardProps> = ({
  title,
  description,
  creator,
  createTime,
  spaceKey,
  isCollected = false,
  onCollectionChange,
  avatarSrc = spaceIcon
}) => {
  const navigate = useNavigate()

  // 处理收藏点击事件
  const handleCollectionClick = (e: React.MouseEvent) => {
    e.stopPropagation()
    if (onCollectionChange) {
      onCollectionChange(spaceKey, !isCollected)
    }
  }

  // 处理链接复制事件
  const handleCopyLink = (e: React.MouseEvent) => {
    e.stopPropagation()
    const spaceUrl = `${window.location.origin}/space/${spaceKey}`
    // copyToClipboard(spaceUrl)
    message.success('空间链接已复制到剪贴板')
  }

  // 处理卡片点击事件
  const handleCardClick = () => {
    navigate(`/space/${spaceKey}`)
  }

  const actions: React.ReactNode[] = [
    isCollected ? (
      <StarFilled key="collection" onClick={handleCollectionClick} style={{ color: '#faad14' }} />
    ) : (
      <StarOutlined key="collection" onClick={handleCollectionClick} />
    ),
    <ExportOutlined key="link" onClick={handleCardClick} />
  ]

  return (
    <Card
      style={{ marginBottom: 24, cursor: 'pointer' }}
      actions={actions}
      bordered={true}
      onClick={handleCardClick}
    >
      <Card.Meta
        avatar={<Avatar src={avatarSrc} />}
        title={title}
        description={
          <div>
            <p>{description}</p>
            <div
              style={{
                fontSize: '12px',
                color: '#8c8c8c',
                display: 'flex',
                gap: '16px',
                marginTop: '12px'
              }}
            >
              <span
                style={{
                  display: 'flex',
                  alignItems: 'center',
                  gap: '4px'
                }}
              >
                <UserOutlined style={{ fontSize: '14px' }} />
                {creator}
              </span>
              <span
                style={{
                  display: 'flex',
                  alignItems: 'center',
                  gap: '4px'
                }}
              >
                <ClockCircleOutlined style={{ fontSize: '14px' }} />
                {createTime}
              </span>
            </div>
          </div>
        }
      />
    </Card>
  )
}

export default SpaceCardUI
