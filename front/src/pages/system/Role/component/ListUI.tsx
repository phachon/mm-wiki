import { Button, Popconfirm, Space, Table, TablePaginationConfig } from 'antd'
import { RoleInfoType, RoleListItemType, RoleTypeAccountRootRole } from '@/types/roleType'
import { CloseSquareOutlined, FormOutlined, TeamOutlined, LockOutlined } from '@ant-design/icons'
import { RoleTypeTagUI } from './ToolsUI'
import ActionButton from '@/components/Action/ActionButton'

interface RoleListUIProps {
  /**
   * 表格 loading
   */
  listLoading: boolean
  /**
   * 角色列表数据
   */
  roleList: RoleListItemType[]
  /**
   * 分页信息
   */
  pagination: TablePaginationConfig
  /**
   * 分页操作方法
   * @param pageConfig
   * @param filters
   * @param sorter
   */
  onListChange: (pageConfig: TablePaginationConfig, filters: any, sorter: any) => void
  /**
   * 修改点击操作方法
   * @param roleInfo
   */
  onEditClick: (roleInfo: RoleInfoType) => void
  /**
   * 删除操作方法
   * @param roleInfo
   * @returns
   */
  onDeleteConfirm: (roleInfo: RoleInfoType) => void
  /**
   * 账号列表点击方法
   * @param roleInfo
   */
  onAccountListClick?: (roleInfo: RoleInfoType) => void
  /**
   * 权限列表点击操作方法
   * @param roleInfo
   */
  onPrivilegeListClick?: (roleInfo: RoleInfoType) => void
}

/**
 * 角色列表 UI 组件
 * @param props
 * @returns
 */
const RoleListUI = (props: RoleListUIProps) => {
  return (
    <div className="panel-body">
      <Table
        rowKey={'role_id'}
        bordered={true}
        dataSource={props.roleList}
        loading={props.listLoading}
        pagination={props.pagination}
        onChange={props.onListChange}
        footer={() => null}
      >
        <Table.Column
          title={'角色ID'}
          dataIndex="role_id"
          width={80}
          key={'role_id'}
          align={'center'}
        />
        <Table.Column title={'角色名'} dataIndex="name" key={'name'} width={250} />
        <Table.Column title={'备注'} dataIndex="remark" key={'remark'} />
        <Table.Column
          title={'角色类型'}
          dataIndex="role_type"
          key={'role_type'}
          width={120}
          align={'center'}
          render={RoleTypeTagUI}
        />
        <Table.Column
          title={'修改时间'}
          dataIndex="update_time"
          key={'update_time'}
          width={200}
          align={'center'}
        />
        <Table.Column
          title={'操作'}
          width={260}
          key={'action'}
          align={'center'}
          render={(roleListItem: RoleListItemType) => (
            <Space>
              <ActionButton
                text="账号"
                icon={<TeamOutlined />}
                onClick={() =>
                  props.onAccountListClick ? props.onAccountListClick(roleListItem) : null
                }
                havePermission={roleListItem.action?.is_account_list == 1}
              />
              {roleListItem.role_type == RoleTypeAccountRootRole ? (
                <ActionButton
                  text="权限"
                  icon={<LockOutlined />}
                  havePermission={false}
                  tooltipTitle="超级管理员默认拥有所有权限"
                />
              ) : (
                <ActionButton
                  text="权限"
                  icon={<LockOutlined />}
                  onClick={() =>
                    props.onPrivilegeListClick ? props.onPrivilegeListClick(roleListItem) : null
                  }
                  havePermission={roleListItem.action?.is_privilege_edit == 1}
                />
              )}

              <ActionButton
                text="修改"
                icon={<FormOutlined />}
                onClick={() => props.onEditClick(roleListItem)}
                havePermission={roleListItem.action?.is_edit == 1}
              />

              <Popconfirm
                title="确定要删除吗?"
                onConfirm={() => props.onDeleteConfirm(roleListItem)}
                okText="确定"
                cancelText="取消"
              >
                <ActionButton
                  text="删除"
                  icon={<CloseSquareOutlined />}
                  havePermission={roleListItem.action?.is_delete == 1}
                />
              </Popconfirm>
            </Space>
          )}
        />
      </Table>
    </div>
  )
}

export default RoleListUI
