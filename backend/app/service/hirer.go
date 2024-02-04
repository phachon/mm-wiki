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
	HirerListDefaultPageSize = 20 // 租客列表默认一页 20 条
)

// Hirer 租客业务逻辑
type Hirer struct {
	ctx      context.Context
	daoHirer *dao.Hirer
}

// NewHirer 创建租客业务逻辑对象
func NewHirer(ctx context.Context) *Hirer {
	return &Hirer{
		ctx:      ctx,
		daoHirer: dao.NewHirer(ctx),
	}
}

// Create 创建租客
func (a *Hirer) Create(hirerEntity *entity.HirerEntity) errors.BizError {
	// 插入一条记录
	err := a.daoHirer.Insert(hirerEntity)
	if err != nil {
		return err
	}
	return nil
}

// Update 修改租客
func (a *Hirer) Update(hirerEntity entity.HirerEntity) errors.BizError {
	// 查找租客是否存在
	updateHirer, err := a.daoHirer.GetHirerByHirerId(hirerEntity.HirerId)
	if err != nil {
		return err
	}
	if updateHirer == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "租客不存在")
	}
	// 更新字段
	err = a.daoHirer.Update(hirerEntity)
	if err != nil {
		return err
	}
	return nil
}

// GetHirerByHirerId 根据租客ID获取租客详情
func (a *Hirer) GetHirerByHirerId(hirerID int64) (hirer *entity.HirerEntity, err errors.BizError) {
	hirer, err = a.daoHirer.GetHirerByHirerId(hirerID)
	if err != nil {
		return hirer, err
	}
	return hirer, nil
}

// GetHirersByHirerIds 根据租客ID获取租客详情
func (a *Hirer) GetHirersByHirerIds(hirerIDs []int64) (hirers []*entity.HirerEntity, err errors.BizError) {
	if len(hirerIDs) == 0 {
		return hirers, nil
	}
	hirers, err = a.daoHirer.GetHirersByHirerIds(hirerIDs)
	if err != nil {
		return hirers, err
	}
	return hirers, nil
}

// GetHirersByHirerIds 根据租客ID获取租客详情
func (a *Hirer) GetHirersMapByHirerIds(hirerIDs []int64) (hirersMap map[int64]*entity.HirerEntity, err errors.BizError) {
	hirersMap = make(map[int64]*entity.HirerEntity)
	if len(hirerIDs) == 0 {
		return hirersMap, nil
	}
	hirers, err := a.daoHirer.GetHirersByHirerIds(hirerIDs)
	if err != nil {
		return hirersMap, err
	}
	for _, hirer := range hirers {
		hirersMap[hirer.HirerId] = hirer
	}
	return hirersMap, nil
}

// DeleteByHirerId 根据租客ID删除
func (a *Hirer) DeleteByHirerId(hirerId int64) errors.BizError {
	// 查找租客是否存在
	updateHirer, err := a.daoHirer.GetHirerByHirerId(hirerId)
	if err != nil {
		return err
	}
	if updateHirer == nil {
		return errors.Errorf(errors.BusinessRecordNotExistError, "租客id %d 不存在", hirerId)
	}
	// 更新状态字段
	err = a.daoHirer.UpdateStatus(hirerId, entity.HirerStatusDelete)
	if err != nil {
		return err
	}
	return nil
}

// GetHirersByLimit 分页获取租客列表
func (a *Hirer) GetHirersByLimit(pageSize int, pageNum int, keywords *entity.HirerKeywords) (hirers []*entity.HirerEntity, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, HirerListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	offset := (pageNum - 1) * pageSize
	// 无搜索词
	if keywords == nil {
		return a.daoHirer.GetHirersByLimit(pageSize, offset)
	}
	return a.daoHirer.GetHirersByKeywordAndLimit(pageSize, offset, keywords)
}

// GetAllHirers 获取所有的租客
func (a *Hirer) GetAllHirers() (hirers []*entity.HirerEntity, err errors.BizError) {
	return a.daoHirer.GetAllHirers()
}

// GetPageInfoLimit 获取分页信息
func (a *Hirer) GetPageInfoLimit(pageSize int, pageNum int, keywords *entity.HirerKeywords) (pagination *entity.PageInfo, err errors.BizError) {
	pageSize = utils.VerifyUint(pageSize, HirerListDefaultPageSize)
	pageNum = utils.VerifyUint(pageNum, 1)
	pageInfo := new(entity.PageInfo)
	var totalNum int64

	// 无搜索词
	if keywords == nil {
		totalNum, err = a.daoHirer.CountHirers()
	} else {
		totalNum, err = a.daoHirer.CountHirersByKeywords(keywords)
	}
	if err != nil {
		logger.WithContext(a.ctx).Warnf("[service.Hirer] GetPageInfoLimit pageSize=%d pageNum=%d err=%s",
			pageSize, pageNum, err.Error())
		return pageInfo, err
	}
	return entity.GetPageInfo(totalNum, pageSize, pageNum), nil
}
