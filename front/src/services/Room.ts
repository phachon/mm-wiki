import { RoomAddResp, RoomDetailResp, RoomEditResp, RoomListResp } from '../types/roomType'
import httpRequest from './http'
import Base from './Base'

const roomUrl = {
  add: '/system/room/add',
  save: '/system/room/save',
  edit: '/system/room/edit',
  modify: '/system/room/modify',
  delete: '/system/room/delete',
  list: '/system/room/list',
  detail: '/system/room/detail'
}

/**
 * Room 房间服务
 */
class Room extends Base {
  public constructor() {
    super()
  }

  /**
   * getAddRoomInfo 获取添加房间信息
   */
  public getAddRoomInfo(): Promise<any> {
    const roomAddUrl = this.getProxyUrl(roomUrl.add)
    return httpRequest.get<RoomAddResp>(roomAddUrl, {})
  }

  /**
   * saveRoom 添加保存房间
   */
  public saveRoom(roomInfo: {}): Promise<any> {
    const roomSaveUrl = this.getProxyUrl(roomUrl.save)
    return httpRequest.post<any>(roomSaveUrl, {}, roomInfo)
  }

  /**
   * getEditRoomInfo 获取编辑房间信息
   */
  public getEditRoomInfo(room_id: bigint): Promise<any> {
    const roomEditUrl = this.getProxyUrl(roomUrl.edit)
    return httpRequest.get<RoomEditResp>(roomEditUrl, {
      room_id: room_id
    })
  }

  /**
   * modifyRoom 修改保存房间
   */
  public modifyRoom(roomEditInfo: {}): Promise<any> {
    const roomModifyUrl = this.getProxyUrl(roomUrl.modify)
    return httpRequest.post<any>(roomModifyUrl, {}, roomEditInfo)
  }

  /**
   * roomList 房间列表
   */
  public roomList(
    pageSize: number | undefined,
    pageNum: number | undefined,
    keywords: {}
  ): Promise<any> {
    const roomListUrl = this.getProxyUrl(roomUrl.list)
    return httpRequest.get<RoomListResp>(roomListUrl, {
      page_size: pageSize,
      page_num: pageNum,
      keywords: JSON.stringify(keywords)
    })
  }

  /**
   * deleteRoom 删除房间
   */
  public deleteRoom(roomId: bigint): Promise<any> {
    let roomDeleteUrl = this.getProxyUrl(roomUrl.delete)
    return httpRequest.post<any>(
      roomDeleteUrl,
      {},
      {
        room_id: roomId
      }
    )
  }

  /**
   * getRoomDetail 获取房间详情
   */
  public getRoomDetail(roomId: bigint): Promise<any> {
    let roomDetailUrl = this.getProxyUrl(roomUrl.detail)
    return httpRequest.get<RoomDetailResp>(roomDetailUrl, { room_id: roomId })
  }
}

export const RoomService = new Room()
