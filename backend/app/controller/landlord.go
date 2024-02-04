package controller

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
)

// LandlordAdd 添加业主页面
func LandlordAdd(ctx *gin.Context) error {
	return nil
}

// LandlordSave 添加业主保存
func LandlordSave(ctx *gin.Context) error {
	nickName := GetParamString(ctx, "nick_name")
	givenName := GetParamString(ctx, "given_name")
	sex := GetParamIntDef(ctx, "sex", 0)
	email := GetParamString(ctx, "email")
	idCardNumber := GetParamString(ctx, "id_card_number")
	mobile := GetParamString(ctx, "mobile")
	address := GetParamString(ctx, "address")

	// 判断参数合法性
	if nickName == "" {
		logger.WithContext(ctx).Warnf("[LandlordSave] 业主昵称不能为空")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "昵称不能为空")
	}
	if givenName == "" {
		logger.WithContext(ctx).Warnf("[LandlordSave] 姓名不能为空")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "姓名不能为空")
	}
	if mobile == "" {
		logger.WithContext(ctx).Warnf("[LandlordSave] 手机不能为空")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "手机不能为空")
	}
	if address == "" {
		logger.WithContext(ctx).Warnf("[LandlordSave] 住址不能为空")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "住址不能为空")
	}

	// landlord 业主实体
	landlordEntity := &entity.LandlordEntity{
		NickName:     nickName,
		GivenName:    givenName,
		Sex:          sex,
		Email:        email,
		IdCardNumber: idCardNumber,
		Mobile:       mobile,
		Address:      address,
	}
	err := service.NewLandlord(ctx).Create(landlordEntity)
	if err != nil {
		sysLogErrorf(ctx, "[LandlordSave] 添加业主失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[LandlordSave] 添加业主 %+v 成功", landlordEntity.LandlordId)

	return RespJsonSuccess(ctx, map[string]interface{}{
		"landlord_id": landlordEntity.LandlordId,
	})
}

// LandlordEdit 修改业主页面
func LandlordEdit(ctx *gin.Context) error {

	landlordId := GetParamInt64(ctx, "landlord_id")
	// 判断参数合法性
	if landlordId == 0 {
		logger.WithContext(ctx).Warnf("[LandlordEdit] 业主id不能为空")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "业主id不能为空")
	}

	// 获取业主信息
	landlordInfo, err := service.NewLandlord(ctx).GetLandlordByLandlordId(landlordId)
	if err != nil {
		sysLogErrorf(ctx, "[LandlordEdit] 获取业主 %d 信息失败: err=%+v", landlordId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if landlordInfo == nil {
		sysLogWarnf(ctx, "[LandlordEdit] 业主id %d 不合法", landlordId)
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "业主id不存在")
	}

	return RespJsonSuccess(ctx, &entity.LandlordEditResp{
		LandlordInfo: landlordInfo,
	})
}

