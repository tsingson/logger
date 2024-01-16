package main

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/tsingson/logger"
	"go.uber.org/zap/exp/zapslog"
)

func main() {
	// zapLogger, _ := zap.NewProduction()
	log := logger.New(logger.WithDebug())
	log.Info("hello world")

	defer log.Sync()

	myLog := slog.New(zapslog.NewHandler(log.Core(), &zapslog.HandlerOptions{}))
	myLog = myLog.
		With("environment", "dev").
		With("release", "v1.0.0")

	// log error
	myLog.
		With("category", "sql").
		With("query.statement", "SELECT COUNT(*) FROM users;").
		With("query.duration", 1*time.Second).
		With("error", fmt.Errorf("could not count users")).
		Error("caramba!")

	// log user signup
	myLog.
		With(
			slog.Group("user",
				slog.String("id", "user-123"),
				slog.Time("created_at", time.Now()),
			),
		).
		Info("user registration")
}
