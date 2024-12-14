import { SaveOutlined, RollbackOutlined } from '@ant-design/icons'
import { Button, Col, Form, Input, Row, Space } from 'antd'
import { useEffect, useRef, useState } from 'react'
import { ContentEntity, DocEntity } from '@/types/docType'
import 'cherry-markdown/dist/cherry-markdown.css'
import Cherry from 'cherry-markdown'
import './edit.css'

// DocEditUIProps 文档编辑组件属性
type DocEditUIProps = {
  onSaveSubmit?: (values: any) => void
  docInfo?: DocEntity
  content?: ContentEntity
}

// DocEditUIProps 文档编辑组件
const DocEditUI = (props: DocEditUIProps) => {
  const docEditRef = useRef<HTMLDivElement>(null)
  const [name, setName] = useState(props.docInfo?.name || '')
  const [content, setContent] = useState(props.content?.content || '')

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
      editor: {
        id: 'doc-content', // textarea 的id属性值
        name: 'content', // textarea 的name属性值
        autoSave2Textarea: true, // 是否自动将编辑区的内容回写到textarea里
        defaultModel: 'edit&preview',
        keepDocumentScrollAfterInit: true
      },
      callback: {
        afterChange: editorChange
      },
      previewer: {
        enablePreviewerBubble: false
      },
      toolbars: {
        showToolbar: true,
        toolbar: [
          'bold',
          'italic',
          'strikethrough',
          '|',
          'color',
          'header',
          'ruby',
          '|',
          'list',
          'panel',
          // 'justify', // 对齐方式，默认不推荐这么“复杂”的样式要求
          'detail',
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
              'toc',
              'table',
              'line-table',
              'bar-table',
              'pdf',
              'word'
            ]
          },
          'graph',
          'settings'
        ]
      },
      // cherry初始化后是否检查 location.hash 尝试滚动到对应位置
      autoScrollByHashAfterInit: true,
      themeSettings: {
        // 主题列表，用于切换主题
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
        toolbarTheme: 'dark' // light or dark 优先级低于mainTheme
      }
    })
    return () => {
      cherry.destroy() // 在组件卸载时销毁 Cherry 实例
    }
  }, [props.content])

  const handleSave = () => {
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

  return (
    <div style={{ padding: '16px 14px 14px 14px' }}>
      <Row>
        <Col span={20}>
          <Input name="name" placeholder="文档标题" value={name} onChange={handleTitleChange} />
        </Col>
        <Col span={4} style={{ textAlign: 'right' }}>
          <Space>
            <Button type="primary" icon={<SaveOutlined />} onClick={handleSave}>
              保存
            </Button>
            <Button type="default" icon={<RollbackOutlined />}>
              取消
            </Button>
          </Space>
        </Col>
      </Row>
      <div className="doc-edit-body" ref={docEditRef}></div>
    </div>
  )
}

export default DocEditUI
