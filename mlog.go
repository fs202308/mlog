package github.com/bsync-tech/mlog

import (
	"go.uber.org/zap"
)

func init() {

}

func GetAtomicLevel(level string) zap.AtomicLevel {
	switch level {
	case "debug", "DEBUG":
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info", "INFO":
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn", "WARN":
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error", "ERROR":
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "dpanic", "DPANIC":
		return zap.NewAtomicLevelAt(zap.DPanicLevel)
	case "panic", "PANIC":
		return zap.NewAtomicLevelAt(zap.PanicLevel)
	case "fatal", "FATAL":
		return zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	}
}