// LandlordModify 修改业主保存
func LandlordModify(ctx *gin.Context) error {

	landlordId := GetParamInt64(ctx, "landlord_id")
	nickName := GetParamString(ctx, "nick_name")
	givenName := GetParamString(ctx, "given_name")
	sex := GetParamIntDef(ctx, "sex", 0)
	email := GetParamString(ctx, "email")
	idCardNumber := GetParamString(ctx, "id_card_number")
	mobile := GetParamString(ctx, "mobile")
	address := GetParamString(ctx, "address")

	// 判断参数合法性
	if landlordId == 0 {
		logger.WithContext(ctx).Warnf("[LandlordModify] landlord_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "业主id不存在")
	}
	if nickName == "" {
		logger.WithContext(ctx).Warnf("[LandlordModify] nick_name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "业主名不能为空")
	}
	if givenName == "" {
		logger.WithContext(ctx).Warnf("[LandlordModify] given_name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "昵称不能为空")
	}
	if mobile == "" {
		logger.WithContext(ctx).Warnf("[LandlordModify] 手机不能为空")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "手机不能为空")
	}
	if address == "" {
		logger.WithContext(ctx).Warnf("[LandlordModify] 住址不能为空")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "住址不能为空")
	}

	// landlord 业主实体
	landlordEntity := entity.LandlordEntity{
		LandlordId:   landlordId,
		NickName:     nickName,
		GivenName:    givenName,
		Sex:          sex,
		Email:        email,
		IdCardNumber: idCardNumber,
		Mobile:       mobile,
		Address:      address,
	}
	err := service.NewLandlord(ctx).Update(landlordEntity)
	if err != nil {
		sysLogErrorf(ctx, "[LandlordModify] 更新业主 %d 失败 err=%+v", landlordId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	sysLogInfof(ctx, "[LandlordModify] 更新业主 %d 成功", landlordId)
	return RespJsonSuccess(ctx, nil)
}

// LandlordList 业主列表
func LandlordList(ctx *gin.Context) error {

	pageSize := GetParamIntDef(ctx, "page_size", 20)
	pageNum := GetParamIntDef(ctx, "page_num", 1)
	keywordsStr := GetParamStringDef(ctx, "keywords", "")
	isAll := GetParamIntDef(ctx, "is_all", 0)

	var keywords *entity.LandlordKeywords
	if len(keywordsStr) > 0 {
		jErr := json.Unmarshal([]byte(keywordsStr), &keywords)
		if jErr != nil {
			logger.WithContext(ctx).Errorf("[LandlordList] GetLandlordsByLimit err=%s", jErr.Error())
		}
	}

	serviceLandlord := service.NewLandlord(ctx)
	// 获取所有的业主列表
	if isAll == 1 {
		landlords, err := serviceLandlord.GetAllLandlords()
		if err != nil {
			sysLogErrorf(ctx, "[LandlordList] 获取业主列表失败: err=%+v", err)
			return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
		}
		data := map[string]interface{}{
			"list": landlords,
		}
		return RespJsonSuccess(ctx, data)
	}

	// 分页获取业主列表
	landlords, err := serviceLandlord.GetLandlordsByLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[LandlordList] 获取业主列表失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取分页信息
	pageInfo, err := serviceLandlord.GetPageInfoLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[LandlordList] 获取业主分页信息失败: err=%+v", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	data := map[string]interface{}{
		"list":      landlords,
		"page_info": pageInfo,
	}
	return RespJsonSuccess(ctx, data)
}

// LandlordDelete 业主删除
func LandlordDelete(ctx *gin.Context) error {

	landlordId := GetParamInt64(ctx, "landlord_id")

	// 判断参数合法性
	if landlordId == 0 {
		logger.WithContext(ctx).Warnf("[LandlordDelete] landlord_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "业主id不存在")
	}

	// 更新业主状态
	err := service.NewLandlord(ctx).DeleteByLandlordId(landlordId)
	if err != nil {
		sysLogErrorf(ctx, "[LandlordDelete] 删除业主失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[LandlordDelete] 删除业主 %d 成功", landlordId)

	return RespJsonSuccess(ctx, nil)
}

// LandlordDetail 业主详情
func LandlordDetail(ctx *gin.Context) error {

	landlordId := GetParamInt64(ctx, "landlord_id")

	// 判断参数合法性
	if landlordId == 0 {
		logger.WithContext(ctx).Warnf("[LandlordDetail] landlord_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "业主id不存在")
	}

	// 获取业主详情
	landlord, err := service.NewLandlord(ctx).GetLandlordByLandlordId(landlordId)
	if err != nil {
		sysLogErrorf(ctx, "[LandlordDetail] 获取业主详情失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if landlord == nil {
		sysLogWarnf(ctx, "[LandlordDetail] 业主id %d 不存在", landlordId)
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "业主id不存在")
	}
	return RespJsonSuccess(ctx, &entity.LandlordDetailResp{
		LandlordInfo: landlord,
	})
}
