package controller

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
	"github.com/phachon/mm-wiki/utils"
)

// OrderAdd 订单添加
func OrderAdd(ctx *gin.Context) error {
	orderType := GetParamInt(ctx, "order_type")

	var houseService = service.NewHouse(ctx)
	var houses = []*entity.HouseEntity{}
	var err errors.BizError
	// 根据订单类型获取可出租且未出租的房屋
	if orderType == entity.OrderTypeSplit {
		houses, err = houseService.GetSplitLeaseAndNotRentHouses()
	}
	if orderType == entity.OrderTypeWhole {
		houses, err = houseService.GetLeaseAndNotRentHouses()
	}
	if err != nil {
		logger.WithContext(ctx).Warnf("[OrderAdd] GeLeaseHouses err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	// 如果是合租，需要返回房产下的所有订单
	var housesRooms = make(map[int64][]*entity.RoomEntity)
	if orderType == entity.OrderTypeSplit {
		var houseIds []int64
		for _, house := range houses {
			houseIds = append(houseIds, house.HouseId)
		}
		rooms, err := service.NewRoom(ctx).GetNotRentRoomByHouseIds(houseIds)
		if err != nil {
			logger.WithContext(ctx).Warnf("[OrderAdd] GetRooms err=%+v", err)
			return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
		}
		logger.WithContext(ctx).Infof("[OrderAdd] GetRooms=%+v", rooms)
		for _, room := range rooms {
			housesRooms[room.HouseId] = append(housesRooms[room.HouseId], room)
		}
	}

	var houseAddResp = &entity.OrderAddResp{
		SelectHouses:     houses,      // 选择房产列表
		SelectHouseRooms: housesRooms, // 选择房产下订单列表
	}
	return RespJsonSuccess(ctx, houseAddResp)
}

// OrderSave 添加订单保存
func OrderSave(ctx *gin.Context) error {

	orderType := GetParamIntDef(ctx, "order_type", 0)
	houseId := GetParamInt64(ctx, "house_id")
	roomId := GetParamInt64(ctx, "room_id")
	hirerId := GetParamInt64(ctx, "hirer_id")
	startTimeStr := GetParamString(ctx, "start_time")
	endTimeStr := GetParamString(ctx, "end_time")
	unitRent := GetParamIntDef(ctx, "unit_rent", 0)
	payType := GetParamIntDef(ctx, "pay_type", 0)
	remarks := GetParamString(ctx, "remarks")

	// 判断参数合法性
	if houseId <= 0 {
		logger.WithContext(ctx).Warnf("[OrderSave] houseId empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "所属房产不能为空")
	}
	startTime, cErr := time.ParseInLocation("2006-01-02", startTimeStr, time.Local)
	if cErr != nil {
		logger.WithContext(ctx).Warnf("[OrderSave] start_time err:%+v", cErr)
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "开始时间错误")
	}
	endTime, cErr := time.ParseInLocation("2006-01-02", endTimeStr, time.Local)
	if cErr != nil {
		logger.WithContext(ctx).Warnf("[OrderSave] end_time err:%+v", cErr)
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "结束时间错误")
	}
	if startTime.Unix() >= endTime.Unix() {
		logger.WithContext(ctx).Warnf("[OrderSave] end_time err:%+v", cErr)
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "开始时间必须小于结束时间")
	}
	if orderType == entity.OrderTypeSplit && roomId <= 0 {
		logger.WithContext(ctx).Warnf("[OrderSave] split order roomId empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "合租必须选择房间")
	}

	// 订单实体
	orderEntity := &entity.OrderEntity{
		OrderType:   orderType,
		HouseId:     houseId,
		RoomId:      roomId,
		HirerId:     hirerId,
		StartTime:   utils.NewJsonDate(startTime),
		EndTime:     utils.NewJsonDate(endTime),
		PayType:     payType,
		UnitRent:    unitRent,
		AccountId:   global.ContextValueLoginAccountID(ctx),
		AccountName: global.ContextValueLoginAccountName(ctx),
		Remarks:     remarks,
	}
	err := service.NewOrder(ctx).Create(orderEntity)
	if err != nil {
		sysLogErrorf(ctx, "[OrderSave] 添加订单失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	sysLogInfof(ctx, "[OrderSave] 添加订单 %+v 成功", orderEntity.OrderId)
	return RespJsonSuccess(ctx, map[string]interface{}{
		"order_id": orderEntity.OrderId,
	})
}

// OrderList 房间列表
func OrderList(ctx *gin.Context) error {

	pageSize := GetParamIntDef(ctx, "page_size", 20)
	pageNum := GetParamIntDef(ctx, "page_num", 1)
	keywordsStr := GetParamStringDef(ctx, "keywords", "")

	var keywords *entity.OrderKeywords
	if len(keywordsStr) > 0 {
		jErr := json.Unmarshal([]byte(keywordsStr), &keywords)
		if jErr != nil {
			logger.WithContext(ctx).Errorf("[OrderList] GetOrderList err=%s", jErr.Error())
		}
	}

	serviceOrder := service.NewOrder(ctx)
	// 获取订单列表
	orders, err := serviceOrder.GetOrdersByLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[OrderList] 获取订单列表失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取分页信息
	pageInfo, err := serviceOrder.GetPageInfoLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[OrderList] 获取订单分页信息失败: err=%+v", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	// 获取订单 item 信息
	orderList, err := serviceOrder.FormatOrderList(orders)
	if err != nil {
		sysLogErrorf(ctx, "[OrderList] 获取订单分页信息失败: err=%+v", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	data := map[string]interface{}{
		"list":      orderList,
		"page_info": pageInfo,
	}
	return RespJsonSuccess(ctx, data)
}

// OrderDetail 订单详情
func OrderDetail(ctx *gin.Context) error {

	orderId := GetParamInt64(ctx, "order_id")
	// 判断参数合法性
	if orderId == 0 {
		logger.WithContext(ctx).Warnf("[OrderDetail] order_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "订单id不存在")
	}

	// 获取订单信息
	orderInfo, err := service.NewOrder(ctx).GetOrderByOrderId(orderId)
	if err != nil {
		sysLogErrorf(ctx, "[OrderDetail] 获取订单 %d 信息失败: err=%+v", orderId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if orderInfo == nil {
		sysLogWarnf(ctx, "[OrderDetail] 订单id %d 不存在", orderId)
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "房间id不存在")
	}
	return RespJsonSuccess(ctx, &entity.OrderDetailResp{
		OrderInfo: orderInfo,
	})
}

// OrderModify 订单更新
func OrderModify(ctx *gin.Context) error {

	orderId := GetParamInt64(ctx, "order_id")
	status := GetParamIntDef(ctx, "status", -1)
	startTimeStr := GetParamString(ctx, "start_time")
	endTimeStr := GetParamString(ctx, "end_time")
	unitRent := GetParamIntDef(ctx, "unit_rent", 0)
	payType := GetParamIntDef(ctx, "pay_type", 0)
	remarks := GetParamString(ctx, "remarks")

	// 判断参数合法性
	if orderId <= 0 {
		logger.WithContext(ctx).Warnf("[OrderModify] order_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "订单ID不能为空")
	}
	startTime, cErr := time.ParseInLocation("2006-01-02", startTimeStr, time.Local)
	if cErr != nil {
		logger.WithContext(ctx).Warnf("[OrderModify] start_time err:%+v", cErr)
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "开始时间错误")
	}
	endTime, cErr := time.ParseInLocation("2006-01-02", endTimeStr, time.Local)
	if cErr != nil {
		logger.WithContext(ctx).Warnf("[OrderModify] end_time err:%+v", cErr)
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "结束时间错误")
	}
	if startTime.Unix() >= endTime.Unix() {
		logger.WithContext(ctx).Warnf("[OrderModify] end_time err:%+v", cErr)
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "开始时间必须小于结束时间")
	}

	// 订单实体
	orderEntity := entity.OrderEntity{
		OrderId:   orderId,
		StartTime: utils.NewJsonDate(startTime),
		EndTime:   utils.NewJsonDate(endTime),
		PayType:   payType,
		UnitRent:  unitRent,
		Status:    status,
		Remarks:   remarks,
	}
	err := service.NewOrder(ctx).Update(orderEntity)
	if err != nil {
		sysLogErrorf(ctx, "[OrderModify] 修改订单失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	sysLogInfof(ctx, "[OrderModify] 修改订单 %d 成功", orderId)
	return RespJsonSuccess(ctx, nil)
}
