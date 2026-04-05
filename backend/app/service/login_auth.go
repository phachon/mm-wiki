package service

import (
	"context"

	"github.com/phachon/mm-wiki/app/dao"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/global"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
	"github.com/phachon/mm-wiki/utils"
)

const (
	LoginAuthListDefaultPageSize = 20 // 认证列表默认一页 20 条
)

// LoginAuth 登录认证业务逻辑
type LoginAuth struct {
	ctx          context.Context
	daoLoginAuth *dao.LoginAuth
}

// NewLoginAuth 创建登录认证业务逻辑对象
func NewLoginAuth(ctx context.Context) *LoginAuth {
	return &LoginAuth{
		ctx:          ctx,
		daoLoginAuth: dao.NewLoginAuth(ctx),
	}
}

// Create 创建登录认证
func (la *LoginAuth) Create(loginAuthEntity *entity.LoginAuthEntity) errors.BizError {
	// 检查名称是否已存在
	existAuth, err := la.daoLoginAuth.GetLoginAuthByName(loginAuthEntity.Name)
	if err != nil {
		return err
	}
	if existAuth != nil {
		return errors.Errorf(errors.BusinessRecordExistError, "认证名称已存在")
	}
	// 插入一条记录
	err = la.daoLoginAuth.Insert(loginAuthEntity)
	if err != nil {
		return err
	}
	return nil
}

// Update 修改登录认证
func (la *LoginAuth) Update(loginAuthEntity entity.LoginAuthEntity) errors.BizError {
	// 查找认证是否存在
	updateAuth, err := la.daoLoginAuth.GetLoginAuthByLoginAuthId(loginAuthEntity.LoginAuthId)
	if err != nil {
		return err
	}
	if updateAuth == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "认证不存在")
	}
	// 检查名称是否已被其他认证使用
	existAuth, err := la.daoLoginAuth.GetLoginAuthByName(loginAuthEntity.Name)
	if err != nil {
		return err
	}
	if existAuth != nil && existAuth.LoginAuthId != loginAuthEntity.LoginAuthId {
		return errors.Errorf(errors.BusinessRecordExistError, "认证名称已存在")
	}
	// 更新字段
	err = la.daoLoginAuth.Update(loginAuthEntity)
	if err != nil {
		return err
	}
	return nil
}

// GetLoginAuthByLoginAuthId 根据认证ID获取认证详情
func (la *LoginAuth) GetLoginAuthByLoginAuthId(loginAuthId int64) (loginAuth *entity.LoginAuthEntity, err errors.BizError) {
	loginAuth, err = la.daoLoginAuth.GetLoginAuthByLoginAuthId(loginAuthId)
	if err != nil {
		return loginAuth, err
	}
	return loginAuth, nil
}

// GetUsedLoginAuth 获取当前使用的认证配置
func (la *LoginAuth) GetUsedLoginAuth() (loginAuth *entity.LoginAuthEntity, err errors.BizError) {
	return la.daoLoginAuth.GetUsedLoginAuth()
}

// SetUsed 设置认证为使用中（先清除其他使用状态再设置）
func (la *LoginAuth) SetUsed(loginAuthId int64) errors.BizError {
	// 查找认证是否存在
	auth, err := la.daoLoginAuth.GetLoginAuthByLoginAuthId(loginAuthId)
	if err != nil {
		return err
	}
	if auth == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "认证id %d 不存在", loginAuthId)
	}
	// 清除所有使用状态
	err = la.daoLoginAuth.ClearUsed()
	if err != nil {
		return err
	}
	// 设置当前认证为使用中
	err = la.daoLoginAuth.UpdateUsed(loginAuthId, entity.LoginAuthIsUsedYes)
	if err != nil {
		return err
	}
	return nil
}

// DeleteLoginAuth 删除登录认证
func (la *LoginAuth) DeleteLoginAuth(loginAuthId int64) errors.BizError {
	// 查找认证是否存在
	updateAuth, err := la.daoLoginAuth.GetLoginAuthByLoginAuthId(loginAuthId)
	if err != nil {
		return err
	}
	if updateAuth == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "认证id %d 不存在", loginAuthId)
	}
	// 不能删除使用中的认证
	if updateAuth.IsUsed == entity.LoginAuthIsUsedYes {
		return errors.Errorf(errors.BusinessForbiddenError, "不能删除使用中的认证")
	}
	// 删除认证
	err = la.daoLoginAuth.DeleteLoginAuth(loginAuthId)
	if err != nil {
		return err
	}
	return nil
}

// GetLoginAuthsByLimit 分页获取认证列表
func (la *LoginAuth) GetLoginAuthsByLimit(pageSize int, pageNum int, keywords *entity.LoginAuthKeywords) (loginAuths []*entity.LoginAuthEntity, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, LoginAuthListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	offset := (pageNum - 1) * pageSize
	if keywords == nil {
		return la.daoLoginAuth.GetLoginAuthsByLimit(pageSize, offset)
	}
	return la.daoLoginAuth.GetLoginAuthsByKeywordsAndLimit(pageSize, offset, keywords)
}

// GetPageInfoLimit 获取分页信息
func (la *LoginAuth) GetPageInfoLimit(pageSize int, pageNum int, keywords *entity.LoginAuthKeywords) (pagination *entity.PageInfo, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, LoginAuthListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	pageInfo := new(entity.PageInfo)
	var totalNum int64
	if keywords == nil {
		totalNum, err = la.daoLoginAuth.CountLoginAuths()
	} else {
		totalNum, err = la.daoLoginAuth.CountLoginAuthsByKeywords(keywords)
	}
	if err != nil {
		logger.WithContext(la.ctx).Warnf("[service.LoginAuth] GetPageInfoLimit pageSize=%d pageNum=%d err=%s",
			pageSize, pageNum, err.Error())
		return pageInfo, err
	}
	return entity.GetPageInfo(totalNum, pageSize, pageNum), nil
}

// FormatLoginAuthList 格式化认证列表
func (la *LoginAuth) FormatLoginAuthList(loginAuths []*entity.LoginAuthEntity) (loginAuthList []*entity.LoginAuthListItem, err errors.BizError) {
	if len(loginAuths) == 0 {
		return loginAuthList, nil
	}
	for _, loginAuth := range loginAuths {
		action := la.GetListItemAction(loginAuth)
		var loginAuthListItem = &entity.LoginAuthListItem{
			LoginAuthEntity: loginAuth,
			Action:          action,
		}
		loginAuthList = append(loginAuthList, loginAuthListItem)
	}
	return loginAuthList, nil
}

// GetListItemAction 获取列表 action 权限
func (la *LoginAuth) GetListItemAction(loginAuthItem *entity.LoginAuthEntity) *entity.LoginAuthListAction {
	indentifys := global.ContextValueLoginIdentifys(la.ctx)
	action := &entity.LoginAuthListAction{
		IsEdit:   indentifys[global.PrivilegeIndentifyLoginAuthEdit],
		IsDelete: indentifys[global.PrivilegeIndentifyLoginAuthDelete],
	}
	return action
}

// GetAllLoginAuths 获取所有的认证
func (la *LoginAuth) GetAllLoginAuths() (loginAuths []*entity.LoginAuthEntity, err errors.BizError) {
	return la.daoLoginAuth.GetAllLoginAuths()
}
