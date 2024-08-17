import { Descriptions, Tag } from 'antd'
import { AccountDetailResp, AccountListItemType } from '@/types/accountType'
import { AccountDepartmentFullName, AccountStatusTag } from './ToolsUI'
import { RoleTags } from '@/pages/system/Role/component/ToolsUI'

interface AccountDetailUIProps {
  accountDetail?: AccountDetailResp
}

/**
 * 账号详情 UI 组件
 */
const AccountDetailUI = (props: AccountDetailUIProps) => {
  const accountInfo = props.accountDetail
  console.log('AccountDetailUI:', accountInfo)
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
        <Descriptions.Item label="手机" span={12} key="mobile">
          {accountInfo?.mobile.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="电话" span={12} key="phone">
          {accountInfo?.phone.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="邮箱" span={12} key="email">
          {accountInfo?.email.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="部门" span={12} key="department">
          {AccountDepartmentFullName(accountInfo?.department_names)}
        </Descriptions.Item>
        <Descriptions.Item label="职位" span={12} key="position">
          {accountInfo?.position.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="工位" span={12} key="location">
          {accountInfo?.location.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="上次登录IP" span={12} key="last_ip">
          {accountInfo?.last_ip.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="上次登录时间" span={12} key="last_time">
          {accountInfo?.last_time.toString()}
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
