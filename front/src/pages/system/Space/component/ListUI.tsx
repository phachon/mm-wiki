import { Popconfirm, Space, Table, TablePaginationConfig } from 'antd'
import { CloseSquareOutlined, FormOutlined, TeamOutlined, ExportOutlined } from '@ant-design/icons'
import { SpaceTypeTagUI, SpaceVisitLevelTagUI } from './ToolsUI'
import ActionButton from '@/components/Action/ActionButton'
import { SpaceInfoType, SpaceListItemType } from '@/types/spaceType'

interface SpaceListUIProps {
  listLoading: boolean
  spaceList: SpaceListItemType[] // 空间列表数据
  pagination: TablePaginationConfig
  onListChange: (pageConfig: TablePaginationConfig, filters: any, sorter: any) => void
  onEditClick: (spaceInfo: SpaceInfoType) => void
  onDeleteConfirm: (spaceInfo: SpaceInfoType) => void
  onAdminListClick?: (spaceInfo: SpaceInfoType) => void
}

/**
 * 空间列表 UI 组件
 * @param props
 * @returns
 */
const SpaceListUI = (props: SpaceListUIProps) => {
  return (
    <div className="panel-body">
      <Table
        rowKey={'space_id'}
        bordered={true}
        dataSource={props.spaceList}
        loading={props.listLoading}
        pagination={props.pagination}
        onChange={props.onListChange}
        footer={() => null}
      >
        <Table.Column
          title={'空间ID'}
          dataIndex="space_id"
          width={100}
          key={'space_id'}
          align={'center'}
        />
        <Table.Column
          title={'空间Key'}
          dataIndex="space_key"
          width={100}
          key={'space_key'}
          render={(text: string) => (
            <a href={`/space/${text}`} target="_blank">
              <Space>
                {text}
                <ExportOutlined />
              </Space>
            </a>
          )}
        />
        <Table.Column title={'空间名'} dataIndex="name" key={'name'} width={250} />
        <Table.Column title={'描述'} dataIndex="description" key={'description'} />
        <Table.Column
          title={'空间类型'}
          dataIndex="space_type"
          key={'space_type'}
          width={120}
          align={'center'}
          render={SpaceTypeTagUI}
        />
        <Table.Column
          title={'访问级别'}
          dataIndex="visit_level"
          key={'visit_level'}
          width={100}
          align={'center'}
          render={SpaceVisitLevelTagUI}
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
          width={200}
          key={'action'}
          align={'center'}
          render={(spaceListItem: SpaceListItemType) => (
            <Space>
              <ActionButton
                text="管理员"
                icon={<TeamOutlined />}
                onClick={() =>
                  props.onAdminListClick ? props.onAdminListClick(spaceListItem) : null
                }
                havePermission={true}
              />
              <ActionButton
                text="修改"
                icon={<FormOutlined />}
                onClick={() => props.onEditClick(spaceListItem)}
                havePermission={true}
              />

              <Popconfirm
                title="确定要删除吗?"
                onConfirm={() => props.onDeleteConfirm(spaceListItem)}
                okText="确定"
                cancelText="取消"
              >
                <ActionButton text="删除" icon={<CloseSquareOutlined />} havePermission={true} />
              </Popconfirm>
            </Space>
          )}
        />
      </Table>
    </div>
  )
}

export default SpaceListUI
