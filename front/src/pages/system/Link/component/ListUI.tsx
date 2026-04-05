import { Popconfirm, Space, Table, TablePaginationConfig } from 'antd'
import { CloseSquareOutlined, FormOutlined } from '@ant-design/icons'
import ActionButton from '@/components/Action/ActionButton'
import { LinkInfoType, LinkListItemType } from '@/types/linkType'

interface LinkListUIProps {
  listLoading: boolean
  linkList: LinkListItemType[]
  pagination: TablePaginationConfig
  onListChange: (pageConfig: TablePaginationConfig, filters: any, sorter: any) => void
  onEditClick: (linkInfo: LinkInfoType) => void
  onDeleteConfirm: (linkInfo: LinkInfoType) => void
}

const LinkListUI = (props: LinkListUIProps) => {
  return (
    <div className="panel-body">
      <Table
        rowKey={'link_id'}
        bordered={true}
        dataSource={props.linkList}
        loading={props.listLoading}
        pagination={props.pagination}
        onChange={props.onListChange}
        footer={() => null}
      >
        <Table.Column
          title={'链接ID'}
          dataIndex="link_id"
          width={80}
          key={'link_id'}
          align={'center'}
        />
        <Table.Column title={'链接名称'} dataIndex="name" key={'name'} width={200} />
        <Table.Column
          title={'链接地址'}
          dataIndex="url"
          key={'url'}
          render={(url: string) => (
            <a href={url} target="_blank" rel="noreferrer">
              {url}
            </a>
          )}
        />
        <Table.Column
          title={'排序号'}
          dataIndex="sequence"
          key={'sequence'}
          width={80}
          align={'center'}
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
          width={140}
          key={'action'}
          align={'center'}
          render={(linkListItem: LinkListItemType) => (
            <Space>
              <ActionButton
                text="修改"
                icon={<FormOutlined />}
                onClick={() => props.onEditClick(linkListItem)}
                havePermission={linkListItem.action?.is_edit == 1}
              />
              <Popconfirm
                title="确定要删除吗?"
                onConfirm={() => props.onDeleteConfirm(linkListItem)}
                okText="确定"
                cancelText="取消"
              >
                <ActionButton
                  text="删除"
                  icon={<CloseSquareOutlined />}
                  havePermission={linkListItem.action?.is_delete == 1}
                />
              </Popconfirm>
            </Space>
          )}
        />
      </Table>
    </div>
  )
}

export default LinkListUI
