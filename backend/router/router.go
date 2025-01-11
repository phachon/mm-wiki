package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/controller"
	spaceController "github.com/phachon/mm-wiki/app/controller/space"
	systemController "github.com/phachon/mm-wiki/app/controller/system"
	"github.com/phachon/mm-wiki/config"
	"github.com/phachon/mm-wiki/filter"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/gopkg/upload"
)

// router 路由相关

const (
	routerGroupNameSystem = "/system"
	routerGroupNameSpace  = "/space"
	routerGroupNameDoc    = "/doc"
)

var (
	// routerHandleTables 路由处理表，新增一个接口配置一条
	routerHandleTables = []routerHandle{
		// ===================== 系统 =====================
		// 登录
		{group: routerGroupNameSystem, relativePath: "/auth/login", method: http.MethodPost, controllerHandle: systemController.AuthLogin},
		// 个人中心
		{group: routerGroupNameSystem, relativePath: "/profile/privileges", method: http.MethodGet, controllerHandle: systemController.ProfilePrivileges},
		{group: routerGroupNameSystem, relativePath: "/profile/info", method: http.MethodGet, controllerHandle: systemController.ProfileInfo},
		{group: routerGroupNameSystem, relativePath: "/profile/repass", method: http.MethodPost, controllerHandle: systemController.ProfileRePass},
		{group: routerGroupNameSystem, relativePath: "/profile/update", method: http.MethodPost, controllerHandle: systemController.ProfileUpdate},
		// 账号管理
		{group: routerGroupNameSystem, relativePath: "/account/add", method: http.MethodGet, controllerHandle: systemController.AccountAdd},
		{group: routerGroupNameSystem, relativePath: "/account/save", method: http.MethodPost, controllerHandle: systemController.AccountSave},
		{group: routerGroupNameSystem, relativePath: "/account/edit", method: http.MethodGet, controllerHandle: systemController.AccountEdit},
		{group: routerGroupNameSystem, relativePath: "/account/modify", method: http.MethodPost, controllerHandle: systemController.AccountModify},
		{group: routerGroupNameSystem, relativePath: "/account/update_status", method: http.MethodPost, controllerHandle: systemController.AccountUpdateStatus},
		{group: routerGroupNameSystem, relativePath: "/account/detail", method: http.MethodGet, controllerHandle: systemController.AccountDetail},
		{group: routerGroupNameSystem, relativePath: "/account/list", method: http.MethodGet, controllerHandle: systemController.AccountList},
		// 角色管理
		{group: routerGroupNameSystem, relativePath: "/role/add", method: http.MethodGet, controllerHandle: systemController.RoleAdd},
		{group: routerGroupNameSystem, relativePath: "/role/save", method: http.MethodPost, controllerHandle: systemController.RoleSave},
		{group: routerGroupNameSystem, relativePath: "/role/edit", method: http.MethodGet, controllerHandle: systemController.RoleEdit},
		{group: routerGroupNameSystem, relativePath: "/role/modify", method: http.MethodPost, controllerHandle: systemController.RoleModify},
		{group: routerGroupNameSystem, relativePath: "/role/list", method: http.MethodGet, controllerHandle: systemController.RoleList},
		{group: routerGroupNameSystem, relativePath: "/role/account_list", method: http.MethodGet, controllerHandle: systemController.RoleAccountList},
		{group: routerGroupNameSystem, relativePath: "/role/account_remove", method: http.MethodPost, controllerHandle: systemController.RoleAccountRemove},
		{group: routerGroupNameSystem, relativePath: "/role/privilege_edit", method: http.MethodGet, controllerHandle: systemController.RolePrivilegeEdit},
		{group: routerGroupNameSystem, relativePath: "/role/privilege_modify", method: http.MethodPost, controllerHandle: systemController.RolePrivilegeModify},
		{group: routerGroupNameSystem, relativePath: "/role/delete", method: http.MethodPost, controllerHandle: systemController.RoleDelete},
		// 权限管理
		{group: routerGroupNameSystem, relativePath: "/privilege/add", method: http.MethodGet, controllerHandle: systemController.PrivilegeAdd},
		{group: routerGroupNameSystem, relativePath: "/privilege/save", method: http.MethodPost, controllerHandle: systemController.PrivilegeSave},
		{group: routerGroupNameSystem, relativePath: "/privilege/edit", method: http.MethodGet, controllerHandle: systemController.PrivilegeEdit},
		{group: routerGroupNameSystem, relativePath: "/privilege/modify", method: http.MethodPost, controllerHandle: systemController.PrivilegeModify},
		{group: routerGroupNameSystem, relativePath: "/privilege/list", method: http.MethodGet, controllerHandle: systemController.PrivilegeList},
		{group: routerGroupNameSystem, relativePath: "/privilege/delete", method: http.MethodPost, controllerHandle: systemController.PrivilegeDelete},
		// 日志管理
		{group: routerGroupNameSystem, relativePath: "/log/list", method: http.MethodGet, controllerHandle: systemController.LogList},
		// 公告管理
		{group: routerGroupNameSystem, relativePath: "/notice/save", method: http.MethodPost, controllerHandle: systemController.NoticeSave},
		{group: routerGroupNameSystem, relativePath: "/notice/edit", method: http.MethodGet, controllerHandle: systemController.NoticeEdit},
		{group: routerGroupNameSystem, relativePath: "/notice/modify", method: http.MethodPost, controllerHandle: systemController.NoticeModify},
		{group: routerGroupNameSystem, relativePath: "/notice/list", method: http.MethodGet, controllerHandle: systemController.NoticeList},
		{group: routerGroupNameSystem, relativePath: "/notice/publish_list", method: http.MethodGet, controllerHandle: systemController.NoticePublishList},
		{group: routerGroupNameSystem, relativePath: "/notice/delete", method: http.MethodPost, controllerHandle: systemController.NoticeDelete},
		// 部门管理
		{group: routerGroupNameSystem, relativePath: "/department/add", method: http.MethodGet, controllerHandle: systemController.DepartmentAdd},
		{group: routerGroupNameSystem, relativePath: "/department/save", method: http.MethodPost, controllerHandle: systemController.DepartmentSave},
		{group: routerGroupNameSystem, relativePath: "/department/edit", method: http.MethodGet, controllerHandle: systemController.DepartmentEdit},
		{group: routerGroupNameSystem, relativePath: "/department/modify", method: http.MethodPost, controllerHandle: systemController.DepartmentModify},
		{group: routerGroupNameSystem, relativePath: "/department/list", method: http.MethodGet, controllerHandle: systemController.DepartmentList},
		{group: routerGroupNameSystem, relativePath: "/department/delete", method: http.MethodPost, controllerHandle: systemController.DepartmentDelete},
		// 空间管理
		{group: routerGroupNameSystem, relativePath: "/space/add", method: http.MethodGet, controllerHandle: systemController.SpaceAdd},
		{group: routerGroupNameSystem, relativePath: "/space/save", method: http.MethodPost, controllerHandle: systemController.SpaceSave},
		{group: routerGroupNameSystem, relativePath: "/space/edit", method: http.MethodGet, controllerHandle: systemController.SpaceEdit},
		{group: routerGroupNameSystem, relativePath: "/space/modify", method: http.MethodPost, controllerHandle: systemController.SpaceModify},
		{group: routerGroupNameSystem, relativePath: "/space/list", method: http.MethodGet, controllerHandle: systemController.SpaceList},
		{group: routerGroupNameSystem, relativePath: "/space/delete", method: http.MethodPost, controllerHandle: systemController.SpaceDelete},
		{group: routerGroupNameSystem, relativePath: "/space/admin_list", method: http.MethodGet, controllerHandle: systemController.SpaceAdminList},
		{group: routerGroupNameSystem, relativePath: "/space/admin_remove", method: http.MethodPost, controllerHandle: systemController.SpaceAdminRemove},
		{group: routerGroupNameSystem, relativePath: "/space/admin_add", method: http.MethodPost, controllerHandle: systemController.SpaceAdminAdd},
		// ===================== 空间 =====================
		// 空间
		{group: routerGroupNameSpace, relativePath: "/space/all", method: http.MethodGet, controllerHandle: spaceController.AllSpaces},
		{group: routerGroupNameSpace, relativePath: "/space/docs", method: http.MethodGet, controllerHandle: spaceController.SpaceDocs},
		// 文档
		{group: routerGroupNameSpace, relativePath: "/doc/create", method: http.MethodPost, controllerHandle: spaceController.DocCreate},
		{group: routerGroupNameSpace, relativePath: "/doc/info", method: http.MethodGet, controllerHandle: spaceController.DocInfo},
		{group: routerGroupNameSpace, relativePath: "/doc/content_save", method: http.MethodPost, controllerHandle: spaceController.DocContentSave},
		{group: routerGroupNameSpace, relativePath: "/doc/upload_file", method: http.MethodPost, controllerHandle: spaceController.DocUploadFile},
		{group: routerGroupNameSpace, relativePath: "/doc/history", method: http.MethodGet, controllerHandle: spaceController.DocHistoryList},
		{group: routerGroupNameSpace, relativePath: "/doc/content_version", method: http.MethodGet, controllerHandle: spaceController.DocContentVersion},
		{group: routerGroupNameSpace, relativePath: "/doc/recover", method: http.MethodPost, controllerHandle: spaceController.DocRecover},
		{group: routerGroupNameSpace, relativePath: "/doc/content_version_del", method: http.MethodPost, controllerHandle: spaceController.DocContentVersionDel},
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

	// 如果 upload 为 local 添加静态文件路由
	for _, uploadConf := range config.GetAppConf().Upload {
		if uploadConf.UploadType == upload.UploaderLocal {
			global.GinEngine.Static("/static/upload", uploadConf.LocalDir)
		}
	}
}

// HandleWrapper 控制器 Wrapper
func HandleWrapper(controllerHandle controller.HandleFunc) func(*gin.Context) {
	return func(context *gin.Context) {
		_ = controllerHandle(context)
	}
}
