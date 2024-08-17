import { Popconfirm, Space, Table, TablePaginationConfig, Tag } from 'antd'
import { AccountListItemType } from '@/types/accountType'
import {
  CheckSquareOutlined,
  CloseSquareOutlined,
  FormOutlined,
  SelectOutlined
} from '@ant-design/icons'
import { AccountStatusTag } from './ToolsUI'
import { RoleTags } from '@/pages/system/Role/component/ToolsUI'
import ActionButton from '@/components/Action/ActionButton'

interface AccountListUIProps {
  /**
   * 列表 loading
   */
  listLoading: boolean

  /**
   * 账号列表数据
   */
  accountList: AccountListItemType[]

  /**
   * 翻页信息
   */
  pagination: TablePaginationConfig

  /**
   * 列表变更操作
   * @param pageConfig
   * @param filters
   * @param sorter
   * @returns
   */
  onListChange: (pageConfig: TablePaginationConfig, filters: any, sorter: any) => void

  /**
   * 详情点击操作
   * @param accountInfo
   * @returns
   */
  onDetailClick: (accountInfo: AccountListItemType) => void

  /**
   * 修改点击操作
   * @param accountInfo
   * @returns
   */
  onEditClick: (accountInfo: AccountListItemType) => void

  /**
   * 更新状态确认
   * @param accountInfo
   * @param status
   * @returns
   */
  onUpdateStatusConfirm: (accountInfo: AccountListItemType, status: number) => void
}

const AccountListUI = (props: AccountListUIProps) => {
  return (
    <div className="panel-body">
      <Table
        bordered={true}
        dataSource={props.accountList}
        loading={props.listLoading}
        pagination={props.pagination}
        onChange={props.onListChange}
        footer={() => ''}
      >
        <Table.Column
          title={'账号ID'}
          dataIndex="account_id"
          width={100}
          key={'account_id'}
          align={'center'}
        />
        <Table.Column title={'账号名'} dataIndex="name" key={'name'} width={120} />
        <Table.Column title={'昵称'} dataIndex="given_name" key={'given_name'} width={150} />
        <Table.Column title={'手机号'} dataIndex="mobile" key={'mobile'} width={200} />
        <Table.Column title={'邮箱'} dataIndex="email" key={'email'} width={220} />
        <Table.Column
          width={200}
          title={'角色'}
          dataIndex="roles"
          key={'roles'}
          render={RoleTags}
        />
        <Table.Column
          title={'状态'}
          dataIndex="status"
          width={70}
          key={'status'}
          align={'center'}
          render={AccountStatusTag}
        />
        <Table.Column
          title={'操作'}
          width={210}
          key={'action'}
          align={'center'}
          render={(accountListItem: AccountListItemType) => (
            <Space>
              <ActionButton
                text="详情"
                icon={<SelectOutlined />}
                onClick={() => props.onDetailClick(accountListItem)}
                havePermission={accountListItem.action?.is_detail == 1}
              />
              <ActionButton
                text="修改"
                icon={<FormOutlined />}
                onClick={() => props.onEditClick(accountListItem)}
                havePermission={accountListItem.action?.is_edit == 1}
              />
              {accountListItem.status === 0 && (
                <Popconfirm
                  title="确定要禁用吗?"
                  onConfirm={() => {
                    props.onUpdateStatusConfirm(accountListItem, -1)
                  }}
                  okText="确定"
                  cancelText="取消"
                >
                  <ActionButton
                    text="禁用"
                    icon={<CloseSquareOutlined />}
                    havePermission={accountListItem.action?.is_update_status == 1}
                  />
                </Popconfirm>
              )}

              {accountListItem.status === -1 && (
                <Popconfirm
                  title="确定要恢复吗?"
                  onConfirm={() => {
                    props.onUpdateStatusConfirm(accountListItem, 0)
                  }}
                  okText="确定"
                  cancelText="取消"
                >
                  <ActionButton
                    text="恢复"
                    icon={<CheckSquareOutlined />}
                    havePermission={accountListItem.action?.is_update_status == 1}
                  />
                </Popconfirm>
              )}
            </Space>
          )}
        />
      </Table>
    </div>
  )
}

export default AccountListUI
