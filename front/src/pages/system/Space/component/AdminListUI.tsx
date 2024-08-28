import { Button, Form, Input, Popconfirm, Select, Table, TablePaginationConfig } from 'antd'
import { AccountInfoType } from '@/types/accountType'
import { CloseCircleOutlined } from '@ant-design/icons'
import ActionButton from '@/components/Action/ActionButton'
import { useEffect } from 'react'

/**
 * 管理员账号列表 UI 组件 props
 */
interface SpaceAdminListUIProps {
  adminList?: AccountInfoType[] // 管理员列表数据
  selectedAccountList?: AccountInfoType[] // 账号列表数据
  onRemoveAdminChange?: (accountInfo: AccountInfoType) => void // 移除账号操作方法
  onAddAdminChange?: (values: any) => void // 添加账号操作方法
}

/**
 * 账号列表 UI 组件
 * @param props 组件依赖数据
 */
const SpaceAdminListUI = (props: SpaceAdminListUIProps) => {
  const [form] = Form.useForm()
  form.resetFields()
  return (
    <div>
      <div className="search-container" style={{ marginBottom: 16 }}>
        <Form
          layout={'inline'}
          style={{ justifyContent: 'end' }}
          onFinish={props.onAddAdminChange}
          form={form}
        >
          <Form.Item
            label="选择管理员"
            name="admin_account_ids"
            rules={[{ required: true, message: '请选择空间管理员' }]}
          >
            <Select
              mode="multiple"
              placeholder="请选择空间管理员"
              showSearch={true}
              style={{ width: 250 }}
              filterOption={(input: any, option: any) => {
                const optionText = option.children.toLowerCase()
                const optionKey = option.key.toLowerCase()
                return (
                  optionText.indexOf(input.toLowerCase()) >= 0 ||
                  optionKey.indexOf(input.toLowerCase()) >= 0
                )
              }}
            >
              {props.selectedAccountList &&
                props.selectedAccountList.map((accountInfo) => (
                  <Select.Option
                    id={accountInfo.account_id}
                    key={accountInfo.name}
                    value={accountInfo.account_id}
                  >
                    {accountInfo.name + '(' + accountInfo.given_name + ')'}
                  </Select.Option>
                ))}
            </Select>
          </Form.Item>

          <Form.Item style={{ margin: 0 }}>
            <Button type="primary" htmlType="submit">
              添加
            </Button>
          </Form.Item>
        </Form>
      </div>

      <Table
        rowKey={'account_id'}
        bordered={true}
        dataSource={props.adminList}
        footer={() => ''}
        pagination={false}
      >
        <Table.Column
          title={'账号ID'}
          dataIndex="account_id"
          width={80}
          key={'account_id'}
          align={'center'}
        />
        <Table.Column title={'账号名'} dataIndex="name" key={'name'} />
        <Table.Column title={'昵称'} dataIndex="given_name" key={'given_name'} />
        <Table.Column
          title={'创建时间'}
          dataIndex="create_time"
          key={'create_time'}
          align={'center'}
        />
        <Table.Column
          title={'修改时间'}
          dataIndex="update_time"
          key={'update_time'}
          align={'center'}
        />
        <Table.Column
          title={'操作'}
          key={'action'}
          align={'center'}
          render={(accountInfo: AccountInfoType) => (
            <span>
              <Popconfirm
                title="确定要移除吗?"
                onConfirm={() =>
                  props.onRemoveAdminChange ? props.onRemoveAdminChange(accountInfo) : undefined
                }
                okText="确定"
                cancelText="取消"
              >
                <ActionButton text="移除" icon={<CloseCircleOutlined />} havePermission={true} />
              </Popconfirm>
            </span>
          )}
        />
      </Table>
    </div>
  )
}

export default SpaceAdminListUI
