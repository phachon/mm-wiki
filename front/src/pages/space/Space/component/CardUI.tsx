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
import { SpaceInfoType } from '@/types/spaceType'

// SpaceCardProps 空间卡片组件属性
interface SpaceCardProps {
  spaceInfo: SpaceInfoType
  isCollected?: boolean
  onCollectionChange?: (spaceId: number, collected: boolean) => void
}

// SpaceCardUI 空间卡片UI组件
const SpaceCardUI: React.FC<SpaceCardProps> = (props: SpaceCardProps) => {
  const navigate = useNavigate()

  // 处理收藏点击事件
  const handleCollectionClick = (e: React.MouseEvent) => {
    e.stopPropagation()
    if (props.onCollectionChange) {
      props.onCollectionChange(props.spaceInfo.space_id, !props.isCollected)
    }
  }

  // 处理链接复制事件
  const handleCopyLink = (e: React.MouseEvent) => {
    e.stopPropagation()
    const spaceUrl = `${window.location.origin}/space/${props.spaceInfo.space_key}`
    // copyToClipboard(spaceUrl)
    message.success('空间链接已复制到剪贴板')
  }

  // 处理卡片点击事件
  const handleCardClick = () => {
    navigate(`/space/${props.spaceInfo.space_key}`)
  }

  const actions: React.ReactNode[] = [
    props.isCollected ? (
      <StarFilled onClick={handleCollectionClick} className="star-collection" />
    ) : (
      <StarOutlined onClick={handleCollectionClick} />
    ),
    <ExportOutlined onClick={handleCardClick} />
  ]

  return (
    <Card
      style={{ marginBottom: 24, cursor: 'pointer' }}
      actions={actions}
      bordered={true}
      onClick={handleCardClick}
    >
      <Card.Meta
        avatar={<Avatar src={spaceIcon} />}
        title={props.spaceInfo.name}
        description={
          <div>
            <p>{props.spaceInfo.description}</p>
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
                {props.spaceInfo.creator_name}
              </span>
              <span
                style={{
                  display: 'flex',
                  alignItems: 'center',
                  gap: '4px'
                }}
              >
                <ClockCircleOutlined style={{ fontSize: '14px' }} />
                {props.spaceInfo.create_time}
              </span>
            </div>
          </div>
        }
      />
    </Card>
  )
}

export default SpaceCardUI
