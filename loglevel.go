package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogLevel zap adapter log level
// var defaultLevel = zapadapter.NewAtomicLevel() // zapadapter.NewAtomicLevelAt(zapadapter.DebugLevel)
var _defaultLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)

// SetLevel set log level for zap adapter log
func SetLevel(level zapcore.Level) {
	_defaultLevel.SetLevel(level)
}

// SetLevel set log level for zap adapter log
func (s *Logger) SetLevel(level zapcore.Level) {
	_defaultLevel.SetLevel(level)
}
