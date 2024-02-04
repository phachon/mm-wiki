package controller

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
)

// HouseSave 添加房产保存
func HouseSave(ctx *gin.Context) error {

	region := GetParamString(ctx, "region")
	address := GetParamString(ctx, "address")
	houseNumber := GetParamString(ctx, "house_number")
	decorationType := GetParamIntDef(ctx, "decoration_type", 0)
	sizeType := GetParamIntDef(ctx, "size_type", 0)
	area := GetParamIntDef(ctx, "area", 0)
	buildTime := GetParamString(ctx, "build_time")
	landlordId := GetParamIntDef(ctx, "landlord_id", 0)
	allowLease := GetParamIntDef(ctx, "allow_lease", 0)
	allowSplitLease := GetParamIntDef(ctx, "allow_split_lease", 0)
	mouthRent := GetParamIntDef(ctx, "mouth_rent", 0)

	// 判断参数合法性
	if region == "" {
		logger.WithContext(ctx).Warnf("[HouseSave] region empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "所在区域不能为空")
	}
	if address == "" {
		logger.WithContext(ctx).Warnf("[HouseSave] address empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "所在地址不能为空")
	}
	if houseNumber == "" {
		logger.WithContext(ctx).Warnf("[HouseSave] houseNumber empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "门牌号不能为空")
	}
	if area <= 0 {
		logger.WithContext(ctx).Warnf("[HouseSave] area empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "面积不能为空")
	}
	if landlordId <= 0 {
		logger.WithContext(ctx).Warnf("[HouseSave] landlord_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "业主不能为空")
	}
	if allowSplitLease == entity.HouseAllowLeaseYes {
		if mouthRent <= 0 {
			logger.WithContext(ctx).Warnf("[HouseSave] mouth_rent empty")
			return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "租金不能为空")
		}
	}

	// house 房产实体
	houseEntity := &entity.HouseEntity{
		LandlordId:      int64(landlordId),
		Region:          region,
		Address:         address,
		HouseNumber:     houseNumber,
		DecorationType:  decorationType,
		SizeType:        sizeType,
		Area:            area,
		BuildTime:       buildTime,
		AllowLease:      allowLease,
		AllowSplitLease: allowSplitLease,
		MouthRent:       mouthRent,
	}
	err := service.NewHouse(ctx).Create(houseEntity)
	if err != nil {
		sysLogErrorf(ctx, "[HouseSave] 添加房产失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	sysLogInfof(ctx, "[HouseSave] 添加房产 %+v 成功", houseEntity.HouseId)
	return RespJsonSuccess(ctx, map[string]interface{}{
		"house_id": houseEntity.HouseId,
	})
}

// HouseEdit 修改房产页面
func HouseEdit(ctx *gin.Context) error {

	houseId := GetParamInt64(ctx, "house_id")
	// 判断参数合法性
	if houseId == 0 {
		logger.WithContext(ctx).Warnf("[HouseEdit] house_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "房产id不能为空")
	}

	// 获取房产信息
	houseInfo, err := service.NewHouse(ctx).GetHouseByHouseId(houseId)
	if err != nil {
		sysLogErrorf(ctx, "[HouseEdit] 获取房产 %d 信息失败: err=%+v", houseId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if houseInfo == nil {
		sysLogWarnf(ctx, "[HouseEdit] 房产id %d 不合法", houseId)
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "房产id不存在")
	}

	// 获取所有的业主用于选择
	landlords, err := service.NewLandlord(ctx).GetAllLandlords()
	if err != nil {
		sysLogErrorf(ctx, "[HouseEdit] 获取所有的业主信息失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	return RespJsonSuccess(ctx, &entity.HouseEditResp{
		HouseInfo: houseInfo,
		Landlords: landlords,
	})
}

// HouseModify 修改房产保存
func HouseModify(ctx *gin.Context) error {

	houseId := GetParamInt64(ctx, "house_id")
	region := GetParamString(ctx, "region")
	address := GetParamString(ctx, "address")
	houseNumber := GetParamString(ctx, "house_number")
	decorationType := GetParamIntDef(ctx, "decoration_type", 0)
	sizeType := GetParamIntDef(ctx, "size_type", 0)
	area := GetParamIntDef(ctx, "area", 0)
	buildTime := GetParamString(ctx, "build_time")
	landlordId := GetParamIntDef(ctx, "landlord_id", 0)
	allowLease := GetParamIntDef(ctx, "allow_lease", 0)
	allowSplitLease := GetParamIntDef(ctx, "allow_split_lease", 0)
	mouthRent := GetParamIntDef(ctx, "mouth_rent", 0)

	// 判断参数合法性
	if region == "" {
		logger.WithContext(ctx).Warnf("[HouseModify] region empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "所在区域不能为空")
	}
	if address == "" {
		logger.WithContext(ctx).Warnf("[HouseModify] address empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "所在地址不能为空")
	}
	if houseNumber == "" {
		logger.WithContext(ctx).Warnf("[HouseModify] houseNumber empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "门牌号不能为空")
	}
	if area <= 0 {
		logger.WithContext(ctx).Warnf("[HouseModify] area empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "面积不能为空")
	}
	if landlordId <= 0 {
		logger.WithContext(ctx).Warnf("[HouseModify] landlord_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "业主不能为空")
	}
	if allowSplitLease == entity.HouseAllowLeaseYes {
		if mouthRent <= 0 {
			logger.WithContext(ctx).Warnf("[HouseModify] mouth_rent empty")
			return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "租金不能为空")
		}
	}

	// house 房产实体
	houseEntity := entity.HouseEntity{
		HouseId:         houseId,
		LandlordId:      int64(landlordId),
		Region:          region,
		Address:         address,
		HouseNumber:     houseNumber,
		DecorationType:  decorationType,
		SizeType:        sizeType,
		Area:            area,
		BuildTime:       buildTime,
		AllowLease:      allowLease,
		AllowSplitLease: allowSplitLease,
		MouthRent:       mouthRent,
	}
	err := service.NewHouse(ctx).Update(houseEntity)
	if err != nil {
		sysLogErrorf(ctx, "[HouseModify] 更新房产 %d 失败 err=%+v", houseId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	sysLogInfof(ctx, "[HouseModify] 更新房产 %d 成功", houseId)
	return RespJsonSuccess(ctx, nil)
}

// HouseList 房产列表
func HouseList(ctx *gin.Context) error {

	pageSize := GetParamIntDef(ctx, "page_size", 20)
	pageNum := GetParamIntDef(ctx, "page_num", 1)
	keywordsStr := GetParamStringDef(ctx, "keywords", "")
	isAll := GetParamIntDef(ctx, "is_all", 0)

	var keywords *entity.HouseKeywords
	if len(keywordsStr) > 0 {
		jErr := json.Unmarshal([]byte(keywordsStr), &keywords)
		if jErr != nil {
			logger.WithContext(ctx).Errorf("[HouseList] GetHousesByLimit err=%s", jErr.Error())
		}
	}

	serviceHouse := service.NewHouse(ctx)

	if isAll == 1 {
		// 获取所有的房产列表
		houses, err := serviceHouse.GetAllHouses()
		if err != nil {
			sysLogErrorf(ctx, "[HouseList] 获取所有的房产列表失败: err=%+v", err)
			return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
		}
		data := map[string]interface{}{
			"list": houses,
		}
		return RespJsonSuccess(ctx, data)
	}
	// 分页获取房产列表
	houses, err := serviceHouse.GetHousesByLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[HouseList] 获取房产列表失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取分页信息
	pageInfo, err := serviceHouse.GetPageInfoLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[HouseList] 获取房产分页信息失败: err=%+v", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	data := map[string]interface{}{
		"list":      houses,
		"page_info": pageInfo,
	}
	return RespJsonSuccess(ctx, data)
}

// HouseDelete 房产删除
func HouseDelete(ctx *gin.Context) error {

	houseId := GetParamInt64(ctx, "house_id")

	// 判断参数合法性
	if houseId == 0 {
		logger.WithContext(ctx).Warnf("[HouseDelete] house_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "房产id不存在")
	}

	// 删除房产
	err := service.NewHouse(ctx).DeleteByHouseId(houseId)
	if err != nil {
		sysLogErrorf(ctx, "[HouseDelete] 删除房产失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	sysLogInfof(ctx, "[HouseDelete] 删除房产 %d 成功", houseId)
	return RespJsonSuccess(ctx, nil)
}

// HouseDetail 房产详情
func HouseDetail(ctx *gin.Context) error {

	houseId := GetParamInt64(ctx, "house_id")
	// 判断参数合法性
	if houseId == 0 {
		logger.WithContext(ctx).Warnf("[HouseDetail] house_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "房产id不存在")
	}

	// 获取房产信息
	houseInfo, err := service.NewHouse(ctx).GetHouseByHouseId(houseId)
	if err != nil {
		sysLogErrorf(ctx, "[HouseDetail] 获取房产 %d 信息失败: err=%+v", houseId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if houseInfo == nil {
		sysLogWarnf(ctx, "[HouseDetail] 房产id %d 不存在", houseId)
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "房产id不存在")
	}
	return RespJsonSuccess(ctx, &entity.HouseDetailResp{
		HouseInfo: houseInfo,
	})
}
