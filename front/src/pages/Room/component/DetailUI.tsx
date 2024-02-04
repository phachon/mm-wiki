import { Descriptions, Tag } from 'antd'
import { RoomInfoType } from '@/types/roomType'
import {
  RoomAllowLeaseTag,
  RoomBalconyStatusTag,
  RoomDirectionTypeTag,
  RoomLeaseStatusTag,
  RoomToiletStatusTag,
  RoomTypeTag
} from './ToolsUI'

interface RoomDetailUIProps {
  /**
   * 房间信息
   */
  roomInfo?: RoomInfoType
}

/**
 * 房间详情 UI 组件
 */
const RoomDetailUI = (props: RoomDetailUIProps) => {
  const roomInfo = props.roomInfo
  return (
    <div key={roomInfo?.room_id.toString()}>
      <Descriptions bordered size="small" column={12}>
        <Descriptions.Item label="房间ID" span={12} key="room_id">
          {roomInfo?.room_id.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="房产ID" span={12} key="house_id">
          {roomInfo?.house_id.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="房间名称" span={12} key="name">
          {roomInfo?.name.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="房间面积" span={12} key="area">
          {roomInfo?.area.toString()}平米
        </Descriptions.Item>
        <Descriptions.Item label="房间朝向" span={12} key="direction_type">
          {RoomDirectionTypeTag(roomInfo?.direction_type)}
        </Descriptions.Item>
        <Descriptions.Item label="房间类型" span={12} key="room_type">
          {RoomTypeTag(roomInfo?.room_type)}
        </Descriptions.Item>
        <Descriptions.Item label="是否有独卫" span={12} key="has_toilet">
          {RoomToiletStatusTag(roomInfo?.has_toilet)}
        </Descriptions.Item>
        <Descriptions.Item label="是否有阳台" span={12} key="has_balcony">
          {RoomBalconyStatusTag(roomInfo?.has_balcony)}
        </Descriptions.Item>
        <Descriptions.Item label="月租金" span={12} key="mouth_rent">
          {roomInfo?.mouth_rent}¥/月
        </Descriptions.Item>
        <Descriptions.Item label="允许出租" span={12} key="allow_lease">
          {RoomAllowLeaseTag(roomInfo?.allow_lease)}
        </Descriptions.Item>
        <Descriptions.Item label="出租状态" span={12} key="lease_status">
          {RoomLeaseStatusTag(roomInfo?.lease_status)}
        </Descriptions.Item>
        <Descriptions.Item label="创建时间" span={12} key="create_time">
          {roomInfo?.create_time.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="修改时间" span={12} key="update_time">
          {roomInfo?.update_time.toString()}
        </Descriptions.Item>
      </Descriptions>
    </div>
  )
}

export default RoomDetailUI
