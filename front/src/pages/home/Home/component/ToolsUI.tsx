import { Divider, Space } from 'antd'
import {
  RightOutlined,
  LeftOutlined,
  EditOutlined,
  HomeOutlined,
  TeamOutlined,
  FolderOutlined,
  ForwardOutlined,
  VerticalLeftOutlined,
  FolderOpenOutlined,
  FileOutlined,
  FileTextOutlined
} from '@ant-design/icons'
import DynamicIcon from '@/components/DynamicIcon/DynamicIcon'
import React from 'react'
import { SpaceInfoType } from '@/types/spaceType'
import { DocEntity, DocType } from '@/types/docType'

// 带分割线的文字
export const TextDeliver = (props: { name: string; icon?: React.ReactNode }) => {
  return (
    <div
      style={{
        fontSize: 14,
        paddingLeft: 8,
        paddingRight: 8,
        marginTop: 12,
        marginBottom: 0,
        color: '#939393'
      }}
    >
      <Space style={{ marginLeft: 0, fontSize: 15 }}>
        <RightOutlined />
        {props.icon}
        <span style={{ fontWeight: 500 }}>{props.name}</span>
      </Space>
      <Divider
        style={{
          marginTop: 12,
          marginBottom: 0,
          borderWidth: 1,
          borderColor: '#e9e9e9'
        }}
      />
    </div>
  )
}

export const SpaceList = (props: { items: SpaceInfoType[] }) => {
  return (
    <div>
      {props.items.map((item) => (
        <div className="sidebar-space-item">
          <Space>
            <DynamicIcon name={'FolderOpenOutlined'} />
            <a href={'/space/' + item.space_key}>{item.name}</a>
          </Space>
        </div>
      ))}
    </div>
  )
}

export const DocList = (props: { items: DocEntity[] }) => {
  return (
    <div>
      {props.items.map((item) => (
        <div className="sidebar-space-item">
          <Space>
            {item.type == DocType.DOC ? <FileTextOutlined /> : <FileTextOutlined />}
            <a href={'/doc/' + item.doc_id}>{item.name}</a>
          </Space>
        </div>
      ))}
    </div>
  )
}
