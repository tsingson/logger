# **DO NOT USE THIS IN  YOUR BUSINESS PROJECT**
# a thin wraper of [uber-go/zap](https://github.com/uber-go/zap) logger for personal project 

[![GoDoc](https://godoc.org/github.com/tsingson/logger?status.svg)](https://godoc.org/github.com/tsingson/logger)[![Go Report Card](https://goreportcard.com/badge/github.com/tsingson/logger)](https://goreportcard.com/report/github.com/tsingson/logger)


## 0. thanks [uber-go/zap](https://github.com/uber-go/zap)
> Blazing fast, structured, leveled logging in Go.
>
> ---- [uber-go/zap](https://github.com/uber-go/zap)

this thin wrapper of [uber-go/zap](https://github.com/uber-go/zap)

[chinese readme here 中文说明](./README_cn.md)


## 1. feature

* Import from github.com/rs/zerolog/diode Support many-to-one caching to enhance log file write performance
* Add gopkg.in/natefinch/lumberjack.v2 for split log file by file size
* Add github.com/lestrrat-go/file-rotatelogs for split log file  every day
* Support github.com/jackc/pgx/v4 log 
* Support fasthttp log 
* setup default log storage path, log file name....
* debug option that logs to STDOUT 
* add global log level setting


## 2.  example
 in linux os macOS

 go source code in ./example-test/main.go

```
package main

import (
    "go.uber.org/zap"

    "github.com/tsingson/logger"
)

func main() {

		log := logger.New(logger.WithStoreInDay(),
		logger.WithDebug(),
		logger.WithDays(31),
		logger.WithLevel(zapcore.InfoLevel))
		
    defer log.Sync()

    logger.SetLevel(zap.DebugLevel)
    log.Info("info logging enabled 1")
    log.Debug("------------------------------------ 2")
    log.Warn(`logger.SetLevel(zap.ErrorLevel) 3`)
    log.Info("info logging disabled aaaaaaaaaa 4")
    log.Debug("info logging disabled aaaaaaa 5")
    log.Error("info logging disabled aaaaaaaaa 6")
    log.Warn("info logging disabled aaaaaaaa 7")

    logger.SetLevel(zap.ErrorLevel)
    log.Info("info logging disabled 8")
    log.Debug("info logging disabled 9")
    log.Error("info logging disabled 10")
    log.Warn("info logging disabled 11")
    log.Error("------------------------------------ 12")
    log.Info(`	logger.SetLevel(zap.DebugLevel) 13`)

}

 

```

build and runing 

```
 /home/go/bin   ./example-test                      
2019-11-10T23:06:59.399+0800	INFO	logger@/zaploggerfunc.go:21	info logging enabled 1
2019-11-10T23:06:59.399+0800	DEBUG	logger@/zaploggerfunc.go:11	------------------------------------ 2
2019-11-10T23:06:59.399+0800	WARN	logger@/zaploggerfunc.go:31	logger.SetLevel(zap.ErrorLevel) 3
2019-11-10T23:06:59.399+0800	INFO	logger@/zaploggerfunc.go:21	info logging disabled aaaaaaaaaa 4
2019-11-10T23:06:59.399+0800	DEBUG	logger@/zaploggerfunc.go:11	info logging disabled aaaaaaa 5
2019-11-10T23:06:59.399+0800	ERROR	logger@/zaploggerfunc.go:51	info logging disabled aaaaaaaaa 6
2019-11-10T23:06:59.399+0800	WARN	logger@/zaploggerfunc.go:31	info logging disabled aaaaaaaa 7
2019-11-10T23:06:59.399+0800	ERROR	logger@/zaploggerfunc.go:51	info logging disabled 10
2019-11-10T23:06:59.399+0800	ERROR	logger@/zaploggerfunc.go:51	------------------------------------ 12
/home/go/bin   cat ./log/example-test-2019-11-10-23.log 
{"level":"info","ts":"2019-11-10T23:06:59.399+0800","caller":"logger@/zaploggerfunc.go:21","msg":"info logging enabled 1"}
{"level":"debug","ts":"2019-11-10T23:06:59.399+0800","caller":"logger@/zaploggerfunc.go:11","msg":"------------------------------------ 2"}
{"level":"warn","ts":"2019-11-10T23:06:59.399+0800","caller":"logger@/zaploggerfunc.go:31","msg":"logger.SetLevel(zap.ErrorLevel) 3"}
{"level":"info","ts":"2019-11-10T23:06:59.399+0800","caller":"logger@/zaploggerfunc.go:21","msg":"info logging disabled aaaaaaaaaa 4"}
{"level":"debug","ts":"2019-11-10T23:06:59.399+0800","caller":"logger@/zaploggerfunc.go:11","msg":"info logging disabled aaaaaaa 5"}
{"level":"error","ts":"2019-11-10T23:06:59.399+0800","caller":"logger@/zaploggerfunc.go:51","msg":"info logging disabled aaaaaaaaa 6"}
{"level":"warn","ts":"2019-11-10T23:06:59.399+0800","caller":"logger@/zaploggerfunc.go:31","msg":"info logging disabled aaaaaaaa 7"}
{"level":"error","ts":"2019-11-10T23:06:59.399+0800","caller":"logger@/zaploggerfunc.go:51","msg":"info logging disabled 10"}
{"level":"error","ts":"2019-11-10T23:06:59.399+0800","caller":"logger@/zaploggerfunc.go:51","msg":"------------------------------------ 12"}

```


## 3. how to use it

 ### 3.1 go get and  import it

 this repo use go module

 get it
 
```
 go get github.com/tsingson/logger
```


 import in go code

```
 import "github.com/tsingson/logger"
```

logger struct  in ./zaplogger.go
```
// 
// ZapLogger a wrap of uber-go/zap
type ZapLogger struct {
	debug      bool    // debut is true, send log to  STDOUT
	storeInDay bool    // storeInDay is true,  save log file day by day 
	addCaller  bool    // 
	days     int64     // max storage days
	path       string   // the path to save log file
	prefix     string   // the prefix of log file name
	Log        *zap.Logger
	logLevel zapcore.LevelEnabler  // global log level setting, default is info level
}

// Logger  nick name
type Logger = ZapLogger

```

initial the logger 

simple :

```
    log := logger.New()
    defer log.Sync()
```

add options 
```
    log := logger.New(logger.WithDebug(true ), logger.WithAddCaller())
    defer log.Sync()
    
    logger.SetLevel(zap.DebugLevel) // change log level in runtime
```


## 4. limit
hard code for 3G mem for log writer, so , **DO NO USE IN YOUR BUSINESS PROJECT**


## 5. change log

1.  2019/10/28 move code from project as single repo
2.  2019/12/24 add log storage in every day and splite error log 
3.  2020/02/08 clean up for golangci-lint
4.  2024/03/09 adapter to log/slog
