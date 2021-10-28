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

	// nlog tags
	NlogTags map[string]string `yaml:"nlog_tag" json:"nlog_tag"`
	// nlog remote addr
	NlogRemoteAddr string `yaml:"nlog_remote_addr" json:"nlog_remote_addr"`

	DflogSubsys string `yaml:"dflog_subsys" json:"dflog_subsys"`

	DflogModule string `yaml:"dflog_module" json:"dflog_module"`

	DflogRemoteAddr string `yaml:"dflog_remote_addr" json:"dflog_remote_addr"`
}

type MLogger interface {
	Init(c *LogConfig) zapcore.WriteSyncer
}

func Init(c *LogConfig) {

	// 开启开发模式，堆栈跟踪
	// caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置日志级别
	atomicLevel := GetAtomicLevel(c.Level)

	var writers []zapcore.WriteSyncer
	switch c.Mode {
	case "file":
		fileLogger := new(FileLogger)
		fileWriter := fileLogger.Init(c)
		writers = append(writers, fileWriter)
		coren := zapcore.NewCore(zapcore.NewJSONEncoder(normalEncodeConfig()), zapcore.NewMultiWriteSyncer(writers...), atomicLevel)
		log = zap.New(coren, development, getFields(c))

	case "console":
		consoleLogger := new(ConsoleLogger)
		fileWriter := consoleLogger.Init(c)
		writers = append(writers, fileWriter)
		coren := zapcore.NewCore(zapcore.NewJSONEncoder(normalEncodeConfig()), zapcore.NewMultiWriteSyncer(writers...), atomicLevel)
		log = zap.New(coren, development, getFields(c))

	case "nlog":
		cores := make([]zapcore.Core, 0)
		for k, v := range c.NlogTags {
			nlogLogger := new(NlogLogger)
			nlogLogger.tag = v
			nlogLogger.level = k
			nlogWriter := nlogLogger.Init(c)
			cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(normalEncodeConfig()), nlogWriter, GetAtomicLevelEnableFuncEqual(k)))
		}
		coresn := zapcore.NewTee(cores...)
		log = zap.New(coresn, development, getFields(c))

	case "dflog":
		dflogLogger := new(DflogLogger)
		dflogWriter := dflogLogger.Init(c)

		core := zapcore.NewCore(zapcore.NewJSONEncoder(dflogLogger.getEncoderConfig()), dflogWriter, atomicLevel)
		log = zap.New(core, development, dflogLogger.getFields())

	}
	log.Info("mlog init success")
}

//support normal EncoderConfig for file、 console、 nlog
func normalEncodeConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "file",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		// EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
		EncodeName: zapcore.FullNameEncoder,
	}
}

func getFields(c *LogConfig) zap.Option {
	return zap.Fields(zap.String("appName", c.Name))
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
