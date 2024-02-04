package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/controller"
	"github.com/phachon/mm-wiki/filter"
	"github.com/phachon/mm-wiki/global"
)

// router 路由相关

const (
	routerGroupNameSystem = "/system"
)

var (
	// routerHandleTables 路由处理表，新增一个接口配置一条
	routerHandleTables = []routerHandle{
		// 登录
		{group: routerGroupNameSystem, relativePath: "/auth/login", method: http.MethodPost, controllerHandle: controller.AuthLogin},
		// 个人中心
		{group: routerGroupNameSystem, relativePath: "/profile/info", method: http.MethodGet, controllerHandle: controller.ProfileInfo},
		{group: routerGroupNameSystem, relativePath: "/profile/repass", method: http.MethodPost, controllerHandle: controller.ProfileRePass},
		{group: routerGroupNameSystem, relativePath: "/profile/update", method: http.MethodPost, controllerHandle: controller.ProfileUpdate},
		// 账号管理
		{group: routerGroupNameSystem, relativePath: "/account/add", method: http.MethodGet, controllerHandle: controller.AccountAdd},
		{group: routerGroupNameSystem, relativePath: "/account/save", method: http.MethodPost, controllerHandle: controller.AccountSave},
		{group: routerGroupNameSystem, relativePath: "/account/edit", method: http.MethodGet, controllerHandle: controller.AccountEdit},
		{group: routerGroupNameSystem, relativePath: "/account/modify", method: http.MethodPost, controllerHandle: controller.AccountModify},
		{group: routerGroupNameSystem, relativePath: "/account/update_status", method: http.MethodPost, controllerHandle: controller.AccountUpdateStatus},
		{group: routerGroupNameSystem, relativePath: "/account/detail", method: http.MethodGet, controllerHandle: controller.AccountDetail},
		{group: routerGroupNameSystem, relativePath: "/account/list", method: http.MethodGet, controllerHandle: controller.AccountList},
		// 角色管理
		{group: routerGroupNameSystem, relativePath: "/role/add", method: http.MethodGet, controllerHandle: controller.RoleAdd},
		{group: routerGroupNameSystem, relativePath: "/role/save", method: http.MethodPost, controllerHandle: controller.RoleSave},
		{group: routerGroupNameSystem, relativePath: "/role/edit", method: http.MethodGet, controllerHandle: controller.RoleEdit},
		{group: routerGroupNameSystem, relativePath: "/role/modify", method: http.MethodPost, controllerHandle: controller.RoleModify},
		{group: routerGroupNameSystem, relativePath: "/role/list", method: http.MethodGet, controllerHandle: controller.RoleList},
		{group: routerGroupNameSystem, relativePath: "/role/account_list", method: http.MethodGet, controllerHandle: controller.RoleAccountList},
		{group: routerGroupNameSystem, relativePath: "/role/account_remove", method: http.MethodPost, controllerHandle: controller.RoleAccountRemove},
		{group: routerGroupNameSystem, relativePath: "/role/privilege_edit", method: http.MethodGet, controllerHandle: controller.RolePrivilegeEdit},
		{group: routerGroupNameSystem, relativePath: "/role/privilege_modify", method: http.MethodPost, controllerHandle: controller.RolePrivilegeModify},
		{group: routerGroupNameSystem, relativePath: "/role/delete", method: http.MethodPost, controllerHandle: controller.RoleDelete},
		// 权限管理
		{group: routerGroupNameSystem, relativePath: "/privilege/add", method: http.MethodGet, controllerHandle: controller.PrivilegeAdd},
		{group: routerGroupNameSystem, relativePath: "/privilege/save", method: http.MethodPost, controllerHandle: controller.PrivilegeSave},
		{group: routerGroupNameSystem, relativePath: "/privilege/edit", method: http.MethodGet, controllerHandle: controller.PrivilegeEdit},
		{group: routerGroupNameSystem, relativePath: "/privilege/modify", method: http.MethodPost, controllerHandle: controller.PrivilegeModify},
		{group: routerGroupNameSystem, relativePath: "/privilege/list", method: http.MethodGet, controllerHandle: controller.PrivilegeList},
		{group: routerGroupNameSystem, relativePath: "/privilege/delete", method: http.MethodPost, controllerHandle: controller.PrivilegeDelete},
		// 日志管理
		{group: routerGroupNameSystem, relativePath: "/log/list", method: http.MethodGet, controllerHandle: controller.LogList},
		// 公告管理
		{group: routerGroupNameSystem, relativePath: "/notice/save", method: http.MethodPost, controllerHandle: controller.NoticeSave},
		{group: routerGroupNameSystem, relativePath: "/notice/edit", method: http.MethodGet, controllerHandle: controller.NoticeEdit},
		{group: routerGroupNameSystem, relativePath: "/notice/modify", method: http.MethodPost, controllerHandle: controller.NoticeModify},
		{group: routerGroupNameSystem, relativePath: "/notice/list", method: http.MethodGet, controllerHandle: controller.NoticeList},
		{group: routerGroupNameSystem, relativePath: "/notice/publish_list", method: http.MethodGet, controllerHandle: controller.NoticePublishList},
		{group: routerGroupNameSystem, relativePath: "/notice/delete", method: http.MethodPost, controllerHandle: controller.NoticeDelete},
	}
)

// routerHandle 路由处理配置
type routerHandle struct {
	group            string
	relativePath     string
	method           string
	controllerHandle controller.HandleFunc
}

// Init 初始化路由
func Init() {
	global.GinEngine.Use(cors.Default())
	global.GinEngine.Use(gin.Logger())
	global.GinEngine.Use(gin.Recovery())
	global.GinEngine.Use(filter.RequestParse())
	global.GinEngine.Use(filter.AuthLoginJWT())
	global.GinEngine.Use(filter.PermissionCheck())
	global.GinEngine.Use(filter.DebugCosTime())
	initRouter()
}

// RegisterHandle 注册一个路由处理
func RegisterHandle(groupPath string, relativePath string, method string, controllerHandle controller.HandleFunc) {
	routerHandleTables = append(routerHandleTables, routerHandle{
		group:            groupPath,
		relativePath:     relativePath,
		method:           method,
		controllerHandle: controllerHandle,
	})
}

// GetRouterTables 获取路由表
func GetRouterTables() []routerHandle {
	return routerHandleTables
}

func initRouter() {
	// 转化称 group => handles
	routerGroupTables := make(map[string][]routerHandle)
	for _, routerHandleItem := range routerHandleTables {
		if _, ok := routerGroupTables[routerHandleItem.group]; !ok {
			routerGroupTables[routerHandleItem.group] = []routerHandle{}
		}
		routerGroupTables[routerHandleItem.group] = append(routerGroupTables[routerHandleItem.group], routerHandleItem)
	}
	for routerGroup, routerHandles := range routerGroupTables {
		groupRouter := global.GinEngine.Group(routerGroup)
		for _, routerHandleItem := range routerHandles {
			groupRouter.Handle(
				routerHandleItem.method,
				routerHandleItem.relativePath,
				HandleWrapper(routerHandleItem.controllerHandle),
			)
		}
	}
}

// HandleWrapper 控制器 Wrapper
func HandleWrapper(controllerHandle controller.HandleFunc) func(*gin.Context) {
	return func(context *gin.Context) {
		_ = controllerHandle(context)
	}
}
