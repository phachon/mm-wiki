import { Popconfirm, Space, Table, TablePaginationConfig, Tag } from 'antd'
import { CloseSquareOutlined, FormOutlined, CheckCircleOutlined, SendOutlined } from '@ant-design/icons'
import ActionButton from '@/components/Action/ActionButton'
import { EmailInfoType, EmailListItemType } from '@/types/emailType'

interface EmailListUIProps {
  listLoading: boolean
  emailList: EmailListItemType[]
  pagination: TablePaginationConfig
  onListChange: (pageConfig: TablePaginationConfig, filters: any, sorter: any) => void
  onEditClick: (emailInfo: EmailInfoType) => void
  onDeleteConfirm: (emailInfo: EmailInfoType) => void
  onUsedClick: (emailInfo: EmailInfoType) => void
  onTestClick: (emailInfo: EmailInfoType) => void
}

const EmailUsedTag = (isUsed: number) => {
  switch (isUsed) {
    case 0:
      return <Tag color="default">未使用</Tag>
    case 1:
      return <Tag color="success">使用中</Tag>
    default:
      return <Tag color="cyan">未知</Tag>
  }
}

const EmailListUI = (props: EmailListUIProps) => {
  return (
    <div className="panel-body">
      <Table
        rowKey={'email_id'}
        bordered={true}
        dataSource={props.emailList}
        loading={props.listLoading}
        pagination={props.pagination}
        onChange={props.onListChange}
        footer={() => null}
      >
        <Table.Column
          title={'邮箱ID'}
          dataIndex="email_id"
          width={80}
          key={'email_id'}
          align={'center'}
        />
        <Table.Column title={'服务器名称'} dataIndex="name" key={'name'} width={150} />
        <Table.Column title={'发件人地址'} dataIndex="sender_address" key={'sender_address'} width={200} />
        <Table.Column title={'服务器主机'} dataIndex="host" key={'host'} width={150} />
        <Table.Column title={'端口'} dataIndex="port" key={'port'} width={80} align={'center'} />
        <Table.Column
          title={'使用状态'}
          dataIndex="is_used"
          key={'is_used'}
          width={100}
          align={'center'}
          render={EmailUsedTag}
        />
        <Table.Column
          title={'修改时间'}
          dataIndex="update_time"
          key={'update_time'}
          width={180}
          align={'center'}
        />
        <Table.Column
          title={'操作'}
          width={260}
          key={'action'}
          align={'center'}
          render={(emailListItem: EmailListItemType) => (
            <Space>
              <ActionButton
                text="测试"
                icon={<SendOutlined />}
                onClick={() => props.onTestClick(emailListItem)}
                havePermission={emailListItem.action?.is_edit == 1}
              />
              <ActionButton
                text="使用"
                icon={<CheckCircleOutlined />}
                onClick={() => props.onUsedClick(emailListItem)}
                havePermission={emailListItem.action?.is_edit == 1}
              />
              <ActionButton
                text="修改"
                icon={<FormOutlined />}
                onClick={() => props.onEditClick(emailListItem)}
                havePermission={emailListItem.action?.is_edit == 1}
              />
              <Popconfirm
                title="确定要删除吗?"
                onConfirm={() => props.onDeleteConfirm(emailListItem)}
                okText="确定"
                cancelText="取消"
              >
                <ActionButton
                  text="删除"
                  icon={<CloseSquareOutlined />}
                  havePermission={emailListItem.action?.is_delete == 1}
                />
              </Popconfirm>
            </Space>
          )}
        />
      </Table>
    </div>
  )
}

export default EmailListUI
