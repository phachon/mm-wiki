import ActionButton from '@/components/Action/ActionButton'
import { DocVersionEntity } from '@/types/contentType'
import {
  CloseSquareOutlined,
  SelectOutlined,
  ForkOutlined,
  FormOutlined,
  RedoOutlined,
  SplitCellsOutlined,
  TeamOutlined,
  SwapOutlined
} from '@ant-design/icons'
import { Popconfirm, Space, Table, TablePaginationConfig, TableProps } from 'antd'

// DocHistoryUIProps 文档历史组件属性
type DocHistoryUIProps = {
  historyList?: DocVersionEntity[]
  pagination: TablePaginationConfig
  onViewClick?: (docVersion: DocVersionEntity) => void
  onRecoverClick?: (contentVersionId: number, docId: number) => void
  onDeleteConfirm?: (contentVersionId: number, docId: number) => void
  onListChange?: (pageConfig: TablePaginationConfig, filters: any, sorter: any) => void
}

// DocHistoryUI 文档历史组件
const DocHistoryUI = (props: DocHistoryUIProps) => {
  return (
    <div>
      <div className="ant-modal-title">
        <strong>文档历史</strong>
      </div>
      <Table
        style={{ marginTop: 12 }}
        pagination={{
          ...props.pagination,
          showSizeChanger: false
        }}
        bordered
        columns={[
          {
            title: '版本号',
            dataIndex: 'version',
            key: 'version',
            render: (text: string, record: DocVersionEntity) => (
              <strong>{record.content_version_id}</strong>
            ),
            width: 200
          },
          {
            title: '修改账号',
            dataIndex: 'account_name',
            key: 'account_name',
            render: (text: string, record: DocVersionEntity) => <a>{record.edit_account_name}</a>,
            width: 150
          },
          {
            title: '文档修改时间',
            dataIndex: 'update_time',
            key: 'update_time',
            width: 200
          },
          {
            title: '版本创建时间',
            dataIndex: 'create_time',
            key: 'create_time',
            width: 200
          },
          {
            title: '操作',
            key: 'action',
            width: 200,
            render: (text: string, record: DocVersionEntity) => (
              <Space>
                <ActionButton
                  text="Diff"
                  icon={<SwapOutlined />}
                  onClick={() => props.onViewClick && props.onViewClick(record)}
                  havePermission={true}
                />
                <Popconfirm
                  title="确定恢复至当前版本吗?"
                  onConfirm={() =>
                    props.onRecoverClick &&
                    props.onRecoverClick(record.content_version_id, record.doc_id)
                  }
                  okText="确定"
                  cancelText="取消"
                >
                  <ActionButton text="恢复" icon={<RedoOutlined />} havePermission={true} />
                </Popconfirm>
                <Popconfirm
                  title="确定删除吗?"
                  onConfirm={() =>
                    props.onDeleteConfirm &&
                    props.onDeleteConfirm(record.content_version_id, record.doc_id)
                  }
                  okText="确定"
                  cancelText="取消"
                >
                  <ActionButton text="删除" icon={<CloseSquareOutlined />} havePermission={true} />
                </Popconfirm>
              </Space>
            )
          }
        ]}
        dataSource={props.historyList}
        onChange={props.onListChange}
        footer={() => ''}
      />
    </div>
  )
}

export default DocHistoryUI
