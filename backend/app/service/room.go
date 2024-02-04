package service

import (
	"context"

	"github.com/phachon/mm-wiki/app/dao"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
	"github.com/phachon/mm-wiki/utils"
)

const (
	RoomListDefaultPageSize = 20 // 房间列表默认一页 20 条
)

// Room 房间业务逻辑
type Room struct {
	ctx     context.Context
	daoRoom *dao.Room
}

// NewRoom 创建房间业务逻辑对象
func NewRoom(ctx context.Context) *Room {
	return &Room{
		ctx:     ctx,
		daoRoom: dao.NewRoom(ctx),
	}
}

// Create 创建房间
func (r *Room) Create(roomEntity *entity.RoomEntity) errors.BizError {

	// todo 生成房间唯一编号

	// 插入一条记录
	err := r.daoRoom.Insert(roomEntity)
	if err != nil {
		return err
	}
	return nil
}

// Update 修改房间
func (r *Room) Update(roomEntity entity.RoomEntity) errors.BizError {
	// 查找房间是否存在
	updateRoom, err := r.daoRoom.GetRoomByRoomId(roomEntity.RoomId)
	if err != nil {
		return err
	}
	if updateRoom == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "房间不存在")
	}
	if updateRoom.Status == entity.RoomStatusDelete {
		return errors.Errorf(errors.BusinessForbiddenError, "删除的房间无法修改")
	}
	// 更新字段
	err = r.daoRoom.Update(roomEntity)
	if err != nil {
		return err
	}
	return nil
}

// GetRoomByRoomId 根据房间ID获取房间详情
func (r *Room) GetRoomByRoomId(roomID int64) (room *entity.RoomEntity, err errors.BizError) {
	room, err = r.daoRoom.GetRoomByRoomId(roomID)
	if err != nil {
		return room, err
	}
	return room, nil
}

// GetRoomsByRoomIds 根据房间ID获取房间详情
func (r *Room) GetRoomsByRoomIds(roomIDs []int64) (rooms []*entity.RoomEntity, err errors.BizError) {
	if len(roomIDs) == 0 {
		return rooms, nil
	}
	rooms, err = r.daoRoom.GetRoomsByRoomIds(roomIDs)
	if err != nil {
		return rooms, err
	}
	return rooms, nil
}

// GetRoomsMapByRoomIds 根据房间ID获取房间详情Map
func (r *Room) GetRoomsMapByRoomIds(roomIDs []int64) (roomsMap map[int64]*entity.RoomEntity, err errors.BizError) {
	roomsMap = make(map[int64]*entity.RoomEntity)
	if len(roomIDs) == 0 {
		return roomsMap, nil
	}
	rooms, err := r.daoRoom.GetRoomsByRoomIds(roomIDs)
	if err != nil {
		return roomsMap, err
	}
	for _, room := range rooms {
		roomsMap[room.RoomId] = room
	}
	return roomsMap, nil
}

// GetRoomByRoomId 根据多个房间ID获取房间列表
func (r *Room) GetRoomByHouseIds(roomIds []int64) (rooms []*entity.RoomEntity, err errors.BizError) {
	rooms, err = r.daoRoom.GetRoomsByHouseIds(roomIds)
	if err != nil {
		return rooms, err
	}
	return rooms, nil
}

// GetNotRentRoomByHouseIds 根据房间ID获取未出租房间
func (r *Room) GetNotRentRoomByHouseIds(roomIds []int64) (rooms []*entity.RoomEntity, err errors.BizError) {
	rooms, err = r.daoRoom.GetNotRentRoomsByHouseIds(roomIds)
	if err != nil {
		return rooms, err
	}
	return rooms, nil
}

// DeleteByRoomId 根据房间ID删除
func (r *Room) DeleteByRoomId(roomId int64) errors.BizError {
	// 查找房间是否存在
	updateRoom, err := r.daoRoom.GetRoomByRoomId(roomId)
	if err != nil {
		return err
	}
	if updateRoom == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "房间id %d 不存在", roomId)
	}
	// 更新状态字段
	err = r.daoRoom.UpdateStatus(roomId, entity.RoomStatusDelete)
	if err != nil {
		return err
	}
	return nil
}

// GetRoomsByLimit 分页获取房间列表
func (r *Room) GetRoomsByLimit(pageSize int, pageNum int, keywords *entity.RoomKeywords) (rooms []*entity.RoomEntity, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, RoomListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	offset := (pageNum - 1) * pageSize
	// 无搜索词
	if keywords == nil {
		return r.daoRoom.GetRoomsByLimit(pageSize, offset)
	}
	return r.daoRoom.GetRoomsByKeywordAndLimit(pageSize, offset, keywords)
}

// GetAllRooms 获取所有的房间
func (r *Room) GetAllRooms() (rooms []*entity.RoomEntity, err errors.BizError) {
	return r.daoRoom.GetAllRooms()
}

// GetPageInfoLimit 获取分页信息
func (r *Room) GetPageInfoLimit(pageSize int, pageNum int, keywords *entity.RoomKeywords) (pagination *entity.PageInfo, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, RoomListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	pageInfo := new(entity.PageInfo)
	var totalNum int64

	// 无搜索词
	if keywords == nil {
		totalNum, err = r.daoRoom.CountRooms()
	} else {
		totalNum, err = r.daoRoom.CountRoomsByKeywords(keywords)
	}
	if err != nil {
		logger.WithContext(r.ctx).Warnf("[service.Room] GetPageInfoLimit pageSize=%d pageNum=%d err=%s",
			pageSize, pageNum, err.Error())
		return pageInfo, err
	}
	return entity.GetPageInfo(totalNum, pageSize, pageNum), nil
}

// FormatRoomList 格式化房间列表
func (r *Room) FormatRoomList(rooms []*entity.RoomEntity) ([]*entity.RoomListItem, errors.BizError) {
	var (
		houseIds []int64
		roomList []*entity.RoomListItem
	)
	for _, room := range rooms {
		houseIds = append(houseIds, room.HouseId)
	}
	// 获取房间信息
	housesMap, err := NewHouse(r.ctx).GetHousesMapByHouseIds(houseIds)
	if err != nil {
		return roomList, err
	}
	for _, room := range rooms {
		var roomItem = &entity.RoomListItem{
			RoomEntity: room,
		}
		if house, ok := housesMap[room.HouseId]; ok {
			roomItem.HouseInfo = house
		}
		roomList = append(roomList, roomItem)
	}
	return roomList, nil
}
