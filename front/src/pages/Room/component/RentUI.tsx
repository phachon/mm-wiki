import { Descriptions, Tag } from 'antd'
import { RoomInfoType } from '@/types/roomType'

interface RoomRentUIProps {
  /**
   * 房间详情
   */
  roomInfo?: RoomInfoType
}

/**
 * 房间租金 UI 组件
 */
const RoomRentUI = (props: RoomRentUIProps) => {
  const roomInfo = props.roomInfo
  return (
    <div key={roomInfo?.room_id.toString()}>
      <Descriptions bordered size="small" column={12}>
        <Descriptions.Item label="房间ID" span={12} key="room_id">
          {roomInfo?.room_id.toString()}
        </Descriptions.Item>
        <Descriptions.Item label="月租金" span={12} key="mouth_rent">
          {roomInfo?.mouth_rent} CNY
        </Descriptions.Item>
      </Descriptions>
    </div>
  )
}

export default RoomRentUI
