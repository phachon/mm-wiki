import { SaveOutlined, RollbackOutlined } from '@ant-design/icons'
import { Button, Col, Form, Input, message, Popconfirm, Row, Space } from 'antd'
import { useEffect, useRef, useState } from 'react'
import { DocEntity } from '@/types/docType'
import { ContentEntity } from '@/types/contentType'
import 'cherry-markdown/dist/cherry-markdown.css'
import Cherry from 'cherry-markdown'
import { CherryOptions } from 'cherry-markdown/types/cherry'
import './edit.css'
import { SettingConfig } from '@/config/setting'
import { DocUrlProcessor, navigateDocView } from './ToolsUI'
import { useNavigate } from 'react-router-dom'

// DocEditUIProps 文档编辑组件属性
type DocEditUIProps = {
  docInfo?: DocEntity
  content?: ContentEntity
  onSaveSubmit?: (values: any) => void
  onFileUpload?: (file: any, callback: any, docId?: number) => void
}

const editCherryConfig: CherryOptions = {
  externals: {
    // echarts: window.echarts,
    // katex: window.katex,
    // MathJax: window.MathJax
  },
  editor: {
    defaultModel: 'edit&preview',
    keepDocumentScrollAfterInit: true
  },
  previewer: {
    enablePreviewerBubble: false
  },
  drawioIframeUrl: '/doc/draw_io',
  toolbars: {
    showToolbar: true,
    toolbar: [
      'bold',
      'italic',
      {
        strikethrough: ['strikethrough', 'underline', 'sub', 'sup', 'ruby']
      },
      'size',
      '|',
      'color',
      'header',
      'hr',
      'ruby',
      '|',
      'list',
      'table',
      'justify',
      'panel',
      'detail',
      '|',
      {
        insert: [
          'image',
          'audio',
          'video',
          'link',
          'hr',
          'br',
          'code',
          'formula',
          'table',
          'line-table',
          'bar-table',
          'toc',
          'pdf',
          'word'
        ]
      },
      'graph',
      'formula',
      '|',
      'codeTheme',
      'search',
      'settings'
    ],
    toolbarRight: ['fullScreen', 'export', 'wordCount'],
    bubble: [
      'bold',
      'italic',
      'underline',
      'strikethrough',
      'sub',
      'sup',
      'quote',
      'ruby',
      '|',
      'size',
      'color'
    ],
    sidebar: ['mobilePreview', 'copy', 'theme']
  },
  // cherry初始化后是否检查 location.hash 尝试滚动到对应位置
  autoScrollByHashAfterInit: true,
  // 主题列表，用于切换主题
  themeSettings: {
    themeList: [
      { className: 'default', label: '默认' },
      { className: 'dark', label: '暗黑' },
      { className: 'light', label: '明亮' },
      { className: 'green', label: '清新' },
      { className: 'red', label: '热情' },
      { className: 'violet', label: '淡雅' },
      { className: 'blue', label: '清幽' }
    ],
    mainTheme: 'light',
    codeBlockTheme: 'default',
    inlineCodeTheme: 'red', // red or black
    toolbarTheme: 'light' // light or dark 优先级低于mainTheme
  }
}

// DocEditUIProps 文档编辑组件
const DocEditUI = (props: DocEditUIProps) => {
  const docEditRef = useRef<HTMLDivElement>(null)
  const [name, setName] = useState(props.docInfo?.name || '')
  const [content, setContent] = useState(props.content?.content || '')
  const navigate = useNavigate()

  useEffect(() => {
    setName(props.docInfo?.name || '')
  }, [props.docInfo])

  useEffect(() => {
    setContent(props.content?.content || '')
  }, [props.content])

  useEffect(() => {
    if (!docEditRef.current) {
      return
    }
    const cherry = new Cherry({
      el: docEditRef.current,
      value: props.content?.content,
      callback: {
        fileUpload: (file: any, callback: any) => {
          props.onFileUpload && props.onFileUpload(file, callback, props.docInfo?.doc_id)
        },
        urlProcessor: DocUrlProcessor,
        afterChange: editorChange
      },
      ...editCherryConfig
    })
    return () => {
      cherry.destroy() // 在组件卸载时销毁 Cherry 实例
    }
  }, [props.content])

  const onSaveClick = () => {
    if (!props.onSaveSubmit) {
      return
    }
    props.onSaveSubmit({
      name: name,
      content: content,
      doc_id: props.docInfo?.doc_id
    })
  }

  const editorChange = (md: string, html: string) => {
    setContent(md)
  }

  const handleTitleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setName(e.target.value)
  }

  const onCancelConfirm = () => {
    navigateDocView(navigate, props.docInfo?.doc_id)
  }

  return (
    <div style={{ padding: '16px 14px 14px 14px' }}>
      <Row>
        <Col span={20}>
          <Input name="name" placeholder="文档标题" value={name} onChange={handleTitleChange} />
        </Col>
        <Col span={4} style={{ textAlign: 'right' }}>
          <Space>
            <Button type="primary" icon={<SaveOutlined />} onClick={onSaveClick}>
              保存
            </Button>
            <Popconfirm
              title="确定要取消编辑吗?"
              description="未保存的内容将会丢失"
              onConfirm={onCancelConfirm}
              onCancel={() => {}}
              okText="确定取消"
              cancelText="继续编辑"
            >
              <Button type="default" icon={<RollbackOutlined />}>
                取消
              </Button>
            </Popconfirm>
          </Space>
        </Col>
      </Row>
      <div className="doc-edit-body" ref={docEditRef}></div>
    </div>
  )
}

export default DocEditUI
