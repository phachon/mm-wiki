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
	ContactListDefaultPageSize = 20 // 联系人列表默认一页 20 条
)

// Contact 联系人业务逻辑
type Contact struct {
	ctx        context.Context
	daoContact *dao.Contact
}

// NewContact 创建联系人业务逻辑对象
func NewContact(ctx context.Context) *Contact {
	return &Contact{
		ctx:        ctx,
		daoContact: dao.NewContact(ctx),
	}
}

// Create 创建联系人
func (c *Contact) Create(contactEntity *entity.ContactEntity) errors.BizError {
	// 插入一条记录
	err := c.daoContact.Insert(contactEntity)
	if err != nil {
		return err
	}
	return nil
}

// Update 修改联系人
func (c *Contact) Update(contactEntity entity.ContactEntity) errors.BizError {
	// 查找联系人是否存在
	updateContact, err := c.daoContact.GetContactByContactId(contactEntity.ContactId)
	if err != nil {
		return err
	}
	if updateContact == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "联系人不存在")
	}
	// 更新字段
	err = c.daoContact.Update(contactEntity)
	if err != nil {
		return err
	}
	return nil
}

// GetContactByContactId 根据联系人ID获取联系人详情
func (c *Contact) GetContactByContactId(contactId int64) (contact *entity.ContactEntity, err errors.BizError) {
	contact, err = c.daoContact.GetContactByContactId(contactId)
	if err != nil {
		return contact, err
	}
	return contact, nil
}

// DeleteContact 删除联系人
func (c *Contact) DeleteContact(contactId int64) errors.BizError {
	// 查找联系人是否存在
	updateContact, err := c.daoContact.GetContactByContactId(contactId)
	if err != nil {
		return err
	}
	if updateContact == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "联系人id %d 不存在", contactId)
	}
	// 删除联系人
	err = c.daoContact.DeleteContact(contactId)
	if err != nil {
		return err
	}
	return nil
}

// GetContactsByLimit 分页获取联系人列表
func (c *Contact) GetContactsByLimit(pageSize int, pageNum int, keywords *entity.ContactKeywords) (contacts []*entity.ContactEntity, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, ContactListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	offset := (pageNum - 1) * pageSize
	// 无搜索词
	if keywords == nil {
		return c.daoContact.GetContactsByLimit(pageSize, offset)
	}
	return c.daoContact.GetContactsByKeywordsAndLimit(pageSize, offset, keywords)
}

// GetPageInfoLimit 获取分页信息
func (c *Contact) GetPageInfoLimit(pageSize int, pageNum int, keywords *entity.ContactKeywords) (pagination *entity.PageInfo, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, ContactListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	pageInfo := new(entity.PageInfo)
	var totalNum int64
	// 无搜索词
	if keywords == nil {
		totalNum, err = c.daoContact.CountContacts()
	} else {
		totalNum, err = c.daoContact.CountContactsByKeywords(keywords)
	}
	if err != nil {
		logger.WithContext(c.ctx).Warnf("[service.Contact] GetPageInfoLimit pageSize=%d pageNum=%d err=%s",
			pageSize, pageNum, err.Error())
		return pageInfo, err
	}
	return entity.GetPageInfo(totalNum, pageSize, pageNum), nil
}

// FormatContactList 格式化联系人列表
func (c *Contact) FormatContactList(contacts []*entity.ContactEntity) (contactList []*entity.ContactListItem, err errors.BizError) {
	if len(contacts) == 0 {
		return contactList, nil
	}
	for _, contact := range contacts {
		action := c.GetListItemAction(contact)
		var contactListItem = &entity.ContactListItem{
			ContactEntity: contact,
			Action:        action,
		}
		contactList = append(contactList, contactListItem)
	}
	return contactList, nil
}

// GetListItemAction 获取列表 action 权限
func (c *Contact) GetListItemAction(contactItem *entity.ContactEntity) *entity.ContactListAction {
	indentifys := global.ContextValueLoginIdentifys(c.ctx)
	action := &entity.ContactListAction{
		IsEdit:   indentifys[global.PrivilegeIndentifyContactEdit],
		IsDelete: indentifys[global.PrivilegeIndentifyContactDelete],
	}
	return action
}
