import { Button, Popconfirm, Space, Table, TablePaginationConfig, Tag } from 'antd'
import { NoticeInfoType, NoticeListItemType } from '@/types/noticeType'
import { CloseSquareOutlined, FormOutlined, TeamOutlined, LockOutlined } from '@ant-design/icons'
import { NoticePublishStatusTag } from './ToolsUI'
import ActionButton from '@/components/Action/ActionButton'

interface NoticeListUIProps {
  listLoading: boolean
  noticeList: NoticeListItemType[]
  pagination: TablePaginationConfig
  onListChange: (pageConfig: TablePaginationConfig, filters: any, sorter: any) => void
  onEditClick: (noticeInfo: NoticeInfoType) => void
  onDeleteConfirm: (noticeInfo: NoticeInfoType) => void
}

/**
 * 公告列表 UI 组件
 * @param props
 * @returns
 */
const NoticeListUI = (props: NoticeListUIProps) => {
  return (
    <div className="panel-body">
      <Table
        rowKey={'notice_id'}
        bordered={true}
        dataSource={props.noticeList}
        loading={props.listLoading}
        pagination={props.pagination}
        onChange={props.onListChange}
        footer={() => null}
      >
        <Table.Column
          title={'公告ID'}
          dataIndex="notice_id"
          width={80}
          key={'notice_id'}
          align={'center'}
        />
        <Table.Column
          title={'发布账号'}
          dataIndex="account_name"
          key={'account_name'}
          width={100}
        />
        <Table.Column title={'公告标题'} dataIndex="title" key={'title'} width={200} />
        <Table.Column title={'公告内容'} dataIndex="content" key={'content'} />
        <Table.Column
          title={'发布状态'}
          dataIndex="publish_status"
          key={'publish_status'}
          width={100}
          align={'center'}
          render={NoticePublishStatusTag}
        />
        <Table.Column
          title={'开始时间'}
          dataIndex="start_time"
          key={'start_time'}
          width={180}
          align={'center'}
        />
        <Table.Column
          title={'结束时间'}
          dataIndex="end_time"
          key={'end_time'}
          width={180}
          align={'center'}
        />
        <Table.Column
          title={'操作'}
          width={140}
          key={'action'}
          align={'center'}
          render={(noticeListItem: NoticeListItemType) => (
            <Space>
              <ActionButton
                text="修改"
                icon={<FormOutlined />}
                onClick={() => props.onEditClick(noticeListItem)}
                havePermission={noticeListItem.action?.is_edit == 1}
              />
              <Popconfirm
                title="确定要删除吗?"
                onConfirm={() => props.onDeleteConfirm(noticeListItem)}
                okText="确定"
                cancelText="取消"
              >
                <ActionButton
                  text="删除"
                  icon={<CloseSquareOutlined />}
                  havePermission={noticeListItem.action?.is_delete == 1}
                />
              </Popconfirm>
            </Space>
          )}
        />
      </Table>
    </div>
  )
}

export default NoticeListUI
