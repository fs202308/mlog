package mlog

import (
	"net"
	"os"
	"time"

	"go.uber.org/zap/zapcore"
)

type NlogLogger struct {
	tag    string
	remote string
	conn   *net.UDPConn
}

func (n *NlogLogger) Write(p []byte) (int, error) {
	n.conn.Write([]byte(n.getNlogPrefix() + " [Debug] " + string(p)))
	return len(p), nil
}

func (n *NlogLogger) Init(c *LogConfig) zapcore.WriteSyncer {
	udpaddr, err := net.ResolveUDPAddr("udp4", c.NlogRemoteAddr)
	if err != nil {
		panic(err)
	}

	n.conn, err = net.DialUDP("udp", nil, udpaddr)
	if err != nil {
		panic(err)
	}

	return zapcore.AddSync(n)
}

func (n *NlogLogger) getNlogPrefix() string {
	now := time.Now().Format("2006-01-02 15:04:05.999")
	if n.tag != "" {
		hostname, _ := os.Hostname()
		pri := "<191>" + hostname + " " + n.tag + "[0]: "
		return pri + now
	}

	return now
}
