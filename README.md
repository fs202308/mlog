# mlog

log library with go, console、file、syslog3164(nlog) supported, this repo mainly depends on uber/zap

## Installation

`go get -u github.com/fs202308/mlog`

## Quick Start

```go
    f := `name: fapp
        mode: console
        level: debug
        file_path: "./log"
        file_max_size: 128
        file_retain_days: 30
        file_max_backups: 30
        file_compress: true

        nlog_tag:
          debug: 5401
          info: 5402
          warn: 5403
          error: 5404
          default: 5401
        nlog_remote_addr: 192.168.64.113:5211`
        
    var c mlog.LogConfig
    err := yaml.Unmarshal(f, &c)
    if err != nil {
        panic(err)
    }
    
    mlog.Init(&c)
    defer mlog.Sync()
    
    mlog.Debug("debug msg")
    mlog.Info("info msg")
    mlog.Warn("warn msg")
    mlog.Error("error msg")
```

Released under the [MIT License](LICENSE.txt).
