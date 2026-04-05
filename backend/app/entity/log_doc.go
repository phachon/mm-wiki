package entity

const (
	LogDocActionCreate = 1 // 动作 1 创建
	LogDocActionModify = 2 // 动作 2 修改
	LogDocActionDelete = 3 // 动作 3 删除
)

// LogDocEntity 文档操作日志数据结构
type LogDocEntity struct {
	LogDocId   int64  `json:"log_doc_id" gorm:"primary_key"` // 文档日志ID
	DocId      string `json:"doc_id"`                         // 文档id
	SpaceId    int64  `json:"space_id"`                       // 空间id
	AccountId  int64  `json:"account_id"`                     // 账号id
	Action     int    `json:"action"`                         // 动作 1 创建 2 修改 3 删除
	Comment    string `json:"comment"`                        // 备注信息
	CreateTime int64  `json:"create_time"`                    // 创建时间
}

// LogDocListItem 文档日志列表结构
type LogDocListItem struct {
	*LogDocEntity
	AccountName string `json:"account_name"` // 账号名
	DocName     string `json:"doc_name"`     // 文档名称
}

// LogDocKeywords 文档日志搜索关键字
type LogDocKeywords struct {
	DocId   string `json:"doc_id"`   // 文档ID
	SpaceId int64  `json:"space_id"` // 空间ID
}
