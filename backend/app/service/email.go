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
	EmailListDefaultPageSize = 20 // 邮箱列表默认一页 20 条
)

// Email 邮箱业务逻辑
type Email struct {
	ctx      context.Context
	daoEmail *dao.Email
}

// NewEmail 创建邮箱业务逻辑对象
func NewEmail(ctx context.Context) *Email {
	return &Email{
		ctx:      ctx,
		daoEmail: dao.NewEmail(ctx),
	}
}

// Create 创建邮箱
func (e *Email) Create(emailEntity *entity.EmailEntity) errors.BizError {
	// 检查邮箱名称是否已存在
	existEmail, err := e.daoEmail.GetEmailByName(emailEntity.Name)
	if err != nil {
		return err
	}
	if existEmail != nil {
		return errors.Errorf(errors.BusinessRecordExistError, "邮箱名 %s 已经存在", emailEntity.Name)
	}
	// 插入一条记录
	err = e.daoEmail.Insert(emailEntity)
	if err != nil {
		return err
	}
	return nil
}

// Update 修改邮箱
func (e *Email) Update(emailEntity entity.EmailEntity) errors.BizError {
	// 查找邮箱是否存在
	updateEmail, err := e.daoEmail.GetEmailByEmailId(emailEntity.EmailId)
	if err != nil {
		return err
	}
	if updateEmail == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "邮箱不存在")
	}
	// 检查邮箱名称是否已被其他邮箱使用
	existEmail, err := e.daoEmail.GetEmailByName(emailEntity.Name)
	if err != nil {
		return err
	}
	if existEmail != nil && existEmail.EmailId != emailEntity.EmailId {
		return errors.Errorf(errors.BusinessRecordExistError, "邮箱名 %s 已经存在", emailEntity.Name)
	}
	// 更新字段
	err = e.daoEmail.Update(emailEntity)
	if err != nil {
		return err
	}
	return nil
}

// GetEmailByEmailId 根据邮箱ID获取邮箱详情
func (e *Email) GetEmailByEmailId(emailId int64) (email *entity.EmailEntity, err errors.BizError) {
	email, err = e.daoEmail.GetEmailByEmailId(emailId)
	if err != nil {
		return email, err
	}
	return email, nil
}

// DeleteEmail 删除邮箱
func (e *Email) DeleteEmail(emailId int64) errors.BizError {
	// 查找邮箱是否存在
	updateEmail, err := e.daoEmail.GetEmailByEmailId(emailId)
	if err != nil {
		return err
	}
	if updateEmail == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "邮箱id %d 不存在", emailId)
	}
	// 删除邮箱
	err = e.daoEmail.DeleteEmail(emailId)
	if err != nil {
		return err
	}
	return nil
}

// GetEmailsByLimit 分页获取邮箱列表
func (e *Email) GetEmailsByLimit(pageSize int, pageNum int, keywords *entity.EmailKeywords) (emails []*entity.EmailEntity, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, EmailListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	offset := (pageNum - 1) * pageSize
	// 无搜索词
	if keywords == nil {
		return e.daoEmail.GetEmailsByLimit(pageSize, offset)
	}
	return e.daoEmail.GetEmailsByKeywordsAndLimit(pageSize, offset, keywords)
}

// GetPageInfoLimit 获取分页信息
func (e *Email) GetPageInfoLimit(pageSize int, pageNum int, keywords *entity.EmailKeywords) (pagination *entity.PageInfo, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, EmailListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	pageInfo := new(entity.PageInfo)
	var totalNum int64
	// 无搜索词
	if keywords == nil {
		totalNum, err = e.daoEmail.CountEmails()
	} else {
		totalNum, err = e.daoEmail.CountEmailsByKeywords(keywords)
	}
	if err != nil {
		logger.WithContext(e.ctx).Warnf("[service.Email] GetPageInfoLimit pageSize=%d pageNum=%d err=%s",
			pageSize, pageNum, err.Error())
		return pageInfo, err
	}
	return entity.GetPageInfo(totalNum, pageSize, pageNum), nil
}

// FormatEmailList 格式化邮箱列表
func (e *Email) FormatEmailList(emails []*entity.EmailEntity) (emailList []*entity.EmailListItem, err errors.BizError) {
	if len(emails) == 0 {
		return emailList, nil
	}
	for _, email := range emails {
		action := e.GetListItemAction(email)
		var emailListItem = &entity.EmailListItem{
			EmailEntity: email,
			Action:      action,
		}
		emailList = append(emailList, emailListItem)
	}
	return emailList, nil
}

// GetListItemAction 获取列表 action 权限
func (e *Email) GetListItemAction(emailItem *entity.EmailEntity) *entity.EmailListAction {
	indentifys := global.ContextValueLoginIdentifys(e.ctx)
	action := &entity.EmailListAction{
		IsEdit:   indentifys[global.PrivilegeIndentifyEmailEdit],
		IsDelete: indentifys[global.PrivilegeIndentifyEmailDelete],
	}
	return action
}

// GetUsedEmail 获取正在使用的邮箱
func (e *Email) GetUsedEmail() (email *entity.EmailEntity, err errors.BizError) {
	return e.daoEmail.GetUsedEmail()
}

// SetEmailUsed 设置邮箱为使用中
func (e *Email) SetEmailUsed(emailId int64) errors.BizError {
	// 查找邮箱是否存在
	emailEntity, err := e.daoEmail.GetEmailByEmailId(emailId)
	if err != nil {
		return err
	}
	if emailEntity == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "邮箱id %d 不存在", emailId)
	}
	// 先设置所有邮箱为未使用
	err = e.daoEmail.UpdateAllEmailUnused()
	if err != nil {
		return err
	}
	// 设置当前邮箱为使用中
	emailEntity.IsUsed = entity.EmailUsedTrue
	err = e.daoEmail.Update(*emailEntity)
	if err != nil {
		return err
	}
	return nil
}
