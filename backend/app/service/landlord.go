package service

import (
	"context"

	"github.com/phachon/mm-wiki/app/dao"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/logger"
	"github.com/phachon/mm-wiki/utils"
)

const (
	LandlordListDefaultPageSize = 20 // 业主列表默认一页 20 条
)

// Landlord 业主业务逻辑
type Landlord struct {
	ctx         context.Context
	daoLandlord *dao.Landlord
}

// NewLandlord 创建业主业务逻辑对象
func NewLandlord(ctx context.Context) *Landlord {
	return &Landlord{
		ctx:         ctx,
		daoLandlord: dao.NewLandlord(ctx),
	}
}

// Create 创建业主
func (a *Landlord) Create(landlordEntity *entity.LandlordEntity) errors.BizError {
	// 查找昵称是否存在
	landlord, err := a.daoLandlord.GetLandlordByNickName(landlordEntity.NickName)
	if err != nil {
		return err
	}
	if landlord != nil {
		return errors.Errorf(errors.BusinessRecordExistError, "业主昵称 %s 已经存在", landlordEntity.NickName)
	}
	// 插入一条记录
	err = a.daoLandlord.Insert(landlordEntity)
	if err != nil {
		return err
	}
	return nil
}

// Update 修改业主
func (a *Landlord) Update(landlordEntity entity.LandlordEntity) errors.BizError {
	// 查找业主是否存在
	updateLandlord, err := a.daoLandlord.GetLandlordByLandlordId(landlordEntity.LandlordId)
	if err != nil {
		return err
	}
	if updateLandlord == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "业主不存在")
	}
	if updateLandlord.Status == entity.LandlordStatusDelete {
		return errors.Errorf(errors.BusinessForbiddenError, "删除的业主无法修改")
	}
	if landlordEntity.NickName != "" {
		// 查找 name 是否存在
		hasName, err := a.daoLandlord.HasSameNickName(landlordEntity.LandlordId, landlordEntity.NickName)
		if err != nil {
			logger.WithContext(a.ctx).Errorf("[service.Landlord] Update HasSameName landlordId=%d name=%s err=%s",
				landlordEntity.LandlordId, landlordEntity.NickName, err.Error())
		}
		if hasName {
			return errors.Errorf(errors.BusinessRecordExistError, "业主昵称 %s 已经存在", landlordEntity.NickName)
		}
	}
	// 更新字段
	err = a.daoLandlord.Update(landlordEntity)
	if err != nil {
		return err
	}
	return nil
}

// GetLandlordByLandlordId 根据业主ID获取业主详情
func (a *Landlord) GetLandlordByLandlordId(landlordID int64) (landlord *entity.LandlordEntity, err errors.BizError) {
	landlord, err = a.daoLandlord.GetLandlordByLandlordId(landlordID)
	if err != nil {
		return landlord, err
	}
	return landlord, nil
}

// DeleteByLandlordId 根据业主ID删除
func (a *Landlord) DeleteByLandlordId(landlordId int64) errors.BizError {
	// 查找业主是否存在
	updateLandlord, err := a.daoLandlord.GetLandlordByLandlordId(landlordId)
	if err != nil {
		return err
	}
	if updateLandlord == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "业主id %d 不存在", landlordId)
	}
	// 更新状态字段
	err = a.daoLandlord.UpdateStatus(landlordId, entity.LandlordStatusDelete)
	if err != nil {
		return err
	}
	return nil
}

// GetLandlordsByLimit 分页获取业主列表
func (a *Landlord) GetLandlordsByLimit(pageSize int, pageNum int, keywords *entity.LandlordKeywords) (landlords []*entity.LandlordEntity, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, LandlordListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	offset := (pageNum - 1) * pageSize
	// 无搜索词
	if keywords == nil {
		return a.daoLandlord.GetLandlordsByLimit(pageSize, offset)
	}
	return a.daoLandlord.GetLandlordsByKeywordAndLimit(pageSize, offset, keywords)
}

// GetAllLandlords 获取所有的业主
func (a *Landlord) GetAllLandlords() (landlords []*entity.LandlordEntity, err errors.BizError) {
	return a.daoLandlord.GetAllLandlords()
}

// GetPageInfoLimit 获取分页信息
func (a *Landlord) GetPageInfoLimit(pageSize int, pageNum int, keywords *entity.LandlordKeywords) (pagination *entity.PageInfo, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, LandlordListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	pageInfo := new(entity.PageInfo)
	var totalNum int64

	// 无搜索词
	if keywords == nil {
		totalNum, err = a.daoLandlord.CountLandlords()
	} else {
		totalNum, err = a.daoLandlord.CountLandlordsByKeywords(keywords)
	}
	if err != nil {
		logger.WithContext(a.ctx).Warnf("[service.Landlord] GetPageInfoLimit pageSize=%d pageNum=%d err=%s",
			pageSize, pageNum, err.Error())
		return pageInfo, err
	}
	return entity.GetPageInfo(totalNum, pageSize, pageNum), nil
}
