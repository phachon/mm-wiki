package user

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/controller"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/logger"
)

// AccountList 账号列表
func AccountList(ctx *gin.Context) error {

	isAll := controller.GetParamIntDef(ctx, "is_all", 0)
	pageSize := controller.GetParamIntDef(ctx, "page_size", 20)
	pageNum := controller.GetParamIntDef(ctx, "page_num", 1)
	keywordsStr := controller.GetParamStringDef(ctx, "keywords", "")

	serviceAccount := service.NewAccount(ctx)

	if isAll == 1 {
		accounts, err := serviceAccount.GetAllNormalAccounts()
		if err != nil {
			return controller.RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
		}
		data := map[string]interface{}{
			"list": accounts,
		}
		return controller.RespJsonSuccess(ctx, data)
	}

	var keywords *entity.AccountKeywords
	if len(keywordsStr) > 0 {
		jErr := json.Unmarshal([]byte(keywordsStr), &keywords)
		if jErr != nil {
			logger.WithContext(ctx).Errorf("[AccountList] GetAccountsByLimit err=%s", jErr.Error())
		}
	}

	// 获取账号列表
	accounts, err := serviceAccount.GetAccountsByLimit(pageSize, pageNum, keywords)
	if err != nil {
		return controller.RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 获取分页信息
	pageInfo, err := serviceAccount.GetPageInfoLimit(pageSize, pageNum, keywords)
	if err != nil {
		return controller.RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}
	// 格式化账号列表
	accountList, err := serviceAccount.FormatAccountList(accounts)
	if err != nil {
		return controller.RespJsonError(ctx, err.GetErrCode(), err.GetErrMsg())
	}

	data := map[string]interface{}{
		"list":      accountList,
		"page_info": pageInfo,
	}
	return controller.RespJsonSuccess(ctx, data)

}
