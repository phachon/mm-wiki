package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/utils"
	"gorm.io/gorm"
)

const (
	// TableNameRoom 房屋房间表
	TableNameRoom = "hms_room"
	// RoomPrimaryKey 房间表主键ID
	RoomPrimaryKey = "room_id"
)

// Room 房屋房间表数据
type Room struct {
	ctx context.Context
}

// NewRoom 创建房屋房间表数据对象
func NewRoom(ctx context.Context) *Room {
	return &Room{
		ctx: ctx,
	}
}

// Insert 创建房间插入一条房间记录
func (h *Room) Insert(roomEntity *entity.RoomEntity) errors.BizError {
	roomEntity.CreateTime = utils.NewJsonTime(time.Now())
	roomEntity.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(h.ctx).
		Table(TableNameRoom).
		Save(roomEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetRoomByRoomId 根据房间ID获取房间信息
func (h *Room) GetRoomByRoomId(roomId int64) (room *entity.RoomEntity, err errors.BizError) {

	room = &entity.RoomEntity{}
	db := GetDB(dbNameKms).WithContext(h.ctx).Table(TableNameRoom).
		Where(map[string]interface{}{
			RoomPrimaryKey: roomId,
			"status":       entity.RoomStatusDefault,
		}).
		First(&room)

	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return room, nil
}

// GetRoomsByRoomIds 根据多个房间ID批量获取房间ID
func (h *Room) GetRoomsByRoomIds(roomIds []int64) (rooms []*entity.RoomEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(h.ctx).Table(TableNameRoom).
		Where("status = ?", entity.RoomStatusDefault).
		Where("room_id IN (?)", roomIds).
		Find(&rooms)
	if db.Error != nil {
		return rooms, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return rooms, nil
}

// GetRoomsByHouseIds 根据多个房间ID批量获取房间
func (h *Room) GetRoomsByHouseIds(roomIds []int64) (rooms []*entity.RoomEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(h.ctx).Table(TableNameRoom).
		Where("status = ?", entity.RoomStatusDefault).
		Where("house_id IN (?)", roomIds).
		Find(&rooms)
	if db.Error != nil {
		return rooms, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return rooms, nil
}

// GetNotRentRoomsByHouseIds 根据多个房间ID批量获取房间
func (h *Room) GetNotRentRoomsByHouseIds(roomIds []int64) (rooms []*entity.RoomEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(h.ctx).Table(TableNameRoom).
		Where("status = ?", entity.RoomStatusDefault).
		Where("lease_status = ?", entity.RoomLeaseStatusNotRent).
		Where("house_id IN (?)", roomIds).
		Find(&rooms)
	if db.Error != nil {
		return rooms, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return rooms, nil
}

// Update 更新房间，只会更新如下字段
func (h *Room) Update(room entity.RoomEntity) errors.BizError {
	room.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(h.ctx).Table(TableNameRoom).
		Select(
			"Name", "Area", "DirectionType",
			"RoomType", "HasToilet", "HasBalcony",
			"MouthRent", "LeaseStatus", "AllowLease",
			"UpdateTime",
		).
		Where(map[string]interface{}{
			RoomPrimaryKey: room.RoomId,
		}).
		Updates(room)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// UpdateStatus 更新房间状态
func (h *Room) UpdateStatus(roomId int64, status int) errors.BizError {
	updateTime := utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(h.ctx).Table(TableNameRoom).
		Where(map[string]interface{}{
			RoomPrimaryKey: roomId,
		}).
		Update("status", status).
		Update("update_time", updateTime)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// UpdateLeaseStatus 更新房间出租状态
func (h *Room) UpdateLeaseStatus(roomId int64, leaseStatus int) errors.BizError {
	updateTime := utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(h.ctx).Table(TableNameRoom).
		Where(map[string]interface{}{
			RoomPrimaryKey: roomId,
		}).
		Update("lease_status", leaseStatus).
		Update("update_time", updateTime)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// GetAllRooms 获取所有的房间
func (h *Room) GetAllRooms() (room []*entity.RoomEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(h.ctx).
		Table(TableNameRoom).
		Where("status = ?", entity.RoomStatusDefault).
		Find(&room)
	if db.Error != nil {
		return room, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return room, nil
}

// GetRoomsByKeywords 根据房间名模糊匹配房间
func (h *Room) GetRoomsByKeywords(keywords *entity.RoomKeywords) (rooms []*entity.RoomEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(h.ctx).
		Table(TableNameRoom).
		Where("status = ?", entity.RoomStatusDefault)
	db = h.getRoomWhereByKeywords(keywords, db)
	db.Find(&rooms)
	if db.Error != nil {
		return rooms, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return rooms, nil
}

func (*Room) getRoomWhereByKeywords(keywords *entity.RoomKeywords, db *gorm.DB) *gorm.DB {
	if keywords.DirectionType != nil && *keywords.DirectionType >= 0 {
		db = db.Where("direction_type = ?", *keywords.DirectionType)
	}
	if keywords.RoomType != nil && *keywords.RoomType >= 0 {
		db = db.Where("room_type = ?", *keywords.RoomType)
	}
	if keywords.LeaseStatus != nil && *keywords.LeaseStatus >= 0 {
		db = db.Where("lease_status = ?", *keywords.LeaseStatus)
	}
	if keywords.AllowLease != nil && *keywords.AllowLease >= 0 {
		db = db.Where("allow_lease = ?", *keywords.AllowLease)
	}
	if keywords.HouseId != "" {
		db = db.Where("house_id = ?", keywords.HouseId)
	}
	if keywords.Name != "" {
		db = db.Where("name LIKE ?", "%"+keywords.Name+"%")
	}
	return db
}

// GetRoomsByKeywordAndLimit 根据关键字分页获取房间
func (h *Room) GetRoomsByKeywordAndLimit(limit int, offset int, keywords *entity.RoomKeywords) (rooms []*entity.RoomEntity, err errors.BizError) {
	if keywords == nil {
		return h.GetRoomsByLimit(limit, offset)
	}
	db := GetDB(dbNameKms).WithContext(h.ctx).
		Table(TableNameRoom).
		Where("status = ?", entity.RoomStatusDefault)
	db = h.getRoomWhereByKeywords(keywords, db)
	db = db.Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", RoomPrimaryKey)).
		Find(&rooms)
	if db.Error != nil {
		return rooms, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return rooms, nil
}

// GetRoomsByLimit 分页获取房间列表
func (h *Room) GetRoomsByLimit(limit int, offset int) (rooms []*entity.RoomEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(h.ctx).
		Table(TableNameRoom).
		Where("status = ?", entity.RoomStatusDefault).
		Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", RoomPrimaryKey)).
		Find(&rooms)
	if db.Error != nil {
		return rooms, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return rooms, nil
}

// CountRooms 房间总数
func (h *Room) CountRooms() (count int64, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(h.ctx).
		Table(TableNameRoom).
		Where("status = ?", entity.RoomStatusDefault).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// CountRoomsByKeywords 根据关键词搜索房间总数
func (h *Room) CountRoomsByKeywords(keywords *entity.RoomKeywords) (count int64, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(h.ctx).
		Table(TableNameRoom).
		Where("status = ?", entity.RoomStatusDefault)
	db = h.getRoomWhereByKeywords(keywords, db)
	db.Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}
