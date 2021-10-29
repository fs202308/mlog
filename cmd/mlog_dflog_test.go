package main

import (
	"io/ioutil"
	"testing"

	"github.com/bsync-tech/mlog"
	"gopkg.in/yaml.v2"
)

func TestDflog(t *testing.T) {

	f, err := ioutil.ReadFile("./dflog_test.yaml")
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

	//json message support test
	mlog.Info("{k1:\"v1\", k2:'v2'}")

}
