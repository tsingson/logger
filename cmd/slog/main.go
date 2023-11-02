package main

import (
	logger2 "github.com/tsingson/logger"
	"go.uber.org/zap"
)

func main() {
	log := logger2.New(logger2.WithDays(31), logger2.WithDebug())
	defer func() {
		log.Sync()
	}()

	log.Info("test", zap.String("name", "tsingson"))

	//zapLogger := log.Named("slog")
	//sl := slog.New(slogzap.Option{Level: slog.LevelDebug, Logger: zapLogger}.NewZapHandler())
	//sl = sl.
	//	With("environment", "dev").
	//	With("release", "v1.0.0")

	// log error
	//sl.
	//	With("category", "sql").
	//	With("query.statement", "SELECT COUNT(*) FROM users;").
	//	With("query.duration", 1*time.Second).
	//	With("error", fmt.Errorf("could not count users")).
	//	Error("caramba!")

	//// log user signup
	//sl.
	//	With(
	//		slog.Group("user",
	//			slog.String("id", "user-123"),
	//			slog.Time("created_at", time.Now()),
	//		),
	//	).
	//	Info("user registration")
}
