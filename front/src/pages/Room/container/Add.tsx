import React, { useEffect, useState } from 'react'
import { RoomService } from '@/services/Room'
import { message } from 'antd'
import RoomFormUI from '../component/FormUI'
import { RoomAddResp, RoomInfoType } from '@/types/roomType'
import { useNavigate } from 'react-router-dom'
import { HouseInfoType } from '@/types/houseType'

const RoomAdd: React.FC = () => {
  const [houseList, setHouseList] = useState<HouseInfoType[]>([])
  const navigate = useNavigate()

  useEffect(() => {
    getRoomAddInfo()
  }, [])

  /**
   * 获取添加房产信息
   */
  const getRoomAddInfo = () => {
    RoomService.getAddRoomInfo()
      .then((roomAddInfoRes: RoomAddResp) => {
        setHouseList(roomAddInfoRes.houses)
      })
      .catch((e) => {
        console.log('获取添加房产信息 err:', e)
      })
  }

  /**
   * 添加房间保存
   * @param roomInfo
   */
  const addOnFinishCallback = (roomInfo: RoomInfoType) => {
    RoomService.saveRoom(roomInfo)
      .then(() => {
        message.success('保存成功', 2, () => {
          navigate('/room/list')
        })
      })
      .catch((e) => {
        console.log(e)
      })
  }

  /**
   * 返回组件
   */
  return (
    <div>
      <RoomFormUI houseList={houseList} onFinishCallback={addOnFinishCallback} />
    </div>
  )
}

export default RoomAdd
