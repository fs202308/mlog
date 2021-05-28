package github.com/bsync-tech/mlog

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type FileLogger struct {
}

func (f *FileLogger) Init(c *LogConfig) zapcore.WriteSyncer {
	logger := lumberjack.Logger{
		Filename:   c.FilePath,       // 日志文件路径
		MaxSize:    c.FileMaxSize,    // 每个日志文件保存的大小 单位:M
		MaxAge:     c.FileRetainDays, // 文件最多保存多少天
		MaxBackups: c.FileMaxBackups, // 日志文件最多保存多少个备份
		Compress:   c.FileCompress,   // 是否压缩
	}

	writer := zapcore.Lock(zapcore.AddSync(&logger))

	return writer
}
