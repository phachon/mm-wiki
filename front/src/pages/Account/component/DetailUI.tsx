import { Descriptions, Tag } from 'antd'
import { AccountListItemType } from '@/types/accountType'
import { AccountStatusTag } from './ToolsUI'
import { RoleTags } from '@/pages/Role/component/ToolsUI'

interface AccountDetailUIProps {
  accountDetail?: AccountListItemType
}

/**
 * 账号详情 UI 组件
 */
const AccountDetailUI = (props: AccountDetailUIProps) => {
  const accountInfo = props.accountDetail
  return (
    <div key={accountInfo?.account_id.toString()}>
      <Descriptions bordered size="small" column={12}>
        <Descriptions.Item label="账号ID" span={12} key="account_id">
          {accountInfo?.account_id.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="账号名" span={12} key="name">
          {accountInfo?.name.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="昵称" span={12} key="given_name">
          {accountInfo?.given_name.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="邮箱" span={12} key="email">
          {accountInfo?.email.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="电话" span={12} key="phone">
          {accountInfo?.phone.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="手机" span={12} key="mobile">
          {accountInfo?.mobile.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="角色" span={12} key="roles">
          {RoleTags(accountInfo?.roles)}
        </Descriptions.Item>
        <Descriptions.Item label="状态" span={12} key="status">
          {AccountStatusTag(accountInfo?.status)}
        </Descriptions.Item>
        <Descriptions.Item label="创建时间" span={12} key="create_time">
          {accountInfo?.create_time.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="修改时间" span={12} key="update_time">
          {accountInfo?.update_time.toString()}
        </Descriptions.Item>
      </Descriptions>
    </div>
  )
}

export default AccountDetailUI
