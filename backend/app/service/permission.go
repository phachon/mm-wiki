package service

import (
	"context"
	"strings"

	"github.com/phachon/mm-wiki/app/cache"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
)

var ignoreApiPath = map[string]bool{
	"/admin/auth/login": true,
}

// Permission 权限标识
type Permission struct {
	ctx             context.Context
	permissionCache *cache.PermissionCache
}

// NewPermission 创建权限标识
func NewPermission(ctx context.Context) *Permission {
	permission := &Permission{
		ctx:             ctx,
		permissionCache: cache.NewPermissionCache(ctx),
	}
	return permission
}

// GetAccountPrivileges 获取账号列表权限带缓存
func (p *Permission) GetAccountPrivilegesCache(accountId int64) ([]*entity.PrivilegeEntity, errors.BizError) {
	// 先从缓存获取
	privilegesCache, err := p.permissionCache.GetAccountPrivileges(accountId)
	if err == nil && len(privilegesCache) > 0 {
		return privilegesCache, nil
	}
	logger.WithContext(p.ctx).Errorf("[Permission] GetAccountPrivileges cache len=%+v err=%+v",
		len(privilegesCache), err)
	// 从数据库获取
	privileges, err := p.GetAccountPrivileges(accountId)
	if err != nil {
		return privileges, err
	}
	// 更新缓存
	err = p.permissionCache.SetAccountPrivileges(accountId, privileges)
	if err != nil {
		logger.WithContext(p.ctx).Errorf("[Permission] SetAccountPrivileges cache err=%+v", err)
	}
	return privileges, nil
}

// getAccountPrivileges 获取账号的权限列表
func (p *Permission) GetAccountPrivileges(accountId int64) ([]*entity.PrivilegeEntity, errors.BizError) {
	// 获取账号所有的角色ID
	roleIds, err := NewAccountRole(p.ctx).GetRoleIdsByAccountId(accountId)
	if err != nil {
		return []*entity.PrivilegeEntity{}, err
	}
	// 获取角色所有的权限
	privileges, err := NewRole(p.ctx).GetPrivilegesByRoleIds(roleIds)
	if err != nil {
		return []*entity.PrivilegeEntity{}, err
	}
	return privileges, nil
}

// GetAccountApiMarks 获取账号接口标识
func (p *Permission) GetAccountApiMarks(accountId int64) ([]string, error) {
	privileges, err := p.GetAccountPrivilegesCache(accountId)
	if err != nil {
		return []string{}, err
	}
	// 获取所有的接口标识
	var apiMarks []string
	for _, privilege := range privileges {
		apiMarksStr := privilege.ApiMarks
		if apiMarksStr == "" {
			continue
		}
		apiMarksList := strings.Split(apiMarksStr, ",")
		if len(apiMarksList) > 0 {
			apiMarks = append(apiMarks, apiMarksList...)
		}
	}
	return apiMarks, nil
}

// GetPrivilegesIdentifys 获取权限的标识 map
func (p *Permission) GetPrivilegesIdentifys(privileges []*entity.PrivilegeEntity) map[string]int {
	var indentifys = make(map[string]int)
	for _, privilege := range privileges {
		indentifys[privilege.Identify] = 1
	}
	return indentifys
}

// GetPrivilegesApiMarks 获取权限接口标识 map
func (p *Permission) GetPrivilegesApiMarks(privileges []*entity.PrivilegeEntity) map[string]bool {
	var apiMarks = make(map[string]bool)
	for _, privilege := range privileges {
		apiMarksStr := privilege.ApiMarks
		if apiMarksStr == "" {
			continue
		}
		apiMarksList := strings.Split(apiMarksStr, ",")
		for _, apiMark := range apiMarksList {
			apiMarks[apiMark] = true
		}
	}
	return apiMarks
}

// GetAllApiMarks 获取所有的接口标识
func GetAllApiRouterPath() []string {
	var allPath []string
	routers := global.GinEngine.Routes()
	for _, router := range routers {
		allPath = append(allPath, router.Path)
	}
	return allPath
}

// GetAllApiMarks 获取所有的接口标识
func GetAllApiMarks() []string {
	var apiMarks []string
	allPath := GetAllApiRouterPath()
	for _, path := range allPath {
		if _, ok := ignoreApiPath[path]; ok {
			continue
		}
		apiMark := GetApiMarkByRouterPath(path)
		apiMarks = append(apiMarks, apiMark)
	}
	return apiMarks
}

// RouterPathToMarks 接口 path 转为 api 标识
// 当前采用比较简单的映射方式：格式 /aaa/bbb => aaa:bbb
func GetApiMarkByRouterPath(path string) string {
	if path == "" {
		return ""
	}
	path = strings.TrimLeft(path, "/")
	path = strings.TrimRight(path, "/")
	pathNames := strings.Split(path, "/")
	return strings.Join(pathNames, ":")
}
