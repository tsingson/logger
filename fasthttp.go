package logger

import (
	"fmt"

	"go.uber.org/zap"
)

// FromZap initial
func FromZap(log *zap.Logger) *ZapLogger {
	return &ZapLogger{
		Log: log,
	}
}

// Printf logs a message at level Info on the ZapLogger.
func (s *Logger) Printf(format string, args ...interface{}) {
	s.Log.Info(fmt.Sprintf(format, args...))
}

// Println logs a message at level Info on the ZapLogger.
func (s *Logger) Println(v ...interface{}) {
	s.Log.Info(fmt.Sprintln(v...))
}

/**
zapadapter logger for fasthttp
*/
