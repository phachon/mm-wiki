package filter

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/gopkg/errors"
)

var (
	// 需要忽略的校验用户权限的接口
	ignorePermissionPath = map[string]bool{
		"/system/auth/login": true,
	}
)

// PermissionCheck 校验用户的接口权限
func PermissionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		if authPermissionHandle(c) {
			c.Next()
		} else {
			c.Abort()
		}
	}
}

func authPermissionHandle(ctx *gin.Context) (keepNext bool) {
	// 需要忽略权限的接口
	path := ctx.Request.URL.Path
	_, ok := ignorePermissionPath[path]
	if ok {
		return true
	}

	res := &entity.Response{
		Code:    int32(errors.SuccessCode),
		Message: "",
		Data:    nil,
	}
	// 获取登录的用户ID
	loginAccountId := global.ContextValueLoginAccountID(ctx)
	if loginAccountId <= 0 {
		res.Code = int32(errors.ClientReqCommonParamErr)
		ctx.JSON(http.StatusForbidden, res) // 无访问权限
		return false
	}

	permissionService := service.NewPermission(ctx)
	// 获取账号所有的权限
	privileges, err := permissionService.GetAccountPrivilegesCache(loginAccountId)
	if err != nil {
		res.Code = int32(errors.ClientUnknownError)
		ctx.JSON(http.StatusForbidden, res) // 无访问权限
		return false
	}
	// ctx 传递权限标识
	identifys := permissionService.GetPrivilegesIdentifys(privileges)
	global.ContextWithLoginIdentifys(ctx, identifys)

	// 从 ctx 获取 debug 参数 跳过权限
	debugParam := global.ContextValueDebugParam(ctx)
	if debugParam != nil && debugParam.IsSkipPremission() {
		return true
	}
	// 校验接口访问权限
	if checkUrlPathPermission(path, privileges) {
		return true
	}
	// 无访问权限
	res.Code = int32(errors.ClientUnknownError)
	ctx.JSON(http.StatusForbidden, res)
	return false
}

func checkUrlPathPermission(urlPath string, privileges []*entity.PrivilegeEntity) bool {
	urlPathMark := service.GetApiMarkByRouterPath(urlPath)
	for _, privilege := range privileges {
		apiMarksStr := privilege.ApiMarks
		if apiMarksStr == "" {
			continue
		}
		apiMarksList := strings.Split(apiMarksStr, ",")
		for _, apiMark := range apiMarksList {
			if apiMark == urlPathMark {
				return true
			}
		}
	}
	return false
}
