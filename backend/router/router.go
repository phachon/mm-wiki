package router

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/controller"
	homeController "github.com/phachon/mm-wiki/app/controller/home"
	spaceController "github.com/phachon/mm-wiki/app/controller/space"
	systemController "github.com/phachon/mm-wiki/app/controller/system"
	userController "github.com/phachon/mm-wiki/app/controller/user"
	"github.com/phachon/mm-wiki/config"
	"github.com/phachon/mm-wiki/filter"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/gopkg/upload"
)

// router 路由相关

const (
	routerGroupNameHome   = "/home"
	routerGroupNameSpace  = "/space"
	routerGroupNameUser   = "/user"
	routerGroupNameSystem = "/system"
)

var (
	// routerHandleTables 路由处理表，新增一个接口配置一条
	routerHandleTables = []routerHandle{
		// ===================== 系统 =====================
		// 登录
		{group: routerGroupNameSystem, relativePath: "/auth/captcha", method: http.MethodGet, controllerHandle: systemController.AuthCaptcha},
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
		// 邮箱管理
		{group: routerGroupNameSystem, relativePath: "/email/save", method: http.MethodPost, controllerHandle: systemController.EmailSave},
		{group: routerGroupNameSystem, relativePath: "/email/edit", method: http.MethodGet, controllerHandle: systemController.EmailEdit},
		{group: routerGroupNameSystem, relativePath: "/email/modify", method: http.MethodPost, controllerHandle: systemController.EmailModify},
		{group: routerGroupNameSystem, relativePath: "/email/list", method: http.MethodGet, controllerHandle: systemController.EmailList},
		{group: routerGroupNameSystem, relativePath: "/email/delete", method: http.MethodPost, controllerHandle: systemController.EmailDelete},
		{group: routerGroupNameSystem, relativePath: "/email/used", method: http.MethodPost, controllerHandle: systemController.EmailUsed},
		// 链接管理
		{group: routerGroupNameSystem, relativePath: "/link/save", method: http.MethodPost, controllerHandle: systemController.LinkSave},
		{group: routerGroupNameSystem, relativePath: "/link/edit", method: http.MethodGet, controllerHandle: systemController.LinkEdit},
		{group: routerGroupNameSystem, relativePath: "/link/modify", method: http.MethodPost, controllerHandle: systemController.LinkModify},
		{group: routerGroupNameSystem, relativePath: "/link/list", method: http.MethodGet, controllerHandle: systemController.LinkList},
		{group: routerGroupNameSystem, relativePath: "/link/delete", method: http.MethodPost, controllerHandle: systemController.LinkDelete},
		// 联系人管理
		{group: routerGroupNameSystem, relativePath: "/contact/save", method: http.MethodPost, controllerHandle: systemController.ContactSave},
		{group: routerGroupNameSystem, relativePath: "/contact/edit", method: http.MethodGet, controllerHandle: systemController.ContactEdit},
		{group: routerGroupNameSystem, relativePath: "/contact/modify", method: http.MethodPost, controllerHandle: systemController.ContactModify},
		{group: routerGroupNameSystem, relativePath: "/contact/list", method: http.MethodGet, controllerHandle: systemController.ContactList},
		{group: routerGroupNameSystem, relativePath: "/contact/delete", method: http.MethodPost, controllerHandle: systemController.ContactDelete},
		// 系统配置
		{group: routerGroupNameSystem, relativePath: "/config/list", method: http.MethodGet, controllerHandle: systemController.ConfigList},
		{group: routerGroupNameSystem, relativePath: "/config/modify", method: http.MethodPost, controllerHandle: systemController.ConfigModify},
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
		// 设置
		{group: routerGroupNameSpace, relativePath: "/setting/basic_modify", method: http.MethodPost, controllerHandle: spaceController.SpaceBasicSettingModify},
		{group: routerGroupNameSpace, relativePath: "/setting/permission_list", method: http.MethodGet, controllerHandle: spaceController.SpacePermissionList},
		// 文档
		{group: routerGroupNameSpace, relativePath: "/doc/create", method: http.MethodPost, controllerHandle: spaceController.DocCreate},
		{group: routerGroupNameSpace, relativePath: "/doc/info", method: http.MethodGet, controllerHandle: spaceController.DocInfo},
		{group: routerGroupNameSpace, relativePath: "/doc/content_save", method: http.MethodPost, controllerHandle: spaceController.DocContentSave},
		{group: routerGroupNameSpace, relativePath: "/doc/upload_file", method: http.MethodPost, controllerHandle: spaceController.DocUploadFile},
		{group: routerGroupNameSpace, relativePath: "/doc/history", method: http.MethodGet, controllerHandle: spaceController.DocHistoryList},
		{group: routerGroupNameSpace, relativePath: "/doc/content_version", method: http.MethodGet, controllerHandle: spaceController.DocContentVersion},
		{group: routerGroupNameSpace, relativePath: "/doc/recover", method: http.MethodPost, controllerHandle: spaceController.DocRecover},
		{group: routerGroupNameSpace, relativePath: "/doc/content_version_del", method: http.MethodPost, controllerHandle: spaceController.DocContentVersionDel},
		// 文档排序
		{group: routerGroupNameSpace, relativePath: "/doc/sort", method: http.MethodPost, controllerHandle: spaceController.DocSort},
		// 文档删除
		{group: routerGroupNameSpace, relativePath: "/doc/delete", method: http.MethodPost, controllerHandle: spaceController.DocDelete},
		// 文档移动
		{group: routerGroupNameSpace, relativePath: "/doc/move", method: http.MethodPost, controllerHandle: spaceController.DocMove},
		// 搜索
		{group: routerGroupNameSpace, relativePath: "/doc/search", method: http.MethodGet, controllerHandle: spaceController.DocSearch},
		// ===================== 用户 =====================
		// 互动
		{group: routerGroupNameUser, relativePath: "/interaction/collection", method: http.MethodPost, controllerHandle: userController.CollectionAdd},
		{group: routerGroupNameUser, relativePath: "/interaction/collection_cancel", method: http.MethodPost, controllerHandle: userController.CollectionCancel},
		{group: routerGroupNameUser, relativePath: "/interaction/collection_status", method: http.MethodGet, controllerHandle: userController.CollectionStatus},
		// 账号
		{group: routerGroupNameUser, relativePath: "/account/list", method: http.MethodGet, controllerHandle: userController.AccountList},
		{group: routerGroupNameUser, relativePath: "/department/list", method: http.MethodGet, controllerHandle: userController.DepartmentList},
		// 关注
		{group: routerGroupNameUser, relativePath: "/interaction/follow", method: http.MethodPost, controllerHandle: userController.FollowAdd},
		{group: routerGroupNameUser, relativePath: "/interaction/follow_cancel", method: http.MethodPost, controllerHandle: userController.FollowCancel},
		{group: routerGroupNameUser, relativePath: "/interaction/follow_status", method: http.MethodGet, controllerHandle: userController.FollowStatus},

		// ===================== 首页 =====================
		// 首页
		{group: routerGroupNameHome, relativePath: "/my_spaces", method: http.MethodGet, controllerHandle: homeController.GetMySpaces},
		{group: routerGroupNameHome, relativePath: "/collection_spaces", method: http.MethodGet, controllerHandle: homeController.GetCollectionSpaces},
		{group: routerGroupNameHome, relativePath: "/collection_docs", method: http.MethodGet, controllerHandle: homeController.GetCollectionDocs},
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
	// CORS 配置
	corsConf := config.GetAppConf().GetCORSConf()
	if len(corsConf.AllowOrigins) > 0 {
		corsConfig := cors.Config{
			AllowOrigins:     corsConf.AllowOrigins,
			AllowMethods:     corsConf.AllowMethods,
			AllowHeaders:     corsConf.AllowHeaders,
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}
		if len(corsConfig.AllowMethods) == 0 {
			corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
		}
		if len(corsConfig.AllowHeaders) == 0 {
			corsConfig.AllowHeaders = []string{
				"Origin", "Content-Type", "Accept",
				global.HeaderKTRequestID, global.HeaderKTTimestamp,
				global.HeaderKTLoginToken, global.HeaderKTDebug,
			}
		}
		global.GinEngine.Use(cors.New(corsConfig))
	} else {
		global.GinEngine.Use(cors.Default())
	}
	global.GinEngine.Use(gin.Logger())
	global.GinEngine.Use(gin.Recovery())
	global.GinEngine.Use(filter.RateLimit())
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
