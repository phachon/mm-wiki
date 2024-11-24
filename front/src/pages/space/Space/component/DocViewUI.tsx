import React, { useEffect, useRef, useState } from 'react'
import { Button, Col, Divider, Row, Space } from 'antd'
import {
  FolderOutlined,
  CalendarOutlined,
  HistoryOutlined,
  PaperClipOutlined,
  EditOutlined,
  StarOutlined,
  ShareAltOutlined,
  ExportOutlined
} from '@ant-design/icons'
import ButtonGroup from 'antd/es/button/button-group'
import 'cherry-markdown/dist/cherry-markdown.css'
import Cherry from 'cherry-markdown'

type SpaceDocViewProps = {
  content: string
}

const SpaceDocViewUI = (props: SpaceDocViewProps) => {
  const editorRef = useRef<HTMLDivElement>(null)

  useEffect(() => {
    if (editorRef.current) {
      const cherry = new Cherry({
        el: editorRef.current,
        value: props.content,
        editor: {
          defaultModel: 'previewOnly', // 仅预览模式
          keepDocumentScrollAfterInit: true
        },
        toolbars: {
          showToolbar: false,
          toolbar: [],
          hiddenToolbar: [],
          // 配置目录
          toc: {
            updateLocationHash: true, // 要不要更新URL的hash
            defaultModel: 'full', // pure: 精简模式/缩略模式，只有一排小点； full: 完整模式，会展示所有标题
            position: 'fixed', // 悬浮目录的悬浮方式。当滚动条在cherry内部时，用absolute；当滚动条在cherry外部时，用fixed
            cssText: 'right: 12px;'
          }
        }
      })

      return () => {
        cherry.destroy() // 在组件卸载时销毁 Cherry 实例
      }
    }
  }, [props.content])

  return (
    <div className="doc-view">
      <div className="doc-view-header">
        <Row gutter={24}>
          <Col span={14}>
            <h3 className="doc-view-page-title">测试空间的标题</h3>
            <p className="doc-view-page-path">
              <Space>
                <FolderOutlined />
                <a className="text text-info">开发部门</a>/ <a>开发部门</a>
              </Space>
            </p>
            <p className="doc-view-page-time">
              <Space>
                <CalendarOutlined />
                <a>xxx（root）</a>
                创建于 2024-17-08 19:09:08，
                <a>xxx（root）</a>
                更新于 2024-17-08 19:09:08
                <a data-link="/document/history?document_id=x">
                  <Space>
                    <HistoryOutlined />
                    查看修改历史
                  </Space>
                </a>
                <a>
                  <Space>
                    <PaperClipOutlined />
                    查看附件
                  </Space>
                </a>
              </Space>
            </p>
          </Col>
          <Col span={10} style={{ textAlign: 'right' }}>
            <div className="doc-view-header-actions">
              <ButtonGroup>
                <Button type="default" icon={<EditOutlined />} href="/page/edit?document_id=">
                  编辑
                </Button>
                <Button type="default" icon={<StarOutlined />}>
                  收藏
                </Button>
                <Button type="default" icon={<StarOutlined />}>
                  取消
                </Button>
                <Button type="default" icon={<ShareAltOutlined />}>
                  分享
                </Button>
                <Button type="default" icon={<ExportOutlined />}>
                  导出
                </Button>
              </ButtonGroup>
            </div>
          </Col>
        </Row>
        <Divider className="doc-view-divider" />
      </div>
      <div className="doc-view-body" ref={editorRef}></div>
    </div>
  )
}

export default SpaceDocViewUI
