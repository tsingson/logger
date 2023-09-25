#   [uber-go/zap](go.uber.org/zap) 的一个简单封装, 用于个人项目 



## 0. 感谢 [uber-go/zap](https://github.com/uber-go/zap)

> Blazing fast, structured, leveled logging in Go.
>
> ---- [uber-go/zap](https://github.com/uber-go/zap)

这是 [uber-go/zap](https://github.com/uber-go/zap) 的一个简单封装, 用于个人项目

[english readme here](./README.md)

### 0.1 特色

* 导入   [github.com/rs/zerolog/diode](github.com/rs/zerolog/diode) 支持 many-to-one 缓存, 加强日志文件写入性能
* 支持  [gopkg.in/natefinch/lumberjack.v2](gopkg.in/natefinch/lumberjack.v2)  按日志文件大小进行存储拆分
* 支持 [github.com/lestrrat-go/file-rotatelogs](github.com/lestrrat-go/file-rotatelogs 按每一天进行存储拆分
* 支持 [github.com/jackc/pgx](github.com/jackc/pgx) 日志接口( interface )
* 支持 fasthttp 日志接口 ( interface )
* 添加默认的日志存储路径, 日志文件名, 简化配置
* 添加 debug 选项, 方便在开发时输出日志到 STDOUT
* 添加全局日志级别设置

## 1. 示例

以下示例在 macOS 上进行示范

 示例代码在 ./example-test/main.go

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

编译与运行

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


## 2. 如何使用

 ### 2.1  go 导入

 使用 go module 

 get it
 ```
 go get github.com/tsingson/logger
 ```


在代码中声明导入

 ```
 import "github.com/tsingson/logger"
 ```

logger struct  in ./zaplogger.go
```
// 
// ZapLogger a wrap of uber-go/zap
type ZapLogger struct {
	debug      bool // 输出日志到 STDOUT
	storeInDay bool // 是否将日志按天存储的标志位, 为 true 时按天拆分日志
	addCaller  bool //
	days     int64  // 日志文件保存天数
	path       string // 日志存储路径
	prefix     string // 日志文件名前缀
	Log        *zap.Logger
	logLevel zapcore.LevelEnabler // 全局日志级别设置项
}

// Logger  别名
type Logger = ZapLogger

```

initial the logger 

简单引用 :
```
    log := logger.New() // 定义日志
    defer log.Sync()   // 同步日志到文件
```

设置一些配置项 
```
 	log := logger.New(logger.WithDebug(), logger.WithAddCaller(), logger.WithStoreInDay()) // 定义
	defer log.Sync() // 同步日志到文件
	logger.SetLevel(zap.DebugLevel) // 运行时设置全局日志输出级别
```



## 3. 变更记录

1.  2019/10/28 从商用项目中拆分出基于 zap 的日志封装
2.  2019/12/24 增加每一天存储一个日志文件, 增加单独的错误日志文件