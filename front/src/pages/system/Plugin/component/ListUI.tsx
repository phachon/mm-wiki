import { Popconfirm, Space, Table, TablePaginationConfig, Tag } from 'antd'
import { CloseSquareOutlined, FormOutlined, CheckCircleOutlined, StopOutlined } from '@ant-design/icons'
import ActionButton from '@/components/Action/ActionButton'
import { PluginInfoType, PluginListItemType } from '@/types/pluginType'

interface PluginListUIProps {
  listLoading: boolean
  pluginList: PluginListItemType[]
  pagination: TablePaginationConfig
  onListChange: (pageConfig: TablePaginationConfig, filters: any, sorter: any) => void
  onEditClick: (pluginInfo: PluginInfoType) => void
  onDeleteConfirm: (pluginInfo: PluginInfoType) => void
  onStatusClick: (pluginInfo: PluginInfoType) => void
}

const PluginStatusTag = (status: number) => {
  switch (status) {
    case 0:
      return <Tag color="default">禁用</Tag>
    case 1:
      return <Tag color="success">启用</Tag>
    default:
      return <Tag color="cyan">未知</Tag>
  }
}

const PluginListUI = (props: PluginListUIProps) => {
  return (
    <div className="panel-body">
      <Table
        rowKey={'plugin_id'}
        bordered={true}
        dataSource={props.pluginList}
        loading={props.listLoading}
        pagination={props.pagination}
        onChange={props.onListChange}
        footer={() => null}
      >
        <Table.Column
          title={'插件ID'}
          dataIndex="plugin_id"
          width={80}
          key={'plugin_id'}
          align={'center'}
        />
        <Table.Column title={'插件名称'} dataIndex="name" key={'name'} width={150} />
        <Table.Column title={'插件标识'} dataIndex="key" key={'key'} width={120} />
        <Table.Column title={'版本号'} dataIndex="version" key={'version'} width={80} align={'center'} />
        <Table.Column title={'作者'} dataIndex="author" key={'author'} width={100} />
        <Table.Column
          title={'状态'}
          dataIndex="status"
          key={'status'}
          width={100}
          align={'center'}
          render={PluginStatusTag}
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
          width={240}
          key={'action'}
          align={'center'}
          render={(pluginListItem: PluginListItemType) => (
            <Space>
              <ActionButton
                text={pluginListItem.status === 1 ? '禁用' : '启用'}
                icon={pluginListItem.status === 1 ? <StopOutlined /> : <CheckCircleOutlined />}
                onClick={() => props.onStatusClick(pluginListItem)}
                havePermission={pluginListItem.action?.is_edit == 1}
              />
              <ActionButton
                text="修改"
                icon={<FormOutlined />}
                onClick={() => props.onEditClick(pluginListItem)}
                havePermission={pluginListItem.action?.is_edit == 1}
              />
              <Popconfirm
                title="确定要删除吗?"
                onConfirm={() => props.onDeleteConfirm(pluginListItem)}
                okText="确定"
                cancelText="取消"
              >
                <ActionButton
                  text="删除"
                  icon={<CloseSquareOutlined />}
                  havePermission={pluginListItem.action?.is_delete == 1}
                />
              </Popconfirm>
            </Space>
          )}
        />
      </Table>
    </div>
  )
}

export default PluginListUI
