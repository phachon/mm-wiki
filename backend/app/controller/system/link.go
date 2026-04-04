package system

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
)

// LinkSave 添加链接保存
func LinkSave(ctx *gin.Context) error {

	name := GetParamString(ctx, "name")
	url := GetParamString(ctx, "url")
	sequence := GetParamInt(ctx, "sequence")
	// 判断参数合法性
	if name == "" {
		logger.WithContext(ctx).Warnf("[LinkSave] name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "链接名称不能为空")
	}
	if url == "" {
		logger.WithContext(ctx).Warnf("[LinkSave] url empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "链接地址不能为空")
	}

	// link 链接实体
	linkEntity := &entity.LinkEntity{
		Name:     name,
		URL:      url,
		Sequence: sequence,
		Status:   entity.LinkStatusDefault,
	}
	// 创建链接
	err := service.NewLink(ctx).Create(linkEntity)
	if err != nil {
		sysLogErrorf(ctx, "[LinkSave] 添加链接失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[LinkSave] 添加链接 %d 成功", linkEntity.LinkId)

	return RespJsonSuccess(ctx, nil)
}

// LinkEdit 链接修改页面
func LinkEdit(ctx *gin.Context) error {

	linkId := GetParamInt64(ctx, "link_id")
	// 判断参数合法性
	if linkId == 0 {
		logger.WithContext(ctx).Warnf("[LinkEdit] 链接id不存在")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "链接id不存在")
	}

	link, err := service.NewLink(ctx).GetLinkByLinkId(linkId)
	if err != nil {
		sysLogErrorf(ctx, "[LinkEdit] 获取链接 %d 信息失败: err=%+v", linkId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	if link == nil {
		logger.WithContext(ctx).Warnf("[LinkEdit] 链接 %d 不存在", linkId)
		return RespJsonError(ctx, int32(errors.ClientReqParamWrongful), "链接id错误")
	}

	return RespJsonSuccess(ctx, map[string]interface{}{
		"link_info": link,
	})
}

// LinkModify 修改链接保存
func LinkModify(ctx *gin.Context) error {

	linkId := GetParamInt64(ctx, "link_id")
	name := GetParamString(ctx, "name")
	url := GetParamString(ctx, "url")
	sequence := GetParamInt(ctx, "sequence")
	// 判断参数合法性
	if linkId <= 0 {
		logger.WithContext(ctx).Warnf("[LinkModify] link_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "链接ID不合法")
	}
	if name == "" {
		logger.WithContext(ctx).Warnf("[LinkModify] name empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "链接名称不能为空")
	}
	if url == "" {
		logger.WithContext(ctx).Warnf("[LinkModify] url empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "链接地址不能为空")
	}

	// link 链接实体
	linkEntity := entity.LinkEntity{
		LinkId:   linkId,
		Name:     name,
		URL:      url,
		Sequence: sequence,
	}

	// 更新链接
	err := service.NewLink(ctx).Update(linkEntity)
	if err != nil {
		sysLogErrorf(ctx, "[LinkModify] 更新链接 %d 失败: err=%+v", linkId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[LinkModify] 更新链接 %d 成功", linkId)

	return RespJsonSuccess(ctx, nil)
}

// LinkList 链接列表
func LinkList(ctx *gin.Context) error {

	pageSize := GetParamIntDef(ctx, "page_size", 20)
	pageNum := GetParamIntDef(ctx, "page_num", 1)
	keywordsStr := GetParamStringDef(ctx, "keywords", "")

	var keywords *entity.LinkKeywords
	if len(keywordsStr) > 0 {
		jErr := json.Unmarshal([]byte(keywordsStr), &keywords)
		if jErr != nil {
			logger.WithContext(ctx).Errorf("[LinkList] GetLinksByLimit err=%s", jErr.Error())
		}
	}

	linkService := service.NewLink(ctx)

	// 获取链接列表
	links, err := linkService.GetLinksByLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[LinkList] 获取链接列表失败: err=%+v", err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取分页信息
	pageInfo, err := linkService.GetPageInfoLimit(pageSize, pageNum, keywords)
	if err != nil {
		sysLogErrorf(ctx, "[LinkList] 获取链接分页信息失败: err=%s", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 格式化列表结构
	linkList, err := linkService.FormatLinkList(links)
	if err != nil {
		sysLogErrorf(ctx, "[LinkList] 格式化链接列表失败: err=%s", err.Error())
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	data := map[string]interface{}{
		"list":      linkList,
		"page_info": pageInfo,
	}
	return RespJsonSuccess(ctx, data)
}

// LinkDelete 链接删除
func LinkDelete(ctx *gin.Context) error {

	linkId := GetParamInt64(ctx, "link_id")

	// 判断参数合法性
	if linkId <= 0 {
		logger.WithContext(ctx).Warnf("[LinkDelete] link_id empty")
		return RespJsonError(ctx, int32(errors.ClientReqParamEmpty), "链接id不存在")
	}

	// 删除链接
	err := service.NewLink(ctx).DeleteLink(linkId)
	if err != nil {
		sysLogErrorf(ctx, "[LinkDelete] 删除链接 %d 失败: err=%+v", linkId, err)
		return RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	sysLogInfof(ctx, "[LinkDelete] 删除链接 %d 成功", linkId)

	return RespJsonSuccess(ctx, nil)
}
