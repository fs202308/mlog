package main

import (
	"io/ioutil"
	"github.com/bsync-tech/mlog"

	"gopkg.in/yaml.v2"
)

func main() {
	f, err := ioutil.ReadFile("./log.yaml")
	if err != nil {
		panic(err)
	}
	var c mlog.LogConfig
	err = yaml.Unmarshal(f, &c)
	if err != nil {
		panic(err)
	}
	mlog.Init(&c)
	defer mlog.Sync()
	mlog.Debug("debug msg")
	mlog.Info("info msg")
	mlog.Warn("warn msg")
	mlog.Error("error msg")
	// mlog.DPanic("dpanic msg")
	// mlog.Panic("panic msg")
	// mlog.Fatal("fatal msg")
}
