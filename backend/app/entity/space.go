package entity

import "github.com/phachon/mm-wiki/utils"

const (
	SpaceStatusDefault = 0  // 空间状态 0 正常
	SpaceStatusDelete  = -1 // 空间状态 -1 删除

	SpaceVisitLevelDefaultPublic  = 0 // 空间访问级别 0 公开
	SpaceVisitLevelDefaultPrivate = 1 // 空间访问级别 1 私有

	SpaceIsShareDefault = 0 // 空间是否允许分享 0 否
	SpaceIsShareYes     = 1 // 空间是否允许分享 1 是

	SpaceIsExportDefault = 0 // 空间是否允许导出 0 否
	SpaceIsExportYes     = 1 // 空间是否允许导出 1 是

	SpaceTypeTeam            = 0 // 空间类型 0 团队
	SpaceTypeDefaultPersonal = 1 // 空间类型 1 个人
)

// SpaceEntity space 空间表结构
type SpaceEntity struct {
	SpaceId          int64          `json:"space_id" gorm:"primary_key"` // 空间ID
	SpaceKey         string         `json:"space_key"`                   // 空间key
	Name             string         `json:"name"`                        // 空间名
	Description      string         `json:"description"`                 // 空间描述
	SpaceType        int            `json:"space_type"`                  // 空间类型 0 团队 1 个人
	VisitLevel       int            `json:"visit_level"`                 // 访问级别
	IsShare          int            `json:"is_share"`                    // 文档是否允许分享 0 否 1 是
	IsExport         int            `json:"is_export"`                   // 文档是否允许导出 0 否 1 是
	CreatorAccountId int64          `json:"creator_account_id"`          // 创建者账户ID
	CreatorName      string         `json:"creator_name"`                // 创建者名称
	Status           int            `json:"status"`                      // 空间状态
	CreateTime       utils.JsonTime `json:"create_time"`                 // 创建时间
	UpdateTime       utils.JsonTime `json:"update_time"`                 // 更新时间
}

// SpaceListItem 空间列表结构
type SpaceListItem struct {
	*SpaceEntity
	Action *SpaceListAction `json:"action"` // 数据操作
}

// SpaceListAction 空间列表操作权限
type SpaceListAction struct {
	IsEdit   int `json:"is_edit,omitempty"`   // 空间修改
	IsDelete int `json:"is_delete,omitempty"` // 空间删除
}

// SpaceKeywords 空间搜索词
type SpaceKeywords struct {
	SpaceName string `json:"space_name"` // 空间名称
	SpaceType *int   `json:"space_type"` // 空间类型
}
