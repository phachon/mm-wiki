import React from 'react'
import { DiffEditor } from '@monaco-editor/react'
import { ArrowLeftOutlined, RedoOutlined, SwapOutlined } from '@ant-design/icons'
import { Button, Popconfirm, Space } from 'antd'
import { DocVersionEntity } from '@/types/contentType'

interface DiffViewUIProps {
  historyDocVersion?: DocVersionEntity // 历史版本内容
  onlineContent: string // 当前版本内容
  onReturnClick: () => void
  onRecoverClick: (contentVersionId?: number, docId?: number) => void
}

// DiffViewUI Diff 组件
const DiffViewUI: React.FC<DiffViewUIProps> = (props: DiffViewUIProps) => {
  console.log('DiffViewUI:', props)

  return (
    <div>
      <div className="ant-modal-title">
        <strong>
          版本差异（历史版本 <SwapOutlined /> 当前版本）
        </strong>
        <Space style={{ float: 'right' }}>
          <Button icon={<ArrowLeftOutlined />} size="small" onClick={props.onReturnClick}>
            返回
          </Button>
          <Popconfirm
            title="确定恢复至当前版本吗?"
            onConfirm={() =>
              props.onRecoverClick(
                props.historyDocVersion?.content_version_id,
                props.historyDocVersion?.doc_id
              )
            }
            okText="确定"
            cancelText="取消"
          >
            <Button icon={<RedoOutlined />} size="small" type="primary">
              恢复
            </Button>
          </Popconfirm>
        </Space>
      </div>
      <div style={{ height: '70vh', border: '1px solid #ddd', marginTop: 12 }}>
        <DiffEditor
          original={props.onlineContent}
          modified={props.historyDocVersion && props.historyDocVersion.content}
          language="markdown"
          theme="light"
          options={{
            readOnly: true,
            renderSideBySide: true,
            minimap: { enabled: true }
          }}
        />
      </div>
    </div>
  )
}

export default DiffViewUI
