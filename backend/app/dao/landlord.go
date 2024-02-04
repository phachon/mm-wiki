package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
	"github.com/phachon/mm-wiki/utils"
	"gorm.io/gorm"
)

const (
	// TableNameLandlord 房屋业主表
	TableNameLandlord = "hms_landlord"
	// LandlordPrimaryKey 业主表主键ID
	LandlordPrimaryKey = "landlord_id"
)

// Landlord 房屋业主表数据
type Landlord struct {
	ctx context.Context
}

// NewLandlord 创建房屋业主表数据对象
func NewLandlord(ctx context.Context) *Landlord {
	return &Landlord{
		ctx: ctx,
	}
}

// Insert 创建业主插入一条业主记录
func (l *Landlord) Insert(landlordEntity *entity.LandlordEntity) errors.BizError {
	landlordEntity.CreateTime = utils.NewJsonTime(time.Now())
	landlordEntity.UpdateTime = utils.NewJsonTime(time.Now())
	landlordEntity.Status = entity.LandlordStatusDefault
	db := GetDB(dbNameKms).WithContext(l.ctx).
		Table(TableNameLandlord).
		Save(landlordEntity)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlInsertErr, db.Error.Error())
	}
	return nil
}

// GetLandlordByName 根据业主昵称查找业主
func (l *Landlord) GetLandlordByNickName(nickName string) (landlord *entity.LandlordEntity, err errors.BizError) {
	landlord = &entity.LandlordEntity{}
	db := GetDB(dbNameKms).WithContext(l.ctx).Table(TableNameLandlord).
		Where(map[string]interface{}{
			"status":    entity.LandlordStatusDefault,
			"nick_name": nickName,
		}).
		First(&landlord)
	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return landlord, nil
}

