package mlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

type LogConfig struct {
	// AppName
	Name string `yaml:"name" json:"name"`
	// console file nlog
	Mode string `yaml:"mode" json:"mode"`
	// debug info warn error
	Level string `yaml:"level" json:"level"`
	// file path
	FilePath string `yaml:"file_path" json:"file_path"`
	// file max size
	FileMaxSize int `yaml:"file_max_size" json:"file_max_size"`
	// file retain days
	FileRetainDays int `yaml:"file_retain_days" json:"file_retain_days"`
	// file max backups
	FileMaxBackups int `yaml:"file_max_backups" json:"file_max_backups"`
	// file compress
	FileCompress bool `yaml:"file_compress" json:"file_compress"`

	// nlog tag
	NlogTag string `yaml:"nlog_tag" json:"nlog_tag"`
	// nlog remote addr tag
	NlogRemoteAddr string `yaml:"nlog_remote_addr" json:"nlog_remote_addr"`
}

type MLogger interface {
	Init(c *LogConfig) zapcore.WriteSyncer
}

func Init(c *LogConfig) {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "file",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		// EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	writers := []zapcore.WriteSyncer{}
	switch c.Mode {
	case "file":
		fileLogger := new(FileLogger)
		fileWriter := fileLogger.Init(c)
		writers = append(writers, fileWriter)
	case "console":
		consoleLogger := new(ConsoleLogger)
		fileWriter := consoleLogger.Init(c)
		writers = append(writers, fileWriter)
	case "nlog":
		nlogLogger := new(NlogLogger)
		nlogWriter := nlogLogger.Init(c)
		writers = append(writers, nlogWriter)
	}

	// 设置日志级别
	atomicLevel := GetAtomicLevel(c.Level)
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(writers...), atomicLevel)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	field := zap.Fields(zap.String("appName", c.Name))
	// 构造日志
	log = zap.New(core, caller, development, field)
	log.Info("mlog init success")
}

func GetLogger() *zap.Logger {
	return log
}

func Debug(msg string, fields ...zapcore.Field) {
	log.Debug(msg, fields...)
}

func Info(msg string, fields ...zapcore.Field) {
	log.Info(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	log.Warn(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	log.Error(msg, fields...)
}

func DPanic(msg string, fields ...zapcore.Field) {
	log.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zapcore.Field) {
	log.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zapcore.Field) {
	log.Fatal(msg, fields...)
}

func Sync() {
	log.Sync()
}
