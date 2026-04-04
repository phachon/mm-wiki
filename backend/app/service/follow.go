package service

import (
	"context"

	"github.com/phachon/mm-wiki/app/dao"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
)

// Follow 关注业务逻辑
type Follow struct {
	ctx       context.Context
	daoFollow *dao.Follow
}

// NewFollow 创建关注业务逻辑对象
func NewFollow(ctx context.Context) *Follow {
	return &Follow{
		ctx:       ctx,
		daoFollow: dao.NewFollow(ctx),
	}
}

// Create 创建关注
func (f *Follow) Create(accountId int64, followType int, objectId string) errors.BizError {
	if accountId <= 0 || objectId == "" {
		return errors.Errorf(errors.ClientReqParamEmpty, "参数不能为空")
	}
	// 检查是否已关注
	existing, err := f.daoFollow.GetFollowByAccountIdAndTypeAndObjectId(accountId, followType, objectId)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.Errorf(errors.BusinessRecordExistError, "已经关注过")
	}
	followEntity := &entity.FollowEntity{
		AccountId:  accountId,
		FollowType: followType,
		ObjectId:   objectId,
	}
	return f.daoFollow.Insert(followEntity)
}

// Cancel 取消关注
func (f *Follow) Cancel(accountId int64, followType int, objectId string) errors.BizError {
	if accountId <= 0 || objectId == "" {
		return nil
	}
	return f.daoFollow.DeleteByAccountIdTypeAndObjectId(accountId, followType, objectId)
}

// GetFollowStatus 获取关注状态
func (f *Follow) GetFollowStatus(accountId int64, followType int, objectId string) (bool, errors.BizError) {
	if accountId <= 0 || objectId == "" {
		return false, nil
	}
	follow, err := f.daoFollow.GetFollowByAccountIdAndTypeAndObjectId(accountId, followType, objectId)
	if err != nil {
		return false, err
	}
	return follow != nil, nil
}
