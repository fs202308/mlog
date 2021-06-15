package mlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

func GetAtomicLevelEnableFuncEqual(level string) zap.LevelEnablerFunc {
	switch level {
	case "debug", "DEBUG":
		return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.DebugLevel
		})
	case "info", "INFO":
		return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.InfoLevel
		})
	case "warn", "WARN":
		return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.WarnLevel
		})
	case "error", "ERROR":
		return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.ErrorLevel
		})
	case "dpanic", "DPANIC":
		return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.DPanicLevel
		})
	case "panic", "PANIC":
		return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.PanicLevel
		})
	case "fatal", "FATAL":
		return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.FatalLevel
		})
	default:
		return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.DebugLevel
		})
	}
}
