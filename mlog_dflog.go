package mlog

import (
	"net"
	"strconv"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const messageKey string = "msg"
const levelKey string = "level"
const timeKey string = "timestamp"

type DflogLogger struct {
	subsys string
	module string
	remote string
	conn   *net.UDPConn
}

func (d *DflogLogger) Write(p []byte) (int, error) {
	d.conn.Write(p)

	return len(p), nil
}

func (d *DflogLogger) Init(c *LogConfig) zapcore.WriteSyncer {

	d.subsys = c.DflogSubsys
	d.module = c.DflogModule
	d.remote = c.DflogRemoteAddr

	udpAddr, err := net.ResolveUDPAddr("udp4", c.DflogRemoteAddr)
	if err != nil {
		panic(err)
	}

	d.conn, err = net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		panic(err)
	}

	return zapcore.AddSync(d)
}

func (d *DflogLogger) getFields() zap.Option {

	return zap.Fields(zap.String("subsys", d.subsys),
		zap.String("module", d.module))
}

func (d *DflogLogger) getEncoderConfig() zapcore.EncoderConfig {

	return zapcore.EncoderConfig{
		MessageKey:    messageKey,
		LevelKey:      levelKey,
		TimeKey:       timeKey,
		NameKey:       "logger",
		CallerKey:     "file",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(strconv.FormatInt(t.Unix(), 10))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		// EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName: zapcore.FullNameEncoder,
	}
}
