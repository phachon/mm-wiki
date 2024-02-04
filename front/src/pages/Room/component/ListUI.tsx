import { Popconfirm, Space, Table, TablePaginationConfig, Tag } from 'antd'
import { RoomInfoType, RoomListItemType } from '@/types/roomType'
import { CloseSquareOutlined, FormOutlined, SelectOutlined } from '@ant-design/icons'
import { RoomAllowLeaseTag, RoomLeaseStatusTag, RoomShowAllTag } from './ToolsUI'
import { HouseShowWholeAddress } from '@/pages/House/component/ToolsUI'
import ActionButton from '@/components/Action/ActionButton'

interface RoomListUIProps {
  listLoading: boolean

  /**
   * 房间列表
   */
  roomList: RoomListItemType[]

  /**
   * 翻页信息
   */
  pagination: TablePaginationConfig

  /**
   * 列表点击操作
   * @param pageConfig
   * @param filters
   * @param sorter
   * @returns
   */
  onRoomListChange: (pageConfig: TablePaginationConfig, filters: any, sorter: any) => void

  /**
   * 房间详情点击
   * @param roomInfo
   * @returns
   */
  onRoomDetailClick: (roomListItem: RoomListItemType) => void

  /**
   * 房产详情点击
   * @param roomInfo
   * @returns
   */
  onHouseDetailClick: (roomListItem: RoomListItemType) => void

  /**
   * 房间修改点击
   * @param roomInfo
   * @returns
   */
  onRoomEditClick: (roomListItem: RoomListItemType) => void

  /**
   * 房间删除操作
   * @param roomInfo
   * @param status
   * @returns
   */
  onRoomDeleteClick: (roomListItem: RoomListItemType) => void
}

const RoomListUI = (props: RoomListUIProps) => {
  return (
    <div className="panel-body">
      <Table
        rowKey={'room_id'}
        bordered={true}
        dataSource={props.roomList}
        loading={props.listLoading}
        pagination={props.pagination}
        onChange={props.onRoomListChange}
        footer={() => ''}
      >
        <Table.Column
          title={'房间ID'}
          dataIndex="room_id"
          width={80}
          key={'room_id'}
          align={'center'}
        />
        <Table.Column
          title={'所属房产'}
          key={'house_id'}
          render={(roomListItem: RoomListItemType) => (
            <a onClick={() => props.onHouseDetailClick(roomListItem)}>
              <Space>
                {HouseShowWholeAddress(roomListItem.house_info)}
                <SelectOutlined />
              </Space>
            </a>
          )}
        />
        <Table.Column title={'房间名称'} key={'name'} dataIndex="name" />
        <Table.Column
          title={'房间配置'}
          key={'config'}
          width={200}
          render={(roomInfo: RoomInfoType) => RoomShowAllTag(roomInfo)}
        />
        <Table.Column
          title={'房间面积'}
          dataIndex="area"
          key={'area'}
          width={100}
          align={'center'}
          render={(area: number) => <>{area} 平米</>}
        />
        <Table.Column
          title={'月租金'}
          key={'mouth_rent'}
          width={120}
          align={'center'}
          render={(roomInfo: RoomInfoType) => <>{roomInfo.mouth_rent}$/月</>}
        />
        <Table.Column
          title={'允许出租'}
          dataIndex="allow_lease"
          width={100}
          key={'allow_lease'}
          align={'center'}
          render={RoomAllowLeaseTag}
        />
        <Table.Column
          title={'出租状态'}
          dataIndex="lease_status"
          width={100}
          key={'lease_status'}
          align={'center'}
          render={RoomLeaseStatusTag}
        />
        <Table.Column
          title={'操作'}
          width={200}
          key={'action'}
          align={'center'}
          render={(roomListItem: RoomListItemType) => (
            <Space>
              <ActionButton
                text="详情"
                icon={<SelectOutlined />}
                onClick={() => props.onRoomDetailClick(roomListItem)}
                havePermission={true}
              />
              <ActionButton
                text="修改"
                icon={<FormOutlined />}
                onClick={() => props.onRoomEditClick(roomListItem)}
                havePermission={roomListItem.action?.is_edit == 1}
              />
              <Popconfirm
                title="确定要删除吗?"
                onConfirm={() => {
                  props.onRoomDeleteClick(roomListItem)
                }}
                okText="确定"
                cancelText="取消"
              >
                <ActionButton
                  text="删除"
                  icon={<CloseSquareOutlined />}
                  havePermission={roomListItem.action?.is_delete == 1}
                />
              </Popconfirm>
            </Space>
          )}
        />
      </Table>
    </div>
  )
}

export default RoomListUI
