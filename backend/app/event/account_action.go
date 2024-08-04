package event

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/app/service"
	"github.com/phachon/mm-wiki/global"
	klog "github.com/phachon/mm-wiki/gopkg/log"
)

// DispatchInfo 分发信息
type DispatchInfo struct {
	ActionType string                 `json:"action_type"` // 分发类型
	Message    string                 `json:"message"`     // 事件信息
	Fields     map[string]interface{} `json:"fields"`      // 扩展字段
}

// AccountActionEvent 账号操作事件
type AccountActionEvent struct {
	ctx        context.Context
	logService *service.Log
}

// NewAccountActionEvent 创建账号操作事件
func NewAccountActionEvent(ctx context.Context) *AccountActionEvent {
	return &AccountActionEvent{
		ctx:        ctx,
		logService: service.NewLog(ctx),
	}
}

// Dispatch 分发事件
func (ae *AccountActionEvent) Dispatch(actionType string, fields map[string]interface{}) {
	message := ActionTypeMessage[actionType]
	if len(fields) > 0 {
		fieldsJson, _ := json.Marshal(fields)
		message = fmt.Sprintf("%s %s", message, string(fieldsJson))
	}
	ae.consumerSystemLogDB(message)
}

// consumerSystemLogDB 消费事件到 system_log 表
func (ae *AccountActionEvent) consumerSystemLogDB(message string) {
	logEntity := new(entity.LogEntity)
	logEntity.Message = message
	logEntity.Level = int(klog.LevelInfo)
	ae.setLogEntityExtra(ae.ctx, logEntity)
	ae.logService.Create(logEntity)
}

func (ae *AccountActionEvent) setLogEntityExtra(ctx context.Context, logEntity *entity.LogEntity) {
	if ctx == nil || logEntity == nil {
		return
	}
	ginCtx, ok := ctx.(*gin.Context)
	if !ok || ginCtx == nil {
		return
	}
	logEntity.Uri = ginCtx.Request.RequestURI
	logEntity.Get = ginCtx.Request.URL.Query().Encode()
	logEntity.Post = ginCtx.Request.PostForm.Encode()
	logEntity.Ip = ginCtx.ClientIP()
	logEntity.AccountId = global.ContextValueLoginAccountID(ctx)
	logEntity.AccountName = global.ContextValueLoginAccountName(ctx)
}
