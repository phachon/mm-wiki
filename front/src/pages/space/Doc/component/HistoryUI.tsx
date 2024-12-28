import ActionButton from '@/components/Action/ActionButton'
import { DocVersionEntity } from '@/types/contentType'
import {
  CloseSquareOutlined,
  ForkOutlined,
  FormOutlined,
  RedoOutlined,
  SplitCellsOutlined,
  TeamOutlined
} from '@ant-design/icons'
import { Popconfirm, Space, Table, TablePaginationConfig } from 'antd'

// DocHistoryUIProps 文档历史组件属性
type DocHistoryUIProps = {
  historyList?: DocVersionEntity[]
  pagination: TablePaginationConfig
}

// DocHistoryUI 文档历史组件
const DocHistoryUI = (props: DocHistoryUIProps) => {
  return (
    <div>
      <Table
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
              <a>{'V' + record.doc_id + '.' + record.content_version_id}</a>
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
            render: (text: string, record: any) => (
              <Space>
                <ActionButton
                  text="恢复"
                  icon={<RedoOutlined />}
                  // onClick={() =>
                  //   // props.onAdminListClick ? props.onAdminListClick(spaceListItem) : null
                  // }
                  havePermission={true}
                />
                <Popconfirm
                  title="确定删除吗?"
                  // onConfirm={() => props.onDeleteConfirm(spaceListItem)}
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
      />
    </div>
  )
}

export default DocHistoryUI
