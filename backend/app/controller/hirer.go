package controller

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
)

// HirerAdd 添加租客页面
func HirerAdd(ctx *gin.Context) error {
	return nil
}

// HirerSave 添加租客保存
func HirerSave(ctx *gin.Context) error {
	name := GetParamString(ctx, "name")
	sex := GetParamIntDef(ctx, "sex", 0)
	email := GetParamString(ctx, "email")
	idCardNumber := GetParamString(ctx, "id_card_number")
	mobile := GetParamString(ctx, "mobile")
	emergencyContact := GetParamString(ctx, "emergency_contact")
	emergencyMobile := GetParamString(ctx, "emergency_mobile")

	// 判断参数合法性
	if name == "" {
		logger.WithContext(ctx).Warnf("[HirerSave] 姓名不能为空")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "姓名不能为空")
	}
	if mobile == "" {
		logger.WithContext(ctx).Warnf("[HirerSave] 手机不能为空")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "手机不能为空")
	}

	// hirer 租客实体
	hirerEntity := &entity.HirerEntity{
		Name:             name,
		Sex:              sex,
		Email:            email,
		IdCardNumber:     idCardNumber,
		Mobile:           mobile,
		EmergencyContact: emergencyContact,
		EmergencyMobile:  emergencyMobile,
	}
	err := service.NewHirer(ctx).Create(hirerEntity)
	if err != nil {
		sysLogErrorf(ctx, "[HirerSave] 添加租客失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[HirerSave] 添加租客 %+v 成功", hirerEntity.HirerId)

	return RespJsonSuccess(ctx, map[string]interface{}{
		"hirer_id": hirerEntity.HirerId,
	})
}

// HirerEdit 修改租客页面
func HirerEdit(ctx *gin.Context) error {

	hirerId := GetParamInt64(ctx, "hirer_id")
	// 判断参数合法性
	if hirerId == 0 {
		logger.WithContext(ctx).Warnf("[HirerEdit] 租客id不能为空")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "租客id不能为空")
	}

	// 获取租客信息
	hirerInfo, err := service.NewHirer(ctx).GetHirerByHirerId(hirerId)
	if err != nil {
		sysLogErrorf(ctx, "[HirerEdit] 获取租客 %d 信息失败: err=%+v", hirerId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if hirerInfo == nil {
		sysLogWarnf(ctx, "[HirerEdit] 租客id %d 不合法", hirerId)
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "租客id不存在")
	}

	return RespJsonSuccess(ctx, &entity.HirerEditResp{
		HirerInfo: hirerInfo,
	})
}

// HirerModify 修改租客保存
func HirerModify(ctx *gin.Context) error {

	hirerId := GetParamInt64(ctx, "hirer_id")
	name := GetParamString(ctx, "name")
	sex := GetParamIntDef(ctx, "sex", 0)
	email := GetParamString(ctx, "email")
	idCardNumber := GetParamString(ctx, "id_card_number")
	mobile := GetParamString(ctx, "mobile")
	emergencyContact := GetParamString(ctx, "emergency_contact")
	emergencyMobile := GetParamString(ctx, "emergency_mobile")

	// 判断参数合法性
	if hirerId == 0 {
		logger.WithContext(ctx).Warnf("[HirerModify] hirer_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "租客id不存在")
	}
	if name == "" {
		logger.WithContext(ctx).Warnf("[HirerModify] name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "昵称不能为空")
	}
	if mobile == "" {
		logger.WithContext(ctx).Warnf("[HirerModify] 手机不能为空")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "手机不能为空")
	}

	// hirer 租客实体
	hirerEntity := entity.HirerEntity{
		HirerId:          hirerId,
		Name:             name,
		Sex:              sex,
		Email:            email,
		IdCardNumber:     idCardNumber,
		Mobile:           mobile,
		EmergencyContact: emergencyContact,
		EmergencyMobile:  emergencyMobile,
	}
	err := service.NewHirer(ctx).Update(hirerEntity)
	if err != nil {
		sysLogErrorf(ctx, "[HirerModify] 更新租客 %d 失败 err=%+v", hirerId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	sysLogInfof(ctx, "[HirerModify] 更新租客 %d 成功", hirerId)
	return RespJsonSuccess(ctx, nil)
}

// HirerList 租客列表
func HirerList(ctx *gin.Context) error {

	pageSize := GetParamIntDef(ctx, "page_size", 20)
	pageNum := GetParamIntDef(ctx, "page_num", 1)
	keywordsStr := GetParamStringDef(ctx, "keywords", "")
	isAll := GetParamIntDef(ctx, "is_all", 0)

	var keywords *entity.HirerKeywords
	if len(keywordsStr) > 0 {
		jErr := json.Unmarshal([]byte(keywordsStr), &keywords)
		if jErr != nil {
			logger.WithContext(ctx).Errorf("[HirerList] GetHirersByLimit err=%s", jErr.Error())
		}
	}

	serviceHirer := service.NewHirer(ctx)
	// 获取所有的租客列表
	if isAll == 1 {
		hirers, err := serviceHirer.GetAllHirers()
		if err != nil {
			sysLogErrorf(ctx, "[HirerList] 获取租客列表失败: err=%+v", err)
			return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
		}
		data := map[string]interface{}{
			"list": hirers,
		}
		return RespJsonSuccess(ctx, data)
	}

	// 分页获取租客列表
	hirers, err := serviceHirer.GetHirersByLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[HirerList] 获取租客列表失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取分页信息
	pageInfo, err := serviceHirer.GetPageInfoLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[HirerList] 获取租客分页信息失败: err=%+v", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	data := map[string]interface{}{
		"list":      hirers,
		"page_info": pageInfo,
	}
	return RespJsonSuccess(ctx, data)
}

// HirerDelete 租客删除
func HirerDelete(ctx *gin.Context) error {

	hirerId := GetParamInt64(ctx, "hirer_id")

	// 判断参数合法性
	if hirerId == 0 {
		logger.WithContext(ctx).Warnf("[HirerDelete] hirer_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "租客id不存在")
	}

	// 更新租客状态
	err := service.NewHirer(ctx).DeleteByHirerId(hirerId)
	if err != nil {
		sysLogErrorf(ctx, "[HirerDelete] 删除租客失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[HirerDelete] 删除租客 %d 成功", hirerId)

	return RespJsonSuccess(ctx, nil)
}

// HirerDetail 租客详情
func HirerDetail(ctx *gin.Context) error {

	hirerId := GetParamInt64(ctx, "hirer_id")

	// 判断参数合法性
	if hirerId == 0 {
		logger.WithContext(ctx).Warnf("[LandlorDetail] hirer_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "租客id不存在")
	}

	// 获取租客详情
	hirer, err := service.NewHirer(ctx).GetHirerByHirerId(hirerId)
	if err != nil {
		sysLogErrorf(ctx, "[LandlorDetail] 获取租客详情失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if hirer == nil {
		sysLogWarnf(ctx, "[LandlorDetail] 租客id %d 不存在", hirerId)
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "租客id不存在")
	}
	return RespJsonSuccess(ctx, &entity.HirerDetailResp{
		HirerInfo: hirer,
	})
}
