import { Popconfirm, Space, Table, TablePaginationConfig, Tag } from 'antd'
import { CloseSquareOutlined, FormOutlined, CheckCircleOutlined } from '@ant-design/icons'
import ActionButton from '@/components/Action/ActionButton'
import { LoginAuthInfoType, LoginAuthListItemType } from '@/types/loginAuthType'

interface LoginAuthListUIProps {
  listLoading: boolean
  loginAuthList: LoginAuthListItemType[]
  pagination: TablePaginationConfig
  onListChange: (pageConfig: TablePaginationConfig, filters: any, sorter: any) => void
  onEditClick: (loginAuthInfo: LoginAuthInfoType) => void
  onDeleteConfirm: (loginAuthInfo: LoginAuthInfoType) => void
  onUsedClick: (loginAuthInfo: LoginAuthInfoType) => void
}

const LoginAuthUsedTag = (isUsed: number) => {
  switch (isUsed) {
    case 0:
      return <Tag color="default">未使用</Tag>
    case 1:
      return <Tag color="success">使用中</Tag>
    default:
      return <Tag color="cyan">未知</Tag>
  }
}

const LoginAuthListUI = (props: LoginAuthListUIProps) => {
  return (
    <div className="panel-body">
      <Table
        rowKey={'login_auth_id'}
        bordered={true}
        dataSource={props.loginAuthList}
        loading={props.listLoading}
        pagination={props.pagination}
        onChange={props.onListChange}
        footer={() => null}
      >
        <Table.Column
          title={'认证ID'}
          dataIndex="login_auth_id"
          width={80}
          key={'login_auth_id'}
          align={'center'}
        />
        <Table.Column title={'认证名称'} dataIndex="name" key={'name'} width={150} />
        <Table.Column title={'账号前缀'} dataIndex="account_prefix" key={'account_prefix'} width={120} />
        <Table.Column title={'认证接口'} dataIndex="url" key={'url'} />
        <Table.Column
          title={'使用状态'}
          dataIndex="is_used"
          key={'is_used'}
          width={100}
          align={'center'}
          render={LoginAuthUsedTag}
        />
        <Table.Column
          title={'操作'}
          width={200}
          key={'action'}
          align={'center'}
          render={(loginAuthListItem: LoginAuthListItemType) => (
            <Space>
              <ActionButton
                text="使用"
                icon={<CheckCircleOutlined />}
                onClick={() => props.onUsedClick(loginAuthListItem)}
                havePermission={loginAuthListItem.action?.is_edit == 1}
              />
              <ActionButton
                text="修改"
                icon={<FormOutlined />}
                onClick={() => props.onEditClick(loginAuthListItem)}
                havePermission={loginAuthListItem.action?.is_edit == 1}
              />
              <Popconfirm
                title="确定要删除吗?"
                onConfirm={() => props.onDeleteConfirm(loginAuthListItem)}
                okText="确定"
                cancelText="取消"
              >
                <ActionButton
                  text="删除"
                  icon={<CloseSquareOutlined />}
                  havePermission={loginAuthListItem.action?.is_delete == 1}
                />
              </Popconfirm>
            </Space>
          )}
        />
      </Table>
    </div>
  )
}

export default LoginAuthListUI
