package logger

import (
	"fmt"

	"go.uber.org/zap"
)

// Debug logs a message at level DebugMode on the ZapLogger.
func (s *Logger) Debug(args ...interface{}) {
	s.Log.Debug(fmt.Sprint(args...))
}

// Debugf logs a message at level DebugMode on the ZapLogger.
func (s *Logger) Debugf(template string, args ...interface{}) {
	s.Log.Debug(fmt.Sprintf(template, args...))
}

// Info logs a message at level Info on the ZapLogger.
func (s *Logger) Info(args ...interface{}) {
	s.Log.Info(fmt.Sprint(args...))
}

// Infof logs a message at level Info on the ZapLogger.
func (s *Logger) Infof(template string, args ...interface{}) {
	s.Log.Info(fmt.Sprintf(template, args...))
}

// Warn logs a message at level Warn on the ZapLogger.
func (s *Logger) Warn(args ...interface{}) {
	s.Log.Warn(fmt.Sprint(args...))
}

// Warning logs a message at level Warn on the ZapLogger.
func (s *Logger) Warning(args ...interface{}) {
	s.Log.Warn(fmt.Sprint(args...))
}

// Warnf logs a message at level Warn on the ZapLogger.
func (s *Logger) Warnf(template string, args ...interface{}) {
	s.Log.Warn(fmt.Sprintf(template, args...))
}

// Warningf logs a message at level Warn on the ZapLogger.
func (s *Logger) Warningf(template string, args ...interface{}) {
	s.Log.Warn(fmt.Sprintf(template, args...))
}

// Error logs a message at level Error on the ZapLogger.
func (s *Logger) Error(args ...interface{}) {
	s.Log.Error(fmt.Sprint(args...))
}

// Errorf logs a message at level Warn on the ZapLogger.
func (s *Logger) Errorf(template string, args ...interface{}) {
	s.Log.Error(fmt.Sprintf(template, args...))
}

// Fatal logs a message at level Fatal on the ZapLogger.
func (s *Logger) Fatal(args ...interface{}) {
	s.Log.Fatal(fmt.Sprint(args...))
}

// Fatalf logs a message at level Warn on the ZapLogger.
func (s *Logger) Fatalf(template string, args ...interface{}) {
	s.Log.Fatal(fmt.Sprintf(template, args...))
}

// Panic logs a message at level Painc on the ZapLogger.
func (s *Logger) Panic(args ...interface{}) {
	s.Log.Panic(fmt.Sprint(args...))
}

// DPanic logs a message at level Painc on the ZapLogger.
func (s *Logger) DPanic(args ...interface{}) {
	s.Log.DPanic(fmt.Sprint(args...))
}

// Panicf logs a message at level Warn on the ZapLogger.
func (s *Logger) Panicf(template string, args ...interface{}) {
	s.Log.Panic(fmt.Sprintf(template, args...))
}

// Print logs a message at level Info on the ZapLogger.
func (s *Logger) Print(args ...interface{}) {
	s.Log.Info(fmt.Sprint(args...))
}

// With return a log with an extra field.
func (s *Logger) With(k string, v interface{}) *ZapLogger {
	s.Log.With(zap.Any(k, v))
	return s
}

// WithField return a log with an extra field.
func (s *Logger) WithField(k string, v interface{}) *ZapLogger {
	s.Log.With(zap.Any(k, v))
	return s
}

// WithFields return a log with extra fields.
func (s *Logger) WithFields(fields map[string]interface{}) *ZapLogger {
	clog := s
	i := 0
	for k, v := range fields {
		if i == 0 {
			clog = s.WithField(k, v)
		} else {
			clog = clog.WithField(k, v)
		}
		i++
	}
	return clog
}
