package log

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"
)

// Config 日志配置 可设置多个输出方式
type Config []OutputConfig

// OutputConfig 日志输出方式配置
type OutputConfig struct {
	Writer       string       `yaml:"writer"`           // 日志输出端 (console, file)
	Formatter    string       `yaml:"formatter"`        // 日志输出格式 (console, json)
	Level        string       `yaml:"level"`            // 控制日志级别 debug info error
	WriteConfig  WriteConfig  `yaml:"writer_config"`    // 写文件配置
	FormatConfig FormatConfig `yaml:"formatter_config"` // 配置格式
	RemoteConfig yaml.Node    `yaml:"remote_config"`    // 远程日志配置
}

// WriteConfig 日志写文件配置
type WriteConfig struct {
	LogPath    string `yaml:"log_path"`    // 日志路径
	Filename   string `yaml:"filename"`    // 日志路径文件名，默认 app.log
	MaxAge     int    `yaml:"max_age"`     // 日志最大保留时间, 天
	MaxBackups int    `yaml:"max_backups"` // 日志最大文件数
	Compress   bool   `yaml:"compress"`    // 日志文件是否压缩
	MaxSize    int    `yaml:"max_size"`    // 日志文件最大大小（单位MB）
}

// FormatConfig 日志格式配置
type FormatConfig struct {
	TimeFmt       string `yaml:"time_fmt"`       // 日志输出时间格式默认为"2006-01-02 15:04:05.000"
	TimeKey       string `yaml:"time_key"`       // 日志输出时间key， 默认为"T"
	LevelKey      string `yaml:"level_key"`      // 日志级别输出key， 默认为"L"
	NameKey       string `yaml:"name_key"`       // 日志名称key， 默认为"N"
	CallerKey     string `yaml:"caller_key"`     // 日志输出调用者key， 默认"C"
	MessageKey    string `yaml:"message_key"`    // 日志输出消息体key，默认"M"
	StacktraceKey string `yaml:"stacktrace_key"` // 日志输出堆栈trace key， 默认"S"
}

func makeEncoder(c *OutputConfig) zapcore.Encoder {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        GetLogEncoderKey("Time", c.FormatConfig.TimeKey),
		LevelKey:       GetLogEncoderKey("Level", c.FormatConfig.LevelKey),
		NameKey:        GetLogEncoderKey("Name", c.FormatConfig.NameKey),
		CallerKey:      GetLogEncoderKey("Caller", c.FormatConfig.CallerKey),
		MessageKey:     GetLogEncoderKey("Message", c.FormatConfig.MessageKey),
		StacktraceKey:  GetLogEncoderKey("StackTrace", c.FormatConfig.StacktraceKey),
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     NewTimeEncoder(c.FormatConfig.TimeFmt),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	encoder := zapcore.NewConsoleEncoder(encoderCfg)
	switch c.Formatter {
	case "console":
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	default:
	}
	return encoder
}

// GetLogEncoderKey 获取用户自定义log输出字段名
func GetLogEncoderKey(defKey, key string) string {
	if key == "" {
		return defKey
	}
	return key
}
