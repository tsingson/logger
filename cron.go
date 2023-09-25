package logger

import (
	"fmt"

	"go.uber.org/zap"
)

//	type Logger interface {
//		// Info logs routine messages about cron's operation.
//		Info(msg string, keysAndValues ...interface{})
//		// Error logs an error condition.
//		Error(err error, msg string, keysAndValues ...interface{})
//	}
type CronLogger struct {
	log *Logger
}

func (c *CronLogger) Info(msg string, keysAndValues ...interface{}) {
	c.log.Info(msg, fmt.Sprint(keysAndValues...))
}

func (c *CronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	c.log.Error(msg, zap.Error(err), fmt.Sprint(keysAndValues...))
}

func (s *Logger) CronLogger() *CronLogger {
	return &CronLogger{
		log: s,
	}
}
