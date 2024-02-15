package main

import (
	"log/slog"

	"github.com/tsingson/logger"
	"go.uber.org/zap/exp/zapslog"
	"go.uber.org/zap/zapcore"
)

func main() {
	myLog := logger.New(
		logger.WithDebug(),
		logger.WithDays(31),
		logger.WithCaller(),
		logger.WithPrefix("test"),
		logger.WithLevel(zapcore.InfoLevel))
	defer myLog.Sync()

	//logger := zap.Must(zap.NewProduction())
	//
	//defer logger.Sync()

	sl := slog.New(zapslog.NewHandler(myLog.Core(), nil))

	sl.Info(
		"incoming request",
		slog.String("method", "GET"),
		slog.String("path", "/api/user"),
		slog.Int("status", 200),
	)
}
