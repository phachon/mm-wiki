import { Popconfirm, Space, Table, TablePaginationConfig } from 'antd'
import { CloseSquareOutlined, FormOutlined } from '@ant-design/icons'
import ActionButton from '@/components/Action/ActionButton'
import { ContactInfoType, ContactListItemType } from '@/types/contactType'

interface ContactListUIProps {
  listLoading: boolean
  contactList: ContactListItemType[]
  pagination: TablePaginationConfig
  onListChange: (pageConfig: TablePaginationConfig, filters: any, sorter: any) => void
  onEditClick: (contactInfo: ContactInfoType) => void
  onDeleteConfirm: (contactInfo: ContactInfoType) => void
}

const ContactListUI = (props: ContactListUIProps) => {
  return (
    <div className="panel-body">
      <Table
        rowKey={'contact_id'}
        bordered={true}
        dataSource={props.contactList}
        loading={props.listLoading}
        pagination={props.pagination}
        onChange={props.onListChange}
        footer={() => null}
      >
        <Table.Column
          title={'联系人ID'}
          dataIndex="contact_id"
          width={100}
          key={'contact_id'}
          align={'center'}
        />
        <Table.Column title={'联系人名称'} dataIndex="name" key={'name'} width={150} />
        <Table.Column title={'联系电话'} dataIndex="mobile" key={'mobile'} width={150} />
        <Table.Column title={'邮箱'} dataIndex="email" key={'email'} width={200} />
        <Table.Column title={'职位'} dataIndex="position" key={'position'} />
        <Table.Column
          title={'修改时间'}
          dataIndex="update_time"
          key={'update_time'}
          width={180}
          align={'center'}
        />
        <Table.Column
          title={'操作'}
          width={140}
          key={'action'}
          align={'center'}
          render={(contactListItem: ContactListItemType) => (
            <Space>
              <ActionButton
                text="修改"
                icon={<FormOutlined />}
                onClick={() => props.onEditClick(contactListItem)}
                havePermission={contactListItem.action?.is_edit == 1}
              />
              <Popconfirm
                title="确定要删除吗?"
                onConfirm={() => props.onDeleteConfirm(contactListItem)}
                okText="确定"
                cancelText="取消"
              >
                <ActionButton
                  text="删除"
                  icon={<CloseSquareOutlined />}
                  havePermission={contactListItem.action?.is_delete == 1}
                />
              </Popconfirm>
            </Space>
          )}
        />
      </Table>
    </div>
  )
}

export default ContactListUI