// GetLandlordByLandlordId 根据业主ID获取业主信息
func (l *Landlord) GetLandlordByLandlordId(landlordId int64) (landlord *entity.LandlordEntity, err errors.BizError) {

	landlord = &entity.LandlordEntity{}
	db := GetDB(dbNameKms).WithContext(l.ctx).Table(TableNameLandlord).
		Where(map[string]interface{}{
			LandlordPrimaryKey: landlordId,
		}).
		First(&landlord)

	if db.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if db.Error != nil {
		return nil, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return landlord, nil
}

// GetLandlordsByLandlordIds 根据多个业主ID批量获取业主ID
func (l *Landlord) GetLandlordsByLandlordIds(landlordIds []int64) (landlords []*entity.LandlordEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(l.ctx).Table(TableNameLandlord).
		Where("landlord_id IN (?)", landlordIds).
		Find(&landlords)
	if db.Error != nil {
		return landlords, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return landlords, nil
}

// CountByNickName 根据业主昵称获取业主数
func (l *Landlord) CountByNickName(nickName string) (count int64, err error) {
	if nickName == "" {
		return 0, nil
	}
	db := GetDB(dbNameKms).WithContext(l.ctx).Table(TableNameLandlord).
		Where(map[string]interface{}{
			"status":    entity.LandlordStatusDefault,
			"nick_name": nickName,
		}).Count(&count)
	if db.Error != nil {
		return 0, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// CheckNameExists 检查业主名是否存在
func (l *Landlord) CheckNameExists(name string) (bool, error) {
	count, err := l.CountByNickName(name)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// HasSameNickName 业主ID和业主昵称是否存在
func (l *Landlord) HasSameNickName(landlordId int64, nickName string) (has bool, err errors.BizError) {
	var count int64
	db := GetDB(dbNameKms).WithContext(l.ctx).Table(TableNameLandlord).
		Where("nick_name = ?", nickName).
		Where("landlord_id <> ?", landlordId).
		Count(&count)
	if db.Error != nil {
		return false, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count > 0, nil
}

// Update 更新业主，只会更新如下字段
func (l *Landlord) Update(landlord entity.LandlordEntity) errors.BizError {
	landlord.UpdateTime = utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(l.ctx).Table(TableNameLandlord).
		Select("NickName", "GivenName", "Sex", "Email", "IdCardNumber", "Mobile", "Address", "UpdateTime").
		Where(map[string]interface{}{
			LandlordPrimaryKey: landlord.LandlordId,
		}).
		Updates(landlord)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// UpdateStatus 更新业主状态
func (l *Landlord) UpdateStatus(landlordId int64, status int) errors.BizError {
	updateTime := utils.NewJsonTime(time.Now())
	db := GetDB(dbNameKms).WithContext(l.ctx).Table(TableNameLandlord).
		Where(map[string]interface{}{
			LandlordPrimaryKey: landlordId,
		}).
		Update("status", status).
		Update("update_time", updateTime)
	if db.Error != nil {
		return errors.Errorf(errors.DalMysqlUpdateErr, db.Error.Error())
	}
	return nil
}

// GetAllLandlords 获取所有的业主
func (l *Landlord) GetAllLandlords() (landlord []*entity.LandlordEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(l.ctx).
		Table(TableNameLandlord).
		Where("status = ?", entity.LandlordStatusDefault).
		Find(&landlord)
	if db.Error != nil {
		return landlord, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return landlord, nil
}

// GetLandlordsByKeywords 根据业主名模糊匹配业主
func (l *Landlord) GetLandlordsByKeywords(keywords *entity.LandlordKeywords) (landlords []*entity.LandlordEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(l.ctx).
		Table(TableNameLandlord).
		Where("status = ?", entity.LandlordStatusDefault)
	if keywords.GivenName != "" {
		db = db.Where("given_name LIKE ?", "%"+keywords.GivenName+"%")
	}
	db.Find(&landlords)
	if db.Error != nil {
		return landlords, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return landlords, nil
}

// GetLandlordsByKeywordAndLimit 根据关键字分页获取业主
func (l *Landlord) GetLandlordsByKeywordAndLimit(limit int, offset int, keywords *entity.LandlordKeywords) (landlords []*entity.LandlordEntity, err errors.BizError) {
	if keywords == nil {
		return l.GetLandlordsByLimit(limit, offset)
	}
	db := GetDB(dbNameKms).WithContext(l.ctx).
		Table(TableNameLandlord).
		Where("status = ?", entity.LandlordStatusDefault)
	if keywords.GivenName != "" {
		db = db.Where("given_name LIKE ?", "%"+keywords.GivenName+"%")
	}
	db = db.Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", LandlordPrimaryKey)).
		Find(&landlords)
	if db.Error != nil {
		return landlords, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return landlords, nil
}

// GetLandlordsByLimit 分页获取业主列表
func (l *Landlord) GetLandlordsByLimit(limit int, offset int) (landlords []*entity.LandlordEntity, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(l.ctx).
		Table(TableNameLandlord).
		Where("status = ?", entity.LandlordStatusDefault).
		Limit(limit).
		Offset(offset).
		Order(fmt.Sprintf("%s DESC", LandlordPrimaryKey)).
		Find(&landlords)
	if db.Error != nil {
		return landlords, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return landlords, nil
}

// CountLandlords 业主总数
func (l *Landlord) CountLandlords() (count int64, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(l.ctx).
		Table(TableNameLandlord).
		Where("status = ?", entity.LandlordStatusDefault).
		Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}

// CountLandlords 业主总数
func (l *Landlord) CountLandlordsByKeywords(keywords *entity.LandlordKeywords) (count int64, err errors.BizError) {
	db := GetDB(dbNameKms).WithContext(l.ctx).
		Table(TableNameLandlord).
		Where("status = ?", entity.LandlordStatusDefault)
	if keywords.GivenName != "" {
		db = db.Where("given_name LIKE ?", "%"+keywords.GivenName+"%")
	}
	db.Count(&count)
	if db.Error != nil {
		return count, errors.Errorf(errors.DalMysqlSelectErr, db.Error.Error())
	}
	return count, nil
}
