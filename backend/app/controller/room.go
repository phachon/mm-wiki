package controller

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
)

// RoomAdd 添加房间页面
func RoomAdd(ctx *gin.Context) error {
	// 获取所有的合租房间
	splitLeaseHouses, err := service.NewHouse(ctx).GetAllowSplitLeaseHouses()
	if err != nil {
		sysLogErrorf(ctx, "[AccountAdd] 获取所有合租房间失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	return RespJsonSuccess(ctx, map[string]interface{}{
		"houses": splitLeaseHouses,
	})
}

// RoomSave 添加房间保存
func RoomSave(ctx *gin.Context) error {

	houseId := GetParamInt64(ctx, "house_id")
	name := GetParamString(ctx, "name")
	area := GetParamIntDef(ctx, "area", 0)
	directionType := GetParamIntDef(ctx, "direction_type", 0)
	roomType := GetParamIntDef(ctx, "room_type", 0)
	hasToilet := GetParamIntDef(ctx, "has_toilet", 0)
	hasBalcony := GetParamIntDef(ctx, "has_balcony", 0)
	mouthRent := GetParamIntDef(ctx, "mouth_rent", 0)
	allowLease := GetParamIntDef(ctx, "allow_lease", 0)

	// 判断参数合法性
	if houseId <= 0 {
		logger.WithContext(ctx).Warnf("[RoomSave] houseId empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "所属房产ID不能为空")
	}
	if name == "" {
		logger.WithContext(ctx).Warnf("[RoomSave] name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "房间名不能为空")
	}
	if area <= 0 {
		logger.WithContext(ctx).Warnf("[RoomSave] area empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "面积不能为空")
	}

	// room 房间实体
	roomEntity := &entity.RoomEntity{
		HouseId:       houseId,
		Name:          name,
		Area:          area,
		DirectionType: directionType,
		RoomType:      roomType,
		HasToilet:     hasToilet,
		HasBalcony:    hasBalcony,
		MouthRent:     mouthRent,
		AllowLease:    allowLease,
	}
	err := service.NewRoom(ctx).Create(roomEntity)
	if err != nil {
		sysLogErrorf(ctx, "[RoomSave] 添加房间失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	sysLogInfof(ctx, "[RoomSave] 添加房间 %+v 成功", roomEntity.RoomId)
	return RespJsonSuccess(ctx, map[string]interface{}{
		"room_id": roomEntity.RoomId,
	})
}

// RoomEdit 修改房间页面
func RoomEdit(ctx *gin.Context) error {

	roomId := GetParamInt64(ctx, "room_id")
	// 判断参数合法性
	if roomId == 0 {
		logger.WithContext(ctx).Warnf("[RoomEdit] room_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "房间id不能为空")
	}

	// 获取房间信息
	roomInfo, err := service.NewRoom(ctx).GetRoomByRoomId(roomId)
	if err != nil {
		sysLogErrorf(ctx, "[RoomEdit] 获取房间 %d 信息失败: err=%+v", roomId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if roomInfo == nil {
		sysLogWarnf(ctx, "[RoomEdit] 房间id %d 不合法", roomId)
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "房间id不存在")
	}

	// 获取所有的房产用于选择
	houses, err := service.NewHouse(ctx).GetAllHouses()
	if err != nil {
		sysLogErrorf(ctx, "[RoomEdit] 获取所有的房产信息失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	return RespJsonSuccess(ctx, &entity.RoomEditResp{
		RoomInfo: roomInfo,
		Houses:   houses,
	})
}

// RoomModify 修改房间保存
func RoomModify(ctx *gin.Context) error {

	roomId := GetParamInt64(ctx, "room_id")
	name := GetParamString(ctx, "name")
	area := GetParamIntDef(ctx, "area", 0)
	directionType := GetParamIntDef(ctx, "direction_type", 0)
	roomType := GetParamIntDef(ctx, "room_type", 0)
	hasToilet := GetParamIntDef(ctx, "has_toilet", 0)
	hasBalcony := GetParamIntDef(ctx, "has_balcony", 0)
	mouthRent := GetParamIntDef(ctx, "mouth_rent", 0)
	allowLease := GetParamIntDef(ctx, "allow_lease", 0)

	// 判断参数合法性
	if name == "" {
		logger.WithContext(ctx).Warnf("[RoomModify] name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "房间名不能为空")
	}
	if area <= 0 {
		logger.WithContext(ctx).Warnf("[RoomModify] area empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "面积不能为空")
	}

	// room 房间实体
	roomEntity := entity.RoomEntity{
		RoomId:        roomId,
		Name:          name,
		Area:          area,
		DirectionType: directionType,
		RoomType:      roomType,
		HasToilet:     hasToilet,
		HasBalcony:    hasBalcony,
		MouthRent:     mouthRent,
		AllowLease:    allowLease,
	}
	err := service.NewRoom(ctx).Update(roomEntity)
	if err != nil {
		sysLogErrorf(ctx, "[RoomModify] 更新房间 %d 失败 err=%+v", roomId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	sysLogInfof(ctx, "[RoomModify] 更新房间 %d 成功", roomId)
	return RespJsonSuccess(ctx, nil)
}

// RoomList 房间列表
func RoomList(ctx *gin.Context) error {

	pageSize := GetParamIntDef(ctx, "page_size", 20)
	pageNum := GetParamIntDef(ctx, "page_num", 1)
	keywordsStr := GetParamStringDef(ctx, "keywords", "")

	var keywords *entity.RoomKeywords
	if len(keywordsStr) > 0 {
		jErr := json.Unmarshal([]byte(keywordsStr), &keywords)
		if jErr != nil {
			logger.WithContext(ctx).Errorf("[RoomList] GetRoomsByLimit err=%s", jErr.Error())
		}
	}

	serviceRoom := service.NewRoom(ctx)
	// 获取房间列表
	rooms, err := serviceRoom.GetRoomsByLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[RoomList] 获取房间列表失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取分页信息
	pageInfo, err := serviceRoom.GetPageInfoLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[RoomList] 获取房间分页信息失败: err=%+v", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取所有的房间信息
	roomList, err := serviceRoom.FormatRoomList(rooms)
	if err != nil {
		sysLogErrorf(ctx, "[RoomList] 获取房间列表信息失败: err=%+v", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	data := map[string]interface{}{
		"list":      roomList,
		"page_info": pageInfo,
	}
	return RespJsonSuccess(ctx, data)
}

// RoomDelete 房间删除
func RoomDelete(ctx *gin.Context) error {

	roomId := GetParamInt64(ctx, "room_id")

	// 判断参数合法性
	if roomId == 0 {
		logger.WithContext(ctx).Warnf("[RoomDelete] room_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "房间id不存在")
	}

	// 删除房间
	err := service.NewRoom(ctx).DeleteByRoomId(roomId)
	if err != nil {
		sysLogErrorf(ctx, "[RoomDelete] 删除房间失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	sysLogInfof(ctx, "[RoomDelete] 删除房间 %d 成功", roomId)
	return RespJsonSuccess(ctx, nil)
}

// RoomDetail 房间详情
func RoomDetail(ctx *gin.Context) error {

	roomId := GetParamInt64(ctx, "room_id")
	// 判断参数合法性
	if roomId == 0 {
		logger.WithContext(ctx).Warnf("[RoomDetail] room_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "房间id不存在")
	}

	// 获取房间信息
	roomInfo, err := service.NewRoom(ctx).GetRoomByRoomId(roomId)
	if err != nil {
		sysLogErrorf(ctx, "[RoomDetail] 获取房间 %d 信息失败: err=%+v", roomId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if roomInfo == nil {
		sysLogWarnf(ctx, "[RoomDetail] 房间id %d 不存在", roomId)
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "房间id不存在")
	}
	return RespJsonSuccess(ctx, &entity.RoomDetailResp{
		RoomInfo: roomInfo,
	})
}
