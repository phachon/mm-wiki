package log

import (
	"fmt"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const WriterFile = "file"

// FileWriter console writer
type FileWriter struct {
	cfg      *OutputConfig
	zapCore  zapcore.Core
	zapLevel zap.AtomicLevel
}

// NewFileWriter 输出到 file
func NewFileWriter() Writer {
	return &FileWriter{}
}

// Name 日志输出名
func (cf *FileWriter) Name() string {
	return WriterFile
}

// Init 初始化 console writer 配置
func (cf *FileWriter) Init(name string, cfg *OutputConfig) error {
	if cfg == nil {
		return fmt.Errorf("console writer:%+v outputConfig empty", name)
	}
	if cfg.WriteConfig.LogPath != "" {
		cfg.WriteConfig.Filename = filepath.Join(cfg.WriteConfig.LogPath, cfg.WriteConfig.Filename)
	}
	cf.cfg = cfg
	cf.makeFileCore()
	return nil
}

// GetZapCore 获取 zap core
func (cf *FileWriter) GetZapCore() zapcore.Core {
	return cf.zapCore
}

// GetZapLevel 获取对应的 zap level
func (cf *FileWriter) GetZapLevel() zap.AtomicLevel {
	return cf.zapLevel
}

func (cf *FileWriter) makeFileCore() {
	cf.zapLevel = zap.NewAtomicLevelAt(zapCoreLevels[cf.cfg.Level])
	encoder := makeEncoder(cf.cfg)

	lumberJackLogger := &lumberjack.Logger{
		Filename:   cf.cfg.WriteConfig.Filename,   // 日志文件位置
		MaxSize:    cf.cfg.WriteConfig.MaxSize,    // 进行切割之前,日志文件的最大大小(MB为单位)
		MaxBackups: cf.cfg.WriteConfig.MaxBackups, // 保留旧文件的最大个数
		MaxAge:     cf.cfg.WriteConfig.MaxAge,     // 保留旧文件的最大天数
		Compress:   cf.cfg.WriteConfig.Compress,   // 是否压缩/归档旧文件
	}
	syncFileWriter := zapcore.AddSync(lumberJackLogger)
	cf.zapCore = zapcore.NewCore(encoder, zapcore.Lock(syncFileWriter), cf.zapLevel)
}
