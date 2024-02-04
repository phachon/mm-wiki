package cache

import (
	"context"

	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
)

// PermissionCache 权限缓存
type PermissionCache struct {
	ctx context.Context
}

// NewPermissionCache 创建权限标识缓存
func NewPermissionCache(ctx context.Context) *PermissionCache {
	permission := &PermissionCache{
		ctx: ctx,
	}
	return permission
}

// GetAccountPrivileges 读取账号权限信息
func (p *PermissionCache) GetAccountPrivileges(accountId int64) ([]*entity.PrivilegeEntity, errors.BizError) {

	return []*entity.PrivilegeEntity{}, nil
}

// SetAccountPrivileges 写入账号权限信息
func (p *PermissionCache) SetAccountPrivileges(accountId int64, privileges []*entity.PrivilegeEntity) errors.BizError {

	return nil
}
