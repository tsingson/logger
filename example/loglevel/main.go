package main

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/tsingson/logger"
)

func main() {
	myLog := logger.New(
		logger.WithDebug(),
		logger.WithDays(31),
		logger.WithCaller(),
		logger.WithPrefix("test"),
		logger.WithLevel(zapcore.InfoLevel))
	defer myLog.Sync()
	// myLog.SetLevel(zap.DebugLevel)
	myLog.Info("info logging enabled 1")
	myLog.Debug("------------------------------------ 2")
	myLog.Warn(`logger.SetLevel(zapadapter.ErrorLevel) 3`)
	myLog.Info("info logging disabled aaaaaaaaaa 4")
	myLog.Debug("info logging disabled aaaaaaa 5")
	myLog.Error("info logging disabled aaaaaaaaa 6")
	// myLog.Fatal("info logging disabled aaaaaaaaaaaaaa")
	myLog.Info("info logging disabled 8")
	myLog.Debug("info logging disabled 9")
	myLog.Error("info logging disabled 10")
	myLog.DPanic("info logging disabled 11")
	// myLog.Fatal("info logging disabled")
	logger.SetLevel(zap.InfoLevel)
	myLog.Error("------------------------------------ 12")
	myLog.Info(`	logger.SetLevel(zapadapter.DebugLevel) 13`)
	log := myLog.Named("name logged")
	log.Error(`	logger.SetLevel(zapadapter.DebugLevel) 13`)
	log.Info(`ipipnet`, zap.String("ipipnet", "127.0.0.3"), zap.String("sn", "xxxxxxxxxxxx"), zap.String("country", "China"))
	log.Debug(`ipipnet`, zap.String("ipipnet", "127.0.0.3"), zap.String("sn", "xxxxxxxxxxxx"), zap.String("country", "China"))
	log.Warn(`ipipnet`, zap.String("ipipnet", "127.0.0.3"), zap.String("sn", "xxxxxxxxxxxx"), zap.String("country", "China"))
	log.DPanic(`ipipnet`, zap.String("ipipnet", "127.0.0.3"), zap.String("sn", "xxxxxxxxxxxx"), zap.String("country", "China"))

	time.Sleep(time.Duration(1) * time.Second)
}
