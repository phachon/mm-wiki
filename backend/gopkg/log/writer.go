package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Writer 日志输出定义
type Writer interface {
	// Writer 输出名称
	Name() string
	// Init 初始化配置
	Init(name string, cfg *OutputConfig) error
	// GetZapCore 获取 zap core
	GetZapCore() zapcore.Core
	// GetZapLevel 获取 zap level
	GetZapLevel() zap.AtomicLevel
}