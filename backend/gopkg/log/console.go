package log

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const WriterConsole = "console"

// ConsoleWriter console writer
type ConsoleWriter struct {
	cfg      *OutputConfig
	zapCore  zapcore.Core
	zapLevel zap.AtomicLevel
}

// NewConsoleWriter 输出到 console
func NewConsoleWriter() Writer {
	return &ConsoleWriter{}
}

// Name 日志输出名
func (cw *ConsoleWriter) Name() string {
	return WriterConsole
}

// Init 初始化 console writer 配置
func (cw *ConsoleWriter) Init(name string, cfg *OutputConfig) error {
	if cfg == nil {
		return fmt.Errorf("console writer:%+v outputConfig empty", name)
	}
	cw.cfg = cfg
	cw.makeConsoleCore()
	return nil
}

// GetZapCore 获取 zap core
func (cw *ConsoleWriter) GetZapCore() zapcore.Core {
	return cw.zapCore
}

// GetZapLevel 获取对应的 zap level
func (cw *ConsoleWriter) GetZapLevel() zap.AtomicLevel {
	return cw.zapLevel
}

func (cw *ConsoleWriter) makeConsoleCore() {
	cw.zapLevel = zap.NewAtomicLevelAt(zapCoreLevels[cw.cfg.Level])
	encoder := makeEncoder(cw.cfg)
	cw.zapCore = zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), cw.zapLevel)
}
