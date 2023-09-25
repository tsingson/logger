package logger

import (
	"fmt"

	"go.uber.org/zap"
)

type Sentinel struct {
	log *Logger
}

func (s *Logger) Sentinel() *Sentinel {
	return &Sentinel{
		log: s,
	}
}

func (s *Sentinel) Debug(msg string, args ...interface{}) {
	s.log.Log.Debug(msg + " " + fmt.Sprint(args...))
}

func (s *Sentinel) Info(msg string, args ...interface{}) {
	s.log.Log.Info(msg + " " + fmt.Sprint(args...))
}

func (s *Sentinel) Error(err error, msg string, args ...interface{}) {
	s.log.Log.Error(msg+" "+fmt.Sprint(args...), zap.Error(err))
}

func (s *Sentinel) Warn(msg string, args ...interface{}) {
	s.log.Log.Warn(msg + " " + fmt.Sprint(args...))
}

func (s *Sentinel) DebugEnabled() bool {
	return true
}

func (s *Sentinel) InfoEnabled() bool {
	return true
}

func (s *Sentinel) WarnEnabled() bool {
	return true
}

func (s *Sentinel) ErrorEnabled() bool {
	return true
}
