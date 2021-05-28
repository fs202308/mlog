package mlog

import (
	"os"

	"go.uber.org/zap/zapcore"
)

type ConsoleLogger struct {
}

func (l *ConsoleLogger) Init(c *LogConfig) zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}
