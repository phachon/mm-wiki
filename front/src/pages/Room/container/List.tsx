import React, { useEffect, useState } from 'react'
import RoomListUI from '../component/ListUI'
import {
  RoomDetailResp,
  RoomEditResp,
  RoomInfoType,
  RoomListItemType,
  RoomListResp
} from '@/types/roomType'
import { RoomService } from '@/services/Room'
import { message, Modal, TablePaginationConfig } from 'antd'
import RoomSearchUI from '../component/SearchUI'
import RoomFormUI from '../component/FormUI'
import RoomDetailUI from '../component/DetailUI'
import { initPagination } from '@/types/adminType'
import { HouseDetailResp, HouseInfoType } from '@/types/houseType'
import { HouseService } from '@/services/House'
import HouseDetailUI from '@/pages/House/component/DetailUI'
import RoomRentUI from '../component/RentUI'

let searchValues = {}

const RoomList: React.FC = () => {
  // 房间列表
  const [roomList, setRoomList] = useState<RoomListItemType[]>([])
  const [pagination, setPagination] = useState(initPagination)

  // 房间修改
  const [editModalOpen, setEditModalOpen] = useState<boolean>(false)
  const [editRoomInfo, setEditRoomInfo] = useState<RoomInfoType>()
  const [houseList, setHouseList] = useState<HouseInfoType[]>([])

  // 房间详情
  const [detailRoomInfo, setDetailRoomInfo] = useState<RoomInfoType>()
  const [detailModalOpen, setDetailModalOpen] = useState<boolean>(false)

  // 所属房产
  const [detailHouseModalOpen, setDetailHouseModalOpen] = useState<boolean>(false)
  const [detailHouseInfo, setDetailHouseInfo] = useState<HouseInfoType>()

  useEffect(() => {
    getRoomList(initPagination, {})
  }, [])

  /**
   * 请求房间列表
   * @param pageConfig 翻页信息
   * @param searchKeywords 搜索信息
   */
  const getRoomList = (pageConfig: TablePaginationConfig, searchKeywords: {}) => {
    searchValues = searchKeywords
    RoomService.roomList(pageConfig.pageSize, pageConfig.current, searchKeywords)
      .then((res: RoomListResp) => {
        setRoomList(res.list)
        setPagination({
          ...initPagination,
          current: res.page_info?.page_num,
          pageSize: res.page_info?.page_size,
          total: res.page_info?.total_num
        })
      })
      .catch((e) => {
        console.log('获取房间列表失败 err:', e)
      })
  }
  /**
   * 搜索查询操作
   * @param values 搜索表单数据
   */
  const searchChangeCallback = (values: {}) => {
    getRoomList(initPagination, values)
  }
  /**
   * 搜索重置操作
   */
  const searchResetCallback = () => {
    getRoomList(initPagination, {})
  }
  /**
   * 列表分页请求
   * @param pageConfig
   * @param filters
   * @param sorter
   */
  const onRoomListChange = (pageConfig: TablePaginationConfig, filters: any, sorter: any) => {
    getRoomList(pageConfig, searchValues)
  }
  /**
   * 修改点击操作
   * @param roomInfo 房间信息
   */
  const onRoomEditClick = (roomInfo: RoomListItemType) => {
    RoomService.getEditRoomInfo(roomInfo.room_id)
      .then((res: RoomEditResp) => {
        setEditRoomInfo(res.room_info)
        setHouseList(res.houses)
        setEditModalOpen(true)
      })
      .catch((e) => {
        console.log('获取修改信息 err:', e)
      })
  }

  /**
   * 详情点击操作
   * @param roomInfo 房间信息
   */
  const onRoomDetailClick = (roomInfo: RoomListItemType) => {
    RoomService.getRoomDetail(roomInfo.room_id)
      .then((roomDetail: RoomDetailResp) => {
        setDetailRoomInfo(roomDetail.room_info)
        setDetailModalOpen(true)
      })
      .catch((e) => {
        console.log('获取房间详情 err:', e)
      })
  }

  /**
   * 房间详情点击操作
   * @param roomInfo 房间信息
   */
  const onHouseDetailClick = (roomListItem: RoomListItemType) => {
    setDetailHouseInfo(roomListItem.house_info)
    setDetailHouseModalOpen(true)
  }

  /**
   * 删除房间操作
   * @param roomInfo
   * @param status
   */
  const onRoomDeleteClick = (roomListItem: RoomListItemType) => {
    RoomService.deleteRoom(roomListItem.room_id)
      .then(() => {
        message.success('删除成功', 2, () => {
          getRoomList(pagination, searchValues)
        })
      })
      .catch((e) => {
        console.log('删除房间 err:', e)
      })
  }
  /**
   * 修改弹框取消操作
   */
  const editModalCancelCallback = () => {
    setEditModalOpen(false)
  }
  /**
   * 修改完成操作
   * @param roomInfo RoomInfoType
   */
  const editOnFinishCallback = (roomInfo: RoomInfoType) => {
    RoomService.modifyRoom(roomInfo)
      .then(() => {
        message.success('修改成功', 2, () => {
          setEditModalOpen(false)
          getRoomList(pagination, searchValues)
        })
      })
      .catch((e) => {
        console.log('修改房间保存失败 err:', e)
      })
  }
  return (
    <div className="panel">
      <RoomSearchUI
        searchChangeCallback={searchChangeCallback}
        searchResetCallback={searchResetCallback}
      />
      <RoomListUI
        listLoading={false}
        roomList={roomList}
        pagination={pagination}
        onRoomListChange={onRoomListChange}
        onRoomEditClick={onRoomEditClick}
        onRoomDetailClick={onRoomDetailClick}
        onHouseDetailClick={onHouseDetailClick}
        onRoomDeleteClick={onRoomDeleteClick}
      />
      <Modal
        title="房间修改"
        width={1200}
        open={editModalOpen}
        onCancel={editModalCancelCallback}
        footer={null}
      >
        <RoomFormUI
          houseList={houseList}
          roomInfo={editRoomInfo}
          onFinishCallback={editOnFinishCallback}
        />
      </Modal>
      <Modal
        title="房间详情"
        width={570}
        open={detailModalOpen}
        onCancel={() => {
          setDetailModalOpen(false)
        }}
        footer={null}
      >
        <RoomDetailUI roomInfo={detailRoomInfo} />
      </Modal>
      <Modal
        title="所属房产"
        width={570}
        open={detailHouseModalOpen}
        onCancel={() => {
          setDetailHouseModalOpen(false)
        }}
        footer={null}
      >
        <HouseDetailUI detailHouseInfo={detailHouseInfo} />
      </Modal>
    </div>
  )
}

export default RoomList
