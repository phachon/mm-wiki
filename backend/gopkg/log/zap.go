package log

import (
	"fmt"
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Levels zapcore level
var zapCoreLevels = map[string]zapcore.Level{
	"":      zapcore.DebugLevel,
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

var levelToZapLevel = map[Level]zapcore.Level{
	LevelTrace: zapcore.DebugLevel,
	LevelDebug: zapcore.DebugLevel,
	LevelInfo:  zapcore.InfoLevel,
	LevelWarn:  zapcore.WarnLevel,
	LevelError: zapcore.ErrorLevel,
	LevelFatal: zapcore.FatalLevel,
}

var zapLevelToLevel = map[zapcore.Level]Level{
	zapcore.DebugLevel: LevelDebug,
	zapcore.InfoLevel:  LevelInfo,
	zapcore.WarnLevel:  LevelWarn,
	zapcore.ErrorLevel: LevelError,
	zapcore.FatalLevel: LevelFatal,
}

// zapLog zap log 实现的 log
type zapLog struct {
	levels []zap.AtomicLevel
	logger *zap.Logger
}

// NewZapLog 创建 zap log
func NewZapLog(cf Config) Logger {
	return NewZapLogWithCallerSkip(cf, 2)
}

// NewZapLogWithCallerSkip 创建一个 zap log 带函数调用深度
func NewZapLogWithCallerSkip(outputs Config, callerSkip int) Logger {
	cores := make([]zapcore.Core, 0, len(outputs))
	levels := make([]zap.AtomicLevel, 0, len(outputs))
	for _, output := range outputs {
		writer, ok := logWriters[output.Writer]
		if !ok {
			fmt.Printf("zapLog writer:%+v no support!\n", output.Writer)
			return nil
		}
		// 初始化配置
		if err := writer.Init(output.Writer, &output); err != nil {
			fmt.Printf("zapLog writer:%+v Init err:%v!\n", output.Writer, err)
			return nil
		}
		cores = append(cores, writer.GetZapCore())
		levels = append(levels, writer.GetZapLevel())
	}
	logger := zap.New(
		zapcore.NewTee(cores...),
		zap.AddCallerSkip(callerSkip),
		zap.AddCaller(),
	)
	return &zapLog{
		levels: levels,
		logger: logger,
	}
}

// WithFields 自定义一些字段
func (l *zapLog) WithFields(fields ...string) Logger {
	zapFields := make([]zap.Field, len(fields)/2)
	for index := range zapFields {
		zapFields[index] = zap.String(fields[2*index], fields[2*index+1])
	}
	return &ZapLogWrapper{l: &zapLog{logger: l.logger.With(zapFields...)}}
}

// Trace _
func (l *zapLog) Trace(args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.DebugLevel) {
		l.logger.Debug(fmt.Sprint(args...))
	}
}

// Tracef _
func (l *zapLog) Tracef(format string, args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.DebugLevel) {
		l.logger.Debug(fmt.Sprintf(format, args...))
	}
}

// Debug _
func (l *zapLog) Debug(args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.DebugLevel) {
		l.logger.Debug(fmt.Sprint(args...))
	}
}

// Debugf _
func (l *zapLog) Debugf(format string, args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.DebugLevel) {
		l.logger.Debug(fmt.Sprintf(format, args...))
	}
}

// Info _
func (l *zapLog) Info(args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.InfoLevel) {
		l.logger.Info(fmt.Sprint(args...))
	}
}

// Infof _
func (l *zapLog) Infof(format string, args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.InfoLevel) {
		l.logger.Info(fmt.Sprintf(format, args...))
	}
}

// Warn _
func (l *zapLog) Warn(args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.WarnLevel) {
		l.logger.Warn(fmt.Sprint(args...))
	}
}

// Warnf _
func (l *zapLog) Warnf(format string, args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.WarnLevel) {
		l.logger.Warn(fmt.Sprintf(format, args...))
	}
}

// Error _
func (l *zapLog) Error(args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.ErrorLevel) {
		l.logger.Error(fmt.Sprint(args...))
	}
}

// Errorf _
func (l *zapLog) Errorf(format string, args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.ErrorLevel) {
		l.logger.Error(fmt.Sprintf(format, args...))
	}
}

// Fatal _
func (l *zapLog) Fatal(args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.FatalLevel) {
		l.logger.Fatal(fmt.Sprint(args...))
	}
}

// Fatalf _
func (l *zapLog) Fatalf(format string, args ...interface{}) {
	if l.logger.Core().Enabled(zapcore.FatalLevel) {
		l.logger.Fatal(fmt.Sprintf(format, args...))
	}
}

// Sync _
func (l *zapLog) Sync() error {
	return l.logger.Sync()
}

// SetLevel 设置输出端日志级别
func (l *zapLog) SetLevel(output string, level Level) {
	i, e := strconv.Atoi(output)
	if e != nil {
		return
	}
	if i < 0 || i >= len(l.levels) {
		return
	}
	l.levels[i].SetLevel(levelToZapLevel[level])
}

// GetLevel 获取输出端日志级别
func (l *zapLog) GetLevel(output string) Level {
	i, e := strconv.Atoi(output)
	if e != nil {
		return LevelDebug
	}
	if i < 0 || i >= len(l.levels) {
		return LevelDebug
	}
	return zapLevelToLevel[l.levels[i].Level()]
}

// ZapLogWrapper
type ZapLogWrapper struct {
	l *zapLog
}

// GetLogger 返回 zapLog
func (z *ZapLogWrapper) GetLogger() Logger {
	return z.l
}

// Trace _
func (z *ZapLogWrapper) Trace(args ...interface{}) {
	z.l.Trace(args...)
}

// Tracef _
func (z *ZapLogWrapper) Tracef(format string, args ...interface{}) {
	z.l.Tracef(format, args...)
}

// Debug _
func (z *ZapLogWrapper) Debug(args ...interface{}) {
	z.l.Debug(args...)
}

// Debugf _
func (z *ZapLogWrapper) Debugf(format string, args ...interface{}) {
	z.l.Debugf(format, args...)
}

// Info _
func (z *ZapLogWrapper) Info(args ...interface{}) {
	z.l.Info(args...)
}

// Infof _
func (z *ZapLogWrapper) Infof(format string, args ...interface{}) {
	z.l.Infof(format, args...)
}

// Warn _
func (z *ZapLogWrapper) Warn(args ...interface{}) {
	z.l.Warn(args...)
}

// Warnf _
func (z *ZapLogWrapper) Warnf(format string, args ...interface{}) {
	z.l.Warnf(format, args...)
}

// Error _
func (z *ZapLogWrapper) Error(args ...interface{}) {
	z.l.Error(args...)
}

// Errorf _
func (z *ZapLogWrapper) Errorf(format string, args ...interface{}) {
	z.l.Errorf(format, args...)
}

// Fatal _
func (z *ZapLogWrapper) Fatal(args ...interface{}) {
	z.l.Fatal(args...)
}

// Fatalf _
func (z *ZapLogWrapper) Fatalf(format string, args ...interface{}) {
	z.l.Fatalf(format, args...)
}

// Sync _
func (z *ZapLogWrapper) Sync() error {
	return z.l.Sync()
}

// SetLevel 设置输出端日志级别
func (z *ZapLogWrapper) SetLevel(output string, level Level) {
	z.l.SetLevel(output, level)
}

// GetLevel 获取输出端日志级别
func (z *ZapLogWrapper) GetLevel(output string) Level {
	return z.l.GetLevel(output)
}

// WithFields _
func (z *ZapLogWrapper) WithFields(fields ...string) Logger {
	return z.l.WithFields(fields...)
}