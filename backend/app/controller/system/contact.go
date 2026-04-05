package system

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
)

// ContactSave 添加联系人保存
func ContactSave(ctx *gin.Context) error {

	name := GetParamString(ctx, "name")
	mobile := GetParamString(ctx, "mobile")
	email := GetParamString(ctx, "email")
	position := GetParamString(ctx, "position")
	// 判断参数合法性
	if name == "" {
		logger.WithContext(ctx).Warnf("[ContactSave] name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "联系人名称不能为空")
	}

	// contact 联系人实体
	contactEntity := &entity.ContactEntity{
		Name:     name,
		Mobile:   mobile,
		Email:    email,
		Position: position,
		Status:   entity.ContactStatusDefault,
	}
	// 创建联系人
	err := service.NewContact(ctx).Create(contactEntity)
	if err != nil {
		sysLogErrorf(ctx, "[ContactSave] 添加联系人失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[ContactSave] 添加联系人 %d 成功", contactEntity.ContactId)

	return RespJsonSuccess(ctx, nil)
}

// ContactEdit 联系人修改页面
func ContactEdit(ctx *gin.Context) error {

	contactId := GetParamInt64(ctx, "contact_id")
	// 判断参数合法性
	if contactId == 0 {
		logger.WithContext(ctx).Warnf("[ContactEdit] 联系人id不存在")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "联系人id不存在")
	}

	contact, err := service.NewContact(ctx).GetContactByContactId(contactId)
	if err != nil {
		sysLogErrorf(ctx, "[ContactEdit] 获取联系人 %d 信息失败: err=%+v", contactId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if contact == nil {
		logger.WithContext(ctx).Warnf("[ContactEdit] 联系人 %d 不存在", contactId)
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "联系人id错误")
	}

	return RespJsonSuccess(ctx, map[string]interface{}{
		"contact_info": contact,
	})
}

// ContactModify 修改联系人保存
func ContactModify(ctx *gin.Context) error {

	contactId := GetParamInt64(ctx, "contact_id")
	name := GetParamString(ctx, "name")
	mobile := GetParamString(ctx, "mobile")
	email := GetParamString(ctx, "email")
	position := GetParamString(ctx, "position")
	// 判断参数合法性
	if contactId <= 0 {
		logger.WithContext(ctx).Warnf("[ContactModify] contact_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "联系人ID不合法")
	}
	if name == "" {
		logger.WithContext(ctx).Warnf("[ContactModify] name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "联系人名称不能为空")
	}

	// contact 联系人实体
	contactEntity := entity.ContactEntity{
		ContactId: contactId,
		Name:      name,
		Mobile:    mobile,
		Email:     email,
		Position:  position,
	}

	// 更新联系人
	err := service.NewContact(ctx).Update(contactEntity)
	if err != nil {
		sysLogErrorf(ctx, "[ContactModify] 更新联系人 %d 失败: err=%+v", contactId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[ContactModify] 更新联系人 %d 成功", contactId)

	return RespJsonSuccess(ctx, nil)
}

// ContactList 联系人列表
func ContactList(ctx *gin.Context) error {

	pageSize := GetParamIntDef(ctx, "page_size", 20)
	pageNum := GetParamIntDef(ctx, "page_num", 1)
	keywordsStr := GetParamStringDef(ctx, "keywords", "")

	var keywords *entity.ContactKeywords
	if len(keywordsStr) > 0 {
		jErr := json.Unmarshal([]byte(keywordsStr), &keywords)
		if jErr != nil {
			logger.WithContext(ctx).Errorf("[ContactList] GetContactsByLimit err=%s", jErr.Error())
		}
	}

	contactService := service.NewContact(ctx)

	// 获取联系人列表
	contacts, err := contactService.GetContactsByLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[ContactList] 获取联系人列表失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取分页信息
	pageInfo, err := contactService.GetPageInfoLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[ContactList] 获取联系人分页信息失败: err=%s", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 格式化列表结构
	contactList, err := contactService.FormatContactList(contacts)
	if err != nil {
		sysLogErrorf(ctx, "[ContactList] 格式化联系人列表失败: err=%s", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	data := map[string]interface{}{
		"list":      contactList,
		"page_info": pageInfo,
	}
	return RespJsonSuccess(ctx, data)
}

// ContactDelete 联系人删除
func ContactDelete(ctx *gin.Context) error {

	contactId := GetParamInt64(ctx, "contact_id")

	// 判断参数合法性
	if contactId <= 0 {
		logger.WithContext(ctx).Warnf("[ContactDelete] contact_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "联系人id不存在")
	}

	// 删除联系人
	err := service.NewContact(ctx).DeleteContact(contactId)
	if err != nil {
		sysLogErrorf(ctx, "[ContactDelete] 删除联系人 %d 失败: err=%+v", contactId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[ContactDelete] 删除联系人 %d 成功", contactId)

	return RespJsonSuccess(ctx, nil)
}
