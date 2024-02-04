package log

var (
	logAdapters = make(map[string]Logger)
	logWriters  = make(map[string]Writer)
)

var (
	DefaultLogger Logger
	defaultConfig = []OutputConfig{
		{
			Writer:    "console",
			Formatter: "console",
			Level:     "debug",
		},
	}
)

func init() {
	RegisterWriter(WriterConsole, NewConsoleWriter())
	RegisterWriter(WriterFile, NewFileWriter())
	DefaultLogger = NewZapLog(defaultConfig)
}

// Register 注册日志 Logger 实现
func Register(name string, logger Logger) {
	logAdapters[name] = logger
}

// GetLogger 获取一个 Logger 实现
func GetLogger(name string) Logger {
	return logAdapters[name]
}

// RegisterWriter 注册日志输出 Writer
func RegisterWriter(name string, writer Writer) {
	logWriters[name] = writer
}

// GetWriter 获取一个 Writer 输出实现
func GetWriter(name string) Writer {
	return logWriters[name]
}

// SetLogger 设置默认logger
func SetLogger(logger Logger) {
	DefaultLogger = logger
}

// Debug logs to DEBUG log.
func Debug(args ...interface{}) {
	DefaultLogger.Debug(args...)
}

// Debugf logs to DEBUG log.
func Debugf(format string, args ...interface{}) {
	DefaultLogger.Debugf(format, args...)
}

// Info logs to INFO log.
func Info(args ...interface{}) {
	DefaultLogger.Info(args...)
}

// Infof logs to INFO log.
func Infof(format string, args ...interface{}) {
	DefaultLogger.Infof(format, args...)
}

// Warn logs to WARNING log.
func Warn(args ...interface{}) {
	DefaultLogger.Warn(args...)
}

// Warnf logs to WARNING log.
func Warnf(format string, args ...interface{}) {
	DefaultLogger.Warnf(format, args...)
}

// Error logs to ERROR log.
func Error(args ...interface{}) {
	DefaultLogger.Error(args...)
}

// Errorf logs to ERROR log.
func Errorf(format string, args ...interface{}) {
	DefaultLogger.Errorf(format, args...)
}

// Fatal logs to Fatal log.
func Fatal(args ...interface{}) {
	DefaultLogger.Fatal(args...)
}

// Fatalf logs to Fatal log.
func Fatalf(format string, args ...interface{}) {
	DefaultLogger.Fatalf(format, args...)
}

// WithFields 设置一些自定义数据
func WithFields(fields ...string) Logger {
	return DefaultLogger.WithFields(fields...)
}
